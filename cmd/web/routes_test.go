package main

import (
	"net/http"
	"testing"

	"github.com/sanjay-xdr/cmd/internals/config"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	mux := Routes(&app)

	switch v := mux.(type) {

	case http.Handler:

	default:

		t.Error("Routes is  not working", v)
	}

}
