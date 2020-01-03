package main

import (
	"net/http"
	"gitlab.com/xblymmx/gopkg1"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello world"))
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
	}


}
