package entity

import "io"

type Image struct {
	body        io.ReadCloser
	contentType string
}

func (e Image) Body() io.ReadCloser {
	return e.body
}

func (e Image) ContentType() string {
	return e.contentType
}

func NewImage(body io.ReadCloser, contentType string) Image {
	return Image{
		body:        body,
		contentType: contentType,
	}
}
