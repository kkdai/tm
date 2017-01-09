package main

import (
	"fmt"

	. "github.com/kkdai/tm"
)

func main() {
	nTM := NewTM()

	//Input State and declare if it is final state
	nTM.InputState("0", false)
	nTM.InputState("1", true)

	//Input config
	// InputConfig parameter as follow:
	//  - SourceState,
	// - Input
	// - Modified Value
	// - DestinationState
	// - Tape Head Move Direction
	nTM.InputConfig("0", "1", "1", "1", MoveRight)
	nTM.InputConfig("0", "0", "1", "0", MoveLeft)
	nTM.InputConfig("1", "0", "1", "0", MoveLeft)
	nTM.InputConfig("1", "1", "1", "1", MoveRight)

	//Input tape data
	nTM.InputTape("0", "0", "1", "1", "0", "0", "0")

	//Run TM to the finish (if exist)
	nTM.Run()

	fmt.Println("New Tape:=", nTM.ExportTape())
}
