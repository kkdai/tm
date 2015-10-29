package tm_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/tm"
)

func TestBasicTM(t *testing.T) {
	nTM := NewTM()
	nTM.InputState("0", false)
	nTM.InputState("1", false)
	nTM.InputState("2", true)

	nTM.InputConfig("0", "1", "1", "1", MoveRight)
	nTM.InputConfig("1", "1", "1", "2", MoveRight)

	nTM.InputTape("1", "1")

	if nTM.Step() {
		t.Errorf("TM: Should not stop")
	}

	if !nTM.Step() {
		t.Errorf("TM: Should be accept")
	}
}

//Rewrite all tape 0 to 1
func TestModifyTM(t *testing.T) {
	nTM := NewTM()
	nTM.InputState("0", false)
	nTM.InputState("1", true)

	nTM.InputConfig("0", "1", "1", "1", MoveRight)
	nTM.InputConfig("0", "0", "1", "0", MoveLeft)
	nTM.InputConfig("1", "0", "1", "0", MoveLeft)
	nTM.InputConfig("1", "1", "1", "1", MoveRight)

	nTM.InputTape("0", "0", "1", "1", "0", "0", "0")

	if !nTM.Run() {
		t.Errorf("TM: Cannot reach to end")
	}
	fmt.Println("New Tape:=", nTM.ExportTape(), nTM.CurrentState)
}
