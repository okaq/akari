// okaq akari webgl swan
// AQ <aq@okaq.com>
package main

import (
	"fmt"
	"net/http"
	"time"
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
	fmt.Printf("the current time is: %s\n", time.Now().String())
	http.HandleFunc("/", AomiHandler)
	http.ListenAndServe(":8080", nil)
}


