package domain

import "fmt"


type Domain struct {
	Name string
}

func Foo() string { return fmt.Sprintf("%s", "foo") }

func New(name string) Domain {
	return Domain{
		Name: name,
	}
}

func (d Domain) String() string {
	return fmt.Sprintf("Name: %s", d.Name)
}
