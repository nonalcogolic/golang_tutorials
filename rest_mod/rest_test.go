package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	r, err := http.Get("http://localhost:8000/phones")
	if err != nil {
		t.Errorf("Error in GET")
	}

	if r.Status != "200 OK" {
		t.Errorf("Error in TestGet()")
	}
	fmt.Println(r)

	r, err = http.Get("http://localhost:8000/phones/John")
	if err != nil {
		t.Errorf("Error in GET")
	}

	if r.Status == "200 OK" {
		t.Errorf("Error in TestGet()")
	}

}
