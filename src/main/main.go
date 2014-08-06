package main

import (
	"net/http"
	"handlers"
	"log"
)


func main() {
	var storageHandler * handlers.StorageHandlerDefinition = handlers.New("/browse/", "/tmp/martial")
	http.HandleFunc(storageHandler.UrlBase, storageHandler.CreateHandler())
	//TODO 404 handling
	error := http.ListenAndServe(":8080", nil)
	if(error!=nil){
		log.Fatal(error)
	}
}
