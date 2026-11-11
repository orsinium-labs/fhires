package fhires

// The purpose of the Claim: predetermination, preauthorization, claim.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/claim-use
type Use string

const (
	// The treatment is complete and this represents a Claim for the services.
	UseClaim Use = "claim"
	// The treatment is proposed and this represents a Pre-authorization for the services.
	UsePreauthorization Use = "preauthorization"
	// The treatment is proposed and this represents a Pre-determination for the services.
	UsePredetermination Use = "predetermination"
)

func (code Use) String() string {
	return code.Code()
}

func (code Use) Code() string {
	return string(code)
}

func (code Use) Display() string {
	switch code {
	case UseClaim:
		return "Claim"
	case UsePreauthorization:
		return "Preauthorization"
	case UsePredetermination:
		return "Predetermination"
	}
	return "<unknown>"
}

func (code Use) Definition() string {
	switch code {
	case UseClaim:
		return "The treatment is complete and this represents a Claim for the services."
	case UsePreauthorization:
		return "The treatment is proposed and this represents a Pre-authorization for the services."
	case UsePredetermination:
		return "The treatment is proposed and this represents a Pre-determination for the services."
	}
	return "<unknown>"
}
