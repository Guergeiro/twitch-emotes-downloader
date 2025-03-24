package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeBody struct{}

func (b fakeBody) Read(p []byte) (int, error) {
	return 0, nil
}

func (b fakeBody) Close() error {
	return nil
}

func TestCreatingNewImage(t *testing.T) {
	contentType := "contentType"
	fakeBody := fakeBody{}

	actual := NewImage(fakeBody, contentType)

	assert.Equal(t, contentType, actual.ContentType())
	assert.Equal(t, fakeBody, actual.Body())
}
