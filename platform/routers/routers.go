package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	// Get the current working directory
	curr_wd, err := os.Getwd()

	if err != nil {

		fmt.Println(err)

		os.Exit(1)
	}

	// Print the current working directory
	fmt.Println((curr_wd))
	BASE_DIIR := filepath.Dir(filepath.Dir(curr_wd))
	fmt.Println(path.Join(BASE_DIIR, "asset/images"))
}
