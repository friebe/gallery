package gallery

import (
	"html/template"
	"net/http"
	"os"

	"github.com/h2non/bimg"
)

// GalleryHandler lädt die Bilder und rendert die Galerie-Seite.
func GalleryHandler(imageDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := loadImages(imageDir)

		t, err := template.ParseFiles("templates/gallery.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}

		t.Execute(w, GalleryData{Images: images})
	}
}

// ResizeHandler skaliert die Bilder serverseitig und sendet sie an den Client.
func ResizeHandler(w http.ResponseWriter, r *http.Request) {
	imagePath := r.URL.Query().Get("path")
	if imagePath == "" {
		http.Error(w, "Image path is required", http.StatusBadRequest)
		return
	}

	file, err := os.ReadFile(imagePath)
	if err != nil {
		http.Error(w, "Unable to read image", http.StatusInternalServerError)
		return
	}

	options := bimg.Options{
		Width:  100,
		Height: 100,
		Crop:   true,
	}

	newImage, err := bimg.NewImage(file).Process(options)
	if err != nil {
		http.Error(w, "Unable to process image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(newImage)
}
