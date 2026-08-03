package mapper

import (
	"io"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type HtmlEmoteMapper interface {
	ToEmotes(html io.ReadCloser, size string) ([]entity.Emote, error)
}
