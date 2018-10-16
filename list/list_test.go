package list

import (
	"github.com/SealNTibbers/gremory/testutils"
	"github.com/SealNTibbers/gremory/utils"
	"testing"
)

func TestCreateNode(t *testing.T) {
	newIntNode := CreateNode(&utils.ValueHolder{42})
	testutils.ASSERT_EQ(t, newIntNode.GetValue(), 42)

	newStrNode := CreateNode(&utils.ValueHolder{"test"})
	testutils.ASSERT_EQ(t, newStrNode.GetValue(), "test")

	newBoolNode := CreateNode(&utils.ValueHolder{true})
	testutils.ASSERT_EQ(t, newBoolNode.GetValue(), true)
}

func TestPushFront(t *testing.T) {
	list := new(List)
	list.PushFront(&utils.ValueHolder{23})
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushFront(&utils.ValueHolder{33})
	testutils.ASSERT_EQ(t, list.head.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPushBack(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{23})
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushBack(&utils.ValueHolder{33})
	testutils.ASSERT_EQ(t, list.head.next.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPopBack(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{23})
	list.PushBack(&utils.ValueHolder{33})
	list.PushBack(&utils.ValueHolder{43})
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.Back(), 43)
	back := list.PopBack()
	testutils.ASSERT_EQ(t, back, 43)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	testutils.ASSERT_EQ(t, list.Back(), 33)
}

func TestPopFront(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{23})
	list.PushBack(&utils.ValueHolder{33})
	list.PushBack(&utils.ValueHolder{43})
	testutils.ASSERT_EQ(t, list.Size(), uint64(3))
	testutils.ASSERT_EQ(t, list.Front(), 23)
	front := list.PopFront()
	testutils.ASSERT_EQ(t, front, 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
	testutils.ASSERT_EQ(t, list.Front(), 33)
}

func TestGet(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{23})
	list.PushBack(&utils.ValueHolder{33})
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
}

func TestInsertAt(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{23})
	list.PushBack(&utils.ValueHolder{33})
	list.PushBack(&utils.ValueHolder{34})
	list.InsertAt(&utils.ValueHolder{11}, 2)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
	testutils.ASSERT_EQ(t, list.At(2), 11)
	testutils.ASSERT_EQ(t, list.At(3), 34)
}

func TestDelete(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{23})
	list.PushBack(&utils.ValueHolder{33})
	list.PushBack(&utils.ValueHolder{34})
	list.PushBack(&utils.ValueHolder{11})
	list.Delete(&utils.ValueHolder{34})
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

func TestSelect(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{1})
	list.PushBack(&utils.ValueHolder{2})
	list.PushBack(&utils.ValueHolder{3})
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
	list := new(List)
	list.PushBack(&utils.ValueHolder{1})
	list.PushBack(&utils.ValueHolder{2})
	list.PushBack(&utils.ValueHolder{3})
	collectList := list.Collect(func(each *ListNode) utils.CollectionObject {
		each.Data.SetValue(each.GetValue().(int) * 10)
		return each.Data
	})
	testutils.ASSERT_EQ(t, collectList.At(0), 10)
	testutils.ASSERT_EQ(t, collectList.At(1), 20)
	testutils.ASSERT_EQ(t, collectList.At(2), 30)
}

func TestBegin(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{1})
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushBack(&utils.ValueHolder{2})
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushBack(&utils.ValueHolder{3})
	testutils.ASSERT_EQ(t, list.Front(), 1)
	list.PushFront(&utils.ValueHolder{4})
	testutils.ASSERT_EQ(t, list.Front(), 4)
}

func TestEnd(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{1})
	testutils.ASSERT_EQ(t, list.Back(), 1)
	list.PushBack(&utils.ValueHolder{2})
	testutils.ASSERT_EQ(t, list.Back(), 2)
	list.PushBack(&utils.ValueHolder{3})
	testutils.ASSERT_EQ(t, list.Back(), 3)
	list.PushFront(&utils.ValueHolder{4})
	testutils.ASSERT_EQ(t, list.Back(), 3)
}

func TestInclides(t *testing.T) {
	list := new(List)
	list.PushBack(&utils.ValueHolder{1})
	list.PushBack(&utils.ValueHolder{2})
	list.PushBack(&utils.ValueHolder{3})
	testutils.ASSERT_EQ(t, list.Includes(2), true)
	testutils.ASSERT_EQ(t, list.Includes(22), false)
}
