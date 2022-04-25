package classic

import (
	"algs4/src/basic"
	"fmt"
)

//
// In order to apply eightQueenBT to a specific class of problems, one must provide the data P for
// the particular instance of the problem that is to be solved, and six procedural parameters, root,
// reject, accept, first, next, and output. These procedures should take the instance data P as a
// parameter and should do the following:
//
// root(P): return the partial candidate at the root of the search tree.
// reject(P,c): return true only if the partial candidate c is not worth completing.
// accept(P,c): return true if c is a solution of P, and false otherwise.
// first(P,c): generate the first extension of candidate c.
// next(P,s): generate the next alternative extension of a candidate, after the extension s.
// output(P,c): use the solution c of P, as appropriate to the application.
//
// The eightQueenBT algorithm reduces the problem to the call backtrack(root(P)), where backtrack is
// the following recursive procedure:
//
// procedure backtrack(c) is
//    if reject(P, c) then return
//    if accept(P, c) then output(P, c)
//    s ← first(P, c)
//    while s ≠ NULL do
//        backtrack(s)
//        s ← next(P, s)
//
// 1. Decision Problem – In this, we search for a feasible solution.
// 2. Optimization Problem – In this, we search for the best solution.
// 3. Enumeration Problem – In this, we find all feasible solutions.

const _nQueen = 8

func accept(pos []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if col == pos[i] ||
			col-pos[i] == row-i ||
			col-pos[i] == i-row {
			return false
		}
	}
	return true
}

func eightQueenBT(pos []int, row int, res *basic.Queue) {
	if row == _nQueen {
		chess := make([]int, _nQueen)
		copy(chess, pos)
		res.Enqueue(chess)
		return
	}
	for col := 0; col < _nQueen; col++ {
		if accept(pos, row, col) {
			pos[row] = col
			eightQueenBT(pos, row+1, res)
		}
	}
}

func EightQueen() *basic.Queue {
	pos := make([]int, _nQueen)
	res := new(basic.Queue)
	eightQueenBT(pos, 0, res)
	return res
}

// ****************************************************** //

// Backtrack(x)
// if x is not a solution
//	 return false
// if x is a new solution
//	 add to list of solutions
// backtrack(expand x)

func intSliceHas(c []int, v int) bool {
	for _, e := range c {
		if e == v {
			return true
		}
	}
	return false
}

func permutationBT(c []int, a []int, n int) {
	if len(c) == len(a) {
		fmt.Println(c)
	}
	for i := 0; i < n; i++ {
		v := a[i]
		if !intSliceHas(c, v) {
			cc := make([]int, len(c), n)
			copy(cc, c)
			cc = append(cc, v)
			permutationBT(cc, a, n)
		}
	}
}

func Permutation(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		c := make([]int, 0, n)
		c = append(c, a[i])
		permutationBT(c, a, n)
	}
}
