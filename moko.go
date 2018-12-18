// okaq questions web
// AQ <aq@okaq.com>
// 2018-12-20
package main

import (
	"fmt"
	"net/http"
)

const (
	INDEX = "moko.html"
)

func MokoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	fmt.Printf("okaq web start on localhost:8080\n%s\n", time.Now().Stirng())
	http.HandleFunc("/", MokoHandler)
	http.ListenAndServe(":8080", nil)
}

// question and answer stream


