package mapper

import (
	"io"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type HtmlEmoteMapper interface {
	ToEmotes(html io.ReadCloser) ([]entity.Emote, error)
}
