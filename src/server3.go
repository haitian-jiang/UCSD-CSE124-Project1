package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
    http.HandleFunc("/gendata", gendata)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func gendata(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
    if v, exist := r.Form["numBytes"]; !exist{
        fmt.Fprintln(w, "Please give the parameter numBytes!")
    } else {
        numBytesStr := v[0]
        if numBytes, err := strconv.Atoi(numBytesStr); err != nil {
            fmt.Fprintln(w, "Please give an integer as parameter!")
        } else {
            fmt.Fprintf(w, "%s", strings.Repeat("A", numBytes))
        }
    }
}
