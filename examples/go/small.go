package example

import "fmt"

type Person struct {
	name string
	mom  *Person
}

func NewPerson(name string, mom *Person) Person {
	return Person{name: name, mom: mom}
}

func (self *Person) GetName() string {
	return self.name
}

func (self *Person) GetMom() *Person {
	return self.mom
}

var p = NewPerson("foo", nil)
var _ = fmt.Println(p)

func f(fmt Person) string {
	return fmt.name // `name` should be a field because fmt shadows fmt
}
