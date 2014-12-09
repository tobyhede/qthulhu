package qthulhu

// "github/tobyhede/gorocks"
import (
	"bytes"
	"encoding/binary"
	"fmt"

	"./../gorocks"
)

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

func (s *RocksDBStore) Put(k uint64, v string) error {
	err := s.db.Put(s.wopts, iToBA(k), []byte(v))
	return err
}

func (s *RocksDBStore) Get(k uint64) (string, error) {
	v, err := s.db.Get(s.ropts, iToBA(k))
	// inspect(string(v))
	return string(v), err
}

func (s *RocksDBStore) Iterator() *gorocks.Iterator {
	s.ropts.SetFillCache(false)
	return s.db.NewIterator(s.ropts)
}

func (s *RocksDBStore) Close() {
	s.env.Close()
	s.cache.Close()
	s.opts.Close()
	s.ropts.Close()
	s.wopts.Close()
	s.topts.Close()
	s.db.Close()
}

func (s *RocksDBStore) Delete() {
	err := gorocks.DestroyDatabase(s.path, s.opts)
	if err != nil {
		// t.Errorf("Unable to remove database directory: %s", dirPath)
	}
	// err := os.RemoveAll(s.path)
}

func iToBA(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func baToI(b []byte) (ret uint64) {
	buf := bytes.NewReader(b)
	binary.Read(buf, binary.BigEndian, &ret)
	return
}
