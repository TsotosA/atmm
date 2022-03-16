package main

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

func ParseMovieFilename(filename string) (parsedFilename ParsedFilename, err error) {
	t, err := getMovieTitle(filename)
	e, err := getMovieYear(filename)

	if err != nil {
		fmt.Printf("skipping %v. could not match title to search for.\n", filename)
		return ParsedFilename{}, errors.New("failed to parse filename")
	}
	t = strings.Replace(t, ".", " ", -1)
	t = strings.TrimSpace(t)

	return ParsedFilename{
		Title: t,
		Year:  e,
	}, nil
}
func getMovieTitle(s string) (string, error) {
	rxp, err := regexp.Compile(MovieTitleRegexp)
	if err != nil {
		return "", err
	}
	title := rxp.FindStringSubmatch(s)
	if len(title) < 3 {
		zap.S().Warnf("failed to parse title from filename: %s", s)
		return "", errors.New("failed to parse title from filename")
	}
	zap.S().Debugf("getMovieTitle() title:%s", title[1])
	return title[1], nil
}

func getMovieYear(s string) (string, error) {
	rxp, err := regexp.Compile(MovieYearRegexp)
	if err != nil {
		return "", err
	}
	title := rxp.FindStringSubmatch(s)
	if len(title) < 3 {
		zap.S().Warnf("failed to parse title from filename: %s", s)
		return "", errors.New("failed to parse title from filename")
	}
	zap.S().Debugf("getMovieYear() title:%s", title[2])
	return title[2], nil
}
