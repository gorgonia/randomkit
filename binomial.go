package randomkit

import (
	"math/rand"
	"sync"
	"time"
)

// BinomialGenerator is a random number generator built on top of math/rand. It generates a value ~ binomial(n, p)
type BinomialGenerator struct {
	sync.Mutex
	*rand.Rand
}

// Int returns X ~ binomial(n, p) as an int
//
// n = number of tries
// p = probability of acceptance
//
// if n > 1000, the process used to generate the result is concurrent.
func (g *BinomialGenerator) Int(n int, p float64) int {
	if n > 1000 {
		return int(g.concurrent(int64(n), p))
	}
	return int(g.basic(int64(n), p))
}

// Int64 returns X ~ binomial(n, p) as an int64
//
// n = number of tries
// p = probability of acceptance
//
// if n > 1000, the process used to generate the result is concurrent.
func (g *BinomialGenerator) Int64(n int64, p float64) int64 {
	if n > 1000 {
		return g.concurrent(n, p)
	}
	return g.basic(n, p)
}

func (g *BinomialGenerator) basic(n int64, p float64) (retVal int64) {
	for i := int64(0); i < n; i++ {
		if g.Float64() < p {
			retVal++
		}
	}
	return
}

func (g *BinomialGenerator) concurrent(n int64, p float64) (retVal int64) {
	workers := 0
	resChan := make(chan int64)
	for n > 0 {
		size := int64(1000)
		if n-1000 < 0 {
			size = n
		}
		go func(s int64) {
			var res int64
			for i := int64(0); i < s; i++ {
				g.Lock()
				if g.Float64() < p {
					res++
				}
				g.Unlock()
			}
			resChan <- res
		}(size)
		n -= size
		workers++
	}
	var result int64
	for i := 0; i < workers; i++ {
		result += <-resChan
	}
	return result
}

// Binomial returns a slice of Xs drawn from a Binomial(n, p) distribution.
func Binomial(n int, p float64, size int) []int {
	retVal := make([]int, 0, size)
	g := &BinomialGenerator{
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for i := 0; i < size; i++ {
		retVal = append(retVal, g.Int(n, p))
	}
	return retVal
}
