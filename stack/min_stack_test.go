package stack

import (
	"testing"
)

// ["MinStack","push","push","push","top","pop","min","pop","min","pop","push","top","min","push","top","min","pop","min"]
// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

//[null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483647,null,-2147483648,-2147483648,null,2147483647]
func TestMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(2147483646)
	// minStack.Push(2147483646)
	// minStack.Push(2147483647)
	minStack.Pop()
	minStack.Top()
	minStack.Pop()
	minStack.Pop()

	// fmt.Println(minStack.Top())
	// minStack.Pop()
	// fmt.Println(minStack.Min())
	// minStack.Pop()
	// fmt.Println(minStack.Min())
	// minStack.Pop()

	// minStack.Push(2147483647)
	// fmt.Println(minStack.Top())
	// fmt.Println(minStack.Min())
	// minStack.Push(-2147483648)
	// fmt.Println(minStack.Top())
	// fmt.Println(minStack.Min())
	// minStack.Pop()
	// fmt.Println(minStack.Min())
}
