package streye

import (
	"fmt"
	"testing"
)

func TestTrieST(t *testing.T) {
	x := 100
	trie := NewTrie()
	keys := GenRandomStrings(x, 10)
	for i := 0; i < x; i += 10 {
		fmt.Println(keys[i : i+10])
	}

	for i := range keys {
		trie.Put(keys[i], keys[i])
	}
	fmt.Println("number of keys:", trie.Size())

	for i := 0; i < x; i++ {
		if keys[i] != trie.Get(keys[i]).(string) {
			fmt.Println("error!")
		}
	}

	all := trie.Keys()
	fmt.Println("number of keys in trie:", len(all))
	for i := 0; i < len(all); i += 10 {
		fmt.Println(all[i : i+10])
	}
}

func TestTrieST_Put_Delete(t *testing.T) {
	var tests = []struct {
		k string
		v int
	}{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"oneee", 111},
		{"twooo", 222},
		{"threeee", 333},
	}
	trie := NewTrie()
	n := len(tests)
	for i := 0; i < n; i++ {
		trie.Put(tests[i].k, tests[i].v)
	}
	for _, test := range tests {
		if got := trie.Get(test.k); got.(int) != test.v {
			t.Errorf("trie.Get(%v) = %v, should be %v", test.k, got, test.v)
		}
	}
	N := trie.Size()
	for i := n - 1; i >= 0; i-- {
		trie.Delete(tests[i].k)
		N--
		if trie.Size() != N {
			t.Errorf("Size = %v, Expected = %v", trie.Size(), N)
		}
	}
}
