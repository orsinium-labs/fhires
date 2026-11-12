package fhires

import "encoding/json/v2"

// This resource provides the insurance enrollment details to the insurer regarding a specified coverage.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
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
	// The Response business identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the request as not currently valid.
	Status *FinancialResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date when this resource was created.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// The Insurer who is target  of the request.
	Insurer *Reference `bson:"insurer,omitempty" json:"insurer,omitempty" xml:"insurer,omitempty"`
	// The practitioner who is responsible for the services rendered to the patient.
	Provider *Reference `bson:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// Patient Resource.
	Candidate *Reference `bson:"candidate,omitempty" json:"candidate,omitempty" xml:"candidate,omitempty"`
	// Reference to the program or plan identification, underwriter or payor.
	Coverage *Reference `bson:"coverage,omitempty" json:"coverage,omitempty" xml:"coverage,omitempty"`
}

func (r EnrollmentRequest) ResourceType() string {
	return "EnrollmentRequest"
}

// MarshalJSON implements json.Marshaler
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	type Wrapper EnrollmentRequest
	return json.Marshal(struct {
		Wrapper
		ResourceType string `json:"resourceType"`
	}{
		ResourceType: r.ResourceType(),
		Wrapper:      Wrapper(r),
	})
}
