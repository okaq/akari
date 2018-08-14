// okaq akari webgl swan
// AQ <aq@okaq.com>
package main

import (
	"fmt"
	"net/http"
)

const (
	INDEX = "aomi.html"
)

func AomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w, r, INDEX)
}

func main() {
	fmt.Println("start okaq web on localhost:8080")
	// time
	http.HandleFunc("/", AomiHandler)
	http.ListenAndServe(":8080", nil)
}


