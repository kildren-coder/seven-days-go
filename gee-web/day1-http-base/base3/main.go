package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	fmt.Println(r)
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %s\n", req.URL.Path)
	})

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%s] = %s\n", k, v)
		}
	})

	r.Run(":9998")
}