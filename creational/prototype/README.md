## Prototype

The Prototype pattern is a creational design pattern that allows creating new objects by copying an existing instance (prototype).
This pattern is useful when object creation is costly or complex, and you need to create variations of objects dynamically without coupling code to their concrete classes.

### Characteristics

- **Cloning Mechanism**:
    Objects can create copies of themselves via a Clone() method, avoiding the need for reinitialization.
- **Prototype Registry (optional)**:
    A registry to store frequently used prototypes, enabling quick access and cloning.
- **Flexibility**:
    New objects can be created at runtime by modifying cloned instances.
- **Reduced Subclassing**:
    Avoids the need for subclassing to create object variations.

### Advantages:

- Simplifies object creation for complex initialization processes.
- Enables dynamic object configuration at runtime.
- Reduces the number of classes compared to factory-based approaches.
- Supports adding and removing objects at runtime via a prototype registry.

### Disadvantages:

- Deep cloning can be complex for objects with circular references or complex structures.
- Cloning might violate encapsulation if objects have private fields.
- Requires careful implementation to ensure clones are independent of the original.

### When to Use?

- When creating an object is more expensive than copying an existing one.
- When system should be independent of how its objects are created and composed.
- When classes are instantiated at runtime or need to be dynamically loaded.
- Common use cases:
    + Game development (spawning enemies/items with similar properties),
    + Configuration templates,
    + Caching mechanisms,
    + Avoiding repetitive initialization (e.g., database connections).

### Required for Implementation

- **Prototype Interface**:
    Declares the Clone() method that all concrete prototypes must implement.
- **Concrete Prototypes**:
    Classes (e.g., ConcretePrototypeA, ConcretePrototypeB) that implement the cloning logic.
- **Client Code**:
    Creates new objects by cloning prototypes, optionally using a registry to manage prototypes.

### Implementation Details in Go

In Go, the Prototype pattern leverages interfaces to define the Clone() method.
Concrete types implement cloning by creating copies of themselves, either through shallow or deep copying.
Unlike class-based languages, Go uses structs and interface embedding to achieve this pattern.
The sync package may be used for thread-safe prototype registration if needed.
