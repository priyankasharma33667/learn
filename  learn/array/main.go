package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Orange"
	// fruitList[0] = "Apple"

	fmt.Println("fruitlist", fruitList)
	fmt.Println("fruitlist", len(fruitList))

	var vegetable = [5]string{"potato", "mushrrom", "tomato"}
	fmt.Println(vegetable)
	fmt.Println(len(vegetable))
}
