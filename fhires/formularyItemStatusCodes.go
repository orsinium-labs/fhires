package fhires

// FormularyItem Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/formularyitem-status
type FormularyItemStatusCodes string

const (
	// The service or product referred to by this FormularyItem is in active use within the drug database or inventory system.
	FormularyItemStatusCodesActive FormularyItemStatusCodes = "active"
	// The service or product referred to by this FormularyItem was entered in error within the drug database or inventory system.
	FormularyItemStatusCodesEnteredInError FormularyItemStatusCodes = "entered-in-error"
	// The service or product referred to by this FormularyItem is not in active use within the drug database or inventory system.
	FormularyItemStatusCodesInactive FormularyItemStatusCodes = "inactive"
)

func (code FormularyItemStatusCodes) String() string {
	return code.Code()
}

func (code FormularyItemStatusCodes) Code() string {
	return string(code)
}

func (code FormularyItemStatusCodes) Display() string {
	switch code {
	case FormularyItemStatusCodesActive:
		return "Active"
	case FormularyItemStatusCodesEnteredInError:
		return "Entered in Error"
	case FormularyItemStatusCodesInactive:
		return "Inactive"
	}
	return "<unknown>"
}

func (code FormularyItemStatusCodes) Definition() string {
	switch code {
	case FormularyItemStatusCodesActive:
		return "The service or product referred to by this FormularyItem is in active use within the drug database or inventory system."
	case FormularyItemStatusCodesEnteredInError:
		return "The service or product referred to by this FormularyItem was entered in error within the drug database or inventory system."
	case FormularyItemStatusCodesInactive:
		return "The service or product referred to by this FormularyItem is not in active use within the drug database or inventory system."
	}
	return "<unknown>"
}
