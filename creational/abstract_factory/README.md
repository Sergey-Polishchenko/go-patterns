## Abstract Factory

The **Abstract Factory** pattern is a **creational** design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes.
It ensures that the created objects are compatible and designed to work together, decoupling client code from platform-specific or family-specific implementations.

### Characteristics

- **Family-Centric Creation**:
    Focuses on producing groups of objects that belong to a single logical family (e.g., UI components for a specific OS).
- **Interface-Driven Design**:
    Clients interact only with abstract interfaces for factories and products, hiding implementation details.
- **Consistency Guarantee**:
    Objects from the same factory are inherently compatible (e.g., buttons, menus, and dialogs for a unified theme).
- **Runtime Flexibility**:
    Enables switching between entire families of objects (e.g., light/dark mode, OS-specific themes) dynamically.

### Advantages

- **Consistent Object Ecosystems**:
    Products from the same factory are guaranteed to work together seamlessly.
- **Decoupling**:
    Client code depends on abstractions, not concrete implementations.
- **Scalability**:
    New families of products can be added without modifying existing client logic.
- **Centralized Control**:
    Object creation rules and compatibility are enforced in one place.

### Disadvantages

- **Complex Initial Setup**:
    Requires defining numerous interfaces and classes, even for small use cases.
- **Rigid Structure**:
    Adding a new product type (e.g., a new UI widget) forces changes across all factory interfaces and implementations.
- **Over-Engineering Risk**:
    Unnecessary for systems that only require a single family of objects.

### When to Use?

- When the system must support multiple independent families of objects (e.g., cross-platform UI kits).
- When objects within a family must be used together (e.g., theme components like fonts, colors, and icons).
- To isolate client code from platform-specific or configuration-specific dependencies.
- Common Use Cases:
    + Cross-platform applications (Windows/macOS UI components),
    + Theming systems (dark/light mode with coordinated elements),
    + Cloud service integrations (AWS vs. Google Cloud SDKs),
    + Game development (rendering pipelines for DirectX/Vulkan).

### Required for Implementation

- **Abstract Factory Interface**:
    Declares methods for creating each product type in the family (e.g., CreateButton(), CreateDialog()).
- **Concrete Factories**:
    Implement the abstract factory interface to produce products for a specific family (e.g., WindowsFactory, MacFactory).
- **Abstract Product Interfaces**:
    Define common APIs for each product type (e.g., Button, Dialog).
- **Concrete Products**:
    Platform- or family-specific implementations of product interfaces (e.g., WindowsButton, MacDialog).

### Implementation Details in Go

In Go, the Abstract Factory pattern leverages interfaces and struct composition due to the language’s lack of classical inheritance.

**Key points**:

- **Interfaces for Abstraction**:
    Define abstract factory and product interfaces to enforce contracts (e.g., GUIFactory, Button).
- **Structs for Concrete Implementations**:
    Implement factory and product interfaces using structs (e.g., WindowsFactory, MacButton).
- **Package Organization**:
    Group related factories and products into packages (e.g., ui/windows, ui/mac).
- **Error Handling**:
    Use Go’s error return pattern if object creation can fail (e.g., unsupported OS).
- **Implicit Interface Satisfaction**:
    Go’s "duck typing" allows structs to implement interfaces without explicit declaration, simplifying extensibility.

**Example Flow**:

1. A client initializes a concrete factory (e.g., mac.NewFactory()).
2. The factory creates compatible products (e.g., mac.Button, mac.Dialog).
3. Client code interacts with products via abstract interfaces.

### Comparison with Factory Method

|Aspect | Abstract Factory | Factory Method |
|---|---|---|
|Focus  | Families of related products| Single product creation |
|Complexity | Multiple factory methods| Simpler, one method |
|Extensibility | Add new product families|Add new products |
|Use Case | Cross-platform UI kits, theme systems| Dynamic object creation |
