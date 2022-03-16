package main

const (
	TvShowEpisodeFilesBucket   = "tv_show_episode_files"
	MovieFilesBucket           = "movie_files"
	GithubApiBaseUrl           = "https://api.github.com/"
	GithubJsonAcceptHeader     = "application/vnd.github.v3+json"
	HttpContentTypeOctetStream = "application/octet-stream"
	TmpDir                     = "./tmp"

	TvShowSeasonRegexp         = `([Ss]|(Season|SEASON))([\d]{1,4})\D`
	TvShowEpisodeSingleRegexp  = `([Ee]|(Episode|EPISODE|Ep|EP))([\d]{1,4})\D`
	TvShowEpisodeIsMultiRegexp = `(([Ee]|(Episode|EPISODE|Ep|EP))([\d]{1,2})){2,}\D`
	TvShowEpisodeTitleRegexp   = `(.*)+(([Ss]|(Season|SEASON))([\d]{1,4})\D)`
	MovieTitleRegexp           = `(.*)(\d{4})\W`
	MovieYearRegexp            = `(.*)(\d{4})\W`

	WindowsGeneratedFileThumbsDb = "Thumbs.db"

	ExampleConfigYaml = `
### https://www.themoviedb.org/
theMovieDbBaseApiUrlV3: https://api.themoviedb.org/3/

### https://www.themoviedb.org/ api key used to query their api
# you will need to provide one by registering to https://www.themoviedb.org/signup 
# and then going to setting -> API -> API Key (v3 auth)
apiKey:

### the root directory where the processed files will end up
rootMediaDir: "./testMediaDir"

### the root directory where the app will look for media files to process
rootScanDir: "./testScanDir"

### ignore for now, eventually it will enable the app to run and log as normally but without creating dirs and moving the files
dryRun: false 

### the log file location and name
# "./atmm.log" - logs to a file named atmm.log located next to the binary
logOutputPath: "./atmm.log" 

### the log level to use. available values are [debug | info | warn | error].
# info should be used as a default in order to avoid spamming the logs
logLevel: info

###  [prod] sets the logging output to json (useful for piping the logs to prometheus/grafana etc..) 
# [dev] sets the logging output to a more human friendly format
environment: prod 

### controls the scan for tv series
# empty string [""] will disable scan completely
# valid examples can be found at the package godoc https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc
cron: ""

###
tvShowEpisodeFormat: "{SeriesTitle} S{Season:00}E{Episode:00} - {EpisodeTitle}"

###
tvShowEpisodeFileRetryFailed: false

###
checkForUpdatesInterval: "@every 60s"

###
githubPersonalToken:

###
githubUsername:

###
isAutoRestartManaged: false

###
scanForMovieInterval: "@every 15s" #"@every 10s"

###
rootMovieScanDir: "./testMovieMediaDir"

###
rootMovieMediaDir: "./testMovieScanDir"

###
movieFileRetryFailed: false

###
movieCustomFormat: "{MovieTitle} ({MovieReleaseYear:0000})"

###
dbBucketsCleanupInterval:"@weekly"

###
logRotateMaxNumOfBackups: 1

###
logRotateMaxAgeOfBackups: 1

###
logRotateMaxLogFileSize: 1

###
logRotateCompressBackups: true
`
)
