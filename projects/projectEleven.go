package main

import "fmt"

func mapOne() {
	user := make(map[string]string)

	user["name"] = "Rahma"
	user["country"] = "Oman"

	fmt.Println(user)
}

func mapTwo() {
	user := map[string]string{
		"name":    "Rahma",
		"country": "Oman",
		"job":     "Backend Developer",
	}

	fmt.Println(user)
}

func mapThree() {
	user := map[string]string{
		"name": "Rahma",
		"city": "Muscat",
	}

	fmt.Println(user["name"])
	fmt.Println(user["city"])
}

func mapFour() {
	scores := make(map[string]int)

	scores["Math"] = 95
	scores["Science"] = 88

	fmt.Println(scores)
}

func mapFive() {
	scores := map[string]int{
		"Math": 80,
	}

	scores["Math"] = 100

	fmt.Println(scores)
}

func mapSix() {
	user := map[string]string{
		"name": "Rahma",
		"city": "Muscat",
	}

	delete(user, "city")

	fmt.Println(user)
}

func mapSeven() {
	user := map[string]string{
		"name": "Rahma",
	}

	// value, exists := user["name"]

	// fmt.Println(value)
	// fmt.Println(exists)

	value, exists := user["age"]

	fmt.Println(value)
	fmt.Println(exists)
}

func mapEight() {
	user := map[string]string{
		"name":    "Rahma",
		"country": "Oman",
		"job":     "Backend Developer",
	}

	for key, value := range user {
		fmt.Println(key, ":", value)
	}
}

func mapNine() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(len(numbers))
}

func mapTen() {
	users := map[string]map[string]string{
		"user1": {
			"name": "Rahma",
			"city": "Muscat",
		},
		"user2": {
			"name": "Ahmed",
			"city": "Nizwa",
		},
	}

	fmt.Println(users["user1"]["name"])
}
