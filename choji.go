// okaq akari 
// "cookies"
// AQ <aq@okaq.com>
// 2018-08-22
package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "choji.html"
)

type Pid struct {
	// cache
	// atomic counters
}

type Player struct {
	Pid string `json:"Pid"`
}

type Id struct {
	Pid string
	Sid string
}

func NewId(s string) *Id {
	i0 := Id{}
	i0.Pid = s
	i0.Sid = Id.Generate()
	return &io
}

func (id *Id) Generate() string {
	// var b0 bytes.Buffer
	s0 := time.Now().UnixNano()
	s1 := rand.New(rand.NewSource(s0)).Int64()
	s2 := fmt.Sprintf("%d:%d", s0, s1)
	return s2
}

func (id *Id) String() string {
	s0 := fmt.Sprintf("%s;%s", id.Pid, id.Sid)
	return s0
}

func ChojiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w, r, INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate player id
	var p0 Player
	err := json.NewDecoder(r.Body).Decode(&p0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p0)

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

// run atomic ops as separate server
// auth is simply generated id in memory cache

// Pid class
// one stop shop for all token pipeline methods
// even handles json responses
