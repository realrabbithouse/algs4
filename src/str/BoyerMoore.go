package str

// BoyerMoore refers to the Boyer-Moore substring search algorithm.
// * Scan characters in pattern from right to left.
// * Can skip as many as M text chars when finding one not in the pattern.
type BoyerMoore struct {
	right []int // i.e. for pattern "abcba", right[a]=4, right[b]=3, right[c]=2.
	pat   string
}

func NewBoyerMoore(pat string) *BoyerMoore {
	right := make([]int, r_)
	for i := 0; i < r_; i++ {
		right[i] = -1
	}
	M := len(pat)
	for i := 0; i < r_; i++ {
		for j := 0; j < M; j++ {
			right[pat[j]] = j
		}
	}
	return &BoyerMoore{
		right: right,
		pat:   pat,
	}
}

func (bm *BoyerMoore) Search(txt string) int {
	N := len(txt)
	M := len(bm.pat)
	var skip int
	for i := 0; i <= N-M; i += skip {
		skip = 0
		for j := M - 1; j >= 0; j-- {
			if bm.pat[j] != txt[i+j] {
				skip = j - bm.right[txt[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return N
}
