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

package cas

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoblistJSON(t *testing.T) {
	var job0 = &JobArchiveListSearchInfo{
		Action: "ArchiveRetrieval",
	}
	var job1 = &JobArchiveSearchInfo{
		Action:    "ArchiveRetrieval",
		ArchiveId: "1234",
	}
	var job2 = &JobArchiveImportInfo{
		Action: "PushToCOS",
	}
	var job3 = &JobArchiveExportInfo{
		Action: "PullFromCOS",
	}

	var list JobsList
	list.JobList = append(list.JobList,
		OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo{
			JobArchiveListSearchInfo: job0,
		},
		OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo{
			JobArchiveSearchInfo: job1,
		},
		OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo{
			JobArchiveImportInfo: job2,
		},
		OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo{
			JobArchiveExportInfo: job3,
		})
	txt, err := json.Marshal(list)
	assert.Nil(t, err)
	var newlist JobsList
	err = json.Unmarshal(txt, &newlist)
	assert.Nil(t, err)
	assert.Equal(t, len(newlist.JobList), int(4))
	assert.Equal(t, len(list.JobList), int(4))
	njob0 := newlist.JobList[0].JobArchiveListSearchInfo
	njob1 := newlist.JobList[1].JobArchiveSearchInfo
	njob2 := newlist.JobList[2].JobArchiveImportInfo
	njob3 := newlist.JobList[3].JobArchiveExportInfo
	assert.NotNil(t, njob0)
	assert.NotNil(t, njob1)
	assert.NotNil(t, njob2)
	assert.NotNil(t, njob3)
	assert.Equal(t, *job0, *njob0)
	assert.Equal(t, *job1, *njob1)
	assert.Equal(t, *job2, *njob2)
	assert.Equal(t, *job3, *njob3)
}
