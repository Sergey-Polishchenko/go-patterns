## Builder

The **Builder** pattern is a **creational** design pattern that separates the construction of a complex object from its representation.
It enables the creation of objects with different configurations using the same construction process while avoiding "telescoping constructors" with excessive parameters.

### Characteristics

- **Step-by-Step Construction**:  
  Breaks object creation into sequential steps, allowing precise control over each stage of assembly.
- **Director Coordination**:  
  Uses a dedicated `Director` class to enforce standardized construction workflows.
- **Representation Flexibility**:  
  Enables creating different object variations using the same construction process.
- **Completion Guarantee**:  
  Ensures objects are fully initialized before use through atomic construction.

### Advantages

- **Configurable Complexity**:  
  Handles objects requiring multiple initialization steps or optional components.
- **Construction Reusability**:  
  Reuse construction algorithms across different object representations.
- **Parameter Encapsulation**:  
  Eliminates long parameter lists in constructors.
- **Immutable Products**:  
  Facilitates thread-safe object creation by finalizing construction before exposure.

### Disadvantages

- **Structural Overhead**:  
  Requires creating multiple classes (Builder, Director, Product) even for simple objects.
- **Coupling Risks**:  
  Tight integration between concrete builders and their products.
- **Runtime Inflexibility**:  
  Construction steps are fixed at design time, limiting dynamic adaptation.

### When to Use?

- When objects require complex multi-step initialization (e.g., documents with headers/body/footers).
- To create different object configurations using the same construction workflow.
- When object construction parameters exceed 4-5 arguments.
- Common Use Cases:
  + Query builders for databases
  + Meal customization systems (burgers with optional ingredients)
  + Configuration object assembly
  + Game entity spawning (characters with modular equipment)

### Required Components

1. **Builder Interface**:  
   Declares construction methods for all possible object components.

2. **Concrete Builders**:  
   Implement specific construction logic and provide product retrieval methods.

3. **Director**:  
   Orchestrates construction sequences using builder interface methods.

4. **Product**:  
   The complex object being constructed, typically immutable after assembly.

### Implementation Principles in Go

While Go lacks classical OOP features, the Builder pattern can be effectively implemented using interface-driven design and struct composition.

**Key considerations**:

- **Interface Contracts**:  
  Define clear builder interfaces specifying mandatory construction steps.

- **Director Neutrality**:  
  Keep directors unaware of concrete builder/product implementations.

- **Product Immutability**:  
  Design products to be unmodifiable after construction for thread safety.

- **Optional Step Handling**:  
  Implement optional construction steps through separate interfaces.

- **Validation Hooks**:  
  Add validation logic in concrete builders before product retrieval.

**Implementation Workflow**:

1. Client selects appropriate concrete builder
2. Client optionally configures director with builder
3. Director executes standardized construction sequence
4. Client retrieves finalized product from builder
