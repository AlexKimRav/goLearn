package main

import (
	"fmt"
	comp "goTest/books"
	"goTest/users"
)

func main() {

}

func trashCan() {
	fmt.Println("Введите название книги")
	var bookFileName string
	fmt.Scan(&bookFileName)

	books := comp.CreateBook("Metamorphosis", "Babylon", "451")
	books.Randomize()
	books.SaveBook(bookFileName)

	fmt.Println(comp.ReadBook(bookFileName))

	user := users.CreateUser("alex", "bb@gmail.com", "wdkwod", "superNickname", 29)
	fmt.Println(user)
}
