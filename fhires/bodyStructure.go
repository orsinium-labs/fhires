package fhires

import "encoding/xml"

// Record details about an anatomical structure.
//
//	This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/BodyStructure
type BodyStructure struct {
	DomainResource
	// Identifier for this instance of the anatomical structure.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Whether this body site is in active use.
	//
	// Comment: This element is labeled as a modifier because it may be used to mark that the resource was created in error.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// The kind of structure being represented by the body structure at `BodyStructure.location`.
	//
	//  This can define both normal and abnormal morphologies.
	//
	// Comment: The minimum cardinality of 0 supports the use case of specifying a location without defining a morphology.
	Morphology *CodeableConcept `bson:"morphology,omitempty" json:"morphology,omitempty" xml:"morphology,omitempty"`
	// The anatomical location(s) or region(s) of the specimen, lesion, or body structure.
	IncludedStructure []BodyStructureIncludedStructure `bson:"includedStructure" json:"includedStructure" xml:"includedStructure"`
	// The anatomical location(s) or region(s) not occupied or represented by the specimen, lesion, or body structure.
	ExcludedStructure []BodyStructureIncludedStructure `bson:"excludedStructure,omitempty" json:"excludedStructure,omitempty"`
	// A summary, characterization or explanation of the body structure.
	//
	// Comment: This description could include any visual markings used to orientate the viewer e.g. external reference points, special sutures, ink markings.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Image or images used to identify a location.
	Image []Attachment `bson:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// The person to which the body site belongs.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
}

type BodyStructureIncludedStructure struct {
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
	// Code that represents the included structure.
	Structure CodeableConcept `bson:"structure" json:"structure" xml:"structure"`
	// Code that represents the included structure laterality.
	Laterality *CodeableConcept `bson:"laterality,omitempty" json:"laterality,omitempty" xml:"laterality,omitempty"`
	// Body locations in relation to a specific body landmark (tatoo, scar, other body structure).
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `bson:"bodyLandmarkOrientation,omitempty" json:"bodyLandmarkOrientation,omitempty" xml:"bodyLandmarkOrientation,omitempty"`
	// XY or XYZ-coordinate orientation for structure.
	SpatialReference []Reference `bson:"spatialReference,omitempty" json:"spatialReference,omitempty" xml:"spatialReference,omitempty"`
	// Code that represents the included structure qualifier.
	Qualifier []CodeableConcept `bson:"qualifier,omitempty" json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
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
	// A description of a landmark on the body used as a reference to locate something else.
	LandmarkDescription []CodeableConcept `bson:"landmarkDescription,omitempty" json:"landmarkDescription,omitempty" xml:"landmarkDescription,omitempty"`
	// An description of the direction away from a landmark something is located based on a radial clock dial.
	ClockFacePosition []CodeableConcept `bson:"clockFacePosition,omitempty" json:"clockFacePosition,omitempty" xml:"clockFacePosition,omitempty"`
	// The distance in centimeters a certain observation is made from a body landmark.
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `bson:"distanceFromLandmark,omitempty" json:"distanceFromLandmark,omitempty" xml:"distanceFromLandmark,omitempty"`
	// The surface area a body location is in relation to a landmark.
	SurfaceOrientation []CodeableConcept `bson:"surfaceOrientation,omitempty" json:"surfaceOrientation,omitempty" xml:"surfaceOrientation,omitempty"`
}

type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
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
	// An instrument, tool, analyzer, etc. used in the measurement.
	Device []CodeableReference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// The measured distance (e.g., in cm) from a body landmark.
	Value []Quantity `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

func (r BodyStructure) ResourceType() string {
	return "BodyStructure"
}

// MarshalJSON implements [json.Marshaler]
func (r BodyStructure) MarshalJSON() ([]byte, error) {
	type Wrapper BodyStructure
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r BodyStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper BodyStructure
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *BodyStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper BodyStructure
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = BodyStructure(w)
	return err
}
