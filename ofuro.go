// animated maps
// AQ <aq@okaq.com>
// 2018-12-30
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "ofuro.html"
)

func OfuroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}


