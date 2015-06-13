package qthulhu

import (
	"log"
	"path/filepath"

	"github.com/tobyhede/gorocks"
)

// Store wraps calls to the underlying rockdb instance
type RocksDBStore struct {
	conf      *Config
	db        *gorocks.DB
	env       *gorocks.Env
	cache     *gorocks.Cache
	opts      *gorocks.Options
	ropts     *gorocks.ReadOptions
	wopts     *gorocks.WriteOptions
	topts     *gorocks.TableOptions
	partition string
	key       uint64
}

// NewStore returns a rocksdb store
func NewRocksDBStore(partition string, conf *Config) (Store, error) {
	s := &RocksDBStore{partition: partition, conf: conf}
	s.init()

	if err := s.open(); err != nil {
		return nil, err
	}

	if err := s.ResetKey(); err != nil {
		return nil, err
	}

	// log.Printf("Last Key: %v\n", s.key)

	return s, nil
}

func (s *RocksDBStore) Append(v []byte) error {
	return s.Set(s.NextKey(), v)
}

func (s *RocksDBStore) Close() error {
	s.env.Close()
	s.cache.Close()
	s.opts.Close()
	s.ropts.Close()
	s.wopts.Close()
	s.topts.Close()
	s.db.Close()
	return nil
}

func (s *RocksDBStore) Destroy() error {
	return gorocks.DestroyDatabase(s.path(), s.opts)
}
func (s *RocksDBStore) Get(k uint64) ([]byte, error) {
	v, err := s.db.Get(s.ropts, uint64ToBytes(k))
	if len(v) == 0 {
		return nil, NewNotFoundError(k)
	}
	return v, err
}

func (s *RocksDBStore) LastKey() uint64 {
	return s.key
}

func (s *RocksDBStore) NextKey() uint64 {
	s.key = s.key + 1
	return s.key
}

func (s *RocksDBStore) Set(k uint64, v []byte) error {
	return s.db.Put(s.wopts, uint64ToBytes(k), v)
}

func (s *RocksDBStore) init() {

	s.opts = gorocks.NewOptions()

	s.env = gorocks.NewDefaultEnv()
	s.opts.SetEnv(s.env)
	s.opts.SetCompression(gorocks.SnappyCompression)
	s.opts.SetCreateIfMissing(true)

	s.ropts = gorocks.NewReadOptions()
	s.ropts.SetVerifyChecksums(true)
	s.ropts.SetFillCache(true)

	s.wopts = gorocks.NewWriteOptions()
	s.wopts.SetSync(false)

	s.cache = gorocks.NewLRUCache(100 * 1048576) // 100MB uncompressed cache
	// options.block_cache_compressed = rocksdb::NewLRUCache(100 * 1048576);  // 100MB compressed cache

	s.topts = gorocks.NewTableOptions()
	s.topts.SetCache(s.cache)
}

func (s *RocksDBStore) ResetKey() error {
	it := s.Iterator()
	defer it.Close()

	it.SeekToLast()

	if it.Valid() {
		s.key = bytesToUint64(it.Key())
	}

	if err := it.GetError(); err != nil {
		s.key = uint64(0)
		return err
	}

	return nil
}

func (s *RocksDBStore) Iterator() *gorocks.Iterator {
	ropts := gorocks.NewReadOptions()
	ropts.SetVerifyChecksums(true)
	ropts.SetFillCache(false)
	defer ropts.Close()
	return s.db.NewIterator(ropts)
}

func (s *RocksDBStore) open() error {
	log.Printf("Open DB: %v\n", s.path())
	db, err := gorocks.Open(s.path(), s.opts)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *RocksDBStore) path() string {
	return filepath.Join(s.conf.DataDir, s.partition)
}

//
// func (s *RocksDBStore) StartBatch() *gorocks.WriteBatch {
// 	return gorocks.NewWriteBatch()
// }
//
// func (s *RocksDBStore) WriteAndCloseBatch(b *gorocks.WriteBatch) error {
// 	defer b.Close()
// 	err := s.db.Write(s.wopts, b)
// 	return err
// }
//
// func (s *RocksDBStore) FirstKey() (uint64, error) {
// 	it := s.Iterator()
//
// 	defer it.Close()
// 	it.SeekToFirst()
// 	return bytesToUint64(it.Key()), nil
// }
//

func (s *RocksDBStore) StartBatch() *gorocks.WriteBatch {
	return gorocks.NewWriteBatch()
}

func (s *RocksDBStore) WriteAndCloseBatch(b *gorocks.WriteBatch) error {
	defer b.Close()
	err := s.db.Write(s.wopts, b)
	return err
}

//
// func (s *RocksDBStore) AppendBatch(msgs []interface{}) error {
// 	b := s.StartBatch()
//
// 	for _, msg := range msgs {
// 		// 	d, _ := encode(l)
// 		b.Put(s.NextKey(), d)
// 	}
//
// 	err := s.WriteAndCloseBatch(b)
// 	return err
// }
