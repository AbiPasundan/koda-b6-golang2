package main

import "fmt"

type best struct {
	best string
}
type are struct {
	the best
}
type we struct {
	are are
}

// task 2 => hello world
type hello struct {
	world string
}

// task 4 =? apel
type favorit struct {
	favorite []int
}

// task 5
type first struct {
	first  []int
	second []int
}

func main() {
	// task 1
	var we we = we{
		are: are{
			the: best{
				best: "Koda",
			},
		},
	}
	fmt.Println(we.are.the.best)

	// task 2
	var hello hello = hello{
		world: "Hello World",
		// world: []string{
		// 	world: "",
		// },
	}
	fmt.Println(hello.world)

	// task 4 = favorite[]int{5},
	my := [1]favorit{}

	fmt.Println(my[0].favorite)

	// task 5
	var num first = first{
		first: []int{
			1, 30,
		},
		second: []int{
			1, 2,
		},
	}

	fmt.Println(num.first[1] + num.second[1])

}
