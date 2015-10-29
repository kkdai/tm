package tm_test

import (
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
