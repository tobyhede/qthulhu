package qthulhu

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

func uint64ToBytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func encode(o interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)

	if err := enc.Encode(o); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decode(b []byte, out interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	return dec.Decode(out)
}
