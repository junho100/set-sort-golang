package main

import (
	"log"
	"os"
	"set-sort-golang/internal"
	"set-sort-golang/internal/case1"
)

func main() {
	internal.GenerateData()

	entries, err := os.ReadDir("./test/")
	if err != nil {
		log.Fatal(err)
	}

	case1.Solution(entries)
}
