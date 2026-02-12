package main

import (
	"fmt"
	"slices"
)

func main() {
	// lib.Slice()
	// initial value
	scores := []int{50, 75, 66, 20, 32, 90}
	// fmt.Println(scores)
	// make empty sliceusing make
	container := make([]int, len(scores))
	// fmt.Println(container)
	// fmt.Println("container")
	// copy value from scores and append into container
	copy(container, scores)
	// take value scores until 66
	i := slices.Index(scores, 66)
	rawI := scores[:i+1]
	result := append(rawI, 88)
	// take value container start with 20
	j := slices.Index(container, 20)
	rawJ := container[j:]

	var finalResult []int = append(result, rawJ...)

	for v := range finalResult {
		// looping process
		fmt.Println(finalResult[v])
	}
	// final result
	fmt.Println(finalResult)
}
