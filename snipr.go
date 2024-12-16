package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/oliamb/cutter"
	hook "github.com/robotn/gohook"
	"github.com/vova616/screenshot"
	"golang.design/x/clipboard"
)

func main() {

	//wait for left click
	for {
		ok := hook.AddMouse("left")
		if ok {
			fmt.Println("click")
			break
		}
	}

	//get cursor coordinates after left click (left corner of screenshot)
	x1, y1 := robotgo.Location()
	fmt.Println(x1, y1)

	//wait for left click
	for {
		ok := hook.AddMouse("left")
		if ok {
			fmt.Println("click")
			break
		}
	}

	//get cursor coordinates after left click (bottom right corner of the screenshot)
	x2, y2 := robotgo.Location()
	fmt.Println(x2, y2)
	
	//take a screenshot of the whole screen (will later be cropped based on selection)
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	
	//calculate cropped image size based on click locations
	w := x2 - x1
	h := y2 - y1
	
	//crop image
	cimg, err := cutter.Crop(img, cutter.Config{
		Width:  w,
		Height: h,
		Anchor: image.Point{x1, y1},
	})
	
	//create a temporary image file that gets removed after copying
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	
	//encode the screenshot as png
	if err := png.Encode(f, cimg); err != nil {
		f.Close()
		panic(err)
	}
	
	//read saved (cropped) screenshot to cpimg
	cpimg, err := os.ReadFile("image.png")
	if err != nil {
		panic(err)
	}
	
	//write cpimg to clipboard
	clipboard.Write(clipboard.FmtImage, []byte(cpimg))
	os.Remove("image.png")
	
	//stay open until a change in the clipboard is detected
	for {
		time.Sleep(5 * time.Second)
		a := clipboard.Read(clipboard.FmtImage)
		if string(a) != string(cpimg) {
			break
		}
	}

}
