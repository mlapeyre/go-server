package main

import (
	"net/http"
	"github.com/mlapeyre/prettyFileServe"
	"log"
)


func main() {
	var storageHandler * prettyFileServe.StorageHandlerDefinition = prettyFileServe.New("/browse/", "/tmp/martial")
	http.HandleFunc(storageHandler.UrlBase, storageHandler.CreateHandler())
	//TODO 404 handling
	error := http.ListenAndServe(":8080", nil)
	if(error!=nil){
		log.Fatal(error)
	}
}
