package app

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
)

func DownloadEmotes(urls []string, output string) {
	emotes := parseHtmls(urls)
	images := downloadImages(emotes)
	writeZip(images, output)
}

type htmlEmote struct {
	name string
	href string
}

func getContent(href string) (io.ReadCloser, error) {
	log.Printf("Downloading %s\n", href)
	u, err := url.Parse(href)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("Invalid StatusCode %d", res.StatusCode)
	}
	return res.Body, nil
}

func parseHtml(url string) []htmlEmote {
	emotes := []htmlEmote{}

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
		emotes = append(emotes, htmlEmote{
			name: single.Text(),
			href: href,
		})
	}

	return emotes
}

func parseHtmls(urls []string) []htmlEmote {
	emotes := []htmlEmote{}
	ch := make(chan []htmlEmote)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string, ch chan<- []htmlEmote, wg *sync.WaitGroup) {
			defer wg.Done()
			parsedEmotes := parseHtml(url)
			ch <- parsedEmotes
		}(url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for emote := range ch {
		emotes = append(emotes, emote...)
	}
	return emotes
}

type imageEmote struct {
	name  string
	href  string
	image io.ReadCloser
}

func downloadImages(emotes []htmlEmote) []imageEmote {
	images := []imageEmote{}
	ch := make(chan imageEmote)

	var wg sync.WaitGroup
	for _, emote := range emotes {
		wg.Add(1)
		go func(emote htmlEmote, ch chan<- imageEmote, wg *sync.WaitGroup) {
			defer wg.Done()

			indexed := imageEmote{
				href: emote.href,
				name: emote.name,
			}

			imageBody, err := getContent(emote.href)
			if err != nil {
				log.Fatal(err)
				indexed.image = nil
				ch <- indexed
			}
			indexed.image = imageBody
			ch <- indexed
		}(emote, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for indexed := range ch {
		images = append(images, indexed)
	}

	return images
}

func writeZip(emotes []imageEmote, output string) {
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	z := zip.NewWriter(f)
	defer z.Close()

	for _, emote := range emotes {
		filename := url.PathEscape(fmt.Sprintf("%s.png", emote.name))
		file, err := z.Create(filename)
		if err != nil {
			log.Fatal(err)
			continue
		}
		imageBody := emote.image
		if imageBody == nil {
			continue
		}

		if _, err := io.Copy(file, imageBody); err != nil {
			log.Fatal(err)
		}
		imageBody.Close()
	}
}
