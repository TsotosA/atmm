package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseTvShowFilename(t *testing.T) {
	var tests = []struct {
		input    string
		expected ParsedFilename
		error    error
	}{
		{"The.Book.of.Boba.Fett.S01E07.Chapter.7.720p.DSNP.WEB-DL.DDP5.1.Atmos.H.264-NOSiViD.mkv", ParsedFilename{Name: "The Book of Boba Fett", SeasonNumber: "01", EpisodeNumber: "07"}, nil},
	}

	for _, tt := range tests {
		testHName := fmt.Sprintf("input, expected, error:[%s, %s, %v]", tt.input, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := ParseTvShowFilename(tt.input)
			if r != tt.expected {
				if err != nil && tt.error != nil {
					t.Errorf("got [%#v] [%v] , wanted [%s] [%v]", r, err, tt.expected, tt.error)
				}
				t.Errorf("got [%#v], wanted [%#v]", r, tt.expected)
			}
		})
	}
}

func TestGetSeason(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
		error    error
	}{
		{"The.Book.of.Boba.Fett.S01E07.Chapter.7.720p.DSNP.WEB-DL.DDP5.1.Atmos.H.264-NOSiViD.mkv", "01", nil},
		{"SEAL.Team.S02E03.XviD-AFG-[GTtvSD].avi", "02", nil},
		{"Vikings.Valhalla.S01E02.Viking.1080p.NF.WEB-DL.DDP5.1.Atmos.x264-TEPES.mkv", "01", nil},
		{"The.Shannara.Chronicles.S01E01E02.mkv", "01", nil},
		{"The.Shannara.Chronicles.S1E01E02.mkv", "1", nil},
		{"The.Shannara.Chronicles.Season1E01E02.mkv", "1", nil},
		{"The.Shannara.Chronicles.Season01E01E02.mkv", "01", nil},
		{"The.Shannara.Chronicles.Season0111.E01E02.mkv", "0111", nil},
		{"The.Shannara.Chronicles.SEASON01.E01E02.mkv", "01", nil},
		{"The.Shannara.Chronicles..E01E02.mkv", "", errors.New("failed to parse season from filename")},
	}

	for _, tt := range tests {
		testHName := fmt.Sprintf("input, expected, error:[%s, %s, %v]", tt.input, tt.expected, tt.error)
		t.Run(testHName, func(t *testing.T) {
			r, err := getSeason(tt.input)
			if r != tt.expected {
				if err != nil && tt.error != nil {
					t.Errorf("got [%s] [%v] , wanted [%s] [%v]", r, err, tt.expected, tt.error)
				}
				t.Errorf("got [%s], wanted [%s]", r, tt.expected)
			}
		})
	}
}

func TestGetEpisode(t *testing.T) {
	var tests = []struct {
		s string
		e string
	}{
		{"The.Book.of.Boba.Fett.S01E07.Chapter.7.720p.DSNP.WEB-DL.DDP5.1.Atmos.H.264-NOSiViD.mkv", "07"},
		{"SEAL.Team.S02E03.XviD-AFG-[GTtvSD].avi", "03"},
		{"Vikings.Valhalla.S01E02.Viking.1080p.NF.WEB-DL.DDP5.1.Atmos.x264-TEPES.mkv", "02"},
		{"The.Shannara.Chronicles.S01E01E02.mkv", "01"},
		{"The.Shannara.Chronicles.S1E01E02.mkv", "01"},
		{"The.Shannara.Chronicles.Season1E01E02.mkv", "01"},
		{"The.Shannara.Chronicles.Season01E01E02.mkv", "01"},
		{"The.Shannara.Chronicles.Season01.E0111E02.mkv", "0111"},
		{"The.Shannara.Chronicles.SEASON01.E01E02.mkv", "01"},
		{"The.Shannara.Chronicles..E01E02.mkv", "01"},
	}

	for _, tt := range tests {
		testHName := fmt.Sprintf("input,expected:[%s,%s]", tt.s, tt.e)
		t.Run(testHName, func(t *testing.T) {
			r, _ := getEpisode(tt.s)
			if r != tt.e {
				t.Errorf("got [%s], wanted [%s]", r, tt.e)
			}
		})
	}
}

func TestGetTitle(t *testing.T) {
	var tests = []struct {
		s string
		e string
	}{
		{"The.Book.of.Boba.Fett.S01E07.Chapter.7.720p.DSNP.WEB-DL.DDP5.1.Atmos.H.264-NOSiViD.mkv", "The.Book.of.Boba.Fett."},
		{"SEAL.Team.S02E03.XviD-AFG-[GTtvSD].avi", "SEAL.Team."},
		{"Vikings.Valhalla.S01E02.Viking.1080p.NF.WEB-DL.DDP5.1.Atmos.x264-TEPES.mkv", "Vikings.Valhalla."},
		{"The.Shannara.Chronicles.S01E01E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles.S1E01E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles.Season1E01E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles.Season01E01E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles.Season01.E01E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles.SEASON01.E01E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles.SEASON01.E0111E02.mkv", "The.Shannara.Chronicles."},
		{"The.Shannara.Chronicles..E01E02.mkv", ""},
	}

	for _, tt := range tests {
		testHName := fmt.Sprintf("input,expected:[%s,%s]", tt.s, tt.e)
		t.Run(testHName, func(t *testing.T) {
			r, _ := getTitle(tt.s)
			if r != tt.e {
				t.Errorf("got [%s], wanted [%s]", r, tt.e)
			}
		})
	}
}
