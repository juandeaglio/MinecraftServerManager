package controls

type Controls interface {
	Start() bool
	Stop() bool
	Status() *Status
}

type Status struct {
	Players int
}
