package main

import "fmt"

type we struct {
	we []string
}
type are struct {
	are string
}
type the struct {
	the []string
}
type best struct {
	name string
}

// task 2 => hello world
type hello struct {
	world string
	// world []string
}

// task 5
type first struct {
	first  []int
	second []int
	// world []string
}

func main() {
	// 	var we we = we{
	// 		are: []are{
	// 			the: []the{
	// 				best: []best{
	// 					name: "koda",
	// 				},
	// 			},
	// 		},
	// 	}

	// task 2
	var hello hello = hello{
		world: "Hello World",
		// world: []string{
		// 	world: "",
		// },
	}
	fmt.Println(hello.world)

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
