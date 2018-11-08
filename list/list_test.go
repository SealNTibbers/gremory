package list

import (
	"fmt"
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
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushFront(23)
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushFront(33)
	testutils.ASSERT_EQ(t, list.head.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPushBack(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(23)
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushBack(33)
	testutils.ASSERT_EQ(t, list.head.next.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPopBack(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack(43)
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.Back(), 43)
	back := list.PopBack()
	testutils.ASSERT_EQ(t, back, 43)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	testutils.ASSERT_EQ(t, list.Back(), 33)
}

func TestPopFront(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack(43)
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.Front(), 23)
	front := list.PopFront()
	testutils.ASSERT_EQ(t, front, 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	testutils.ASSERT_EQ(t, list.Front(), 33)
}

func TestGet(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(23)
	list.PushBack(33)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
}

func TestInsertAt(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(3)
	list.InsertAt(1, 0)
	testutils.ASSERT_EQ(t, list.At(0), 1)
	testutils.ASSERT_EQ(t, list.At(1), 3)
	list.InsertAt(2, 1)
	testutils.ASSERT_EQ(t, list.At(0), 1)
	testutils.ASSERT_EQ(t, list.At(1), 2)
	testutils.ASSERT_EQ(t, list.At(2), 3)
	list.InsertAt(4, 4)
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.At(0), 1)
	testutils.ASSERT_EQ(t, list.At(1), 2)
	testutils.ASSERT_EQ(t, list.At(2), 3)
}

func TestDelete(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack('a')
	list.PushBack(11)
	list.Delete(11)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
	testutils.ASSERT_EQ(t, list.At(2), 'a')
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	list.Delete(23)
	testutils.ASSERT_EQ(t, list.At(0), 33)
	testutils.ASSERT_EQ(t, list.At(1), 'a')
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	list.DeleteAt(0)
	testutils.ASSERT_EQ(t, list.At(0), 'a')
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushFront(33)
	list.DeleteAt(1)
	testutils.ASSERT_EQ(t, list.At(0), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.DeleteAt(2)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	testutils.ASSERT_EQ(t, list.At(0), 33)
	list.DeleteAt(0)
	testutils.ASSERT_EQ(t, list.Size(), uint64(0))
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack('a')
	list.PushBack(11)
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
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushBack(2)
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushBack(3)
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushFront(4)
	testutils.ASSERT_EQ(t, list.Front(), 4)
}

func TestEnd(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	testutils.ASSERT_EQ(t, list.Back(), 1)
	list.PushBack(2)
	testutils.ASSERT_EQ(t, list.Back(), 2)
	list.PushBack(3)
	testutils.ASSERT_EQ(t, list.Back(), 3)
	list.PushFront(4)
	testutils.ASSERT_EQ(t, list.Back(), 3)
}

func TestInclides(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	testutils.ASSERT_EQ(t, list.Includes(2), true)
	testutils.ASSERT_EQ(t, list.Includes(22), false)
}

func TestSort(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(5)
	list.PushBack(20)
	list.PushBack(4)
	list.PushBack(3)
	list.PushBack(30)
	list.Sort()
	counter := 0
	expectedValues := []int{3, 4, 5, 20, 30}
	list.Do(func(each *ListNode) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestAsSortedList(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(5)
	list.PushBack(20)
	list.PushBack(4)
	list.PushBack(3)
	list.PushBack(30)
	counter := 0
	expectedValues := []int{5, 20, 4, 3, 30}
	list.Do(func(each *ListNode) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})

	sortedList := list.AsSortedList()
	counter = 0
	expectedValues = []int{3, 4, 5, 20, 30}
	sortedList.Do(func(each *ListNode) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestWorkingWithTestType(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &testutils.TestTypeHolder{value.(testutils.TestType)}
	}
	list := NewSmartList(valueGen)
	list.PushBack(testutils.TestType{1, "john"})
	list.PushBack(testutils.TestType{2, "garry"})
	counter := 0
	expectedValues := []testutils.TestType{{1, "john"}, {2, "garry"}}
	list.Do(func(each *ListNode) {
		testutils.ASSERT_EQ(t, each.GetValue().(testutils.TestType).Id, expectedValues[counter].Id)
		testutils.ASSERT_EQ(t, each.GetValue().(testutils.TestType).Name, expectedValues[counter].Name)
		counter = counter + 1
	})
	list.Delete(testutils.TestType{2, "garry"})
	list.Do(func(each *ListNode) {
		testutils.ASSERT_EQ(t, each.GetValue().(testutils.TestType).Id, expectedValues[0].Id)
		testutils.ASSERT_EQ(t, each.GetValue().(testutils.TestType).Name, expectedValues[0].Name)
	})
}

func TestSample(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack(11)

	list.PushFront(0)

	front := list.PopFront()
	back := list.PopBack()
	fmt.Println(front)
	fmt.Println(back)

	list.InsertAt(0, 0)
	list.Do(func(each *ListNode) {
		fmt.Println(each.GetValue().(int))
	})
}
