// okaq akari hiko
// bitmap sample and encode
// AQ <aq@okaq.com>
// 2018-11-02
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	INDEX = "hiko.html"
	// png image folder
	INPUT = "./hikoa/"
	// sample data, json, base64 string
	OUPUT = "./hikob/"
)

var (
	// list of png bitmaps
	Png []os.FileInfo
	// list json
	Json []os.FileInfo
)

func HikoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func load() {
	fmt.Println("reading file directories")
}

func motd() {
	fmt.Println("okaq hiko web started on localhost:8080")
	fmt.Println(time.Now().String())
}

func main() {
	motd()
	load()
	http.HandleFunc("/", HikoHandler)
	http.ListenAndServe(":8080", nil)
}

// assign bitmap data directories


