package main

import (
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {

	// fmt.Print("Testing this testno surve")
	var myh myHandler

	h := NoSruve(&myh)

	switch v := h.(type) {

	case http.Handler:

	default:
		t.Error("NoSurve function is failing ", v)

	}
}

func TestSessionLoad(t *testing.T) {

	// fmt.Print("Testing this testno surve")
	var myh myHandler

	h := SessionLoad(&myh)

	switch v := h.(type) {

	case http.Handler:

	default:
		t.Error("Session Load function is failing ", v)

	}
}
