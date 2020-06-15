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

// PackagesValidateUploadMaven struct for PackagesValidateUploadMaven
type PackagesValidateUploadMaven struct {
	// The ID of the artifact.
	ArtifactId string `json:"artifact_id,omitempty"`
	// Artifact's group ID.
	GroupId string `json:"group_id,omitempty"`
	// Adds bundled Java documentation to the Maven package
	JavadocFile string `json:"javadoc_file,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// Artifact's Maven packaging type.
	Packaging string `json:"packaging,omitempty"`
	// The POM file is an XML file containing the Maven coordinates.
	PomFile string `json:"pom_file,omitempty"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish bool `json:"republish,omitempty"`
	// Adds bundled Java source code to the Maven package.
	SourcesFile string `json:"sources_file,omitempty"`
	// Adds bundled Java tests to the Maven package.
	TestsFile string `json:"tests_file,omitempty"`
	// The raw version for this package.
	Version string `json:"version,omitempty"`
}
