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

// WebhooksPartialUpdate struct for WebhooksPartialUpdate
type WebhooksPartialUpdate struct {
	// None
	Events []string `json:"events,omitempty"`
	// If enabled, the webhook will trigger on events and send payloads to the configured target URL.
	IsActive bool `json:"is_active,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyFormat string `json:"request_body_format,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyTemplateFormat string `json:"request_body_template_format,omitempty"`
	// The value that will be sent for the 'Content Type' header.
	RequestContentType string `json:"request_content_type,omitempty"`
	// The header to send the predefined secret in. This must be unique from existing headers or it won't be sent. You can use this as a form of authentication on the endpoint side.
	SecretHeader string `json:"secret_header,omitempty"`
	// The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side.
	SecretValue string `json:"secret_value,omitempty"`
	// The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn't modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it).
	SignatureKey string `json:"signature_key,omitempty"`
	// The destination URL that webhook payloads will be POST'ed to.
	TargetUrl string `json:"target_url,omitempty"`
	// None
	Templates []WebhooksOwnerRepoTemplates `json:"templates,omitempty"`
	// If enabled, SSL certificates is verified when webhooks are sent. It's recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks.
	VerifySsl bool `json:"verify_ssl,omitempty"`
}
