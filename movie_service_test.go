package main

import (
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"testing"
)

func TestScanForMovieFiles(t *testing.T) {
	_ = os.Mkdir(gconst.TmpDir, 0777)
	tmpDir1, _ := ioutil.TempDir(gconst.TmpDir, "testDir*")
	_ = os.WriteFile("./tmp/"+"testFile1", []byte("test input"), 0777)
	_ = os.WriteFile(tmpDir1+"/"+"testFile", []byte("test input"), 0777)
	t.Cleanup(func() {
		zap.S().Infof("cleanup function")
		_ = os.RemoveAll("./tmp/")
	})
	tmpArr := make([]model.MovieFile, 0)
	_ = ScanForMovieFiles(gconst.TmpDir, &tmpArr)
	if len(tmpArr) != 2 {
		t.Errorf("expeceted length 2, got [%d]", len(tmpArr))
	}
}
