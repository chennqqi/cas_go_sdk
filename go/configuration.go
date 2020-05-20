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
	BasePath        string            `json:"basePath,omitempty" yaml:"basePath"`
	Host            string            `json:"host,omitempty" yaml:"host"`
	Scheme          string            `json:"scheme,omitempty" yaml:"scheme"`
	DefaultHeader   map[string]string `json:"defaultHeader,omitempty" yaml:"defaultHeader"`
	UserAgent       string            `json:"userAgent,omitempty" yaml:"userAgent"`
	AppId           string            `json:"appId" yaml:"appId"`
	AccessKey       string            `json:"accessKey" yaml:"accessKey"`
	AccessSecret    string            `json:"accessSecret" yaml:"accessSecret"`
	SignKey         string            `json:"signKey,omitempty" yaml:"signKey"`
	SignKeyStart    time.Time         `json:"signKeyStart,omitempty" yaml:"signKeyStart"`
	SignKeyExpireAt time.Time         `json:"signKeyExpire,omitempty" yaml:"signKeyExpire"`
	HTTPClient      *http.Client      `json:"-"`
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "",
		Scheme:        "http",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
