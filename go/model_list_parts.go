/*
 * Sample Access Code Flow OAuth2 Project
 *
 * This is an example of using OAuth2 Access Code Flow in a specification to describe security to your API.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type ListParts struct {
	ArchiveDescription string `json:"ArchiveDescription,omitempty"`
	CreationDate time.Time `json:"CreationDate,omitempty"`
	Marker string `json:"Marker,omitempty"`
	MultipartUploadId string `json:"MultipartUploadId,omitempty"`
	PartSizeInBytes int64 `json:"PartSizeInBytes,omitempty"`
	VaultQCS string `json:"VaultQCS,omitempty"`
	Parts []ListPartsParts `json:"Parts,omitempty"`
}
