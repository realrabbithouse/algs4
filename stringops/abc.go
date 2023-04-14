package stringops

import (
	"math/rand"
	"time"
)

func GenRandomStrings(length, stringLength int) []string {
	rand.Seed(time.Now().UnixNano())
	ret := make([]string, length)
	for i := 0; i < length; i++ {
		tmp := make([]int, rand.Intn(stringLength)+1)
		for i := range tmp {
			tmp[i] = rand.Intn(26)
		}
		ret[i] = LOWERCASE.ToChars(tmp)
	}
	return ret
}

func BruteForceSearch(text, pat string) int {
	var (
		i, j int
		N    = len(text)
		M    = len(pat)
	)
	for i = 0; i < N && j < M; i++ {
		if text[i] == pat[j] {
			j++
		} else {
			i -= j // 显式回退（explicit backup）
			j = 0
		}
	}
	if j == M {
		return i - M
	} else {
		return N
	}
}
