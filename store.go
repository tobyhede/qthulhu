package qthulhu

import (
	"log"

	"github.com/hashicorp/raft"
)

type Store interface {
}

type PartitionStore struct {
	path   string
	rstore *RocksDBStore
	logger *log.Logger
}

func NewPartitionStore(path string, logger *log.Logger) (*PartitionStore, error) {
	logger.Print("Connecting Store ... ", path)
	rstore := NewRocksDBStore(path)
	return &PartitionStore{rstore: rstore, logger: logger}, nil
}

func (s *PartitionStore) Close() error {
	return s.rstore.Close()
}

func (s *PartitionStore) FirstIndex() (uint64, error) {
	return s.rstore.FirstKey()
}

func (s *PartitionStore) LastIndex() (uint64, error) {
	return s.rstore.LastKey()
}

func (s *PartitionStore) GetLog(index uint64, log *raft.Log) error {
	v, err := s.rstore.Get(uint64ToBytes(index))
	if err != nil {
		return err
	}
	return decode(v, log)
}

func (s *PartitionStore) StoreLog(log *raft.Log) error {
	return s.StoreLogs([]*raft.Log{log})
}

func (s *PartitionStore) StoreLogs(logs []*raft.Log) error {
	b := s.rstore.StartBatch()

	for _, l := range logs {
		k := uint64ToBytes(l.Index)
		d, _ := encode(l)
		b.Put(k, d)
	}

	err := s.rstore.WriteAndCloseBatch(b)
	return err
}

func (s *PartitionStore) DeleteRange(min, max uint64) error {
	b := s.rstore.StartBatch()

	for i := min; i <= max; i++ {
		k := uint64ToBytes(i)
		b.Delete(k)
	}

	return s.rstore.WriteAndCloseBatch(b)
}

func (s *PartitionStore) Set(key []byte, val []byte) error {
	return s.rstore.Put(key, val)
}

func (s *PartitionStore) Get(key []byte) ([]byte, error) {
	return s.rstore.Get(key)
}

func (s *PartitionStore) SetUint64(key []byte, val uint64) error {
	s.logger.Print(key)
	s.logger.Print(val)
	return s.rstore.Put(key, uint64ToBytes(val))
}

func (s *PartitionStore) GetUint64(key []byte) (uint64, error) {
	v, err := s.rstore.Get(key)
	if err != nil {
		return 0, err
	}
	return bytesToUint64(v), err
}
