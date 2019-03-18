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
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

const (
	DEFAULT_CHUNK = 1024 * 1024
)

func ComputeHashFromFile(file string, offset, size, chunk_size int64) (string, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	if chunk_size == 0 {
		chunk_size = DEFAULT_CHUNK
	}

	fi, err := f.Stat()
	if err != nil {
		return "", "", err
	}
	if size == 0 {
		size = fi.Size() - offset
	}
	f.Seek(offset, os.SEEK_SET)

	buf := bufio.NewReader(f)

	h := sha256.New()
	t := New(int(chunk_size), sha256.New())
	mw := io.MultiWriter(h, t)
	io.CopyN(mw, buf, size)

	treeHash := hex.EncodeToString(t.Sum(nil))
	contentHash := hex.EncodeToString(h.Sum(nil))
	return contentHash, treeHash, nil
}
