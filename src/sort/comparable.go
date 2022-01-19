package sort

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
func (s IntSlice) Copy() ComparableSlice {
	cp := make([]int, s.Len())
	for i := range s {
		cp[i] = s[i]
	}
	return IntSlice(cp)
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
func (s StringSlice) New(sz int) ComparableSlice {
	return StringSlice(make([]string, sz))
}
func (s StringSlice) Copy() ComparableSlice {
	cp := make([]string, s.Len())
	for i := range s {
		cp[i] = s[i]
	}
	return StringSlice(cp)
}
