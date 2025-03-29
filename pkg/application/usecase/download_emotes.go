package usecase

import (
	"net/http"
	"net/url"

	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/mapper"
)

type DownloadEmotesUseCase struct {
	download        func(url.URL) (*http.Response, error)
	htmlEmoteMapper mapper.HtmlEmoteMapper
}

func NewDownloadEmotesUseCase(
	download func(url.URL) (*http.Response, error),
	htmlEmoteMapper mapper.HtmlEmoteMapper,
) DownloadEmotesUseCase {
	return DownloadEmotesUseCase{
		download,
		htmlEmoteMapper,
	}
}

func (u DownloadEmotesUseCase) Execute(url url.URL) ([]entity.Emote, error) {
	res, err := u.download(url)
	if err != nil {
		return []entity.Emote{}, err
	}
	defer res.Body.Close()

	return u.htmlEmoteMapper.ToEmotes(res.Body)
}
