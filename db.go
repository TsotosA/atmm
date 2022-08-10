package main

import (
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/model"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

var database string
var fileMode os.FileMode = 0600 // owner can read and write
var db *bbolt.DB

// InitBolt inits bolt database. Create the file if not exist.
// By default, it opens the file in 0600 mode, with a 10 seconds timeout period
func InitBolt(path string, buckets []string) error {
	//log.Println("Trying to open database")
	database = path
	var err error
	// open the target file, file mode fileMode, and a 10 seconds timeout period
	db, err = bbolt.Open(database, fileMode, &bbolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		//log.Fatal(err)
		zap.S().Fatalf("failed to open db with error: %v", err)
	}
	zap.S().Infof("opened db")

	zap.S().Infof("Trying to create buckets")
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, value := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

// Close bolt db
func Close() error {
	err := db.Close()
	if err != nil {
		return err
	}
	zap.S().Infof("closed db")
	return nil
}

// Get value from bucket by key
func Get(bucket []byte, key []byte) []byte {
	var value []byte

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get(key)
		if v != nil {
			value = append(value, b.Get(key)...)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value
}

// Put a key/value pair into target bucket
func Put(bucket []byte, key []byte, value []byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})
	return err
}

// Delete a key from target bucket
func Delete(bucket []byte, key []byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(key)
		return err
	})

	return err
}

// GetAllKeys get all keys from the target bucket
func GetAllKeys(bucket []byte) [][]byte {
	var keys [][]byte

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			// Due to
			// Byte slices returned from Bolt are only valid during a transaction. Once the transaction has been committed or rolled back then the memory they point to can be reused by a new page or can be unmapped from virtual memory and you'll see an unexpected fault address panic when accessing it.
			// We copy the slice to retain it
			dst := make([]byte, len(k))
			copy(dst, k)

			keys = append(keys, dst)
			return nil
		})
		return nil
	})

	return keys
}

// BoltPair is a struct to store key/value pair data
type BoltPair struct {
	Key   []byte
	Value []byte
}

// GetAllKeyValues get all key/value pairs from a bucket in BoltPair struct format
func GetAllKeyValues(bucket []byte) []BoltPair {
	var pairs []BoltPair

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			// Due to
			// Byte slices returned from Bolt are only valid during a transaction. Once the transaction has been committed or rolled back then the memory they point to can be reused by a new page or can be unmapped from virtual memory and you'll see an unexpected fault address panic when accessing it.
			// We copy the slice to retain it
			dstk := make([]byte, len(k))
			dstv := make([]byte, len(v))
			copy(dstk, k)
			copy(dstv, v)

			pair := BoltPair{dstk, dstv}
			pairs = append(pairs, pair)
			return nil
		})

		return nil
	})

	return pairs
}

//func InitDb() (db *bbolt.DB, err error) {
//	database, err := bbolt.Open("atmm.db", 0600, &bbolt.Options{
//		Timeout: 1 * time.Second,
//	})
//	if err != nil {
//		log.Fatalf("%s", err)
//	}
//	return database, err
//}
//
//func CheckForBuckets(db bbolt.DB) error {
//	buckets := []string{"tv_show_episode_files"}
//	for _, bucket := range buckets {
//		err := db.Update(func(tx *bbolt.Tx) error {
//			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
//			if err != nil {
//				return fmt.Errorf("create bucket: %s", err)
//			}
//			return nil
//		})
//		if err != nil {
//			return fmt.Errorf("view buckets: %s", err)
//		}
//	}
//	zap.S().Infof("%#v", db.Stats())
//	return nil
//}

func CleanupBucket(bucket string, currentKeys []string) bool {
	dbKeys := GetAllKeys([]byte(bucket))
	existsInFS := false
	for _, dbKey := range dbKeys {
		for _, currentKey := range currentKeys {
			if currentKey == string(dbKey) {
				existsInFS = true
			}
		}
		if existsInFS == true {
			zap.S().Debugf("keep dbKey: %s", dbKey)
		} else {
			zap.S().Debugf("remove dbKey: %s", dbKey)
			err := Delete([]byte(bucket), dbKey)
			if err != nil {
				zap.S().Debugf("failed to remove dbKey: %s", dbKey)
				return false
			}
		}
		existsInFS = false
	}
	return true
}

func HandleDbCleanup() error {
	rootMovieScanDir := config.Conf.RootMovieScanDir
	movieFilesFoundInScan := make([]model.MovieFile, 0)
	err := ScanForMovieFiles(rootMovieScanDir, &movieFilesFoundInScan)
	zap.S().Infof("found %v files in scan", len(movieFilesFoundInScan))
	if err != nil {
		zap.S().Info(err)
		return err
	}
	var currentFsMoviePaths []string

	for _, file := range movieFilesFoundInScan {
		currentFsMoviePaths = append(currentFsMoviePaths, file.AbsolutePath)
	}
	CleanupBucket(gconst.MovieFilesBucket, currentFsMoviePaths)

	rootScanDir := config.Conf.RootScanDir
	tvShowFilesFoundInScan := make([]model.TvShowEpisodeFile, 0)
	err = ScanForTvShowFiles(rootScanDir, &tvShowFilesFoundInScan)
	zap.S().Infof("found %v files in scan", len(tvShowFilesFoundInScan))
	if err != nil {
		zap.S().Info(err)
		return err
	}
	var currentFsTvShowPaths []string

	for _, file := range tvShowFilesFoundInScan {
		currentFsTvShowPaths = append(currentFsTvShowPaths, file.AbsolutePath)
	}
	CleanupBucket(gconst.TvShowEpisodeFilesBucket, currentFsTvShowPaths)

	return nil
}
