// enums are a special case of sum types
// @todo read about sum types -> https://en.wikipedia.org/wiki/Algebraic_data_type

// Go does not have enum type as a language feature but simple to implement using existing language idioms

package main

import "fmt"

// ServerState enum has underlying type int
type ServerState int

// possible values of ServerState are defined as constants
// iota generates successive constant values automatically eg 0, 1, 2 etc
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// implement fmt.Stringer interface so that values of ServerState can be printed or converted to strings
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) Stringer() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(ss ServerState) ServerState {
	switch ss {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unkown error: %s", ss))

	}
}
