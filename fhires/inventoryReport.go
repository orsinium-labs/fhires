package fhires

import "encoding/xml"

// A report of inventory or stock items.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/InventoryReport
type InventoryReport struct {
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
	// Business identifier for the InventoryReport.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the inventory check or notification - whether this is draft (e.g.
	//
	// the report is still pending some updates) or active.
	Status InventoryReportStatus `bson:"status" json:"status" xml:"status"`
	// Whether the report is about the current inventory count (snapshot) or a differential change in inventory (change).
	CountType InventoryCountType `bson:"countType" json:"countType" xml:"countType"`
	// What type of operation is being performed - addition or subtraction.
	OperationType *CodeableConcept `bson:"operationType,omitempty" json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The reason for this count - regular count, ad-hoc count, new arrivals, etc.
	OperationTypeReason *CodeableConcept `bson:"operationTypeReason,omitempty" json:"operationTypeReason,omitempty" xml:"operationTypeReason,omitempty"`
	// When the report has been submitted.
	ReportedDateTime string `bson:"reportedDateTime" json:"reportedDateTime" xml:"reportedDateTime"`
	// Who submits the report.
	Reporter *Reference `bson:"reporter,omitempty" json:"reporter,omitempty" xml:"reporter,omitempty"`
	// The period the report refers to.
	ReportingPeriod *Period `bson:"reportingPeriod,omitempty" json:"reportingPeriod,omitempty" xml:"reportingPeriod,omitempty"`
	// An inventory listing section (grouped by any of the attributes).
	InventoryListing []InventoryReportInventoryListing `bson:"inventoryListing,omitempty" json:"inventoryListing,omitempty" xml:"inventoryListing,omitempty"`
	// A note associated with the InventoryReport.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type InventoryReportInventoryListing struct {
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
	// Location of the inventory items.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The status of the items.
	ItemStatus *CodeableConcept `bson:"itemStatus,omitempty" json:"itemStatus,omitempty" xml:"itemStatus,omitempty"`
	// The date and time when the items were counted.
	CountingDateTime *string `bson:"countingDateTime,omitempty" json:"countingDateTime,omitempty" xml:"countingDateTime,omitempty"`
	// The item or items in this listing.
	Item []InventoryReportInventoryListingItem `bson:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
}

type InventoryReportInventoryListingItem struct {
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
	// The inventory category or classification of the items being reported.
	//
	// This is meant not for defining the product, but for inventory categories e.g. 'pending recount' or 'damaged'.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// The quantity of the item or items being reported.
	Quantity Quantity `bson:"quantity" json:"quantity" xml:"quantity"`
	// The code or reference to the item type.
	Item CodeableReference `bson:"item" json:"item" xml:"item"`
}

func (r InventoryReport) ResourceType() string {
	return "InventoryReport"
}

// MarshalJSON implements [json.Marshaler]
func (r InventoryReport) MarshalJSON() ([]byte, error) {
	type Wrapper InventoryReport
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *InventoryReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper InventoryReport
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = InventoryReport(w)
	return err
}
