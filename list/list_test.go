package list

import (
	"github.com/SealNTibbers/gremory/testutils"
	. "github.com/SealNTibbers/gremory/utils"
	"testing"
)

func TestCreateNode(t *testing.T) {
	newIntNode := CreateNode(&ValueHolder{42})
	testutils.ASSERT_EQ(t, newIntNode.GetValue(), 42)

	newStrNode := CreateNode(&ValueHolder{"test"})
	testutils.ASSERT_EQ(t, newStrNode.GetValue(), "test")

	newBoolNode := CreateNode(&ValueHolder{true})
	testutils.ASSERT_EQ(t, newBoolNode.GetValue(), true)
}

func TestPushFront(t *testing.T) {
	list := new(List)
	list.PushFrontValueHolder(&ValueHolder{23})
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushFrontValueHolder(&ValueHolder{33})
	testutils.ASSERT_EQ(t, list.head.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPushBack(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{23})
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushBackValueHolder(&ValueHolder{33})
	testutils.ASSERT_EQ(t, list.head.next.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPopBack(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{23})
	list.PushBackValueHolder(&ValueHolder{33})
	list.PushBackValueHolder(&ValueHolder{43})
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.Back(), 43)
	back := list.PopBack()
	testutils.ASSERT_EQ(t, back, 43)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	testutils.ASSERT_EQ(t, list.Back(), 33)
}

func TestPopFront(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{23})
	list.PushBackValueHolder(&ValueHolder{33})
	list.PushBackValueHolder(&ValueHolder{43})
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.Front(), 23)
	front := list.PopFront()
	testutils.ASSERT_EQ(t, front, 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	testutils.ASSERT_EQ(t, list.Front(), 33)
}

func TestGet(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{23})
	list.PushBackValueHolder(&ValueHolder{33})
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
}

func TestInsertAt(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{23})
	list.PushBackValueHolder(&ValueHolder{33})
	list.PushBackValueHolder(&ValueHolder{34})
	list.InsertAtValueHolder(&ValueHolder{11}, 2)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
	testutils.ASSERT_EQ(t, list.At(2), 11)
	testutils.ASSERT_EQ(t, list.At(3), 34)
}

func TestDelete(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{23})
	list.PushBackValueHolder(&ValueHolder{33})
	list.PushBackValueHolder(&ValueHolder{34})
	list.PushBackValueHolder(&ValueHolder{11})
	list.DeleteValueHolder(&ValueHolder{34})
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
	testutils.ASSERT_EQ(t, list.At(2), 11)
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	list.DeleteAt(1)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 11)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	list.DeleteAll()
	testutils.ASSERT_EQ(t, list.Size(), uint64(0))
}

func TestDo(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	counter := 0
	expectedValues := []int{1, 2, 3}
	list.Do(func(each *ListNode) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestSelect(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	selectList := list.Select(func(each *ListNode) bool {
		if each.GetValue().(int) > 1 {
			return true
		}
		return false
	})
	testutils.ASSERT_EQ(t, selectList.At(0), 2)
	testutils.ASSERT_EQ(t, selectList.At(1), 3)
}

func TestCollect(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	collectList := list.Collect(func(each *ListNode) CollectionObject {
		return &ValueHolder{each.GetValue().(int) * 10}
	})
	testutils.ASSERT_EQ(t, collectList.At(0), 10)
	testutils.ASSERT_EQ(t, collectList.At(1), 20)
	testutils.ASSERT_EQ(t, collectList.At(2), 30)
}

func TestBegin(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{1})
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushBackValueHolder(&ValueHolder{2})
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushBackValueHolder(&ValueHolder{3})
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushFrontValueHolder(&ValueHolder{4})
	testutils.ASSERT_EQ(t, list.Front(), 4)
}

func TestEnd(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{1})
	testutils.ASSERT_EQ(t, list.Back(), 1)
	list.PushBackValueHolder(&ValueHolder{2})
	testutils.ASSERT_EQ(t, list.Back(), 2)
	list.PushBackValueHolder(&ValueHolder{3})
	testutils.ASSERT_EQ(t, list.Back(), 3)
	list.PushFrontValueHolder(&ValueHolder{4})
	testutils.ASSERT_EQ(t, list.Back(), 3)
}

func TestInclides(t *testing.T) {
	list := new(List)
	list.PushBackValueHolder(&ValueHolder{1})
	list.PushBackValueHolder(&ValueHolder{2})
	list.PushBackValueHolder(&ValueHolder{3})
	testutils.ASSERT_EQ(t, list.Includes(2), true)
	testutils.ASSERT_EQ(t, list.Includes(22), false)
}
