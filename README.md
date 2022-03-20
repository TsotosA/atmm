# @ media manager (atmm)

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/tsotosa/atmm/graphs/commit-activity)
[![Windows](https://svgshare.com/i/ZhY.svg)](https://svgshare.com/i/ZhY.svg)
[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/TsotosA/atmm/main)

A `golang` based application for `managing tv-shows & movies` (only renaming and copying for now).  
It is developed as a part of my home server infrastructure since existing solutions were too overpowered and heavy for
the use case.   
Also, it serves as an excuse to
learn [![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)

## 🚧 disclaimer 🚧

> **this software is in a very early stage of development**  
> i.e. everything is prone to change (functionality, api, configurations, stability etc..)

## features

- `windows` & `linux` binaries available
- supports `tv series` & `movies`
- `auto-update` to latest release
- `minimal` cpu and memory `footprint`
- `minimal maintenance` required. configure once and forget.
- config.yaml file `hot reload`

## configuration

The app will look for a `config.yaml` file next to the app binary. All configuration/settings are controlled from this
file.  
In case no such file exists it will create one (`example.config.yaml`) with some placeholder values and explanations (
this file needs to be renamed after setup !).   
Also, at every update it will check that all the configuration keys are available at config.yaml. In case some of them
are missing it will print them out at the logs before exiting.

#### *config.yaml*

```yaml
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
```

## custom filenames

the filename format to be used is controlled by the `movieCustomFormat` and `tvShowEpisodeFormat` configurations.  
the available variables for `movies` and `tv shows` are respectively:

| **Variable**              | **Explanation** |
|---------------------------|-----------------|
| `{MovieTitle}`            |                 |
| `{MovieReleaseYear:0000}` |                 |
| `{Resolution}`            |                 |

| **Variable**     | **Explanation** |
|------------------|-----------------|
| `{SeriesTitle}`  |                 |
| `{Episode:00}`   |                 |
| `{Season:00}`    |                 |
| `{EpisodeTitle}` |                 |

