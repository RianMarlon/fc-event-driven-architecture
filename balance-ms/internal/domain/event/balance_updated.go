package event

import "time"

type BalanceUdated struct {
	Name    string
	Payload interface{}
}

func NewBalanceUpdated() *BalanceUdated {
	return &BalanceUdated{
		Name: "BalanceUpdated",
	}
}

func (e *BalanceUdated) GetName() string {
	return e.Name
}

func (e *BalanceUdated) GetPayload() interface{} {
	return e.Payload
}

func (e *BalanceUdated) GetDateTime() time.Time {
	return time.Now()
}

func (e *BalanceUdated) SetPayload(payload interface{}) {
	e.Payload = payload
}
