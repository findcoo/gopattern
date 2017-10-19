package gopattern

import (
	"log"
	"testing"
)

func TestBookShelf(t *testing.T) {
	bookShelf := NewBookShelf()
	bookShelf.AppendBook(
		Book{Name: "로마인 이야기"},
		Book{Name: "나무"},
	)

	it := bookShelf.Iterator()
	for it.HasNext() {
		log.Print(it.Next().(Book))
	}
}
