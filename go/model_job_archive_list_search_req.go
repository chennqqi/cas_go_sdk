/*
 * Sample Access Code Flow OAuth2 Project
 *
 * This is an example of using OAuth2 Access Code Flow in a specification to describe security to your API.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type JobArchiveListSearchReq struct {
	// archive-retrieval
	Type string `json:"Type,omitempty"`
	CallBackUrl string `json:"CallBackUrl,omitempty"`
	// 默认JSON
	Format string `json:"Format,omitempty"`
	Description string `json:"Description,omitempty"`
	InventoryRetrievalParameters JobInventoryRetrievalParameters `json:"InventoryRetrievalParameters,omitempty"`
}
