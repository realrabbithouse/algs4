package sort

type Comparable interface {
	Swap(i, j int)
	Compare(i, j int) bool
	Length() int
	IsSorted() bool
}

type IntSlice []int

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Compare(i, j int) bool {
	return s[i] <= s[j]
}
func (s IntSlice) Length() int {
	return len(s)
}
func (s IntSlice) IsSorted() bool {
	var n = s.Length()
	for i := 0; i < n-1; i++ {
		if !s.Compare(i, i+1) {
			return false
		}
	}
	return true
}
