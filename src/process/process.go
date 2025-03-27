package process

type Process interface {
	Start() error
	Stop() error
	PID() int
}
