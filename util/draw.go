package util

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
)

func jpegEncode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

func Draw(m [][]bool, fn string) error {
	ext := path.Ext(fn)
	encodeFunc := jpegEncode
	switch ext {
	case ".jpeg", ".jpg":
	case ".png":
		encodeFunc = png.Encode
	default:
		return errors.New("only supports JPEG/PNG")
	}
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	// increase size
	increaseFactor := 5
	height := len(m) * increaseFactor
	width := len(m[0]) * increaseFactor

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if m[y/increaseFactor][x/increaseFactor] {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}

	return encodeFunc(f, img)
}
