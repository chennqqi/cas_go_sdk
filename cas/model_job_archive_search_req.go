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
// JobArchiveSearchReq struct for JobArchiveSearchReq
type JobArchiveSearchReq struct {
	// archive-retrieval
	Type string `json:"Type,omitempty"`
	ArchiveId string `json:"ArchiveId,omitempty"`
	CallBackUrl string `json:"CallBackUrl,omitempty"`
	Description string `json:"Description,omitempty"`
	RetrievalByteRange string `json:"RetrievalByteRange,omitempty"`
	Tier string `json:"Tier,omitempty"`
}
