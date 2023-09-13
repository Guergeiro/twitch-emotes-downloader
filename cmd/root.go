package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type emote struct {
	name string
	url  string
}

func collectEmotes(url string) ([]emote, error) {
	emotes := []emote{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("Invalid")
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("samp").Each(func(idx int, s *goquery.Selection) {
		img := s.Prev().Find("img")
		if href, exists := img.Attr("src"); exists {
			emotes = append(emotes, emote{
				name: s.Text(),
				url:  href,
			})
		}
	})
	fmt.Println(emotes)

	return emotes, nil
}

func Execute() {
	url := "https://www.twitchmetrics.net/emotes"
	output := "output.zip"

	rootCmd := &cobra.Command{
		Use:   "twe-dl",
		Short: "Download twitch emotes in bulk",
		Long: `(tw)itch (e)motes (d)own(l)oader is a cli that downloads emotes in bulk
                                   from https://www.twitchmetrics.net/`,
		Run: func(cmd *cobra.Command, args []string) {
			if emotes, err := collectEmotes(url); err == nil {
				fmt.Println(emotes)
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&url, "url", "U", "https://www.twitchmetrics.net/emotes", "URL of the emotes you want.")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "O", "output.zip", "Output filename of zip.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
