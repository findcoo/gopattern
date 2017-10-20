package gopattern

import "testing"

func TestFactoryPattern(t *testing.T) {
	idCardFactory := NewIDCardFactory()
	card1 := idCardFactory.CreateProduct("findcoo")
	card1.Use()

	idCardFactory.RegisterProduct(card1)
}
