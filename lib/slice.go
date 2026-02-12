package lib

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

// task 3 => Tech Academy

type tech struct {
	tech string
}

type man struct {
	man []tech
}

type obj struct {
	str [][][]man
}

// task 4 => apel
type favorite struct {
	favourite []fruits
}
type fruits struct {
	fruits []is
}
type is struct {
	is []string
}

// task 5
type first struct {
	first  []int
	second []int
}

func Slice() {
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
	}
	fmt.Println(hello.world)

	// task 3
	// var obj obj = struct{
	// 	str [][][]struct
	// }
	// fmt.Println(obj)

	// task 4 = favorite[]int{5},
	// var my []favorite = []favorite{
	// {
	// 	favorite: fruits{""},
	// },
	// }

	// fmt.Println(my[0].favourite)
	// fmt.Println(my[0].favourite[3].fruits.is)

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

	// minitask wjdbiwfubew
}
