package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type emote struct {
	name string
	url  string
}

func getContent(url string) (io.ReadCloser, error) {
	log.Printf("Downloading %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("Invalid StatusCode %d", res.StatusCode)
	}
	return res.Body, nil
}

func collectEmotes(url string) []emote {
	emotes := []emote{}

	pageBody, err := getContent(url)
	if err != nil {
		log.Fatal(err)
		return emotes
	}
	defer pageBody.Close()

	doc, err := goquery.NewDocumentFromReader(pageBody)
	if err != nil {
		log.Fatal(err)
		return emotes
	}

	selection := doc.Find("samp")
	for i := range selection.Nodes {
		single := selection.Eq(i)
		img := single.Prev().Find("img")
		href, exists := img.Attr("src")
		if exists == false {
			continue
		}
		emotes = append(emotes, emote{
			name: single.Text(),
			url:  href,
		})
	}

	return emotes
}

type indexedReadCloser struct {
	idx   int
	value io.ReadCloser
}

func downloadImages(emotes []emote) map[int]io.ReadCloser {
	images := map[int]io.ReadCloser{}
	ch := make(chan indexedReadCloser)

	var wg sync.WaitGroup
	for i, emote := range emotes {
		wg.Add(1)
		go func(idx int, url string, ch chan<- indexedReadCloser, wg *sync.WaitGroup) {
			defer wg.Done()

			indexed := indexedReadCloser{
				idx: idx,
			}

			imageBody, err := getContent(url)
			if err != nil {
				log.Fatal(err)
				indexed.value = nil
				ch <- indexed
			}
			indexed.value = imageBody
			ch <- indexed
		}(i, emote.url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for indexed := range ch {
		images[indexed.idx] = indexed.value
	}

	return images
}

func writeZip(emotes []emote, output string) {
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	z := zip.NewWriter(f)
	defer z.Close()

	images := downloadImages(emotes)

	for i, emote := range emotes {
		filename := url.PathEscape(fmt.Sprintf("%s.png", emote.name))
		file, err := z.Create(filename)
		if err != nil {
			log.Fatal(err)
			continue
		}
		imageBody := images[i]
		if imageBody == nil {
			continue
		}

		if _, err := io.Copy(file, imageBody); err != nil {
			log.Fatal(err)
		}
		imageBody.Close()
	}
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
			emotes := collectEmotes(url)
			writeZip(emotes, output)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&url, "url", "U", "https://www.twitchmetrics.net/emotes", "URL of the emotes you want.")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "O", "output.zip", "Output filename of zip.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
