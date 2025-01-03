package case3

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
)

func Solution(entries []os.DirEntry) []int {
	filesChan := make(chan string, len(entries))
	resultMap := make(map[int]struct{})
	var mutex sync.Mutex

	workerCount := runtime.NumCPU() * 2

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// 파일 경로를 채널에 전송
	for _, entry := range entries {
		filesChan <- fmt.Sprintf("./test/%s", entry.Name())
	}
	close(filesChan)

	// 워커들이 파일을 직접 처리
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			localMap := make(map[int]struct{}, 10000)

			for filename := range filesChan {
				file, err := os.Open(filename)
				if err != nil {
					continue
				}

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					if num, err := strconv.Atoi(scanner.Text()); err == nil {
						localMap[num] = struct{}{}
					}
				}
				file.Close()

				// 주기적으로 메인 맵에 병합 (메모리 사용량 조절)
				if len(localMap) > 10000 {
					mutex.Lock()
					for k := range localMap {
						resultMap[k] = struct{}{}
					}
					mutex.Unlock()
					localMap = make(map[int]struct{}, 10000)
				}
			}

			// 남은 데이터 병합
			mutex.Lock()
			for k := range localMap {
				resultMap[k] = struct{}{}
			}
			mutex.Unlock()
		}()
	}

	wg.Wait()

	result := make([]int, 0, len(resultMap))
	for num := range resultMap {
		result = append(result, num)
	}
	sort.Ints(result)
	return result
}
