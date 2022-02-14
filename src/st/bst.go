package st

type Node struct {
	key   Key
	value Value
	left  *Node
	right *Node
	sz    int
}

// ***************************** helper function ***************************** //

func contain(node *Node, key Key) bool {
	if node == nil {
		return false
	}
	cmp := key.CompareTo(node.key)
	if cmp < 0 {
		return contain(node.left, key)
	} else if cmp > 0 {
		return contain(node.right, key)
	} else {
		return true
	}
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.sz
}

func min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return min(node.left)
}

func max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return max(node.right)
}

func put(node *Node, key Key, value Value) *Node {
	if node == nil {
		return &Node{
			key:   key,
			value: value,
			sz:    1,
		}
	}
	cmp := key.CompareTo(node.key)
	if cmp < 0 {
		node.left = put(node.left, key, value)
	} else if cmp > 0 {
		node.right = put(node.right, key, value)
	} else {
		node.value = value
	}
	node.sz = size(node.left) + size(node.right) + 1
	return node
}

func get(node *Node, key Key) Value {
	if node == nil {
		return nil
	}
	cmp := key.CompareTo(node.key)
	if cmp < 0 {
		return get(node.left, key)
	} else if cmp > 0 {
		return get(node.right, key)
	} else {
		return node.value
	}
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	node.sz = size(node.left) + size(node.right) + 1
	return node
}

func deleteMax(node *Node) *Node {
	if node.right == nil {
		return node.left
	}
	node.right = deleteMax(node.right)
	node.sz = size(node.left) + size(node.right) + 1
	return node
}

func del(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.CompareTo(node.key)
	if cmp < 0 {
		node.left = del(node.left, key)
	} else if cmp > 0 {
		node.right = del(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		var tmp *Node = node
		node = min(tmp.right)
		node.right = deleteMin(tmp.right)
		node.left = tmp.left
	}
	node.sz = size(node.left) + size(node.right) + 1
	return node
}

// floor returns the maximum node less than the given key.
func floor(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.CompareTo(node.key)
	if cmp == 0 {
		return node
	}
	if cmp < 0 {
		return floor(node.left, key)
	}
	t := floor(node.right, key)
	if t != nil {
		return t
	} else {
		return node
	}
}

// ceiling returns the minimum node greater than the given key.
func ceiling(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.CompareTo(node.key)
	if cmp == 0 {
		return node
	}
	if cmp < 0 {
		t := ceiling(node.left, key)
		if t != nil {
			return t
		} else {
			return node
		}
	}
	return ceiling(node.right, key)
}

// choose the node whose rank is k (k = 1,2,3...).
// 树中正好有k个小于它的键
func choose(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	sz := size(node.left)
	if sz > k {
		return choose(node.left, k)
	} else if sz < k {
		return choose(node.right, k-sz-1)
	} else {
		return node
	}
}

// 返回小于key的键的数量
// rank 与 choose 互为逆操作
func rank(node *Node, key Key) int {
	if node == nil {
		return 0
	}
	cmp := key.CompareTo(node.key)
	if cmp < 0 {
		return rank(node.left, key)
	}
	if cmp > 0 {
		return 1 + size(node.left) + rank(node.right, key)
	}
	return size(node)
}

// ********************************************************************* //

type BST struct {
	root *Node
}

func (bst BST) Contains(key Key) bool {
	return contain(bst.root, key)
}

func (bst BST) IsEmpty() bool {
	return bst.root == nil
}

func (bst BST) Size() int {
	return bst.root.sz
}

func (bst *BST) Put(key Key, value Value) {
	bst.root = put(bst.root, key, value)
}

func (bst BST) Get(key Key) Value {
	return get(bst.root, key)
}

func (bst *BST) DeleteMin() {
	bst.root = deleteMin(bst.root)
}

func (bst *BST) DeleteMax() {
	bst.root = deleteMax(bst.root)
}

func (bst BST) Min() Key {
	if bst.root == nil {
		return nil
	}
	return min(bst.root).key
}

func (bst BST) Max() Key {
	if bst.root == nil {
		return nil
	}
	return max(bst.root).key
}

func (bst *BST) Delete(key Key) {
	bst.root = del(bst.root, key)
}

func (bst BST) Floor(key Key) Key {
	if bst.root == nil || key == nil {
		return nil
	}
	r := floor(bst.root, key)
	if r == nil {
		return nil
	}
	return r.key
}

func (bst BST) Ceiling(key Key) Key {
	if bst.root == nil || key == nil {
		return nil
	}
	r := ceiling(bst.root, key)
	if r == nil {
		return nil
	}
	return r.key
}

func (bst BST) Choose(k int) Key {
	if k < 0 || k >= bst.Size() {
		panic("out of range")
	}
	return choose(bst.root, k).key
}

func (bst BST) Rank(key Key) int {
	return rank(bst.root, key)
}

// Get all keys in the symbol table in the given range. //

func keys(node *Node, lo, hi Key, sli *[]Key) {
	if node == nil {
		return
	}
	cmplo := lo.CompareTo(node.key)
	cmphi := hi.CompareTo(node.key)
	if cmplo < 0 {
		keys(node.left, lo, hi, sli)
	}
	if cmplo <= 0 && cmphi >= 0 {
		*sli = append(*sli, node.key)
	}
	if cmphi > 0 {
		keys(node.right, lo, hi, sli)
	}
}

func (bst BST) ToSlice() []Key {
	sli := make([]Key, 0, bst.Size())
	keys(bst.root, bst.Min(), bst.Max(), &sli)
	return sli
}
