package core

import (
	"fmt"
	"net/http"
	"time"
)

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
