package fhires

// How the system supports versioning for a resource.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/versioning-policy
type ResourceVersionPolicy string

const (
	// VersionId meta-property is not supported (server) or used (client).
	ResourceVersionPolicyNoVersion ResourceVersionPolicy = "no-version"
	// VersionId meta-property is supported (server) or used (client).
	ResourceVersionPolicyVersioned ResourceVersionPolicy = "versioned"
	// Supports version-aware updates (server) or will be specified (If-match header) for updates (client).
	ResourceVersionPolicyVersionedUpdate ResourceVersionPolicy = "versioned-update"
)

func (code ResourceVersionPolicy) String() string {
	return code.Code()
}

func (code ResourceVersionPolicy) Code() string {
	return string(code)
}

func (code ResourceVersionPolicy) Display() string {
	switch code {
	case ResourceVersionPolicyNoVersion:
		return "No VersionId Support"
	case ResourceVersionPolicyVersioned:
		return "Versioned"
	case ResourceVersionPolicyVersionedUpdate:
		return "VersionId tracked fully"
	}
	return "<unknown>"
}

func (code ResourceVersionPolicy) Definition() string {
	switch code {
	case ResourceVersionPolicyNoVersion:
		return "VersionId meta-property is not supported (server) or used (client)."
	case ResourceVersionPolicyVersioned:
		return "VersionId meta-property is supported (server) or used (client)."
	case ResourceVersionPolicyVersionedUpdate:
		return "Supports version-aware updates (server) or will be specified (If-match header) for updates (client)."
	}
	return "<unknown>"
}
