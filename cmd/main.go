package main

import (
	"log"
	"os"
	"set-sort-golang/internal"
	"set-sort-golang/internal/case2"
)

func main() {
	internal.GenerateData()

	entries, err := os.ReadDir("./test/")
	if err != nil {
		log.Fatal(err)
	}

	case2.Solution(entries)
}
