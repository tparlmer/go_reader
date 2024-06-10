package api

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/tparlmer/go_reader/internal/epubparser"
)

func EpubHandler(w http.ResponseWriter, r *http.Request) {
	epubID := strings.TrimPrefix(r.URL.Path, "/api/epubs/")
	if epubID == "" {
		http.Error(w, "EPUB ID not specified", http.StatusBadRequest)
		return
	}

	epub, err := epubparser.ParseEPUB("uploads/" + epubID)
	if err != nil {
		http.Error(w, "Failed to parse EPUB", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Parsed EPUB successfully: %s", epub.Title)
}