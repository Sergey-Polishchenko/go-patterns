# Go Design Patterns 

This repository is a personal collection of 22 classic design patterns implemented in Go.
It serves as a space for exploring, experimenting, and documenting these patterns with code examples.
While the project is open for reference, its primary purpose is for personal study and practice.

---

## Patterns

<INSERT:creational>
---

<INSERT:structural>
---

<INSERT:behavioral>
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
