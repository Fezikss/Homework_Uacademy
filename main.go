package main

import (
	"fmt"
	"math"
)

func SumOfMaxAndMin(a, b, c float64) float64 {
	max := 0.0
	min := 0.0

	if a > b && a > c {
		max = a
	} else if b > c {
		max = b
	} else {
		max = c
	}
	if a < b && a < c {
		min = a
	} else if b < c {
		min = b
	} else {
		min = c
	}

	return max + min
}

func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func solveQuadratic(a, b, c float64) (float64, float64, bool) {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return 0, 0, false
	}

	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return x1, x2, true
}

func fib(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		temp := a + b
		a = b
		b = temp
	}
	return a
}

func photo1(n int) {

	for i := n; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()

	}

}
func photo2(n int) {

	for i := n; i >= 1; i-- {
		for j := i; 0 <= j-1; j-- {
			fmt.Print(" ")
		}
		for k := 0; k <= n-i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
func main() {

	// var (
	// 	a float64
	// 	b float64
	// 	c float64
	// )
	// fmt.Print("a : ")
	// fmt.Scan(&a)
	// fmt.Print("b : ")
	// fmt.Scan(&b)
	// fmt.Print("c : ")
	// fmt.Scan(&c)

	// sum := SumOfMaxAndMin(a, b, c)
	// fmt.Println(sum)

	//2
	// x1, y1 := 1.0, 2.0
	// x2, y2 := 4.0, 6.0

	// dist := distance(x1, y1, x2, y2)
	// fmt.Printf("The distance between (%.2f, %.2f) and (%.2f, %.2f) is %.2f\n", x1, y1, x2, y2, dist)

	//3
	// a, b, c := 1.0, -3.0, 2.0

	// x1, x2, ok := solveQuadratic(a, b, c)
	// if ok {
	// 	fmt.Printf("The solutions are x1 = %.2f and x2 = %.2f\n", x1, x2)
	// } else {
	// 	fmt.Println("No real solutions exist.")
	// }

	//4
	fib := fib(4)

	fmt.Println(fib)

	// photo1(5)
	// photo2(5)

}
