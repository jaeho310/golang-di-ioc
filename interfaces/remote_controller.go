package interfaces

type RemoteController struct {
	machineImpl Machine
}

type Machine interface {
	TurnOn() (string, error)
}

func (RemoteController) New(machine Machine) *RemoteController {
	return &RemoteController{machine}
}

func (remoteController *RemoteController) TrunOnMachine() (string, error) {
	return remoteController.machineImpl.TurnOn()
}
