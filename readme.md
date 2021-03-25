# Image Processor

A GO library for manipulating and processing images


[![GoDoc](https://pkg.go.dev/badge/github.com/rayne22/file-processor)](https://pkg.go.dev/github.com/rayne22/file-processor)
[![Gitter](https://badges.gitter.im/go-thots/community.svg)](https://gitter.im/go-thots/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Prerequisites
This package needs at least Go `1.15.6`

## Installation

```bash
 go get -u github.com/rayne22/file-processor
````

Easy right...

## Usage
 Import package using

```go
import "github.com/rayne22/file-processor/process"

````

### Image Resizing
- `process.ResizeImage` produces a scaled image based on the width and height provided using the interpolation. If there is need to preserve the aspect ratio, the width or height can be set to 0.

#####  Struct

````go
type UploadedImage struct {
	Path        string `json:"path"`
	ImageName   string `json:"image_name"`
	Width       uint   `json:"width"`
	Height      uint   `json:"height"`
	Request     multipart.File
	ImageHeader *multipart.FileHeader
}
````

##### Function

````go
func (u *UploadedImage) ResizeImage() string 
````

#### Example

````go

processor := process.UploadedImage{}  //Initializing the Struct


processor.Height = 200  // Adding height
processor.Width = 200  // Adding width
processor.Path = "temp"  // Adding temporary image storage directory

processor.Request = file //  If image is being posted through a form, the file is stored in this field
processor.ImageHeader = fileHeader // If image is being posted through a form, the fileHeader is stored in this field

_= processor.ResizeImage() // Function for resizing images

````


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
