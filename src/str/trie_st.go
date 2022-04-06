package str

import "strings"

const r_ = 256 // 基数

type charNode struct {
	val  interface{}
	next []*charNode
}

func newCharNode(val interface{}) *charNode {
	return &charNode{
		val:  val,
		next: make([]*charNode, r_),
	}
}

type TrieST struct {
	root *charNode
	n    int
}

func NewTrie() TrieST {
	return TrieST{
		root: newCharNode(nil),
	}
}

func get(x *charNode, key string, d int) *charNode {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := charAt(key, d)
	return get(x.next[c], key, d+1)
}

func (trie *TrieST) Get(key string) interface{} {
	if key == "" {
		return nil
	}
	x := get(trie.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func put(trie *TrieST, x *charNode, key string, val interface{}, d int) *charNode {
	if x == nil {
		x = newCharNode(nil)
	}
	if d == len(key) {
		x.val = val
		trie.n++
		return x
	}
	c := charAt(key, d)
	x.next[c] = put(trie, x.next[c], key, val, d+1)
	return x
}

func (trie *TrieST) Put(key string, val interface{}) {
	if key == "" {
		return
	}
	trie.root = put(trie, trie.root, key, val, 0)
}

func del(trie *TrieST, x *charNode, key string, d int) *charNode {
	// key doesn't exist at the first place
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.val != nil {
			trie.n--
		}
		x.val = nil
	} else {
		c := charAt(key, d)
		x.next[c] = del(trie, x.next[c], key, d+1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.val != nil {
		return x
	}
	for c := 0; c < r_; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

func (trie *TrieST) Delete(key string) {
	if key == "" {
		return
	}
	trie.root = del(trie, trie.root, key, 0)
}

func (trie *TrieST) Contains(key string) bool {
	if key == "" {
		return false
	}
	return trie.Get(key) != nil
}

func (trie *TrieST) Size() int {
	return trie.n
}

func (trie *TrieST) IsEmpty() bool {
	return trie.Size() == 0
}

func collect(x *charNode, prefix *strings.Builder, results *[]string) {
	if x == nil {
		return
	}
	if x.val != nil {
		*results = append(*results, prefix.String())
	}
	for c := 0; c < r_; c++ {
		prefix.WriteByte(byte(c))
		collect(x.next[c], prefix, results)
		prev := prefix.String()[:prefix.Len()-1]
		prefix.Reset()
		prefix.WriteString(prev)
	}
}

func (trie *TrieST) KeysWithPrefix(prefix string) (results []string) {
	x := get(trie.root, prefix, 0)
	var builder strings.Builder
	builder.WriteString(prefix)
	collect(x, &builder, &results)
	return
}

func (trie *TrieST) Keys() []string {
	return trie.KeysWithPrefix("")
}

//
// Related algorithms:
// Patricia trie (save memory)
// Suffix tree
//

// TODO: longestPrefixOf & keysThatMatch

// ********************************************************************* //

type tstNode struct {
	char             byte
	left, mid, right *tstNode
	val              interface{}
}

// TST Ternary Search Tries
type TST struct {
	root *tstNode
	n    int
}
