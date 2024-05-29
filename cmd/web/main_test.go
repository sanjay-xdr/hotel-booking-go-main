package main

import (
	"testing"
)

func TestRun(t *testing.T) {

	// fmt.Println("THis is going to run the main func")

	err := run()

	if err != nil {
		t.Error("Not ale to run ")
	}

}
