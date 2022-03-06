package main

import (
	"container/list"
	"fmt"
)

func Display(queue *list.List) {
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
