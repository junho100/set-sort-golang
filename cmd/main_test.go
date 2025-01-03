package main

import (
	"os"
	"set-sort-golang/internal"
	"set-sort-golang/internal/case1"
	"set-sort-golang/internal/case2"
	"set-sort-golang/internal/case3"
	"testing"
)

func setupTestData() []os.DirEntry {
	internal.GenerateData()
	entries, _ := os.ReadDir("./test/")
	return entries
}

func BenchmarkCase1(b *testing.B) {
	entries := setupTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		case1.Solution(entries)
	}
}

func BenchmarkCase2(b *testing.B) {
	entries := setupTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		case2.Solution(entries)
	}
}

func BenchmarkCase3(b *testing.B) {
	entries := setupTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		case3.Solution(entries)
	}
}
