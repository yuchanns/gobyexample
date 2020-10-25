package finite_statement_machine

import "errors"

type transState int

const (
	TransRun transState = iota + 1
	TransCritical
	TransDown
)

type BinResult struct {
	TrueState  transState
	FalseState transState
}

type Action func(currentState, stateSets interface{}, others ...interface{}) (interface{}, error)

type MapBinResult map[transState]*BinResult

type FiniteState struct {
	action    Action
	stateSets interface{}
}

func NewFiniteState(stateSets interface{}, action Action) *FiniteState {
	return &FiniteState{
		action:    action,
		stateSets: stateSets,
	}
}

func (fs *FiniteState) Transfer(currentState interface{}, others ...interface{}) (interface{}, error) {
	return fs.action(currentState, fs.stateSets, others...)
}

type TMonitorStatus struct {
	CurrentState transState
	Count        int
	BinState     bool
}

func InitNodeFs() *FiniteState {
	var transTable = MapBinResult{
		TransRun: &BinResult{
			FalseState: TransCritical,
		},
		TransCritical: &BinResult{
			TrueState:  TransRun,
			FalseState: TransDown,
		},
		TransDown: &BinResult{
			TrueState: TransRun,
		},
	}

	var action = func(currentStateInterface, stateSetsInterface interface{}, others ...interface{}) (interface{}, error) {
		currentState, ok := currentStateInterface.(transState)
		if !ok {
			return nil, errors.New("failed to assert currentState")
		}
		stateSets, ok := stateSetsInterface.(MapBinResult)
		if !ok {
			return nil, errors.New("failed to assert stateSets")
		}
		node, ok := others[0].(*TMonitorStatus)
		if !ok {
			return nil, errors.New("failed to assert node")
		}
		boolResult, ok := others[1].(bool)
		if !ok {
			return nil, errors.New("failed to assert flag")
		}

		node.BinState = boolResult

		if node.CurrentState == currentState {
			node.Count++
		} else {
			node.CurrentState = currentState
			node.Count = 1
		}

		if node.Count >= 10 {
			transItem := stateSets[node.CurrentState]
			if node.BinState {
				node.CurrentState = transItem.TrueState
			} else {
				node.CurrentState = transItem.FalseState

			}

			node.Count = 0
		}

		return node, nil
	}

	return NewFiniteState(transTable, action)
}
