package main

import "fmt"

var doall []DoInterface

func RegisterDo(a DoInterface) {
	doall = append(doall, a)
}

func main() {
	for _, a := range doall {
		a.Do()
	}
}

func init() {
	fmt.Println("init start")
	RegisterDo(&ArraySlices{})
	RegisterDo(&MapStruct{})
	RegisterDo(&Variafun{})
	RegisterDo(&RangeUse{})
	RegisterDo(&StringRunes{})
	RegisterDo(&EnumsInt{})
	RegisterDo(&EmbededStruct{})
	RegisterDo(&GenericList{})
	RegisterDo(&GorountieTest{})
	RegisterDo(&ChannelTest{})
	RegisterDo(&SwitchTest{})
	RegisterDo(&TypeConvert{})
}
