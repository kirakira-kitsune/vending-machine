package main

import "fmt"

var bevMap map[string]int

var bottleCost float64

func main() {

	bottleCost = 3

	bevMap = make(map[string]int)

	bevMap["Sencha"] = 1
	bevMap["Hojicha"] = 10
	bevMap["Genmaicha"] = 10
	bevMap["Sparkling Water"] = 10
	bevMap["Lemon Water"] = 10
	bevMap["Iced Coffee"] = 10
	bevMap["Iced Chocolate"] = 10
	bevMap["Mikan Juice"] = 10

	fmt.Println("The vending machine contains: ")

	for key := range bevMap {

		fmt.Printf("%s\n", key)
	}

	fmt.Printf("Each beverage is: $%f\n", bottleCost)


	for {
		selectBeverage()

		var myBeverage string

		fmt.Println("Do you want to continue? Enter y or yes")

		fmt.Scan(&myBeverage)

		if myBeverage != "y" && myBeverage != "yes" {
			return
		}

	}

}

func selectBeverage() {

	var myChoice string

	fmt.Println("What do you want to order?")

	fmt.Scan(&myChoice)

	bottles, ok := bevMap[myChoice]

	// if not ok (key not in map)
	// return
	if !ok {
		return
	}

	fmt.Println("Please insert coins now")

	var coinsReceived float64

	fmt.Scan(&coinsReceived)

	if coinsReceived < bottleCost {
		fmt.Println("Not enough coins received")
		giveChange(0.0, coinsReceived)
		return
	}
	// code from here is only executed if we didn't return above

	if bottles >= 1 {
		fmt.Printf("Here's your %s\n", myChoice)
		bevMap[myChoice] -= 1
		giveChange(bottleCost, coinsReceived)
	} else {
		fmt.Println("Sorry, this is sold out. Please select another drink")
		giveChange(0.0, coinsReceived)
	}
}

func giveChange(price float64, paid float64)  {

	change := paid - price

	fmt.Printf("Your change is: $%f\n", change)
}
