// okaq web server
// webgl labyrinth
// eiga = image (jp)
// AQ <aq@okaq.com>
// 2018-09-20
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "eiga.html"
)

func motd() {
	fmt.Printf("okaq web on localhost:8080\n%s\n", time.Now().String())
}

func EigaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", EigaHandler)
	http.ListenAndServe(":8080", nil)
}


