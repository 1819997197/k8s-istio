package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		s := "hello world!"
		fmt.Println("response: ", s)
		w.Write([]byte(s))
	})

	fmt.Println("listen 50050")
	http.ListenAndServe("0.0.0.0:50050", nil)
}
