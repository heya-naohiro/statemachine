package main

func main() {
	ds := NewDeviceState()
	ds.CheckGraph()

	machine := NewStateMachine(&ds)
	machine.Run(ds.RecieveInitialized, nil)

	machine.Run(ds.RecieveRPCRequest, nil)
	machine.Run(ds.ResponseRPCRequest, nil)

	// for example duplicate
	machine.Run(ds.RecieveRPCRequest, nil)
	machine.Run(ds.ResponseRPCRequest, nil)

}
