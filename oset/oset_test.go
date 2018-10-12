package oset

import (
	"fmt"
	"github.com/SealNTibbers/gremory/testutils"
	"github.com/SealNTibbers/gremory/utils"
	"testing"
)

func TestDo(t *testing.T) {
	set := NewOSet()
	set.AddValue(3)
	set.AddValue(2)
	set.AddValue(3)
	set.AddValue(15)
	set.AddValue(20)
	set.AddValue(1)
	set.AddValue(3)
	set.Do(func(each utils.TreeNodeInterface) {
		fmt.Println(each.GetValue())
	})
}

func TestSelect(t *testing.T) {
	set := NewOSet()
	set.AddValue(7)
	set.AddValue(10)
	set.AddValue(12)
	set.AddValue(24)
	selectList := set.Select(func(each utils.TreeNodeInterface) bool {
		if each.GetValue().(int) > 8 {
			return true
		}
		return false
	})
	selectList.Do(func(each utils.TreeNodeInterface) {
		fmt.Println(each.GetValue())
	})
}

func TestCollect(t *testing.T) {
	set := NewOSet()
	set.AddValue(1)
	set.AddValue(2)
	set.AddValue(3)
	collectList := set.Collect(func(each utils.TreeNodeInterface) interface{} {
		each.GetData().SetValue(each.GetValue().(int) * 10)
		return each.GetValue()
	})
	collectList.Do(func(each utils.TreeNodeInterface) {
		fmt.Println(each.GetValue())
	})
}

func TestInclides(t *testing.T) {
	set := NewOSet()
	set.AddValue(1)
	set.AddValue(2)
	set.AddValue(3)
	testutils.ASSERT_EQ(t, set.Includes(2), true)
	testutils.ASSERT_EQ(t, set.Includes(22), false)
}
