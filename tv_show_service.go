package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"path/filepath"
)

func ScanForTvShowFiles(rootScanDir string, filesToHandle *[]TvShowEpisodeFile) error {
	err := filepath.WalkDir(rootScanDir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			s := TvShowEpisodeFile{
				FilenameOriginal: d.Name(),
				AbsolutePath:     fmt.Sprintf("%v", path),
			}
			*filesToHandle = append(*filesToHandle, s)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func RemoveAlreadyHandledTvShows(filesToHandle *[]TvShowEpisodeFile) {
	var x []TvShowEpisodeFile
	for _, file := range *filesToHandle {
		y := Get([]byte(TvShowEpisodeFilesBucket), []byte(file.AbsolutePath))
		//zap.S().Infof("found in db: [%s]", y)
		if y == nil {
			//zap.S().Infof("not found in db")
			x = append(x, file)
			continue
		}
		var s TvShowEpisodeFile
		err := json.Unmarshal(y, &s)
		if err != nil {
			zap.S().Warnf("failed to unmarshall: [%s]", y)
			continue
		}
		//zap.S().Infof("%#v", s)
		if s.SuccessfulCopyFile && s.SuccessfulParseOriginal {
			//zap.S().Infof("remove from slice: [%s]", y)
			continue
		}
		if (!s.SuccessfulCopyFile || !s.SuccessfulParseOriginal) && !Conf.TvShowEpisodeFileRetryFailed {
			//zap.S().Infof("remove from slice: [%s]", y)
			continue
		}
		x = append(x, file)
	}
	*filesToHandle = x
}

func SaveTvShowEpisodeFileToDb(f TvShowEpisodeFile) error {
	encoded, err := json.Marshal(f)
	if err != nil {
		zap.S().Warnf("failed to marshal show, cannot store it to db. struct is: %#v", f)
		return err
		//continue
	}
	err = Put([]byte(TvShowEpisodeFilesBucket), []byte(f.AbsolutePath), encoded)
	if err != nil {
		zap.S().Warnf("failed to save show, cannot store it to db. encoded is: %#v", encoded)
		return err
		//continue
	}
	return err
}
