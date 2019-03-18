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

package openapi

import (
	"time"
)

type ListParts struct {
	ArchiveDescription string           `json:"ArchiveDescription,omitempty"`
	CreationDate       time.Time        `json:"CreationDate,omitempty"`
	Marker             string           `json:"Marker,omitempty"`
	MultipartUploadId  string           `json:"MultipartUploadId,omitempty"`
	PartSizeInBytes    int64            `json:"PartSizeInBytes,omitempty"`
	VaultQCS           string           `json:"VaultQCS,omitempty"`
	Parts              []ListPartsParts `json:"Parts,omitempty"`
}
