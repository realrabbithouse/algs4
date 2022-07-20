package dp

import (
	"fmt"
	"testing"
)

func TestKnapsack(t *testing.T) {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	n := 3
	W := 50
	fmt.Println("Maximum value is", Knapsack(wt, val, n, W))
}

func TestUnboundedKnapsack(t *testing.T) {
	val := []int{10, 30, 20}
	wt := []int{5, 10, 15}
	n := 3
	W := 100
	fmt.Println("Maximum value is", UnboundedKnapsack(wt, val, n, W))
}

func TestLongestCommonSubstring(t *testing.T) {
	X := "OldSite:GeeksforGeeks.org"
	Y := "NewSite:GeeksQuiz.com"
	fmt.Println("Length of Longest Common Substring is", LongestCommonSubstring(X, Y, len(X), len(Y)))
}

func TestRob(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(
		recurRob(nums),
		recurMemoRob(nums),
		iterMemoRob(nums),
		iterNRob(nums),
	)
}
