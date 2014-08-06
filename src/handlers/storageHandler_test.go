package handlers

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"
)

var testData = []struct {
	basePattern  string
	urlRequest   string
	internalPath string
	resolvedPath string
}{
	{"/browse/", "/browse", "/home/martial/", "/home/martial"},
	{"/browse/", "/browse/", "/home/martial/", "/home/martial"},
	{"/browse/", "/browse/mama", "/home/martial/", "/home/martial/mama"},
	{"/browse/", "/browse/mama", "/home/martial", "/home/martial/mama"},
	{"/browse/", "/browse/mama/", "/home/martial", "/home/martial/mama"},
	{"/browse/film", "/browse/film/mama/", "/home/martial", "/home/martial/mama"},
}

//========================================================================================
//  Tests
//========================================================================================
func TestResolveRealPath(t *testing.T) {
	for _, element := range testData {
		storageHandler := New(element.basePattern, element.internalPath)
		value := resolveRealPath(storageHandler, element.urlRequest)
		assertEquals(value, element.resolvedPath, t)
	}
}

func TestShouldHasIndexBeTrueWhenTheFolderContainsAnIndexFile(t *testing.T) {
	tmpPath, _ := ioutil.TempDir(os.TempDir(), "go-test")
	defer os.RemoveAll(tmpPath)
	os.Create(filepath.Join(tmpPath, "index.html"))
	assertTrue(hasIndexFile(tmpPath),t)
}

func TestShouldHasIndexBeFalseWhenTheFolderDoesntContainAnIndexFile(t *testing.T) {
	tmpPath, _ := ioutil.TempDir(os.TempDir(), "go-test")
	defer os.RemoveAll(tmpPath)
	assertFalse(hasIndexFile(tmpPath),t)
}


//========================================================================================
//  Utils
//========================================================================================
func assertTrue(predicate bool,t *testing.T) {
	if(predicate!=true){
		t.Fail()
	}
}
func assertFalse(predicate bool,t *testing.T) {
	assertTrue(!predicate,t)
}

func assertEquals(first string, second string, t *testing.T) {
	if (first == second) {
		return;
	}else {
		t.Errorf("%s != %s", first, second)
		t.Fail()
	}
}
