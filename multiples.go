/*
* If we list all the natural numbers below 10 that are multiples of 3 or 5,
* we get 3, 5, 6 and 9. The sum of these multiples is 23.
* Find the sum of all the multiples of 3 or 5 below 1000.
 */

package main

import (
	"fmt"
)

func sum_multiples(N int) (sum int) {
	sum = 0
	for i := 0; i < N; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return
}

func main() {
	s := sum_multiples(100)
	fmt.Println(s)
}
