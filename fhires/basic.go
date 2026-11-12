package fhires

import "encoding/xml"

// Basic is used for handling concepts not yet defined in FHIR, narrative-only resources that don't map to an existing resource, and custom resources not appropriate for inclusion in the FHIR specification.
//
// Purpose: Need some way to safely (without breaking interoperability) allow implementers to exchange content not supported by the initial set of declared resources.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
	BaseResource
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	//
	// Comment: Contained resources do not have a narrative. Resources that are not contained SHOULD have a narrative. In some cases, a resource may only have text with little or no additional discrete data (as long as all minOccurs=1 elements are satisfied).  This may be necessary for data from legacy systems where information is captured as a "text blob" or where text is additionally entered raw or narrated and encoded information is added later.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope.
	//
	// This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	//
	// Comment: This should never be done when the content can be identified properly, as once identification is lost, it is extremely difficult (and context dependent) to restore it again. Contained resources may have profiles and tags in their meta elements, but SHALL NOT have security labels.
	// May be used to represent additional information that is not part of the basic definition of the resource.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Identifier assigned to the resource for business purposes, outside the context of FHIR.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Identifies the 'type' of resource - equivalent to the resource name for other resources.
	//
	// Comment: Because resource references will only be able to indicate 'Basic', the type of reference will need to be specified in a Profile identified as part of the resource.  Refer to the resource notes section for information on appropriate terminologies for this code.
	//
	// This element is labeled as a modifier because it defines the meaning of the resource and cannot be ignored.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Identifies the patient, practitioner, device or any other resource that is the "focus" of this resource.
	//
	// Comment: Optional as not all potential resources will have subjects.  Resources associated with multiple subjects can handle this via extension.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Identifies when the resource was first created.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// Indicates who was responsible for creating the resource instance.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
}

func (r Basic) ResourceType() string {
	return "Basic"
}

// MarshalJSON implements [json.Marshaler]
func (r Basic) MarshalJSON() ([]byte, error) {
	type Wrapper Basic
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Basic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Basic
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Basic(w)
	return err
}
