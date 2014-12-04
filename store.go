package qthulhu

// "github/tobyhede/gorocks"
import (
	"fmt"

	"./../gorocks"
)

type Store struct {
	db    *gorocks.DB
	env   *gorocks.Env
	cache *gorocks.Cache
	opts  *gorocks.Options
	ropts *gorocks.ReadOptions
	wopts *gorocks.WriteOptions
	topts *gorocks.TableOptions
	path  string
}

func NewStore(path string) *Store {
	s := &Store{}

	s.path = path
	s.env = gorocks.NewDefaultEnv()
	s.cache = gorocks.NewLRUCache(1 << 20)

	s.opts = gorocks.NewOptions()
	s.opts.SetEnv(s.env)
	s.opts.SetCompression(gorocks.SnappyCompression)
	s.opts.SetCreateIfMissing(true)

	s.ropts = gorocks.NewReadOptions()
	s.ropts.SetVerifyChecksums(true)
	s.ropts.SetFillCache(false)

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

func (s *Store) Put(k, v string) error {
	err := s.db.Put(s.wopts, []byte(k), []byte(v))
	return err
}

func (s *Store) Get(k string) (string, error) {
	v, err := s.db.Get(s.ropts, []byte(k))
	// inspect(string(v))
	return string(v), err
}

func (s *Store) Close() {
	s.env.Close()
	s.cache.Close()
	s.opts.Close()
	s.ropts.Close()
	s.wopts.Close()
	s.topts.Close()
	s.db.Close()
}

func (s *Store) Delete() {
	err := gorocks.DestroyDatabase(s.path, s.opts)
	if err != nil {
		// t.Errorf("Unable to remove database directory: %s", dirPath)
	}
	// err := os.RemoveAll(s.path)
}
