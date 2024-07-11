package ImageOperations

import (
	"fmt"

	"github.com/h2non/bimg"
)

func Resize(image []byte, height, width int) ([]byte, error) {
	if height < 0 || width < 0 {
		return nil, fmt.Errorf("height and width must be positive numbers")
	}
	options := bimg.Options{
		Width:  width,
		Height: height,
		Embed:  true,
	}
	return bimg.NewImage(image).Process(options)
}
