package fhires

import v2 "encoding/json/v2"

// Detailed definition of a medicinal product, typically for uses other than direct patient care (e.g.
//
// regulatory use, drug catalogs, to support prescribing, adverse events management etc.).
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinition struct {
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
	// Business identifier for this product.
	//
	// Could be an MPID. When in development or being regulated, products are typically referenced by official identifiers, assigned by a manufacturer or regulator, and unique to a product (which, when compared to a product instance being prescribed, is actually a product type). See also MedicinalProductDefinition.code.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Regulatory type, e.g. Investigational or Authorized.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// If this medicine applies to human or veterinary uses.
	Domain *CodeableConcept `bson:"domain,omitempty" json:"domain,omitempty" xml:"domain,omitempty"`
	// A business identifier relating to a specific version of the product, this is commonly used to support revisions to an existing product.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// The status within the lifecycle of this product record.
	//
	// A high-level status, this is not intended to duplicate details carried elsewhere such as legal status, or authorization status.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date at which the given status became applicable.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// General description of this product.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The dose form for a single part product, or combined form of a multiple part product.
	//
	// This is one concept that describes all the components. It does not represent the form with components physically mixed, if that might be necessary, for which see (AdministrableProductDefinition.administrableDoseForm).
	CombinedPharmaceuticalDoseForm *CodeableConcept `bson:"combinedPharmaceuticalDoseForm,omitempty" json:"combinedPharmaceuticalDoseForm,omitempty" xml:"combinedPharmaceuticalDoseForm,omitempty"`
	// The path by which the product is taken into or makes contact with the body.
	//
	// In some regions this is referred to as the licenced or approved route. See also AdministrableProductDefinition resource. MedicinalProductDefinition.route is the same concept as AdministrableProductDefinition.routeOfAdministration.code, and they cannot be used together.
	Route []CodeableConcept `bson:"route,omitempty" json:"route,omitempty" xml:"route,omitempty"`
	// Description of indication(s) for this product, used when structured indications are not required.
	//
	// In cases where structured indications are required, they are captured using the ClinicalUseDefinition resource. An indication is a medical situation for which using the product is appropriate.
	Indication *string `bson:"indication,omitempty" json:"indication,omitempty" xml:"indication,omitempty"`
	// The legal status of supply of the medicinal product as classified by the regulator.
	LegalStatusOfSupply *CodeableConcept `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty" xml:"legalStatusOfSupply,omitempty"`
	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons, such as heightened reporting requirements.
	AdditionalMonitoringIndicator *CodeableConcept `bson:"additionalMonitoringIndicator,omitempty" json:"additionalMonitoringIndicator,omitempty" xml:"additionalMonitoringIndicator,omitempty"`
	// Whether the Medicinal Product is subject to special measures for regulatory reasons, such as a requirement to conduct post-authorization studies.
	SpecialMeasures []CodeableConcept `bson:"specialMeasures,omitempty" json:"specialMeasures,omitempty" xml:"specialMeasures,omitempty"`
	// If authorised for use in children, or infants, neonates etc.
	PediatricUseIndicator *CodeableConcept `bson:"pediatricUseIndicator,omitempty" json:"pediatricUseIndicator,omitempty" xml:"pediatricUseIndicator,omitempty"`
	// Allows the product to be classified by various systems, commonly WHO ATC.
	Classification []CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty" xml:"classification,omitempty"`
	// Marketing status of the medicinal product, in contrast to marketing authorization.
	//
	// This refers to the product being actually 'on the market' as opposed to being allowed to be on the market (which is an authorization).
	MarketingStatus []MarketingStatus `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty" xml:"marketingStatus,omitempty"`
	// Package type for the product. See also the PackagedProductDefinition resource.
	PackagedMedicinalProduct []CodeableConcept `bson:"packagedMedicinalProduct,omitempty" json:"packagedMedicinalProduct,omitempty" xml:"packagedMedicinalProduct,omitempty"`
	// Types of medicinal manufactured items and/or devices that this product consists of, such as tablets, capsule, or syringes.
	//
	// Used as a direct link when the item's packaging is not being recorded (see also PackagedProductDefinition.package.containedItem.item).
	ComprisedOf []Reference `bson:"comprisedOf,omitempty" json:"comprisedOf,omitempty" xml:"comprisedOf,omitempty"`
	// The ingredients of this medicinal product - when not detailed in other resources.
	//
	// This is only needed if the ingredients are not specified by incoming references from the Ingredient resource, or indirectly via incoming AdministrableProductDefinition, PackagedProductDefinition or ManufacturedItemDefinition references. In cases where those levels of detail are not used, the ingredients may be specified directly here as codes.
	Ingredient []CodeableConcept `bson:"ingredient,omitempty" json:"ingredient,omitempty" xml:"ingredient,omitempty"`
	// Any component of the drug product which is not the chemical entity defined as the drug substance, or an excipient in the drug product.
	//
	// This includes process-related impurities and contaminants, product-related impurities including degradation products.
	Impurity []CodeableReference `bson:"impurity,omitempty" json:"impurity,omitempty" xml:"impurity,omitempty"`
	// Additional information or supporting documentation about the medicinal product.
	AttachedDocument []Reference `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty" xml:"attachedDocument,omitempty"`
	// A master file for the medicinal product (e.g.
	//
	// Pharmacovigilance System Master File). Drug master files (DMFs) are documents submitted to regulatory agencies to provide confidential detailed information about facilities, processes or articles used in the manufacturing, processing, packaging and storing of drug products.
	MasterFile []Reference `bson:"masterFile,omitempty" json:"masterFile,omitempty" xml:"masterFile,omitempty"`
	// A product specific contact, person (in a role), or an organization.
	Contact []MedicinalProductDefinitionContact `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// Clinical trials or studies that this product is involved in.
	ClinicalTrial []Reference `bson:"clinicalTrial,omitempty" json:"clinicalTrial,omitempty" xml:"clinicalTrial,omitempty"`
	// A code that this product is known by, usually within some formal terminology, perhaps assigned by a third party (i.e.
	//
	// not the manufacturer or regulator). Products (types of medications) tend to be known by identifiers during development and within regulatory process. However when they are prescribed they tend to be identified by codes. The same product may be have multiple codes, applied to it by multiple organizations.
	Code []Coding `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The product's name, including full name and possibly coded parts.
	Name []MedicinalProductDefinitionName `bson:"name" json:"name" xml:"name"`
	// Reference to another product, e.g.
	//
	// for linking authorised to investigational product, or a virtual product.
	CrossReference []MedicinalProductDefinitionCrossReference `bson:"crossReference,omitempty" json:"crossReference,omitempty" xml:"crossReference,omitempty"`
	// A manufacturing or administrative process or step associated with (or performed on) the medicinal product.
	Operation []MedicinalProductDefinitionOperation `bson:"operation,omitempty" json:"operation,omitempty" xml:"operation,omitempty"`
	// Allows the key product features to be recorded, such as "sugar free", "modified release", "parallel import".
	Characteristic []MedicinalProductDefinitionCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty" xml:"characteristic,omitempty"`
}

type MedicinalProductDefinitionContact struct {
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
	// Allows the contact to be classified, for example QPPV, Pharmacovigilance Enquiry Information.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// A product specific contact, person (in a role), or an organization.
	Contact Reference `bson:"contact" json:"contact" xml:"contact"`
}

type MedicinalProductDefinitionName struct {
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
	// The full product name.
	ProductName string `bson:"productName" json:"productName" xml:"productName"`
	// Type of product name, such as rINN, BAN, Proprietary, Non-Proprietary.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Coding words or phrases of the name.
	Part []MedicinalProductDefinitionNamePart `bson:"part,omitempty" json:"part,omitempty" xml:"part,omitempty"`
	// Country and jurisdiction where the name applies, and associated language.
	Usage []MedicinalProductDefinitionNameUsage `bson:"usage,omitempty" json:"usage,omitempty" xml:"usage,omitempty"`
}

type MedicinalProductDefinitionNamePart struct {
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
	// A fragment of a product name.
	Part string `bson:"part" json:"part" xml:"part"`
	// Identifying type for this part of the name (e.g. strength part).
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
}

type MedicinalProductDefinitionNameUsage struct {
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
	// Country code for where this name applies.
	Country CodeableConcept `bson:"country" json:"country" xml:"country"`
	// Jurisdiction code for where this name applies.
	//
	// A jurisdiction may be a sub- or supra-national entity (e.g. a state or a geographic region).
	Jurisdiction *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Language code for this name.
	Language CodeableConcept `bson:"language" json:"language" xml:"language"`
}

type MedicinalProductDefinitionCrossReference struct {
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
	// Reference to another product, e.g. for linking authorised to investigational product.
	Product CodeableReference `bson:"product" json:"product" xml:"product"`
	// The type of relationship, for instance branded to generic, virtual to actual product, product to development product (investigational), parallel import version.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

type MedicinalProductDefinitionOperation struct {
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
	// The type of manufacturing operation e.g.
	//
	// manufacturing itself, re-packaging. For the authorization of this, a RegulatedAuthorization would point to the same plan or activity referenced here.
	Type *CodeableReference `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Date range of applicability.
	EffectiveDate *Period `bson:"effectiveDate,omitempty" json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	// The organization or establishment responsible for (or associated with) the particular process or step, examples include the manufacturer, importer, agent.
	Organization []Reference `bson:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
	// Specifies whether this particular business or manufacturing process is considered proprietary or confidential.
	ConfidentialityIndicator *CodeableConcept `bson:"confidentialityIndicator,omitempty" json:"confidentialityIndicator,omitempty" xml:"confidentialityIndicator,omitempty"`
}

type MedicinalProductDefinitionCharacteristic struct {
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
	// A code expressing the type of characteristic.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A value for the characteristic.text.
	//
	// Comment: Text, as markdown, is intended to be used only when the item cannot ever reasonably be coded. Otherwise, for cases where a code is not available, but is feasible, use CodeableConcept.text.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueMarkdown        *string          `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty" xml:"valueMarkdown,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueInteger         *int             `bson:"valueInteger,omitempty" json:"valueInteger,omitempty" xml:"valueInteger,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty" xml:"valueDate,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
}

func (r MedicinalProductDefinition) ResourceType() string {
	return "MedicinalProductDefinition"
}

// MarshalJSON implements json.Marshaler
func (r MedicinalProductDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		MedicinalProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		MedicinalProductDefinition: r,
		ResourceType:               r.ResourceType(),
	})
}
