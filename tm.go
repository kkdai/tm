package tm

const ( // iota is reset to 0
	MoveLeft  = iota //0
	MoveRight = iota //1
)

func NewTM() *TM {
	newTM := new(TM)
	newTM.States = make(map[string]bool)
	newTM.FinalStates = make(map[string]bool)
	newTM.Inputs = make(map[string]bool)
	newTM.Configs = make(map[ConfigIn]ConfigOut)
	return newTM
}

type ConfigIn struct {
	SrcState string
	Input    string
}

type ConfigOut struct {
	DstState    string
	ModifiedVal string
	TapeMove    int
}

type TM struct {
	Input        *Tape
	States       map[string]bool
	FinalStates  map[string]bool
	Inputs       map[string]bool
	Configs      map[ConfigIn]ConfigOut
	CurrentState string
}

// Input State and declare if it is final state
func (t *TM) InputState(state string, isFinal bool) {
	//First input state will be init state
	if t.CurrentState == "" {
		t.CurrentState = state
	}

	//Add newState
	t.States[state] = true
	if isFinal {
		t.FinalStates[state] = true
	}
}

// Input turing machine tape data
func (t *TM) InputTape(tracks ...string) {
	newTape := NewTape(tracks)
	t.Input = newTape
}

//Input config
// InputConfig parameter as follow:
// - SourceState,
// - Input
// - Modified Value
// - DestinationState
// - Tape Head Move Direction
func (t *TM) InputConfig(srcState string, input string, modifiedVal string, dstState string, tapeMove int) {
	//Check state
	if _, exist := t.States[srcState]; !exist {
		return
	}

	if _, exist := t.States[dstState]; !exist {
		return
	}

	//Add Input
	t.Inputs[input] = true

	//Add config
	newConfigIn := ConfigIn{SrcState: srcState, Input: input}
	newConfigOut := ConfigOut{ModifiedVal: modifiedVal, TapeMove: tapeMove, DstState: dstState}
	t.Configs[newConfigIn] = newConfigOut
}

// Step will read a input and run single step, return if it is accept
func (t *TM) Step() bool {
	input := t.Input.ReadSymbol()
	inConfig := ConfigIn{SrcState: t.CurrentState, Input: input}

	if newOut, exist := t.Configs[inConfig]; !exist {
		return false
	} else {
		t.Input.DoOption(newOut.ModifiedVal, newOut.TapeMove == MoveRight)
		t.CurrentState = newOut.DstState
	}

	if _, exist := t.FinalStates[t.CurrentState]; exist {
		return true
	}

	return false
}

// Run turing machine all tape data to the end
func (t *TM) Run() bool {
	var latestResult bool
	for !t.Input.EndInput() {
		//fmt.Println("run  index=", t.Input.Head)
		latestResult = t.Step()
	}
	return latestResult
}

func (t *TM) ExportTape() []string {
	return t.Input.Symbol
}
