// okaq web server
// six connect
// AQ <aq@okaq.com>
// 2018-10-06
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	INDEX = "fogu.html"
)

var (
	P *Poeme
)

// Server side player identity
type Id struct {
	Date string
	Number string
}

// Cache, concurrency safe key value store
type Poeme struct {
	Dict map[string]string
	mu sync.Mutex
}

func (p *Poeme) Update(k, v string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Dict[k] = v
}

func (p *Poeme) Glean(k string) string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.Dict[k]
}

func FoguHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// invoke atomic counter here
	Increment()
	http.ServeFile(w,r,INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// output count
	s0 := P.Glean("count")
	s1 := fmt.Sprintf("count: %s", s0)
	b0 := byte(s1)
	w.Write(b0)
}

func Increment() {
	// atomic counter
	s0 := P.Glean("count")
	i0, _ := strconv.Atoi(s0)
	i1 := i0 + 1
	s1 := strconv.Itoa(i1)
	P.Update("count",s1)
	s2 := fmt.Sprintf("hit count: %s\n", s1)
	w.Write([]byte(s2))
	// return err on fail
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate and write player id
	var p Id
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Date: %s\nNumber: %s\n", p.Date, p.Number)
	// cache this id
	s0 := fmt.Sprintf("%s|%s", p.Number, p.Date)
	P.Update(s0, "ok")
	fmt.Println(P.Dict)
	// cache pretty print function
}

func motd() {
	fmt.Println("okaq web serving on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func cache() {
	P = &Poeme{}
	P.Dict = make(map[string]string)
	P.Update("count", "0")
}

func main() {
	motd()
	cache()
	http.HandleFunc("/", FoguHandler)
	http.HandleFunc("/s", StatHandler)
	http.HandleFunc("/p", PidHandler)
	http.ListenAndServe(":8080", nil)
}


