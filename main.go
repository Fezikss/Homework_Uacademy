package main

import (
	"math"
)

func main() {
	//1
	// var slice1 = []int{2, 9, 3, -4, 5}
	// max := 0
	// min := 0

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[max] < slice1[i] {
	// 		max = i
	// 	}
	// }
	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[min] > slice1[i] {
	// 		min = i
	// 	}
	// }
	// fmt.Println("max: ", max, "   min : ", min)

	// slice1[min], slice1[max] = slice1[max], slice1[min]

	// fmt.Println(slice1)

	//2

	// var slice = []int{3, 5, 6, -9, 17, 21, 33, 0, 43, 8}
	// var catchPrime = []int{}
	// counter := 0

	// for i := 0; i < len(slice); i++ {
	// 	if isPrime(slice[i]) {
	// 		counter++
	// 		catchPrime = append(catchPrime, slice[i])
	// 	}
	// }
	// fmt.Println(counter)
	// fmt.Println(catchPrime)

	//3

	// var slice = []int{3, 5, 6, -9, 17, 21, 33, 0, 43, 8}

	// sum1 := 0
	// sum2 := 0
	// for i := 0; i < len(slice); i++ {
	// 	sum1 = sum1 + slice[i]
	// 	sum2 = sum2 + i
	// }
	// if sum2 > sum1 {
	// 	fmt.Println("INDEX")
	// } else if sum2 < sum1 {
	// 	fmt.Println("VALUES")
	// } else {
	// 	fmt.Println("WOW")
	// }

	//4

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

//    s = {1,2,3,4,5,6,4,4,5,5,3}
