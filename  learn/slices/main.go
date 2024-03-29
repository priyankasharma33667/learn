package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")
	var fruitlist = []string{"apple", "mango", "gauava"}
	fmt.Printf("type of fruit list %T\n", fruitlist)

	fruitlist = append(fruitlist, "banana", "chilli")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 254
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 444, 666)

	fmt.Println(highScores)
	sort.Ints(highScores)

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
    courses = append(courses[:index],courses[index+1:]...)
    fmt.Println(courses)



}
