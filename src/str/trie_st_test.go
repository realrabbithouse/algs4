package str

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
