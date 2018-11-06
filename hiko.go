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
	OUTPUT = "./hikob/"
)

var (
	// list of png bitmaps
	Png []os.FileInfo
	// list json
	Json []os.FileInfo
)

func Lister() {
	// pretty print file lists
	fmt.Println("PNG Files...")
	for _, f0 := range(Png) {
		fmt.Printf("file name: %s\n", f0.Name())
		fmt.Printf("file size: %d bytes\n", f0.Size())
	}
	fmt.Println("JSON Files...")
	for _, f0 := range(Json) {
		fmt.Printf("file name: %s\n", f0.Name())
		fmt.Printf("file size: %d bytes\n", f0.Size())
	}
}

func HikoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate and cache player id
}

func TokHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	time.Sleep(time.Second * 2)
	w.Write([]byte("ok"))
}

func load() {
	fmt.Println("reading file directories")
	var err error
	Png, err = ioutil.ReadDir(INPUT)
	if err != nil {
		fmt.Println(err)
	}
	Json, err = ioutil.ReadDir(OUTPUT)
	if err != nil {
		fmt.Println(err)
	}
	Lister()
}

func motd() {
	fmt.Println("okaq hiko web started on localhost:8080")
	fmt.Println(time.Now().String())
}

func main() {
	motd()
	load()
	http.HandleFunc("/", HikoHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/b", TokHandler)
	http.ListenAndServe(":8080", nil)
}

// assign bitmap data directories
// pretty print file info and sync
// server side and browser client states


