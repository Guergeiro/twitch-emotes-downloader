package e2e

import (
	"archive/zip"
	"log"
	"testing"

	"github.com/guergeiro/twitch-emotes-downloader/internal/http"
	"github.com/guergeiro/twitch-emotes-downloader/internal/mapper"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/adapter/controller"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/application/usecase"
	"github.com/stretchr/testify/assert"
)

func TestDownloadEmotesController(t *testing.T) {
	if testing.Short() {
		t.Skip("short test")
	}

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

	hrefs := []string{"https://www.twitchmetrics.net/emotes"}
	output := "output.zip"

	err := c.Handle(hrefs, output)

	assert.Nil(t, err)

	zip, err := zip.OpenReader(output)
	assert.Nil(t, err)
	defer zip.Close()

	containsKappa := false
	for _, f := range zip.File {
		log.Println(f.Name)
		if f.Name == "Kappa.png" {
			containsKappa = true
		}
	}
	assert.True(t, containsKappa)
}
