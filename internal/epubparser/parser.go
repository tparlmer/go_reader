package epubparser

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Epub struct {
	Title string
	Author string
}

func ParseEPUB(filePath string) (*Epub, error) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open EPUB file: %v", err)
	}
	defer r.Close()

	var epub Epub

	for _, f := range r.File {
		if filepath.Base(f.Name) == "content.opf" {
			rc, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("failed to open content.opf: %v", err)
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, fmt.Errorf("failed to read content.opf: %v", err)
			}

			var metadata struct {
				Title string `xml:"metadata>title"`
				Author string `xml:"metadata>creator"`
			}

			err = xml.Unmarshal(data, &metadata)
			if err != nil {
				return nil, fmt.Errorf("failed to parse content.opf: %v", err)
			}

			epub.Title = metadata.Title
			epub.Author = metadata.Author
		}
	}

	return &epub, nil
}