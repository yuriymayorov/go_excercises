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

func TestPartition(t *testing.T) {
	n5 := LinkedListNode{nil, 4}
	n4 := LinkedListNode{&n5, 2}
	n3 := LinkedListNode{&n4, 1}
	n2 := LinkedListNode{&n3, 2}
	n1 := LinkedListNode{&n2, 5}

	res := Partition(&n1, 3)

	i := 0
	for res.next != nil {
		if (i == 3 || i == 4) && res.data < 3 {
			t.Errorf("Partition(%v, 3)", n1)
			printLinkedList(res)
		}
		i++
		res = res.next
	}
}

func TestAddLists(t *testing.T) {
	n23 := LinkedListNode{nil, 5}
	n22 := LinkedListNode{&n23, 9}
	n21 := LinkedListNode{&n22, 2}
	n13 := LinkedListNode{nil, 7}
	n12 := LinkedListNode{&n13, 1}
	n11 := LinkedListNode{&n12, 6}

	res := AddLists(&n11, &n21)

	if res == nil {
		t.Errorf("AddList(%v %v)", n11, n21)
		return
	}

	i := 0
	for res.next != nil {
		if (i == 0 && res.data != 1) || (i == 1 && res.data != 3) ||
			(i == 2 && res.data != 0) || (i == 3 && res.data != 8) {
			t.Errorf("AddList(%v %v)", n11, n21)
			return
		}
		i++
		res = res.next
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
