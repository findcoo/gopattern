package gopattern

// Painter Painter Template 인터페이스
type Painter interface {
	Open()
	Print()
	Close()
}

// Paint Painter를 진행하는 상수 함수
func Paint(p Painter) {
	p.Open()
	for i := 0; i < 5; i++ {
		p.Print()
	}
	p.Close()
}

// CharPainter 문자 하나만 출력하는 하위 구조
type CharPainter struct {
	ch rune
}

// NewCharPainter CharDisplay 생성자
func NewCharPainter(c rune) *CharPainter {
	return &CharPainter{ch: c}
}

// Open implemente AbstractPainter
func (cp *CharPainter) Open() {
	print("<<")
}

// Print implemente AbstractPainter
func (cp *CharPainter) Print() {
	print(string(cp.ch))
}

// Close implemente AbstractPainter
func (cp *CharPainter) Close() {
	print(">>")
}
