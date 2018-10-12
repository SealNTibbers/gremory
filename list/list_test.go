package list

import (
	"github.com/gremory/testutils"
	"testing"
)

func TestCreateNode(t *testing.T) {
	newIntNode := CreateNode(42)
	testutils.ASSERT_EQ(t, newIntNode.GetValue(), 42)

	newStrNode := CreateNode("test")
	testutils.ASSERT_EQ(t, newStrNode.GetValue(), "test")

	newBoolNode := CreateNode(true)
	testutils.ASSERT_EQ(t, newBoolNode.GetValue(), true)
}

func TestPushFront(t *testing.T) {
	list := new(List)
	list.PushFront(23)
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushFront(33)
	testutils.ASSERT_EQ(t, list.head.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestPushBack(t *testing.T) {
	list := new(List)
	list.PushBack(23)
	testutils.ASSERT_EQ(t, list.head.GetValue(), 23)
	testutils.ASSERT_EQ(t, list.Size(), uint64(1))
	list.PushBack(33)
	testutils.ASSERT_EQ(t, list.head.next.GetValue(), 33)
	testutils.ASSERT_EQ(t, list.Size(), uint64(2))
}

func TestGet(t *testing.T) {
	list := new(List)
	list.PushBack(23)
	list.PushBack(33)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
}

func TestInsertAt(t *testing.T) {
	list := new(List)
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack(34)
	list.InsertAt(11, 2)
	testutils.ASSERT_EQ(t, list.At(0), 23)
	testutils.ASSERT_EQ(t, list.At(1), 33)
	testutils.ASSERT_EQ(t, list.At(2), 11)
	testutils.ASSERT_EQ(t, list.At(3), 34)
}

func TestDelete(t *testing.T) {
	list := new(List)
	list.PushBack(23)
	list.PushBack(33)
	list.PushBack(34)
	list.PushBack(11)
	list.Delete(34)
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
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	selectList := list.Select(func(each *ListNode) bool {
		if each.Data.Data.(int) > 1 {
			return true
		}
		return false
	})
	testutils.ASSERT_EQ(t, selectList.At(0), 2)
	testutils.ASSERT_EQ(t, selectList.At(1), 3)
}

func TestCollect(t *testing.T) {
	list := new(List)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	collectList := list.Collect(func(each *ListNode) interface{} {
		each.Data.Data = each.Data.Data.(int) * 10
		return each.GetValue()
	})
	testutils.ASSERT_EQ(t, collectList.At(0), 10)
	testutils.ASSERT_EQ(t, collectList.At(1), 20)
	testutils.ASSERT_EQ(t, collectList.At(2), 30)
}

func TestBegin(t *testing.T) {
	list := new(List)
	list.PushBack(1)
	testutils.ASSERT_EQ(t, list.Begin().GetValue(), 1)
	list.PushBack(2)
	testutils.ASSERT_EQ(t, list.Begin().GetValue(), 1)
	list.PushBack(3)
	testutils.ASSERT_EQ(t, list.Begin().GetValue(), 1)
	list.PushFront(4)
	testutils.ASSERT_EQ(t, list.Begin().GetValue(), 4)
}

func TestEnd(t *testing.T) {
	list := new(List)
	list.PushBack(1)
	testutils.ASSERT_EQ(t, list.End().GetValue(), 1)
	list.PushBack(2)
	testutils.ASSERT_EQ(t, list.End().GetValue(), 2)
	list.PushBack(3)
	testutils.ASSERT_EQ(t, list.End().GetValue(), 3)
	list.PushFront(4)
	testutils.ASSERT_EQ(t, list.End().GetValue(), 3)
}

func TestInclides(t *testing.T) {
	list := new(List)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	testutils.ASSERT_EQ(t, list.Includes(2), true)
	testutils.ASSERT_EQ(t, list.Includes(22), false)
}
