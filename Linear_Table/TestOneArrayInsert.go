// Insert a value into an array in a especific index.

package main

import "fmt"

func insert(scores []int, newScores []int, lenght int, value int, index int) {
	for i := 0; i < lenght; i++ {
		if i < index {
			newScores[i] = scores[i]
		} else {
			newScores[i+1] = scores[i]
		}
	}

	newScores[index] = value

}

func main() {
	var scores []int = []int{90, 70, 50, 60, 85}

	var lenght int = len(scores)

	var newScores []int = make([]int, lenght+1)

	insert(scores, newScores, lenght, 75, 2)

	scores = newScores

	for i := 0; i < lenght+1; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
