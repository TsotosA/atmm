package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func CopyFileToLocation(srcFilePath, destFilePath string) error {
	//zap.S().Info("srcFilePath:", srcFilePath, "destFilePath:", destFilePath)
	zap.S().Infof("copying filepath [%v] to [%v]", srcFilePath, destFilePath)
	//check if dir exists and create it if it doesn't

	tmpFile, err := os.Create(fmt.Sprintf("%s.part", destFilePath)) // creates if file doesn't exist
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	bytesCopied, err := io.Copy(tmpFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		return err
	}
	zap.S().Infof("copied bytes: [%v]", bytesCopied)
	err = tmpFile.Sync()
	if err != nil {
		return err
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	err = os.Rename(fmt.Sprintf("%s.part", destFilePath), destFilePath)
	if err != nil {
		return err
	}
	return nil
}

func SanitizeForWindowsPathOrFile(s string) (string, error) {
	illegalWindowsChars := []string{"<", ">", ":", "\\", "/", "|", "?", "*"}
	for _, char := range illegalWindowsChars {
		if strings.Contains(s, char) {
			zap.S().Debug("found char [", char, "] at ", s)
			s = strings.ReplaceAll(s, char, "")
		}

	}
	return s, nil
}

func CheckIfDirOrFileExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func PrettyPrint(x interface{}) []byte {
	y, _ := json.MarshalIndent(x, "", "  ")
	if y != nil {
		fmt.Printf("%s:", reflect.TypeOf(x))
		fmt.Printf("%s\n", y)
	}
	return y
}

func VerifyFilesizeOfPaths(from, to string) (bool, error) {
	fromStat, err := os.Stat(from)
	if err != nil {
		return false, err
	}
	toStat, err := os.Stat(to)
	if err != nil {
		return false, err
	}
	result := fromStat.Size() == toStat.Size()
	zap.S().Infof("filepaths are the same size: %t\n", result)
	return result, nil
}

func CurrrentBinaryAbsolutePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return ex, nil
}

func IsPartialFile(p string) bool {
	return filepath.Ext(p) == PartialFileExtension
}
