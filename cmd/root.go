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
	"time"

	"github.com/spf13/cobra"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/viper"
)

var bookmarkFile string

type bookmark struct {
	name      string
	url       string
	dateAdded time.Time
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chrbmfzf",
	Short: "chrbmfzf fuzzy-finder for Google Chrome Bookmark",
	Long:  `chrbmfzf fuzzy-finder for Google Chrome Bookmark`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	b, err := ioutil.ReadFile(bookmarkFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = b

	var tracks = []bookmark{
		{"pon", "po3rin.com", time.Now()},
		{"twitter", "twitter.com", time.Now()},
		{"golang", "godoc,org", time.Now()},
	}

	id, err := fuzzyfinder.Find(
		tracks,
		func(i int) string {
			return tracks[i].name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("%s\n\nArtist: %s\nAlbum:  %s",
				tracks[i].name,
				tracks[i].url,
				tracks[i].dateAdded,
			)
		}),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(tracks[id].url)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&bookmarkFile, "file", path.Join(home, `Library/Application Support/Google/Chrome/Default/Bookmarks`), "config file (default is $HOME/Library/Application Support/Google/Chrome/Default/Bookmarks)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
