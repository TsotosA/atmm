package model

import (
	"fmt"
	"github.com/tsotosa/atmm/gconst"
	"gopkg.in/yaml.v3"
)

type TvShowEpisodeFile struct {
	FilenameOriginal        string                         `json:"filenameOriginal,omitempty"`
	FilenameNew             string                         `json:"filenameNew,omitempty"`
	AbsolutePath            string                         `json:"absolutePath,omitempty"`
	TvShow                  TheMovieDbTvShow               `json:"tvShow"`
	TvShowEpisode           TheMovieDbTvShowEpisodeDetails `json:"tvShowEpisode"`
	SuccessfulParseOriginal bool                           `json:"successfulParseOriginal"`
	SuccessfulCopyFile      bool                           `json:"successfulCopyFile"`
	ParsedFilename          ParsedFilename                 `json:"parsedFilename"`
}

type ParsedFilename struct {
	Name          string `json:"name,omitempty"`
	Title         string `json:"title,omitempty"`
	SeasonNumber  string `json:"seasonNumber,omitempty"`
	EpisodeNumber string `json:"episodeNumber,omitempty"`
	Year          string `json:"year,omitempty"`
	Resolution    string `json:"resolution,omitempty"`
}

type AppConf struct {
	RootScanDir                  string `yaml:"rootScanDir" json:"rootScanDir"`
	RootMediaDir                 string `yaml:"rootMediaDir" json:"rootMediaDir"`
	TheMovieDbBaseApiUrlV3       string `yaml:"theMovieDbBaseApiUrlV3" json:"theMovieDbBaseApiUrlV3"`
	ApiKey                       string `yaml:"apiKey" json:"apiKey"`
	DryRun                       bool   `yaml:"dryRun" json:"dryRun"`
	LogOutputPath                string `yaml:"logOutputPath" json:"logOutputPath"`
	Environment                  string `yaml:"environment" json:"environment"`
	LogLevel                     string `yaml:"logLevel" json:"logLevel"`
	Cron                         string `yaml:"cron" json:"cron"`
	TvShowEpisodeFormat          string `yaml:"tvShowEpisodeFormat" json:"tvShowEpisodeFormat"`
	TvShowEpisodeFileRetryFailed bool   `yaml:"tvShowEpisodeFileRetryFailed" json:"tvShowEpisodeFileRetryFailed"`
	CheckForUpdatesInterval      string `yaml:"checkForUpdatesInterval" json:"checkForUpdatesInterval"`
	GithubPersonalToken          string `yaml:"githubPersonalToken" json:"githubPersonalToken"`
	GithubUsername               string `yaml:"githubUsername" json:"githubUsername"`
	IsAutoRestartManaged         bool   `yaml:"isAutoRestartManaged" json:"isAutoRestartManaged"`
	ScanForMovieInterval         string `yaml:"scanForMovieInterval" json:"scanForMovieInterval"`
	RootMovieScanDir             string `yaml:"rootMovieScanDir" json:"rootMovieScanDir"`
	RootMovieMediaDir            string `yaml:"rootMovieMediaDir" json:"rootMovieMediaDir"`
	MovieFileRetryFailed         bool   `yaml:"movieFileRetryFailed" json:"movieFileRetryFailed"`
	MovieCustomFormat            string `yaml:"movieCustomFormat" json:"movieCustomFormat"`
	DbBucketsCleanupInterval     string `yaml:"dbBucketsCleanupInterval" json:"dbBucketsCleanupInterval"`
	LogRotateMaxNumOfBackups     int    `yaml:"logRotateMaxNumOfBackups" json:"logRotateMaxNumOfBackups"`
	LogRotateMaxAgeOfBackups     int    `yaml:"logRotateMaxAgeOfBackups" json:"logRotateMaxAgeOfBackups"`
	LogRotateMaxLogFileSize      int    `yaml:"logRotateMaxLogFileSize" json:"logRotateMaxLogFileSize"`
	LogRotateCompressBackups     bool   `yaml:"logRotateCompressBackups" json:"logRotateCompressBackups"`
	ApiPort                      int    `yaml:"apiPort" json:"apiPort"`
	UiPort                       int    `yaml:"uiPort" json:"uiPort"`
}

type AppConfUpdate struct {
	RootScanDir                  interface{} `yaml:"rootScanDir" json:"rootScanDir,omitempty"`
	RootMediaDir                 interface{} `yaml:"rootMediaDir" json:"rootMediaDir,omitempty"`
	TheMovieDbBaseApiUrlV3       interface{} `yaml:"theMovieDbBaseApiUrlV3" json:"theMovieDbBaseApiUrlV3,omitempty"`
	ApiKey                       interface{} `yaml:"apiKey" json:"apiKey,omitempty"`
	DryRun                       interface{} `yaml:"dryRun" json:"dryRun,omitempty"`
	LogOutputPath                interface{} `yaml:"logOutputPath" json:"logOutputPath,omitempty"`
	Environment                  interface{} `yaml:"environment" json:"environment,omitempty"`
	LogLevel                     interface{} `yaml:"logLevel" json:"logLevel,omitempty"`
	Cron                         interface{} `yaml:"cron" json:"cron,omitempty"`
	TvShowEpisodeFormat          interface{} `yaml:"tvShowEpisodeFormat" json:"tvShowEpisodeFormat,omitempty"`
	TvShowEpisodeFileRetryFailed interface{} `yaml:"tvShowEpisodeFileRetryFailed" json:"tvShowEpisodeFileRetryFailed,omitempty"`
	CheckForUpdatesInterval      interface{} `yaml:"checkForUpdatesInterval" json:"checkForUpdatesInterval,omitempty"`
	GithubPersonalToken          interface{} `yaml:"githubPersonalToken" json:"githubPersonalToken,omitempty"`
	GithubUsername               interface{} `yaml:"githubUsername" json:"githubUsername,omitempty"`
	IsAutoRestartManaged         interface{} `yaml:"isAutoRestartManaged" json:"isAutoRestartManaged,omitempty"`
	ScanForMovieInterval         interface{} `yaml:"scanForMovieInterval" json:"scanForMovieInterval,omitempty"`
	RootMovieScanDir             interface{} `yaml:"rootMovieScanDir" json:"rootMovieScanDir,omitempty"`
	RootMovieMediaDir            interface{} `yaml:"rootMovieMediaDir" json:"rootMovieMediaDir,omitempty"`
	MovieFileRetryFailed         interface{} `yaml:"movieFileRetryFailed" json:"movieFileRetryFailed,omitempty"`
	MovieCustomFormat            interface{} `yaml:"movieCustomFormat" json:"movieCustomFormat,omitempty"`
	DbBucketsCleanupInterval     interface{} `yaml:"dbBucketsCleanupInterval" json:"dbBucketsCleanupInterval,omitempty"`
	LogRotateMaxNumOfBackups     interface{} `yaml:"logRotateMaxNumOfBackups" json:"logRotateMaxNumOfBackups,omitempty"`
	LogRotateMaxAgeOfBackups     interface{} `yaml:"logRotateMaxAgeOfBackups" json:"logRotateMaxAgeOfBackups,omitempty"`
	LogRotateMaxLogFileSize      interface{} `yaml:"logRotateMaxLogFileSize" json:"logRotateMaxLogFileSize,omitempty"`
	LogRotateCompressBackups     interface{} `yaml:"logRotateCompressBackups" json:"logRotateCompressBackups,omitempty"`
	ApiPort                      interface{} `yaml:"apiPort" json:"apiPort,omitempty"`
	UiPort                       interface{} `yaml:"uiPort" json:"uiPort,omitempty"`
}

func (receiver AppConf) Mask() AppConf {
	if receiver.GithubPersonalToken != "" {
		receiver.GithubPersonalToken = "******"
	}
	if receiver.ApiKey != "" {
		receiver.ApiKey = "******"
	}
	if receiver.GithubUsername != "" {
		receiver.GithubUsername = "******"
	}
	return receiver
}

func (receiver AppConf) UpdateFields(update *AppConfUpdate) {
	var n yaml.Node

	y := []byte(gconst.ExampleConfigYaml)
	err := yaml.Unmarshal(y, &n)
	//fmt.Printf("----- r %+v", n)
	d, err := yaml.Marshal(&n)

	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", d)
	//if update.LogRotateMaxNumOfBackups != nil {
	//	viper.Set("LogRotateMaxNumOfBackups", update.LogRotateMaxNumOfBackups)
	//	fmt.Printf("%+v", receiver)
	//}
}

type FilenameFormatPair struct {
	StartIndex    int
	EndIndex      int
	PropertyName  string
	PropertyValue string
}

type CronJob struct {
	ScheduleToUse string
	MethodToRun   func()
}

type MovieFile struct {
	FilenameOriginal        string          `json:"filenameOriginal,omitempty"`
	FilenameNew             string          `json:"filenameNew,omitempty"`
	AbsolutePath            string          `json:"absolutePath,omitempty"`
	Movie                   TheMovieDbMovie `json:"movie,omitempty"`
	SuccessfulParseOriginal bool            `json:"successfulParseOriginal,omitempty"`
	SuccessfulCopyFile      bool            `json:"successfulCopyFile,omitempty"`
	ParsedFilename          ParsedFilename  `json:"parsedFilename,omitempty"`
}
