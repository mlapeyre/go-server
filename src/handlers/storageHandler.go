package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type StorageHandler struct {
	UrlBase      string
	InternalPath string
}

func New(urlBase string, internalPath string) *StorageHandler {
	var storage = new(StorageHandler)
	storage.InternalPath = internalPath
	storage.UrlBase = urlBase
	fmt.Println("")
	return storage
}

func (handler *StorageHandler) CreateHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		realPath := resolveRealPath(handler, r.URL.Path)
		fmt.Fprintf(w, "Trying to get %s", realPath)
	}
}

func resolveRealPath(handler *StorageHandler, requestedPath string) string {
  	base := filepath.Clean(handler.UrlBase)
	value := requestedPath[len(base):len(requestedPath)]
	fmt.Println("value="+value)
	return filepath.Join(handler.InternalPath,value)
}
