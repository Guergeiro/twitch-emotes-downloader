package usecase

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"mime"
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
		extension, err := mime.ExtensionsByType(emote.Image().ContentType())
		if err != nil {
			return err
		}
		if extension == nil || len(extension) == 0 {
			return errors.New("no extension available for the content type")
		}
		filename := url.PathEscape(
			fmt.Sprintf("%s%s", emote.Name(), extension[0]),
		)
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
