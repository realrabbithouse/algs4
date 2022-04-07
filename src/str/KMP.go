package str

// KMP Knuth-Morris-Pratt
// 提前判断如何重新开始查找，这种判断只取决于模式本身
// Deterministic Finite Automaton
type KMP struct {
	r   int
	m   int
	dfa [][]int // 状态机
}

func initDFA(R, N int) [][]int {
	dfa := make([][]int, R)
	for i := 0; i < R; i++ {
		dfa[i] = make([]int, N)
	}
	return dfa
}

func NewKMP(pat string) *KMP {
	R := 256
	m := len(pat)
	dfa := initDFA(R, m)
	x := 0

	// 1) If in state j and next char matches (such that: c == pat.charAt(j)), go to j+1.
	// 2) If in state j and next char doesn't match (such that: c != pat.charAt(j)), then
	//    the last j-1 characters of input are pat[1...j-1], followed by c.
	//    To compute dfa[c][j]: Simulate pat[1...j-1] (state x) on DFA and take transition c.
	//    Mismatch transition: dfa[c][j] = dfa[c][x]
	dfa[pat[0]][0] = 1 // Match succeed at j=0.
	for i := 1; i < m; i++ {
		for c := 0; c < R; c++ {
			dfa[c][i] = dfa[c][x] // Copy mismatch cases.
		}
		dfa[pat[i]][i] = i + 1 // Target hit.
		x = dfa[pat[i]][x]     // Update restart state.
		// i.e. for pattern 'A B A B A C'
		// Compute mismatch dfa in state 1 <=> Compute state transition in state X where X = simulation of empty string
		// Compute mismatch dfa in state 2 <=> Compute state transition in state X where X = simulation of 'B'
		// Compute mismatch dfa in state 3 <=> Compute state transition in state X where X = simulation of 'B A'
		// Compute mismatch dfa in state 4 <=> Compute state transition in state X where X = simulation of 'B A B'
		// Compute mismatch dfa in state 5 <=> Compute state transition in state X where X = simulation of 'B A B A'
	}

	return &KMP{
		r:   R,
		m:   m,
		dfa: dfa,
	}
}

func (kmp KMP) Search(txt string) int {
	n := len(txt)
	i := 0
	j := 0
	for i = 0; i < n && j < kmp.m; i++ {
		j = kmp.dfa[txt[i]][j] // reading a character txt[i] under state j trans to a new state
	}
	if j == kmp.m {
		return i - j
	}
	return n
}
