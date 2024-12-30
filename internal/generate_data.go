package internal

import (
	"log"
	"os"
	"slices"
)

func GenerateData() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	if slices.ContainsFunc(entries, func(e os.DirEntry) bool {
		return e.Name() == "test"
	}) {
		if err := os.RemoveAll("./test"); err != nil {
			log.Fatal(err)
		}
	}
}
