package main

import (
	"bytes"
	"fmt"
	"github.com/dthung1602/image2css/pkg"
	"image"
	"log"
	"syscall/js"
	"time"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("image2css", wrapper())
	<-make(chan bool)
}

func wrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		startTime := time.Now()
		jsImgBuffer := args[0].Get("image")
		byteLength := args[0].Get("byteLength").Int()
		width := args[0].Get("width").Int()
		height := args[0].Get("height").Int()
		resolution := args[0].Get("resolution").Int()
		selector := args[0].Get("selector").String()
		img := jsBufferToImg(jsImgBuffer, byteLength)
		log.Println("Read value from JS time:", time.Since(startTime))

		startTime = time.Now()
		img = image2css.ScaleImage(img, width, height)
		css := image2css.GenCSS(img, resolution, selector)
		log.Println("Process time:", time.Since(startTime))

		return css
	})
	return jsonFunc
}

func jsBufferToImg(jsBuff js.Value, byteLength int) image.Image {
	buff := make([]byte, byteLength)
	js.CopyBytesToGo(buff, jsBuff)
	reader := bytes.NewReader(buff)
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	return img
}
