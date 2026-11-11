package fhires

// The mode of a RESTful capability statement.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/restful-capability-mode
type RestfulCapabilityMode string

const (
	// The application acts as a client for this resource.
	RestfulCapabilityModeClient RestfulCapabilityMode = "client"
	// The application acts as a server for this resource.
	RestfulCapabilityModeServer RestfulCapabilityMode = "server"
)

func (code RestfulCapabilityMode) String() string {
	return code.Code()
}

func (code RestfulCapabilityMode) Code() string {
	return string(code)
}

func (code RestfulCapabilityMode) Display() string {
	switch code {
	case RestfulCapabilityModeClient:
		return "Client"
	case RestfulCapabilityModeServer:
		return "Server"
	}
	return "<unknown>"
}

func (code RestfulCapabilityMode) Definition() string {
	switch code {
	case RestfulCapabilityModeClient:
		return "The application acts as a client for this resource."
	case RestfulCapabilityModeServer:
		return "The application acts as a server for this resource."
	}
	return "<unknown>"
}
