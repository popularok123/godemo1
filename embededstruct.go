package main

import "fmt"

type EmbededStruct struct{}

type describer interface {
	describe() string
}

func (es *EmbededStruct) Do() {
	co := container{
		base: base{
			num: 10,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describle())

	dd := co

	fmt.Println("describer:", dd.describle())

}

type base struct {
	num int
}

func (b base) describle() string {
	return fmt.Sprintf("num: %d", b.num)
}

type container struct {
	base
	str string
}
