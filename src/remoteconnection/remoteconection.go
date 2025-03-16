package remoteconnection

type RemoteConnection interface {
	PollServer() StatusResponse
	IsAvailable(portNum int) bool
}
