// okaq web server
// "space races"
// daiji = "important" (jp)
// AQ <aq@okaq.com>
// 2018-09-10
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "daiji.html"
)

type Grid struct {
	// cells
	Cells [][]byte
	// stride
	Stride uint64
}

func motd() {
	fmt.Printf("okaq web serve on localhost:8080\n%s\n", time.Now().String())
}

func DaijiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// channel sync for updating cache and counter
	// start receiver in main
	// store and increment on each send
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// total request count
}

func main() {
	motd()
	http.HandleFunc("/", DaijiHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/b", StatsHandler)
	http.ListenAndServe(":8080", nil)
}

// grid generate and api
// player id and peer net broker


