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
// JobArchiveImportReq struct for JobArchiveImportReq
type JobArchiveImportReq struct {
	// push-to-cos
	Type string `json:"Type,omitempty"`
	Description string `json:"Description,omitempty"`
	ArchiveId string `json:"ArchiveId,omitempty"`
	CallBackUrl string `json:"CallBackUrl,omitempty"`
	RetrievalByteRange string `json:"RetrievalByteRange,omitempty"`
	Bucket string `json:"Bucket,omitempty"`
	Object string `json:"Object,omitempty"`
	Tier string `json:"Tier,omitempty"`
}
