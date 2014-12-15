package qthulhu

import "github.com/hashicorp/raft"

type Server struct {
	conf *Config
	raft *raft.Raft
}

func NewServer(conf *Config) (*Server, error) {
	s := &Server{conf: conf}

	return s, nil
}

func (s *Server) Shutdown() error {
	return nil
}
