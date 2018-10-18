package oset

import (
	"github.com/SealNTibbers/gremory/testutils"
	. "github.com/SealNTibbers/gremory/utils"
	"testing"
)

type TestType struct {
	id   int
	name string
}

type TestTypeHolder struct {
	data TestType
}

func (holder *TestTypeHolder) GetValue() interface{} {
	return holder.data
}

func (lv *TestTypeHolder) Less(rv CollectionObject) bool {
	rvValue, ok := rv.(*TestTypeHolder)
	if !ok {
		return false
	}
	return lv.data.id < rvValue.data.id
}

func (lv *TestTypeHolder) Greater(rv CollectionObject) bool {
	rvValue, ok := rv.(*TestTypeHolder)
	if !ok {
		return false
	}
	return lv.data.id > rvValue.data.id
}

func (lv *TestTypeHolder) Equal(rv CollectionObject) bool {
	rvValue, ok := rv.(*TestTypeHolder)
	if !ok {
		return false
	}
	return lv.data.id > rvValue.data.id
}

func TestDoForTestType(t *testing.T) {
	set := NewOSet()
	set.AddValueHolders(&TestTypeHolder{TestType{1, "john"}})
	set.AddValueHolders(&TestTypeHolder{TestType{2, "garry"}})
	set.AddValueHolders(&TestTypeHolder{TestType{3, "marry"}})
	counter := 0
	expectedValues := []TestType{{1, "john"}, {2, "garry"}, {3, "marry"}}
	set.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(TestType).id, expectedValues[counter].id)
		testutils.ASSERT_EQ(t, each.GetValue().(TestType).name, expectedValues[counter].name)
		counter = counter + 1
	})
}

func TestSelectForTestType(t *testing.T) {
	set := NewOSet()
	set.AddValueHolders(&TestTypeHolder{TestType{1, "john"}})
	set.AddValueHolders(&TestTypeHolder{TestType{2, "garry"}})
	set.AddValueHolders(&TestTypeHolder{TestType{3, "marry"}})
	selectedSet := set.Select(func(each TreeNodeInterface) bool {
		if each.GetValue().(TestType).id > 1 {
			return true
		}
		return false
	})
	counter := 0
	expectedValues := []TestType{{2, "garry"}, {3, "marry"}}
	selectedSet.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(TestType).id, expectedValues[counter].id)
		testutils.ASSERT_EQ(t, each.GetValue().(TestType).name, expectedValues[counter].name)
		counter = counter + 1
	})
}

func TestCollectForTestType(t *testing.T) {
	set := NewOSet()
	set.AddValueHolders(&TestTypeHolder{TestType{1, "john"}})
	set.AddValueHolders(&TestTypeHolder{TestType{2, "garry"}})
	set.AddValueHolders(&TestTypeHolder{TestType{3, "marry"}})

	collectedSet := set.Collect(func(each TreeNodeInterface) CollectionObject {
		return &TestTypeHolder{TestType{each.GetValue().(TestType).id * 10, each.GetValue().(TestType).name}}
	})
	counter := 0
	expectedValues := []TestType{{10, "john"}, {20, "garry"}, {30, "marry"}}
	collectedSet.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(TestType).id, expectedValues[counter].id)
		testutils.ASSERT_EQ(t, each.GetValue().(TestType).name, expectedValues[counter].name)
		counter = counter + 1
	})
}

func TestDo(t *testing.T) {
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	set := NewSmartOSet(valueGen)
	set.AddValue(3)
	set.AddValue(2)
	set.AddValue(3)
	set.AddValue(15)
	set.AddValue(20)
	set.AddValue(1)
	set.AddValue(3)
	counter := 0
	expectedValues := []int{1, 2, 3, 15, 20}
	set.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter + 1
	})
}

func TestReverseDo(t *testing.T) {
	set := NewOSet()
	set.AddValueHolders(&ValueHolder{3})
	set.AddValueHolders(&ValueHolder{2})
	set.AddValueHolders(&ValueHolder{3})
	set.AddValueHolders(&ValueHolder{15})
	set.AddValueHolders(&ValueHolder{20})
	set.AddValueHolders(&ValueHolder{1})
	set.AddValueHolders(&ValueHolder{3})
	counter := 4
	expectedValues := []int{20, 15, 3, 2, 1}
	set.Do(func(each TreeNodeInterface) {
		testutils.ASSERT_EQ(t, each.GetValue().(int), expectedValues[counter])
		counter = counter - 1
	})
}

func TestSelect(t *testing.T) {
	set := NewOSet()
	set.AddValueHolders(&ValueHolder{7})
	set.AddValueHolders(&ValueHolder{10})
	set.AddValueHolders(&ValueHolder{12})
	set.AddValueHolders(&ValueHolder{24})
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
	set.AddValueHolders(&ValueHolder{1})
	set.AddValueHolders(&ValueHolder{2})
	set.AddValueHolders(&ValueHolder{3})
	collectSet := set.Collect(func(each TreeNodeInterface) CollectionObject {
		return &ValueHolder{each.GetValue().(int) * 10}
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
	set.AddValueHolders(&ValueHolder{1})
	set.AddValueHolders(&ValueHolder{2})
	set.AddValueHolders(&ValueHolder{3})
	testutils.ASSERT_EQ(t, set.Includes(2), true)
	testutils.ASSERT_EQ(t, set.Includes(22), false)
}
