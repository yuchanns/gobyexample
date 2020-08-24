package finite_statement_machine

import "testing"

func TestDoorOpenAndClose(t *testing.T) {
	currentState := Opened
	isAccept := false
	for i := 0; i < 10; i++ {
		currentState, isAccept = DoorOpenAndClose(currentState)
		if !isAccept {
			break
		}
	}
}

func TestNewDoorMachine(t *testing.T) {
	dm := NewDoorMachine(Opened, map[State]func() (State, bool){
		Opened: func() (State, bool) {
			return Closed, true
		},
		Closed: func() (State, bool) {
			return Locked, true
		},
		Locked: func() (State, bool) {
			return Unlocked, true
		},
		Unlocked: func() (State, bool) {
			return Opened, true
		},
	})

	for i := 0; i < 10; i++ {
		_, isAccept := dm.Action()
		if !isAccept {
			break
		}
	}
}
