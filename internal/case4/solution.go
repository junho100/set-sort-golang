package case4

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
)

func Solution(entries []os.DirEntry) []int {
	var wg sync.WaitGroup
	numOfWorkers := runtime.NumCPU() * 4
	wg.Add(numOfWorkers)

	fileChannel := make(chan string, len(entries))
	for _, entry := range entries {
		fileChannel <- fmt.Sprintf("./test/%s", entry.Name())
	}
	close(fileChannel)

	s := NewSet()

	for i := 0; i < numOfWorkers; i++ {
		go func() {
			defer wg.Done()

			tempMap := make(map[int]struct{})

			for filePath := range fileChannel {
				file, err := os.Open(filePath)
				if err != nil {
					log.Fatal(err)
				}
				rdr := csv.NewReader(bufio.NewReader(file))
				rows, _ := rdr.ReadAll()
				for _, row := range rows {
					i, err := strconv.Atoi(row[0])
					if err != nil {
						log.Fatal(err)
					}
					tempMap[i] = struct{}{}
				}
				file.Close()

				if len(tempMap) > 10000 {
					keys := make([]int, 0, len(tempMap))
					for k := range tempMap {
						keys = append(keys, k)
					}
					s.AddAll(keys)
					tempMap = make(map[int]struct{})
				}
			}

			keys := make([]int, 0, len(tempMap))
			for k := range tempMap {
				keys = append(keys, k)
			}
			s.AddAll(keys)
		}()
	}

	wg.Wait()
	sl := s.ToSlice()
	sort.Ints(sl)
	return sl
}
