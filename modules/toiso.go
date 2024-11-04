package modules

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	logName = "[RVZ TO ISO] "
)

func CovertRVZToISO(info *DownloadInfo) {

	fullPath := info.LastFilePath
	fileName := filepath.Base(fullPath)

	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	fmt.Printf(logName+"Converting %s\n", name)

	// info.TaskStatus <- fmt.Sprintf(logName+"Converting %s", name)

	// Rvz To ISO
	outputPath := filepath.Join(Config.IsoDirPath, name+".iso")

	info.LastFilePath = outputPath

	err := cliRvzToISO(fullPath, outputPath)

	if err != nil {
		fmt.Println(logName+"Error converting file:", err)
		//info.TaskStatus <- logName + "Error converting file"
	} else {
		fmt.Println(logName+"Successfully converted file:", fullPath)
		//info.TaskStatus <- fmt.Sprintf(logName+"Successfully converted file %s", name)
	}

	Remove(fullPath)
}

func cliRvzToISO(inputFilePath string, outputPath string) error {
	fmt.Printf("%sWriting %s\n", logName, outputPath)

	cmd := exec.Command(Config.DolphinToolPath, "convert",
		"--format=iso",
		"--input="+inputFilePath,
		"--output="+outputPath)

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(logName + string(stdoutStderr))

	return nil
}
