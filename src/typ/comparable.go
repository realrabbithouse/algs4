package typ

// ComparableSlice is the interface that wraps the comparable slice objects for sorting.
type ComparableSlice interface {
	Swap(i, j int)
	Compare(i, j int) bool
	Len() int
	IsSorted() bool
}

type IntSlice []int

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Compare(i, j int) bool {
	return s[i] <= s[j]
}
func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) IsSorted() bool {
	var n = s.Len()
	for i := 0; i < n-1; i++ {
		if !s.Compare(i, i+1) {
			return false
		}
	}
	return true
}
func (s IntSlice) New(sz int) ComparableSlice {
	return IntSlice(make([]int, sz))
}

type StringSlice []string

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringSlice) Compare(i, j int) bool {
	return s[i] <= s[j]
}
func (s StringSlice) Len() int {
	return len(s)
}
func (s StringSlice) IsSorted() bool {
	var n = s.Len()
	for i := 0; i < n-1; i++ {
		if !s.Compare(i, i+1) {
			return false
		}
	}
	return true
}

// ******************************************************************************************* //

// Comparable is the interface that wraps the compare method for comparable objects.
type Comparable interface {
	CompareTo(Comparable) int
}

func Equal(left, right Comparable) bool {
	return left.CompareTo(right) == 0
}

func Less(left, right Comparable) bool {
	return left.CompareTo(right) < 0
}

func LessEqual(left, right Comparable) bool {
	res := left.CompareTo(right)
	return res <= 0
}

func Greater(left, right Comparable) bool {
	return left.CompareTo(right) > 0
}

func GreaterEqual(left, right Comparable) bool {
	res := left.CompareTo(right)
	return res >= 0
}

type ComparableInt int

func (c ComparableInt) CompareTo(obj Comparable) int {
	if c < obj.(ComparableInt) {
		return -1
	} else if c > obj.(ComparableInt) {
		return 1
	} else {
		return 0
	}
}

type ComparableString string

func (c ComparableString) CompareTo(obj Comparable) int {
	if c < obj.(ComparableString) {
		return -1
	} else if c > obj.(ComparableString) {
		return 1
	} else {
		return 0
	}
}

type ComparableFloat64 float64

func (c ComparableFloat64) CompareTo(obj Comparable) int {
	if c < obj.(ComparableFloat64) {
		return -1
	} else if c > obj.(ComparableFloat64) {
		return 1
	} else {
		return 0
	}
}
