package modules

import (
	"fmt"
	"os/exec"
)

const (
	cliToolPath    = "C:\\Program Files (x86)\\ISO to WBFS\\wbfs_file.exe"
	convertDirPath = "F:\\wbfs"
)

func ConvertToWBFS(info *DownloadInfo) {
	fullPath := info.LastFilePath

	logName := "[ISO TO WBFS] "

	err := cliToWBFS(fullPath)

	if err != nil {
		fmt.Println(logName+info.Name+" Error converting file:", err)
	} else {
		fmt.Println(logName+"Successfully converted file:", info.Name)
		Remove(fullPath)
	}
}

func cliToWBFS(file string) error {
	cmd := exec.Command(cliToolPath, file, "convert", convertDirPath)

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(string(stdoutStderr))

	return nil
}
