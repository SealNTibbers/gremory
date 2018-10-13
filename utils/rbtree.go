package utils

type NodeColor byte

const (
	Red         NodeColor = 0
	Black                 = 1
	DoubleBlack           = 2
)

type Direction byte

const (
	LEFT  Direction = 0
	RIGHT           = 1
	NODIR           = 2
)

var nilNodeSingle = new(NilNode)

type TreeNodeInterface interface {
	IsNilNode() bool
	IsNotNilNode() bool
	GetKey() CollectionObject
	GetData() CollectionObject
	GetKeyValue() interface{}
	GetValue() interface{}
	GetColor() NodeColor

	SetColor(color NodeColor)

	acceptNode(rv TreeNodeInterface)

	isEqualTo(nodeInterface TreeNodeInterface) bool
	isLesserThan(nodeInterface TreeNodeInterface) bool
	isGreaterThan(nodeInterface TreeNodeInterface) bool

	setValueFrom(nodeInterface TreeNodeInterface)

	getLeft() TreeNodeInterface
	getRight() TreeNodeInterface
	getParent() TreeNodeInterface

	setLeft(nodeInterface TreeNodeInterface)
	setRight(nodeInterface TreeNodeInterface)
	setParent(nodeInterface TreeNodeInterface)
}

type NilNode struct {
}

func GetNilNode() *NilNode {
	return nilNodeSingle
}

func (*NilNode) IsNilNode() bool {
	return true
}

func (*NilNode) IsNotNilNode() bool {
	return false
}

func (*NilNode) GetValue() interface{} {
	panic("implement me")
}

func (*NilNode) GetKey() CollectionObject {
	panic("implement me")
}
func (*NilNode) GetKeyValue() interface{} {
	panic("implement me")
}

func (*NilNode) GetData() CollectionObject {
	panic("implement me")
}

func (*NilNode) GetColor() NodeColor {
	return Black
}

func (*NilNode) SetColor(color NodeColor) {
	panic("implement me")
}

func (*NilNode) acceptNode(rv TreeNodeInterface) {
	panic("implement me")
}

func (*NilNode) isEqualTo(nodeInterface TreeNodeInterface) bool {
	return false
}

func (*NilNode) isLesserThan(nodeInterface TreeNodeInterface) bool {
	return false
}

func (*NilNode) isGreaterThan(nodeInterface TreeNodeInterface) bool {
	return false
}

func (*NilNode) setValueFrom(nodeInterface TreeNodeInterface) {
	panic("implement me")
}

func (*NilNode) getLeft() TreeNodeInterface {
	return nil
}

func (*NilNode) getRight() TreeNodeInterface {
	return nil
}

func (*NilNode) getParent() TreeNodeInterface {
	return nil
}

func (*NilNode) setLeft(nodeInterface TreeNodeInterface) {
	panic("implement me")
}

func (*NilNode) setRight(nodeInterface TreeNodeInterface) {
	panic("implement me")
}

func (*NilNode) setParent(nodeInterface TreeNodeInterface) {
	panic("implement me")
}

type TreeNode struct {
	Data  CollectionObject
	color NodeColor

	left   TreeNodeInterface
	right  TreeNodeInterface
	parent TreeNodeInterface
}

func (node *TreeNode) IsNilNode() bool {
	return false
}

func (node *TreeNode) IsNotNilNode() bool {
	return true
}

func (node *TreeNode) getLeft() TreeNodeInterface {
	return node.left
}
func (node *TreeNode) getRight() TreeNodeInterface {
	return node.right
}
func (node *TreeNode) getParent() TreeNodeInterface {
	return node.parent
}

func (node *TreeNode) setLeft(nodeInterface TreeNodeInterface) {
	node.left = nodeInterface
}
func (node *TreeNode) setRight(nodeInterface TreeNodeInterface) {
	node.right = nodeInterface
}
func (node *TreeNode) setParent(nodeInterface TreeNodeInterface) {
	node.parent = nodeInterface
}

func (node *TreeNode) GetValue() interface{} {
	return node.Data.GetValue()
}

func (node *TreeNode) GetData() CollectionObject {
	return node.Data
}

func (node *TreeNode) GetColor() NodeColor {
	return node.color
}

func (node *TreeNode) SetColor(color NodeColor) {
	node.color = color
}

func SwapColors(lv TreeNodeInterface, rv TreeNodeInterface) {
	tmpColor := lv.GetColor()
	lv.SetColor(rv.GetColor())
	rv.SetColor(tmpColor)
}

func InsertBST(root TreeNodeInterface, node TreeNodeInterface) TreeNodeInterface {
	if root.IsNilNode() {
		return node
	}
	if root.isEqualTo(node) {
		root.setValueFrom(node)
		return root
	}
	if node.isLesserThan(root) {
		left := InsertBST(root.getLeft(), node)
		root.setLeft(left)
		root.getLeft().setParent(root)
	} else {
		right := InsertBST(root.getRight(), node)
		root.setRight(right)
		root.getRight().setParent(root)
	}
	return root
}

func DeleteBST(root TreeNodeInterface, node TreeNodeInterface) TreeNodeInterface {
	if root.IsNilNode() {
		return root
	}

	if node.isLesserThan(root) {
		return DeleteBST(root.getLeft(), node)
	}

	if node.isGreaterThan(root) {
		return DeleteBST(root.getRight(), node)
	}

	if root.getLeft().IsNilNode() || root.getRight().IsNilNode() {
		return root
	}
	temp := minValueNode(root.getRight())
	root.acceptNode(temp)
	return DeleteBST(root.getRight(), temp)
}

func FixAddRBTree(root TreeNodeInterface, node TreeNodeInterface) TreeNodeInterface {
	var parent TreeNodeInterface
	var grandparent TreeNodeInterface

	for node != root && node.GetColor() == Red && node.getParent().GetColor() == Red {
		parent = node.getParent()
		grandparent = parent.getParent()
		if parent == grandparent.getLeft() {
			uncle := grandparent.getRight()
			if uncle.GetColor() == Red {
				uncle.SetColor(Black)
				parent.SetColor(Black)
				grandparent.SetColor(Red)
				node = grandparent
			} else {
				if node == parent.getRight() {
					root = rotateLeft(root, parent)
					node = parent
					parent = node.getParent()
				}
				root = rotateRight(root, grandparent)
				SwapColors(parent, grandparent)
				node = parent
			}
		} else {
			uncle := grandparent.getLeft()
			if uncle.GetColor() == Red {
				uncle.SetColor(Black)
				parent.SetColor(Black)
				grandparent.SetColor(Red)
				node = grandparent
			} else {
				if node == parent.getLeft() {
					root = rotateRight(root, parent)
					node = parent
					parent = node.getParent()
				}
				root = rotateLeft(root, grandparent)
				SwapColors(parent, grandparent)
				node = parent
			}
		}
	}
	root.SetColor(Black)
	return root
}

func FixDeleteRBTree(root TreeNodeInterface, node TreeNodeInterface) TreeNodeInterface {
	if node.IsNilNode() {
		return nil
	}

	if node == root {
		root = GetNilNode()
		return nil
	}

	if node.GetColor() == Red || node.getLeft().GetColor() == Red || node.getRight().GetColor() == Red {
		var child TreeNodeInterface
		if !node.getLeft().IsNilNode() {
			child = node.getLeft()
		} else {
			child = node.getRight()
		}

		if node == node.getParent().getLeft() {
			node.getParent().setLeft(child)
			if !child.IsNilNode() {
				child.setParent(node.getParent())
			}
		} else {
			node.getParent().setRight(child)
			if !child.IsNilNode() {
				child.setParent(node.getParent())
			}
		}
		child.SetColor(Black)
		node.setParent(GetNilNode())
		node.setLeft(GetNilNode())
		node.setRight(GetNilNode())

	} else {
		var sibling TreeNodeInterface
		var parent TreeNodeInterface
		ptr := node
		ptr.SetColor(DoubleBlack)
		for ptr != root && ptr.GetColor() == DoubleBlack {
			parent = ptr.getParent()
			if ptr == parent.getLeft() {
				sibling = parent.getRight()
				if sibling.GetColor() == Red {
					sibling.SetColor(Black)
					parent.SetColor(Red)
					root = rotateLeft(root, parent)
				} else {
					if sibling.getLeft().GetColor() == Black && sibling.getRight().GetColor() == Black {
						sibling.SetColor(Red)
						if parent.GetColor() == Red {
							parent.SetColor(Black)
						} else {
							parent.SetColor(DoubleBlack)
						}
						ptr = parent
					} else {
						if sibling.getRight().GetColor() == Black {
							sibling.getLeft().SetColor(Black)
							sibling.SetColor(Red)
							root = rotateRight(root, sibling)
							sibling = parent.getRight()
						}
						sibling.SetColor(parent.GetColor())
						parent.SetColor(Black)
						sibling.getRight().SetColor(Black)
						root = rotateLeft(root, parent)
						break
					}
				}
			} else {
				sibling = parent.getLeft()
				if sibling.GetColor() == Red {
					sibling.SetColor(Black)
					parent.SetColor(Red)
					root = rotateRight(root, parent)
				} else {
					if sibling.getLeft().GetColor() == Black && sibling.getRight().GetColor() == Black {
						sibling.SetColor(Red)
						if parent.GetColor() == Red {
							parent.SetColor(Black)
						} else {
							parent.SetColor(DoubleBlack)
						}
						ptr = parent
					} else {
						if sibling.getLeft().GetColor() == Black {
							sibling.getRight().SetColor(Black)
							sibling.SetColor(Red)
							root = rotateLeft(root, sibling)
							sibling = parent.getLeft()
						}
						sibling.SetColor(parent.GetColor())
						parent.SetColor(Black)
						sibling.getLeft().SetColor(Black)
						root = rotateRight(root, parent)
						break
					}
				}
			}
		}
		if node == node.getParent().getLeft() {
			node.getParent().setLeft(GetNilNode())
		} else {
			node.getParent().setRight(GetNilNode())
		}
		node.setParent(GetNilNode())
		node.setLeft(GetNilNode())
		node.setRight(GetNilNode())
		root.SetColor(Black)
	}
	return root
}

func rotateLeft(root TreeNodeInterface, node TreeNodeInterface) TreeNodeInterface {
	nodeToTurnLeft := node.getRight()
	/* Turn nodeToTurnLeft's left sub-tree into node's right sub-tree */
	node.setRight(nodeToTurnLeft.getLeft())

	if nodeToTurnLeft.getLeft().IsNotNilNode() {
		nodeToTurnLeft.getLeft().setParent(node)
	}
	/* nodeToTurnLeft's new parent was node's parent */
	nodeToTurnLeft.setParent(node.getParent())
	/* Set the parent to point to nodeToTurnLeft instead of node */
	/* First see whether we're at the root */
	if node.getParent().IsNilNode() {
		root = nodeToTurnLeft
	} else if node == node.getParent().getLeft() {
		/* node was on the left of its parent */
		node.getParent().setLeft(nodeToTurnLeft)
	} else {
		/* node must have been on the right */
		node.getParent().setRight(nodeToTurnLeft)
	}
	/* Finally, put node on nodeToTurnLeft's left */
	nodeToTurnLeft.setLeft(node)
	node.setParent(nodeToTurnLeft)
	return root
}

func rotateRight(root TreeNodeInterface, node TreeNodeInterface) TreeNodeInterface {
	nodeToTurnRight := node.getLeft()
	/* Turn nodeToTurnRight's left sub-tree into node's right sub-tree */
	node.setLeft(nodeToTurnRight.getRight())

	if nodeToTurnRight.getRight().IsNotNilNode() {
		nodeToTurnRight.getRight().setParent(node)
	}
	/* nodeToTurnRight's new parent was node's parent */
	nodeToTurnRight.setParent(node.getParent())
	/* Set the parent to point to nodeToTurnRight instead of node */
	/* First see whether we're at the root */
	if node.getParent().IsNilNode() {
		root = nodeToTurnRight
	} else if node == node.getParent().getRight() {
		/* node was on the right of its parent */
		node.getParent().setRight(nodeToTurnRight)
	} else {
		/* node must have been on the left */
		node.getParent().setLeft(nodeToTurnRight)
	}
	/* Finally, put node on nodeToTurnRight's left */
	nodeToTurnRight.setRight(node)
	node.setParent(nodeToTurnRight)
	return root
}

func minValueNode(node TreeNodeInterface) TreeNodeInterface {
	ptr := node

	for !ptr.getLeft().IsNilNode() {
		ptr = ptr.getLeft()
	}

	return ptr
}

type Visitor interface {
	Visit(node TreeNodeInterface)
}

type DoVisitor struct {
	Action func(each TreeNodeInterface)
}

func (v *DoVisitor) Visit(node TreeNodeInterface) {
	if node.IsNilNode() {
		return
	}

	v.Visit(node.getLeft())
	v.Action(node)
	v.Visit(node.getRight())
}

type ValueNode struct {
	*TreeNode
}

func NewRBNode() *ValueNode {
	node := new(ValueNode)
	node.TreeNode = new(TreeNode)
	node.parent = GetNilNode()
	node.right = GetNilNode()
	node.left = GetNilNode()
	return node
}

func (node *ValueNode) isEqualTo(nodeInterface TreeNodeInterface) bool {
	return node.Data.Equal(nodeInterface.(*ValueNode).Data)
}

func (node *ValueNode) isLesserThan(nodeInterface TreeNodeInterface) bool {
	return node.Data.Less(nodeInterface.(*ValueNode).Data)
}

func (node *ValueNode) isGreaterThan(nodeInterface TreeNodeInterface) bool {
	return node.Data.Greater(nodeInterface.(*ValueNode).Data)
}

func (node *ValueNode) setValueFrom(nodeInterface TreeNodeInterface) {
}

func (*ValueNode) GetKey() CollectionObject {
	return nil
}

func (*ValueNode) GetKeyValue() interface{} {
	return nil
}

func (node *ValueNode) acceptNode(rv TreeNodeInterface) {
	node.Data = rv.(*ValueNode).Data
}

type KeyValueNode struct {
	*TreeNode
	Key CollectionObject
}

func NewDictNode() *KeyValueNode {
	node := new(KeyValueNode)
	node.TreeNode = new(TreeNode)
	node.parent = GetNilNode()
	node.right = GetNilNode()
	node.left = GetNilNode()
	return node
}

func (node *KeyValueNode) isEqualTo(nodeInterface TreeNodeInterface) bool {
	if nodeInterface.IsNilNode() {
		return false
	}
	return node.Key.Equal(nodeInterface.(*KeyValueNode).Key)
}

func (node *KeyValueNode) isLesserThan(nodeInterface TreeNodeInterface) bool {
	if nodeInterface.IsNilNode() {
		return false
	}
	return node.Key.Less(nodeInterface.(*KeyValueNode).Key)
}

func (node *KeyValueNode) isGreaterThan(nodeInterface TreeNodeInterface) bool {
	if nodeInterface.IsNilNode() {
		return false
	}
	return node.Key.Greater(nodeInterface.(*KeyValueNode).Key)
}

func (node *KeyValueNode) setValueFrom(nodeInterface TreeNodeInterface) {
	node.Data = nodeInterface.(*KeyValueNode).Data
}

func (node *KeyValueNode) acceptNode(rv TreeNodeInterface) {
	node.Key = rv.(*KeyValueNode).Key
	node.Data = rv.(*KeyValueNode).Data
}

func (node *KeyValueNode) GetKey() CollectionObject {
	return node.Key
}

func (node *KeyValueNode) GetKeyValue() interface{} {
	return node.Key.GetValue()
}

func internalLookup(parent TreeNodeInterface, this TreeNodeInterface, key interface{}, dir Direction) (TreeNodeInterface, bool, Direction) {
	tmpNode := new(KeyValueNode)
	tmpNode.Key = &ValueHolder{key}
	switch {
	case this == nil:
		return parent, false, dir
	case tmpNode.isEqualTo(this):
		return parent, true, dir
	case tmpNode.isLesserThan(this):
		return internalLookup(this, this.getLeft(), key, LEFT)
	case tmpNode.isGreaterThan(this):
		return internalLookup(this, this.getRight(), key, RIGHT)
	default:
		return parent, false, NODIR
	}
}

func getParent(root TreeNodeInterface, key interface{}) (parent TreeNodeInterface, found bool, dir Direction) {
	if root == nil {
		return nil, false, NODIR
	}

	return internalLookup(nil, root, key, NODIR)
}

func GetNode(root TreeNodeInterface, key interface{}) (TreeNodeInterface, bool) {
	parent, found, dir := getParent(root, key)
	if found {
		if parent == nil {
			return root, true
		} else {
			var node TreeNodeInterface
			switch dir {
			case LEFT:
				node = parent.getLeft()
			case RIGHT:
				node = parent.getRight()
			}

			if node != nil {
				return node, true
			}
		}
	}
	return nil, false
}
