package tm_test

import (
	"testing"

	. "github.com/kkdai/tm"
)

func TestTape(t *testing.T) {
	tapeData := []string{"1", "0", "1", "1"}
	newT := NewTape(tapeData)

	//Test readsymbol
	if newT.ReadSymbol() != "1" {
		t.Errorf("ReadSymbol: expect %s, get %s \n", "1", newT.ReadSymbol())
	}

	//Basic op
	newT.DoOption("B", true)
	if newT.ReadSymbol() != "0" {
		t.Errorf("ReadSymbol: expect %s, get %s \n", "0", newT.ReadSymbol())
	}

	//Test move left
	newT.DoOption(newT.ReadSymbol(), false)
	if newT.ReadSymbol() != "B" {
		t.Errorf("ReadSymbol: expect %s, get %s \n", "B", newT.ReadSymbol())
	}
}
