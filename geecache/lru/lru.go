package lru

import "container/list"

// Value use Len to count how many bytes it takes.
type Value interface {
	Len() int64
}

type entry struct {
	key   string
	value Value
}

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64
	nBytes   int64
	l        *list.List
	cache    map[string]*list.Element
	// Optional and executed when an entry is purged.
	OnEvicted func(string, Value)
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		l:         list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	return
}

func (c *Cache) Add(key string, value Value) {

}

func (c *Cache) Len() int {
	return c.l.Len()
}
