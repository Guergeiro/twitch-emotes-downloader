package usecase

import (
	"net/http"
	"net/url"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type DownloadImageUseCase struct {
	download func(url.URL) (*http.Response, error)
}

func NewDownloadImageUseCase(
	download func(url.URL) (*http.Response, error),
) DownloadImageUseCase {
	return DownloadImageUseCase{
		download,
	}
}

func (u DownloadImageUseCase) Execute(emote entity.Emote) (entity.Image, error) {
	res, err := u.download(emote.Href())
	if err != nil {
		return entity.Image{}, err
	}

	contentType := res.Header.Get("content-type")

	if contentType == "" {
		contentType = "image/png"
	}

	return entity.NewImage(res.Body, contentType), nil
}
