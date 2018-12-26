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

func main() {
	T = time.Now()
	fmt.Printf("okaq favicon generator start: %s\n", time.Now().String())
	Files()
	defer Pn.Close()
	defer Js.Close()
	// generator
	Solid()
	// Block()
}


