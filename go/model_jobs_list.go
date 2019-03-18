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
	"encoding/json"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

/*
	func UnmarshalJSON(txt []byte) error
	func MarshalJSON() ([]byte,error)
*/
type OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo struct {
	*JobArchiveListSearchInfo `json:"-"`
	*JobArchiveSearchInfo     `json:"-"`
	*JobArchiveExportInfo     `json:"-"`
	*JobArchiveImportInfo     `json:"-"`
}

type JobsList struct {
	JobList []OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo `json:"JobList,omitempty"`
	Marker  string                                                                                      `json:"Marker,omitempty"`
}

func (t OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo) MarshalJSON() ([]byte, error) {
	if t.JobArchiveListSearchInfo != nil {
		return json.Marshal(t.JobArchiveListSearchInfo)
	} else if t.JobArchiveSearchInfo != nil {
		return json.Marshal(t.JobArchiveSearchInfo)
	} else if t.JobArchiveExportInfo != nil {
		return json.Marshal(t.JobArchiveExportInfo)
	} else if t.JobArchiveImportInfo != nil {
		return json.Marshal(t.JobArchiveImportInfo)
	}
	return nil, nil
}

func (t *OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo) UnmarshalJSON(data []byte) error {
	action, _ := jsonparser.GetString(data, "Action")
	_, err := jsonparser.GetString(data, "ArchiveId")

	switch action {
	case "ArchiveRetrieval":
		if err == jsonparser.KeyPathNotFoundError {
			job := new(JobArchiveListSearchInfo)
			err := json.Unmarshal(data, job)
			if err != nil {
				return err
			}
			t.JobArchiveListSearchInfo = job
		} else {
			job := new(JobArchiveSearchInfo)
			err := json.Unmarshal(data, job)
			if err != nil {
				return err
			}
			t.JobArchiveSearchInfo = job
		}

	case "PushToCOS":
		job := new(JobArchiveImportInfo)
		err := json.Unmarshal(data, job)
		if err != nil {
			return err
		}
		t.JobArchiveImportInfo = job

	case "PullFromCOS":
		job := new(JobArchiveExportInfo)
		err := json.Unmarshal(data, job)
		if err != nil {
			return err
		}
		t.JobArchiveExportInfo = job

	default:
		return errors.Errorf("unknow action=[%v]", action)
	}
	return nil
}
