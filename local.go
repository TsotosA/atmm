package main

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
	RootScanDir                  string `yaml:"rootScanDir"`
	RootMediaDir                 string `yaml:"rootMediaDir"`
	TheMovieDbBaseApiUrlV3       string `yaml:"theMovieDbBaseApiUrlV3"`
	ApiKey                       string `yaml:"apiKey"`
	DryRun                       bool   `yaml:"dryRun"`
	LogOutputPath                string `yaml:"logOutputPath"`
	Environment                  string `yaml:"environment"`
	LogLevel                     string `yaml:"logLevel"`
	Cron                         string `yaml:"cron"`
	TvShowEpisodeFormat          string `yaml:"tvShowEpisodeFormat"`
	TvShowEpisodeFileRetryFailed bool   `yaml:"tvShowEpisodeFileRetryFailed"`
	CheckForUpdatesInterval      string `yaml:"checkForUpdatesInterval"`
	GithubPersonalToken          string `yaml:"githubPersonalToken"`
	GithubUsername               string `yaml:"githubUsername"`
	IsAutoRestartManaged         bool   `yaml:"isAutoRestartManaged"`
	ScanForMovieInterval         string `yaml:"scanForMovieInterval"`
	RootMovieScanDir             string `yaml:"rootMovieScanDir"`
	RootMovieMediaDir            string `yaml:"rootMovieMediaDir"`
	MovieFileRetryFailed         bool   `yaml:"movieFileRetryFailed"`
	MovieCustomFormat            string `yaml:"movieCustomFormat"`
	DbBucketsCleanupInterval     string `yaml:"dbBucketsCleanupInterval"`
	LogRotateMaxNumOfBackups     int    `yaml:"logRotateMaxNumOfBackups"`
	LogRotateMaxAgeOfBackups     int    `yaml:"logRotateMaxAgeOfBackups"`
	LogRotateMaxLogFileSize      int    `yaml:"logRotateMaxLogFileSize"`
	LogRotateCompressBackups     bool   `yaml:"logRotateCompressBackups"`
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
