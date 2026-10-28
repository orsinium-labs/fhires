package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Identifies the style of unique identifier used to identify a namespace.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/namingsystem-identifier-type
type NamingSystemIdentifierType int

const (
	NamingSystemIdentifierTypeOid NamingSystemIdentifierType = iota
	// A universally unique identifier of the form a5afddf4-e880-459b-876e-e4591b0acc11.
	NamingSystemIdentifierTypeUuid
	// A uniform resource identifier (ideally a URL - uniform resource locator); e.g. http://unitsofmeasure.org.
	NamingSystemIdentifierTypeUri
	// An IRI string that can be prepended to the code to obtain a concept IRI for RDF applications. This should be a valid, absolute IRI as defined in RFC 3987. See rdf.html#iri-stem for details on how this value may be used.
	NamingSystemIdentifierTypeIriStem
	// A short string published by HL7 for use in the V2 family of standsrds to idenfify a code system in the V12 coded data types CWE, CNE, and CF. The code values are also published by HL7 at http://www.hl7.org/Special/committees/vocab/table_0396/index.cfm
	NamingSystemIdentifierTypeV2csmnemonic
	// Some other type of unique identifier; e.g. HL7-assigned reserved string such as LN for LOINC.
	NamingSystemIdentifierTypeOther
)

func (code NamingSystemIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *NamingSystemIdentifierType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "oid":
		*code = NamingSystemIdentifierTypeOid
	case "uuid":
		*code = NamingSystemIdentifierTypeUuid
	case "uri":
		*code = NamingSystemIdentifierTypeUri
	case "iri-stem":
		*code = NamingSystemIdentifierTypeIriStem
	case "json.smnemonic":
		*code = NamingSystemIdentifierTypeV2csmnemonic
	case "other":
		*code = NamingSystemIdentifierTypeOther
	default:
		return fmt.Errorf("unknown NamingSystemIdentifierType code `%s`", s)
	}
	return nil
}

func (code NamingSystemIdentifierType) String() string {
	return code.Code()
}

func (code NamingSystemIdentifierType) Code() string {
	switch code {
	case NamingSystemIdentifierTypeOid:
		return "oid"
	case NamingSystemIdentifierTypeUuid:
		return "uuid"
	case NamingSystemIdentifierTypeUri:
		return "uri"
	case NamingSystemIdentifierTypeIriStem:
		return "iri-stem"
	case NamingSystemIdentifierTypeV2csmnemonic:
		return "json.smnemonic"
	case NamingSystemIdentifierTypeOther:
		return "other"
	}
	return "<unknown>"
}

func (code NamingSystemIdentifierType) Display() string {
	switch code {
	case NamingSystemIdentifierTypeOid:
		return "OID"
	case NamingSystemIdentifierTypeUuid:
		return "UUID"
	case NamingSystemIdentifierTypeUri:
		return "URI"
	case NamingSystemIdentifierTypeIriStem:
		return "IRI stem"
	case NamingSystemIdentifierTypeV2csmnemonic:
		return "V2CSMNemonic"
	case NamingSystemIdentifierTypeOther:
		return "Other"
	}
	return "<unknown>"
}

func (code NamingSystemIdentifierType) Definition() string {
	switch code {
	case NamingSystemIdentifierTypeOid:
		return "An ISO object identifier; e.g. 1.2.3.4.5."
	case NamingSystemIdentifierTypeUuid:
		return "A universally unique identifier of the form a5afddf4-e880-459b-876e-e4591b0acc11."
	case NamingSystemIdentifierTypeUri:
		return "A uniform resource identifier (ideally a URL - uniform resource locator); e.g. http://unitsofmeasure.org."
	case NamingSystemIdentifierTypeIriStem:
		return "An IRI string that can be prepended to the code to obtain a concept IRI for RDF applications. This should be a valid, absolute IRI as defined in RFC 3987. See rdf.html#iri-stem for details on how this value may be used."
	case NamingSystemIdentifierTypeV2csmnemonic:
		return "A short string published by HL7 for use in the V2 family of standsrds to idenfify a code system in the V12 coded data types CWE, CNE, and CF. The code values are also published by HL7 at http://www.hl7.org/Special/committees/vocab/table_0396/index.cfm"
	case NamingSystemIdentifierTypeOther:
		return "Some other type of unique identifier; e.g. HL7-assigned reserved string such as LN for LOINC."
	}
	return "<unknown>"
}
