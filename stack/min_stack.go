package stack

import (
	"log"
	"math"
)

type MinStack struct {
	items  []int
	minVal int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		items:  make([]int, 0),
		minVal: math.MaxInt,
	}
}

func (ms *MinStack) Push(x int) {
	if x < ms.minVal {
		ms.minVal = x
	}
	ms.items = append(ms.items, x, ms.minVal)

}

func (ms *MinStack) Pop() {
	// prevent panic
	if len(ms.items) == 0 {
		log.Fatalln("stack is empty")
	}
	ms.items = ms.items[:len(ms.items)-2]
	if len(ms.items) > 0 {
		ms.minVal = ms.items[len(ms.items)-1]
	} else {
		ms.minVal = math.MaxInt
	}
}

func (ms *MinStack) Top() int {
	if len(ms.items) == 0 {
		log.Fatalln("stack is empty")
	}
	return ms.items[len(ms.items)-2]
}

func (ms *MinStack) Min() int {
	return ms.minVal
}
