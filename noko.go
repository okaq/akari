// favicon generate and encode
// AQ <aq@okaq.com>
// 2018-12-25
// merry xmas!
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"time"
)

const (
	PATH = "nokoa/"
	PNG = "favicon_"
	PNG_E = ".png"
	JSON = "dataUri_"
	JSON_E = ".json"
)

var (
	// png path
	P string
	// png file
	Pn *os.File
	// json path
	J string
	// json file
	Js *os.File
	// time
	T time.Time
	// pixels
	I *image.RGBA
)

func Files() {
	fmt.Println("opening files to save on disk")
	t := T.UnixNano()
	p := fmt.Sprintf("%s%s%d%s", PATH, PNG, t, PNG_E)
	j := fmt.Sprintf("%s%s%d%s", PATH, JSON, t, JSON_E)
	fmt.Println(p, j)
	var err error
	Pn, err = os.Create(p)
	if err != nil {
		panic(err)
	}
	Js, err = os.Create(j)
	if err != nil {
		panic(err)
	}
}

func Pixels() {
	r := image.Rect(0,0,32,32)
	I = image.NewRGBA(r)
}

func Solid() {
	// generate a solid color favicon
	c := color.RGBA{128,128,128,255}
	dx := I.Bounds().Dx()
	dy := I.Bounds().Dy()
	b := dx * dy
	for i := 0; i < b; i++ {
		x := i % dx
		y := i / dx
		I.SetRGBA(x,y,c)
	}
}

func Square() {
	// two color square within a square
}

func main() {
	T = time.Now()
	fmt.Printf("okaq favicon generator start: %s\n", time.Now().String())
	Files()
	defer Pn.Close()
	defer Js.Close()
	Pixels()
	// generator
	Solid()
	// Block()
}


