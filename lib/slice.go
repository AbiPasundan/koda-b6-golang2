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

// type tech struct {
// 	tech string
// }

// type man struct {
// 	man []tech
// }

// type obj struct {
// 	str [][][]man
// }

type Tech struct {
	academy string
}

type Man struct {
	tech Tech
}

type StrItem struct {
	man []Man
}

type Object struct {
	str [][][]StrItem
}

// task 4 => apel
type is struct {
	is string
}
type fruit struct {
	is string
}

// di dalam favorite harus ada fruit. ...{ favorite: []favorite{ ...
type favorite struct {
	fruit fruit
}

// didalam my harus ada favorite. var my = []my{ { favorite: []favorite{ ...
type my struct {
	favorite []favorite
}

// type item struct {
// 	favourite []favourite
// }

// type fruits struct {
// 	fruits []is
// }
// type is struct {
// 	is []string
// }

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

	targetMan := Man{
		tech: Tech{
			academy: "Tech Academy",
		},
	}

	strSlice := make([][][]StrItem, 4)
	strSlice[3] = make([][]StrItem, 2)
	strSlice[3][1] = make([]StrItem, 3)

	strSlice[3][1][2] = StrItem{
		man: []Man{targetMan},
	}

	obj := Object{
		str: strSlice,
	}

	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

	// task 4 = favorite[]int{5},
	// var my []favorite = []favorite{
	// {
	// 	favorite: fruits{""},
	// },
	// }

	// fmt.Println(my[0].favourite)
	// fmt.Println(my[0].favourite[3].fruits.is)

	// var favorite []favourite

	// var my = [1]slice{
	// 	favorite: []slice{
	// 		"ass",
	// 	},
	// }

	var my = []my{
		{
			favorite: []favorite{
				{
					fruit: fruit{
						is: "Apel",
					},
					// is{"check"},
					// fruit: fruit{
					// 	is: is{"Apel"},
					// },
				},
			},
		},
	}

	fmt.Println(my[0].favorite[0].fruit.is)

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
