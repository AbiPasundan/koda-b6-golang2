package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	// lib.Slice()
	// initial value
	scores := []int{50, 75, 66, 20, 32, 90}
	container := make([]int, len(scores))

	m := scores[2]
	fmt.Print("\n specific index ", m, "\n")

	copy(container, scores)

	i := slices.Index(scores, 66)
	test := scores[:i]
	fmt.Print("\n", test, "\n")

	rawI := scores[:i]
	fmt.Print("\n raw i", test, "\n")
	result := append(rawI, 88)
	fmt.Print("\n", result, "\n")

	j := slices.Index(container, 20)
	testj := container[j:]
	fmt.Print("\n testj ", testj, "\n")

	rawJ := container[j:]

	var finalResult []int = append(result, rawJ...)
	fmt.Println(finalResult)

	alter := append(test, rawJ...)
	fmt.Print("\n alternative ", alter, "\n")

	alterIndex := rand.Intn(len(alter))
	fmt.Print("\n index  ", alterIndex, "\n")

	// fmt.Println(rand.Intn(len(alter)))
	// fmt.Println(rand.Intn(len(alter)))
	// fmt.Println(rand.Intn(len(alter)))

	randomResult := slices.Insert(alter, alterIndex, 88)
	fmt.Print("\n Next Result  ", randomResult, "\n")

}
