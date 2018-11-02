package oset

import (
	. "github.com/SealNTibbers/gremory/utils"
)

type OSet struct {
	root           TreeNodeInterface
	size           uint64
	valueGenerator func(interface{}) CollectionObject
}

func NewSmartOSet(valueGenerator func(interface{}) CollectionObject) *OSet {
	oset := new(OSet)
	oset.root = GetNilNode()
	oset.valueGenerator = valueGenerator
	return oset
}

func NewOSet() *OSet {
	oset := new(OSet)
	oset.root = GetNilNode()
	return oset
}

func (s *OSet) AddValue(value interface{}) {
	if s.valueGenerator == nil {
		return
	}
	s.AddValueHolders(s.valueGenerator(value))
}

func (s *OSet) AddValueHolders(value CollectionObject) {
	node := NewRBNode()
	node.Data = value
	s.Add(node)
}

func (s *OSet) Add(node *ValueNode) {
	s.root = InsertBST(s.root, node).(*ValueNode)
	s.size = s.size + 1
	s.root = FixAddRBTree(s.root, node)
}

func (s *OSet) Delete(value interface{}) {
	valueNode := NewRBNode()
	valueNode.Data = &ValueHolder{value}
	node := DeleteBST(s.root, valueNode)
	if node != nil {
		s.size = s.size - 1
	}
	s.root = FixDeleteRBTree(s.root, node)
}

func (s *OSet) Size() uint64 {
	if s.root == nil {
		return 0
	}
	return s.size
}

func (s *OSet) walk(visitor Visitor) {
	visitor.Visit(s.root)
}

func (s *OSet) Do(action func(each TreeNodeInterface)) {
	if s.root == nil {
		return
	}
	visitor := &DoVisitor{action}
	s.walk(visitor)
}

func (s *OSet) Collect(collectAction func(each TreeNodeInterface) CollectionObject) *OSet {
	result := NewOSet()
	doAction := func(e TreeNodeInterface) {
		result.AddValueHolders(collectAction(e))
	}
	s.Do(doAction)
	return result
}

func (s *OSet) Select(selectAction func(each TreeNodeInterface) bool) *OSet {
	result := NewOSet()
	doAction := func(e TreeNodeInterface) {
		if selectAction(e) {
			result.AddValueHolders(e.GetData())
		}
	}
	s.Do(doAction)
	return result
}

func (s *OSet) Includes(data interface{}) bool {
	if s.root == nil {
		return false
	}
	selectAction := func(each TreeNodeInterface) bool {
		if each.(*ValueNode).Data.GetValue() == data {
			return true
		}
		return false
	}
	result := s.Select(selectAction)
	return result.size > 0
}
