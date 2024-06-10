package main

import (
    "log"
    "net/http"
    "path/filepath"
    "text/template"

    "github.com/tparlmer/go_reader/internal/api"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/upload", api.UploadHandler)
    mux.HandleFunc("/api/epubs/", api.EpubHandler)
    mux.HandleFunc("/api/metadata/", api.MetadataHandler)
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "index.html")
    })
    mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "upload.html")
    })
    mux.HandleFunc("/reader", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "reader.html")
    })

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles(filepath.Join("web/templates", tmpl))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}