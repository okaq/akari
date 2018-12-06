// okaq akari jiyu
// mobile web layout spec
// AQ <aq@okaq.com>
// 2018-11-28
package main

import (
	"encoding/json"
	"fmt"
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


