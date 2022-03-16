package main

import (
	"go.uber.org/zap"
	"os"
	"reflect"
	"testing"
)

func TestInitBolt(t *testing.T) {
	_ = os.Mkdir(TmpDir, 0777)
	t.Cleanup(func() {
		_ = Close()
		_ = os.RemoveAll("./tmp/")
	})
	err := InitBolt("./tmp/test.db", []string{TvShowEpisodeFilesBucket})
	if err != nil {
		zap.S().Fatalf("failed to init db: error is: %v\n", err)
		return
	}
	created := CheckIfDirOrFileExists("./tmp/test.db")
	if !created {
		t.Errorf("failed to create the db file")
	}
}

func TestPutGetDelete(t *testing.T) {
	_ = os.Mkdir(TmpDir, 0777)
	t.Cleanup(func() {
		_ = Close()
		_ = os.RemoveAll("./tmp/")
	})
	_ = InitBolt("./tmp/test.db", []string{TvShowEpisodeFilesBucket})
	_ = Put([]byte(TvShowEpisodeFilesBucket), []byte("testKey"), []byte("testValue"))
	r := Get([]byte(TvShowEpisodeFilesBucket), []byte("testKey"))
	if string(r) != "testValue" {
		t.Errorf("expexted %s got %s", "testValue", string(r))
	}
	_ = Delete([]byte(TvShowEpisodeFilesBucket), []byte("testKey"))
	r = Get([]byte(TvShowEpisodeFilesBucket), []byte("testKey"))
	if r != nil {
		t.Errorf("expexted nil got %s", string(r))
	}
}

func TestGetAllKeysAndAllKeyValuePairs(t *testing.T) {
	_ = os.Mkdir(TmpDir, 0777)
	t.Cleanup(func() {
		_ = Close()
		_ = os.RemoveAll("./tmp/")
	})
	_ = InitBolt("./tmp/test.db", []string{TvShowEpisodeFilesBucket})
	r := GetAllKeys([]byte(TvShowEpisodeFilesBucket))
	if len(r) != 0 {
		t.Errorf("expexted 0 got %s", r)
	}
	_ = Put([]byte(TvShowEpisodeFilesBucket), []byte("testKey1"), []byte("testValue1"))
	_ = Put([]byte(TvShowEpisodeFilesBucket), []byte("testKey2"), []byte("testValue2"))
	_ = Put([]byte(TvShowEpisodeFilesBucket), []byte("testKey3"), []byte("testValue3"))
	r = GetAllKeys([]byte(TvShowEpisodeFilesBucket))
	if len(r) != 3 {
		t.Errorf("expexted 3 got %s", r)
	}
	x := GetAllKeyValues([]byte(TvShowEpisodeFilesBucket))
	expected := []BoltPair{
		{
			Key:   []byte("testKey1"),
			Value: []byte("testValue1"),
		}, {
			Key:   []byte("testKey2"),
			Value: []byte("testValue2"),
		}, {
			Key:   []byte("testKey3"),
			Value: []byte("testValue3"),
		},
	}
	if !reflect.DeepEqual(x, expected) {
		t.Errorf("expexted [%s] got [%s]", expected, x)
	}
}
