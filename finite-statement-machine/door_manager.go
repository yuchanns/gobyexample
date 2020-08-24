package finite_statement_machine

import "fmt"

type State int

const (
	Opened State = iota
	Closed
	Locked
	Unlocked
)

var stateDesc = map[State]string{
	Opened:   "开启",
	Closed:   "关闭",
	Locked:   "锁定",
	Unlocked: "解锁",
}

func (s State) String() string {
	if desc, ok := stateDesc[s]; ok {
		return desc
	}

	return "invalid state"
}

// simple two state function
func DoorOpenAndClose(state State) (State, bool) {
	var nextState State
	isAccept := false
	if state == Opened {
		nextState = Closed
		isAccept = true
	} else if state == Closed {
		nextState = Opened
		isAccept = true
	}

	fmt.Printf("current state is %s, the next state is %s\n", state, nextState)

	return nextState, isAccept
}

// multiple state machine
type DoorMachine struct {
	currentState State
	actions      map[State]func() (State, bool)
}

func NewDoorMachine(
	current State,
	actions map[State]func() (State, bool),
) *DoorMachine {
	return &DoorMachine{
		currentState: current,
		actions:      actions,
	}
}
func (m *DoorMachine) Action() (State, bool) {
	action, ok := m.actions[m.currentState]
	if !ok {
		return m.currentState, false
	}

	nextState, isAccept := action()

	fmt.Printf("current state is %s, the next state is %s\n", m.currentState, nextState)
	m.currentState = nextState

	return nextState, isAccept
}
