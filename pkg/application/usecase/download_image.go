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

	return entity.NewImage(res.Body, res.Header.Get("content-type")), nil
}
