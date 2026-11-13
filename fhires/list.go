package fhires

import "encoding/xml"

// A List is a curated collection of resources, for things such as problem lists, allergy lists, facility list, organization list, etc.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/List
type List struct {
	DomainResource
	// Identifier for the List assigned for business purposes outside the context of FHIR.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates the current state of this list.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status ListStatus `bson:"status" json:"status" xml:"status"`
	// How this list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
	//
	// Comment: This element is labeled as a modifier because a change list must not be misunderstood as a complete list.
	Mode ListMode `bson:"mode" json:"mode" xml:"mode"`
	// A label for the list assigned by the author.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// This code defines the purpose of the list - why it was created.
	//
	// Comment: If there is no code, the purpose of the list is implied where it is used, such as in a document section using Document.section.code.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The common subject(s) (or patient(s)) of the resources that are in the list if there is one (or a set of subjects).
	//
	// Comment: Some purely arbitrary lists do not have a common subject, so this is optional. Lists having multiple subjects may raise challenges for systems that base access control or other security considerations on linkages to 'subject'. E.g. if a List is associated with multiple Patient subjects, then it would be nominally part of multiple patient compartments, which might drive access rights.
	Subject []Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The encounter that is the context in which this list was created.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Date list was last reviewed/revised and determined to be 'current'.
	//
	// Comment: The actual important date is the date of currency of the resources that were summarized, but it is usually assumed that these are current when the preparation occurs.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The entity responsible for deciding what the contents of the list were.
	//
	// Where the list was created by a human, this is the same as the author of the list.
	//
	// Comment: The primary source is the entity that made the decisions what items are in the list. This may be software or user.
	Source *Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// What order applies to the items in the list.
	//
	// Comment: Applications SHOULD render ordered lists in the order provided, but MAY allow users to re-order based on their own preferences as well. If there is no order specified, the order is unknown, though there may still be some order.
	OrderedBy *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty" xml:"orderedBy,omitempty"`
	// Comments that apply to the overall list.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Entries in this list.
	//
	// Comment: If there are no entries in the list, an emptyReason SHOULD be provided.
	Entry []ListEntry `bson:"entry,omitempty" json:"entry,omitempty" xml:"entry,omitempty"`
	// If the list is empty, why the list is empty.
	//
	// Comment: The various reasons for an empty list make a significant interpretation to its interpretation. Note that this code is for use when the entire list has been suppressed, and not for when individual items are omitted - implementers may consider using a text note or a flag on an entry in these cases.
	EmptyReason *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty" xml:"emptyReason,omitempty"`
}

type ListEntry struct {
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
	// The flag allows the system constructing the list to indicate the role and significance of the item in the list.
	//
	// Comment: The flag can only be understood in the context of the List.code. If the flag means that the entry has actually been deleted from the list, the deleted element SHALL be true
	Flag *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty" xml:"flag,omitempty"`
	// True if this item is marked as deleted in the list.
	//
	// Comment: If the flag means that the entry has actually been deleted from the list, the deleted element SHALL be true. A deleted entry should be displayed in narrative as deleted.  This element is labeled as a modifier because it indicates that an item is (to be) no longer in the list.
	Deleted *bool `bson:"deleted,omitempty" json:"deleted,omitempty" xml:"deleted,omitempty"`
	// When this item was added to the list.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// A reference to the actual resource from which data was derived.
	Item Reference `bson:"item" json:"item" xml:"item"`
}

func (r List) ResourceType() string {
	return "List"
}

// MarshalJSON implements [json.Marshaler]
func (r List) MarshalJSON() ([]byte, error) {
	type Wrapper List
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r List) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper List
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *List) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper List
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = List(w)
	return err
}
