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
	TvShowResolutionRegexp     = `\W(\d+p)\W`
	MovieTitleRegexp           = `(.*)(\d{4})\W`
	MovieYearRegexp            = `(.*)(\d{4})\W`
	MovieResolutionRegexp      = `\W(\d+p)\W`

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

### the file to write logs to
# "./atmm.log" - logs to a file named atmm.log located next to the binary.
logOutputPath: "./atmm.log" 

### the log level to use. available values are [debug | info | warn | error].
# [info] should be used as a default in order to avoid spamming the logs.
logLevel: info

###  [prod] sets the logging output to json (useful for piping the logs to prometheus/grafana etc..) .
# [dev] sets the logging output to a more human friendly format
environment: prod 

### controls the scan for tv series.
# empty string [""] will disable scan completely.
# valid examples can be found at the package godoc https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc.
cron: ""

### naming convention for tv shows. see README for available variables.
tvShowEpisodeFormat: "{SeriesTitle} S{Season:00}E{Episode:00} - {EpisodeTitle}"

### retry failed tv show files on every scan.
tvShowEpisodeFileRetryFailed: false

### how often to check for new version release.
checkForUpdatesInterval: "@every 1d"

### ignore for now.
githubPersonalToken:

### ignore for now.
githubUsername:

### in case this is false when auto updating the app will not exit itself. this means that the new version will be applied after a reboot only.
# if this is true the app will kill itself after updating, 
# and it is expected that it will be auto restarted afterwards (like if its setup as a service in systemd etc..).
isAutoRestartManaged: false

### controls the scan for new movie files.
# empty string [""] will disable scan completely.
# valid examples can be found at the package godoc https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc.
scanForMovieInterval: "@every 15s" #"@every 10s"

### the root directory where the app will look for movie media files to process
rootMovieScanDir: "./testMovieMediaDir"

### the root directory where processed movie files will end up
rootMovieMediaDir: "./testMovieScanDir"

### retry failed movies files on every scan.
movieFileRetryFailed: false

###  naming convention for movies. see README for available variables
movieCustomFormat: "{MovieTitle} ({MovieReleaseYear:0000})"

### aligns the actual files of the filesystem (movie and series scan dirs) with the db entries. 
# weekly is a good default value.
dbBucketsCleanupInterval: "@weekly"

### the maximum number of old log files to retain
logRotateMaxNumOfBackups: 1

### the maximum number of days to retain old log files based on the timestamp encoded in their filename
logRotateMaxAgeOfBackups: 1

### the maximum size in megabytes of the log file before it gets rotated
logRotateMaxLogFileSize: 1

### determines if the rotated log files should be compressed  using gzip
logRotateCompressBackups: true
`
)
