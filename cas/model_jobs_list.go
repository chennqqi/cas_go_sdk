/*
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * API version: 1.5.0
 * Contact: chennqqi@qq.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// JobsList struct for JobsList
type JobsList struct {
	JobList []OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo `json:"JobList,omitempty"`
	Marker string `json:"Marker,omitempty"`
}
