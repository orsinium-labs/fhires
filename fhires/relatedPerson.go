package fhires

import v2 "encoding/json/v2"

// Information about a person that is involved in a patient's health or the care for a patient, but who is not the target of healthcare, nor has a formal responsibility in the care process.
//
// Purpose: Need to track persons related to the patient or the healthcare process.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	// The logical id of the resource, as used in the URL for the resource.
	//
	// Once assigned, this value never changes.
	//
	// Comment: Within the context of the FHIR RESTful interactions, the resource has an id except for cases like the create and conditional update. Otherwise, the use of the resouce id depends on the given use case.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The metadata about the resource.
	//
	// This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content.
	//
	// Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	//
	// Comment: Asserting this rule set restricts the content to be only understood by a limited set of trading partners. This inherently limits the usefulness of the data in the long term. However, the existing health eco-system is highly fractured, and not yet ready to define, collect, and exchange data in a generally computable sense. Wherever possible, implementers and/or specification writers should avoid using this element. Often, when used, the URL is a reference to an implementation guide that defines these special rules as part of its narrative along with other profiles, value sets, etc.
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty" xml:"implicitRules,omitempty"`
	// The base language in which the resource is written.
	//
	// Comment: Language is provided to support indexing and accessibility (typically, services such as text to speech use the language tag). The html language tag in the narrative applies  to the narrative. The language tag on the resource may be used to specify the language of other presentations generated from the data in the resource. Not all the content has to be in the base language. The Resource.language should not be assumed to apply to the narrative automatically. If a language is specified, it should it also be specified on the div element in the html (see rules in HTML5 for information about the relationship between xml:lang and the html lang attribute).
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
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
	// Identifier for a person within a particular scope.
	//
	// Comment: RelatedPerson identifiers might not be unique across instances within a system, as a single human individual may be represented as many different RelatedPerson resources with different roles, periods, or relationships.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Whether this related person record is in active use.
	//
	// Comment: This element is labeled as a modifier because it may be used to mark that the resource was created in error.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// The patient this person is related to.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The nature of the relationship between the related person and the patient.
	//
	// Comment: The directionality of the relationship is from the RelatedPerson to the Patient. For example, if the Patient is a child, and the RelatedPerson is the mother, the relationship would be PRN (parent) or MTH (mother).
	Relationship []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// A name associated with the person.
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A contact detail for the person, e.g. a telephone number or an email address.
	//
	// Comment: Person may have multiple ways to be contacted with different uses or applicable periods.  May need to have options for contacting the person urgently, and also to help with identification.
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty" xml:"telecom,omitempty"`
	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes.
	Gender *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	// The date on which the related person was born.
	BirthDate *string `bson:"birthDate,omitempty" json:"birthDate,omitempty" xml:"birthDate,omitempty"`
	// Address where the related person can be contacted or visited.
	Address []Address `bson:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Image of the person.
	Photo []Attachment `bson:"photo,omitempty" json:"photo,omitempty" xml:"photo,omitempty"`
	// The period of time during which this relationship is or was active.
	//
	// If there are no dates defined, then the interval is unknown.
	//
	// Comment: If an individual has a relationship with a patient over multiple, non-adjacent periods, there should be a distinct RelatedPerson instance for each period.  For example, if a person is a roommate for a period of time, moves out, and is later a roommate with the same person again, you would have two RelatedPerson instances.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// A language which may be used to communicate with the related person about the patient's health.
	//
	// Comment: If no language is specified, this *implies* that the default local language is spoken.  If you need to convey proficiency for multiple modes, then you need multiple RelatedPerson.Communication associations.   If the RelatedPerson does not speak the default local language, then the Interpreter Required Standard can be used to explicitly declare that an interpreter is required.
	Communication []RelatedPersonCommunication `bson:"communication,omitempty" json:"communication,omitempty" xml:"communication,omitempty"`
}

type RelatedPersonCommunication struct {
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
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case; e.g.
	//
	// "en" for English, or "en-US" for American English versus "en-AU" for Australian English.
	//
	// Comment: The structure aa-BB with this exact casing is one the most widely used notations for locale. However not all systems actually code this but instead have it as free text. Hence CodeableConcept instead of code as the data type.
	Language CodeableConcept `bson:"language" json:"language" xml:"language"`
	// Indicates whether or not the related person prefers this language (over other languages he or she masters up a certain level).
	//
	// Comment: This language is specifically identified for communicating healthcare information.
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty" xml:"preferred,omitempty"`
}

func (r RelatedPerson) ResourceType() string {
	return "RelatedPerson"
}

// MarshalJSON implements json.Marshaler
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		RelatedPerson
		ResourceType string `json:"resourceType"`
	}{
		RelatedPerson: r,
		ResourceType:  r.ResourceType(),
	})
}
