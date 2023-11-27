package main

import (
	"fmt"
)

type Student struct {
	Name        string
	Scholarship int
	Course      int
}
type Airport struct {
	PlaneType     string
	FlightDate    string
	FromCity      string
	ToCity        string
	FlightTime    string
	flightingTime float32
}
type Team struct {
	Name         string
	Coach        string
	PlayersCount int
	Captain      string
}

func main() {
	//1
	// students := []Student{
	// 	{Name: "John Doe", Scholarship: 100, Course: 1},
	// 	{Name: "Jane Smith", Scholarship: 200, Course: 2},
	// 	{Name: "Michael Johnson", Scholarship: 150, Course: 2},
	// 	{Name: "Emily Davis", Scholarship: 300, Course: 3},
	// 	{Name: "William Brown", Scholarship: 250, Course: 1},
	// 	{Name: "Olivia Wilson", Scholarship: 200, Course: 2},
	// 	{Name: "James Taylor", Scholarship: 150, Course: 1},
	// 	{Name: "Sophia Miller", Scholarship: 250, Course: 2},
	// 	{Name: "Benjamin Anderson", Scholarship: 200, Course: 3},
	// 	{Name: "Ava Martinez", Scholarship: 300, Course: 2},
	// }
	// sumOfScholarship := 0

	// for _, student := range students {
	// 	if student.Course == 2 {
	// 		sumOfScholarship += student.Scholarship
	// 	}
	// }

	// fmt.Printf("2-kurs talabalari uchun umumiy scholarship miqdori: %d\n", sumOfScholarship)

	//2

	// students := []Student{
	// 	{Name: "Jam", Scholarship: 100, Course: 1},
	// 	{Name: "Jane Smith", Scholarship: 200, Course: 2},
	// 	{Name: "Michael Johnson", Scholarship: 150, Course: 2},
	// 	{Name: "Emily Davis", Scholarship: 300, Course: 3},
	// 	{Name: "William Brown", Scholarship: 250, Course: 1},
	// 	{Name: "Olivia Wilson", Scholarship: 200, Course: 2},
	// 	{Name: "James Taylor", Scholarship: 150, Course: 1},
	// 	{Name: "Sophia Miller", Scholarship: 250, Course: 2},
	// 	{Name: "azim", Scholarship: 200, Course: 3},
	// 	{Name: "Ava Martinez", Scholarship: 300, Course: 2},
	// }
	// for i := 0; i < len(students); i++ {

	// 	if len(students[i].Name) < 5 {
	// 		fmt.Println(students[i])

	// 	}

	// }

	//3

	// airports := []Airport{
	// 	{PlaneType: "Boeing 747", FlightDate: "2023-07-15", FromCity: "New York", ToCity: "London", FlightTime: "08:00"},
	// 	{PlaneType: "Airbus A380", FlightDate: "2023-09-20", FromCity: "Paris", ToCity: "Dubai", FlightTime: "12:30"},
	// 	{PlaneType: "Boeing 737", FlightDate: "2023-03-25", FromCity: "London", ToCity: "Berlin", FlightTime: "10:15"},
	// 	{PlaneType: "Airbus A320", FlightDate: "2023-07-28", FromCity: "Tokyo", ToCity: "Seoul", FlightTime: "05:45"},
	// 	{PlaneType: "Boeing 787", FlightDate: "2023-07-18", FromCity: "Moscow", ToCity: "Beijing", FlightTime: "14:20"},
	// 	{PlaneType: "Airbus A350", FlightDate: "2023-06-22", FromCity: "Los Angeles", ToCity: "Sydney", FlightTime: "19:50"},
	// 	{PlaneType: "Boeing 777", FlightDate: "2023-07-29", FromCity: "Dublin", ToCity: "Toronto", FlightTime: "09:30"},
	// 	{PlaneType: "Airbus A330", FlightDate: "2023-09-16", FromCity: "Rome", ToCity: "Madrid", FlightTime: "11:10"},
	// 	{PlaneType: "Boeing 767", FlightDate: "2023-12-23", FromCity: "Berlin", ToCity: "Amsterdam", FlightTime: "06:20"},
	// 	{PlaneType: "Airbus A380", FlightDate: "2023-04-26", FromCity: "Singapore", ToCity: "Hong Kong", FlightTime: "15:40"},
	// }

	// fmt.Println("Yoz oylarida uchadigan samolyotlar haqida ma'lumotlar:")
	// for _, airport := range airports {
	// 	if airport.FlightDate[5:7] == "06" || airport.FlightDate[5:7] == "07" || airport.FlightDate[5:7] == "08" {
	// 		fmt.Printf("Samolyot turi: %s, Qayerdan: %s, Qayerga: %s, Uchish vaqti: %s\n", airport.PlaneType, airport.FromCity, airport.ToCity, airport.FlightTime)
	// 	}
	// }

	//4

	// flights := []Airport{
	// 	{PlaneType: "Boeing 747", FlightDate: "2023-11-28", FromCity: "London", ToCity: "Tashkent", FlightTime: "02:30"},
	// 	{PlaneType: "Airbus A380", FlightDate: "2023-11-29", FromCity: "Paris", ToCity: "Tashkent", FlightTime: "03:15"},
	// 	{PlaneType: "Boeing 737", FlightDate: "2023-11-28", FromCity: "Tashkent", ToCity: "New York", FlightTime: "02:45"},
	// 	{PlaneType: "Airbus A320", FlightDate: "2023-11-30", FromCity: "Tashkent", ToCity: "Moscow", FlightTime: "03:30"},
	// 	{PlaneType: "Boeing 787", FlightDate: "2023-11-28", FromCity: "Tashkent", ToCity: "Dubai", FlightTime: "02:15"},
	// 	{PlaneType: "Airbus A350", FlightDate: "2023-11-29", FromCity: "Tashkent", ToCity: "Istanbul", FlightTime: "03:05"},
	// 	{PlaneType: "Boeing 777", FlightDate: "2023-11-28", FromCity: "New Delhi", ToCity: "Tashkent", FlightTime: "02:50"},
	// 	{PlaneType: "Airbus A330", FlightDate: "2023-11-30", FromCity: "Tashkent", ToCity: "Bangkok", FlightTime: "03:40"},
	// 	{PlaneType: "Boeing 767", FlightDate: "2023-11-28", FromCity: "Tashkent", ToCity: "Beijing", FlightTime: "02:20"},
	// 	{PlaneType: "Airbus A380", FlightDate: "2023-11-29", FromCity: "Tashkent", ToCity: "Seoul", FlightTime: "03:10"},
	// }

	// fmt.Println("Uchish soati 2 va 3 oralig'idagi va qo'nish shaxri 'Toshkent' bo'lgan samolyotlar:")
	// for _, flight := range flights {
	// 	if flight.FlightTime >= "02:00" && flight.FlightTime <= "03:59" && flight.FromCity == "Tashkent" {
	// 		fmt.Printf("Samolyot turi: %s, Qayerdan: %s, Qayerga: %s, Uchish vaqti: %s\n", flight.PlaneType, flight.FromCity, flight.ToCity, flight.FlightTime)
	// 	}
	// }

	//5
	// teams := []Team{
	// 	{
	// 		Name:         "Barcelona",
	// 		Coach:        "Ronald Koeman",
	// 		PlayersCount: 25,
	// 		Captain:      "Lionel Messi"},

	// 	{
	// 		Name:         "Real Madrid",
	// 		Coach:        "Carlo Ancelotti",
	// 		PlayersCount: 23,
	// 		Captain:      "Sergio Ramos"},
	// 	{
	// 		Name:         "Manchester United",
	// 		Coach:        "Ole Gunnar Solskjaer",
	// 		PlayersCount: 28,
	// 		Captain:      "Harry Maguire"},
	// 	{
	// 		Name:         "Bayern Munich",
	// 		Coach:        "Julian Nagelsmann",
	// 		PlayersCount: 22,
	// 		Captain:      "Manuel Neuer"},
	// }

	// for i := 0; i < len(teams)-1; i++ {
	// 	for j := i + 1; j < len(teams); j++ {
	// 		if teams[i].PlayersCount < teams[j].PlayersCount {
	// 			teams[i], teams[j] = teams[j], teams[i]
	// 		}
	// 	}
	// }

	// for _, team := range teams {
	// 	fmt.Printf("Team: %s\n", team.Name)
	// 	fmt.Printf("Coach: %s\n", team.Coach)
	// 	fmt.Printf("Players Count: %d\n", team.PlayersCount)
	// 	fmt.Printf("Captain: %s\n", team.Captain)
	// 	fmt.Println()
	// }

}
