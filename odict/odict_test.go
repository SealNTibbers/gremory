package odict

import (
	"fmt"
	"github.com/SealNTibbers/gremory/testutils"
	. "github.com/SealNTibbers/gremory/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	dict.AddPair(&ValueHolder{4}, &ValueHolder{'d'})
	testutils.ASSERT_EQ(t, dict.Size(), uint64(4))
}

func TestAt(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	testutils.ASSERT_EQ(t, dict.Size(), uint64(3))
	testutils.ASSERT_EQ(t, dict.At(1).(int32), int32(97))
	testutils.ASSERT_EQ(t, dict.At(2).(int32), int32(98))
	testutils.ASSERT_EQ(t, dict.At(3).(int32), int32(99))
	testutils.ASSERT_EQ(t, dict.At(4), nil)

}

func TestDelete(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	dict.AddPair(&ValueHolder{4}, &ValueHolder{'d'})
	testutils.ASSERT_EQ(t, dict.Size(), uint64(4))
	dict.DeleteKey(1)
	testutils.ASSERT_EQ(t, dict.Size(), uint64(3))
}

func TestDo(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	dict.AddPair(&ValueHolder{4}, &ValueHolder{'d'})
	dict.Do(func(each TreeNodeInterface) {
		fmt.Println(each.GetKey(), "->", each.GetValue())
	})
}

func TestSelect(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	dict.AddPair(&ValueHolder{4}, &ValueHolder{'d'})
	selected := dict.Select(func(each TreeNodeInterface) bool {
		if each.GetKeyValue().(int) > 1 {
			return true
		}
		return false
	})
	selected.Do(func(each TreeNodeInterface) {
		fmt.Println(each.GetKey(), "->", each.GetValue())
	})
}

func TestCollect(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	dict.AddPair(&ValueHolder{4}, &ValueHolder{'d'})
	collected := dict.Collect(func(each TreeNodeInterface) (CollectionObject, CollectionObject) {
		each.GetData().SetValue(each.GetValue().(int32) + 5)
		return each.GetKey(), each.GetData()
	})
	collected.Do(func(each TreeNodeInterface) {
		fmt.Println(each.GetKey(), "->", each.GetValue())
	})
}

func TestInclides(t *testing.T) {
	dict := NewODict()
	dict.AddPair(&ValueHolder{1}, &ValueHolder{'a'})
	dict.AddPair(&ValueHolder{2}, &ValueHolder{'b'})
	dict.AddPair(&ValueHolder{3}, &ValueHolder{'c'})
	dict.AddPair(&ValueHolder{4}, &ValueHolder{'d'})
	testutils.ASSERT_EQ(t, dict.Includes(2), true)
	testutils.ASSERT_EQ(t, dict.Includes(22), false)
}
