package gallery

import (
	"html/template"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
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

	// Prüfe, ob die Datei existiert
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Lade das Bild
	img, err := imaging.Open(imagePath)
	if err != nil {
		http.Error(w, "Unable to read image", http.StatusInternalServerError)
		return
	}

	// Skalieren des Bildes
	resizedImg := imaging.Resize(img, 100, 100, imaging.Lanczos)

	// Setze den Content-Type Header
	w.Header().Set("Content-Type", "image/jpeg")

	// Sende das verkleinerte Bild an den Client
	err = imaging.Encode(w, resizedImg, imaging.JPEG)
	if err != nil {
		http.Error(w, "Unable to encode image", http.StatusInternalServerError)
		return
	}
}
