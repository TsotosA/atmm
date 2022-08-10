package main

import (
	"fmt"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"path/filepath"
	"regexp"
)

func ParseMovieFilenameCustomFormat(f string) ([]model.FilenameFormatPair, error) {
	openingCurly := 0
	closingCurly := 0
	currentFind := ""
	pairs := make([]model.FilenameFormatPair, 0)

	for i, v := range f {
		if string(v) == "{" {
			openingCurly = i
			currentFind = "openingCurly"
		}

		if string(v) == "}" {
			closingCurly = i
			currentFind = "closingCurly"
		}

		addPair := closingCurly != 0 && currentFind == "closingCurly"

		if addPair {
			pairs = append(pairs, model.FilenameFormatPair{
				StartIndex:   openingCurly,
				EndIndex:     closingCurly,
				PropertyName: f[openingCurly+1 : closingCurly],
			})
			closingCurly = 0
			openingCurly = 0
			zap.S().Debugf("pairs %#v", pairs)
		}
	}
	return pairs, nil
}

func MapMovieFilenameProperties(o model.MovieFile, ffp []model.FilenameFormatPair) error {
	for i, pair := range ffp {
		switch pair.PropertyName {
		case "MovieTitle":
			ffp[i].PropertyValue = o.Movie.Title
			break
		case "MovieReleaseYear:0000":
			ffp[i].PropertyValue = GetMovieYearFromReleaseDate(o.Movie.ReleaseDate)
			break
		case "Resolution":
			ffp[i].PropertyValue = o.ParsedFilename.Resolution
			break
		}
	}
	return nil
}

func ReplaceCustomMovieFormatStringToTitle(s string, p []model.FilenameFormatPair) (string, error) {
	t := s
	for _, pair := range p {
		reg, err := regexp.Compile(fmt.Sprintf("{%s}", pair.PropertyName))
		if err != nil {
			return "", err
		}
		t = reg.ReplaceAllString(t, pair.PropertyValue)
	}
	return t, nil
}

func MakeMovieFilename(o model.MovieFile) (string, error) {
	customFormat := config.Conf.MovieCustomFormat
	formatPairs, err := ParseMovieFilenameCustomFormat(customFormat)
	if err != nil {
		return "", err
	}
	err = MapMovieFilenameProperties(o, formatPairs)
	if err != nil {
		return "", err
	}
	title, err := ReplaceCustomMovieFormatStringToTitle(customFormat, formatPairs)
	title = title + filepath.Ext(o.FilenameOriginal)
	if err != nil {
		return "", err
	}
	return title, nil
}

func GetMovieYearFromReleaseDate(s string) string {
	r := `\d{4}`
	reg, err := regexp.Compile(r)
	if err != nil {
		return ""
	}
	return reg.FindString(s)
}
