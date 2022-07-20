package streye

import (
	"algs4/graph"
)

// The NFA provides a testdata type for creating a nondeterministic finite state automaton (NFA)
// from a regular expression and testing whether a given string is matched by that regular
// expression. It supports the following operations: concatenation, closure, binary or,
// and parentheses. It does not support multiway or, character classes, metacharacters
// (either in the text or pattern), capturing capabilities, greedy or reluctant modifiers,
// and other features in industrial-strength implementations such as java.util.regex.Pattern
// and java.util.regex.Matcher.
//
// This implementation builds the NFA using a digraph and a stack and simulates the NFA
// using digraph search. The constructor takes time proportional to m, where m is the number
// of characters in the regular expression. The recognizes method takes time proportional to m*n,
// where n is the number of characters in the text.
type NFA struct {
	dfs    *graph.Digraph // digraph of epsilon transitions
	regexp string         // regular expression
	m      int            // number of characters in regular expression
}

func MakeNFA(regexp string) (*NFA, error) {
	return nil, nil
}

func (nfa *NFA) Recognize(txt string) (tf bool, matchFrag string, matchIndex int) {
	return
}
