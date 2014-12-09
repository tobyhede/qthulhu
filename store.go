package qthulhu

import "github.com/hashicorp/raft"

type PartitionStore struct {
	path   string
	rstore *RocksDBStore
}

func NewPartitionStore(path string) (*PartitionStore, error) {
	rstore := NewRocksDBStore(path)
	return &PartitionStore{rstore: rstore}, nil
}

// func (m *PartitionStore) initialize() error {
// }

func (s *PartitionStore) Close() error {
	s.rstore.Close()
	return nil
}

func (s *PartitionStore) FirstIndex() (uint64, error) {
	return s.rstore.FirstKey()
}

func (s *PartitionStore) LastIndex() (uint64, error) {
	return s.rstore.LastKey()
}

func (s *PartitionStore) GetLog(index uint64, log *raft.Log) error {
	v, err := s.rstore.Get(index)
	if err != nil {
		//something
	}
	log.Data = v
	return err
}

func (s *PartitionStore) StoreLog(log *raft.Log) error {
	return s.StoreLogs([]*raft.Log{log})
}

func (s *PartitionStore) StoreLogs(logs []*raft.Log) error {
	s.rstore.PutBatch(logs)
	return nil
}

func (s *PartitionStore) DeleteRange(min, max uint64) error {
	return nil
}

// type StableStore interface {
//     Set(key []byte, val []byte) error
//     Get(key []byte) ([]byte, error)

//     SetUint64(key []byte, val uint64) error
//     GetUint64(key []byte) (uint64, error)
// }

// type SnapshotStore interface {
//     // Create is used to begin a snapshot at a given index and term,
//     // with the current peer set already encoded
//     Create(index, term uint64, peers []byte) (SnapshotSink, error)

//     // List is used to list the available snapshots in the store.
//     // It should return then in descending order, with the highest index first.
//     List() ([]*SnapshotMeta, error)

//     // Open takes a snapshot ID and provides a ReadCloser. Once close is
//     // called it is assumed the snapshot is no longer needed.
//     Open(id string) (*SnapshotMeta, io.ReadCloser, error)
// }

// type LogStore interface {
//     // Returns the first index written. 0 for no entries.
//     FirstIndex() (uint64, error)

//     // Returns the last index written. 0 for no entries.
//     LastIndex() (uint64, error)

//     // Gets a log entry at a given index
//     GetLog(index uint64, log *raft.Log) error

//     // Stores a log entry
//     StoreLog(log *raft.Log) error

//     // Stores multiple log entries
//     StoreLogs(logs []*raft.Log) error

//     // Deletes a range of log entries. The range is inclusive.
//     DeleteRange(min, max uint64) error
// }
