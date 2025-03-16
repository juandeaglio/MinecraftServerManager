package remoteconnection

type RCONResponse struct {
	totalPlayers int
}

func (r *RCONResponse) TotalPlayers() int {
	return r.totalPlayers
}
