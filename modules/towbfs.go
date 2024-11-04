package modules

import (
	"fmt"
	"os/exec"
)

const (
	_logName = "[ISO TO WBFS] "
)

func ConvertToWBFS(info *DownloadInfo) {
	fullPath := info.LastFilePath

	// fmt.Println(_logName + fullPath)

	fmt.Printf("%sWorking on %s\n", _logName, info.Name)

	// info.TaskStatus <- fmt.Sprintf(_logName+"Working on %s", info.Name)

	err := cliToWBFS(fullPath)

	if err != nil {
		//info.TaskStatus <- _logName + info.Name + " Error Converting"
		fmt.Println(_logName+info.Name+" Error converting file:", err)
	} else {
		fmt.Println(_logName+"Successfully converted file:", info.Name)
		// info.TaskStatus <- _logName + "Successfully converted file: " + info.Name
		Remove(fullPath)
	}
}

func cliToWBFS(file string) error {
	cmd := exec.Command(Config.ISOtoWBFSPath, file, "convert", Config.WBFSDirPath)

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(_logName + string(stdoutStderr))

	return nil
}
