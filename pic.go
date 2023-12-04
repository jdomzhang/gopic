package gopic

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
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
	if err := obj.CheckVideoFile(videoFile); err != nil {
		return err
	}

	// check output folder existing
	// create file if not existing
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		// create folder first
		baseFolder := path.Dir(outputFile)
		if _, err := os.Stat(baseFolder); os.IsNotExist(err) {
			if err := os.MkdirAll(baseFolder, 0755); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// create file
		f, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		f.Close()
	} else if err != nil {
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

func (obj *Pic) CheckVideoFile(videoFile string) error {
	isOnlineFile := strings.HasPrefix(videoFile, "http://") || strings.HasPrefix(videoFile, "https://")
	if isOnlineFile {
		// check online file available
		resp, err := http.Get(videoFile)
		if err != nil {
			return err
		}

		// status code 2xx
		if resp.StatusCode/100 != 2 {
			return fmt.Errorf("url: %s | error, code: %d", videoFile, resp.StatusCode)
		}
	} else {
		_, err := os.Stat(videoFile)
		if err != nil {
			return err
		}
	}

	return nil
}
