package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

type Image struct {
    Year  string
    Month string
    Path  string
}

func main() {
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        images, err := loadImages("images")
        if err != nil {
            http.Error(w, "Unable to load images", http.StatusInternalServerError)
            return
        }
		
        t, err := template.ParseFiles("template.html")
        if err != nil {
            http.Error(w, "Unable to load template", http.StatusInternalServerError)
            return
        }

        t.Execute(w, images)
    })

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

func loadImages(root string) ([]Image, error) {
    var images []Image

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && isImageFile(path) {
            relPath, err := filepath.Rel(root, path)
            if err != nil {
                return err
            }
            parts := strings.Split(relPath, string(os.PathSeparator))
            if len(parts) >= 3 {
                images = append(images, Image{
                    Year:  parts[0],
                    Month: parts[1],
                    Path:  "/images/" + relPath,
                })
            }
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    return images, nil
}

func isImageFile(path string) bool {
    ext := strings.ToLower(filepath.Ext(path))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}
