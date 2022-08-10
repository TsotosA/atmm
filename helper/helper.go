package helper

import (
	"encoding/json"
	"fmt"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/global"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func CopyFileToLocation(srcFilePath, destFilePath string, t string) error {
	//zap.S().Info("srcFilePath:", srcFilePath, "destFilePath:", destFilePath)
	zap.S().Infof("copying filepath [%v] to [%v]", srcFilePath, destFilePath)
	if t == gconst.TvShow {
		global.WaitingSeriesToFinishCopying = true
		defer func() {
			global.WaitingSeriesToFinishCopying = false
		}()
	}
	if t == gconst.Movie {
		global.WaitingMoviesToFinishCopying = true
		defer func() {
			global.WaitingMoviesToFinishCopying = false
		}()
	}
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
	return filepath.Ext(p) == gconst.PartialFileExtension
}

func IsFileDoneBeingWritten(path string, sleepFor time.Duration, t string) bool {
	if t == gconst.TvShow {
		global.WaitingSeriesToFinishCopying = true
		defer func() {
			global.WaitingSeriesToFinishCopying = false
		}()
	}
	if t == gconst.Movie {
		global.WaitingMoviesToFinishCopying = true
		defer func() {
			global.WaitingMoviesToFinishCopying = false
		}()
	}
	response := false
	stat1, err := os.Stat(path)
	time.Sleep(sleepFor)
	stat2, err := os.Stat(path)
	if err != nil {
		return false
	}
	response = stat1.ModTime() == stat2.ModTime()
	return response
}

func GetLastNLinesWithSeek(filepath string, nLines int) string {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		//todo: handle this?
		panic("Cannot open file")
		os.Exit(1)
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	lineBreaksFound := 0
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			if lineBreaksFound > nLines {
				break
			}
			lineBreaksFound++
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}
