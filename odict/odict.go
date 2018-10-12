package odict

import . "github.com/SealNTibbers/gremory/utils"

type ODict struct {
	root TreeNodeInterface
	size uint64
}

func NewODict() *ODict {
	odict := new(ODict)
	odict.root = GetNilNode()
	return odict
}

func (d *ODict) AddPair(key interface{}, value interface{}) {
	node := NewDictNode()
	node.Key = new(ValueHolder)
	node.Key.Data = key
	node.Data = new(ValueHolder)
	node.Data.Data = value

	d.Add(node)
}

func (d *ODict) Add(node *KeyValueNode) {
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

func (d *ODict) Size() uint64 {
	if d.root == nil {
		return 0
	}
	return d.size
}

func (d *ODict) walk(visitor Visitor) {
	visitor.Visit(d.root)
}

func (d *ODict) Do(action func(each TreeNodeInterface)) {
	if d.root == nil {
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

func (d *ODict) Collect(collectAction func(each TreeNodeInterface) (interface{}, interface{})) *ODict {
	result := NewODict()
	doAction := func(e TreeNodeInterface) {
		result.AddPair(collectAction(e))
	}
	d.Do(doAction)
	return result
}

func (d *ODict) Select(selectAction func(each TreeNodeInterface) bool) *ODict {
	result := NewODict()
	doAction := func(e TreeNodeInterface) {
		if selectAction(e) {
			result.AddPair(e.(*KeyValueNode).GetKey(), e.(*KeyValueNode).GetValue())
		}
	}
	d.Do(doAction)
	return result
}

func (d *ODict) Includes(key interface{}) bool {
	if d.root == nil {
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
