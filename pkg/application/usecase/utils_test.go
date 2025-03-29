package usecase

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type fakeBody struct{}

func (b fakeBody) Read(p []byte) (int, error) {
	return 0, nil
}

func (b fakeBody) Close() error {
	return nil
}

type fakeDownloader struct {
	shouldError bool
	contentType string
}

func (d fakeDownloader) download(u url.URL) (*http.Response, error) {
	if d.shouldError {
		return nil, errors.New("some error occurred")
	}
	return &http.Response{
		Body: fakeBody{},
		Header: http.Header{
			"Content-Type": []string{d.contentType},
		},
	}, nil
}

type fakeMapper struct {
	shouldError bool
}

func (m fakeMapper) ToEmotes(body io.ReadCloser) ([]entity.Emote, error) {
	if m.shouldError {
		return []entity.Emote{}, errors.New("some error occurred")
	}
	return []entity.Emote{}, nil
}
