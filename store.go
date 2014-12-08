package qthulhu

// "github/tobyhede/gorocks"

type QStore struct {
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
//     GetLog(index uint64, log *Log) error

//     // Stores a log entry
//     StoreLog(log *Log) error

//     // Stores multiple log entries
//     StoreLogs(logs []*Log) error

//     // Deletes a range of log entries. The range is inclusive.
//     DeleteRange(min, max uint64) error
// }
func (s *QStore) FirstIndex() (uint64, error) {
	return 0, nil
}
