package list

import "github.com/SealNTibbers/gremory/utils"

type DataType = utils.ValueHolder

type ListNode struct {
	Data *DataType
	next *ListNode
	prev *ListNode
	head *ListNode
	tail *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func CreateNode(data interface{}) *ListNode {
	node := new(ListNode)
	dataHolder := new(utils.ValueHolder)
	dataHolder.Data = data
	node.Data = dataHolder
	node.next = nil
	node.prev = nil
	return node
}

func (node *ListNode) GetValue() interface{} {
	if node == nil {
		return nil
	}
	return node.Data.Data
}

func (l *List) Begin() *ListNode {
	return l.head
}

func (l *List) End() *ListNode {
	return l.tail
}

func (l *List) PushFront(data interface{}) {
	newNode := CreateNode(data)
	if l.head == nil {
		l.head = newNode
		return
	}
	l.head.prev = newNode
	newNode.next = l.head
	l.head = newNode
}

func (l *List) Delete(data interface{}) {
	if l.head == nil {
		return
	}
	currentElement := l.head
	var removedNode *ListNode
	for currentElement.next != nil {
		if currentElement.GetValue() == data {
			removedNode = currentElement
		}
		currentElement = currentElement.next
	}
	if removedNode == nil {
		return
	}
	if removedNode.next != nil {
		removedNode.next.prev = removedNode.prev
	}
	if removedNode.prev != nil {
		removedNode.prev.next = removedNode.next
	}
	removedNode.next = nil
	removedNode.prev = nil
}

func (l *List) DeleteAt(index uint64) {
	if l.head == nil {
		return
	}
	currentElement := l.head
	var removedNode *ListNode
	var counter uint64 = 0
	for currentElement.next != nil {
		if counter == index {
			removedNode = currentElement
		}
		currentElement = currentElement.next
		counter = counter + 1
	}
	if removedNode == nil {
		return
	}
	if removedNode.next != nil {
		removedNode.next.prev = removedNode.prev
	}
	if removedNode.prev != nil {
		removedNode.prev.next = removedNode.next
	}
	removedNode.next = nil
	removedNode.prev = nil
}

func (l *List) DeleteAll() {
	if l.head == nil {
		return
	}
	l.head = nil
}

func (l *List) PushBack(data interface{}) {
	temp := l.head
	newNode := CreateNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	newNode.prev = temp
	l.tail = newNode
}

func (l *List) Size() uint64 {
	if l.head == nil {
		return 0
	}
	currentElement := l.head
	var counter uint64 = 1
	for currentElement.next != nil {
		currentElement = currentElement.next
		counter = counter + 1
	}
	return counter
}

func (l *List) At(index uint64) interface{} {
	if l.head == nil || index > l.Size() {
		return nil
	}
	currentNode := l.head
	var counter uint64 = 0
	for counter != index {
		currentNode = currentNode.next
		counter = counter + 1
	}
	return currentNode.Data.Data
}

func (l *List) InsertAt(data interface{}, index uint64) {
	if l.head == nil || index > l.Size() {
		return
	}
	currentNode := l.head
	var counter uint64 = 1
	for counter != index {
		currentNode = currentNode.next
		counter = counter + 1
	}
	newNode := CreateNode(data)

	newNode.next = currentNode.next
	currentNode.next = newNode
	newNode.prev = currentNode

	if newNode.next != nil {
		newNode.next.prev = newNode
	}
}

func (l *List) Do(action func(each *ListNode)) {
	if l.head == nil {
		return
	}
	currentElement := l.head
	action(currentElement)
	for currentElement.next != nil {
		currentElement = currentElement.next
		action(currentElement)
	}
}

func (l *List) Select(selectAction func(each *ListNode) bool) *List {
	result := new(List)
	doAction := func(e *ListNode) {
		if selectAction(e) {
			result.PushBack(e.GetValue())
		}
	}
	l.Do(doAction)
	return result
}

func (l *List) Collect(collectAction func(each *ListNode) interface{}) *List {
	result := new(List)
	doAction := func(e *ListNode) {
		result.PushBack(collectAction(e))
	}
	l.Do(doAction)
	return result
}

func (l *List) Includes(data interface{}) bool {
	if l.head == nil {
		return false
	}
	selectAction := func(each *ListNode) bool {
		if each.Data.GetValue() == data {
			return true
		}
		return false
	}
	result := l.Select(selectAction)
	return result.Size() > 0
}
