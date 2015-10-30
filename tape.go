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
	if t.Head >= len(t.Symbol) {
		return
	}

	//fmt.Println("index:", t.Head)
	t.Symbol[t.Head] = modifiedSym

	if directToRight {
		t.moveRight()
	} else {
		t.moveLeft()
	}
}

//Return if the tape is run to the end
func (t *Tape) EndInput() bool {
	return t.Head == len(t.Symbol) //when head exceed slice size, it is end
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
