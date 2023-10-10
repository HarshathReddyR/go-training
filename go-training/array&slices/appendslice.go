package main

import "fmt"

func main() {
	// a := []int{10, 20, 30, 40, 50, 60}
	// b := a[2:5]
	// b = append(b, 999)
	// fmt.Println(a)
	// fmt.Println(b)
	a := []int{10, 20, 30, 40, 50}
	i := a[1:5]

	//fmt.Println(i)
	fmt.Println(len(i), cap(i))
	fmt.Println(len(a), cap(a))
	i = append(i, 60)
	fmt.Println(len(i), cap(i))
	fmt.Println(len(a), cap(a))
	fmt.Printf("%p", a)
	fmt.Println()
	a = append(a, 60) //here creating the new memory due to capacit and the lenght is same for a and update the capacity of the a will become the double of before
	i[0] = 99
	fmt.Println(len(i), cap(i))
	fmt.Println(len(a), cap(a))
	fmt.Printf("%p", a)
	fmt.Println()
	fmt.Println(a)
	fmt.Println(i)
}
