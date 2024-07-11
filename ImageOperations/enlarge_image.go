package ImageOperations

import "github.com/h2non/bimg"

func Enlarge(image []byte, height, width int) ([]byte, error) {
	options := bimg.Options{
		Width:   width,
		Height:  height,
		Enlarge: true,
	}
	return bimg.NewImage(image).Process(options)
}
