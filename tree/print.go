package tree

import "fmt"

// 结构的方法 go的参数都是传值
//拷贝
func (node Node) Print() {
	fmt.Print(node.Value)
}
