package main

import (
	"errors"
	"fmt"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

func ParseTvShowFilename(filename string) (parsedFilename model.ParsedFilename, err error) {
	s, err := getSeason(filename)
	e, err := getEpisode(filename)
	t, err := getTitle(filename)
	r, _ := getResolution(filename)

	if err != nil {
		fmt.Printf("skipping %v. could not match title to search for.\n", filename)
		return model.ParsedFilename{}, errors.New("failed to parse filename")
	}
	//titleRegexp, _ := regexp.Compile(Conf.TvShowTitleRegexp)
	//x := titleRegexp.FindStringSubmatch(filename)
	//
	//if len(x) <= 3 {
	//	fmt.Printf("skipping %v. could not match title to search for.\n", filename)
	//	return ParsedFilename{}, errors.New("failed to parse filename")
	//}

	//title := x[1]
	//season := x[2]
	//episode := x[3]

	//title = strings.Replace(title, ".", " ", -1)
	t = strings.Replace(t, ".", " ", -1)
	t = strings.TrimSpace(t)

	//return ParsedFilename{
	//	Name:          title,
	//	SeasonNumber:  season,
	//	EpisodeNumber: episode,
	//}, nil

	return model.ParsedFilename{
		Name:          t,
		SeasonNumber:  s,
		EpisodeNumber: e,
		Resolution:    r,
	}, nil
}

func getSeason(s string) (string, error) {
	rxp, err := regexp.Compile(gconst.TvShowSeasonRegexp)
	if err != nil {
		return "", err
	}
	season := rxp.FindStringSubmatch(s)
	if len(season) < 3 {
		zap.S().Warnf("failed to parse season from filename: %s", s)
		return "", errors.New("failed to parse season from filename")
	}
	zap.S().Debugf("getSeason() season:%s", season[3])
	return season[3], nil
}

func getEpisode(s string) (string, error) {
	rxp, err := regexp.Compile(gconst.TvShowEpisodeSingleRegexp)
	if err != nil {
		return "", err
	}
	episode := rxp.FindStringSubmatch(s)
	if len(episode) < 3 {
		zap.S().Warnf("failed to parse episode from filename: %s", s)
		return "", errors.New("failed to parse episode from filename")
	}
	zap.S().Debugf("getEpisode() episode:%s", episode[3])
	return episode[3], nil
}

func getTitle(s string) (string, error) {
	rxp, err := regexp.Compile(gconst.TvShowEpisodeTitleRegexp)
	if err != nil {
		return "", err
	}
	title := rxp.FindStringSubmatch(s)
	if len(title) < 3 {
		zap.S().Warnf("failed to parse title from filename: %s", s)
		return "", errors.New("failed to parse title from filename")
	}
	zap.S().Debugf("getTitle() title:%s", title[1])
	return title[1], nil
}

func getResolution(s string) (string, error) {
	rxp, err := regexp.Compile(gconst.TvShowResolutionRegexp)
	if err != nil {
		return "", err
	}
	res := rxp.FindStringSubmatch(s)
	if len(res) < 1 {
		zap.S().Warnf("failed to parse resolution from filename: %s", s)
		return "", errors.New("failed to parse resolution from filename")
	}
	zap.S().Debugf("getResolution() res:%s", res[0])
	return res[1], nil
}
