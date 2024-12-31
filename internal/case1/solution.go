package case1

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Solution(entries []os.DirEntry) {
	s := NewSet()

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
