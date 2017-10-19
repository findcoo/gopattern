package gopattern

import "log"

// Painter Display 메소드를 정의하기 위한 선행 구조
type Painter struct {
	Open  func()
	Print func()
	Close func()
}

// Display Painter의 전체 진행을 시작
func (p *Painter) Display() {
	p.Open()
	for i := 0; i < 5; i++ {
		p.Print()
	}
	p.Close()
}

// AbstractPainter Painter Template 인터페이스
type AbstractPainter interface {
	Open()
	Print()
	Close()
	Display()
}

// CharDisplay 문자 하나만 출력하는 하위 구조
type CharDisplay struct {
	ch rune
	*Painter
}

// NewCharDisplay CharDisplay 생성자
func NewCharDisplay(c rune) *CharDisplay {
	return &CharDisplay{ch: c}
}

// Open implemente AbstractPainter
func (charD *CharDisplay) Open() {
	charD.Painter.Open = func() {
		log.Print("<<")
	}
	charD.Painter.Open()
}

// Print implemente AbstractPainter
func (charD *CharDisplay) Print() {
	log.Print(charD.ch)
}

// Close implemente AbstractPainter
func (charD *CharDisplay) Close() {
	log.Print(">>")
}
