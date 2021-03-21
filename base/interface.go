package base

import "fmt"

type Animal interface {
	Say() string
}

type Dog struct{}

func (d Dog) Say() string {
	word := "Dog!"
	fmt.Println(word)
	return word
}

type Cat struct{}

func (d Cat) Say() string {
	word := "Cat!"
	fmt.Println(word)
	return word
}

func Test() {
	t := []Animal{&Dog{}, Cat{}}
	for _, tt := range t {
		tt.Say()
	}
}
