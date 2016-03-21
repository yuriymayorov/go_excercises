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

func Partition(node *LinkedListNode, x int) *LinkedListNode {
	var beforeStart *LinkedListNode
	var afterStart *LinkedListNode
	var beforeEnd *LinkedListNode
	var afterEnd *LinkedListNode

	for node != nil {
		next := node.next
		node.next = nil

		if node.data < x {
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
		node = next
	}

	if beforeStart == nil {
		return afterStart
	}

	beforeEnd.next = afterStart
	return beforeStart
}
