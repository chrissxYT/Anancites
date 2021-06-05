package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RemoteAddr)
		req.Host = "mobileapi.dsbcontrol.de"
		req.URL.Host = req.Host
		req.URL.Scheme = "https"
		req.RequestURI = ""
		res, err := client.Do(req)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err)
		} else {
			w.Header().Add("Access-Control-Allow-Origin", "https://*.chrissx.de")
			w.WriteHeader(res.StatusCode)
			io.Copy(w, res.Body)
		}
	})

	http.ListenAndServe(":8066", nil)
}
