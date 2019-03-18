// Copyright 2019 chennqqi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package treehash

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"hash"

	"gitlab.com/NebulousLabs/merkletree"
)

type TreeHash struct {
	buf  *bytes.Buffer
	tree *merkletree.Tree
	hash hash.Hash

	segmentSize int
	cache       []byte
}

func New(segmentSize int, h hash.Hash) hash.Hash {
	return &TreeHash{
		buf:         bytes.NewBuffer(nil),
		tree:        merkletree.New(h),
		hash:        h,
		segmentSize: segmentSize,
		cache:       make([]byte, segmentSize),
	}
}

func (h *TreeHash) Write(p []byte) (n int, err error) {
	buf := h.buf
	tree := h.tree
	n, err = buf.Write(p)

	for {
		if buf.Len() >= h.segmentSize {
			buf.Read(h.cache)
			tree.Push(h.cache)
		} else {
			break
		}
	}
	return n, err
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (h *TreeHash) Sum(b []byte) []byte {
	buf := h.buf
	tree := h.tree
	if buf.Len() > 0 {
		tree.Push(buf.Bytes())
	}
	return tree.Root()
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (h *TreeHash) Reset() {
}

func (h *TreeHash) Size() int {
	return h.hash.Size()
}

func (h *TreeHash) BlockSize() int {
	return h.segmentSize
}

func ComputeHashFromList(list []string) string {
	h := merkletree.New(sha256.New())
	for i := 0; i < len(list); i++ {
		h.Push([]byte(list[i]))
	}
	return hex.EncodeToString(h.Root())
}
