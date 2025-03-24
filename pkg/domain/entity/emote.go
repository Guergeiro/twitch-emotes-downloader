package entity

import "net/url"

type Emote struct {
	name  string
	href  url.URL
	image *Image
}

type EmoteOption func(*Emote)

func (e Emote) Name() string {
	return e.name
}

func (e Emote) Href() url.URL {
	return e.href
}

func (e *Emote) SetImage(image *Image) {
	e.image = image
}

func (e Emote) Image() *Image {
	return e.image
}

func WithImage(image *Image) EmoteOption {
	return func(e *Emote) {
		e.SetImage(image)
	}
}

func NewEmote(name string, href url.URL, opts ...EmoteOption) Emote {
	emote := Emote{
		name: name,
		href: href,
	}

	for _, opt := range opts {
		opt(&emote)
	}

	return emote
}
