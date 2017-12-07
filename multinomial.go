package randomkit

import (
	"math/rand"
	"time"
)

// Multinomial returns a slice of ints drawn from the probabilities provided.
func Multinomial(n int, probs []float64, retSize int) []int {
	d := len(probs)
	if Kahan(probs) > (1.0 + 1e-12) {
		panic("Probabilities add up to greater than 1!")
	}
	// retSize must be wholely divisible by d
	if retSize%d != 0 {
		panic("The size of the return vector has to be wholely divisible by the size of the probabilities")
	}

	g := &BinomialGenerator{
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	retVal := make([]int, retSize)
	for i := 0; i < len(retVal); {
		sum := 1.0
		dn := n
		for j := 0; j < len(probs)-1; j++ {
			retVal[i+j] = g.Int(dn, probs[j]/sum)
			dn -= retVal[i+j]
			if dn <= 0 {
				break
			}
			sum -= probs[j]
		}
		if dn > 0 {
			retVal[i+d-1] = dn
		}
		i += d
	}
	return retVal
}
