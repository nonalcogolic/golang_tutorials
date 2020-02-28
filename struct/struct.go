package main

import "fmt"

type Star struct {
	name          string
	distance      float64
	constellation string
}

type Base struct {
	name string
}

func (Base) new(name string) *Base {
	ptr := &Base{name}
	return ptr
}

type Derived struct {
	Base
	surname string
}

func sayName(class Base) {
	fmt.Println(class.name)
}

func (base Base) sayName() {
	fmt.Println(base.name)
}

func main() {
	aldebaran := Star{"Aldebaran", 65.3, "Taurus"}
	fmt.Println(aldebaran)
	aldebaran2 := Star{"Aldebaran", 65.3, "Taurus"}
	fmt.Println(aldebaran2)
	if aldebaran == aldebaran2 {
		fmt.Println("Star is the same")
	}

	vega := Star{
		name:          "Vega",
		distance:      25,
		constellation: "Lira",
	}
	fmt.Println(vega)

	var base Base = Base{"some name"}

	sayName(base)

	var derived Derived = Derived{Base{"Some other name"}, "surname"}
	fmt.Println(derived)

	sayName(derived.Base)

	base.sayName()
	derived.sayName()

	newBase := Base{}.new("value")
	fmt.Println(newBase)
} //declare receiver string in the module string
