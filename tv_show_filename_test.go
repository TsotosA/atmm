package main

import (
	"fmt"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/model"
	"reflect"
	"testing"
)

func TestParseFilenameCustomFormat(t *testing.T) {
	var tests = []struct {
		input    string
		expected []model.FilenameFormatPair
		error    error
	}{
		{"{SeriesTitle} S{Season:00}E{Episode:00} - {EpisodeTitle}", []model.FilenameFormatPair{
			{
				StartIndex:    0,
				EndIndex:      12,
				PropertyName:  "SeriesTitle",
				PropertyValue: "",
			},
			{
				StartIndex:    15,
				EndIndex:      25,
				PropertyName:  "Season:00",
				PropertyValue: "",
			},
			{
				StartIndex:    27,
				EndIndex:      38,
				PropertyName:  "Episode:00",
				PropertyValue: "",
			},
			{
				StartIndex:    42,
				EndIndex:      55,
				PropertyName:  "EpisodeTitle",
				PropertyValue: "",
			},
		}, nil},
	}
	for _, tt := range tests {
		testHName := fmt.Sprintf("input, expected, error:[%s, %#v, %v]", tt.input, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := ParseFilenameCustomFormat(tt.input)
			if !reflect.DeepEqual(r, tt.expected) {
				if err != nil && tt.error != nil {
					t.Errorf("got [%#v] [%v] , wanted [%#v] [%v]", r, err, tt.expected, tt.error)
				}
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}

func TestMapProperties(t *testing.T) {
	var tests = []struct {
		a        model.TvShowEpisodeFile
		b        []model.FilenameFormatPair
		expected []model.FilenameFormatPair
		error    error
	}{
		{model.TvShowEpisodeFile{
			FilenameOriginal: "",
			FilenameNew:      "",
			AbsolutePath:     "",
			TvShow: model.TheMovieDbTvShow{
				Name: "tv show name",
			},
			TvShowEpisode: model.TheMovieDbTvShowEpisodeDetails{
				EpisodeNumber: 11,
				Name:          "tv show episode name",
				SeasonNumber:  22,
			},
			SuccessfulParseOriginal: false,
			SuccessfulCopyFile:      false,
		},
			[]model.FilenameFormatPair{
				{
					PropertyName:  "SeriesTitle",
					PropertyValue: "",
				},
				{
					PropertyName:  "Season:00",
					PropertyValue: "",
				},
				{
					PropertyName:  "Episode:00",
					PropertyValue: "",
				},
				{
					PropertyName:  "EpisodeTitle",
					PropertyValue: "",
				},
			},
			[]model.FilenameFormatPair{
				{
					PropertyName:  "SeriesTitle",
					PropertyValue: "tv show name",
				},
				{
					PropertyName:  "Season:00",
					PropertyValue: "22",
				},
				{
					PropertyName:  "Episode:00",
					PropertyValue: "11",
				},
				{
					PropertyName:  "EpisodeTitle",
					PropertyValue: "tv show episode name",
				},
			},
			nil},
	}
	for _, tt := range tests {
		testHName := fmt.Sprintf("a,b,expected,error:[%v],[%v],[%#v],[%v]", tt.a, tt.b, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			err := MapProperties(tt.a, tt.b)
			if !reflect.DeepEqual(tt.b, tt.expected) {
				if err != nil && tt.error != nil {
					t.Errorf("got [%v],wanted [%#v] [%v]", tt.b, err, tt.expected)
				}
				t.Errorf("got [%#v], wanted [%#v]", tt.b, tt.expected)
			}
		})
	}

}

func TestReplaceCustomFormatStringToTitle(t *testing.T) {
	var tests = []struct {
		a        string
		b        []model.FilenameFormatPair
		expected string
		error    error
	}{
		{"{SeriesTitle} S{Season:00}E{Episode:00} - {EpisodeTitle}",
			[]model.FilenameFormatPair{
				{
					PropertyName:  "SeriesTitle",
					PropertyValue: "series title",
				},
				{
					PropertyName:  "Season:00",
					PropertyValue: "season",
				},
				{
					PropertyName:  "Episode:00",
					PropertyValue: "episode",
				},
				{
					PropertyName:  "EpisodeTitle",
					PropertyValue: "episode title",
				},
			},
			"series title SseasonEepisode - episode title",
			nil,
		},
	}
	for _, tt := range tests {
		testHName := fmt.Sprintf("a,b,expected,error:[%v],[%v],[%#v],[%v]", tt.a, tt.b, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := ReplaceCustomFormatStringToTitle(tt.a, tt.b)
			if r != tt.expected {
				if err != nil && tt.error != nil {
					t.Errorf("got [%v],wanted [%#v] [%v]", r, tt.expected, err)
				}
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}

func TestMakeFilename(t *testing.T) {
	var tests = []struct {
		a        model.TvShowEpisodeFile
		expected string
		error    error
	}{
		{
			model.TvShowEpisodeFile{
				FilenameOriginal: "",
				FilenameNew:      "",
				AbsolutePath:     "",
				TvShow: model.TheMovieDbTvShow{
					Name: "tv show name",
				},
				TvShowEpisode: model.TheMovieDbTvShowEpisodeDetails{
					EpisodeNumber: 11,
					Name:          "tv show episode name",
					SeasonNumber:  22,
				},
				SuccessfulParseOriginal: false,
				SuccessfulCopyFile:      false,
			},
			"tv show name S22E11 - tv show episode name",
			nil,
		},
	}

	config.Conf.TvShowEpisodeFormat = "{SeriesTitle} S{Season:00}E{Episode:00} - {EpisodeTitle}"
	for _, tt := range tests {
		testHName := fmt.Sprintf("a,expected,error:[%v],[%v],[%v]", tt.a, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := MakeFilename(tt.a)
			if r != tt.expected {
				if err != nil && tt.error != nil {
					t.Errorf("got [%v],wanted [%#v] [%v]", r, tt.expected, err)
				}
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}
