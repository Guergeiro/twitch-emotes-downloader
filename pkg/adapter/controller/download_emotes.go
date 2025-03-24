package controller

import (
	"net/url"
	"sync"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/application/usecase"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type DownloadEmotesController struct {
	downloadEmotesUseCase usecase.DownloadEmotesUseCase
	downloadImageUseCase  usecase.DownloadImageUseCase
	writeZipUseCase       usecase.WriteZipUseCase
}

func NewDownloadEmotesController(
	downloadEmotesUseCase usecase.DownloadEmotesUseCase,
	downloadImageUseCase usecase.DownloadImageUseCase,
	writeZipUseCase usecase.WriteZipUseCase,
) DownloadEmotesController {
	return DownloadEmotesController{
		downloadEmotesUseCase,
		downloadImageUseCase,
		writeZipUseCase,
	}
}

func (c DownloadEmotesController) Handle(
	hrefs []string,
	output string,
) error {

	channel :=
		c.startDownloadImages(
			c.startDownloadEmotes(
				c.startParseHrefs(hrefs),
			),
		)

	emotes := []entity.Emote{}
	for emote := range channel {
		emotes = append(emotes, emote)
	}

	return c.writeZipUseCase.Execute(output, emotes)
}

func (c DownloadEmotesController) startParseHrefs(hrefs []string) <-chan url.URL {
	out := make(chan url.URL)

	go func() {
		defer close(out)
		for _, href := range hrefs {
			if u, err := url.Parse(href); err == nil {
				out <- *u
			}
		}
	}()

	return out
}

func (c DownloadEmotesController) startDownloadEmotes(channel <-chan url.URL) <-chan []entity.Emote {
	out := make(chan []entity.Emote)

	go func() {
		defer close(out)
		for u := range channel {
			if emotes, err := c.downloadEmotesUseCase.Execute(u); err == nil {
				out <- emotes
			}
		}
	}()

	return out
}

func (c DownloadEmotesController) startDownloadImages(channel <-chan []entity.Emote) <-chan entity.Emote {
	out := make(chan entity.Emote)

	go func() {
		defer close(out)
		for emotes := range channel {
			for emote := range c.downloadImage(emotes) {
				out <- emote
			}
		}
	}()

	return out
}

func (c DownloadEmotesController) downloadImage(emotes []entity.Emote) <-chan entity.Emote {
	channel := make(chan entity.Emote)
	wg := sync.WaitGroup{}

	for _, emote := range emotes {
		wg.Add(1)
		go func(e entity.Emote, wg *sync.WaitGroup) {
			defer wg.Done()
			if image, err := c.downloadImageUseCase.Execute(e); err == nil {
				channel <- entity.NewEmote(
					e.Name(),
					e.Href(),
					entity.WithImage(&image),
				)
			}
		}(emote, &wg)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	return channel
}
