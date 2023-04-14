package stringops

// The RabinKarp finds the first occurrence of a pattern string in a text string.
// This implementation uses the Rabin-Karp algorithm.
type RabinKarp struct {
	pat     string // the pattern string
	patHash int64  // pattern hash value
	_M      int    // pat length
	_Q      int64  // a large prime, small enough to avoid long overflow
	_R      int    // radix
	_RM     int64  // R^(M-1) % Q
}
