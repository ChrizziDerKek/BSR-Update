package agent

import (
	"os"
	"path/filepath"
)

func CleanupPostInstallRoot(source string) error {
	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(source, entry.Name())
		destPath := entry.Name()

		err := os.Rename(srcPath, destPath)
		if err != nil {
			return err
		}
	}

	err = os.Remove("update.zip")
	if err != nil {
		return err
	}

	if *IsServerUpdate {
		err = os.Remove("BSR_Client.exe")
		if err != nil {
			return err
		}

		err = os.Remove("sounds")
		if err != nil {
			return err
		}

	} else {
		err = os.Remove("BSR_Server.exe")
		if err != nil {
			return err
		}
	}

	return os.Remove(source)
}

func CleanupPreInstallRoot() error {
	var err error
	if *IsServerUpdate {
		err = os.Remove("BSR_Server.exe")
		if err != nil {
			return err
		}

	} else {
		err = os.Remove("BSR_Client.exe")
		if err != nil {
			return err
		}

		err = os.RemoveAll("sounds")
		if err != nil {
			return err
		}

	}

	return err
}
