package tests

import (
	"github.com/kenjitheman/url-probe/core"
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	testArgs := []string{"-source", "json", "-file", "test.json", "url1", "url2"}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = append([]string{"programName"}, testArgs...)

	parsedArgs, err := core.ParseArgs()

	if err != nil {
		t.Fatalf("ParseArgs returned an error: %v", err)
	}

	if parsedArgs.Source != "json" {
		t.Errorf("Expected Source to be 'json', but got '%s'", parsedArgs.Source)
	}

	if parsedArgs.File != "test.json" {
		t.Errorf("Expected File to be 'test.json', but got '%s'", parsedArgs.File)
	}

	if len(parsedArgs.URLs) != 2 {
		t.Errorf("Expected 2 URLs, but got %d", len(parsedArgs.URLs))
	} else {
		if parsedArgs.URLs[0] != "url1" {
			t.Errorf("Expected URL 1 to be 'url1', but got '%s'", parsedArgs.URLs[0])
		}
		if parsedArgs.URLs[1] != "url2" {
			t.Errorf("Expected URL 2 to be 'url2', but got '%s'", parsedArgs.URLs[1])
		}
	}
}
