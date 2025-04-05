package main

import (
	"fmt"
	"io"
	"net/http"
	"proxy-cache/cache"
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
	cachedResp, cachedHeader, found := cache.Get(URL)
	if found {
		for k, v := range cachedHeader {
			for _, vv := range v {
				w.Header().Set(k, vv)
			}
		}
		w.Write(cachedResp)
		return
	}
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
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error trying to read body")
		return
	}
	fmt.Println("Cache miss")
	cache.Set(URL, bodyBytes, resp.Header)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyBytes)
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
