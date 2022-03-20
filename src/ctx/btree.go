package ctx

import (
	"algs4/src/rabbit"
	"errors"
	"fmt"
	"strings"
)

const _M uint32 = 4 // B-tree of order M, an even number

type state uint8

const (
	_delete state = iota + 1
	_replace
	_insert
	_cancel
)

type entry struct {
	key  rabbit.Comparable
	val  interface{}
	next *node
}

func newEntry(key rabbit.Comparable, val interface{}) entry {
	return entry{
		key: key,
		val: val,
	}
}

type node struct {
	x          uint32 // number of children
	children   []entry
	isExternal bool
}

func newNode(x uint32) *node {
	return &node{
		x:          x,
		children:   make([]entry, _M),
		isExternal: true,
	}
}

type BTree struct {
	root   *node  // root of the B-tree
	n      uint64 // number of elements in the B-Tree
	height uint32 // height of the B-Tree
}

func NewBTree() *BTree {
	return &BTree{
		root: newNode(0),
	}
}

func (t BTree) IsEmpty() bool {
	return t.n == 0
}

func (t BTree) Size() uint64 {
	return t.n
}

func (t BTree) Height() uint32 {
	return t.height
}

func search(node *node, key rabbit.Comparable, h uint32) (val interface{}) {
	if h == 0 {
		for i := uint32(0); i < node.x; i++ {
			if rabbit.Equal(key, node.children[i].key) {
				return node.children[i].val
			}
		}
	} else {
		for i := uint32(0); i < node.x; i++ {
			if i+1 == node.x || rabbit.Less(key, node.children[i+1].key) {
				return search(node.children[i].next, key, h-1)
			}
		}
	}

	return nil
}

// Get returns the value associated with the given key, and returns nil if the
// key is not in the symbol table.
func (t BTree) Get(key rabbit.Comparable) (val interface{}) {
	if key == nil {
		fmt.Println("error: nil key")
		return nil
	}
	return search(t.root, key, t.height)
}

func insert(node *node, key rabbit.Comparable, val interface{}, h uint32) (r *node, code state) {
	var (
		j   uint32               // which position to insert?
		ent = newEntry(key, val) // will bring cascading insert?
	)

	if h == 0 {
		// external node
		for j = uint32(0); j < node.x; j++ {
			if rabbit.Equal(key, node.children[j].key) {
				if val == nil {
					// delete
					if node.x == 1 {
						node.x--
						return nil, _delete
					}
					for i := j; i < node.x-1; i++ {
						node.children[j] = node.children[j+1]
					}
					node.x--
					return nil, _delete
				}
				// replace
				node.children[j].val = val
				return nil, _replace
			}
			if rabbit.Less(key, node.children[j].key) {
				// insert, later
				break
			}
		}
	} else {
		// internal node
		for j = 0; j < node.x; j++ {
			if j+1 == node.x || rabbit.Less(key, node.children[j+1].key) {
				r, code = insert(node.children[j].next, key, val, h-1)
				j++
				if r == nil {
					return
				}
				// 理解这一步的逻辑十分关键！
				ent.next = r
				ent.key = r.children[0].key
				ent.val = nil
				break
			}
		}
	}

	// cancel
	if val == nil {
		return nil, _cancel
	}

	// insert, external or internal, same logic
	for i := node.x; i > j; i-- {
		node.children[i] = node.children[i-1]
	}
	node.children[j] = ent
	node.x++
	code = _insert
	if node.x < _M {
		r = nil
	} else {
		fmt.Printf("split when insert %v:%v\n", key, val)
		r = split(node)
	}
	return
}

func split(a *node) *node {
	b := newNode(_M / 2)
	a.x = _M / 2 // lazy, no data erase
	for i := uint32(0); i < _M/2; i++ {
		b.children[i] = a.children[i+_M/2]
	}
	a.isExternal, b.isExternal = false, false
	return b
}

// Put inserts the key-value pair into the symbol table, overwriting the old value
// with the new value if the key is already in the symbol table.
// If the put value is nil, this effectively deletes the key from the symbol table.
func (t *BTree) Put(key rabbit.Comparable, val interface{}) {
	if key == nil {
		fmt.Println("error: nil key")
		return
	}
	r, code := insert(t.root, key, val, t.height)
	switch code {
	case _delete:
		fmt.Printf("delete %v:%v\n", key, val)
		t.n--
	case _replace:
		fmt.Printf("repalce %v:%v\n", key, val)
	case _insert:
		t.n++
		fmt.Printf("insert %v:%v\n", key, val)
	case _cancel:
		fmt.Printf("cancel %v:%v\n", key, val)
	}
	if r != nil {
		x := newNode(2)
		x.children[0] = entry{
			key:  t.root.children[0].key,
			val:  nil,
			next: t.root,
		}
		x.children[1] = entry{
			key:  r.children[0].key,
			val:  nil,
			next: r,
		}
		t.root = x
		t.height++
	}
}

func validate(node *node, lo rabbit.Comparable) (err error) {
	if node.isExternal {
		for i := uint32(0); i < node.x; i++ {
			if i == 0 && !rabbit.Equal(lo, node.children[i].key) {
				err = errors.New(fmt.Sprintf("internal: first entry unequal: %v", lo))
				return
			}
			if i > 0 && !rabbit.Less(lo, node.children[i].key) {
				err = errors.New(fmt.Sprintf("internal: relation discrepancy: %v", lo))
				return
			}
		}
	} else {
		for i := uint32(0); i < node.x; i++ {
			if i == 0 && !rabbit.Equal(lo, node.children[i].key) {
				err = errors.New(fmt.Sprintf("external: first entry unequal: %v", lo))
				return
			}
			if i > 0 && !rabbit.Less(lo, node.children[i].key) {
				err = errors.New(fmt.Sprintf("external: relation discrepancy: %v", lo))
				return
			}
			err = validate(node.children[i].next, node.children[i].key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t BTree) Validate() (err error) {
	if t.root.x > 0 {
		err = validate(t.root, t.root.children[0].key)
	}
	return
}

func (t BTree) String() string {
	return t.toString(t.root, t.height, " ")
}

func (t BTree) toString(node *node, h uint32, indent string) string {
	builder := strings.Builder{}
	children := node.children
	if h == 0 {
		for i := uint32(0); i < node.x; i++ {
			builder.WriteString(fmt.Sprintf("%s[%v %v]\n", indent, children[i].key, children[i].val))
		}
	} else {
		for i := uint32(0); i < node.x; i++ {
			builder.WriteString(fmt.Sprintf("(%v):\n", children[i].key))
			builder.WriteString(t.toString(children[i].next, h-1, "	"))
		}
	}
	return builder.String()
}
