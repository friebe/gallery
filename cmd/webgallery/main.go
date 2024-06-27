package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Image struct {
	Year  string
	Month string
	Path  string
}

type Gallery struct {
	Images map[string]map[string][]Image
}

//serve local img for test purpose
//const imageDir = "./static/images"

// mounted drive on my pi
const imageDir = "/mnt/external"
const port = 8080

func main() {
	http.HandleFunc("/", galleryHandler)
	http.Handle("/images/", cacheControl(http.StripPrefix("/images/", http.FileServer(http.Dir(imageDir)))))

	fmt.Printf("Starting server at port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func cacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		h.ServeHTTP(w, r)
	})
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	gallery := Gallery{
		Images: make(map[string]map[string][]Image),
	}

	err := filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isImage(info.Name()) {
			relPath, err := filepath.Rel(imageDir, path)
			if err != nil {
				return err
			}

			parts := strings.Split(relPath, string(os.PathSeparator))
			if len(parts) >= 3 {
				year := parts[0]
				month := parts[1]
				image := Image{
					Year:  year,
					Month: month,
					Path:  "/images/" + relPath,
				}

				if _, ok := gallery.Images[year]; !ok {
					gallery.Images[year] = make(map[string][]Image)
				}
				gallery.Images[year][month] = append(gallery.Images[year][month], image)
			}
		}
		return nil
	})

	if err != nil {
		http.Error(w, "Unable to read image directory", http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("templates/gallery.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, gallery)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}
