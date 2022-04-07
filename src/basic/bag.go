package basic

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type linked struct {
	item interface{}
	next *linked
}

type Bag struct {
	first *linked // first item
	iter  *linked // iter 用于遍历元素，每次遍历完毕都将重置
	n     int
}

func (b Bag) IsEmpty() bool {
	return b.first == nil
}

func (b *Bag) Add(item interface{}) {
	orig := b.first
	b.first = &linked{
		item: item,
		next: orig,
	}
	b.iter = b.first
	b.n++
}

func (b Bag) Size() int {
	return b.n
}

func (b *Bag) HasNext() bool {
	if b.iter != nil {
		return true
	}
	b.iter = b.first
	return false
}

func (b *Bag) Next() (ret interface{}) {
	ret = b.iter.item
	b.iter = b.iter.next
	return
}
