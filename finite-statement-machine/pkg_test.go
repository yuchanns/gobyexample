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

func TestNodeFs(t *testing.T) {
	table := []struct {
		CurrentState transState
		Flag         bool
		Node         *TMonitorStatus
		Expect       transState
	}{
		{TransRun, false, &TMonitorStatus{TransRun, 9, true}, TransCritical},
		{TransCritical, true, &TMonitorStatus{TransCritical, 9, true}, TransRun},
		{TransCritical, false, &TMonitorStatus{TransCritical, 9, false}, TransDown},
	}

	fs := InitNodeFs()

	for _, item := range table {
		node, err := fs.Transfer(item.CurrentState, item.Node, item.Flag)
		if err != nil {
			t.Errorf("failed transfer: %+v", err)
		}

		n, ok := node.(*TMonitorStatus)
		if !ok {
			t.Errorf("failed to assert node: %v", node)
		}
		if n.CurrentState != item.Expect {
			t.Errorf("next status not as expected: %d", n.CurrentState)
		}
	}

}
