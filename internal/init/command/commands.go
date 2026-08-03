package command

import (
	"fmt"

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
	size := "3.0"

	rootCmd := &cobra.Command{
		Use:   "twe-dl {...urls}",
		Short: "Download twitch emotes in bulk",
		Long: `(tw)itch (e)motes (d)own(l)oader is a cli that downloads emotes in bulk
                                   from https://www.twitchmetrics.net/`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				args = append(args, url)
			}

			if err := validateEmoteSize(size); err != nil {
				return err
			}

			return c.Handle(args, output, size)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&output, "output", "O", "output.zip", "Output filename of zip.")
	rootCmd.PersistentFlags().StringVarP(&size, "size", "s", "3.0", "Emote size to download: 1.0 (28px), 2.0 (56px), or 3.0 (112px).")

	return rootCmd
}

func validateEmoteSize(size string) error {
	switch size {
	case "1.0", "2.0", "3.0":
		return nil
	default:
		return fmt.Errorf("invalid size %q: expected one of 1.0, 2.0, or 3.0", size)
	}
}
