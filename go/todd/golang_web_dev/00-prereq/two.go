package main

import "fmt"

type person struct {
	name string
	age  int
}
type agent struct {
	name string
	age  int
	id   int
}

func (r *person) speak() (spoken string) {
	fmt.Printf("Hey, Im %s and im %v years old\n", r.name, r.age)
	return
}
func (r *agent) speak() (spoken string) {
	// fmt.Printf("Hey, Im agent %s and im %v years old and my id is %v", r.credentials.name, r.credentials.age, r.id)
	fmt.Printf("Hey, Im agent %s and im %v years old and my id is %v\n", r.name, r.age, r.id)
	return
}

func main() {
	person1 := person{name: "jack", age: 20}
	agent1 := agent{name: "Smith", age: 40, id: 22004}
	fmt.Println(person1.name)
	fmt.Println(agent1.name)
	person1.speak()
	agent1.speak()

}
