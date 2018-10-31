// okaq akari hiko
// bitmap sample and encode
// AQ <aq@okaq.com>
// 2018-11-02
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "hiko.html"
)

func HikoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println("okaq hiko web started on localhost:8080")
	fmt.Println(time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", HikoHandler)
	http.ListenAndServe(":8080", nil)
}

// assign bitmap data directories


