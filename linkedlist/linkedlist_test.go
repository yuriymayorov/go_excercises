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
		printLinkedList(&n1)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	n5 := LinkedListNode {nil, 5}
	n4 := LinkedListNode {&n5, 2}
	n3 := LinkedListNode {&n4, 3}
	n2 := LinkedListNode {&n3, 2}
	n1 := LinkedListNode {&n2, 1}

	DeleteDuplicates(&n1)
	
	if n3.next != nil && n3.next.data != 5 {
		t.Errorf("DeleteDuplicates(%v)", n1)
		printLinkedList(&n1)
	}
}

func TestNthToLast(t *testing.T) {
	n5 := LinkedListNode {nil, 5}
	n4 := LinkedListNode {&n5, 2}
	n3 := LinkedListNode {&n4, 3}
	n2 := LinkedListNode {&n3, 2}
	n1 := LinkedListNode {&n2, 1}

	in := 3
	res := NthToLast(&n1, in)	

	if res != &n3 {
		t.Errorf("NthToLast(%p, %d) = %p, want %p", &n1, in, &res, &n3)
		printLinkedList(&n1)
	}
}

func TestDeleteNode(t *testing.T) {
	n5 := LinkedListNode {nil, 5}
	n4 := LinkedListNode {&n5, 2}
	n3 := LinkedListNode {&n4, 3}
	n2 := LinkedListNode {&n3, 2}
	n1 := LinkedListNode {&n2, 1}

	//in := n3;
	printLinkedList(&n1)

	res := DeleteNode(&n3)
	fmt.Printf("\n\n\n\n")

	//fmt.Printf("%t\n\n\n\n", res)
	printLinkedList(&n1)

	if n2.next != &n4 || !res {
		t.Errorf("DeleteNode(%p) = %t, want %t and n2.next = %p want %p", &n3, res, true, &n2.next, &n4)
		//printLinkedList(&n1)
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