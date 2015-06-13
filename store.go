package qthulhu

import "github.com/tobyhede/gorocks"

type Store interface {
	Append(v []byte) error
	Close() error
	Destroy() error
	Get(key uint64) ([]byte, error)
	Iterator() *gorocks.Iterator
	LastKey() uint64
	ResetKey() error
	NextKey() uint64
	Set(key uint64, val []byte) error
}

func NewStore(partition string, conf *Config) (Store, error) {
	return NewRocksDBStore(partition, conf)
}
