package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed ananke categories images posts tags  
//go:embed 404.html go.mod index.html index.xml server server.go sitemap.xml  
var siteData embed.FS

func main() {
	listenAddr := ":8081"
	if len(os.Getenv("LISTEN_ADDR")) != 0 {
		listenAddr = os.Getenv("LISTEN_ADDR")

	}
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(http.FS(siteData))
	mux.Handle("/", staticFileServer)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
