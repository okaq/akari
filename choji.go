// okaq akari 
// "cookies"
// AQ <aq@okaq.com>
// 2018-08-22
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "choji.html"
)

func ChojiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w, r, INDEX)
}

func main() {
	fmt.Println("start okaq web on localhost:8080")
	// time
	fmt.Printf("start time: %s\n", time.Now().String())
	http.HandleFunc("/", ChojiHandler)
	http.ListenAndServe(":8080", nil)
}


