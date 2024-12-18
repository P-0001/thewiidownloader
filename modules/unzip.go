package modules

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func UnzipTask(info *DownloadInfo) {
	fmt.Printf("Unzipping Processing: %s\n", info.Name)

	//info.TaskStatus <- fmt.Sprintf("Unzipping %s", info.Name)

	zipFilePath := info.LastFilePath

	err := unzip(zipFilePath, Config.RvzDirPath, info)

	if err != nil {
		fmt.Printf("[%s] Error during unzip: %s\n", info.Name, err)
		return
	}

	Remove(zipFilePath)

	//	info.TaskStatus <- fmt.Sprintf("Unzipped successfully: %s", info.Name)

	fmt.Printf("Unzipped successfully: %s\n", info.Name)

}

func unzip(src string, dest string, info *DownloadInfo) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		fpath := filepath.Join(dest, file.Name)

		// Check for directory or create parent directories
		if file.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		info.LastFilePath = fpath

		// Create directories for files
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		// Open destination file
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		// Copy the file content
		_, err = io.Copy(outFile, rc)

		// Close the source and destination files
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}

		// Update the modification time (optional)
		mtime := file.Modified
		err = os.Chtimes(fpath, time.Now(), mtime)
		if err != nil {
			return err
		}
	}
	return nil
}
