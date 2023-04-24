package web

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed build
var buildFiles embed.FS

func Handler() http.Handler {
	serveFiles, err := fs.Sub(buildFiles, "build")
	if err != nil {
		panic(err) // shouldnt happen
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		path := filepath.Clean(r.URL.Path)
		if path == "/" { // Add other paths that you route on the UI side here
			path = "index.html"
		}
		path = strings.TrimPrefix(path, "/")

		file, err := serveFiles.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				log.Println("file", path, "not found:", err)
				http.NotFound(w, r)
				return
			}
			log.Println("file", path, "cannot be read:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		contentType := mime.TypeByExtension(filepath.Ext(path))
		w.Header().Set("Content-Type", contentType)
		if strings.HasPrefix(path, "static/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000")
		}
		stat, err := file.Stat()
		if err == nil && stat.Size() > 0 {
			w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
		}

		n, _ := io.Copy(w, file)
		log.Println("file", path, "copied", n, "bytes")
	})
}
