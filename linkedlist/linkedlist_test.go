package linkedlist

import (
	"fmt"
	"testing"
)

func TestAppendToTail(t *testing.T) {
	n3 := LinkedListNode{nil, 3}
	n2 := LinkedListNode{&n3, 2}
	n1 := LinkedListNode{&n2, 1}
	newVal := 0

	n1.AppendToTail(newVal)

	if n3.next == nil || n3.next.data != newVal {
		t.Errorf("AppendToTail(%d)", newVal)
		printLinkedList(&n1)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	n5 := LinkedListNode{nil, 5}
	n4 := LinkedListNode{&n5, 2}
	n3 := LinkedListNode{&n4, 3}
	n2 := LinkedListNode{&n3, 2}
	n1 := LinkedListNode{&n2, 1}

	DeleteDuplicates(&n1)

	if n3.next != nil && n3.next.data != 5 {
		t.Errorf("DeleteDuplicates(%v)", n1)
		printLinkedList(&n1)
	}
}

func TestNthToLast(t *testing.T) {
	n5 := LinkedListNode{nil, 5}
	n4 := LinkedListNode{&n5, 2}
	n3 := LinkedListNode{&n4, 3}
	n2 := LinkedListNode{&n3, 2}
	n1 := LinkedListNode{&n2, 1}

	in := 3
	res := NthToLast(&n1, in)

	if res != &n3 {
		t.Errorf("NthToLast(%p, %d) = %p, want %p", &n1, in, &res, &n3)
		printLinkedList(&n1)
	}
}

func TestDeleteNode(t *testing.T) {
	n5 := LinkedListNode{nil, 5}
	n4 := LinkedListNode{&n5, 2}
	n3 := LinkedListNode{&n4, 3}
	n2 := LinkedListNode{&n3, 2}
	n1 := LinkedListNode{&n2, 1}

	res := DeleteNode(&n3)

	if n3.data != 2 || !res {
		t.Errorf("DeleteNode(%p) = %t, want %t and n3.data = %d want %d", &n3, res, true, &n3.data, 2)
		printLinkedList(&n1)
	}
}

func TestBreakList(t *testing.T) {
	n5 := LinkedListNode{nil, 4}
	n4 := LinkedListNode{&n5, 2}
	n3 := LinkedListNode{&n4, 3}
	n2 := LinkedListNode{&n3, 2}
	n1 := LinkedListNode{&n2, 5}

	BreakList(&n1, 3)
	node := &n1
	i := 0

	for node.next != nil {
		if (i == 3 || i == 4) && node.data < 3 {
			t.Errorf("BreakList(%v, 3)", n1)
			printLinkedList(&n1)
		}
		node = node.next
		i++
	}
}

func printLinkedList(node *LinkedListNode) {
	i := 1
	for node.next != nil {
		fmt.Printf("node%d: %d <%p>\n", i, node.data, node)
		i++
		node = node.next
	}
	fmt.Printf("node%d: %d <%p>\n", i, node.data, node)
}
