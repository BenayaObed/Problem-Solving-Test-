package main

import (
	"fmt"
)

func getDenseRanking(scores []int) []int {
	rank := []int{}
	rankMap := make(map[int]int)
	rankCounter := 1

	for _, score := range scores {
		if _, exists := rankMap[score]; !exists {
			rankMap[score] = rankCounter
			rank = append(rank, score)
			rankCounter++
		}
	}
	return rank
}

func getGitsRankings(scores []int, gitsScores []int) []int {
	rankings := getDenseRanking(scores)
	results := []int{}

	for _, gitsScore := range gitsScores {
		position := 1
		for position <= len(rankings) && gitsScore < rankings[position-1] {
			position++
		}
		results = append(results, position)
	}
	return results
}

func main() {
	var n int
	fmt.Scan(&n)
	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	var m int
	fmt.Scan(&m)
	gitsScores := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&gitsScores[i])
	}

	gitsRankings := getGitsRankings(scores, gitsScores)

	for _, rank := range gitsRankings {
		fmt.Print(rank, " ")
	}
	fmt.Println()
}
