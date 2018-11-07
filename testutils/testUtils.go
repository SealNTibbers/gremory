package testutils

import "testing"
import . "github.com/SealNTibbers/gremory/utils"

type TestType struct {
	Id   int
	Name string
}

type TestTypeHolder struct {
	Data TestType
}

func (holder *TestTypeHolder) GetValue() interface{} {
	return holder.Data
}

func (lv *TestTypeHolder) Less(rv CollectionObject) bool {
	rvValue, ok := rv.(*TestTypeHolder)
	if !ok {
		return false
	}
	return lv.Data.Id < rvValue.Data.Id
}

func (lv *TestTypeHolder) Greater(rv CollectionObject) bool {
	rvValue, ok := rv.(*TestTypeHolder)
	if !ok {
		return false
	}
	return lv.Data.Id > rvValue.Data.Id
}

func (lv *TestTypeHolder) Equal(rv CollectionObject) bool {
	rvValue, ok := rv.(*TestTypeHolder)
	if !ok {
		return false
	}
	return lv.Data.Id > rvValue.Data.Id
}

func ASSERT_EQ(t *testing.T, actual interface{}, expected interface{}) {
	if expected != actual {
		t.Fatalf("expected=%d, got=%d", expected, actual)
	}
}
