package modules

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	rvzDirPath      = "E:\\ssdwiigames"
	dolphinToolPath = "C:\\Users\\chrsh\\OneDrive\\Desktop\\Dolphin-x64\\DolphinTool.exe"
)

func CovertRVZToISO(info *DownloadInfo) {
	logName := "[RVZ TO ISO] "

	fullPath := info.LastFilePath
	fileName := filepath.Base(fullPath)

	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	fmt.Printf(logName+"Converting %s\n", name)

	// Rvz To ISO
	outputPath := filepath.Join(rvzDirPath, name+".iso")

	info.LastFilePath = outputPath

	err := cliRvzToISO(fullPath, outputPath)

	if err != nil {
		fmt.Println(logName+"Error converting file:", err)
	} else {
		fmt.Println(logName+"Successfully converted file:", fullPath)
		Remove(fullPath)
	}

}

func cliRvzToISO(inputFilePath string, outputPath string) error {
	cmd := exec.Command(dolphinToolPath, "convert",
		"--format=iso",
		"--input="+inputFilePath,
		"--output="+outputPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// Create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure it's called to release resources

	// Start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Create scanners to read from stdout and stderr
	stdoutScanner := bufio.NewScanner(stdout)
	stderrScanner := bufio.NewScanner(stderr)

	// Read from stdout in a separate goroutine
	go func() {
		for stdoutScanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(stdoutScanner.Text()) // Print stdout line
				time.Sleep(2 * time.Second)       // Delay for 2 seconds
			}
		}
	}()

	// Read from stderr in a separate goroutine
	go func() {
		for stderrScanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Fprintln(os.Stderr, stderrScanner.Text()) // Print stderr line
				time.Sleep(2 * time.Second)                   // Delay for 2 seconds
			}
		}
	}()

	// Wait for the command to finish
	err = cmd.Wait()

	// Cancel the context to stop the goroutines
	cancel()

	// Check for any errors
	if err != nil {
		return err
	}

	return nil
}
