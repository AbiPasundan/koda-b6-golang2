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

// hello world
type hello struct {
	world string
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

}
