package basic

func PrimeFactor(n uint64) (primes []uint64) {
	var i uint64
	for i = 2; i*i <= n; i++ {
		for n%i == 0 {
			primes = append(primes, i)
			n /= i
		}
	}
	if n > 1 {
		primes = append(primes, n)
	}
	return
}
