package qthulhu

type Server struct {
	conf *Config
	raft *Raft
}

func NewServer(conf *Config) (*Server, error) {
	r, err := NewRaft(conf)
	if err != nil {

	}
	s := &Server{conf: conf, raft: r}
	return s, err
}

func (s *Server) Shutdown() error {
	s.raft.Close()
	return nil
}
