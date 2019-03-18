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

type JobArchiveListSearchInfo struct {
	Action                       string                          `json:"Action,omitempty"`
	JobId                        string                          `json:"JobId,omitempty"`
	JobDescription               string                          `json:"JobDescription,omitempty"`
	CallBackUrl                  string                          `json:"CallBackUrl,omitempty"`
	CreationDate                 time.Time                       `json:"CreationDate,omitempty"`
	CompletionDate               time.Time                       `json:"CompletionDate,omitempty"`
	Completed                    bool                            `json:"Completed,omitempty"`
	StatusCode                   string                          `json:"StatusCode,omitempty"`
	StatusMessage                string                          `json:"StatusMessage,omitempty"`
	VaultQCS                     string                          `json:"VaultQCS,omitempty"`
	InventorySizeInBytes         string                          `json:"InventorySizeInBytes,omitempty"`
	InventoryRetrievalParameters JobInventoryRetrievalParameters `json:"InventoryRetrievalParameters,omitempty"`
}
