package usecase

import (
	"archive/zip"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type WriteZipUseCase struct{}

func NewWriteZipUseCase() WriteZipUseCase {
	return WriteZipUseCase{}
}

func (u WriteZipUseCase) Execute(output string, emotes []entity.Emote) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	z := zip.NewWriter(f)
	defer z.Close()

	for _, emote := range emotes {
		filename := url.PathEscape(fmt.Sprintf("%s.png", emote.Name()))
		file, err := z.Create(filename)
		if err != nil {
			return err
		}
		imageBody := emote.Image().Body()

		if _, err := io.Copy(file, imageBody); err != nil {
			return err
		}
		imageBody.Close()
	}

	return nil
}
