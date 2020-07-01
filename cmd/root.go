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

	"github.com/spf13/cobra"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/po3rin/chrbmfzf/chrbm"
	"github.com/spf13/viper"
)

var bookmarkFile string

var rootCmd = &cobra.Command{
	Use:   "chrbmfzf",
	Short: "chrbmfzf fuzzy-finder for Google Chrome Bookmark",
	Long:  `chrbmfzf fuzzy-finder for Google Chrome Bookmark`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(bookmarkFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tracks, err := chrbm.NewBookmark(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		id, err := fuzzyfinder.Find(
			tracks,
			func(i int) string {
				return tracks[i].Name
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return fmt.Sprintf("%s\n\nArtist: %s\nAlbum:  %s",
					tracks[i].Name,
					tracks[i].URL,
					tracks[i].DateAdded,
				)
			}),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("https://" + tracks[id].URL)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&bookmarkFile, "file", "f", path.Join(home, `Library/Application Support/Google/Chrome/Default/Bookmarks`), "bookmark file path")
}

func initConfig() {
	viper.AutomaticEnv()
}
