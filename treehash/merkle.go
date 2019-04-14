package treehash

/*
inspired by https://github.com/cyberdelia/treehash
*/

import (
	"bytes"
	"crypto/sha256"
	"hash"
)

type digest struct {
	h           hash.Hash
	chunks      [][]byte
	buf         *bytes.Buffer
	segmentSize int
}

// New creates a new hash.Hash computing the Tree Hash checksum.
func New(segmentSize int, h hash.Hash) hash.Hash {
	return &digest{
		h:           h,
		chunks:      make([][]byte, 0),
		buf:         bytes.NewBuffer(nil),
		segmentSize: segmentSize,
	}
}

func (d *digest) Size() int { return d.h.Size() }

func (d *digest) BlockSize() int { return d.h.BlockSize() }

func (d *digest) Reset() {
	panic("NOT SUPPORT")
}

func compute(chunks [][]byte) []byte {
	sha := sha256.New()
	previousLevel := chunks
	for {
		if len(previousLevel) == 1 {
			break
		}

		length := len(previousLevel) / 2
		if len(previousLevel)%2 != 0 {
			length++
		}

		currentLevel := make([][]byte, length)
		for i, j := 0, 0; i < len(previousLevel); i, j = i+2, j+1 {
			if len(previousLevel)-i > 1 {
				sha.Reset()
				sha.Write(previousLevel[i])
				sha.Write(previousLevel[i+1])
				currentLevel[j] = sha.Sum(nil)
			} else {
				currentLevel[j] = previousLevel[i]
			}
		}

		previousLevel = currentLevel
	}
	return previousLevel[0]
}

func (d *digest) Sum(in []byte) []byte {
	h := d.h
	buf := d.buf

	if buf.Len() > 0 {
		chunk := make([]byte, d.segmentSize)
		l, _ := buf.Read(chunk)
		h.Reset()
		h.Write(chunk[:l])
		d.chunks = append(d.chunks, h.Sum(nil))
	}
	return compute(d.chunks)
}

func (d *digest) Write(p []byte) (n int, err error) {
	h := d.h
	buf := d.buf

	buf.Write(p)
	for {
		if buf.Len() >= d.segmentSize {
			chunk := make([]byte, d.segmentSize)
			buf.Read(chunk)

			h.Reset()
			h.Write(chunk)
			d.chunks = append(d.chunks, h.Sum(nil))
		} else {
			break
		}
	}
	return len(p), nil
}
