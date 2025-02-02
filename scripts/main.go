package main

import (
	"fmt"
	"os"
)

const (
	mainReadme    = "README.md"
	markerPattern = `<INSERT:(\w+)>`
)

func main() {
	err := UpdateReadme(mainReadme, markerPattern)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	fmt.Println("README.md was successfully updated")
}
