package main

import (
	"container/list"
	"fmt"
)

func listTest() {
	l := list.New()

	l.PushBack("asd")
	l.PushFront("89n")
	ele := l.PushBack("gr09")
	l.InsertAfter("insertafter", ele)
	l.InsertBefore("before", ele)
	l.Remove(ele)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Printf("%s\n", i.Value)
	}
}
