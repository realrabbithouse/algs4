package basic

type Queue struct {
	first, last *linked
	iter        *linked
	n           int
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Enqueue(val interface{}) {
	oldLast := q.last
	q.last = &linked{
		item: val,
	}
	if q.first == nil {
		q.first = q.last
		q.iter = q.first
	} else {
		oldLast.next = q.last
	}
	q.n++
}

func (q *Queue) Dequeue() interface{} {
	if q.first == nil {
		return nil
	}
	val := q.first.item
	q.first = q.first.next
	q.iter = q.first
	q.n--
	if q.first == nil {
		q.last = nil
	}
	return val
}

func (q *Queue) HasNext() bool {
	if q.iter != nil {
		return true
	}
	q.iter = q.first
	return false
}

func (q *Queue) Next() interface{} {
	val := q.iter.item
	q.iter = q.iter.next
	return val
}
