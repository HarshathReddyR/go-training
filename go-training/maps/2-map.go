package main

import "fmt"

// var dic = make(map[int][]string) //in global we can't use" : "syntax
// func search(k int) {
// 	h, ok := dic[k]
// 	if !ok {
// 		fmt.Println("id not there")
// 		return
// 	}
// 	fmt.Println(h)
// }
func main() {
	dic := make(map[int][]string)
	dic[1] = []string{"play", "dance"}
	dic[2] = []string{"music", "dance"}
	dic[3] = []string{"play", "sleeping"}
	for key, value := range dic {
		fmt.Println(key, value)
	}
	fmt.Println(len(dic))
	delete(dic, 1)
	//search(1)
	fmt.Println(len(dic))
}
