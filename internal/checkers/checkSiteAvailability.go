package checker

import (
	"net/http"
	"time"
)

// CheckSiteAvailability ...
func CheckSiteAvailability(url string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
