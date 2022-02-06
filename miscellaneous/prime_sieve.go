package miscellaneous

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// print the first n primes (n > 0)
func SievePrimeConcurrent(n int) {
	if n < 1 {
		return
	}
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < n; i++ {
		prime := <-ch
		print(prime, ",")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

func SievePrimeNormal(n int) {
	if n < 1 {
		return
	}

	var primes []int
	primes = append(primes, 2)
	counter := 1
	for i := 3; ; i++ {
		if counter >= n {
			break
		}
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			counter++
		}
	}
	fmt.Println(primes)
}
