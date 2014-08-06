package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"os"
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
		fileInfo, err := os.Stat(realPath)
		if (err != nil) {
			if (os.IsNotExist(err)) {
				fmt.Fprintf(w, "File %s not found", fileInfo.Name())
				fmt.Fprintf(w, "File %s not found", err)
			}
		}else {
			if(hasIndexFile(realPath)){
				fmt.Fprintf(w, "RETURN THE INDEX FILE ! %s", fileInfo.Name())
			}else{
				if(fileInfo.IsDir()){
					fmt.Fprintf(w, "LIST FILE CHILDREN %s", fileInfo.Name())
				}else{
					fmt.Fprintf(w, "Return FILE %s", fileInfo.Name())
				}
			}
		}
	}
}


func hasIndexFile(folderPath string) bool {
	fileInfo , err := os.Stat(filepath.Join(folderPath, "index.html"))
	if(err != nil){
		//isNOtExistOrAnother problem lets assume that we won't be able to use it
		return false;
	}else{
		return fileInfo.IsDir()==false
	}
}

func resolveRealPath(handler *StorageHandler, requestedPath string) string {
	base := filepath.Clean(handler.UrlBase)
	value := requestedPath[len(base):len(requestedPath)]
	return filepath.Join(handler.InternalPath, value)
}
