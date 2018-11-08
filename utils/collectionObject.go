package utils

import (
	"math"
	"reflect"
)

var FLOAT_EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < FLOAT_EPSILON {
		return true
	}
	return false
}

type Comparable interface {
	Less(rv CollectionObject) bool
	Greater(rv CollectionObject) bool
	Equal(rv CollectionObject) bool
}

type CollectionObject interface {
	Comparable
	GetValue() interface{}
}

type ValueHolder struct {
	Data interface{}
}

func (holder *ValueHolder) GetValue() interface{} {
	return holder.Data
}

func (lv *ValueHolder) Less(rv CollectionObject) bool {
	if reflect.TypeOf(lv.Data) != reflect.TypeOf(rv.(*ValueHolder).Data) {
		return false
	}

	lvInt, lvIntOk := lv.Data.(int)
	rvInt, rvIntOk := rv.(*ValueHolder).Data.(int)
	if lvIntOk && rvIntOk {
		return lvInt < rvInt
	}

	lvFloat, lvFloatOk := lv.Data.(float64)
	rvFloat, rvFloatOk := rv.(*ValueHolder).Data.(float64)
	if lvFloatOk && rvFloatOk {
		return lvFloat < rvFloat
	}

	lvString, lvStringOk := lv.Data.(string)
	rvString, rvStringOk := rv.(*ValueHolder).Data.(string)
	if lvStringOk && rvStringOk {
		return lvString < rvString
	}

	lvRune, lvRuneOk := lv.Data.(rune)
	rvRune, rvRuneOk := rv.(*ValueHolder).Data.(rune)
	if lvRuneOk && rvRuneOk {
		return lvRune < rvRune
	}

	panic("This ValueHolder only for basic types.")
}

func (lv *ValueHolder) Greater(rv CollectionObject) bool {
	if reflect.TypeOf(lv.Data) != reflect.TypeOf(rv.(*ValueHolder).Data) {
		return false
	}

	lvInt, lvIntOk := lv.Data.(int)
	rvInt, rvIntOk := rv.(*ValueHolder).Data.(int)
	if lvIntOk && rvIntOk {
		return lvInt > rvInt
	}

	lvFloat, lvFloatOk := lv.Data.(float64)
	rvFloat, rvFloatOk := rv.(*ValueHolder).Data.(float64)
	if lvFloatOk && rvFloatOk {
		return lvFloat > rvFloat
	}

	lvString, lvStringOk := lv.Data.(string)
	rvString, rvStringOk := rv.(*ValueHolder).Data.(string)
	if lvStringOk && rvStringOk {
		return lvString > rvString
	}

	lvRune, lvRuneOk := lv.Data.(rune)
	rvRune, rvRuneOk := rv.(*ValueHolder).Data.(rune)
	if lvRuneOk && rvRuneOk {
		return lvRune > rvRune
	}

	panic("This ValueHolder only for basic types.")
}

func (lv *ValueHolder) Equal(rv CollectionObject) bool {
	if reflect.TypeOf(lv.Data) != reflect.TypeOf(rv.(*ValueHolder).Data) {
		return false
	}

	lvInt, lvIntOk := lv.Data.(int)
	rvInt, rvIntOk := rv.(*ValueHolder).Data.(int)

	if lvIntOk && rvIntOk {
		return lvInt == rvInt
	}

	lvFloat, lvFloatOk := lv.Data.(float64)
	rvFloat, rvFloatOk := rv.(*ValueHolder).Data.(float64)
	if lvFloatOk && rvFloatOk {
		return floatEquals(lvFloat, rvFloat)
	}

	lvString, lvStringOk := lv.Data.(string)
	rvString, rvStringOk := rv.(*ValueHolder).Data.(string)
	if lvStringOk && rvStringOk {
		return lvString == rvString
	}

	lvRune, lvRuneOk := lv.Data.(rune)
	rvRune, rvRuneOk := rv.(*ValueHolder).Data.(rune)
	if lvRuneOk && rvRuneOk {
		return lvRune == rvRune
	}

	panic("This ValueHolder only for basic types.")
}
