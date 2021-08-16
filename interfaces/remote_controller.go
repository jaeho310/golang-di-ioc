package interfaces

// 리모콘은 Machine이라는 의존성이 필요합니다.
// 외부에서 생성할때 주입해줍니다.
type RemoteController struct {
	machineImpl Machine
}

type Machine interface {
	TurnOn() (string, error)
}

func (RemoteController) New(machine Machine) *RemoteController {
	return &RemoteController{machine}
}

// 리모콘은 뭐를 키는지 몰라도 됩니다.
func (remoteController *RemoteController) TrunOnMachine() (string, error) {
	return remoteController.machineImpl.TurnOn()
}
