/**
  * Author: JeffreyBool
  * Date: 2019/5/2
  * Time: 19:15
  * Software: GoLand
*/

package bst

type node struct {
	e           int
	left, right *node
}

func newNode(e int) *node {
	return &node{e:e}
}

type Bst struct {
	root *node
	size int
}

func NewBst() *Bst {
	return &Bst{}
}

func (bst *Bst) Size() int {
	return bst.size
}

func (bst *Bst) IsEmpty() bool {
	return bst.size == 0
}

// 向二分搜索树中添加新的元素e
func (bst *Bst) Add(e int)  {
	// tree 为空，设置根节点
	if bst.root == nil {
		bst.root = newNode(e)
		bst.size ++
	}else{
		bst.add(bst.root,e)
	}
}

// 向以node为根的二分搜索树中插入元素e，递归算法
func (bst *Bst) add(node *node,e int) {
	// 不处理重复数据的节点
	if node.e == e{
		return
	// 左子树递归终止条件
	}else if e < node.e && node.left == nil{
		node.left = newNode(e)
		bst.size ++
		return
	}else if e > node.e && node.right == nil{
		node.right = newNode(e)
		bst.size ++
	}

	// 递归调用
	if e < node.e {
		bst.add(node.left,e)
	}else{
		bst.add(node.right,e)
	}
}