package builtins

/**
 * Possible Optimizations:
 *   - Only save odd numbers in bit-array
 *   - Don't save multiples of five
 *   - Look into starting at the top for factorizations?
 */

import (
	"fmt"
	"math"
	"strconv"
)

func Factorize(argv []string) error {
	// var number, _ = strconv.ParseUint(argv[0], 10, 64)
	var number, _ = strconv.Atoi(argv[0])
	// var sieve = sievePrimes(number)
	var sieve = sievePrimes(int(math.Sqrt(float64(number))))
	var bitcount = len(sieve) * 64 // Maximum Bits allocated

	// var counter int
	for i := 0; i < bitcount; i++ {
		if getBit(sieve, i) == 0 {
			// counter++
			for ; number%i == 0; number /= i {
				fmt.Print(i, " ")
			}
		}

		// // Print progress every percent
		// if i%(number/100) == 0 {
		// 	// fmt.Print(float32(i)/float32(number), "#")
		// 	fmt.Print("#")
		// }
	}
	// fmt.Println("\nNumber of Primes:", counter)
	fmt.Println()

	return nil
}

func sievePrimes(number int) []uint64 {
	var sieve = make([]uint64, number/64+1)
	for i := 2; i < number; i++ {
		for j := i * i; j < number; j += i {
			setBit(sieve, j, 1)
		}

		// // Print progress every percent
		// if i%(number/100) == 0 {
		// 	// fmt.Print(float32(i)/float32(number), "#")
		// 	fmt.Print("#")
		// }
	}

	setBit(sieve, 0, 1) // Manually mark 0 as not prime
	setBit(sieve, 1, 1) // Manually mark 1 as not prime

	return sieve
}

func setBit(array []uint64, index int, value uint64) {
	x, y := getIndex(uint64(index))
	array[x] |= (value << y)
}

func getBit(array []uint64, index int) uint64 {
	x, y := getIndex(uint64(index))
	return (array[x] >> y) & 0x1
}

/**
 *
 *             +-- 1, 0 = 8
 *             |
 * 0000 0000 | 0000 0000
 *  |             |
 *  +-- 0, 1 = 1  +-- 1, 3 = 11
 *
 *                    +-- 1, 63 = 127
 *                    |
 * 00.. ..00 | 00.. ..00
 *         |
 *         +-- 0, 63 = 63
 *
 * x * sizeof(T) + y <=> x = index / sizeof(T),
 *                       y = index % sizeof(T)
 * x = number in array
 * y = bit in number
 */
func getIndex(index uint64) (uint64, uint64) {
	return index / 64, index % 64
}
