package process

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	_ "net/http"
	"os"
	"path/filepath"
	"strings"
)

type UploadedImage struct {
	Path      string `json:"path"`
	ImageName string `json:"image_name"`
	Width     uint   `json:"width"`
	Height    uint   `json:"height"`
	Request   multipart.File
	ImageHeader *multipart.FileHeader
}

func(u *UploadedImage, ) ResizeImage()   string {

	var imagePath string

	path := CreateDir(u.Path)



	if u.Request != nil {
		file := u.Request
		s := strings.Split(u.ImageHeader.Filename, ".")

		imagePath = path +"/" + u.ImageHeader.Filename

		if s[1] == "jpg" || s[1] == "jpeg" || s[1] == "jpe" || s[1] == "jif" || s[1] == "jfif" || s[1] == "jfi" {
			// decode image.Image
			img, err := jpeg.Decode(file)
			if err != nil {
				log.Println("JPEG Decode Error",err)
			}
			file.Close()

			// resize
			// and preserve aspect ratio
			m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

			out, err := os.Create(path + u.ImageHeader.Filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			_ = jpeg.Encode(out, m, nil)
		} else if s[1] == "png" || s[1] == "PNG" {
			img, err := png.Decode(file)
			if err != nil {
				log.Println("PNG Decode Error",err)
			}
			file.Close()

			// resize to width 1000 using Lanczos resampling
			// and preserve aspect ratio
			m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

			out, err := os.Create(path +u.ImageHeader.Filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			_ = png.Encode(out, m)


		}

	} else {
		s := strings.Split(u.ImageName, ".")
		imagePath = path +"/" + u.ImageName

		// open image
		file, err := os.Open(u.ImageName)
		if err != nil {
			log.Fatal(err)
		}

		if s[1] == "jpg" || s[1] == "jpeg" || s[1] == "jpe" || s[1] == "jif" || s[1] == "jfif" || s[1] == "jfi" {
			// decode image.Image
			img, err := jpeg.Decode(file)
			if err != nil {
				log.Println("JPEG Decode Error",err)
			}
			file.Close()

			// resize
			// and preserve aspect ratio
			m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

			out, err := os.Create(path + u.ImageName)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			_ = jpeg.Encode(out, m, nil)
		} else if s[1] == "png" || s[1] == "PNG" {
			img, err := png.Decode(file)
			if err != nil {
				log.Println("PNG Decode Error",err)
			}
			file.Close()

			// resize to width 1000 using Lanczos resampling
			// and preserve aspect ratio
			m := resize.Resize(u.Width, u.Height, img, resize.Lanczos3)

			out, err := os.Create(path +u.ImageName)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			_ = png.Encode(out, m)




		}
	}



	return imagePath
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


