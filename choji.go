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

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate player id

	// pipeline 
	// parse json request {"pid":value}
	// generate server id
	// return json response
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// return text stats

	// unique page txt/plain
	// with basic page stats
	// include atomic requests counter
}

func main() {
	fmt.Println("start okaq web on localhost:8080")
	// time
	fmt.Printf("start time: %s\n", time.Now().String())
	http.HandleFunc("/", ChojiHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/s", StatsHandler)
	http.ListenAndServe(":8080", nil)
}


