package case1

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"set-sort-golang/internal"
	"sort"
	"strconv"
)

func Solution() {
	s := internal.NewSet()

	entries, err := os.ReadDir("./test/")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		file, _ := os.Open(fmt.Sprintf("./test/%s", entry.Name()))
		rdr := csv.NewReader(bufio.NewReader(file))

		rows, _ := rdr.ReadAll()

		for _, row := range rows {
			i, err := strconv.Atoi(row[0])
			if err != nil {
				log.Fatal(err)
			}
			s.Add(i)
		}
	}

	sl := s.ToSlice()
	sort.Ints(sl)
	fmt.Println(sl)
}
