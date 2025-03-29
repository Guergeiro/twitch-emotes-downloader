package usecase

import (
	"log"
	"net/url"
	"testing"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestDownloadImageUseCaseWithFailedDownload(t *testing.T) {
	fakeDownloader := fakeDownloader{shouldError: true}

	entity := entity.NewEmote("some name", url.URL{})

	usecase := NewDownloadImageUseCase(fakeDownloader.download)

	actual, err := usecase.Execute(entity)

	assert.Error(t, err)
	assert.Empty(t, actual)
}

func TestDownloadImageUseCaseNoContentType(t *testing.T) {
	fakeDownloader := fakeDownloader{shouldError: false}

	entity := entity.NewEmote("some name", url.URL{})

	usecase := NewDownloadImageUseCase(fakeDownloader.download)

	actual, err := usecase.Execute(entity)
	log.Println(actual)

	assert.Nil(t, err)
	assert.Equal(t, "image/png", actual.ContentType())
}

func TestDownloadImageUseCaseWithContentType(t *testing.T) {
	fakeDownloader := fakeDownloader{shouldError: false, contentType: "text/html"}

	entity := entity.NewEmote("some name", url.URL{})

	usecase := NewDownloadImageUseCase(fakeDownloader.download)

	actual, err := usecase.Execute(entity)
	log.Println(actual)

	assert.Nil(t, err)
	assert.Equal(t, "text/html", actual.ContentType())
}
