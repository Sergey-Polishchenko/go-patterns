## Factory Method

The **Factory Method** pattern is a **creational** design pattern that provides an interface for creating objects but allows subclasses to alter the type of objects that will be created.
It is particularly useful when a system needs to remain independent of how its objects are created, composed, and represented.

### Characteristics

- **Encapsulation of Object Creation**:
    The creation logic is centralized in a factory method, decoupling it from the main business logic.
- **Interface-Based Products**:
    All products implement a common interface, ensuring consistency and interchangeability.
- **Flexibility**:
    New product types can be added without modifying existing client code.
- **Dynamic Object Creation**:
    The type of object to create is determined at runtime based on input parameters.

### Advantages

- **Decoupling**:
    Client code depends on interfaces, not concrete implementations.
- **Extensibility**:
    New product types can be introduced with minimal changes.
- **Single Responsibility**:
    Object creation logic is isolated in one place.
- **Reusability**:
    Factory methods can be reused across different parts of the application.

### Disadvantages

- **Complexity**:
    Introducing multiple factory methods can increase the complexity of the codebase.
- **OCP Limitation**:
    Adding new product types may require modifying the factory method (e.g., extending a switch-case statement).
- **Overhead**:
    May introduce unnecessary abstraction if the system only deals with a single product type.

### When to Use?

- When the system needs to work with multiple interchangeable product types.
- When the creation logic involves conditions or external parameters.
- When you want to centralize object creation to improve maintainability.
- Common use cases include:
    + File format handlers (e.g., processing different document types),
    + Cross-platform UI components,
    + Payment gateway integrations,
    + Database connector factories,
    + Required for Implementation.

### Required for Implementation

- **Product Interface**:
    Defines a common interface that all products must implement. This ensures that all products are interchangeable.
- **Concrete Products**:
    Classes that implement the Product interface. Each class represents a specific type of product.
- **Factory Method**:
    A method that encapsulates the creation logic. It takes a parameter (e.g., an enum or string) to determine which product to create.
- **Error Handling**:
    The factory method should handle invalid input gracefully, returning an error if an unsupported product type is requested.

### Implementation Details in Go

In Go, the Factory Method pattern is implemented using interfaces and functions.
Unlike languages with classical inheritance, Go relies on composition and interfaces to achieve polymorphism.
The factory method is typically a function that returns an interface type, allowing for flexible and dynamic object creation.

**Key points**:

- Interfaces: Define the contract for products.
- Functions: Act as factory methods, encapsulating creation logic.
- Error Handling: Ensures robustness by handling invalid input.

### Comparison with Abstract Factory

|Aspect | Factory Method | Abstract Factory|
|---|---|---|
|Focus  | Single product creation | Families of related products|
|Complexity | Simpler, one method | Multiple factory methods|
|Extensibility | Add new products | Add new product families|
|Use Case | Dynamic object creation | Cross-platform UI kits, theme systems|
