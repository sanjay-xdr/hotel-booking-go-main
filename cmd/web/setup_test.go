package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// run then exit

	fmt.Println("Running the main test")
	os.Exit(m.Run())
}

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
