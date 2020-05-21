/*
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * API version: 1.5.0
 * Contact: chennqqi@qq.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cas
// JobArchiveExportInfoCondition struct for JobArchiveExportInfoCondition
type JobArchiveExportInfoCondition struct {
	IfModifiedSince string `json:"If-Modified-Since,omitempty"`
	IfUmodifiedSince string `json:"If-Umodified-Since,omitempty"`
	IfMatch string `json:"If-Match,omitempty"`
	IfNoneMatch string `json:"If-None-Match,omitempty"`
}
