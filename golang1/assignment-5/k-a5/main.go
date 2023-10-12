package main

import "fmt"

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
func main() {
	numbers := []int{1, 2, 3, 2, 4, 1, 5}
	var duplicateNumbers []int
	d := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				d++
				if d > 1 {
					duplicateNumbers = append(duplicateNumbers, numbers[i])
					numbers = deleteElement(numbers, i)

				}
			}
		}
	}
	fmt.Println(duplicateNumbers)
}
