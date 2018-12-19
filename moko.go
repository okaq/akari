// okaq questions web
// AQ <aq@okaq.com>
// 2018-12-20
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "moko.html"
)

func MokoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	fmt.Printf("okaq web start on localhost:8080\n%s\n", time.Now().String())
	http.HandleFunc("/", MokoHandler)
	http.ListenAndServe(":8080", nil)
}

// question and answer stream

// each individual question is 
// stored in a json flat file on disk
// on server start, cache in memory as []byte
// select one from key list to ask
// keep track of player stream


