// +build ignore

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

// START SERVER OMIT
var async_count int64 = 0
var atomic_count int64 = 0

func main() {
	http.HandleFunc("/async", func(res http.ResponseWriter, req *http.Request) {
		async_count += 1
		fmt.Fprintf(os.Stdout, "ASYNC: %d\n", async_count)
	})

	http.HandleFunc("/atomic", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(os.Stdout, "ATOMIC: %d\n", atomic.AddInt64(&atomic_count, 1))
	})

	http.ListenAndServe(":3000", nil)
}

// END SERVER OMIT
