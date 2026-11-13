package fhires

import "encoding/xml"

// Identifies two or more records (resource instances) that refer to the same real-world "occurrence".
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Linkage
type Linkage struct {
	DomainResource
	// Indicates whether the asserted set of linkages are considered to be "in effect".
	//
	// Comment: If false, any asserted linkages should not be considered current/relevant/applicable.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// Identifies the user or organization responsible for asserting the linkages as well as the user or organization who establishes the context in which the nature of each linkage is evaluated.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Identifies which record considered as the reference to the same real-world occurrence as well as how the items should be evaluated within the collection of linked items.
	Item []LinkageItem `bson:"item" json:"item" xml:"item"`
}

type LinkageItem struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Distinguishes which item is "source of truth" (if any) and which items are no longer considered to be current representations.
	Type LinkageType `bson:"type" json:"type" xml:"type"`
	// The resource instance being linked as part of the group.
	Resource Reference `bson:"resource" json:"resource" xml:"resource"`
}

func (r Linkage) ResourceType() string {
	return "Linkage"
}

// MarshalJSON implements [json.Marshaler]
func (r Linkage) MarshalJSON() ([]byte, error) {
	type Wrapper Linkage
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Linkage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Linkage
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Linkage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Linkage
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Linkage(w)
	return err
}
