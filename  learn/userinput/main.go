package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating of our pizza:")

	input, _ := reader.ReadString('\n') 
	fmt.Println("thanks for reading", input)
	fmt.Printf("type of the rating is %T", input)
}
