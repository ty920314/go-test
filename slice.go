package main

import "fmt"

func slice() {
	highRiseBuilding := [30]int{}
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间的所有元素
	fmt.Println(highRiseBuilding[:2])
}
