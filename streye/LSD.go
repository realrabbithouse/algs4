package streye

//
// 思想：索引计数法，对单个“位”排序，这个“位”可以是数字的某一位或字符串的某一位
// 1. 统计频率；
// 2. 累加计算索引；
// 3. 分类；
// 4. 写回。
//

// LSD examines the characters in the keys in a right-to-left order.
// LSD Least-Significant-Digit string sort.
// Every string has at least W digits.
func LSD(a []string, W int) {
	var (
		N   = len(a)
		R   = 256
		aux = make([]string, N)
	)

	for d := W - 1; d >= 0; d-- {
		count := make([]int, R+1)
		for i := 0; i < N; i++ { // Compute frequency counts base on d-th digit.
			count[a[i][d]+1]++
		}
		for r := 0; r < R; r++ { // Compute cumulates.
			count[r+1] += count[r]
		}
		for i := 0; i < N; i++ { // Classify all elements base on the d-th digit.
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}
		for i := 0; i < N; i++ { // Write back.
			a[i] = aux[i]
		}
	}
}
