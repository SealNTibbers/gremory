package odict

import (
	"fmt"
	"github.com/SealNTibbers/gremory/testutils"
	"github.com/SealNTibbers/gremory/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	dict.AddPair(4, 'd')
	testutils.ASSERT_EQ(t, dict.Size(), uint64(4))
}

func TestAt(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	testutils.ASSERT_EQ(t, dict.Size(), uint64(3))
	testutils.ASSERT_EQ(t, dict.At(1).(int32), int32(97))
	testutils.ASSERT_EQ(t, dict.At(2).(int32), int32(98))
	testutils.ASSERT_EQ(t, dict.At(3).(int32), int32(99))
	testutils.ASSERT_EQ(t, dict.At(4), nil)

}

func TestDelete(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	dict.AddPair(4, 'd')
	testutils.ASSERT_EQ(t, dict.Size(), uint64(4))
	dict.DeleteKey(1)
	testutils.ASSERT_EQ(t, dict.Size(), uint64(3))
}

func TestDo(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	dict.AddPair(4, 'd')
	dict.Do(func(each utils.TreeNodeInterface) {
		fmt.Println(each.GetKey(), "->", each.GetValue())
	})
}

func TestSelect(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	dict.AddPair(4, 'd')
	selected := dict.Select(func(each utils.TreeNodeInterface) bool {
		if each.GetKey().(int) > 1 {
			return true
		}
		return false
	})
	selected.Do(func(each utils.TreeNodeInterface) {
		fmt.Println(each.GetKey(), "->", each.GetValue())
	})
}

func TestCollect(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	dict.AddPair(4, 'd')
	collected := dict.Collect(func(each utils.TreeNodeInterface) (interface{}, interface{}) {
		each.GetData().SetValue(each.GetValue().(int32) + 5)
		return each.GetKey(), each.GetValue()
	})
	collected.Do(func(each utils.TreeNodeInterface) {
		fmt.Println(each.GetKey(), "->", each.GetValue())
	})
}

func TestInclides(t *testing.T) {
	dict := NewODict()
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
	dict.AddPair(4, 'd')
	testutils.ASSERT_EQ(t, dict.Includes(2), true)
	testutils.ASSERT_EQ(t, dict.Includes(22), false)
}
