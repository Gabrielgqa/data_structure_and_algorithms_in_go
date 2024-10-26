package main

import "fmt"

// Bubble Sorting Algorithm
func bubbleSort(scores []int, lenght int) {
	for i := 0; i < lenght-1; i++ {
		var isSwap bool = false
		for j := 0; j < lenght-1-i; j++ {
			if scores[j] > scores[j+1] {
				var temp int = scores[j]
				scores[j] = scores[j+1]
				scores[j+1] = temp
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

func main() {
	var scores []int = []int{60, 50, 95, 80, 70}

	var lenght int = len(scores)

	bubbleSort(scores, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
