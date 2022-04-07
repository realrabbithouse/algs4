package str

type RabinKarp struct {
	pat     string
	patHash uint64
	_M      int    // pat length
	_Q      uint64 // a big prime number
	_R      int    // base
	_RM     uint64 // R^(M - 1) % Q
}
