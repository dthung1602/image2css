package image2css

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"
)

func ScaleImage(img image.Image, width, height int) image.Image {
	if width == -1 && height == -1 {
		return img
	}

	if width == -1 {
		size := img.Bounds().Size()
		width = int(float32(height) * float32(size.X) / float32(size.Y))
	}

	if height == -1 {
		size := img.Bounds().Size()
		height = int(float32(width) * float32(size.Y) / float32(size.X))
	}

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(newImg, newImg.Rect, img, img.Bounds(), draw.Over, nil)
	return newImg
}

func GenCSS(img image.Image, resolution int, cssSelector string) string {
	blur := resolution - 1
	if blur == 0 {
		blur = 1
	}

	imgSize := img.Bounds().Size()
	shadows := make([]string, 0)

	for i := 0; i < imgSize.X; i += resolution {
		for j := 0; j < imgSize.Y; j += resolution {
			r, g, b, a := avgColorAt(img, i, j, resolution)
			line := fmt.Sprintf(
				"%dpx %dpx %dpx %dpx rgba(%d, %d, %d, %.2f)",
				i, j, blur, resolution,
				scaleColor(r), scaleColor(g), scaleColor(b), float32(a)/65536,
			)
			shadows = append(shadows, line)
		}
	}

	css := "%s {\n    width: 0;\n    height: 0;\n    box-shadow:\n        %s\n}"
	return fmt.Sprintf(css, cssSelector, strings.Join(shadows, ",\n        "))
}

func avgColorAt(img image.Image, x, y, resolution int) (r uint32, g uint32, b uint32, a uint32) {
	size := img.Bounds().Size()
	sizeX := minInt(size.X, x+resolution)
	sizeY := minInt(size.Y, y+resolution)
	total := uint32((sizeX - x) * (sizeY - y))

	for i := x; i < sizeX; i++ {
		for j := y; j < sizeY; j++ {
			pr, pg, pb, pa := img.At(i, j).RGBA()
			r += pr
			g += pg
			b += pb
			a += pa
		}
	}
	return r / total, g / total, b / total, a / total
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scaleColor(n uint32) uint8 {
	return uint8(float32(n) / 65536 * 256)
}
