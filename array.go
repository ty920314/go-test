package main

import "fmt"

func array() {
	//var a [3]int
	//for i, v := range a {
	//	fmt.Printf("%d %d\n", i, v)
	//}
	q := [...]int{1, 2, 3, 325, 3, 46, 457, 56, 8}
	for _, v := range q {
		fmt.Printf("%d\n", v)
	}
}
