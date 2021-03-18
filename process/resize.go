package process

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type UploadedImage struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Width uint `json:"width"`
	Height uint `json:"height"`
}

func(u *UploadedImage) ResizeImage()  {

	s := strings.Split(u.Name, ".")

	path := CreateDir(u.Path)

	fmt.Println("TET", path)


	// open image
	file, err := os.Open(u.Name)
	if err != nil {
		log.Fatal(err)
	}

	if s[1] == "jpg" || s[1] == "jpeg" || s[1] == "jpe" || s[1] == "jif" || s[1] == "jfif" || s[1] == "jfi" {
		// decode image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize
		// and preserve aspect ratio
		m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

		out, err := os.Create(path + u.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		_ = jpeg.Encode(out, m, nil)
	} else if s[1] == "png" || s[1] == "PNG" {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

		out, err := os.Create(path +u.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		_ = png.Encode(out, m)


	}
}

// basePath is a fixed directory path
func CreateDir(basePath string) (dataString string) {
	folderPath := filepath.Join(basePath, "/")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create folder first
		os.Mkdir(folderPath, 0777)
		// modify permissions again
		os.Chmod(folderPath, 0777)
	}
	result := folderPath + "/"
	return result
}
