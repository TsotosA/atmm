package main

import (
	"github.com/tsotosa/atmm/model"
	"reflect"
	"testing"
)

func TestParseMovieFilenameCustomFormat(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.FilenameFormatPair
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{f: "{MovieTitle} {MovieReleaseYear:0000} {Resolution}"},
			want: []model.FilenameFormatPair{
				{
					StartIndex:    0,
					EndIndex:      11,
					PropertyName:  "MovieTitle",
					PropertyValue: "",
				},
				{
					StartIndex:    13,
					EndIndex:      35,
					PropertyName:  "MovieReleaseYear:0000",
					PropertyValue: "",
				},
				{
					StartIndex:    37,
					EndIndex:      48,
					PropertyName:  "Resolution",
					PropertyValue: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMovieFilenameCustomFormat(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMovieFilenameCustomFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMovieFilenameCustomFormat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
