package builder

import "fmt"

type Builder interface {
	MakeConfig()
	MakeLogging()
	MakeDB()
	GetProduct() *Product
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.MakeConfig()
	d.builder.MakeLogging()
	d.builder.MakeDB()
}

type SetupBuilder struct {
	product *Product
}

func NewSetupBuilder() *SetupBuilder {
	return &SetupBuilder{product: &Product{}}
}

func (b *SetupBuilder) MakeConfig() {
	b.product.Config = "Config initialized"
}

func (b *SetupBuilder) MakeLogging() {
	b.product.Logging = true
}

func (b *SetupBuilder) MakeDB() {
	b.product.DB = "DB connected"
}

func (b *SetupBuilder) GetProduct() *Product {
	return b.product
}

type Product struct {
	Config  string
	Logging bool
	DB      string
}

func (p *Product) Show() string {
	return fmt.Sprintf("Config: %s, Logging: %t, DB: %s", p.Config, p.Logging, p.DB)
}
