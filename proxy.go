package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func startProxy(port int, nodes []Node, currentNodeId *int) {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["nodes"] = nodes
		data["current"] = nodes[*currentNodeId].Url.String()

		js, err := json.MarshalIndent(data, "", "  ")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		currentNodeUrl := nodes[*currentNodeId].Url

		originHost := currentNodeUrl.Host
		originPathPrefix := currentNodeUrl.Path

		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", originHost)
		req.Host = originHost
		req.URL.Scheme = currentNodeUrl.Scheme
		req.URL.Host = originHost
		req.URL.Path = originPathPrefix + req.URL.Path
	}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if currentNodeId == nil {
			http.Error(w, "No available nodes", http.StatusInternalServerError)
			return
		}

		proxy.ServeHTTP(w, r)
	})

	Info.Printf("Starting proxy on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
