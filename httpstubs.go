package main

import (
	"net/http"
	"github.com/dnaeon/go-vcr/recorder"
	"os"
	"fmt"
	"io"
)

func main() {
	file := "example"

	if _, err := os.Stat(fmt.Sprintf("%s.yaml", file)); os.IsNotExist(err) {
		// missing example file should halt server boot.
		panic(err)
	}

	var vcr *recorder.Recorder
	if rec, err := recorder.New("example"); err != nil {
		// invalid cassette should halt server boot.
	 	panic(err)
	} else {
		vcr = rec
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if response, err := vcr.Transport.RoundTrip(r); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error: %s", err)
		} else {
			w.WriteHeader(response.StatusCode)
			for k, v := range response.Header {
				for _, i := range v {
					response.Header.Add(k, i)
				}
			}
			io.Copy(w, response.Body)
		}
	})
	http.ListenAndServe("0.0.0.0:3000", nil)
}