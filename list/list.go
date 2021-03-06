package list

import (
	. "github.com/SealNTibbers/gremory/utils"
)

type ListNode struct {
	Data CollectionObject
	next *ListNode
	prev *ListNode
	head *ListNode
	tail *ListNode
}

type List struct {
	head           *ListNode
	tail           *ListNode
	valueGenerator func(interface{}) CollectionObject
}

func NewSmartList(valueGenerator func(interface{}) CollectionObject) *List {
	list := new(List)
	list.valueGenerator = valueGenerator
	return list
}

func CreateNode(data CollectionObject) *ListNode {
	node := new(ListNode)
	node.Data = data
	node.next = nil
	node.prev = nil
	return node
}

func (node *ListNode) GetValue() interface{} {
	if node == nil {
		return nil
	}
	return node.Data.GetValue()
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) PushFront(data interface{}) {
	if l.valueGenerator == nil {
		return
	}
	l.pushFrontValueHolder(l.valueGenerator(data))
}

func (l *List) pushFrontValueHolder(data CollectionObject) {
	newNode := CreateNode(data)
	if l.IsEmpty() {
		l.head = newNode
		return
	}
	l.head.prev = newNode
	newNode.next = l.head
	l.head = newNode
}

func (l *List) Delete(data interface{}) {
	if l.valueGenerator == nil {
		return
	}
	l.deleteCollectionObject(l.valueGenerator(data))
}

func (l *List) deleteCollectionObject(data CollectionObject) {
	if l.IsEmpty() {
		return
	}
	currentElement := l.head
	var removedNode *ListNode
	for removedNode == nil && currentElement != nil {
		if data.Equal(currentElement.Data) {
			removedNode = currentElement
		}
		currentElement = currentElement.next
	}
	if removedNode.Data.Equal(l.head.Data) {
		l.head = removedNode.next
	}

	if removedNode.Data.Equal(l.tail.Data) {
		l.tail = removedNode.prev
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
	listSize := l.Size()
	if l.IsEmpty() || index > listSize {
		return
	}
	currentElement := l.head
	var removedNode *ListNode
	if index > 0 {
		var counter uint64 = 1
		for currentElement.next != nil {
			if counter == index {
				removedNode = currentElement
			}
			currentElement = currentElement.next
			counter = counter + 1
		}
	} else {
		removedNode = currentElement
		l.head = currentElement.next
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
	if l.IsEmpty() {
		return
	}
	l.head = nil
}

func (l *List) PushBack(data interface{}) {
	if l.valueGenerator == nil {
		return
	}
	l.pushBackCollectionObject(l.valueGenerator(data))
}

func (l *List) pushBackCollectionObject(data CollectionObject) {
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

func (l *List) PopBack() interface{} {
	back := l.tail
	l.tail = back.prev
	back.prev.next = nil
	return back.Data.GetValue()
}

func (l *List) Front() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.head.GetValue()
}

func (l *List) Back() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.tail.GetValue()
}

func (l *List) PopFront() interface{} {
	front := l.head
	l.head = front.next
	front.next.prev = nil
	return front.Data.GetValue()
}

func (l *List) Size() uint64 {
	if l.IsEmpty() {
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
	node := l.atNode(index)
	if node == nil {
		return nil
	}
	return node.Data.GetValue()
}

func (l *List) atNode(index uint64) *ListNode {
	if l.IsEmpty() || index > l.Size() {
		return nil
	}
	currentNode := l.head
	var counter uint64 = 0
	for counter != index {
		currentNode = currentNode.next
		counter = counter + 1
	}
	return currentNode
}

func (l *List) InsertAt(data interface{}, index uint64) {
	if l.valueGenerator == nil {
		return
	}

	l.insertAtCollectionObject(l.valueGenerator(data), index)
}

func (l *List) insertAtCollectionObject(data CollectionObject, index uint64) {
	listSize := l.Size()
	if l.IsEmpty() || index > listSize {
		return
	}
	currentNode := l.head
	newNode := CreateNode(data)

	if index > 0 {
		var counter uint64 = 1
		for counter != index {
			currentNode = currentNode.next
			counter = counter + 1
		}
		newNode.next = currentNode.next
		currentNode.next = newNode
		newNode.prev = currentNode
		if listSize == index {
			l.tail = currentNode
		}
	} else {
		currentNode.prev = newNode
		newNode.next = currentNode
		l.head = newNode
	}

	if newNode.next != nil {
		newNode.next.prev = newNode
	}
}

func (l *List) ReverseDo(action func(each *ListNode)) {
	if l.IsEmpty() {
		return
	}
	currentElement := l.tail
	action(currentElement)
	for currentElement.prev != nil {
		currentElement = currentElement.prev
		action(currentElement)
	}
}

func (l *List) Do(action func(each *ListNode)) {
	if l.IsEmpty() {
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
	result := NewSmartList(l.valueGenerator)
	doAction := func(e *ListNode) {
		if selectAction(e) {
			result.pushBackCollectionObject(e.Data)
		}
	}
	l.Do(doAction)
	return result
}

func (l *List) Collect(collectAction func(each *ListNode) CollectionObject) *List {
	result := NewSmartList(l.valueGenerator)
	doAction := func(e *ListNode) {
		result.pushBackCollectionObject(collectAction(e))
	}
	l.Do(doAction)
	return result
}

func swapData(lv *ListNode, rv *ListNode) {
	var tmpData = lv.Data
	lv.Data = rv.Data
	rv.Data = tmpData
}

func partition(low *ListNode, high *ListNode) *ListNode {
	var i *ListNode = nil
	var j = low
	var pivot = high

	for j != high {
		if pivot.Data.Greater(j.Data) {
			if i == nil {
				i = low
			} else {
				i = i.next
			}
			swapData(i, j)
		}
		j = j.next
	}
	if i == nil {
		i = low
	} else {
		i = i.next
	}
	swapData(i, j)

	return i
}

func quickSort(low *ListNode, high *ListNode) {
	if high != nil && low != high && low != high.next {
		var currentPartition = partition(low, high)
		quickSort(low, currentPartition.prev)
		quickSort(currentPartition.next, high)
	}
}

func (l *List) AsSortedList() *List {
	copyList := l.Select(func(each *ListNode) bool {
		return true
	})
	copyList.Sort()
	return copyList
}

func (l *List) Sort() {
	quickSort(l.head, l.tail)
}

func (l *List) Includes(data interface{}) bool {
	if l.IsEmpty() {
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

func (l *List) SwapIndex(firstIndex uint64, secondIndex uint64) {
	if firstIndex == secondIndex {
		return
	}
	if firstIndex < 0 || firstIndex > l.Size()-1 || secondIndex > l.Size()-1 {
		return
	}
	firstElement := l.atNode(firstIndex)
	secondElement := l.atNode(secondIndex)
	swapData(firstElement, secondElement)
}
