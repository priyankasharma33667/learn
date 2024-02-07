//map is basically hash(key value pair)

package main

import "fmt"

func main() {
	fmt.Println("maps in  golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("js for ", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loops are interesting in golangs

	for _, value := range languages {
		fmt.Println("for value is %v\n", value)
	}

}
