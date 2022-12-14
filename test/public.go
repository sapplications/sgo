package test

import (
	"fmt"

	"github.com/sapplications/sgo/app"
)

type Runner interface {
	Run()
}

type RunnerImpl struct{}

type Item1 struct {
	Int1      int
	Float1    float32
	Field1    Field1
	Field1V2  Field1
	Field2    Field2
	Field3    Field3
	Runner    Runner
	Logger    app.Logger
	Hello     func(string)
	EmptyFunc func()
}

type Field1 struct{}

type Field2 struct {
	Name string
}

type Field3 struct {
	Field Field1
}

func NewField1() Field1 {
	return Field1{}
}

func NewField1V2(name string, value string) Field1 {
	return Field1{}
}

func NewField2(name string) Field2 {
	return Field2{name}
}

func NewField3(field Field1) Field3 {
	return Field3{field}
}

func Hello(name string) {
	fmt.Printf("Hello %s!", name)
}

func EmptyFunc() {

}

func (r *RunnerImpl) Run() {

}

func (i *Item1) Execute() {

}
