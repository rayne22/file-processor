package process

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type UploadedImage struct {
	Name string `json:"name"`
	Width uint `json:"width"`
	Height uint `json:"height"`
}

func(u *UploadedImage) ResizeImage()  {

	// open image
	file, err := os.Open(u.Name)
	if err != nil {
		log.Fatal(err)
	}

	// decode image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

	out, err := os.Create("temp/test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	
}