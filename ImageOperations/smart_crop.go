package ImageOperations

import "github.com/h2non/bimg"

func SmartCrop(image []byte, height, width int) ([]byte, error) {
	options := bimg.Options{
		Width:   width,
		Height:  height,
		Crop:    true,
		Gravity: bimg.GravitySmart,
	}
	return bimg.NewImage(image).Process(options)
}
