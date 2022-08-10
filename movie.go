package main

import (
	"fmt"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/helper"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"os"
	"time"
)

func HandleMovies() error {
	rootMovieScanDir := config.Conf.RootMovieScanDir
	rootMovieMediaDir := config.Conf.RootMovieMediaDir
	filesFoundInScan := make([]model.MovieFile, 0)
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
		sanitised, err := helper.SanitizeForWindowsPathOrFile(movieTitleFormat)
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

		sanitisedTitle, err := helper.SanitizeForWindowsPathOrFile(movie.Movie.Title)
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
		fileExists := helper.CheckIfDirOrFileExists(destination)
		if fileExists {
			zap.S().Infof("skipping file copy, already exists at destination: [%v]", destination)
			filesFoundInScan[i].SuccessfulCopyFile = true
			err = SaveMovieFileToDb(filesFoundInScan[i])
			if err != nil {
				continue
			}
			continue
		}

		isDone := helper.IsFileDoneBeingWritten(movie.AbsolutePath, 1*time.Second, gconst.Movie)
		zap.S().Infof("isDone: %t - movie: %s", isDone, movie.AbsolutePath)
		if !isDone {
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}

		err = helper.CopyFileToLocation(movie.AbsolutePath, destination, gconst.Movie)
		if err != nil {
			fmt.Println(err)
			_ = SaveMovieFileToDb(filesFoundInScan[i])
			continue
		}

		_, err = helper.VerifyFilesizeOfPaths(movie.AbsolutePath, destination)
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
