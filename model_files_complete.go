/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.44
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// FilesComplete struct for FilesComplete
type FilesComplete struct {
	// Filename for the package file upload.
	Filename string `json:"filename"`
	// MD5 checksum for a POST-based package file upload.
	Md5Checksum string `json:"md5_checksum,omitempty"`
	// The method to use for package file upload.
	Method string `json:"method,omitempty"`
	// SHA256 checksum for a PUT-based package file upload.
	Sha256Checksum string `json:"sha256_checksum,omitempty"`
}
