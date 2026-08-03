package mapper

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoQueryHtmlEmoteMapperUsesRequestedSize(t *testing.T) {
	html := `<div><div><img src="https://static-cdn.jtvnw.net/emoticons/v2/test/default/dark/3.0" /></div><samp>test</samp></div>`

	m := NewGoQueryHtmlEmoteMapper()
	emotes, err := m.ToEmotes(io.NopCloser(strings.NewReader(html)), "2.0")

	assert.NoError(t, err)
	assert.Len(t, emotes, 1)
	href := emotes[0].Href()
	assert.Contains(t, href.String(), "/default/")
	assert.Contains(t, href.String(), "/2.0")
	assert.NotContains(t, href.String(), "/3.0")
}
