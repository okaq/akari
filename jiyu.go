// okaq akari jiyu
// mobile web layout spec
// AQ <aq@okaq.com>
// 2018-11-28
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	INDEX = "jiyu.html"
)

var (
	// Pid cache
	// C map[string]string
	Store *Cache
)

type Cache struct {
	C map[string]string
	*sync.Mutex
}

type Screen struct {
	// DOM window.screen details
	W int64 `json:"Width"`
	H int64 `json:"Height"`
	Pid int64 `json:"Pid"`
}

func JiyuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate and store key
	j0 := struct {
		Pid int64 `json:"Pid"`
	} {
		Pid: 0,
	}
	err := json.NewDecoder(r.Body).Decode(&j0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Player id: %d\n", j0.Pid)
}

func ScreenHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// get browser resolution
	j0 := new(Screen)
	err := json.NewDecoder(r.Body).Decode(&j0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Browser width = %d. Height = %d.\n", j0.W, j0.H)
	// store in cache
	// simple string format "W:H"
	// otherwise json.Marshal(Screen)
	s0 := fmt.Sprintf("%d:%d", j0.W, j0.H)
	// requires pid
	// refactor to do everything in one call
	// browser sends POST server side with screen resolution
	// server responds with generated Pid
	// start goroutine to store in cache
	// or sync via channel running for loop reciever
	fmt.Println(s0)
}

func ScreenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// parse req body json
	j0 := new(Screen)
	err := json.NewDecoder(r.Body).Decode(&j0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Browser screen resolution: %d x %d.\n", j0.W, j0.H)
	// generate player id
	// timestamp
	t0 := time.Now().UnixNano()
	// random int32
	n0 := rand.New(rand.NewSource(t0))
	n1 := n0.Uint32()
	p0 := fmt.Sprintf("%d:%d", n1, t0)
	fmt.Printf("Player id: %s.\n", p0)
	// cache
	// send pid response
	j1 := struct {
		Pid string `json:"Pid"`
	} {
		Pid: p0,
	}
	j2, err := json.Marshal(j1)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(j2)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// atomic counter
}

func motd() {
	fmt.Println("okaq jiyu on localhost:8080")
	fmt.Println(time.Now().String())
}

func data() {
	Store = &Cache{}
	Store.C = make(map[string]string)
	Store.C["0"] = "app start"
	fmt.Println(Store.C)
}

func main() {
	motd()
	// init cache
	data()
	http.HandleFunc("/", JiyuHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/b", ScreenHandler)
	http.HandleFunc("/s", StatsHandler)
	http.ListenAndServe(":8080", nil)
}

// localhost tunnel to public www
// test on multiple mobile web screens

// no logging
// cached live stats served on /s

// triva generator
// interface for admin (edit,add,delete)
// upload user generated and curate
// serve one at a time to client
// query answer on server


