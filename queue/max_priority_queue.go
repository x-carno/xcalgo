package queue

import "container/heap"

type MaxPQItem struct {
	value    string
	priority int
	index    int
}

type MaxPriorityQueue []*MaxPQItem

func (mpq MaxPriorityQueue) Len() int {
	return len(mpq)
}

func (mpq MaxPriorityQueue) Less(i, j int) bool {
	return mpq[i].priority > mpq[j].priority
}

func (mpq MaxPriorityQueue) Swap(i, j int) {
	mpq[i], mpq[j] = mpq[j], mpq[i]
	mpq[i].index = i
	mpq[j].index = j
}

func (mpq *MaxPriorityQueue) Push(x interface{}) {
	n := len(*mpq)
	item := x.(*MaxPQItem)
	item.index = n
	*mpq = append(*mpq, item)
}

func (mpq *MaxPriorityQueue) Pop() interface{} {
	old := *mpq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*mpq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (mpq *MaxPriorityQueue) Update(item *MaxPQItem, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(mpq, item.index)
}
