package command

import (
	"github.com/guergeiro/twitch-emotes-downloader/internal/http"
	"github.com/guergeiro/twitch-emotes-downloader/internal/mapper"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/adapter/controller"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/application/usecase"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	c := controller.NewDownloadEmotesController(
		usecase.NewDownloadEmotesUseCase(
			http.Download,
			mapper.NewGoQueryHtmlEmoteMapper(),
		),
		usecase.NewDownloadImageUseCase(
			http.Download,
		),
		usecase.NewWriteZipUseCase(),
	)

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

			c.Handle(args, output)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&output, "output", "O", "output.zip", "Output filename of zip.")

	return rootCmd
}
