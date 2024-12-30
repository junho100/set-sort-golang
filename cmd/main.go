package main

import (
	"set-sort-golang/internal"
	"set-sort-golang/internal/case1"
)

func main() {
	internal.GenerateData()

	case1.Solution()
}
