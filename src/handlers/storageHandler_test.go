package handlers

import (
	"testing"
)
var testData = []struct {
		basePattern string
		urlRequest string
		internalPath string
		resolvedPath string
}{
	{"/browse/","/browse","/home/martial/","/home/martial"},
	{"/browse/","/browse/","/home/martial/","/home/martial"},
	{"/browse/","/browse/mama","/home/martial/","/home/martial/mama"},
	{"/browse/","/browse/mama","/home/martial","/home/martial/mama"},
	{"/browse/","/browse/mama/","/home/martial","/home/martial/mama"},
	{"/browse/film","/browse/film/mama/","/home/martial","/home/martial/mama"},
}

func TestResolveRealPath(t *testing.T) {
	for _,element := range testData {
		storageHandler := New(element.basePattern, element.internalPath)
		value := resolveRealPath(storageHandler,element.urlRequest)
		assertEquals(value,element.resolvedPath,t)
	}
}

func assertEquals(first string, second string,t *testing.T){
	if(first == second){
		return;
	}else{
		t.Errorf("%s != %s",first,second)
		t.Fail()
	}
}
