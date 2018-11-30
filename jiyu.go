// okaq akari jiyu
// mobile web layout spec
// AQ <aq@okaq.com>
// 2018-11-28
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "jiyu.html"
)

func JiyuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate and store key
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// atomic counter
}

func motd() {
	fmt.Println("okaq jiyu on localhost:8080")
	fmt.Println(time.Now().String())
}

func main() {
	motd()
	// init cache
	http.HandleFunc("/", JiyuHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/s", StatsHandler)
	http.ListenAndServe(":8080", nil)
}

// localhost tunnel to public www
// test on multiple mobile web screens

// no logging
// cached live stats served on /s

// triva generator
// interface for admin (edit,add,delete)
// upload user generated and curate
// serve one at a time to client
// query answer on server


