package leetcode

import (
	"fmt"
	"math"
	"sort"
)

// 785. Is Graph Bipartite?

type Color bool

type Bipartite struct {
	graph       [][]int
	visited     []bool
	colors      []Color
	isBipartite bool
}

func (b *Bipartite) dfs(u int) {
	b.visited[u] = true
	for _, v := range b.graph[u] {
		if !b.visited[v] {
			b.colors[v] = !b.colors[u]
			b.dfs(v)
		} else if b.colors[v] == b.colors[u] {
			b.isBipartite = false
			break
		}
	}
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	b := Bipartite{
		graph:       graph,
		visited:     make([]bool, n),
		colors:      make([]Color, n),
		isBipartite: true,
	}
	for i := 0; i < n; i++ {
		if !b.visited[i] && b.isBipartite {
			b.dfs(i)
		}
	}
	return b.isBipartite
}

// primeFactors 质因子分解
func primeFactors(x int) {
	for x%2 == 0 {
		fmt.Print(2, " ")
		x /= 2
	}
	rx := int(math.Floor(math.Sqrt(float64(x))))
	for i := 3; i <= rx; i++ {
		for x%i == 0 {
			fmt.Print(i, " ")
			x /= i
		}
	}
	if x > 1 {
		fmt.Print(x, " ")
	}
	fmt.Print("\n")
}

// Dynamic Programming
// longestIncSeq The Longest Increasing Subsequence (LIS):
// given x as index, f(x) as length of LIS at x, then
// f(x) = max{f(p)} + 1, for p < x and a_p < a_x.
func longestIncSeq(a []int) int {
	var (
		n    = len(a)
		f    = make([]int, n)
		fx   = 1
		fmax = 1
	)
	for i := 0; i < n; i++ {
		f[i] = 1
	}
	for i := 1; i < n; i++ {
		fx = 1
		for j := i - 1; j >= 0; j-- {
			if a[i] > a[j] {
				if f[j]+1 > fx {
					fx = f[j] + 1
				}
			}
		}
		f[i] = fx
		if fx > fmax {
			fmax = fx
		}
	}
	return fmax
}

func maxInt(a ...int) int {
	n := len(a)
	max := 0
	for i := 0; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func maxSum(a []int, lo, hi int) int {
	if lo == hi {
		return a[lo]
	}
	mid := lo + (hi-lo)/2
	l := maxSum(a, lo, mid)
	r := maxSum(a, mid+1, hi)
	var p, q int
	ml := a[mid]
	mr := a[mid+1]
	for i := mid; i >= lo; i-- {
		p += a[i]
		if p > ml {
			ml = p
		}
	}
	for i := mid + 1; i <= hi; i++ {
		q += a[i]
		if q > mr {
			mr = q
		}
	}
	return maxInt(l, r, ml+mr)
}

func hanoiTower(height int) int {
	if height < 1 {
		return 0
	}
	if height == 1 {
		return 1
	}
	return 2*hanoiTower(height-1) + 1
}

// hanoiTowerMove 汉诺塔
func hanoiTowerMove(height int, a, b, c string) int {
	// 将 height 的东西按序从 a 移动到 c
	if height < 1 {
		return 0
	}
	if height == 1 {
		fmt.Println(a, "->", c)
		return 1
	}
	// 将 height-1 的东西按序从 a 移动到 b
	i := hanoiTowerMove(height-1, a, c, b)
	fmt.Println(a, "->", c) // 将 height 该层的东西移动到 c
	// 将 height-1 的东西按序从 b 移动到 c
	j := hanoiTowerMove(height-1, b, a, c)
	return i + j + 1
}

// lengthOfLongestSubstring leetcode 3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	var left, right, res int
	chars := make([]int, 256)
	for right < length {
		r := s[right]
		chars[r]++
		for chars[r] > 1 {
			l := s[left]
			chars[l]--
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
	}
	return res
}

// groupAnagrams leetcode 49. Group Anagrams
func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	cp := make([]string, length)
	for i := 0; i < length; i++ {
		b := []byte(strs[i])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		cp[i] = string(b)
	}
	m := make(map[string][]int)
	for i := 0; i < length; i++ {
		m[cp[i]] = append(m[cp[i]], i)
	}
	var res [][]string
	for _, indices := range m {
		anagrams := make([]string, len(indices))
		for i, index := range indices {
			anagrams[i] = strs[index]
		}
		res = append(res, anagrams)
	}
	return res
}

// findSubstring leetcode 30. Substring with Concatenation of All Words
func findSubstring(s string, words []string) (ret []int) {
	n := len(s)
	wordNum := len(words)
	wordLength := len(words[0])
	substringLength := wordNum * wordLength
	wordFreq := make(map[string]int, wordNum)

	for i := range words {
		wordFreq[words[i]]++
	}

	for i := 0; i < wordLength; i++ {
		answer := slidingWindow(i, s, wordFreq, n, wordNum, wordLength, substringLength)
		ret = append(ret, answer...)
	}

	return ret
}

func slidingWindow(left int, s string, wordFreq map[string]int, n, wordNum, wordLength, substringLength int) []int {
	var (
		wordFound = make(map[string]int)
		answer    []int
		wordUsed  int
		excessive bool
	)

	for right := left; right <= n-wordLength; right += wordLength {
		sub := s[right : right+wordLength]

		if _, ok := wordFreq[sub]; !ok {
			wordFound = make(map[string]int)
			wordUsed = 0
			excessive = false
			left = right + wordLength // reset left pointer
		} else {
			// not match, reset left pointer
			for right-left == substringLength || excessive {
				leftMostWord := s[left : left+wordLength]
				left += wordLength
				wordFound[leftMostWord]--
				if wordFound[leftMostWord] >= wordFreq[leftMostWord] {
					// not match due to excessive
					excessive = false
				} else {
					wordUsed--
				}
			}

			wordFound[sub]++
			if wordFound[sub] <= wordFreq[sub] {
				wordUsed++
			} else {
				excessive = true
			}

			if wordUsed == wordNum && !excessive {
				// found a valid substring
				answer = append(answer, left)
			}
		}
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// threeSumClosest leetcode 16. 3Sum Closest
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	var (
		j, k int
		sum  int
		rel  = math.MaxInt
	)
	for i := 0; i < length-2; i++ {
		j = i + 1
		k = length - 1
		for j < k {
			r := nums[i] + nums[j] + nums[k] - target
			if rel > abs(r) {
				rel = abs(r)
				sum = nums[i] + nums[j] + nums[k]
			}
			if r > 0 {
				k--
			} else if r < 0 {
				j++
			} else {
				return sum
			}
		}
	}
	return sum
}

// leetcode 86. Partition List
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	return nil
}

// MyCalendarThree leetcode 732. My Calendar III
type MyCalendarThree struct {
	timeline map[int]int
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.timeline[start]++
	this.timeline[end]--
	keys := make([]int, 0, len(this.timeline))
	for k := range this.timeline {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var ongoing, max int
	for i := range keys {
		ongoing += this.timeline[keys[i]]
		if ongoing > max {
			max = ongoing
		}
	}
	return max
}

// MyCircularDeque leetcode 641. Design Circular Deque
type MyCircularDeque struct {
	N     int
	Cap   int
	data  []int
	first int
	last  int
}

//func Constructor(k int) MyCircularDeque {
//	return MyCircularDeque{
//		Cap:   k,
//		testdata:  make([]int, k),
//		first: -1,
//		last:  -1,
//	}
//}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.first = 0
		this.last = 0
		this.data[this.first] = value
	} else {
		this.first--
		if this.first < 0 {
			this.first = this.Cap - 1
		}
		this.data[this.first] = value
	}
	this.N++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.last = 0
		this.first = 0
		this.data[this.last] = value
	} else {
		this.last = (this.last + 1) % this.Cap
		this.data[this.last] = value
	}
	this.N++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.N--
	if this.IsEmpty() {
		this.first = -1
		this.last = -1
	} else {
		this.first = (this.first + 1) % this.Cap
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.N--
	if this.IsEmpty() {
		this.last = -1
		this.first = -1
	} else {
		this.last--
		if this.last < 0 {
			this.last = this.Cap - 1
		}
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.first]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.last]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.N == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.N == this.Cap
}

func validateStackSequences(pushed []int, popped []int) bool {
	var (
		n         = len(pushed)
		iNextPush = 1
		iNextPop  = 0
		iTop      = -1
		done      = make([]bool, n)
	)
	for iNextPush < n || iNextPop < n {
		if iTop != -1 && pushed[iTop] == popped[iNextPop] {
			iNextPop++
			done[iTop] = true
			for i := iTop - 1; i >= 0; i-- {
				if !done[i] {
					iTop = i
					break
				}
			}
		} else if iNextPush < n {
			iTop = iNextPush
			iNextPush++
		} else if iNextPop < n {
			return false
		}
	}
	return true
}

// leetcode 105. Construct Binary Tree from Preorder and Inorder Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	r, lpre, rpre, lin, rin := split(preorder, inorder)
	t := new(TreeNode)
	t.Val = r
	t.Left = buildTree(lpre, lin)
	t.Right = buildTree(rpre, rin)
	return t
}

func index(v int, slice []int) int {
	for i, val := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func split(pre, in []int) (root int, lpre, rpre, lin, rin []int) {
	root = pre[0]
	mid := index(root, in)
	lin = in[:mid]
	rin = in[mid+1:]
	lpre = pre[1 : len(lin)+1]
	rpre = pre[len(lin)+1:]
	return
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	n := len(people)
	ret := make([][]int, n)
	for i, person := range people {
		for j := i; j > person[1]; j-- {
			ret[j] = ret[j-1]
		}
		ret[person[1]] = person
	}
	return ret
}

// leetcode 116. Populating Next Right Pointers in Each Node

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	// 递归
	//root.Left = connect(root.Left)
	//root.Right = connect(root.Right)
	//if root.Left != nil {
	//	root.Left.Next = root.Right
	//	l, r := root.Left, root.Right
	//	for l.Right != nil {
	//		l.Right.Next = r.Left
	//		l, r = l.Right, r.Left
	//	}
	//}

	// 非递归
	head := root
	var pre *Node
	for head.Left != nil {
		pre = head.Left
		pre.Next = head.Right
		for pre.Next != nil {
			pre.Right.Next = pre.Next.Left
			pre = pre.Next
		}
		head = head.Left
	}

	return root
}

// leetcode 117. Populating Next Right Pointers in Each Node

func connectNormal(root *Node) *Node {
	/*var head *Node // left most node on the next level
	var prev *Node // leading node on the next level
	cur := root    // current level node
	for cur != nil {
		for cur != nil {
			if cur.Left != nil {
				if prev != nil {
					prev.Next = cur.Left
				} else {
					head = cur.Left
				}
				prev = cur.Left
			}
			if cur.Right != nil {
				if prev != nil {
					prev.Next = cur.Right
				} else {
					head = cur.Right
				}
				prev = cur.Right
			}
			cur = cur.Next
		}
		cur = head
		head = nil
		prev = nil
	}*/

	// level-order traversal
	r := root
	for root != nil {
		tmpChild := new(Node)
		curChild := tmpChild
		for root != nil {
			if root.Left != nil {
				curChild.Next = root.Left
				curChild = root.Left
			}
			if root.Right != nil {
				curChild.Next = root.Right
				curChild = root.Right
			}
			root = root.Next
		}
		root = tmpChild.Next
	}

	return r
}

func binarySearchInt(slice []int, v int) int {
	lo, hi := 0, len(slice)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if slice[mid] > v {
			hi = mid - 1
		} else if slice[mid] < v {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func leastBricks(wall [][]int) int {
	N := len(wall)
	internalEdge := make(map[int]struct{})
	for i := range wall {
		n := len(wall[i])
		for j := 0; j < n-1; j++ {
			wall[i][j+1] += wall[i][j]
		}
		for j := 0; j < n-1; j++ {
			internalEdge[wall[i][j]] = struct{}{}
		}
	}
	min := N
	for v := range internalEdge {
		var nBrick int
		for i := 0; i < N; i++ {
			if binarySearchInt(wall[i], v) == -1 {
				nBrick++
			}
		}
		if nBrick < min {
			min = nBrick
		}
	}
	return min
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt, math.MaxInt)
}

func check(n *TreeNode, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Val <= min || n.Val >= max {
		return false
	}
	return check(n.Left, min, n.Val) && check(n.Right, n.Val, max)
}

// shortestPathBinaryMatrix leetcode 1091. Shortest Path in Binary Matrix
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == -1 || grid[n-1][n-1] == -1 {
		return -1
	}
	q := make([][3]int, 0, n)
	q = append(q, [3]int{0, 0, 1})
	for len(q) > 0 {
		i, j, d := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		if i == n-1 && j == n-1 {
			return d
		}
		elems := [][2]int{{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1}, {i, j - 1},
			{i, j + 1}, {i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}}
		for _, elem := range elems {
			if elem[0] >= 0 && elem[0] < n && elem[1] >= 0 && elem[1] < n && grid[elem[0]][elem[1]] != 1 {
				q = append(q, [3]int{elem[0], elem[1], d + 1})
				grid[elem[0]][elem[1]] = 1
			}
		}
	}
	return -1
}

// shortestAlternatingPaths 1129. Shortest Path with Alternating Colors
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	g := make([][][2]int, n) // add: <index, color>
	for _, edge := range redEdges {
		// red = 0
		g[edge[0]] = append(g[edge[0]], [2]int{edge[1], 0})
	}
	for _, edge := range blueEdges {
		// blue = 1
		g[edge[0]] = append(g[edge[0]], [2]int{edge[1], 1})
	}
	queue := [][2]int{
		{0, 0},
		{0, 1},
	} // <index, color>
	cost := make([][2]int, n) // <red cost, blue cost>
	for i := 0; i < n; i++ {
		cost[i] = [2]int{-1, -1}
	}
	cost[0] = [2]int{0, 0}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range g[u[0]] {
			if cost[v[0]][v[1]] != -1 || u[1] == v[1] {
				continue
			}
			cost[v[0]][v[1]] = 1 + cost[u[0]][u[1]]
			queue = append(queue, v)
		}
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		rx, bx := cost[i][0], cost[i][1]
		if rx < bx {
			if rx == -1 {
				ret[i] = bx
			} else {
				ret[i] = rx
			}
		} else {
			if bx == -1 {
				ret[i] = rx
			} else {
				ret[i] = bx
			}
		}
	}
	return ret
}

// levelOrder 102. Binary Tree Level Order Traversal
func levelOrder(root *TreeNode) [][]int {
	var (
		queue []*TreeNode
		ret   [][]int
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		var newQueue []*TreeNode
		for i := range queue {
			if queue[i] != nil {
				level = append(level, queue[i].Val)
			}
		}
		for i := range queue {
			if queue[i] != nil {
				newQueue = append(newQueue, queue[i].Left, queue[i].Right)
			}
		}
		if level != nil {
			ret = append(ret, level)
		}
		queue = newQueue
	}
	return ret
}

func next(nums, sum []int, target int) int {
	sz := len(nums)
	if sz == 1 {
		if nums[0] == target || nums[0] == -target {
			if target == 0 {
				return 2
			}
			return 1
		}
		return 0
	} else {
		max := sum[len(sum)-sz]
		if target > max || target < -max {
			return 0
		} else {
			return next(nums[1:], sum, target-nums[0]) + next(nums[1:], sum, target+nums[0])
		}
	}
}

// findTargetSumWays 494. Target Sum
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := make([]int, n)
	for i := n - 1; i >= 1; i-- {
		sum[i-1] += sum[i]
	}
	return next(nums, sum, target)
}

// canPartition 416. Partition Equal Subset Sum
func canPartition(nums []int) bool {
	// DP
	return false
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	r := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(i, visited, isConnected)
		}
		r++
	}
	return r
}

func bfs(s int, visited []bool, isConnected [][]int) {
	n := len(isConnected)
	q := make([]int, 0, n)
	q = append(q, s)
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		visited[u] = true
		for v, ok := range isConnected[u] {
			if !visited[v] && v != u && ok == 1 {
				q = append(q, v)
			}
		}
	}
}

// closedIsland 1254. Number of Closed Islands
func closedIsland(grid [][]int) int {
	var (
		M               = len(grid)
		N               = len(grid[0])
		visited         = make([][]bool, M)
		nInternalIsland int
		isInternal      = true
		queue           [][2]int
	)
	for i := 0; i < M; i++ {
		visited[i] = make([]bool, N)
	}
	for ix := 0; ix < M; ix++ {
		for jx := 0; jx < N; jx++ {
			if grid[ix][jx] == 0 && !visited[ix][jx] {
				queue = append(queue, [2]int{ix, jx})
				for len(queue) > 0 {
					n := queue[0]
					queue = queue[1:]
					i, j := n[0], n[1]
					visited[i][j] = true
					if isInternal && i == 0 || i == M-1 ||
						j == 0 || j == N-1 {
						isInternal = false
					}
					elems := [][2]int{
						{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
					}
					for _, elem := range elems {
						if elem[0] >= 0 && elem[0] < M && elem[1] >= 0 && elem[1] < N &&
							!visited[elem[0]][elem[1]] && grid[elem[0]][elem[1]] == 0 {
							queue = append(queue, elem)
						}
					}
				}
				if isInternal {
					nInternalIsland++
				}
				isInternal = true
			}
		}
	}
	return nInternalIsland
}

// openLock 752. Open the Lock
func openLock(deadends []string, target string) int {
	visitedMap := make(map[string]struct{})
	for i := range deadends {
		visitedMap[deadends[i]] = struct{}{}
	}
	var queue []string
	var dist int
	if _, ok := visitedMap["0000"]; !ok {
		queue = append(queue, "0000")
		visitedMap["0000"] = struct{}{}
	}
	for len(queue) > 0 {
		var newQueue []string
		for _, u := range queue {
			if u == target {
				return dist
			}
			adjs := adj(u)
			for _, adj := range adjs {
				if _, ok := visitedMap[adj]; !ok {
					visitedMap[adj] = struct{}{}
					newQueue = append(newQueue, adj)
				}
			}
		}
		queue = newQueue
		dist++
	}
	return -1
}

func adj(s string) []string {
	ret := make([]string, 0, 8)
	bs := []byte(s)
	for i, b := range bs {
		bplus := make([]byte, 4)
		bdec := make([]byte, 4)
		copy(bplus, bs)
		copy(bdec, bs)
		bplus[i] = (b-47)%10 + 48
		bdec[i] = (b-39)%10 + 48
		ret = append(ret, string(bplus), string(bdec))
	}
	return ret
}

// reverseKGroup 25. Reverse Nodes in k-Group
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		size                      int
		prevNewTail, prevNextHead *ListNode
		ret                       *ListNode
	)
	prevNextHead = head
	for it := head; it != nil; it = it.Next {
		size++
	}
	for i := 0; i < size-k+1; i += k {
		newHead, newTail, nextHead := reverseK(prevNextHead, k)
		if i == 0 {
			ret = newHead
		}
		if prevNewTail != nil {
			prevNewTail.Next = newHead
		}
		prevNewTail = newTail
		prevNextHead = nextHead
	}
	prevNewTail.Next = prevNextHead
	return ret
}

func reverseK(head *ListNode, k int) (newHead, newTail, nextHead *ListNode) {
	newTail = head
	var prev, current, next *ListNode
	current = head
	for i := 0; i < k; i++ {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	newHead = prev
	nextHead = next
	return
}

type FreqStack struct {
	freq      map[int]int
	freqStack map[int][]int
	mf        int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:      make(map[int]int),
		freqStack: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	f := this.freq[val]
	this.freqStack[f] = append(this.freqStack[f], val)
	if f > this.mf {
		this.mf = f
	}
}

func (this *FreqStack) Pop() int {
	stack := this.freqStack[this.mf]
	n := len(stack)
	ret := stack[n-1]
	if n == 1 {
		delete(this.freqStack, this.mf)
		for f := this.mf - 1; f >= 0; f-- {
			if len(this.freqStack[f]) > 0 || f == 0 {
				this.mf = f
				break
			}
		}
	} else {
		this.freqStack[this.mf] = stack[:n-1]
	}
	this.freq[ret]--
	return ret
}

func orderlyQueue(s string, k int) string {
	if k > 1 {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	} else {
		min := s
		for i := 1; i < len(s); i++ {
			tmp := s[i:] + s[:i]
			if tmp < min {
				min = tmp
			}
		}
		return min
	}
}
