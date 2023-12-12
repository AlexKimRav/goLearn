package main

import "fmt"

func main() {
	bocuNoPico := createHent("boku no piko", "some random jap dude")
	fmt.Println(bocuNoPico)
	bocuNoPico.changeNameHent("babo")
	fmt.Println(bocuNoPico)
}

func testingPointers() {
	a := 5
	b := &a
	*b = 10
	fmt.Printf("a: %d \nb: %d", a, *b)
}

type Hent struct {
	name   string
	author string
}

func (h Hent) String() string {
	return fmt.Sprintln(h.name, "\n"+h.author, "lol")
}

func createHent(name, author string) Hent {
	bocNPic := Hent{name, author}
	return bocNPic
}

func (h *Hent) changeNameHent(newName string) {
	(*h).name = newName
}
