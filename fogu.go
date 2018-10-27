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

var (
	P *Poeme
)

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

func cache() {
	P = Poeme{}
	P.Dict = make(map[string]string)
}

func main() {
	motd()
	cache()
	http.HandleFunc("/", FoguHandler)
	http.HandleFunc("/s", StatHandler)
	http.HandleFunc("/p", PidHandler)
	http.ListenAndServe(":8080", nil)
}


