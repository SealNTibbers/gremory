package oset

import (
	"github.com/SealNTibbers/gremory/testutils"
	. "github.com/SealNTibbers/gremory/utils"
	"testing"
)

func TestDo(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{3})
	set.AddValue(&ValueHolder{2})
	set.AddValue(&ValueHolder{3})
	set.AddValue(&ValueHolder{15})
	set.AddValue(&ValueHolder{20})
	set.AddValue(&ValueHolder{1})
	set.AddValue(&ValueHolder{3})
	counter := 0
	expectedValues := []int{1, 2, 3, 15, 20}
	set.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestReverseDo(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{3})
	set.AddValue(&ValueHolder{2})
	set.AddValue(&ValueHolder{3})
	set.AddValue(&ValueHolder{15})
	set.AddValue(&ValueHolder{20})
	set.AddValue(&ValueHolder{1})
	set.AddValue(&ValueHolder{3})
	counter := 4
	expectedValues := []int{20, 15, 3, 2, 1}
	set.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter - 1
	})
}

func TestSelect(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{7})
	set.AddValue(&ValueHolder{10})
	set.AddValue(&ValueHolder{12})
	set.AddValue(&ValueHolder{24})
	selectedSet := set.Select(func(each TreeNodeInterface) bool {
		if each.GetValue().(int) > 8 {
			return true
		}
		return false
	})
	counter := 0
	expectedValues := []int{10, 12, 24}
	selectedSet.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestCollect(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{1})
	set.AddValue(&ValueHolder{2})
	set.AddValue(&ValueHolder{3})
	collectSet := set.Collect(func(each TreeNodeInterface) CollectionObject {
		each.GetData().SetValue(each.GetValue().(int) * 10)
		return each.GetData()
	})
	counter := 0
	expectedValues := []int{10, 20, 30}
	collectSet.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestInclides(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{1})
	set.AddValue(&ValueHolder{2})
	set.AddValue(&ValueHolder{3})
	testutils.ASSERT_EQ(t, set.Includes(2), true)
	testutils.ASSERT_EQ(t, set.Includes(22), false)
}
