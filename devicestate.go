package main

import (
	"fmt"
	"time"

	"github.com/bykof/stateful"
	"github.com/bykof/stateful/statefulGraph"
)

var (
	BEGIN         = stateful.DefaultState("BEGIN")
	RPC_WAIT      = stateful.DefaultState("RPC_WAIT")
	RPC_REQUESTED = stateful.DefaultState("RPC_REQUESTED")
)

type DeviceState struct {
	state stateful.State
	log   string
}

func NewDeviceState() DeviceState {
	return DeviceState{
		state: BEGIN,
		log:   "",
	}
}

func (s *DeviceState) SetState(state stateful.State) error {
	s.state = state
	return nil
}

func (s *DeviceState) RecieveInitialized(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Println("RecieveInitialized")
	return RPC_WAIT, nil
}

func (s *DeviceState) RecieveRPCRequest(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Println("Hello RPC")

	time.Sleep(time.Second)
	return RPC_REQUESTED, nil
}

func (s *DeviceState) ResponseRPCRequest(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Println("End RPC")
	return RPC_WAIT, nil

}

func (s *DeviceState) State() stateful.State {
	return s.state
}

func (s *DeviceState) CheckGraph() {
	//check graph
	m := NewStateMachine(s)
	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: m}
	_ = stateMachineGraph.DrawGraph()
}

func NewStateMachine(s *DeviceState) stateful.StateMachine {
	stateMachine := stateful.StateMachine{
		StatefulObject: s,
	}
	stateMachine.AddTransition(s.RecieveInitialized, stateful.States{BEGIN}, stateful.States{RPC_WAIT})
	stateMachine.AddTransition(s.RecieveRPCRequest, stateful.States{RPC_WAIT}, stateful.States{RPC_REQUESTED})
	stateMachine.AddTransition(s.ResponseRPCRequest, stateful.States{RPC_REQUESTED}, stateful.States{RPC_WAIT})

	return stateMachine
}
