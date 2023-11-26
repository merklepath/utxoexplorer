package main

import (
	"log"
	"net/http"

	"github.com/merklepath/utxoexplorer/static"
)

func main() {
	// Initialize the router
	http.HandleFunc("/", webui.RenderIndex)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.Static))))

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
