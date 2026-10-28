package fhires

// Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).
//
//	This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
//
// The ISO21090-codedString may be used to provide a coded representation of the contents of strings in an Address.
//
// Purpose: Need to be able to record postal addresses, along with notes about their use.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Address
type Address struct {
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
	// The purpose of this address.
	//
	// Comment: Applications can assume that an address is current unless it explicitly says that it is temporary or old.
	Use *AddressUse `bson:"use,omitempty" json:"use,omitempty" xml:"use,omitempty"`
	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g.
	//
	// PO Boxes and care-of addresses). Most addresses are both.
	//
	// Comment: The definition of Address states that "address is intended to describe postal addresses, not physical locations". However, many applications track whether an address has a dual purpose of being a location that can be visited as well as being a valid delivery destination, and Postal addresses are often used as proxies for physical locations (also see the [Location](location.html#) resource).
	Type *AddressType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Specifies the entire address as it should be displayed e.g.
	//
	// on a postal label. This may be provided instead of or as well as the specific parts.
	//
	// Comment: Can provide both a text representation and parts. Applications updating an address SHALL ensure that  when both text and parts are present,  no content is included in the text that isn't found in a part.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// This component contains the house number, apartment number, street name, street direction,  P.O.
	//
	// Box number, delivery hints, and similar address information.
	Line []string `bson:"line,omitempty" json:"line,omitempty" xml:"line,omitempty"`
	// The name of the city, town, suburb, village or other community or delivery center.
	City *string `bson:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// The name of the administrative area (county).
	//
	// Comment: District is sometimes known as county, but in some regions 'county' is used in place of city (municipality), so county name should be conveyed in city instead.
	District *string `bson:"district,omitempty" json:"district,omitempty" xml:"district,omitempty"`
	// Sub-unit of a country with limited sovereignty in a federally organized country.
	//
	// A code may be used if codes are in common use (e.g. US 2 letter state codes).
	State *string `bson:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// A postal code designating a region defined by the postal service.
	PostalCode *string `bson:"postalCode,omitempty" json:"postalCode,omitempty" xml:"postalCode,omitempty"`
	// Country - a nation as commonly understood or generally accepted.
	//
	// Comment: ISO 3166 2- or 3- letter codes MAY be used in place of a human readable country name.
	Country *string `bson:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Time period when address was/is in use.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}
