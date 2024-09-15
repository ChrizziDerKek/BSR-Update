package agent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func GetLatestRelease(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	if len(release.Assets) == 0 {
		return "", fmt.Errorf("no assets found in the latest release")
	}

	var correctPath string

	if strings.Contains(release.Assets[0].BrowserDownloadURL, "Release.zip") {
		correctPath = release.Assets[0].BrowserDownloadURL
	} else if strings.Contains(release.Assets[1].BrowserDownloadURL, "Release.zip") {
		correctPath = release.Assets[1].BrowserDownloadURL
	} else {
		panic("No release zip available in release! Please raise a github issue to fix this!")
	}

	return correctPath, nil
}
