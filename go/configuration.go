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

/*
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * API version: 1.4.3-3.0
 * Contact: chennqqi@qq.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"time"

	"github.com/chennqqi/goutils/utime"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	AppId         string            `json:"appId"`
	AccessKey     string            `json:"accessKey"`
	AccessSecret  string            `json:"accessSecret"`
	SignKeyExpire utime.Duration    `json:"signKeyExpire,omitempty"`
	HTTPClient    *http.Client      `json:"-"`
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "",
		Scheme:        "http",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		SignKeyExpire: utime.Duration(time.Second * 1200),
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
