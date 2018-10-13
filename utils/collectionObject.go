package utils

type Comparable interface {
	Less(rv CollectionObject) bool
	Greater(rv CollectionObject) bool
	Equal(rv CollectionObject) bool
}

type CollectionObject interface {
	Comparable
	GetValue() interface{}
	SetValue(value interface{})
}

type ValueHolder struct {
	Data interface{}
}

func (holder *ValueHolder) GetValue() interface{} {
	return holder.Data
}

func (holder *ValueHolder) SetValue(value interface{}) {
	holder.Data = value
}

func (lv *ValueHolder) Less(rv CollectionObject) bool {
	lvInt, lvIntOk := lv.Data.(int)
	rvInt, rvIntOk := rv.(*ValueHolder).Data.(int)
	if lvIntOk && rvIntOk {
		return lvInt < rvInt
	}

	lvString, lvStringOk := lv.Data.(string)
	rvString, rvStringOk := rv.(*ValueHolder).Data.(string)
	if lvStringOk && rvStringOk {
		return lvString < rvString
	}

	panic("This ValueHolder only for basic types.")
}

func (lv *ValueHolder) Greater(rv CollectionObject) bool {
	lvInt, lvIntOk := lv.Data.(int)
	rvInt, rvIntOk := rv.(*ValueHolder).Data.(int)
	if lvIntOk && rvIntOk {
		return lvInt > rvInt
	}

	lvString, lvStringOk := lv.Data.(string)
	rvString, rvStringOk := rv.(*ValueHolder).Data.(string)
	if lvStringOk && rvStringOk {
		return lvString > rvString
	}

	panic("This ValueHolder only for basic types.")
}

func (lv *ValueHolder) Equal(rv CollectionObject) bool {
	lvInt, lvIntOk := lv.Data.(int)
	rvInt, rvIntOk := rv.(*ValueHolder).Data.(int)
	if lvIntOk && rvIntOk {
		return lvInt == rvInt
	}

	lvString, lvStringOk := lv.Data.(string)
	rvString, rvStringOk := rv.(*ValueHolder).Data.(string)
	if lvStringOk && rvStringOk {
		return lvString == rvString
	}

	panic("This ValueHolder only for basic types.")
}
