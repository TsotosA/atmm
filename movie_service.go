package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"path/filepath"
	"strings"
)

func ScanForMovieFiles(rootScanDir string, filesToHandle *[]MovieFile) error {
	err := filepath.WalkDir(rootScanDir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.EqualFold(d.Name(), WindowsGeneratedFileThumbsDb) {
			s := MovieFile{
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

func RemoveAlreadyHandledMovies(filesToHandle *[]MovieFile) {
	var x []MovieFile
	for _, file := range *filesToHandle {
		y := Get([]byte(MovieFilesBucket), []byte(file.AbsolutePath))
		if y == nil {
			x = append(x, file)
			continue
		}
		var s MovieFile
		err := json.Unmarshal(y, &s)
		if err != nil {
			zap.S().Warnf("failed to unmarshall: [%s]", y)
			continue
		}
		if s.SuccessfulCopyFile && s.SuccessfulParseOriginal {
			continue
		}
		if (!s.SuccessfulCopyFile || !s.SuccessfulParseOriginal) && !Conf.MovieFileRetryFailed {
			continue
		}
		x = append(x, file)
	}
	*filesToHandle = x
}

func SaveMovieFileToDb(f MovieFile) error {
	encoded, err := json.Marshal(f)
	if err != nil {
		zap.S().Warnf("failed to marshal show, cannot store it to db. struct is: %#v", f)
		return err
		//continue
	}
	err = Put([]byte(MovieFilesBucket), []byte(f.AbsolutePath), encoded)
	if err != nil {
		zap.S().Warnf("failed to save show, cannot store it to db. encoded is: %#v", encoded)
		return err
		//continue
	}
	return err
}
