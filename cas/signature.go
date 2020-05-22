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
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

func (s *Signature) Sign() string {
	return ""
}

func CalSignKey(start int64, expire time.Duration, accessSecret string) string {
	ts := time.Unix(start, 0)
	timeRange := fmt.Sprintf(`%d;%d`, start, ts.Add(expire).Unix())
	mac := hmac.New(sha1.New, []byte(accessSecret))
	mac.Write([]byte(timeRange))
	return hex.EncodeToString(mac.Sum(nil))
}
