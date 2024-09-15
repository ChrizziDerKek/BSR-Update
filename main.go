package main

import (
	"bsr/update-agent/agent"
	"fmt"
	"time"
)

func main() {

	agent.GetArguments()

	err := agent.KillProcess()
	if err != nil {
		panic("Error killing process:" + err.Error())
	}

	time.Sleep(2 * time.Second)

	err = agent.CleanupPreInstallRoot()
	if err != nil {
		panic("Error cleaning up root PRE-INST:" + err.Error())
	}

	apiURL := "https://api.github.com/repos/chrizziderkek/bugshotroulette/releases/latest"
	zipPath := "update.zip"
	destPath := "extracted"

	downloadURL, err := agent.GetLatestRelease(apiURL)
	if err != nil {
		panic("Error getting latest release:" + err.Error())
	}

	err = agent.DownloadFile(downloadURL, zipPath)
	if err != nil {
		panic("Error downloading file:" + err.Error())
	}

	err = agent.UnzipSource(zipPath, destPath)
	if err != nil {
		panic("Error unzipping file:" + err.Error())
	}

	err = agent.CleanupPostInstallRoot(destPath)
	if err != nil {
		panic("Error cleaning up root POST-INST:" + err.Error())
	}

	err = agent.OpenProcess()
	if err != nil {
		panic("Error opening process:" + err.Error())
	}

	fmt.Println("Update finished.")
}
