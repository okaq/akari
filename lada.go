// okaq bitmap json decode and render
// AQ <aq@okaq.com>
// 2018-12-18
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	INDEX = "lada.html"
)

var (
	R *rand.Rand
	Id uint64
)

func LadaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println("start okaq web on localhost:8080...")
	fmt.Println(time.Now().String())
	fmt.Println(Id)
}

func id() {
	t0 := time.Now().UnixNano()
	s0 := rand.NewSource(t0)
	R = rand.New(s0)
	Id = R.Uint64()
}

func main() {
	id()
	motd()
	http.HandleFunc("/", LadaHandler)
	http.ListenAndServe(":8080", nil)

}


