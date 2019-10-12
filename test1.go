package main

import "net/http"

var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	}
)
var in int = 100
var _ = 200

func test1() {
	http.Handle("/", http.FileServer(http.Dir('.')))
	_ = http.ListenAndServe(":8080", nil)
}
