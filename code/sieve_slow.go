package main

const n = 100000

func main() {
	primes := make([]int, 0)
	discovered := 0
OUTER:
	for i := 2; discovered < n; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue OUTER
			}
		}
		primes = append(primes, i)
		discovered++
	}
	// fmt.Println(primes)
}
