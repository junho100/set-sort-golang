package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"set-sort-golang/internal"
	"set-sort-golang/internal/case1"
	"set-sort-golang/internal/case2"
	"set-sort-golang/internal/case3"
	"set-sort-golang/internal/case4"
)

func runCase1() []int {
	entries, _ := os.ReadDir("./test/")
	return case1.Solution(entries)
}

func runCase2() []int {
	entries, _ := os.ReadDir("./test/")
	return case2.Solution(entries)
}

func runCase3() []int {
	entries, _ := os.ReadDir("./test/")
	return case3.Solution(entries)
}

func runCase4() []int {
	entries, _ := os.ReadDir("./test/")
	return case4.Solution(entries)
}

func profileCase(caseNum int, runFunc func() []int) {
	// 트레이스 파일 생성
	f, err := os.Create(fmt.Sprintf("traces/trace_case%d.out", caseNum))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 트레이스 시작
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop() // trace.Stop() 호출 보장
	// CPU 프로파일
	cpuFile, err := os.Create(fmt.Sprintf("profiles/cpu_case%d.prof", caseNum))
	if err != nil {
		log.Fatal(err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// 솔루션 실행
	result := runFunc()
	fmt.Printf("Case%d Result Length: %d\n", caseNum, len(result))

	// 메모리 프로파일
	memFile, err := os.Create(fmt.Sprintf("profiles/mem_case%d.prof", caseNum))
	if err != nil {
		log.Fatal(err)
	}
	defer memFile.Close()

	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 테스트 데이터 생성
	internal.GenerateData()

	// 각 케이스별로 프로파일링
	profileCase(1, runCase1)
	profileCase(2, runCase2)
	profileCase(3, runCase3)
	profileCase(4, runCase4)
}
