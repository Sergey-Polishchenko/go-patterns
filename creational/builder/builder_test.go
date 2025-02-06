package builder

import (
	"reflect"
	"strings"
	"testing"
)

type MockBuilder struct {
	calls []string
}

func (m *MockBuilder) MakeConfig() {
	m.calls = append(m.calls, "config")
}

func (m *MockBuilder) MakeLogging() {
	m.calls = append(m.calls, "logging")
}

func (m *MockBuilder) MakeDB() {
	m.calls = append(m.calls, "db")
}

func (m *MockBuilder) GetProduct() *Product {
	return &Product{}
}

func TestNewDirector(t *testing.T) {
	mb := &MockBuilder{}
	director := NewDirector(mb)

	if director.builder != mb {
		t.Errorf("Expected builder %v, got %v", mb, director.builder)
	}
}

func TestDirector_Construct(t *testing.T) {
	mb := &MockBuilder{}
	director := NewDirector(mb)

	director.Construct()

	expectedCalls := []string{"config", "logging", "db"}
	if !reflect.DeepEqual(mb.calls, expectedCalls) {
		t.Errorf("Expected calls %v, got %v", expectedCalls, mb.calls)
	}
}

func TestNewSetupBuilder(t *testing.T) {
	builder := NewSetupBuilder()

	if builder.product == nil {
		t.Error("Expected product to be initialized")
	}
}

func TestSetupBuilder_MakeConfig(t *testing.T) {
	builder := NewSetupBuilder()
	builder.MakeConfig()

	if !strings.Contains(builder.product.Config, "Config initialized") {
		t.Errorf("Expected config initialized, got: %s", builder.product.Config)
	}
}

func TestSetupBuilder_MakeLogging(t *testing.T) {
	builder := NewSetupBuilder()
	builder.MakeLogging()

	if !builder.product.Logging {
		t.Error("Expected logging to be enabled")
	}
}

func TestSetupBuilder_MakeDB(t *testing.T) {
	builder := NewSetupBuilder()
	builder.MakeDB()

	if !strings.Contains(builder.product.DB, "DB connected") {
		t.Errorf("Expected DB connection, got: %s", builder.product.DB)
	}
}

func TestSetupBuilder_GetProduct(t *testing.T) {
	builder := NewSetupBuilder()
	builder.MakeConfig()
	product := builder.GetProduct()

	if product.Config == "" {
		t.Error("Expected product with config, got empty")
	}
}

func TestProduct_Show(t *testing.T) {
	p := &Product{
		Config:  "test_config",
		Logging: true,
		DB:      "test_db",
	}

	expected := "Config: test_config, Logging: true, DB: test_db"
	if result := p.Show(); result != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, result)
	}
}
