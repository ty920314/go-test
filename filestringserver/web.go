package main

import (
	"go-test/filestringserver/filelist"
	"net/http"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandle) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc("/list/", errWrapper(filelist.HandleFileList))
	serve := http.ListenAndServe(":8889", nil)
	if serve != nil {
		panic(serve)
	}
}
