package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	ll := make([]int, 0)
	for {
		if head == nil {
			break
		}
		ll = append(ll, 0)
		copy(ll[1:], ll)
		ll[0] = head.Val
		head = head.Next
	}
	return ll
}

func ReverseList(head *ListNode) *ListNode {
	var tmp *ListNode
	var reverse *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = reverse
		reverse = head
		head = tmp
	}
	return reverse
}
