package main

import "fmt"

func append(scores []int, value int) []int {
	var lenght int = len(scores)

	var newScores []int = make([]int, lenght+1)

	for i := 0; i < lenght; i++ {
		newScores[i] = scores[i]
	}

	newScores[lenght] = value

	return newScores

}

func main() {
	var scores []int = []int{90, 70, 50, 60, 85}

	scores = append(scores, 75)

	var lenght int = len(scores)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
