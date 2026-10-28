package fhires

import v2 "encoding/json/v2"

// Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively, and are not formally or legally recognized; i.e.
//
// a collection of entities that isn't an Organization.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Group
type Group struct {
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
	// Business identifiers assigned to this participant by one of the applications involved.
	//
	//  These identifiers remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: Note: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates whether the record for the group is available for use or is merely being retained for historical purposes.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// Identifies the broad classification of the kind of resources the group includes.
	//
	// Comment: Group members SHALL be of the appropriate resource type (Patient for person or animal; or Practitioner, PractitionerRole, Device, CareTeam, HealthcareService, Location, Organization, RelatedPerson, or Specimen for the other types.), or a Group of the resources of the appropriate type.
	Type GroupType `bson:"type" json:"type" xml:"type"`
	// Basis for membership in the Group:
	//
	// * 'definitional': The Group.characteristics specified are both necessary and sufficient to determine membership.
	//
	// All entities that meet the criteria are considered to be members of the group, whether referenced by the group or not. If members are present, they are individuals that happen to be known as meeting the Group.characteristics. The list cannot be presumed to be complete.
	//
	// * 'enumerated': The Group.characteristics are necessary but not sufficient to determine membership. Membership is determined by being listed as one of the Group.member.
	Membership GroupMembershipBasis `bson:"membership" json:"membership" xml:"membership"`
	// Provides a specific type of resource the group includes; e.g. "cow", "syringe", etc.
	//
	// Comment: This would generally be omitted for Person resources.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// A label assigned to the group for human identification and communication.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Explanation of what the group represents and how it is intended to be used.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A count of the number of resource instances that are part of the group.
	//
	// Comment: Note that the quantity may be less than the number of members if some of the members are not active.
	Quantity *int `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Entity responsible for defining and maintaining Group characteristics and/or registered members.
	//
	// Comment: This does not strictly align with ownership of a herd or flock, but may suffice to represent that relationship in simple cases. More complex cases will require an extension.
	ManagingEntity *Reference `bson:"managingEntity,omitempty" json:"managingEntity,omitempty" xml:"managingEntity,omitempty"`
	// Identifies traits whose presence r absence is shared by members of the group.
	//
	// Comment: All the identified characteristics must be true for an entity to a member of the group.
	Characteristic []GroupCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty" xml:"characteristic,omitempty"`
	// Identifies the resource instances that are members of the group.
	Member []GroupMember `bson:"member,omitempty" json:"member,omitempty" xml:"member,omitempty"`
}

type GroupCharacteristic struct {
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
	// A code that identifies the kind of trait being asserted.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The value of the trait that holds (or does not hold - see 'exclude') for members of the group.
	//
	// Comment: For Range, it means members of the group have a value that falls somewhere within the specified range.
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference" xml:"valueReference"`
	// If true, indicates the characteristic is one that is NOT held by members of the group.
	//
	// Comment: This is labeled as "Is Modifier" because applications cannot wrongly include excluded members as included or vice versa.
	Exclude bool `bson:"exclude" json:"exclude" xml:"exclude"`
	// The period over which the characteristic is tested; e.g.
	//
	// the patient had an operation during the month of June.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

type GroupMember struct {
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
	// A reference to the entity that is a member of the group.
	//
	// Must be consistent with Group.type. If the entity is another group, then the type must be the same.
	Entity Reference `bson:"entity" json:"entity" xml:"entity"`
	// The period that the member was in the group, if known.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// A flag to indicate that the member is no longer in the group, but previously may have been a member.
	Inactive *bool `bson:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
}

func (r Group) ResourceType() string {
	return "Group"
}

// MarshalJSON implements json.Marshaler
func (r Group) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Group
		ResourceType string `json:"resourceType"`
	}{
		Group:        r,
		ResourceType: r.ResourceType(),
	})
}
