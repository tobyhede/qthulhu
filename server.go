package qthulhu

import "github.com/hashicorp/raft"

type Server struct {
	conf *Config
	raft *raft.Raft
}

func NewServer(conf *Config) (*Server, error) {
	r, err := NewRaft(conf)
	if err != nil {

	}
	s := &Server{conf: conf, raft: r}
	return s, err
}

func (s *Server) Shutdown() error {
	return nil
}
