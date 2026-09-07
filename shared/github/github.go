package github

import (
	"net/http"
	"os"
)

// token holds an optional GitHub API token set programmatically.
var token string

// SetToken sets the GitHub API token used by Get.
func SetToken(t string) { token = t }

// Get performs an HTTP GET request to the given URL.
// If a token was set via SetToken or the GITHUB_TOKEN environment variable
// is set, the request includes an Authorization header to avoid API rate limits.
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	t := token
	if t == "" {
		t = os.Getenv("GITHUB_TOKEN")
	}
	if t != "" {
		req.Header.Set("Authorization", "Bearer "+t)
	}

	return http.DefaultClient.Do(req)
}
