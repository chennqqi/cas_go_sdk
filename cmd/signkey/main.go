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
	"time"
)

const (
	LAYOUT = "2006-01-02 15:04:05"
)

func main() {
	var secret string
	var after, at string
	flag.StringVar(&secret, "secret", "", "set secret key")
	flag.StringVar(&after, "after", "", "set key expire time from now")
	flag.StringVar(&at, "at", "", "set key expire at time")
	flag.Parse()

	if after == "" {
		after = time.Now().Format(LAYOUT)
	}

	start, err := time.Parse(LAYOUT, after)
	if err != nil {
		fmt.Println("parse start", err)
		return
	}
	fmt.Println("start:", start.Unix())
	end, err := time.Parse(LAYOUT, at)
	if err != nil {
		fmt.Println("parse end", err)
		return
	}
	fmt.Println("end:", end.Unix())
	timeRange := fmt.Sprintf("%d;%d", start.Unix(), end.Unix())
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(timeRange))
	signKey := hex.EncodeToString(mac.Sum(nil))
	fmt.Println("sign_key", signKey)
}
