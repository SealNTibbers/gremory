package list

import (
	. "github.com/SealNTibbers/gremory/oset"
	. "github.com/SealNTibbers/gremory/utils"
)

type DataType = CollectionObject

type ListNode struct {
	Data DataType
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

func (l *List) PushFront(data interface{}) {
	if l.valueGenerator == nil {
		return
	}
	l.PushFrontValueHolder(l.valueGenerator(data))
}

func (l *List) PushFrontValueHolder(data CollectionObject) {
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
	if l.valueGenerator == nil {
		return
	}
	l.DeleteValueHolder(l.valueGenerator(data))
}

func (l *List) DeleteValueHolder(data CollectionObject) {
	if l.head == nil {
		return
	}
	currentElement := l.head
	var removedNode *ListNode
	for currentElement.next != nil {
		if currentElement.GetValue() == data.GetValue() {
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
	if l.valueGenerator == nil {
		return
	}
	l.PushBackValueHolder(l.valueGenerator(data))
}

func (l *List) PushBackValueHolder(data CollectionObject) {
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
	if l.head == nil {
		return nil
	}
	return l.head.GetValue()
}

func (l *List) Back() interface{} {
	if l.tail == nil {
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
	return currentNode.Data.GetValue()
}

func (l *List) InsertAt(data CollectionObject, index uint64) {
	if l.valueGenerator == nil {
		return
	}
	l.InsertAtValueHolder(l.valueGenerator(data), index)
}

func (l *List) InsertAtValueHolder(data CollectionObject, index uint64) {
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

func (l *List) ReverseDo(action func(each *ListNode)) {
	if l.head == nil {
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
	result := NewSmartList(l.valueGenerator)
	doAction := func(e *ListNode) {
		if selectAction(e) {
			result.PushBackValueHolder(e.Data)
		}
	}
	l.Do(doAction)
	return result
}

func (l *List) Collect(collectAction func(each *ListNode) CollectionObject) *List {
	result := NewSmartList(l.valueGenerator)
	doAction := func(e *ListNode) {
		result.PushBackValueHolder(collectAction(e))
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

func (l *List) AsOSet() *OSet {
	set := NewSmartOSet(l.valueGenerator)
	l.Do(func(each *ListNode) {
		set.AddValue(each.GetValue().(int))
	})
	return set
}
