package core

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
)

func ReadURLsFromJSON(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var urls []string
	if err := json.Unmarshal(data, &urls); err != nil {
		return nil, err
	}

	return urls, nil
}

func ReadURLsFromCSV(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, record := range records {
		urls = append(urls, record[0])
	}

	return urls, nil
}

func ReadURLsFromText(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := string(data)
	urls := []string{}

	for _, line := range strings.Split(lines, "\n") {
		urls = append(urls, line)
	}

	return urls, nil
}
