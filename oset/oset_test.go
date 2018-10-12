package oset

import (
	"fmt"
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
	set.Do(func(each TreeNodeInterface) {
		fmt.Println(each.GetData())
	})
}

func TestSelect(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{7})
	set.AddValue(&ValueHolder{10})
	set.AddValue(&ValueHolder{12})
	set.AddValue(&ValueHolder{24})
	selectList := set.Select(func(each TreeNodeInterface) bool {
		if each.GetValue().(int) > 8 {
			return true
		}
		return false
	})
	selectList.Do(func(each TreeNodeInterface) {
		fmt.Println(each.GetValue())
	})
}

func TestCollect(t *testing.T) {
	set := NewOSet()
	set.AddValue(&ValueHolder{1})
	set.AddValue(&ValueHolder{2})
	set.AddValue(&ValueHolder{3})
	collectList := set.Collect(func(each TreeNodeInterface) CollectionObject {
		each.GetData().SetValue(each.GetValue().(int) * 10)
		return each.GetData()
	})
	collectList.Do(func(each TreeNodeInterface) {
		fmt.Println(each.GetValue())
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
