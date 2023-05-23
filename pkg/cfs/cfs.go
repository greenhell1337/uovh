package cfs

import (
	"log"
	"net/http"
	"path/filepath"
)

type CloseFileSystem struct {
	Fs http.FileSystem
}

func (cfs CloseFileSystem) Open(path string) (http.File, error) {
	f, err := cfs.Fs.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := cfs.Fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}