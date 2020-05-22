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
// JobOutput struct for JobOutput
type JobOutput struct {
	VaultQCS string `json:"VaultQCS,omitempty"`
	InventoryDate time.Time `json:"InventoryDate,omitempty"`
	ArchiveList []JobOutputArchive `json:"ArchiveList,omitempty"`
}