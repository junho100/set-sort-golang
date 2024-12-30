package internal

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"

	"math/rand"
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

	if err := os.Mkdir("test", 0750); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		filePath := fmt.Sprintf("./test/file_%s.csv", fmt.Sprint(i))
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		wr := csv.NewWriter(bufio.NewWriter(file))

		for j := 0; j < 10000; j++ {
			wr.Write([]string{fmt.Sprint(rand.Intn(999999) + 1)})
		}
		wr.Flush()
	}
}
