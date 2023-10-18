package core

import (
	"fmt"
	"net/http"
	"time"
)

func TestURLs(urls []string) {
	for _, url := range urls {
		if err := testURL(url); err != nil {
			fmt.Println("[!] Error testing URL:", url)
			fmt.Println("  -- [?] error:", err)
		}
	}
}

func testURL(url string) error {
	startTime := time.Now()
	resp, err := http.Get(url)
	elapsedTime := time.Since(startTime)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("[+] URL:", url)
	} else {
		fmt.Println("[!] URL:", url)
	}

	fmt.Println("  -- [?] status:", resp.Status)
	fmt.Println("  -- [?] response time:", elapsedTime)

	return nil
}
