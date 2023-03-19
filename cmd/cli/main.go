package main

import (
	"flag"
	"fmt"
	"github.com/dthung1602/image2css/pkg"
	"image"
	"log"
	"os"
	"strings"
	"time"
)

type Args struct {
	imagePath  string
	outputPath string
	selector   string
	width      int
	height     int
	resolution int
}

func main() {
	args := parseArgs()
	img := getImageFromFilePath(args.imagePath)

	startTime := time.Now()
	img = image2css.ScaleImage(img, args.width, args.height)
	css := image2css.GenCSS(img, args.resolution, args.selector)
	log.Println("Process time:", time.Since(startTime))

	writeCss(css, args.outputPath)
}

func parseArgs() *Args {
	args := &Args{}
	flag.StringVar(&args.imagePath, "i", "", "path to image")
	flag.StringVar(&args.imagePath, "input", "", "path to image")
	flag.StringVar(&args.outputPath, "o", "", "path to css output file")
	flag.StringVar(&args.outputPath, "output", "", "path to css output file")
	flag.StringVar(&args.selector, "s", "#css-image", "css selector to render the image")
	flag.StringVar(&args.selector, "selector", "#css-image", "css selector to render the image")
	flag.IntVar(&args.width, "w", -1, "output image width")
	flag.IntVar(&args.width, "width", -1, "output image width")
	flag.IntVar(&args.height, "h", -1, "output image height")
	flag.IntVar(&args.height, "height", -1, "output image height")
	flag.IntVar(&args.resolution, "r", 5, "shadow resolution (in pixel)")
	flag.IntVar(&args.resolution, "resolution", 5, "shadow resolution (in pixel)")
	flag.Parse()
	if args.imagePath == "" {
		panic("image path is required")
	}
	if args.outputPath == "" {
		args.outputPath = strings.Replace(args.imagePath, ".png", ".css", 1)
		args.outputPath = strings.Replace(args.outputPath, ".jpg", ".css", 1)
		args.outputPath = strings.Replace(args.outputPath, ".jpeg", ".css", 1)
		args.outputPath = strings.Replace(args.outputPath, ".gif", ".css", 1)
		fmt.Println(args.imagePath, args.outputPath)
	}
	return args
}

func getImageFromFilePath(filePath string) image.Image {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		panic("cannot read image")
	}

	img, _, err := image.Decode(f)
	if err != nil {
		panic("cannot read image")
	}
	return img
}

func writeCss(css, outputPath string) {
	f, err := os.Create(outputPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(css)
	if err != nil {
		panic(err)
	}
}
