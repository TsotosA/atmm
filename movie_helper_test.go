package main

import (
	"fmt"
	"testing"
)

func TestGetResolution(t *testing.T) {
	var tests = []struct {
		s string
		e string
	}{
		{"The.Book.of.Boba.Fett.S01E07.Chapter.7.720p.DSNP.WEB-DL.DDP5.1.Atmos.H.264-NOSiViD.mkv", "720p"},
		{"Vikings.Valhalla.S01E02.Viking.1080p.NF.WEB-DL.DDP5.1.Atmos.x264-TEPES.mkv", "1080p"},
		{"The.Shannara.Chronicles..E01E02.mkv", ""},
		{"Escape.from.Pretoria.2020.WEB-DL.DD5.1.H264-FGT", ""},
		{"Cars.3.2017.BluRay.720p.Greek.Audio.x264", "720p"},
	}

	for _, tt := range tests {
		testHName := fmt.Sprintf("input,expected:[%s,%s]", tt.s, tt.e)
		t.Run(testHName, func(t *testing.T) {
			r, _ := getMovieResolution(tt.s)
			if r != tt.e {
				t.Errorf("got [%s], wanted [%s]", r, tt.e)
			}
		})
	}
}
