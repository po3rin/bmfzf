package chrbm_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/po3rin/chrbmfzf/chrbm"
)

func TestNewBookmark(t *testing.T) {
	tests := []struct {
		name string
		file string
		want []chrbm.Bookmark
	}{
		{
			name: "simple",
			file: "./testdata/test.json",
			want: []chrbm.Bookmark{
				{
					Name: "Google",
					Path: "/Google",
					URL:  "https://www.google.co.jp/?gws_rd=ssl",
				},
				{
					Name: "Facebook",
					Path: "/private/Facebook",
					URL:  "https://www.facebook.com/",
				},
				{
					Name: "Twitter",
					Path: "/private/Twitter",
					URL:  "https://twitter.com/",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteJSON, err := ioutil.ReadFile(tt.file)
			if err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

			got, err := chrbm.NewBookmark(byteJSON)
			if err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nwant : %+v\ngot  : %+v\n", tt.want, got)
			}
		})
	}
}
