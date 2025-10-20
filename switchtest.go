package main

import "fmt"

type SwitchTest struct{}

const (
	Enone = iota
	Eio
	Einval
)

func (st *SwitchTest) Do() {
	i := 1
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	_ = a
	fmt.Print("Write ", i, " as ")
	on := true
	switch i {
	case 1:
		if on {
			fmt.Println("break out switch")
			break
		}
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
