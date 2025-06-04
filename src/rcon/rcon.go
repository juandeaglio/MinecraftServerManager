package rcon

import (
	"log"
	"time"

	mcrcon "github.com/gorcon/rcon"
)

type RCONAdapter interface {
	GetStatus() *Status
}

type Status struct {
	Players int
	Online  bool
}

func NewStubRCONAdapter() *StubRCONAdapter {
	return &StubRCONAdapter{}
}

type StubRCONAdapter struct {
}

func (s *StubRCONAdapter) GetStatus() *Status {
	return &Status{
		Players: 0,
		Online:  true,
	}
}

var _ RCONAdapter = (*StubRCONAdapter)(nil)

type MinecraftRCONAdapter struct {
	timeout time.Duration
}

func (m *MinecraftRCONAdapter) GetStatus() *Status {
	done := make(chan struct{})
	var conn *mcrcon.Conn
	var err error

	go func() {
		conn, err = mcrcon.Dial("localhost:25575", "password")
		close(done)
	}()

	select {
	case <-done:
		if err != nil {
			return nil
		}
	case <-time.After(m.timeout):
		return nil
	}

	defer func(conn *mcrcon.Conn) {
		_ = conn.Close()
	}(conn)

	_, err = conn.Execute("help")
	if err != nil {
		log.Fatal(err)
	}

	return &Status{
		Players: 1,
		Online:  true,
	}
}

func (m *MinecraftRCONAdapter) WithTimeout(timeout time.Duration) *MinecraftRCONAdapter {
	m.timeout = timeout
	return m
}

func NewMinecraftRCONAdapter() *MinecraftRCONAdapter {
	return &MinecraftRCONAdapter{}
}

var _ RCONAdapter = (*MinecraftRCONAdapter)(nil)
