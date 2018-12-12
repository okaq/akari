// okaq bitmap sampler
// AQ <aq@okaq.com>
// 2018-12-12
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	INDEX = "kira.html"
	DATA = "kirad/"
)

func KiraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func meta() {
	// load persistant log file
	// unmarshal to local object
	// launch goroutine to update
	// marshal json and write to new log
}

func motd() {
	t0 := time.Now()
	fmt.Println("okaq web serve on localhost:8080")
	r0 := rand.New(rand.NewSource(t0.UnixNano())).Uint32()
	fmt.Printf("id: %d.\ntime: %s.\n", r0, t0.String())
}

func main() {
	motd()
	meta()
	http.HandleFunc("/", KiraHandler)
	http.ListenAndServe(":8080", nil)
}

// single data dir
// bat.png, bat.json, data.js

// persistent log
// write json metadata to filesystem


