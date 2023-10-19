package core

import (
	"fmt"
	"net/http"
	"time"
)

func ProbeURL(url string) error {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	elapsed := time.Since(startTime)

	fmt.Printf("[+] URL: %s\n", url)
	fmt.Printf("    - Status Code: %d\n", resp.StatusCode)
	fmt.Printf("    - Response Time: %s\n", elapsed)

	return nil
}

func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No URLs provided -> usage: url-probe <url1> <url2>")
	}

	for _, url := range args {
		err := ProbeURL(url)
		if err != nil {
			fmt.Printf("ERROR probing URL %s: %v\n", url, err)
		}
	}

	return nil
}
