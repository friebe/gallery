package main

import (
	"fmt"
	"net/http"

	"github.com/friebe/webgallery/internal/gallery"
)

const imageDir = "static/images"

func main() {
	http.HandleFunc("/", gallery.GalleryHandler(imageDir))
	http.HandleFunc("/resize", gallery.ResizeHandler)

	port := 8080
	fmt.Printf("Starting server at port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
