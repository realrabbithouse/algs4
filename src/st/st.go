package st

import (
	"algs4/src/typ"
)

type (
	Key   typ.Comparable
	Value interface{}
)

type ST interface {
	Contains(k Key) bool
	IsEmpty() bool
	Size() int
	Put(k Key, v Value)
	Get(k Key) Value
	Delete(k Key)
	Min() Key
	Max() Key
	Floor(k Key) Key   // 小于等于key的最大键
	Ceiling(k Key) Key // 大于等于key的最小键
	Rank(k Key) int    // 小于key的键的数量
	Choose(k int) Key  // 排名为k的键，k = 0,1,2...
	DeleteMin()
	DeleteMax()
}
