package tests

import (
	comp "goTest/books"
	"testing"
)

func TestNewBook(t *testing.T) {
	books := comp.CreateBook("451", "Kafka", "AmericanPsycho")
	if len(books) != 3 {
		t.Errorf("expected size of 4 but got %d", len(books))
	}
}
