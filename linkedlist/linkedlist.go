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


func NthToLast(head *LinkedListNode, k int) (int, *LinkedListNode) {
	if (head == nil) {
		return 0, head
	}

	i, res := NthToLast(head.next, k) + 1, head

	if (i == k) {
		return k, head
	}

	return i, head
}