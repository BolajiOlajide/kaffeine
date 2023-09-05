package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLs_Set(t *testing.T) {
	var urls URLs

	// Test appending a valid url
	err := urls.Set("http://example.com")
	assert.NoError(t, err)
	assert.Equal(t, urls.Len(), 1)

	// Test appending multiple urls
	err = urls.Set("https://test.com")
	assert.NoError(t, err)
	assert.Equal(t, urls.Len(), 2)
}

func TestURLs_String(t *testing.T) {
	urls := URLs{"http://example.com", "https://test.com"}

	// Test String() joins urls correctly
	assert.Equal(t, "http://example.com, https://test.com", urls.String())
}

func TestURLs_Validate(t *testing.T) {
	var urls URLs

	// Test empty urls
	err := urls.Validate()
	assert.Error(t, err)

	// Test invalid scheme
	urls = URLs{"ftp://test.com"}
	err = urls.Validate()
	assert.Error(t, err)

	// Test missing host
	urls = URLs{"https://"}
	err = urls.Validate()
	assert.Error(t, err)

	// Test valid urls
	urls = URLs{"http://example.com", "https://test.com"}
	err = urls.Validate()
	assert.NoError(t, err)
}
