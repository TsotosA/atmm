package main

import (
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"testing"
)

func TestCopyFileToLocation(t *testing.T) {
	_ = os.Mkdir(TmpDir, 0777)
	tmpDir1, _ := ioutil.TempDir(TmpDir, "testDir*")
	tmpDir2, _ := ioutil.TempDir(TmpDir, "testDir*")
	_ = os.WriteFile(tmpDir1+"/"+"testFile", []byte("test input"), 0777)
	t.Cleanup(func() {
		zap.S().Infof("cleanup function")
		_ = os.RemoveAll("./tmp/")
	})
	zap.S().Infof("tmpDir1: %s, tmpDir2: %s", tmpDir1, tmpDir2)
	_ = CopyFileToLocation(tmpDir1+"/"+"testFile", tmpDir2+"/"+"testFile")
	created := CheckIfDirOrFileExists(tmpDir2 + "/" + "testFile")
	if !created {
		t.Errorf("failed to copy file to correct location")
	}
}

func TestSanitizeForWindowsPathOrFile(t *testing.T) {
	tests := []struct {
		a        string
		expected string
		error    error
	}{
		{
			a:        "a?b",
			expected: "ab",
			error:    nil,
		},
	}
	for _, tt := range tests {
		testHName := fmt.Sprintf("a,expected,error:[%v],[%v],[%v]", tt.a, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := SanitizeForWindowsPathOrFile(tt.a)
			if r != tt.expected {
				if err != nil && tt.error != nil {
					t.Errorf("got [%v],wanted [%#v] [%v]", r, tt.expected, err)
				}
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}

func TestCheckIfDirOrFileExists(t *testing.T) {
	tests := []struct {
		a        string
		expected bool
	}{
		{
			a:        "./helper_test.go",
			expected: true,
		},
		{
			a:        "./whatever.go",
			expected: false,
		},
	}
	for _, tt := range tests {
		testHName := fmt.Sprintf("a,expected:[%v],[%v]", tt.a, tt.expected)
		t.Run(testHName, func(t *testing.T) {
			r := CheckIfDirOrFileExists(tt.a)
			if r != tt.expected {
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}

func TestVerifyFilesizeOfPaths(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
		error    error
	}{
		{
			a:        "./helper.go",
			b:        "./helper.go",
			expected: true,
			error:    nil,
		},
		{
			a:        "./helper.go",
			b:        "./helper_test.go",
			expected: false,
			error:    nil,
		},
	}
	for _, tt := range tests {
		testHName := fmt.Sprintf("a,expected,error:[%v],[%v],[%v]", tt.a, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := VerifyFilesizeOfPaths(tt.a, tt.b)
			if r != tt.expected {
				if err != nil && tt.error != nil {
					t.Errorf("got [%v],wanted [%#v] [%v]", r, tt.expected, err)
				}
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}
