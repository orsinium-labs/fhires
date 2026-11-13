package fhires

import "encoding/xml"

// Regulatory approval, clearance or licencing related to a regulated product, treatment, facility or activity that is cited in a guidance, regulation, rule or legislative act.
//
// An example is Market Authorization relating to a Medicinal Product.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorization struct {
	DomainResource
	// Business identifier for the authorization, typically assigned by the authorizing body.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The product type, treatment, facility or activity that is being authorized.
	Subject []Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Overall type of this authorization, for example drug marketing approval, orphan drug designation.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// General textual supporting information.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The territory (e.g., country, jurisdiction etc.) in which the authorization has been granted.
	Region []CodeableConcept `bson:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// The status that is authorised e.g.
	//
	// approved. Intermediate states and actions can be tracked with cases and applications.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date at which the current status was assigned.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// The time period in which the regulatory approval, clearance or licencing is in effect.
	//
	// As an example, a Marketing Authorization includes the date of authorization and/or an expiration date.
	ValidityPeriod *Period `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty" xml:"validityPeriod,omitempty"`
	// Condition for which the use of the regulated product applies.
	Indication []CodeableReference `bson:"indication,omitempty" json:"indication,omitempty" xml:"indication,omitempty"`
	// The intended use of the product, e.g. prevention, treatment, diagnosis.
	IntendedUse *CodeableConcept `bson:"intendedUse,omitempty" json:"intendedUse,omitempty" xml:"intendedUse,omitempty"`
	// The legal or regulatory framework against which this authorization is granted, or other reasons for it.
	Basis []CodeableConcept `bson:"basis,omitempty" json:"basis,omitempty" xml:"basis,omitempty"`
	// The organization that has been granted this authorization, by some authoritative body (the 'regulator').
	Holder *Reference `bson:"holder,omitempty" json:"holder,omitempty" xml:"holder,omitempty"`
	// The regulatory authority or authorizing body granting the authorization.
	//
	// For example, European Medicines Agency (EMA), Food and Drug Administration (FDA), Health Canada (HC), etc.
	Regulator *Reference `bson:"regulator,omitempty" json:"regulator,omitempty" xml:"regulator,omitempty"`
	// Additional information or supporting documentation about the authorization.
	AttachedDocument []Reference `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty" xml:"attachedDocument,omitempty"`
	// The case or regulatory procedure for granting or amending a regulated authorization.
	//
	// An authorization is granted in response to submissions/applications by those seeking authorization. A case is the administrative process that deals with the application(s) that relate to this and assesses them. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page).
	Case *RegulatedAuthorizationCase `bson:"case,omitempty" json:"case,omitempty" xml:"case,omitempty"`
}

type RegulatedAuthorizationCase struct {
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
	// Identifier by which this case can be referenced.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The defining type of case.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The status associated with the case.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Relevant date for this case.
	DatePeriod   *Period `bson:"datePeriod,omitempty" json:"datePeriod,omitempty" xml:"datePeriod,omitempty"`
	DateDateTime *string `bson:"dateDateTime,omitempty" json:"dateDateTime,omitempty" xml:"dateDateTime,omitempty"`
	// A regulatory submission from an organization to a regulator, as part of an assessing case.
	//
	// Multiple applications may occur over time, with more or different information to support or modify the submission or the authorization. The applications can be considered as steps within the longer running case or procedure for this authorization process.
	Application []RegulatedAuthorizationCase `bson:"application,omitempty" json:"application,omitempty"`
}

func (r RegulatedAuthorization) ResourceType() string {
	return "RegulatedAuthorization"
}

// MarshalJSON implements [json.Marshaler]
func (r RegulatedAuthorization) MarshalJSON() ([]byte, error) {
	type Wrapper RegulatedAuthorization
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r RegulatedAuthorization) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper RegulatedAuthorization
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *RegulatedAuthorization) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper RegulatedAuthorization
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = RegulatedAuthorization(w)
	return err
}
