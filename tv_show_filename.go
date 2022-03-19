package main

import (
	"fmt"
	"go.uber.org/zap"
	"path/filepath"
	"regexp"
)

func ParseFilenameCustomFormat(f string) ([]FilenameFormatPair, error) {
	openingCurly := 0
	closingCurly := 0
	currentFind := ""
	pairs := make([]FilenameFormatPair, 0)

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
			pairs = append(pairs, FilenameFormatPair{
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

func MapProperties(o TvShowEpisodeFile, ffp []FilenameFormatPair) error {
	for i, pair := range ffp {
		switch pair.PropertyName {
		case "SeriesTitle":
			ffp[i].PropertyValue = o.TvShow.Name
			break
		case "EpisodeTitle":
			ffp[i].PropertyValue = o.TvShowEpisode.Name
			break
		case "Season":
			ffp[i].PropertyValue = fmt.Sprintf("%d", o.TvShowEpisode.SeasonNumber)
			break
		case "Episode":
			ffp[i].PropertyValue = fmt.Sprintf("%d", o.TvShowEpisode.EpisodeNumber)
			break
		case "Season:00":
			ffp[i].PropertyValue = fmt.Sprintf("%02d", o.TvShowEpisode.SeasonNumber)
			break
		case "Episode:00":
			ffp[i].PropertyValue = fmt.Sprintf("%02d", o.TvShowEpisode.EpisodeNumber)
			break
		case "Resolution":
			ffp[i].PropertyValue = o.ParsedFilename.Resolution
			break
		}
	}
	return nil
}

func ReplaceCustomFormatStringToTitle(s string, p []FilenameFormatPair) (string, error) {
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

func MakeFilename(o TvShowEpisodeFile) (string, error) {
	customFormat := Conf.TvShowEpisodeFormat
	formatPairs, err := ParseFilenameCustomFormat(customFormat)
	if err != nil {
		return "", err
	}
	err = MapProperties(o, formatPairs)
	if err != nil {
		return "", err
	}
	title, err := ReplaceCustomFormatStringToTitle(customFormat, formatPairs)
	title = title + filepath.Ext(o.FilenameOriginal)
	if err != nil {
		return "", err
	}
	return title, nil
}
