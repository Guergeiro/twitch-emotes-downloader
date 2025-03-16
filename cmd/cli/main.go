package main

import (
	"fmt"
	"os"

	"github.com/guergeiro/twitch-emotes-downloader/internal/app"
	"github.com/spf13/cobra"
)

func main() {
	url := "https://www.twitchmetrics.net/emotes"
	output := "output.zip"

	rootCmd := &cobra.Command{
		Use:   "twe-dl {...urls}",
		Short: "Download twitch emotes in bulk",
		Long: `(tw)itch (e)motes (d)own(l)oader is a cli that downloads emotes in bulk
                                   from https://www.twitchmetrics.net/`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = append(args, url)
			}
			app.DownloadEmotes(args, output)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&output, "output", "O", "output.zip", "Output filename of zip.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
