package main

import (
	"github.com/rayne22/image-processor/process"
)

func main()  {
	r := process.UploadedImage{}

	r.name = "test.jpg"
	r.width = 500
	r.height = 500

	r.ResizeImage()
}
