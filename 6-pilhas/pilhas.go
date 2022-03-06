package main

import (
	"container/list"
	"fmt"
)

func Display(queue *list.List) {
	for e := queue.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
