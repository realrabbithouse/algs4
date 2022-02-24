package str

type KMP struct {
	R   int
	m   int
	dfa [][]int // 状态机
}

func makeDFA(R, N int) [][]int {
	dfa := make([][]int, R)
	for i := 0; i < R; i++ {
		dfa[i] = make([]int, N)
	}
	return dfa
}

func NewKMP(pat string) *KMP {
	R := 256
	m := len(pat)
	dfa := makeDFA(R, m)
	x := 0

	dfa[pat[0]][0] = 1 // Match succeed at j=0.
	for i := 1; i < m; i++ {
		for c := 0; c < R; c++ {
			dfa[c][i] = dfa[c][x] // Copy mismatch cases.
		}
		dfa[pat[i]][i] = i + 1 // Target hit.
		x = dfa[pat[i]][x]     // Update restart state.
	}

	return &KMP{
		R:   R,
		m:   m,
		dfa: dfa,
	}
}

func (kmp KMP) Search(txt string) int {
	n := len(txt)
	i := 0
	j := 0
	for i = 0; i < n && j < kmp.m; i++ {
		j = kmp.dfa[txt[i]][j]
	}
	if j == kmp.m {
		return i - j
	}
	return n
}
