package tm

type Tape struct {
	Symbol []string
	Head   int
}

func (t *Tape) moveLeft() {
	if t.Head > 0 {
		t.Head--
	}
}

func (t *Tape) moveRight() {
	t.Head++
}

//Run tape head move and modify value
func (t *Tape) DoOption(modifiedSym string, directToRight bool) {
	if t.Head < len(t.Symbol) {
		t.Symbol[t.Head] = modifiedSym
	} else { //Append new Blank if goes to Right
		t.Symbol = append(t.Symbol, "B")
	}

	if directToRight {
		t.moveRight()
	} else {
		t.moveLeft()
	}
}

//Return if the tape is run to the end
func (t *Tape) EndInput() bool {
	return t.Head == len(t.Symbol)-1
}

//Read one symbol from tape
func (t *Tape) ReadSymbol() string {
	return t.Symbol[t.Head]
}

func NewTape(Symbols []string) *Tape {
	newT := new(Tape)
	newT.Symbol = Symbols
	return newT
}
