# gopic

ffmpeg wrapper for extract video frame to image

# Example Code
```go
package main

import (
	"fmt"

	"github.com/jdomzhang/gopic"
)

const (
	sampleInputFile     = "./sample/input/sample.mp4"
	sampleOutputJpgFile = "./sample/output/sample.jpg"
	sampleOutputPngFile = "./sample/output/sample.png"
)

func main() {
	pic, err := gopic.NewPic()
	if err != nil {
		panic(err)
	}

	second := "1"
	// second := "00:00:01"
	err = pic.Extract(sampleInputFile, second, sampleOutputJpgFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
```

# Thanks
- goffmpeg: `https://github.com/xfrr/goffmpeg`
