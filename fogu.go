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

func FoguHandler(w http.Responsewriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// atomic counter
}

func main() {
	motd()
	http.HandleFunc("/", FoguHandler)
	http.HandleFunc("/s", StatHandler)
	http.ListenAndServe(":8080", nil)
}


