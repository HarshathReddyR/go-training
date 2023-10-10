package main

import "fmt"

type user struct {
	name string
	age  int
}

type movies struct {
	movieName string
	// anonymous field which have no name
	// anonymous field get there names from the types
	user // embedding
	Book
	Game
}
type Book struct {
	Bname string
}
type Game struct {
	Gname string
}

func main() {
	m := movies{
		movieName: "abc",
		user: user{
			name: "xyz",
			age:  29,
		},
		Book: Book{
			Bname: "read",
		},
		Game: Game{
			Gname: "Cricket",
		},
	}

	fmt.Println(m.age)
	m.update(30)
	fmt.Println(m.age)
}
func (m *movies) update(a int) {
	m.age = a
}
