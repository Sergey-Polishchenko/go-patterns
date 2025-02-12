package prototype

type Prototyper interface {
	Clone() Prototyper
	GetName() string
}

type ConcreteProduct struct {
	name string
}

func NewConcreteProduct(name string) Prototyper {
	return &ConcreteProduct{
		name: name,
	}
}

func (p *ConcreteProduct) GetName() string {
	return p.name
}

func (p *ConcreteProduct) Clone() Prototyper {
	return &ConcreteProduct{p.name}
}
