package main

import (
	"net/http"
	"handlers"
	"log"
)


func main() {
	var storageHandler * handlers.StorageHandler = handlers.New("/browse/", "/tmp/martial")
	http.HandleFunc(storageHandler.UrlBase, storageHandler.CreateHandler())
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	error := http.ListenAndServe(":8080", nil)
	if(error!=nil){
		log.Fatal(error)
	}
}
