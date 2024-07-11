package ImageOperations

import "github.com/h2non/bimg"

func CropImage(image []byte, height, width int) ([]byte, error) {
	options := bimg.Options{
		Height: height,
		Width:  width,
		Crop:   true,
	}
	processedImage, err := bimg.NewImage(image).Process(options)
	if err != nil {
		return nil, err // Return the specific error from bimg
	}
	return processedImage, nil

}
