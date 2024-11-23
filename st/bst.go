package st

type Node struct {
	key   Key
	value Value
	left  *Node
	right *Node
	size  int
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.size
}

// ********************************************************************* //

// BST 二叉树
type BST struct {
	root *Node
}

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

func (bst BST) Contains(key Key) bool {
	return contain(bst.root, key)
}

func (bst BST) IsEmpty() bool {
	return bst.root == nil
}

func (bst BST) Size() int {
	return bst.root.size
}

func put(node *Node, k Key, v Value) *Node {
	if node == nil {
		return &Node{
			key:   k,
			value: v,
			size:  1,
		}
	}
	cmp := k.CompareTo(node.key)
	if cmp < 0 {
		node.left = put(node.left, k, v)
	} else if cmp > 0 {
		node.right = put(node.right, k, v)
	} else {
		node.value = v
	}
	node.size = size(node.left) + size(node.right) + 1
	return node
}

func (bst *BST) Put(k Key, v Value) {
	bst.root = put(bst.root, k, v)
}

func get(node *Node, k Key) Value {
	if node == nil {
		return nil
	}
	cmp := k.CompareTo(node.key)
	if cmp < 0 {
		return get(node.left, k)
	} else if cmp > 0 {
		return get(node.right, k)
	} else {
		return node.value
	}
}

func (bst BST) Get(k Key) Value {
	return get(bst.root, k)
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	node.size = size(node.left) + size(node.right) + 1
	return node
}

func (bst *BST) DeleteMin() {
	bst.root = deleteMin(bst.root)
}

func deleteMax(node *Node) *Node {
	if node.right == nil {
		return node.left
	}
	node.right = deleteMax(node.right)
	node.size = size(node.left) + size(node.right) + 1
	return node
}

func (bst *BST) DeleteMax() {
	bst.root = deleteMax(bst.root)
}

func min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return min(node.left)
}

func (bst BST) Min() Key {
	if bst.root == nil {
		return nil
	}
	return min(bst.root).key
}

func max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return max(node.right)
}

func (bst BST) Max() Key {
	if bst.root == nil {
		return nil
	}
	return max(bst.root).key
}

// deleteKey removes the node with the specified key from the binary search tree.
// It returns the new root of the subtree after deletion.
//
// Parameters:
//   - node: The root of the subtree from which the key should be deleted.
//   - key: The key of the node to be deleted.
//
// Returns:
//   - *Node: The new root of the subtree after the key has been deleted. If the key is not found, it returns the original node.
func deleteKey(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.CompareTo(node.key)
	if cmp < 0 {
		node.left = deleteKey(node.left, key)
	} else if cmp > 0 {
		node.right = deleteKey(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		var tmp *Node = node
		// Replace the current node with the minimum node of right side.
		node = min(tmp.right)
		node.right = deleteMin(tmp.right)
		node.left = tmp.left
	}
	node.size = size(node.left) + size(node.right) + 1
	return node
}

func (bst *BST) Delete(key Key) {
	bst.root = deleteKey(bst.root, key)
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

// The Floor of 2.31 is 2. The Floor of -2.31 is -3.
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

// The Ceiling of 2.31 is 3. The Ceiling of -2.31 is -2.
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

func (bst BST) Choose(k int) Key {
	if k < 0 || k >= bst.Size() {
		panic("out of range")
	}
	return choose(bst.root, k).key
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

func (bst BST) Rank(key Key) int {
	return rank(bst.root, key)
}

// Get all keys in the symbol table in given range.
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

func (bst BST) GetKeys() []Key {
	sli := make([]Key, 0, bst.Size())
	keys(bst.root, bst.Min(), bst.Max(), &sli)
	return sli
}
