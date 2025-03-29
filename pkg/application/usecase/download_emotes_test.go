package usecase

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadEmotesUseCaseWithFailedDownload(t *testing.T) {
	fakeDownloader := fakeDownloader{shouldError: true}
	fakeMapper := fakeMapper{shouldError: true}

	usecase := NewDownloadEmotesUseCase(fakeDownloader.download, fakeMapper)

	actual, err := usecase.Execute(url.URL{})

	assert.Error(t, err)
	assert.Empty(t, actual)
}

func TestDownloadEmotesUseCaseWithFailedMapper(t *testing.T) {
	fakeDownloader := fakeDownloader{shouldError: false}
	fakeMapper := fakeMapper{shouldError: true}

	usecase := NewDownloadEmotesUseCase(fakeDownloader.download, fakeMapper)

	actual, err := usecase.Execute(url.URL{})

	assert.Error(t, err)
	assert.Empty(t, actual)
}

func TestDownloadEmotesUseCase(t *testing.T) {
	fakeDownloader := fakeDownloader{shouldError: false}
	fakeMapper := fakeMapper{shouldError: false}

	usecase := NewDownloadEmotesUseCase(fakeDownloader.download, fakeMapper)

	actual, err := usecase.Execute(url.URL{})

	assert.Nil(t, err)
	assert.Empty(t, actual)
}
