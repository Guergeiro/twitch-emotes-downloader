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

func (m GoQueryHtmlEmoteMapper) ToEmotes(html io.ReadCloser, size string) ([]entity.Emote, error) {
	emotes := []entity.Emote{}
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return emotes, err
	}

	normalizedSize := normalizeEmoteSize(size)
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
		u.Path = updateEmoteSize(u.Path, normalizedSize)
		emotes = append(emotes, entity.NewEmote(single.Text(), *u))
	}

	return emotes, nil
}

func normalizeEmoteSize(size string) string {
	switch strings.TrimSpace(size) {
	case "1.0", "2.0", "3.0":
		return size
	default:
		return "3.0"
	}
}

func updateEmoteSize(path string, size string) string {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) == 0 {
		return size
	}

	lastIndex := len(segments) - 1
	if lastIndex >= 0 && (segments[lastIndex] == "1.0" || segments[lastIndex] == "2.0" || segments[lastIndex] == "3.0") {
		segments[lastIndex] = size
	} else {
		segments = append(segments, size)
	}

	return "/" + strings.Join(segments, "/")
}
