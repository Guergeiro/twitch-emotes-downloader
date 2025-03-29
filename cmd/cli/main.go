package main

import (
	"fmt"
	"os"

	"github.com/guergeiro/twitch-emotes-downloader/internal/init/command"
)

func main() {
	rootCmd := command.CreateCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
