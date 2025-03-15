package systeminterface

type SystemInterface interface {
	PollExecutable() bool
}
