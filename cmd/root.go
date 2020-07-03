/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/po3rin/bmfzf/chrbm"
	"github.com/spf13/viper"
)

var bookmarkFile string

var rootCmd = &cobra.Command{
	Use:   "bmfzf",
	Short: "bmfzf fuzzy-finder for Google Chrome Bookmark",
	Long:  `bmfzf fuzzy-finder for Google Chrome Bookmark`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(bookmarkFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tracks, err := chrbm.ListBookmarks(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		id, err := fuzzyfinder.Find(
			tracks,
			func(i int) string {
				return tracks[i].Path
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return fmt.Sprintf("%s\n\nPath: %s\nURL: %s",
					tracks[i].Name,
					tracks[i].Path,
					tracks[i].URL,
				)
			}),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(tracks[id].URL)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func bookmarkFileLocation() (string, error) {
	os := runtime.GOOS
	switch os {
	case "windows":
		return `AppData\Local\Google\Chrome\User Data\Default\bookmarks`, nil
	case "darwin":
		return `Library/Application Support/Google/Chrome/Default/Bookmarks`, nil
	case "linux":
		return `.config/google-chrome/Default`, nil
	default:
		return "", fmt.Errorf("sorry... your OS %v is not supported. please specify you bookmark file using -f flag.", os)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	location, err := bookmarkFileLocation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&bookmarkFile, "file", "f", path.Join(home, location), "bookmark file path")
}

func initConfig() {
	viper.AutomaticEnv()
}
