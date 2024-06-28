package gallery

import (
	"os"
	"path/filepath"
	"strings"
)

// Image repräsentiert eine Bilddatei mit Jahr, Monat und Pfad.
type Image struct {
	Year  string
	Month string
	Path  string
}

// GalleryData repräsentiert die Daten für die Galerie-Vorlage.
type GalleryData struct {
	Images map[string]map[string][]Image
}

// loadImages lädt die Bilder aus dem angegebenen Verzeichnis.
func loadImages(imageDir string) map[string]map[string][]Image {
	images := make(map[string]map[string][]Image)

	filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".jpg") || strings.HasSuffix(info.Name(), ".jpeg") || strings.HasSuffix(info.Name(), ".png")) {
			relPath := strings.TrimPrefix(path, imageDir+"/")
			parts := strings.Split(relPath, string(os.PathSeparator))
			if len(parts) >= 3 {
				year, month := parts[0], parts[1]
				if _, ok := images[year]; !ok {
					images[year] = make(map[string][]Image)
				}
				images[year][month] = append(images[year][month], Image{
					Year:  year,
					Month: month,
					Path:  path,
				})
			}
		}
		return nil
	})

	return images
}
