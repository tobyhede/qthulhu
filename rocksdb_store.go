package qthulhu

// "github/tobyhede/gorocks"
import (
	"log"

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
	s.wopts.SetSync(true)

	s.topts = gorocks.NewTableOptions()
	s.topts.SetCache(s.cache)

	db, err := gorocks.Open(path, s.opts)
	if err != nil {
		log.Fatalf("Error opening database %v\n%v", path, err)
	}
	s.db = db
	return s
}

func (s *RocksDBStore) Put(k, v []byte) error {
	return s.db.Put(s.wopts, k, v)
}

func (s *RocksDBStore) StartBatch() *gorocks.WriteBatch {
	return gorocks.NewWriteBatch()
}

func (s *RocksDBStore) WriteAndCloseBatch(b *gorocks.WriteBatch) error {
	defer b.Close()
	err := s.db.Write(s.wopts, b)
	return err
}

func (s *RocksDBStore) Get(k []byte) ([]byte, error) {
	v, err := s.db.Get(s.ropts, k)
	if len(v) == 0 {
		return nil, NewNotFoundError(k)
	}
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
	err := it.GetError()

	if it.Valid() {
		return bytesToUint64(it.Key()), err
	} else {
		return uint64(0), err
	}
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

func (s *RocksDBStore) Destroy() {
	err := gorocks.DestroyDatabase(s.path, s.opts)
	if err != nil {
		// t.Errorf("Unable to remove database directory: %s", dirPath)
	}
	// err := os.RemoveAll(s.path)
}
