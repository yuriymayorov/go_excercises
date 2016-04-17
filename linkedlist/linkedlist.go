// Package contains linkedlist utilities
package linkedlist

type LinkedListNode struct {
	next *LinkedListNode
	data int
}

// Func adds new node to tail of linked list
// Estimate time: O(n)
func (node *LinkedListNode) AppendToTail(data int) {
	end := new(LinkedListNode)
	end.data = data

	for node.next != nil {
		node = node.next
	}
	node.next = end
}

// Func deletes all duplicates in linked list
// Estimate time: O(n) Estimate required memory: O(n)
func DeleteDuplicates(node *LinkedListNode) {
	m := make(map[int]bool)
	var prev *LinkedListNode
	for node != nil {
		if _, ok := m[node.data]; ok {
			prev.next = node.next
		} else {
			m[node.data] = true
			prev = node
		}
		node = node.next
	}
}

// Func returns k element from the end
// Estimate time: O(n) Estimate required memory: O(1)
func NthToLast(head *LinkedListNode, k int) *LinkedListNode {
	if k <= 0 {
		return nil
	}

	p1, p2 := head, head

	for i := 0; i < k-1; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.next
	}
	if p2 == nil {
		return nil
	}

	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next
	}

	return p1
}

// Func deletes node
// Estimate time: O(1) Estimate required memory: O(1)
func DeleteNode(node *LinkedListNode) bool {
	if node == nil || node.next == nil {
		return false
	}

	next := node.next
	node.data = next.data
	node.next = next.next
	node = next

	return true
}

// Func parts LinkedList by 2 around x value
// Estimate time: O(n) Estimate required memory: O(n)
func Partition(pNode *LinkedListNode, x int) *LinkedListNode {
	var beforeStart *LinkedListNode
	var afterStart *LinkedListNode
	var beforeEnd *LinkedListNode
	var afterEnd *LinkedListNode

	for pNode != nil {
		node := &LinkedListNode{pNode.next, pNode.data}

		if pNode.data < x {
			if beforeStart == nil {
				beforeStart = node
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = node
				beforeEnd = node
			}
		} else {
			if afterStart == nil {
				afterStart = node
				afterEnd = afterStart
			} else {
				afterEnd.next = node
				afterEnd = node
			}
		}
		pNode = pNode.next
	}

	if beforeStart == nil {
		return afterStart
	}

	beforeEnd.next = afterStart
	return beforeStart
}

// Func sums 2 LinkedList
// Estimate time: O(n) Estimate required memory: O(n)
func AddLists(n1 *LinkedListNode, n2 *LinkedListNode, carry int) *LinkedListNode {
	if n1 == nil && n2 == nil && carry == 0 {
		return nil
	}

	result := LinkedListNode{nil, carry}

	value := carry

	var n1Next, n2Next *LinkedListNode

	if n1 != nil {
		value += n1.data
		n1Next = n1.next
	}

	if n2 != nil {
		value += n2.data
		n2Next = n2.next
	}

	carryNext := 0

	if value >= 10 {
		carryNext = 1
	}

	result.data = value % 10

	if n1 != nil || n2 != nil || value >= 10 {
		more := AddLists(n1Next, n2Next, carryNext)
		result.next = more
	}

	return &result
}
