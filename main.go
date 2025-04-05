package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var Port string
var Origin string
var Path string
var Method string

func handleClient(w http.ResponseWriter, r *http.Request) {
	Path = strings.Split(r.URL.Path, "/")[1]
	URL := Origin + Path
	req, err := http.NewRequest(r.Method, URL, r.Body)
	fmt.Println("Forwarding the request to -> " + URL)
	if err != nil {
		fmt.Println("Error trying to create the request " + err.Error())
		return
	}
	req.Header = r.Header
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error trying to forward the request to -> " + URL + " error: " + err.Error())
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	fmt.Scanf("caching-proxy --port %s --origin %s", &Port, &Origin)
	fmt.Printf("Port is %s", Port)
	fmt.Println()
	fmt.Printf("Origin is %s", Origin)
	fmt.Println()
	http.HandleFunc("/", handleClient)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}
