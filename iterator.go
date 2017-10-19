package gopattern

// Aggregate 한 집합체의 반복자를 생성하기 위한 인터페이스
type Aggregate interface {
	Iterator() Iterator
}

// Iterator 대상을 하나씩 조회하는 인터페이스
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// BookShelfIterator 서가를 검색하는 구조체
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

// HasNext 책이 있다면 true를 반환한다.
func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.GetLength()
}

// Next 책을 반환하고 다음 책의 위치로 이동
func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}

// BookShelf 서가 구조
type BookShelf struct {
	books []Book
}

// NewBookShelf BookShelf 생성 함수
func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: []Book{},
	}
}

// GetBookAt 위치에 해당하는 책을 반환한다.
func (bs *BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

// AppendBook 서가에 책을 추가한다.
func (bs *BookShelf) AppendBook(book ...Book) {
	bs.books = append(bs.books, book...)
}

// GetLength 서가의 최대 길이를 나타낸다.
func (bs *BookShelf) GetLength() int {
	return len(bs.books)
}

// Iterator 서가에 대한 반복자를 생성한다.
func (bs *BookShelf) Iterator() *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bs}
}

// Book 책 구조
type Book struct {
	Name string
}
