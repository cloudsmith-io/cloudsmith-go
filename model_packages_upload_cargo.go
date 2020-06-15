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

// PackagesUploadCargo struct for PackagesUploadCargo
type PackagesUploadCargo struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish bool `json:"republish,omitempty"`
}
