package computation

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Book []string

func CreateBook(bookNames ...string) Book {
	books := Book{}

	for _, book := range bookNames {
		books = append(books, book)
	}

	return books
}

func (book Book) SaveBook(fileName string) {
	err := book.saveToFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
}

func ReadBook(filename string) Book {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return Book(ss)

}

func (b Book) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(b.toString()), 0666)
}

func (b Book) Randomize() { //Creating normal randomizing seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range b {
		newPosition := r.Intn(len(b) - 1)
		b[i], b[newPosition] = b[newPosition], b[i]
	}
}

func (b Book) toString() string {
	return strings.Join([]string(b), ",")
}
