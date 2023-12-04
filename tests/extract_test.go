package tests

import (
	"testing"

	gopic "github.com/jdomzhang/gopic"
	"github.com/stretchr/testify/assert"
)

const (
	sampleInputFile     = "../sample/input/sample.mp4"
	sampleOutputJpgFile = "../sample/output/sample.jpg"
	sampleOutputPngFile = "../sample/output/sample.png"
)

func TestEnv(t *testing.T) {
	_, err := gopic.NewPic()
	assert.NoError(t, err)
}

func TestExtractJpg(t *testing.T) {
	pic, err := gopic.NewPic()
	assert.NoError(t, err)

	err = pic.Extract(sampleInputFile, "1", sampleOutputJpgFile)
	assert.NoError(t, err)
}

func TestExtractPng(t *testing.T) {
	pic, err := gopic.NewPic()
	assert.NoError(t, err)

	err = pic.Extract(sampleInputFile, "1", sampleOutputPngFile)
	assert.NoError(t, err)
}

func Test_ExtractJpg_From_URL(t *testing.T) {
	pic, err := gopic.NewPic()
	assert.NoError(t, err)

	err = pic.Extract("https://cdn.edesoft.com/youzheng/upload/202312/04/132113.mp4", "15", sampleOutputJpgFile)
	assert.NoError(t, err)
}

func Test_OnlineURL(t *testing.T) {
	pic, err := gopic.NewPic()
	assert.NoError(t, err)

	err = pic.CheckVideoFile("https://cdn.edesoft.com/youzheng/upload/202312/04/132113.mp4")
	assert.NoError(t, err)
}
