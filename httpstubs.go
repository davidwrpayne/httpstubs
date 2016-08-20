package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`hello httpstubs`))
	})
	http.ListenAndServe("0.0.0.0:3000", nil)
}