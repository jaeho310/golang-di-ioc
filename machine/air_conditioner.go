package machine

type AirConditional struct{}

func (AirConditional) New() *AirConditional {
	return &AirConditional{}
}

func (airConditional *AirConditional) TurnOn() (string, error) {
	return "에어콘을 킵니다", nil
}
