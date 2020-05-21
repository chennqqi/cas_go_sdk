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
import (
	"time"
)
// ListParts struct for ListParts
type ListParts struct {
	ArchiveDescription string `json:"ArchiveDescription,omitempty"`
	CreationDate time.Time `json:"CreationDate,omitempty"`
	Marker string `json:"Marker,omitempty"`
	MultipartUploadId string `json:"MultipartUploadId,omitempty"`
	PartSizeInBytes int64 `json:"PartSizeInBytes,omitempty"`
	VaultQCS string `json:"VaultQCS,omitempty"`
	Parts []ListPartsParts `json:"Parts,omitempty"`
}
