package linkedlist

import (
	"testing"
	"fmt"
)

func TestAppendToTail(t *testing.T) {
	n3 := LinkedListNode {nil, 3}
	n2 := LinkedListNode {&n3, 2}
	n1 := LinkedListNode {&n2, 1}
	newVal := 0;

	n1.AppendToTail(newVal)

	if n3.next == nil || n3.next.data != newVal {
		t.Errorf("AppendToTail(%d)", newVal)
		printLinkedList(n1)
	}
}

func TestDeleteDuplicates(t * testing.T) {
	n5 := LinkedListNode {nil, 5}
	n4 := LinkedListNode {&n5, 2}
	n3 := LinkedListNode {&n4, 3}
	n2 := LinkedListNode {&n3, 2}
	n1 := LinkedListNode {&n2, 1}

	DeleteDuplicates(&n1)
	
	if (n3.next != nil && n3.next.data != 5) {
		t.Errorf("DeleteDuplicates(%v)", n1)
		printLinkedList(n1)
	}
}

func TestNthToLast(t *testing.T) {
	n5 := LinkedListNode {nil, 5}
	n4 := LinkedListNode {&n5, 22}
	n3 := LinkedListNode {&n4, 345}
	n2 := LinkedListNode {&n3, 2}
	n1 := LinkedListNode {&n2, 1}
	
	in := 3
	k, res := NthToLast(&n1, in)

	if (res.data != 345) {
		t.Errorf("NthToLast(%v, %d) = %d, %d, wanted = %d, %d", &n1, in, k, res.data, 3, 345)
		printLinkedList(n1)
	}
}

func printLinkedList(node LinkedListNode) {
	i := 1
	for node.next != nil {
		fmt.Printf("node%d: %d\n", i, node.data)
		i++
		node = *node.next	
	} 
	fmt.Printf("node%d: %d\n", i, node.data)
}