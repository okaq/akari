// okaq web server
// six connect
// AQ <aq@okaq.com>
// 2018-10-06
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "fogu.html"
)

func FoguHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// atomic counter
}

func motd() {
	fmt.Println("okaq web serving on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", FoguHandler)
	http.HandleFunc("/s", StatHandler)
	http.ListenAndServe(":8080", nil)
}


