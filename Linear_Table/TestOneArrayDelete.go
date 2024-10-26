package main

import "fmt"

func delete(scores []int, index int) []int {
	var lenght int = len(scores)

	var newScores []int = make([]int, lenght-1)

	for i := 0; i < lenght; i++ {
		if i < index {
			newScores[i] = scores[i]
		}

		if i > index {
			newScores[i-1] = scores[i]
		}
	}

	return newScores
}

func main() {
	var scores []int = []int{90, 70, 50, 80, 60, 85}

	var lenght int = len(scores)

	scores = delete(scores, 2)

	for i := 0; i < lenght-1; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
