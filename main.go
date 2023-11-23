package main

import "fmt"

func main() {
	// sum := 0
	// finder := 0
	// var slice1 = []int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10}
	// a := 0
	// b := 0

	// for i := 0; i < 10; i++ {
	// 	if slice1[i] < 0 {
	// 		finder++
	// 	}
	// 	if finder == 2 {
	// 		a = i
	// 	}
	// 	if finder == 4 {
	// 		b = i
	// 		break
	// 	}
	// }
	// for i := a; i < b; i++ {
	// 	sum += slice1[i]
	// }
	// fmt.Println(sum)
	//2
	// var slice1 = []int{0, -1, 2, -3, 4, 0, 5, -6, 7, 0}
	// musbat := 0
	// manfiy := 0
	// toq := 0
	// juft := 0
	// nol := 0

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] == 0 {
	// 		nol++
	// 	}
	// 	if slice1[i] > 0 {
	// 		musbat++
	// 	}
	// 	if slice1[i] < 0 {
	// 		manfiy++
	// 	}
	// 	if slice1[i]%2 == 0 {
	// 		juft++
	// 	}
	// 	if slice1[i]%2 == 1 {
	// 		toq++
	// 	}
	// }
	// fmt.Println("musbat", musbat, " ta ")
	// fmt.Println("manfiy", manfiy, " ta ")
	// fmt.Println("nol", nol, " ta ")
	// fmt.Println("toq", toq, " ta ")
	// fmt.Println("juft", juft, " ta ")

	//3

	// var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var juft = []int{}
	// var toq = []int{}

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i]%2 == 0 {
	// 		juft = append(juft, slice1[i])
	// 	}
	// 	if slice1[i]%2 == 1 {
	// 		toq = append(toq, slice1[i])
	// 	}
	// }
	// fmt.Println(juft)
	// fmt.Println(toq)

	//4

	// var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var juft = 1
	// var toq = 0

	// for i := 0; i < len(slice1); i++ {
	// 	if i%2 == 0 {
	// 		juft *= slice1[i]
	// 	}
	// 	if slice1[i]%2 == 1 {
	// 		toq += slice1[i]
	// 	}
	// }
	// fmt.Println(juft)
	// fmt.Println(toq)

	//5

	// var slice1 = []int{11, 11, 33, 3, 4, 5, 41, 6, 2, 1, 8, 9, 10}

	// min := slice1[0]
	// secondMin := slice1[1]

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] < min {
	// 		min = slice1[i] // 1
	// 	}
	// }

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] < secondMin && slice1[i] != min {
	// 		secondMin = slice1[i]
	// 	}
	// }

	// fmt.Println(min, secondMin)

	//6

	// var slice1 = []int{11, 11, 33, 3, 4, 5, 41, 6, 2, 1, 8, 9, 10}

	// max := slice1[0]
	// secondMax := slice1[1]

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] > max {
	// 		max = slice1[i] // 1
	// 	}
	// }

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] > secondMax && slice1[i] != max {
	// 		secondMax = slice1[i]
	// 	}
	// }

	// fmt.Println(max, secondMax)

	//7

	// var slice1 = []int{11, 11, 33, 3, 4, 5, 41, 6, 2, 1, 8, 9, 10}

	// slice1[0], slice1[len(slice1)-1] = slice1[len(slice1)-1], slice1[0]

	// fmt.Println(slice1)

	//8
	// var slice1 = []int{11, 11, 33, 3, 4, 5, 41, 6, 2, 1, 8, 9, 10}

	// counter := 0
	// a := 0
	// max := slice1[0]

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] > max {
	// 		max = slice1[i]
	// 		a = i
	// 	}
	// }

	// for i := a + 1; i < len(slice1); i++ {
	// 	counter++
	// }

	// fmt.Println("max dan keyin ", counter, " ta son bor")

	//9

	// var slice1 = []int{11, 11, 33, 3, 4, 5, 41, 6, 2, 1, 8, 9, 10}

	// counter := 0
	// a := 0
	// min := slice1[0]

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] < min {
	// 		min = slice1[i]
	// 		a = i
	// 	}
	// }

	// for i := 0; i < a; i++ {
	// 	counter++
	// }

	// fmt.Println("min dan oldin ", counter, " ta son bor")

	//10

	// var slice1 = []int{11, 11, 33, 3, 4, 5, 41, 6, 2, 8, 9, 10}

	// counter := 0
	// a := 0
	// b := 0
	// min := slice1[0]
	// max := slice1[0]

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] < min {
	// 		min = slice1[i]
	// 		a = i
	// 	}
	// }

	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] > max {
	// 		max = slice1[i]
	// 		b = i
	// 	}
	// }

	// fmt.Println(a, b)
	// if a > b {
	// 	for i := b + 1; i < a; i++ {
	// 		counter++
	// 	}
	// } else {
	// 	for i := a + 1; i < b; i++ {
	// 		counter++
	// 	}
	// }

	// fmt.Println("min dan max ", counter, " ta son bor")

}
