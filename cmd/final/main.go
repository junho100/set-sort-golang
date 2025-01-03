package main

import (
	"fmt"
	"log"
	"os"
	"set-sort-golang/internal"
	"set-sort-golang/internal/case1"
	"set-sort-golang/internal/case2"
	"set-sort-golang/internal/case3"
	"set-sort-golang/internal/case4"
	"slices"
	"testing"
	"time"
)

// 각 솔루션의 결과를 검증하는 함수
func validateResults(results [][]int) error {
	if len(results) == 0 {
		return fmt.Errorf("결과가 비어있습니다")
	}

	baseResult := results[0]
	for i := 1; i < len(results); i++ {
		if !slices.Equal(baseResult, results[i]) {
			fmt.Println(results[0])
			fmt.Println(results[i])
			return fmt.Errorf("솔루션 %d의 결과가 다릅니다", i+1)
		}
	}
	return nil
}

// 벤치마크를 실행하고 결과를 반환하는 함수
func runBenchmark(fn func(b *testing.B)) testing.BenchmarkResult {
	result := testing.Benchmark(fn)
	return result
}

func main() {
	// 테스트 데이터 생성
	internal.GenerateData()
	entries, err := os.ReadDir("./test/")
	if err != nil {
		log.Fatal("테스트 데이터 읽기 실패:", err)
	}

	// 기능 검증
	fmt.Println("\n=== 기능 검증 시작 ===")
	results := [][]int{
		case1.Solution(entries),
		case2.Solution(entries),
		case3.Solution(entries),
		case4.Solution(entries),
	}

	if err := validateResults(results); err != nil {
		fmt.Printf("검증 실패: %v\n", err)
		return
	}
	fmt.Printf("모든 솔루션이 동일한 결과를 반환했습니다. (결과 크기: %d)\n", len(results[0]))

	// 벤치마크 실행
	fmt.Println("\n=== 벤치마크 테스트 시작 ===")
	benchmarks := []struct {
		name string
		fn   func(*testing.B)
	}{
		{"Case1", func(b *testing.B) {
			entries := setupTestData()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				case1.Solution(entries)
			}
		}},
		{"Case2", func(b *testing.B) {
			entries := setupTestData()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				case2.Solution(entries)
			}
		}},
		{"Case3", func(b *testing.B) {
			entries := setupTestData()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				case3.Solution(entries)
			}
		}},
		{"Case4", func(b *testing.B) {
			entries := setupTestData()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				case4.Solution(entries)
			}
		}},
	}

	for _, bm := range benchmarks {
		result := runBenchmark(bm.fn)
		fmt.Printf("\n%s:\n", bm.name)
		fmt.Printf("- 평균 실행 시간: %v\n", result.T/time.Duration(result.N))
		fmt.Printf("- 반복 횟수: %d\n", result.N)
		fmt.Printf("- 메모리 할당: %d bytes/op\n", result.AllocsPerOp())
	}
}

// 테스트 데이터 설정 함수
func setupTestData() []os.DirEntry {
	internal.GenerateData()
	entries, _ := os.ReadDir("./test/")
	return entries
}
