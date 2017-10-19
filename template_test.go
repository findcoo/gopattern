package gopattern

import "testing"

func TestTemplatePattern(t *testing.T) {
	var displayer AbstractPainter = NewCharDisplay('A')
	displayer.Display()
}
