package mapper

import (
	"io"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/guergeiro/twitch-emotes-downloader/pkg/domain/entity"
)

type GoQueryHtmlEmoteMapper struct{}

func NewGoQueryHtmlEmoteMapper() GoQueryHtmlEmoteMapper {
	return GoQueryHtmlEmoteMapper{}
}

func (m GoQueryHtmlEmoteMapper) ToEmotes(html io.ReadCloser) ([]entity.Emote, error) {
	emotes := []entity.Emote{}
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return emotes, err
	}

	selection := doc.Find("samp")
	for i := range selection.Nodes {
		single := selection.Eq(i)
		img := single.Prev().Find("img")
		href, exists := img.Attr("src")
		if exists == false {
			continue
		}
		u, err := url.Parse(strings.ReplaceAll(href, "/static/", "/default/"))
		if err != nil {
			continue
		}
		emotes = append(emotes, entity.NewEmote(single.Text(), *u))
	}

	return emotes, nil
}
