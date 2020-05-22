// Copyright 2019 chennqqi@qq.com
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

package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
)

func init() {
	var secretKey string
	var after, at string
	flag.StringVar(&secretKey, "key", "", "set secret key")
	flag.StringVar(&after, "after", "", "set key expire time from now")
	flag.StringVar(&at, "at", "", "set key expire at time")
	flag.Parse()
	_, _, _ = secretKey, after, at
}

func main() {
	// timeRange = fmt.Sprintf(`%d;%d`, now.Unix(), now.Add(expire).Unix())
	timeRange := "1589817600;1609430400"
	mac := hmac.New(sha1.New, []byte("SgkibEafTCm7D7lAXGoCRSFm7OJzPgiW"))
	mac.Write([]byte(timeRange))
	signKey := hex.EncodeToString(mac.Sum(nil))
	fmt.Println(signKey)
}
