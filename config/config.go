package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"github.com/tsotosa/atmm/cronjob"
	"github.com/tsotosa/atmm/debounce"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/helper"
	"github.com/tsotosa/atmm/model"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	Conf             model.AppConf
	AllConfYamlProps = []string{
		"rootScanDir",
		"rootMediaDir",
		"theMovieDbBaseApiUrlV3",
		"apiKey",
		"dryRun",
		"logOutputPath",
		"environment",
		"logLevel",
		"cron",
		"tvShowEpisodeFormat",
		"tvShowEpisodeFileRetryFailed",
		"checkForUpdatesInterval",
		"githubPersonalToken",
		"githubUsername",
		"isAutoRestartManaged",
		"scanForMovieInterval",
		"rootMovieScanDir",
		"rootMovieMediaDir",
		"movieFileRetryFailed",
		"movieCustomFormat",
		"dbBucketsCleanupInterval",
		"logRotateMaxNumOfBackups",
		"logRotateMaxAgeOfBackups",
		"logRotateMaxLogFileSize",
		"logRotateCompressBackups",
		"apiPort",
		"uiPort",
	}
)

func ConfigInit(c *cron.Cron, mwg *sync.WaitGroup, jobs []model.CronJob) {
	//for k, v := range defaults {
	//	viper.SetDefault(string(k), v)
	//}

	b := helper.CheckIfDirOrFileExists("./config.yaml")
	if b {
		missingProps, err := checkIfAllConfigPropsExistInConfFile()
		if err != nil {
			fmt.Printf("failed to check config for all props existance: %s\n", err)
		}
		if len(missingProps) > 0 {
			panic(fmt.Errorf("missingProps from config file: %v\n", missingProps))
		}
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//err := viper.WriteConfigAs("./example.config.yaml")
			err := os.WriteFile("./example.config.yaml", []byte(gconst.ExampleConfigYaml), 0777)
			if err != nil {
				fmt.Printf("failed to write template config: %s", err)
				panic("config file not found")
			}
			panic("README ==> could not locate configuration file. wrote a new one with placeholder values (example.config.yaml), please fill them in appropriately and copy the file to config.yaml <=== README")
			// AppConf file not found; ignore error if desired
		} else {
			panic(fmt.Errorf("fatal err config file: %s", err))
			// AppConf file was found but another error was produced
		}
	}

	err := viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("could not decode config to struct: %v", err)
	}

	viper.WatchConfig()
	debounced := debounce.New(time.Second * 2)
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		log.Printf("updated configuration file")
		debounced(func() {
			cronjob.RestartCronJobs(c, mwg, jobs)
		})
		if err != nil {
			log.Fatalf("config changed, could not decode to struct: %v", err)
		}
	})
}

func checkIfAllConfigPropsExistInConfFile() (missingProps []string, err error) {
	//r := `^(.*):(.*)[\n]?`
	res := []string{}
	fileData, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Printf("failed to read config: %s", err)
		return res, err
	}
	for _, prop := range AllConfYamlProps {
		if !strings.Contains(string(fileData), prop) {
			res = append(res, prop)
		}
	}
	return res, nil
}
