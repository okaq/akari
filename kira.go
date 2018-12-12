// okaq bitmap sampler
// AQ <aq@okaq.com>
// 2018-12-12
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	INDEX = "kira.html"
	DATA = "kirad/"
)

var (
	P []string
	J []byte
)

func KiraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func files() {
	// generate list of file paths
	// sort in lexicographic order
	f0, err := ioutil.ReadDir(DATA)
	if err != nil {
		fmt.Println(err)
	}
	P = make([]string, len(f0))
	for i, f1 := range f0 {
		// fmt.Println(f1.Name())
		// P[i] = f1.Name()
		f2 := fmt.Sprintf("%s%s", DATA, f1.Name())
		P[i] = f2
	}
	fmt.Println(P)
	var err2 error
	J, err2 = json.Marshal(P)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(J))
}

func meta() {
	// load persistant log file
	// unmarshal to local object
	// launch goroutine to update
	// marshal json and write to new log
}

func motd() {
	t0 := time.Now()
	fmt.Println("okaq web serve on localhost:8080")
	r0 := rand.New(rand.NewSource(t0.UnixNano())).Uint32()
	fmt.Printf("id: %d.\ntime: %s.\n", r0, t0.String())
}

func main() {
	motd()
	meta()
	files()
	http.HandleFunc("/", KiraHandler)
	http.ListenAndServe(":8080", nil)
}

// single data dir
// bat.png, bat.json, data.js

// persistent log
// write json metadata to filesystem


