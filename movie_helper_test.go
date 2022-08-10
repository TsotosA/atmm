package main

import (
	"fmt"
	"github.com/tsotosa/atmm/model"
	"reflect"
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

func TestGetMovieYear(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "year - happy path",
			args:    args{s: "Escape.from.Pretoria.2020.WEB-DL.DD5.1.H264-FGT"},
			want:    "2020",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMovieYear(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMovieYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMovieYear() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMovieTitle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "parse success",
			args:    args{s: "Escape.from.Pretoria.2020.WEB-DL.DD5.1.H264-FGT"},
			want:    "Escape.from.Pretoria.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMovieTitle(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMovieTitle() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMovieTitle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMovieFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name               string
		args               args
		wantParsedFilename model.ParsedFilename
		wantErr            bool
	}{
		{
			name: "year - happy path",
			args: args{"Escape.from.Pretoria.2020.WEB-DL.DD5.1.H264-FGT"},
			wantParsedFilename: model.ParsedFilename{
				Name:          "",
				Title:         "Escape from Pretoria",
				SeasonNumber:  "",
				EpisodeNumber: "",
				Year:          "2020",
				Resolution:    "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParsedFilename, err := ParseMovieFilename(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMovieFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParsedFilename, tt.wantParsedFilename) {
				t.Errorf("ParseMovieFilename() gotParsedFilename = %v, want %v", gotParsedFilename, tt.wantParsedFilename)
			}
		})
	}
}
