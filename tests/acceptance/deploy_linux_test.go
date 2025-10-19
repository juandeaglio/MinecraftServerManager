//go:build linux

package acceptance_test

import (
	"minecraftremote/tests/dsl"
	"testing"
)

func TestAcceptanceDeployServerToLinux(t *testing.T) {
	_ = `Given a server-hoster with a remote on Linux
		when I start the process via Remote
		then the server should be started.`
	_ = dsl.GivenALinuxRemote()

}
