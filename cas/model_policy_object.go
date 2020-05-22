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
// PolicyObject struct for PolicyObject
type PolicyObject struct {
	Version string `json:"version,omitempty"`
	Statement []PolicyState `json:"statement,omitempty"`
	Action []string `json:"action,omitempty"`
	Resource []string `json:"resource,omitempty"`
}