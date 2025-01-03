package case2

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
)

func Solution(entries []os.DirEntry) {
	var wg sync.WaitGroup
	s := NewSet()

	wg.Add(len(entries))
	for _, entry := range entries {
		go func(entry os.DirEntry) {
			defer wg.Done()

			file, _ := os.Open(fmt.Sprintf("./test/%s", entry.Name()))
			rdr := csv.NewReader(bufio.NewReader(file))
			rows, _ := rdr.ReadAll()

			// Lock 획득 후 다른 고루틴 대기로 인해서 오버헤드 증가 -> context switching 증가
			// 임시 Set 만든 후 병합 -> Lock 횟수 감소
			tempSet := make(map[int]struct{})
			for _, row := range rows {
				i, err := strconv.Atoi(row[0])
				if err != nil {
					log.Fatal(err)
				}
				tempSet[i] = struct{}{}
			}

			s.m.Lock()
			for k, v := range tempSet {
				s.data[k] = v
			}
			s.m.Unlock()
		}(entry)
	}

	wg.Wait()
	sl := s.ToSlice()
	sort.Ints(sl)
	fmt.Println(sl)
}
