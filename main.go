package main

import (
	"fmt"
)

type Car struct {
	Name       string
	Price      float64
	Model      string
	HorsePower int
}
type Student struct {
	Name        string
	Age         int
	Scholarship int
	Score       []int
}
type Team struct {
	Name         string
	Coach        string
	PlayersCount int
	Captain      string
}

func main() {
	//1
	// cars := []Car{
	// 	{
	// 		Name:       "aame",
	// 		Price:      0,
	// 		Model:      "String",
	// 		HorsePower: 2,
	// 	},
	// 	{
	// 		Name:       "dame",
	// 		Price:      0,
	// 		Model:      "String",
	// 		HorsePower: 2,
	// 	},
	// 	{
	// 		Name:       "bame",
	// 		Price:      0,
	// 		Model:      "String",
	// 		HorsePower: 2,
	// 	},
	// 	{
	// 		Name:       "came",
	// 		Price:      0,
	// 		Model:      "String",
	// 		HorsePower: 2,
	// 	},
	// 	{
	// 		Name:       "yame",
	// 		Price:      0,
	// 		Model:      "String",
	// 		HorsePower: 2,
	// 	},
	// }

	// for i := 0; i < len(cars); i++ {
	// 	for j := i + 1; j < len(cars); j++ {
	// 		if cars[i].Name > cars[j].Name {
	// 			cars[i], cars[j] = cars[j], cars[i]
	// 		}
	// 	}
	// }

	//	fmt.Println(cars)

	//2

	teams := []Team{
		{
			Name:         "Barcelona",
			Coach:        "Ronald Koeman",
			PlayersCount: 25,
			Captain:      "Lionel Messi"},

		{
			Name:         "Real Madrid",
			Coach:        "Carlo Ancelotti",
			PlayersCount: 23,
			Captain:      "Sergio Ramos"},
		{
			Name:         "Manchester United",
			Coach:        "Ole Gunnar Solskjaer",
			PlayersCount: 28,
			Captain:      "Harry Maguire"},
		{
			Name:         "Bayern Munich",
			Coach:        "Julian Nagelsmann",
			PlayersCount: 22,
			Captain:      "Manuel Neuer"},
	}

	for i := 0; i < len(teams)-1; i++ {
		for j := i + 1; j < len(teams); j++ {
			if teams[i].PlayersCount > teams[j].PlayersCount {
				teams[i], teams[j] = teams[j], teams[i]
			}
		}
	}

	for _, team := range teams {
		fmt.Printf("Team: %s\n", team.Name)
		fmt.Printf("Coach: %s\n", team.Coach)
		fmt.Printf("Players Count: %d\n", team.PlayersCount)
		fmt.Printf("Captain: %s\n", team.Captain)
		fmt.Println()
	}

	//3
	// sumOfScore := 0

	// students := []Student{
	// 	{
	// 		Name:        "Anvar",
	// 		Age:         19,
	// 		Scholarship: 75,
	// 		Score:       []int{4, 5, 4, 3, 5},
	// 	},
	// 	{
	// 		Name:        "Jamshid",
	// 		Age:         19,
	// 		Scholarship: 75,
	// 		Score:       []int{5, 5, 5, 5, 5},
	// 	}, {
	// 		Name:        "Sarvar",
	// 		Age:         19,
	// 		Scholarship: 75,
	// 		Score:       []int{4, 4, 4, 5, 4},
	// 	}, {
	// 		Name:        "Gulom",
	// 		Age:         19,
	// 		Scholarship: 75,
	// 		Score:       []int{3, 3, 3, 3, 3},
	// 	}, {
	// 		Name:        "Polat",
	// 		Age:         19,
	// 		Scholarship: 75,
	// 		Score:       []int{4, 3, 4, 4, 3},
	// 	},
	// }

	// for i := 0; i < len(students); i++ {
	// 	for j := 0; j < len(students[i].Score); j++ {
	// 		sumOfScore += students[i].Score[j]
	// 	}
	// 	if sumOfScore/len(students[i].Score) >= 4 {
	// 		fmt.Println(students[i])
	// 	}
	// 	sumOfScore = 0
	// }
}
