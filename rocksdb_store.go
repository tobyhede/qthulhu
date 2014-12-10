package qthulhu

// "github/tobyhede/gorocks"
import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"

	"github.com/hashicorp/raft"

	"./../gorocks"
)

type KV struct {
	Key   uint64
	Value []byte
}

type RocksDBStore struct {
	db    *gorocks.DB
	env   *gorocks.Env
	cache *gorocks.Cache
	opts  *gorocks.Options
	ropts *gorocks.ReadOptions
	wopts *gorocks.WriteOptions
	topts *gorocks.TableOptions
	path  string
}

func NewRocksDBStore(path string) *RocksDBStore {
	s := &RocksDBStore{}

	s.path = path
	s.env = gorocks.NewDefaultEnv()
	s.cache = gorocks.NewLRUCache(1 << 20)

	s.opts = gorocks.NewOptions()
	s.opts.SetEnv(s.env)
	s.opts.SetCompression(gorocks.SnappyCompression)
	s.opts.SetCreateIfMissing(true)

	s.ropts = gorocks.NewReadOptions()
	s.ropts.SetVerifyChecksums(true)
	s.ropts.SetFillCache(true)

	s.wopts = gorocks.NewWriteOptions()
	s.wopts.SetSync(false)

	s.topts = gorocks.NewTableOptions()
	s.topts.SetCache(s.cache)

	db, err := gorocks.Open(path, s.opts)
	if err != nil {
		fmt.Println("Open failed: %v", err)
		panic("Open database failed")
	}
	s.db = db
	return s
}

func (s *RocksDBStore) Put(k uint64, v []byte) error {
	err := s.db.Put(s.wopts, uint64ToBytes(k), v)
	return err
}

func (s *RocksDBStore) Set(k, v []byte) error {
	err := s.db.Put(s.wopts, k, v)
	return err
}

func (s *RocksDBStore) PutBatch(logs []*raft.Log) error {
	wb := gorocks.NewWriteBatch()
	defer wb.Close()
	for _, l := range logs {
		k := uint64ToBytes(l.Index)
		wb.Put(k, l.Data)
	}
	err := s.db.Write(s.wopts, wb)

	return err
}

func (s *RocksDBStore) Get(k uint64) ([]byte, error) {
	v, err := s.db.Get(s.ropts, uint64ToBytes(k))
	// inspect(string(v))
	return v, err
}

func (s *RocksDBStore) Iterator() *gorocks.Iterator {
	s.ropts.SetFillCache(false)
	return s.db.NewIterator(s.ropts)
}

func (s *RocksDBStore) FirstKey() (uint64, error) {
	it := s.Iterator()

	defer it.Close()
	it.SeekToFirst()
	return bytesToUint64(it.Key()), nil
}

func (s *RocksDBStore) LastKey() (uint64, error) {
	it := s.Iterator()

	defer it.Close()
	it.SeekToLast()
	return bytesToUint64(it.Key()), nil
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

func (s *RocksDBStore) Delete() {
	err := gorocks.DestroyDatabase(s.path, s.opts)
	if err != nil {
		// t.Errorf("Unable to remove database directory: %s", dirPath)
	}
	// err := os.RemoveAll(s.path)
}

func uint64ToBytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func encode(log *raft.Log) ([]bytes, error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)

	err := enc.Encode(log)
	if err != nil {
		return err
	}
	return b.Bytes(), nil
}
