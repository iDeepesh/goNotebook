package bf

import (
	"hash"
	"hash/fnv"
)

//BloomFilter - the interface for bloom filter
type BloomFilter interface {
	Add(s string) int
	Exists(s string) bool
}

//New - create a new BloomFilter
func New(size int) BloomFilter {
	hf := []hash.Hash64{fnv.New64(), fnv.New64a()}
	bf := realBF{size, 0, make([]byte, size/8), hf}
	return &bf
}

type realBF struct {
	size  int
	count int
	bits  []byte
	hf    []hash.Hash64
}

func (bf *realBF) Add(s string) int {
	for _, h := range bf.getHashes(s) {
		i := h % uint64(bf.size)
		b := 1 << (i % 8)
		bf.bits[i/8] = bf.bits[i/8] | byte(b)
	}
	bf.count++
	// fmt.Println(bf.bits)
	return bf.count
}

func (bf *realBF) getHashes(s string) []uint64 {
	h := make([]uint64, len(bf.hf))
	for i, hf := range bf.hf {
		hf.Write([]byte(s))
		h[i] = hf.Sum64()
		hf.Reset()
	}
	return h
}

func (bf *realBF) Exists(s string) bool {
	found := true
	for _, h := range bf.getHashes(s) {
		i := h % uint64(bf.size)
		b := 1 << (i % 8)
		if (bf.bits[i/8] & byte(b)) <= 0 {
			found = false
			break
		}
	}
	return found
}
