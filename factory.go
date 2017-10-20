package gopattern

import "fmt"

// Product factory에 의해 생성되는 대상 인터페이스
type Product interface {
	Use()
}

// ProductFactory Product를 생성하는 factory 인터페이스
type ProductFactory interface {
	CreateProduct() Product
	RegisterProduct()
}

// IDCard 카드 구조체
type IDCard struct {
	Owner string
}

// Use implement Product
func (idc *IDCard) Use() {
	fmt.Printf("%s 의 카드를 사용합니다.", idc.Owner)
}

// IDCardFactory IDCard를 생성하는 factory 구조체
type IDCardFactory struct {
	ownerArray []string
}

// NewIDCardFactory IDCardFactory 생성자
func NewIDCardFactory() *IDCardFactory {
	return &IDCardFactory{ownerArray: []string{}}
}

// CreateProduct implement ProductFactory
func (idcf *IDCardFactory) CreateProduct(owner string) *IDCard {
	return &IDCard{Owner: owner}
}

// RegisterProduct implement RegisterProduct
func (idcf *IDCardFactory) RegisterProduct(product Product) {
	idcf.ownerArray = append(idcf.ownerArray, product.(*IDCard).Owner)
}
