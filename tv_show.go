package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

func HandleTvShows() error {
	rootScanDir := Conf.RootScanDir
	rootMediaDir := Conf.RootMediaDir

	filesFoundInScan := make([]TvShowEpisodeFile, 0)

	err := ScanForTvShowFiles(rootScanDir, &filesFoundInScan)
	zap.S().Infof("found %v files in scan", len(filesFoundInScan))
	//zap.S().Debugf("filesFoundInScan:[%#v]", filesFoundInScan)
	if err != nil {
		zap.S().Info(err)
		return err
	}
	RemoveAlreadyHandledTvShows(&filesFoundInScan)
	zap.S().Infof("removed already processed & failed files, left %v files to process", len(filesFoundInScan))
	for i, v := range filesFoundInScan {

		parsedFilename, err := ParseTvShowFilename(v.FilenameOriginal)
		if err != nil {
			zap.S().Warnf("failed to parse filename index:[%#v] name:[%#v] with error:[%#v]\n", i, v.FilenameOriginal, err)
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			continue
		}

		show, err := searchTvShow(parsedFilename.Name)
		if err != nil {
			zap.S().Warnf("failed to fetch data for index:[%#v] object:[%#v] with error:[%#v]\n", i, v, err)
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			continue
		}

		if show.TotalResults > 0 {
			filesFoundInScan[i].TvShow = show.Results[0]
		}

		tvShowEpisodeDetails, err := getTvShowEpisodeDetails(fmt.Sprintf("%v", filesFoundInScan[i].TvShow.Id), parsedFilename.SeasonNumber, parsedFilename.EpisodeNumber)
		if err != nil {
			zap.S().Warnf("failed to get episode details, skipping [%s]", v.AbsolutePath)
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			continue
		}
		filesFoundInScan[i].TvShowEpisode = tvShowEpisodeDetails

		//if v.SuccessfulParseOriginal == false {
		//	zap.S().Warnf("entry was not processed successfully. skipping %s", v.FilenameOriginal)
		//	continue
		//}

		filename, err := MakeFilename(filesFoundInScan[i])
		if err != nil {
			return err
		}
		//episodeTitleFormat := fmt.Sprintf("%s S%02dE%02d - %s%s", filesFoundInScan[i].TvShow.Name, filesFoundInScan[i].TvShowEpisode.SeasonNumber, filesFoundInScan[i].TvShowEpisode.EpisodeNumber, filesFoundInScan[i].TvShowEpisode.Name, filepath.Ext(filesFoundInScan[i].FilenameOriginal))
		episodeTitleFormat := filename
		sanitised, err := SanitizeForWindowsPathOrFile(episodeTitleFormat)
		if err != nil {
			continue
		}

		filesFoundInScan[i].FilenameNew = sanitised
		filesFoundInScan[i].SuccessfulParseOriginal = true

	}

	for i, show := range filesFoundInScan {
		if show.SuccessfulParseOriginal == false {
			zap.S().Warnf("entry was not processed successfully. skipping %s \n", show.FilenameOriginal)
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			//fmt.Printf("entry was not processed successfully. skipping %s \n", show.FilenameOriginal)
			continue
		}

		dirPath := fmt.Sprintf("%s/%s/Season %02d", rootMediaDir, show.TvShow.Name, show.TvShowEpisode.SeasonNumber)
		destination := fmt.Sprintf("%s/%s/Season %02d/%s", rootMediaDir, show.TvShow.Name, show.TvShowEpisode.SeasonNumber, show.FilenameNew)

		//todo: find more generic way to handle dry runs
		if Conf.DryRun {
			return nil
		}

		err := os.MkdirAll(dirPath, 777)
		if err != nil {
			fmt.Println(err)
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			continue
		}

		fileExists := CheckIfDirOrFileExists(destination)
		if fileExists {
			zap.S().Infof("skipping file copy, already exists at destination: [%v]", destination)
			filesFoundInScan[i].SuccessfulCopyFile = true
			err = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			if err != nil {
				continue
			}
			continue
		}

		err = CopyFileToLocation(show.AbsolutePath, destination)
		if err != nil {
			fmt.Println(err)
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			continue
		}

		_, err = VerifyFilesizeOfPaths(show.AbsolutePath, destination)
		if err != nil {
			_ = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
			continue
		}
		filesFoundInScan[i].SuccessfulCopyFile = true

		err = SaveTvShowEpisodeFileToDb(filesFoundInScan[i])
		if err != nil {
			continue
		}
	}
	return nil
}
