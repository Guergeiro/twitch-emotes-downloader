package entity

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingNewEmoteWithoutOpts(t *testing.T) {
	name := "foo"
	href, err := url.Parse("https://brenosalles.com")
	assert.Nil(t, err)

	actual := NewEmote(name, *href)
	assert.Equal(t, name, actual.Name())
	assert.Equal(t, *href, actual.Href())
}

func TestCreatingNewEmoteWithOpts(t *testing.T) {
	name := "foo"
	href, err := url.Parse("https://brenosalles.com")
	image := &Image{}

	assert.Nil(t, err)

	actual := NewEmote(name, *href, WithImage(image))
	assert.Equal(t, name, actual.Name())
	assert.Equal(t, *href, actual.Href())
	assert.Equal(t, image, actual.Image())
}
