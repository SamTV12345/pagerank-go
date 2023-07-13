package main

import (
	"fmt"
	"math"
)

func main() {
	n := 100

	k := 5

	var fpr = 0.001

	// m = (1/(1-(1-p^1/k)^(1/k/n))

	m := math.Ceil(1 / (1 - math.Pow(1-math.Pow(fpr, 1/float64(k)), 1/(float64(k)*float64(n)))))

	fmt.Println(m)
}
