package chainOfResponsibility

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}
