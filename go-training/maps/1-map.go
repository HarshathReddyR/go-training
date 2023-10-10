package main

import "fmt"

func main() {
	dictionary := make(map[string]string, 5)

	dictionary["up"] = "above"
	dictionary["below"] = "down"

	fmt.Println(dictionary["up"])
	for k, _ := range dictionary {
		fmt.Println(k)
	}
	fmt.Println(len(dictionary))
	fmt.Println(dictionary) //map[below:down up:above]
}
