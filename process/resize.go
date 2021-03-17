package process

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type UploadedImage struct {
	Name string `json:"name"`
	Width uint `json:"width"`
	Height uint `json:"height"`
}

func(u *UploadedImage) ResizeImage()  {

	s := strings.Split(u.Name, ".")

	// open image
	file, err := os.Open(u.Name)
	if err != nil {
		log.Fatal(err)
	}

	// decode image.Image

	if s[1] == "jpg" || s[1] == "jpeg" || s[1] == "jpe" || s[1] == "jif" || s[1] == "jfif" || s[1] == "jfi" {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize
		// and preserve aspect ratio
		m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

		out, err := os.Create("temp/test_resized.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	} else if s[1] == "png" {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

		out, err := os.Create("temp/test_resized.png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		_ = png.Encode(out, m)
	}

	
}