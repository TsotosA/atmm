package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

func HandleMovies() error {
	rootMovieScanDir := Conf.RootMovieScanDir
	rootMovieMediaDir := Conf.RootMovieMediaDir
	filesFoundInScan := make([]MovieFile, 0)
	err := ScanForMovieFiles(rootMovieScanDir, &filesFoundInScan)
	zap.S().Infof("found %v files in scan", len(filesFoundInScan))
	if err != nil {
		zap.S().Info(err)
		return err
	}
	RemoveAlreadyHandledMovies(&filesFoundInScan)
	zap.S().Infof("removed already processed & failed files, left %v files to process", len(filesFoundInScan))
	for i, v := range filesFoundInScan {

		parsedFilename, err := ParseMovieFilename(v.FilenameOriginal)
		if err != nil {
			zap.S().Warnf("failed to parse filename index:[%#v] name:[%#v] with error:[%#v]\n", i, v.FilenameOriginal, err)
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}
		zap.S().Debugf("%#v", parsedFilename)

		filesFoundInScan[i].ParsedFilename = parsedFilename

		movie, err := searchMovie(parsedFilename.Title)
		if err != nil {
			zap.S().Warnf("failed to fetch data for index:[%#v] object:[%#v] with error:[%#v]\n", i, v, err)
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}

		if movie.TotalResults > 0 {
			filesFoundInScan[i].Movie = movie.Results[0]
		}

		filename, err := MakeMovieFilename(filesFoundInScan[i])
		if err != nil {
			return err
		}
		movieTitleFormat := filename
		sanitised, err := SanitizeForWindowsPathOrFile(movieTitleFormat)
		if err != nil {
			continue
		}

		filesFoundInScan[i].FilenameNew = sanitised
		filesFoundInScan[i].SuccessfulParseOriginal = true

	}
	for i, movie := range filesFoundInScan {
		if movie.SuccessfulParseOriginal == false {
			zap.S().Warnf("entry was not processed successfully. skipping %s \n", movie.FilenameOriginal)
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			//fmt.Printf("entry was not processed successfully. skipping %s \n", show.FilenameOriginal)
			continue
		}

		sanitisedTitle, err := SanitizeForWindowsPathOrFile(movie.Movie.Title)
		if err != nil {
			zap.S().Warnf("entry was not processed successfully. skipping %s \n", movie.FilenameOriginal)
			continue
		}

		//dirPath := fmt.Sprintf("%s/%s/Season %02d", rootMediaDir, show.TvShow.Name, show.TvShowEpisode.SeasonNumber)
		dirPath := fmt.Sprintf("%s/%s", rootMovieMediaDir, sanitisedTitle)
		destination := fmt.Sprintf("%s/%s/%s", rootMovieMediaDir, sanitisedTitle, movie.FilenameNew)
		//
		////todo: find more generic way to handle dry runs
		//if Conf.DryRun {
		//	return nil
		//}
		//
		err = os.MkdirAll(dirPath, 777)
		if err != nil {
			fmt.Println(err)
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}
		//
		fileExists := CheckIfDirOrFileExists(destination)
		if fileExists {
			zap.S().Infof("skipping file copy, already exists at destination: [%v]", destination)
			filesFoundInScan[i].SuccessfulCopyFile = true
			err = SaveMovieFileToDb(filesFoundInScan[i])
			if err != nil {
				continue
			}
			continue
		}
		//zap.S().Debugf("%t", fileExists)

		err = CopyFileToLocation(movie.AbsolutePath, destination)
		if err != nil {
			fmt.Println(err)
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}

		_, err = VerifyFilesizeOfPaths(movie.AbsolutePath, destination)
		if err != nil {
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}
		filesFoundInScan[i].SuccessfulCopyFile = true

		err = SaveMovieFileToDb(filesFoundInScan[i])
		if err != nil {
			continue
		}
	}

	return nil
}
