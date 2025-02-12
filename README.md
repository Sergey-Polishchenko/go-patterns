# Go Design Patterns 

This repository is a personal collection of 22 classic design patterns implemented in Go.
It serves as a space for exploring, experimenting, and documenting these patterns with code examples.
While the project is open for reference, its primary purpose is for personal study and practice.

---

## Patterns

### [Creational](creational)

Creational design patterns are concerned with the way objects are created.
They help to make a system independent of how its objects are created, composed, and represented.
These patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code.
They abstract the instantiation process, making it easier to introduce new types of objects without changing the existing code.

- **Class-level patterns**
    Class-level patterns change the class of the created object through inheritance.
    These patterns define how objects are instantiated by delegating the responsibility to subclasses.

- **Object-level patterns**
    Object-level patterns create new objects by using other objects.
    These patterns focus on delegating the instantiation process to another object, promoting flexibility and decoupling.

---

### List of Patterns

- **Class-level patterns**
    + [Factory Method](./creational/factory_method)

- **Object-level patterns**
    + [Abstract Factory](./creational/abstract_factory)
    + [Builder](./creational/builder)
    + [Prototype](./creational/prototype)
    + [Singleton](./creational/singleton)


---

### [Structural](structural)

Structural design patterns deal with object composition or the way objects are structured to form larger structures.
They help to ensure that if one part of a system changes, the entire structure doesn't need to change.
These patterns focus on how classes and objects are composed to form larger structures, simplifying the design by identifying relationships between entities.
Structural patterns often use inheritance to compose interfaces or implementations.

- **Class-level patterns**
    Class-level patterns describe interactions between classes and their subclasses.
    These relationships are expressed through inheritance and class implementation.
    Here, the base class defines the interface, while subclasses provide the implementation.

- **Object-level patterns**
    Object-level patterns describe interactions between objects.
    These relationships are expressed through [associations](#association), [aggregation](#aggregation), and [composition](#composition).
    Structures are built by combining objects of certain classes, allowing for dynamic and flexible designs.

---

### List of Patterns

- **Class-level patterns**
    + [Adapter]()

- **Object-level patterns**
    + [Bridge]()
    + [Composite]()
    + [Decorator]()
    + [Facade]()
    + [Flyweight]()
    + [Proxy]()


---

### [Behavioral](behavioral)

Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects.
They focus on how objects interact and communicate with each other, ensuring that objects can work together in a flexible and maintainable way.
These patterns help to define the ways in which objects interact and distribute responsibility, making it easier to manage complex control flows and communication between objects.

- **Class-level patterns**
    Class-level patterns describe interactions between classes and their subclasses.
    These relationships are expressed through inheritance and class implementation.
    The base class defines the interface, while subclasses provide the implementation.

- **Object-level patterns**
    Object-level patterns describe interactions between objects.
    These relationships are expressed through [associations](#association), [aggregation](#aggregation), and [composition](#composition).
    Structures are built by combining objects of certain classes, enabling dynamic and flexible behavior.

---

### List of Patterns

- **Class-level patterns**
    + [Template Method]()

- **Object-level patterns**
    + [Chain of Responsibility]()
    + [Command]()
    + [Iterator]()
    + [Mediator]()
    + [Memento]()
    + [Observer]()
    + [State]()
    + [Strategy]()
    + [Visitor]()


---

## Theory Box

### **Association**

An association is a relationship where objects of two classes can refer to each other.
For example, a property of one class may hold an instance of another class.
This relationship can be bidirectional or unidirectional, meaning that one object can know about the other, or both objects can know about each other.

Example:
```go
package main

import "fmt"

// Engine represents Car engine.
type Engine struct {
    Power int
}

// Car represents car, which associated with engine.
type Car struct {
    Model  string
    Engine *Engine // Association with Engine
}

func main() {
    engine := &Engine{Power: 150}
    car := Car{Model: "Tesla Model S", Engine: engine}

    fmt.Printf("Car: %s, Engine Power: %d HP\n", car.Model, car.Engine.Power)
}
```

### **Aggregation**

Aggregation is a specialized form of association where one object acts as a container for other objects, but the lifetime of the contained objects does not depend on the lifetime of the container.
In other words, if the container is destroyed, the objects it contains are not affected.
For example, you can create an object and then pass it to a container object, either through a method or by directly assigning it to a property of the container.
When the container is deleted, the contained object remains unaffected and can interact with other containers.

Example:
```go
package main

import "fmt"

// Wheel represents the car wheel.
type Wheel struct {
    Diameter int
}

// Car represents car, which aggregate wheels.
type Car struct {
    Model  string
    Wheels []Wheel // Aggregation of wheels
}

func main() {
    wheels := []Wheel{
        {Diameter: 17},
        {Diameter: 17},
        {Diameter: 17},
        {Diameter: 17},
    }
    car := Car{Model: "BMW X5", Wheels: wheels}

    fmt.Printf("Car: %s, Wheels: %+v\n", car.Model, car.Wheels)
}
```

### **Composition**

Composition is a stronger form of aggregation where the contained objects cannot exist independently of the container.
If the container is destroyed, all its contained objects are also destroyed.
For example, if an object is created within a method of the container and assigned to a property of the container, it is not accessible from the outside.
Therefore, when the container is deleted, the contained object is also deleted, as there are no external references to it.

Example:
```go
package main

import "fmt"

// Heart represents the human heart.
type Heart struct {
    BeatsPerMinute int
}

// Human represents human, which consist of heart (Composition).
type Human struct {
    Name  string
    Heart Heart // Composition: Heart not exist without human
}

func main() {
    human := Human{
        Name:  "John",
        Heart: Heart{BeatsPerMinute: 72},
    }

    fmt.Printf("Human: %s, Heart Rate: %d BPM\n", human.Name, human.Heart.BeatsPerMinute)
}
```

---

## TODO

- [ ] **Implement all classic design patterns**
  - [ ] Creational patterns
  - [ ] Structural patterns
  - [ ] Behavioral patterns

- [ ] **Complete README files for each pattern**
  - [ ] Describe each pattern

- [ ] **Write tests for each pattern**
- [ ] **Add usage examples for each pattern**

---

## Acknowledgements

- [go-patterns by AlexanderGrom](https://github.com/AlexanderGrom/go-patterns)

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
