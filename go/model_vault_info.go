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

type VaultInfo struct {
	CreationDate      string `json:"CreationDate,omitempty"`
	LastInventoryDate string `json:"LastInventoryDate,omitempty"`
	NumberOfArchives  int64  `json:"NumberOfArchives,omitempty"`
	SizeInBytes       int64  `json:"SizeInBytes,omitempty"`
	VaultQCS          string `json:"VaultQCS,omitempty"`
	VaultName         string `json:"VaultName,omitempty"`
}
