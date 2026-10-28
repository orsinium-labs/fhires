package fhires

import v2 "encoding/json/v2"

// The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons.
//
// Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
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
	// Identifiers assigned to this event performer or other systems.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// References the (external) source of pricing information, rules of application for the code this ChargeItem uses.
	DefinitionUri []string `bson:"definitionUri,omitempty" json:"definitionUri,omitempty" xml:"definitionUri,omitempty"`
	// References the source of pricing information, rules of application for the code this ChargeItem uses.
	DefinitionCanonical []string `bson:"definitionCanonical,omitempty" json:"definitionCanonical,omitempty" xml:"definitionCanonical,omitempty"`
	// The current state of the ChargeItem.
	//
	// Comment: Unknown does not represent "other" - one of the defined statuses must apply.  Unknown is used when the authoring system is not sure what the current status is.
	//
	// This element is labeled as a modifier because the status contains the code entered-in-error that marks the charge item as not currently valid.
	Status ChargeItemStatus `bson:"status" json:"status" xml:"status"`
	// ChargeItems can be grouped to larger ChargeItems covering the whole set.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code that identifies the charge, like a billing code.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The individual or set of individuals the action is being or was performed on.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// This ChargeItem has the details of how the associated Encounter should be billed or otherwise be handled by finance systems.
	//
	// Comment: This ChargeItem may be recorded during planning, execution or after the actual encounter takes place.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Date/time(s) or duration when the charged service was applied.
	//
	// Comment: The list of types may be constrained as appropriate for the type of charge item.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// Indicates who or what performed or participated in the charged service.
	Performer []ChargeItemPerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The organization performing the service.
	//
	// Comment: Practitioners and Devices can be associated with multiple organizations. It has to be made clear, on behalf of which Organization the services have been rendered.
	PerformingOrganization *Reference `bson:"performingOrganization,omitempty" json:"performingOrganization,omitempty" xml:"performingOrganization,omitempty"`
	// The organization requesting the service.
	//
	// Comment: The rendered Service might not be associated with a Request. This property indicates which Organization requested the services to be rendered. (In many cases, this may just be the Department associated with the Encounter.location).
	RequestingOrganization *Reference `bson:"requestingOrganization,omitempty" json:"requestingOrganization,omitempty" xml:"requestingOrganization,omitempty"`
	// The financial cost center permits the tracking of charge attribution.
	//
	// Comment: The costCenter could either be given as a reference to an Organization(Role) resource or as the identifier of the cost center determined by Reference.identifier.value and Reference.identifier.system, depending on use case requirements.
	CostCenter *Reference `bson:"costCenter,omitempty" json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	// Quantity of which the charge item has been serviced.
	//
	// Comment: In many cases this may just be a value, if the underlying units are implicit in the definition of the charge item code.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The anatomical location where the related service has been applied.
	//
	// Comment: Only used if not implicit in code found in Condition.code. If the use case requires attributes from the BodySite resource (e.g. to identify and track separately) then use the standard extension [http://hl7.org/fhir/StructureDefinition/bodySite](http://hl7.org/fhir/extensions/StructureDefinition-bodySite.html).  May be a summary code, or a reference to a very precise definition of the location, or both.
	Bodysite []CodeableConcept `bson:"bodysite,omitempty" json:"bodysite,omitempty" xml:"bodysite,omitempty"`
	// The unit price of the chargable item.
	//
	// Comment: This could be communicated in ChargeItemDefinition. But if ChargeItemDefinition is not supported or not available, the unit price can be communicated in this property.
	UnitPriceComponent *MonetaryComponent `bson:"unitPriceComponent,omitempty" json:"unitPriceComponent,omitempty" xml:"unitPriceComponent,omitempty"`
	// The total price for the chargable item, accounting for the quantity.
	//
	// Comment:  Often, the total price may be be calculated and recorded on the Invoice, but if a calculated total price must be exchanged prior to Invoice creation, it can be communicated in this property.
	TotalPriceComponent *MonetaryComponent `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty" xml:"totalPriceComponent,omitempty"`
	// If the list price or the rule-based factor associated with the code is overridden, this attribute can capture a text to indicate the  reason for this action.
	OverrideReason *CodeableConcept `bson:"overrideReason,omitempty" json:"overrideReason,omitempty" xml:"overrideReason,omitempty"`
	// The device, practitioner, etc. who entered the charge item.
	//
	// Comment: The enterer is also the person considered responsible for factor/price overrides if applicable.
	Enterer *Reference `bson:"enterer,omitempty" json:"enterer,omitempty" xml:"enterer,omitempty"`
	// Date the charge item was entered.
	//
	// Comment: The actual date when the service associated with the charge has been rendered is captured in occurrence[x].
	EnteredDate *string `bson:"enteredDate,omitempty" json:"enteredDate,omitempty" xml:"enteredDate,omitempty"`
	// Describes why the event occurred in coded or textual form.
	//
	// Comment: If the application of the charge item requires a reason to be given, it can be captured here. Textual reasons can be captured using reasonCode.text.
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Indicated the rendered service that caused this charge.
	Service []CodeableReference `bson:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Identifies the device, food, drug or other product being charged either by type code or reference to an instance.
	Product []CodeableReference `bson:"product,omitempty" json:"product,omitempty" xml:"product,omitempty"`
	// Account into which this ChargeItems belongs.
	//
	// Comment: Systems posting the ChargeItems might not always be able to determine, which accounts the Items need to be places into. It is up to the postprocessing Financial System to apply internal rules to decide based on the Encounter/EpisodeOfCare/Patient/Coverage context and the type of ChargeItem, which Account is appropriate.
	Account []Reference `bson:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Comments made about the event by the performer, subject or other participants.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Further information supporting this charge.
	SupportingInformation []Reference `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
}

type ChargeItemPerformer struct {
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
	// Describes the type of performance or participation(e.g.
	//
	// primary surgeon, anesthesiologiest, etc.).
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The device, practitioner, etc. who performed or participated in the service.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

func (r ChargeItem) ResourceType() string {
	return "ChargeItem"
}

// MarshalJSON implements json.Marshaler
func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ChargeItem
		ResourceType string `json:"resourceType"`
	}{
		ChargeItem:   r,
		ResourceType: r.ResourceType(),
	})
}
