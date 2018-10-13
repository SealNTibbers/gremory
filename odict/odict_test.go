package odict

import (
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
	expectedValues := []int32{97, 98, 99}
	for i := 0; i < 3; i++ {
		testutils.ASSERT_EQ(t, dict.At(i+1).(int32), expectedValues[i])
	}
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
	counter := 0
	expectedKeys := []int{1, 2, 3, 4}
	expectedValues := []int32{97, 98, 99, 100}
	dict.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetKeyValue().(int), expectedKeys[counter])
		testutils.ASSERT_EQ(t, each.GetValue(), expectedValues[counter])
		counter = counter + 1
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
	counter := 0
	expectedKeys := []int{2, 3, 4}
	expectedValues := []int32{98, 99, 100}
	selected.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetKeyValue().(int), expectedKeys[counter])
		testutils.ASSERT_EQ(t, each.GetValue(), expectedValues[counter])
		counter = counter + 1
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
	counter := 0
	expectedKeys := []int{1, 2, 3, 4}
	expectedValues := []int32{102, 103, 104, 105}
	collected.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetKeyValue().(int), expectedKeys[counter])
		testutils.ASSERT_EQ(t, each.GetValue(), expectedValues[counter])
		counter = counter + 1
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
