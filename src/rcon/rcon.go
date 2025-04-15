package rcon

type RCONAdapter interface {
	GetStatus() *Status
}

type Status struct {
	Players int
	Online  bool
}
