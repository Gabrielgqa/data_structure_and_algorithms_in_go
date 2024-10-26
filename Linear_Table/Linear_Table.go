// Linear Table example

package main

import "fmt"

func main() {
	var scores []int = []int{90, 70, 50, 60, 85}

	var lenght int = len(scores)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
