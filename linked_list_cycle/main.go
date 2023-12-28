package main

func detectCycle(head *EduLinkedListNode) bool {
	return false
}

func main() {
	input := [][]int{
		{2, 4, 6, 8, 10},
		{1, 3, 5, 7, 9},
		{1, 2, 3, 4, 5},
		{0, 2, 3, 5, 6},
		{3, 6, 9, 10, 11},
	}

	ll := &EduLinkedList{}
	for _, i := range input {
		ll.CreateLinkedList(i)
		ll.DisplayLinkedList()
	}
}
