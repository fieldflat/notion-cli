package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"notion-cli/config"
)

func GetHTTPHandler(endpoint string, body io.Reader) *http.Request {
	token := os.Getenv("NOTION_INTEGRATION_TOKEN")

	req, err := http.NewRequest("GET", endpoint, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in GetHTTPHandler")
		fmt.Fprintln(os.Stderr, "endpoint: "+endpoint)
		os.Exit(1)
	}
	req.Header.Set("Notion-Version", config.NotionVersion)
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}
