package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

//go:generate stringer -type=ServerState

func (ss ServerState) String() string {
	return stateName[ss]
}

type EnumsInt struct{}

func (ei *EnumsInt) Do() {

	ns := transition(StateIdle)

	fmt.Println(ns)

	ns2 := transition(ns)

	fmt.Println(ns2)

}

func transition(state ServerState) ServerState {
	switch state {
	case StateIdle:
		return StateConnected
	case StateConnected:
		return StateError
	case StateError:
		return StateRetrying
	default:
		panic(fmt.Errorf("unknown state: %s", state))
	}
}
