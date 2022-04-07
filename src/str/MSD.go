package str

// MSD  Most-Significant-Digit string sort.
type MSD struct {
	r      int // 基数，256
	cutoff int // 小数组切换阈值，小数组采用插入排序
	aux    []string
}

func defaultMSD() MSD {
	return MSD{
		r:      256,
		cutoff: 15,
	}
}

// charAt @return -1 ~ 255
func charAt(s string, d int) int {
	if d < len(s) {
		return int(s[d])
	}
	return -1
}

func swap(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func less(a, b string, d int) bool {
	minLen := len(a)
	if len(b) < len(a) {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if charAt(a, d) < charAt(b, d) {
			return true
		}
		if charAt(a, d) > charAt(b, d) {
			return false
		}
	}
	return len(a) < len(b)
}

func insertion(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(a[j], a[j-1], d); j-- {
			swap(a, j, j-1)
		}
	}
}

func (msd *MSD) recurSort(a []string, lo, hi, d int) {
	if hi <= lo+msd.cutoff {
		insertion(a, lo, hi, d)
		return
	}
	count := make([]int, msd.r+2)
	for i := lo; i <= hi; i++ { // 计算频率
		count[charAt(a[i], d)+2]++
	}
	for r := 0; r < msd.r+1; r++ { // 将频率转换为索引
		count[r+1] += count[r]
	}
	for i := lo; i <= hi; i++ { // 数据分类
		c := charAt(a[i], d)
		msd.aux[count[c+1]] = a[i]
		count[c+1]++
	}
	for i := lo; i <= hi; i++ { // 写回
		a[i] = msd.aux[i-lo]
	}
	// 递归的以每个字符为键进行排序
	for r := 0; r < msd.r; r++ {
		msd.recurSort(a, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

func (msd *MSD) sort(a []string) {
	msd.aux = make([]string, len(a))
	msd.recurSort(a, 0, len(a)-1, 0)
}

func MSDSort(a []string) {
	msd := defaultMSD()
	msd.sort(a)
}
