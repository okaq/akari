// okaq web server
// six connect
// AQ <aq@okaq.com>
// 2018-10-06
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	INDEX = "fogu.html"
)

// Cache, concurrency safe key value store
type Poeme struct {
	Dict map[string]string
	mu sync.Mutex
}

func FoguHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// atomic counter
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate and write player id
}

func motd() {
	fmt.Println("okaq web serving on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func main() {
	motd()
	cache()
	http.HandleFunc("/", FoguHandler)
	http.HandleFunc("/s", StatHandler)
	http.HandleFunc("/p", PidHandler)
	http.ListenAndServe(":8080", nil)
}


