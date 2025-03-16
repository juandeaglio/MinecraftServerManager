package remoteconnection

type RemoteConnection interface {
	PollServer() bool
	IsAvailable(portNum int) bool
}
