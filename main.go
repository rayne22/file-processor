package main

import (
	"github.com/rayne22/image-processor/process"
)

func main()  {
	r := process.UploadedImage{
		Name:   "123.png",
		Width:  500,
		Height: 500,
	}

	r.ResizeImage()
}
