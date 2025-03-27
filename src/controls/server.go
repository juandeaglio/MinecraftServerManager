package controls

type Controls interface {
	Start() bool
	Stop() bool
	Status() *Status
	IsStarted() bool
}

type Status struct {
	Players int
}
