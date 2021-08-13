package main

import (
	"container/list"
	"testing"
)

func Test_list(t *testing.T) {
	lst := list.New()
	lst.PushBack("test")
	head := lst.Front()
	x := head.Value.(string)
	t.Logf("x='%v'\n", x)
}

func Test_heap(t *testing.T) {

}
