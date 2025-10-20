package main

import "fmt"

type TypeConvert struct{}

func (tc *TypeConvert) Do() {

	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

	rs := dotest()

	fmt.Println(rs)
}

func functionOfSomeType() interface{} {
	i := 42
	return &i
}

func functionOfSomeType1() interface{} {
	a := "abc"
	return a
}

type Stringer interface {
	String() string
}

func dotest() string {
	var value interface{} // Value provided by caller.
	value = functionOfSomeType1()
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	default:
		return fmt.Sprintf("%v", value)
	}
}

func dootherway() string {
	var value interface{} // Value provided by caller.
	value = functionOfSomeType1()
	if str, ok := value.(string); ok {
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	} else {
		return ""
	}
}
