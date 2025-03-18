package main

import (
	"fmt"
	"strings"
)

func OEISA000124(n int) []int {
	sequence := make([]int, n)
	sequence[0] = 1

	for i := 1; i < n; i++ {
		sequence[i] = sequence[i-1] + i
	}

	return sequence
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	sequence := OEISA000124(n)

	strSequence := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sequence)), "-"), "[]")
	fmt.Println("Output:", strSequence)
}
