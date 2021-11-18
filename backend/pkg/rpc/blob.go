package rpc

import (
	bytes "bytes"

	pool "github.com/libp2p/go-buffer-pool"
)

type Blob []byte

func (p *Blob) Release() {
	pool.GlobalPool.Put(p.Bytes())
	*p = nil
}

func (p Blob) Bytes() []byte {
	return p
}

func (p Blob) Clone() Blob {
	r := pool.GlobalPool.Get(p.Size())
	copy(r, p)
	return r
}

func (p Blob) Marshal() ([]byte, error) {
	if len(p) == 0 {
		return nil, nil
	}
	return []byte(p), nil
}

func (p Blob) MarshalTo(data []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	return copy(data, p), nil
}

func (p *Blob) Unmarshal(data []byte) error {
	if len(data) == 0 {
		p = nil
		return nil
	}
	id := Blob(pool.GlobalPool.Get(len(data)))
	copy(id, data)
	*p = id
	return nil
}

func (p *Blob) Size() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

func (p *Blob) Equal(that interface{}) bool {
	that1, ok := that.(Blob)
	if !ok {
		return false
	}
	return bytes.Equal(p.Bytes(), that1.Bytes())
}

func (p *Blob) EqualBlob(that Blob) bool {
	return bytes.Equal(p.Bytes(), that.Bytes())
}

func NewPopulatedBlob(r randySchema) *Blob {
	this := Blob(make([]byte, r.Intn(10240)))

	return &this
}
