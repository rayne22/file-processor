package main

import (
	"github.com/rayne22/image-processor/process"
)

func main()  {
	r := process.UploadedImage{
		Name:   "123.PNG",
		Width:  0,
		Height: 0,
		Path: "temp",
	}


	r.ResizeImage()
}
