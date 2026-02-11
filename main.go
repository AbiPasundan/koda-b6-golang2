package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slice()
	// var container []int
	scores := []int{50, 75, 66, 20, 32, 90}
	fmt.Println(scores)
	container := make([]int, len(scores))
	copy(container, scores)
	// fmt.Println(container)

	i := slices.Index(scores, 66)
	// fmt.Println(i)
	rawI := scores[:i+1]

	// fmt.Println(append(rawI, 88))
	result := append(rawI, 88)
	// fmt.Println(rawI)

	// fmt.Println("oihdoshfoehf")

	j := slices.Index(container, 20)
	// fmt.Println(j)
	rawJ := container[j:]
	// fmt.Println(rawJ)

	var finalResult []int = append(result, rawJ...)

	fmt.Println(finalResult)

	for v := range finalResult {
		fmt.Println(finalResult[v])
	}

	// fmt.Println(scores)
	// fmt.Println("test")
	// fmt.Println(i)
	// fmt.Println(scores[i])
	// numbers := []int{1, 2, 3}
	// for v := range scores {
	// 	fmt.Println(scores[v])
	// 	append(container, scores)
	// }
	// test := slices.Contains(scores, 66)
	// i := slices.Index(scores, 66)
	// j := slices.Index(scores, 66)
	// fmt.Println("j")
	// fmt.Println(j)
	// rawI := scores[:i+1]
	// // rawJ := scores[i+1:]
	// var rawJ []int
	// rawJ = scores[j+1:]
	// fmt.Println("rawJ")
	// fmt.Println(rawJ)

	// // fmt.Println(append(scores, 8))

	// // fmt.Println(i)
	// // fmt.Println(j)
	// // fmt.Println(test)
	// // fmt.Println(scores[2])
	// // fmt.Println("rawI")
	// // fmt.Println(rawI)
	// // fmt.Println("rawJ")
	// // fmt.Println(rawJ)
	// fmt.Println("didieu yeuh")
	// fmt.Println(append(rawI, 88))
	// fmt.Println(rawJ)
	// // result := append(rawI, rawJ...)
	// // fmt.Println(result)
	// fmt.Println("result")
	// fmt.Println(append(rawI, rawJ...))

	// for _, v := range scores {
	// 	container = append(container, v)
	// }

	// fmt.Println(container)

	// numbers := []int{0, 42, 8}
	// fmt.Println(slices.Equal(numbers, []int{0, 42, 8}))
	// fmt.Println(slices.Equal(numbers, []int{10}))

	// numbers := []int{0, 1, 2, 3}
	// fmt.Println(slices.Contains(numbers, 2))
	// fmt.Println(slices.Contains(numbers, 4))
}
