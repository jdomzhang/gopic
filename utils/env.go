package utils

import (
	"bytes"
	"os/exec"
	"runtime"
)

func GetFFmpegExec() []string {
	var platform = runtime.GOOS
	var command = []string{"", "ffmpeg"}

	switch platform {
	case "windows":
		command[0] = "where"
	default:
		command[0] = "which"
	}

	return command
}

func GetFFprobeExec() []string {
	var platform = runtime.GOOS
	var command = []string{"", "ffprobe"}

	switch platform {
	case "windows":
		command[0] = "where"
	default:
		command[0] = "which"
	}
	return command
}

func LineSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

// TestCmd ...
func TestCmd(command string, args string) (bytes.Buffer, error) {
	var out bytes.Buffer

	cmd := exec.Command(command, args)

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return out, err
	}

	return out, nil
}
