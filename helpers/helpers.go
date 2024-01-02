package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {

	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	// PREVENT COMMON domain error url types
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}

func getDomain() string {
	domain := ""

	domain = os.Getenv("DOMAIN")
	if os.Getenv("PORT") == "8080" {
		domain = os.Getenv("LOCAL_DOMAIN")
	}
	return domain
}
