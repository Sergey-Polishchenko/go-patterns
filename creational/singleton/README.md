## Singleton

The **Singleton** pattern is a **creational** design pattern that ensures a class has only one instance and provides a global point of access to that instance.
It is useful in scenarios where you need to ensure that only one object of a particular class exists, such as managing a database connection, application settings, logging, etc.

### Characteristics 

- **Single Instance**:
    The Singleton class ensures that only one instance of the class exists in the system.
- **Global Access**:
    The instance of the Singleton is accessible globally through a static method (usually called getInstance()).
- **Lazy Initialization**:
    The instance is created only when it is first accessed, which helps save resources.
- **Thread Safety**:
    In multithreaded applications, the Singleton must be implemented in a way that prevents the creation of multiple instances.

### Advantages:

- Control over a single instance.
- Global access to the instance.
- Lazy initialization (instance is created only when needed).

### Disadvantages:

- Violates the Single Responsibility Principle (the class manages its own creation and behavior).
- Makes testing harder due to global state.
- Can lead to issues in multithreaded applications if not implemented correctly.

### When to Use?

- When there should be only one instance of a class in the system.
- When global access to that instance is required.
- When lazy initialization is important for resource efficiency.
- Common use cases:
    + database connections,
    + configuration management,
    + logging,
    + caching,
    + thread pools.

### Required for Implementation

To implement the Singleton pattern, a method is required that will control the creation of the class instance.
This method is usually called GetInstance.

**This method is responsible for the following actions**:

1. **Checking for an existing instance**:
    If an instance already exists, the method returns the existing instance.
2. **Creating an instance**:
    If no instance exists, the method creates it and stores it for future use.
3. **Returning the instance**:
    In any case, the method returns the single instance of the class.

This method is typically implemented as a static method so that it can be called without the need to create an object of the class.

### Implementation Details in Go

In Go, the Singleton pattern is implemented differently due to the lack of classes and static methods.
Instead, global variables and synchronization mechanisms like sync.Once are used to ensure thread safety and lazy initialization.
This guarantees that only one instance of the object is created and globally accessible.
