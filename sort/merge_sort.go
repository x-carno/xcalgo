package sort

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/panjf2000/ants"
)

type MergeSort struct {
	items []int
	size  int
	// left  int
	// right int

	concurrentNum int

	grPool *ants.Pool
	wg     sync.WaitGroup
}

func NewMergeSort(s []int) *MergeSort {
	var cn int
	if runtime.NumCPU() < 2 {
		cn = 2
	} else {
		cn = runtime.NumCPU()
	}

	p, _ := ants.NewPool(1 << len(s))
	return &MergeSort{
		items: s,
		size:  len(s),
		// left:          0,
		// right:         len(s),
		concurrentNum: cn,
		grPool:        p,
	}
}

func (m *MergeSort) SetConcurrentDepth(n int) {
	// if n > runtime.NumCPU() {
	// 	fmt.Println("max concurrent sort number is equal to CPU number:", runtime.NumCPU())
	// 	m.concurrentNum = runtime.NumCPU()
	// 	return
	// }
	if n < 2 {
		fmt.Println("concurrent sort number should not smaller than 2")
		m.concurrentNum = 2
		return
	}
	m.concurrentNum = n
}

func (m *MergeSort) Sort() []int {
	if m.size >= 1<<10 {
		piece := m.size / m.concurrentNum
		for i := 0; i < m.concurrentNum; i++ {
			m.wg.Add(1)
			if i < m.concurrentNum-1 {
				go func(le, ri int) {
					defer m.wg.Done()
					m.sort(le, ri)
				}(i*piece, (i+1)*piece)
			} else {
				go func(le, ri int) {
					defer m.wg.Done()
					m.sort(le, ri)
				}(i*piece, m.size)
			}
		}
		m.wg.Wait()

		for i := 0; i < m.concurrentNum-1; i++ {
			if i < m.concurrentNum-2 {
				m.merge(0, (i+1)*piece, (i+2)*piece)
			} else {
				m.merge(0, (i+1)*piece, m.size)
			}
		}
	} else {
		m.sort(0, m.size)
	}

	return m.items
}

func (m *MergeSort) sort(le, ri int) {
	if le+1 == ri {
		return
	}
	if le+2 == ri {
		m.sort2(le)
		return
	}
	if le+3 == ri {
		m.sort3(le)
		return
	}

	m.sort(le, le+(ri-le)/2)
	m.sort(le+(ri-le)/2, ri)
	m.merge(le, le+(ri-le)/2, ri)
}

func (m *MergeSort) merge(le, mi, ri int) {
	// if m.items[mi] > m.items[mi-1] {
	// 	return
	// }

	s := make([]int, ri-le)
	i, j := 0, 0
	remainLeft := true
	for {
		if m.items[le+i] < m.items[mi+j] {
			s[i+j] = m.items[le+i]
			i++
			if i == mi-le {
				remainLeft = false
				break
			}
		} else {
			s[i+j] = m.items[mi+j]
			j++
			if j == ri-mi {
				break
			}
		}
	}
	for {
		if i+j == ri-le {
			break
		}
		if remainLeft {
			s[i+j] = m.items[le+i]
			i++
		} else {
			s[i+j] = m.items[mi+j]
			j++
		}
	}

	for i := 0; i < len(s); i++ {
		m.items[le+i] = s[i]
	}
}

func (m *MergeSort) swap(i, j int) {
	k := m.items[i]
	m.items[i] = m.items[j]
	m.items[j] = k
}

func (m *MergeSort) sort2(le int) {
	if m.items[le] > m.items[le+1] {
		m.swap(le, le+1)
	}
}

func (m *MergeSort) sort3(le int) {
	if m.items[le] > m.items[le+2] {
		m.swap(le, le+2)
	}
	if m.items[le] > m.items[le+1] {
		m.swap(le, le+1)
	}
	if m.items[le+1] > m.items[le+2] {
		m.swap(le+1, le+2)
	}
}

func MergeSortFunc(s []int) []int {
	if len(s) == 1 {
		return s
	}
	left := MergeSortFunc(s[:len(s)/2])
	right := MergeSortFunc(s[len(s)/2:])

	s = merge(left, right)
	return s
}

func MergeSortConcurrent(s []int) []int {
	// fmt.Println(runtime.NumGoroutine())
	if len(s) == 1 {
		return s
	}
	left := make(chan []int)
	right := make(chan []int)
	// if len(s) < 10 {
	// 	grPool.Submit(func() {
	// 		left <- MergeSort(s[:len(s)/2])
	// 	})
	// 	grPool.Submit(func() {
	// 		right <- MergeSort(s[len(s)/2:])
	// 	})
	// } else {
	// 	grPool.Submit(func() {
	// 		left <- MergeSortConcurrent(s[:len(s)/2])
	// 	})
	// 	grPool.Submit(func() {
	// 		right <- MergeSortConcurrent(s[len(s)/2:])
	// 	})
	// }
	go func() {
		left <- MergeSortConcurrent(s[:len(s)/2])
	}()
	go func() {
		right <- MergeSortConcurrent(s[len(s)/2:])
	}()
	s = merge(<-left, <-right)
	return s
}

// left and right is already sorted
func merge(left []int, right []int) []int {
	s := make([]int, len(left)+len(right))
	i, j := 0, 0
	remainLeft := true
	for {
		if left[i] < right[j] {
			s[i+j] = left[i]
			i++
			if i == len(left) {
				remainLeft = false
				break
			}
		} else {
			s[i+j] = right[j]
			j++
			if j == len(right) {
				break
			}
		}
	}
	for {
		if i+j == len(s) {
			break
		}
		if remainLeft {
			s[i+j] = left[i]
			i++
		} else {
			s[i+j] = right[j]
			j++
		}
	}
	return s
}
