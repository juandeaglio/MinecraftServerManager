package processtest_test

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"testing"

	"minecraftremote/src/process"

	"github.com/stretchr/testify/assert"
)

type brokenOsOperations struct{}

func (b *brokenOsOperations) FindProcess(pid int) (*os.Process, error) {
	return nil, fmt.Errorf("broken")
}

func (b *brokenOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return fmt.Errorf("broken")
}

func (b *brokenOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	return exec.Command(program, args...)
}

func (b *brokenOsOperations) SetSysProcAttr(cmd *exec.Cmd) {}

func (b *brokenOsOperations) StartCmd(cmd *exec.Cmd) error {
	return fmt.Errorf("broken")
}

func (b *brokenOsOperations) KillProcess(process *os.Process) error {
	return fmt.Errorf("broken")
}

func TestProcess(t *testing.T) {
	fakeProcess := process.NewProcess(&process.FakeOsOperations{}, "fake", "args")
	fakeProcess.Start()
	assert.Truef(t, fakeProcess.Started(), "Process failed to start.")

	fakeProcess.Stop()
	assert.Falsef(t, fakeProcess.Started(), "Process is still marked as started after stopping")
}

func TestProcessError(t *testing.T) {
	fakeProcess := process.NewProcess(&brokenOsOperations{}, "fake", "args")
	err := fakeProcess.Start()
	assert.Error(t, err)
}
