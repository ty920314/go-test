package main

import "fmt"

func pointer() {
	var value string = "asjdoihr segrb1 rg3223!"
	ptr := &value
	fmt.Printf("value:%s\n", value)
	fmt.Printf("address:%p\n", ptr)
	fmt.Printf("type:%T\n", ptr)
	res := *ptr
	fmt.Printf("yin:%s\n", res)

}
