package dp

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// ***************************************************************** //

// Knapsack dp(i, w) = max(dp(i-1, w), val[i] + dp(i-1, w-weight[i]))
func Knapsack(weight []int, val []int, n int, w int) (maxValue int) {
	if len(weight) != n || len(val) != n || w < 0 {
		return
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < w+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= weight[i] {
				tmp := val[i] + dp[i-1][j-weight[i]]
				if tmp > dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	return dp[n-1][w]
}

// UnboundedKnapsack dp(w) = max(dp(w-weight[i]) + val[i], i = 0,1,2,...)
func UnboundedKnapsack(weight []int, val []int, n int, w int) (maxValue int) {
	if len(weight) != n || len(val) != n || w < 0 {
		return
	}
	dp := make([]int, w+1)
	for i := 1; i < w+1; i++ {
		var max int
		for j := 0; j < n; j++ {
			if i >= weight[j] && dp[i-weight[j]]+val[j] > max {
				max = dp[i-weight[j]] + val[j]
			}
		}
		dp[i] = max
	}
	return dp[w]
}

// LongestCommonSubstring given two strings ‘X’ and ‘Y’, find the length of
// the longest common substring.
//
// The longest common suffix has following optimal substructure property.
// If last characters match, then we reduce both lengths by 1
// LCSuff(X, Y, m, n) = LCSuff(X, Y, m-1, n-1) + 1 if X[m-1] = Y[n-1]
// If last characters do not match, then result is 0, i.e.,
// LCSuff(X, Y, m, n) = 0 if (X[m-1] != Y[n-1])
// Now we consider suffixes of different substrings ending at different indexes.
// The maximum length Longest Common Suffix is the longest common substring.
// LCSubStr(X, Y, m, n) = Max(LCSuff(X, Y, i, j)) where 1 <= i <= m and 1 <= j <= n
func LongestCommonSubstring(X, Y string, m, n int) int {
	lcSuff := make([][]int, m+1)
	for i := 0; i < m; i++ {
		lcSuff[i] = make([]int, n+1)
	}
	var lcLen int
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if X[i-1] == Y[j-1] {
				lcSuff[i][j] = lcSuff[i-1][j-1] + 1
				if lcSuff[i][j] > lcLen {
					lcLen = lcSuff[i][j]
				}
			} else {
				lcSuff[i][j] = 0
			}
		}
	}
	return lcLen

	// Another approach: (Space optimized approach).
	// In the above approach, we are only using the last row of the 2-D array only,
	// hence we can optimize the space by using
	// a 2-D array of dimension 2*(min(n,m)).
}

// CanJump You are given an integer array nums. You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.
func CanJump(nums []int) bool {
	n := len(nums)
	var i, reach int
	for ; i < n && i <= reach; i++ {
		tmp := i + nums[i]
		if tmp > reach {
			reach = tmp
		}
	}
	return i == n
}

// Jump Your goal is to reach the last index in the minimum number of jumps.
// You can assume that you can always reach the last index.
func Jump(nums []int) int {
	var (
		n                   = len(nums)
		curEnd, curFurthest int
		nJump               int
	)
	for i := 0; i < n-1; i++ {
		if i+nums[i] > curFurthest {
			curFurthest = i + nums[i]
		}
		if i == curEnd {
			nJump++
			curEnd = curFurthest
		}
	}
	return nJump
}

// House robber to illustrate common idea.
// 1. Find recursive relation
// 2. Recursive (top-down)
// 3. Recursive + memo (top-down)
// 4. Iterative + memo (bottom-up)
// 5. Iterative + N variables (bottom-up)

// 1. Recursive relation: f(n) = max(f(n-2) + a[n], f(n-1))

// 2. Recursive (top-down)
func recurRob(nums []int) int {
	return recurRobFunc(nums, len(nums)-1)
}

func recurRobFunc(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	return maxInt(recurRobFunc(nums, i-1), recurRobFunc(nums, i-2)+nums[i])
}

// 3. Recursive + memo (top-down)
func recurMemoRob(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return recurMemoRobFunc(nums, memo, len(nums)-1)
}

func recurMemoRobFunc(nums, memo []int, i int) int {
	if i < 0 {
		return 0
	}
	if memo[i] >= 0 {
		return memo[i]
	}
	res := maxInt(recurMemoRobFunc(nums, memo, i-1), recurMemoRobFunc(nums, memo, i-2)+nums[i])
	memo[i] = res
	return res
}

// 4. Iterative + memo (bottom-up)
func iterMemoRob(nums []int) int {
	n := len(nums)
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = nums[0]
	for i := 1; i < n; i++ {
		memo[i+1] = maxInt(memo[i], memo[i-1]+nums[i])
	}
	return memo[n]
}

// Iterative + N variables (bottom-up)
func iterNRob(nums []int) int {
	// Notice that in the previous step we use only memo[i] and memo[i-1]!
	n := len(nums)
	prev0 := 0       // prev0 as memo[i-1]
	prev1 := nums[0] // prev1 as memo[i]
	for i := 1; i < n; i++ {
		tmp := maxInt(prev0+nums[i], prev1)
		prev0, prev1 = prev1, tmp
	}
	return prev1
}
