package LinkedList

type gen interface {
	~int | ~string | ~float64
}

type node[T gen] struct {
	value T
	next  *node[T]
}

type Link[T gen] struct {
	head *node[T]
	tail *node[T]

	// Returns the length of the list
	count int
}

// Puts a new value at the Tail end of the list
func (l *Link[T]) Add(val T) {
	myNode := node[T]{0, l.head}
	l.count++
	l.head = &myNode
}

// Inserts a new value at a given index, pushing the existing value at that index to the next index spot, and so on. Insert may ONLY target indices that are currently in use. In other words, if you have 5 elements in your list, you may insert at any index between 0 and 4 inclusive. Index 5 would be considered out of bounds as it is not currently in use during the insertion process. Any index less than zero or equal to or greater than Count should throw an index out of bounds exception.
func (l *Link[T]) Insert(val T, idx int) {
	myNode := node[T]{val}
	curNode := l.head

	// [1,3,4]
	// val = 2
	// idx = 1
	for i := 0; i < idx-1; i++ {
		curNode = curNode.next
	}
	curNode.next = &myNode
}

// Returns the value at the given index. Any index less than zero or equal to or greater than Count should throw an index out of bounds exception.
func (l *Link[T]) Get(idx int) T {
	curNode := l.head

	if idx <= 0 || idx == l.count || idx >= l.count {
		return
	}

	for i := 0; i < idx; i++ {
		curNode = curNode.next
	}

	return curNode.value
}

// Removes and returns the first value in the list
func (l *Link[T]) Remove() T {

	myNode := node[T]{0, l.head}
	l.count--
	l.head = nil
	l.head = &myNode

	return l.head.value
}

// Removes and returns the last value in the list
func (l *Link[T]) RemoveLast() T {

	myNode := node[T]{0, l.tail}
	l.count--
	l.tail = nil
	l.tail & myNode

	return l.tail.value
}

// Removes and returns the value at a given index. Any index less than zero or equal to or greater than Count should throw an index out of bounds exception.
func (l *Link[T]) RemoveAt(idx int) T {
	var t T
	return t
}

// Removes all values in the list.
func (l *Link[T]) Clear() {
	l.head = nil
}

// Searches for a value in the list and returns the first index of that value when found. If the key is not found in the list, the method returns -1.
func (l *Link[T]) Search(val T) int {

	return 1
}

// An override method that creates and returns a string representation of all the values in the list. The string must be in the format of “v0, v1, v2, .., vn-1” where n-1 is the last index in the list. An empty list should return an empty string (but not null). While every value in the string is separated by a comma and space, the string must NOT have any unnecessary commas or spaces at the beginning or end.
func (l *Link[T]) ToString() string {

	return ""
}

// Create a single linked list
func singleLinkList[T gen]() Link[T] {
	return Link[T]{nil, nil, 0}
}

// Create a double linked list
func doubleLinkList() {

}
