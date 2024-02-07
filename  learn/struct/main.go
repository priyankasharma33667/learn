//struct is alternative version of clasess in golang

package main

import "fmt"

func main() {
	fmt.Println("struct in golang")
	//no inheritance in golang
	priyanka := User{"Priyanka", "priyanka@go.dev", true, 23}
	fmt.Println(priyanka)
	fmt.Printf("priyanka details are : %+v\n", priyanka)
	fmt.Printf("name is %v and emailis %v", priyanka.Name, priyanka.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
