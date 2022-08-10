package main

import (
	"encoding/json"
	"fmt"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"io/fs"
	"path/filepath"
	"strings"
)

func ScanForMovieFiles(rootScanDir string, filesToHandle *[]model.MovieFile) error {
	err := filepath.WalkDir(rootScanDir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.EqualFold(d.Name(), gconst.WindowsGeneratedFileThumbsDb) {
			s := model.MovieFile{
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

func RemoveAlreadyHandledMovies(filesToHandle *[]model.MovieFile) {
	var x []model.MovieFile
	for _, file := range *filesToHandle {
		y := Get([]byte(gconst.MovieFilesBucket), []byte(file.AbsolutePath))
		if y == nil {
			x = append(x, file)
			continue
		}
		var s model.MovieFile
		err := json.Unmarshal(y, &s)
		if err != nil {
			zap.S().Warnf("failed to unmarshall: [%s]", y)
			continue
		}
		if s.SuccessfulCopyFile && s.SuccessfulParseOriginal {
			continue
		}
		if (!s.SuccessfulCopyFile || !s.SuccessfulParseOriginal) && !config.Conf.MovieFileRetryFailed {
			continue
		}
		x = append(x, file)
	}
	*filesToHandle = x
}

func SaveMovieFileToDb(f model.MovieFile) error {
	encoded, err := json.Marshal(f)
	if err != nil {
		zap.S().Warnf("failed to marshal show, cannot store it to db. struct is: %#v", f)
		return err
		//continue
	}
	err = Put([]byte(gconst.MovieFilesBucket), []byte(f.AbsolutePath), encoded)
	if err != nil {
		zap.S().Warnf("failed to save show, cannot store it to db. encoded is: %#v", encoded)
		return err
		//continue
	}
	return err
}
