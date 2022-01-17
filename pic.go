package gopic

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jdomzhang/gopic/utils"
)

type Pic struct {
	ffmpeg  string
	ffprobe string
}

func NewPic() (*Pic, error) {
	var outFFmpeg bytes.Buffer
	var outProbe bytes.Buffer

	execFFmpegCommand := utils.GetFFmpegExec()
	execFFprobeCommand := utils.GetFFprobeExec()

	outFFmpeg, err := utils.TestCmd(execFFmpegCommand[0], execFFmpegCommand[1])
	if err != nil {
		return nil, err
	}

	outProbe, err = utils.TestCmd(execFFprobeCommand[0], execFFprobeCommand[1])
	if err != nil {
		return nil, err
	}

	ffmpeg := strings.Replace(strings.Split(outFFmpeg.String(), "\n")[0], utils.LineSeparator(), "", -1)
	ffprobe := strings.Replace(strings.Split(outProbe.String(), "\n")[0], utils.LineSeparator(), "", -1)

	return &Pic{ffmpeg, ffprobe}, nil
}

func (obj *Pic) Extract(videoFile string, second string, outputFile string) error {
	_, err := os.Stat(videoFile)
	if err != nil {
		return err
	}

	// set default second
	if second == "" {
		second = "1"
	}

	var (
		outb, errb bytes.Buffer
	)

	command := []string{"-y", "-i", videoFile, "-ss", second, "-vframes", "1", outputFile}
	cmd := exec.Command(obj.ffmpeg, command...)
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command: (%s) | error: %s | message: %s %s", command, err, outb.String(), errb.String())
	}

	return nil
}
