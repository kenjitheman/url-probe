package tests

import (
	"github.com/kenjitheman/url-probe/core"
	"os"
	"testing"
)

func TestReadURLsFromJSON(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	data := []byte(`["url1", "url2"]`)
	_, err = tempFile.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	urls, err := core.ReadURLsFromJSON(tempFile.Name())
	if err != nil {
		t.Fatalf("ReadURLsFromJSON returned an error: %v", err)
	}

	expectedURLs := []string{"url1", "url2"}
	if len(urls) != len(expectedURLs) {
		t.Errorf("Expected %d URLs, but got %d", len(expectedURLs), len(urls))
	}

	for i, url := range urls {
		if url != expectedURLs[i] {
			t.Errorf("Expected URL %d to be '%s', but got '%s'", i, expectedURLs[i], url)
		}
	}
}
