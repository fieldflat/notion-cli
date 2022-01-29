package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"notion-cli/config"
)

func GetHTTPRequester(endpoint string, body io.Reader) *http.Request {
	token := os.Getenv("NOTION_INTEGRATION_TOKEN")

	req, err := http.NewRequest("GET", endpoint, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in GetHTTPRequester")
		fmt.Fprintln(os.Stderr, "endpoint: "+endpoint)
		os.Exit(1)
	}
	req.Header.Set("Notion-Version", config.NotionVersion)
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}

func PostHTTPRequester(endpoint string, body io.Reader) *http.Request {
	token := os.Getenv("NOTION_INTEGRATION_TOKEN")

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in PostHTTPRequester")
		fmt.Fprintln(os.Stderr, "endpoint: "+endpoint)
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", config.NotionVersion)
	req.Header.Add("Authorization", "Bearer "+token)

	return req
}

func PatchHTTPRequester(endpoint string, body io.Reader) *http.Request {
	token := os.Getenv("NOTION_INTEGRATION_TOKEN")

	req, err := http.NewRequest("PATCH", endpoint, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in PatchHTTPRequester")
		fmt.Fprintln(os.Stderr, "endpoint: "+endpoint)
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", config.NotionVersion)
	req.Header.Add("Authorization", "Bearer "+token)

	return req
}
