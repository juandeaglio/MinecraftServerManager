package remoteconnection

type RemoteConnection interface {
	PollServer() StatusResponse
	IsAvailable() bool
}
