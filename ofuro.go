// animated maps
// AQ <aq@okaq.com>
// 2018-12-30
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "ofuro.html"
)

func OfuroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	fmt.Println("okaq web start on localhost:8080")
	fmt.Println(time.Now().String())
	http.HandleFunc("/", OfuroHandler)
	http.ListenAndServe(":8080", nil)
}


