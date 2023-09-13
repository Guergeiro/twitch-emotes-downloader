package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "twe-dl",
	Short: "Download twitch emotes in bulk",
	Long: `(tw)itch (e)motes (d)own(l)oader is a cli that downloads emotes in bulk
                                 from https://www.twitchmetrics.net/`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
