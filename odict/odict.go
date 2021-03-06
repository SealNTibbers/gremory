package odict

import . "github.com/SealNTibbers/gremory/utils"

type ODict struct {
	root           TreeNodeInterface
	size           uint64
	keyGenerator   func(interface{}) CollectionObject
	valueGenerator func(interface{}) CollectionObject
}

func NewReadyODict() *ODict {
	keyGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	odict := new(ODict)
	odict.root = GetNilNode()
	odict.keyGenerator = keyGen
	odict.valueGenerator = valueGen
	return odict
}

func NewSmartODict(keyGenerator func(interface{}) CollectionObject, valueGenerator func(interface{}) CollectionObject) *ODict {
	odict := new(ODict)
	odict.root = GetNilNode()
	odict.keyGenerator = keyGenerator
	odict.valueGenerator = valueGenerator
	return odict
}

func NewODict() *ODict {
	odict := new(ODict)
	odict.root = GetNilNode()
	odict.keyGenerator = nil
	odict.valueGenerator = nil
	return odict
}

func (d *ODict) AddPair(key interface{}, value interface{}) {
	if d.keyGenerator == nil || d.valueGenerator == nil {
		return
	}
	d.addCollectionObjectsPair(d.keyGenerator(key), d.valueGenerator(value))
}

func (d *ODict) addCollectionObjectsPair(key CollectionObject, value CollectionObject) {
	node := NewDictNode()
	node.Key = key
	node.Data = value

	d.addNode(node)
}

func (d *ODict) addNode(node *KeyValueNode) {
	d.root = InsertBST(d.root, node).(*KeyValueNode)
	d.size = d.size + 1
	d.root = FixAddRBTree(d.root, node)
}

func (d *ODict) DeleteKey(key interface{}) {
	valueNode := NewDictNode()
	valueNode.Key = &ValueHolder{key}
	node := DeleteBST(d.root, valueNode)
	if node != nil {
		d.size = d.size - 1
	}
	d.root = FixDeleteRBTree(d.root, node)
}

func (d *ODict) IsEmpty() bool {
	return d.root == nil
}

func (d *ODict) Size() uint64 {
	if d.IsEmpty() {
		return 0
	}
	return d.size
}

func (d *ODict) walk(visitor Visitor) {
	visitor.Visit(d.root)
}

func (d *ODict) ReverseDo(action func(each TreeNodeInterface)) {
	if d.IsEmpty() {
		return
	}
	visitor := &ReverseDoVisitor{Action: action}
	d.walk(visitor)
}

func (d *ODict) Do(action func(each TreeNodeInterface)) {
	if d.IsEmpty() {
		return
	}
	visitor := &DoVisitor{Action: action}
	d.walk(visitor)
}

func (d *ODict) At(key interface{}) interface{} {
	node, ok := GetNode(d.root, key)

	if ok {
		return node.GetValue()
	} else {
		return nil
	}
}

func (d *ODict) Collect(collectAction func(each TreeNodeInterface) (CollectionObject, CollectionObject)) *ODict {
	result := NewODict()
	doAction := func(e TreeNodeInterface) {
		result.addCollectionObjectsPair(collectAction(e))
	}
	d.Do(doAction)
	return result
}

func (d *ODict) Select(selectAction func(each TreeNodeInterface) bool) *ODict {
	result := NewODict()
	doAction := func(e TreeNodeInterface) {
		if selectAction(e) {
			result.addCollectionObjectsPair(e.GetKey(), e.GetData())
		}
	}
	d.Do(doAction)
	return result
}

func (d *ODict) Includes(key interface{}) bool {
	if d.IsEmpty() {
		return false
	}
	selectAction := func(each TreeNodeInterface) bool {
		if each.(*KeyValueNode).Key.GetValue() == key {
			return true
		}
		return false
	}
	result := d.Select(selectAction)
	return result.size > 0
}
