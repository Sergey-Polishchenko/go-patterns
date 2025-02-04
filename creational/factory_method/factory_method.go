package factorymethod

import "errors"

type actor int

const (
	A actor = iota
	B
	C
)

type Product interface {
	DoSmt()
}

func NewProduct(actor actor) (Product, error) {
	switch actor {
	case A:
		return &ProductA{}, nil
	case B:
		return &ProductB{}, nil
	case C:
		return &ProductC{}, nil
	default:
		return nil, errors.New("unknown actor")
	}
}

type ProductA struct{}

func (p ProductA) DoSmt() {}

type ProductB struct{}

func (p ProductB) DoSmt() {}

type ProductC struct{}

func (p ProductC) DoSmt() {}
