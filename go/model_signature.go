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

type Signature struct {
	// 签名算法默认sha1
	QSignAlgorithm string `json:"q-sign-algorithm,omitempty"`
	// 用于标识用户身份， SecretID 的字段，在腾讯云的 API 密钥页面中可以查看。
	QAk string `json:"q-ak,omitempty"`
	// 签名的有效起止时间，其使用10位Unix时间戳来表示，有效效力精确到秒。该字段通过分号区分起止，起时在前、止时在后。
	QSignTime int64 `json:"q-sign-time,omitempty"`
	// 用户可以自定义 SignKey 有效时间，使用 10 位 Unix 时间戳来表示，有效效力精确到秒。该字段通过分号区分起止，起始时间在前、终止时间在后。一般 q-key-time 的时间范围大于等于 q-sign-time。
	QKeyTime int64 `json:"q-key-time,omitempty"`
	// 提供密文中包含需要校验的 Headers 列表，必须是小写字符，且需要按字典序排序，以\";\"分隔
	QHeaderList string `json:"q-header-list,omitempty"`
	// 提供密文中包含需要校验的 Parameters 列表，必须是小写字符，以\";\"分隔
	QUrlParamList string `json:"q-url-param-list,omitempty"`
	// 经过 HMAC-SHA1 算法加密的请求校验信息。
	QSignature string `json:"q-signature,omitempty"`
}
