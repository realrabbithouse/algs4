package st

import "algs4/src/sort"

type Key sort.Comparable
type Value interface{}

type ST interface {
	Contains(key Key) bool
	IsEmpty() bool
	Size() int
	Put(key Key, value Value)
	Get(key Key) Value
	Delete(key Key)
	Min() Key
	Max() Key
	Floor(key Key) Key   // 小于等于key的最大键
	Ceiling(key Key) Key // 大于等于key的最小键
	Rank(key Key) int    // 小于key的键的数量
	Choose(k int) Key    // 排名为k的键
	DeleteMin()
	DeleteMax()
}
