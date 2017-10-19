package gopattern

import (
	"github.com/fatih/color"
)

// Display 문구를 출력하는 인터페이스
type Display interface {
	PrintWeak()
	PrintStrong()
}

// DisplayBanner 배너를 출력하는 구조
type DisplayBanner struct {
	*Banner
}

// NewDisplayBanner DisplayBanner 생성자
func NewDisplayBanner(p string) *DisplayBanner {
	return &DisplayBanner{
		Banner: NewBanner(p),
	}
}

// PrintWeak 배너 출력의 어댑터 메소드
func (db *DisplayBanner) PrintWeak() {
	db.showBlue()
}

// PrintStrong 배너 출력의 어댑터 메소드
func (db *DisplayBanner) PrintStrong() {
	db.showRed()
}

// Banner 배너
type Banner struct {
	phrase string
}

// NewBanner 배너 생성자
func NewBanner(p string) *Banner {
	return &Banner{phrase: p}
}

func (b *Banner) showRed() {
	color.Red(b.phrase)
}

func (b *Banner) showBlue() {
	color.Blue(b.phrase)
}
