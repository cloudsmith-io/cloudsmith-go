/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.34
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// PackagesValidateUploadRaw struct for PackagesValidateUploadRaw
type PackagesValidateUploadRaw struct {
	// A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here.
	ContentType string `json:"content_type,omitempty"`
	// A textual description of this package.
	Description string `json:"description,omitempty"`
	// The name of this package.
	Name string `json:"name,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish bool `json:"republish,omitempty"`
	// A one-liner synopsis of this package.
	Summary string `json:"summary,omitempty"`
	// The raw version for this package.
	Version string `json:"version,omitempty"`
}
