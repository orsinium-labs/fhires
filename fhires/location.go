package fhires

import (
	"encoding/json/v2"

	"github.com/shopspring/decimal"
)

// Details and position information for a place where services are provided and resources and participants may be stored, found, contained, or accommodated.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
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
	// Unique code or number identifying the location to its users.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status property covers the general availability of the resource, not the current value which may be covered by the operationStatus, or by a schedule/slots if they are configured for the location.
	Status *LocationStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The operational status covers operation values most relevant to beds (but can also apply to rooms/units/chairs/etc.
	//
	// such as an isolation unit/dialysis chair). This typically covers concepts such as contamination, housekeeping, and other activities like maintenance.
	OperationalStatus *Coding `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty" xml:"operationalStatus,omitempty"`
	// Name of the location as used by humans. Does not need to be unique.
	//
	// Comment: If the name of a location changes, consider putting the old name in the alias column so that it can still be located through searches.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A list of alternate names that the location is known as, or was known as, in the past.
	//
	// Comment: There are no dates associated with the alias/historic names, as this is not intended to track when names were used, but to assist in searching so that older names can still result in identifying the location.
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty" xml:"alias,omitempty"`
	// Description of the Location, which helps in finding or referencing the place.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether a resource instance represents a specific location or a class of locations.
	//
	// Comment: This is labeled as a modifier because whether or not the location is a class of locations changes how it can be used and understood.
	Mode *LocationMode `bson:"mode,omitempty" json:"mode,omitempty" xml:"mode,omitempty"`
	// Indicates the type of function performed at the location.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The contact details of communication devices available at the location.
	//
	// This can include addresses, phone numbers, fax numbers, mobile numbers, email addresses and web sites.
	//
	// Comment: The address/telecom use code 'home' are not to be used. Note that these contacts are not the contact details of people who provide the service (that would be through PractitionerRole), these are official contacts for the Healthcare itself for specific purposes. E.g. Mailing Addresses, Billing Addresses, Contact numbers for Booking or Billing Enquiries, general web address, web address for online bookings etc.
	//
	// If this is empty (or the type of interest is empty), refer to the organization's contacts.
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// Physical location.
	//
	// Comment: Additional addresses should be recorded using another instance of the Location resource, or via the Organization.
	Address *Address `bson:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Physical form of the location, e.g. building, room, vehicle, road, virtual.
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty" xml:"form,omitempty"`
	// The absolute geographic location of the Location, expressed using the WGS84 datum (This is the same co-ordinate system used in KML).
	//
	// Comment: To define a boundary shape for this location use the standard extension `[http://hl7.org/fhir/StructureDefinition/location-boundary-geojson](http://hl7.org/fhir/extensions/StructureDefinition-location-boundary-geojson.html)`, and search using the `contains` special search parameter.
	Position *LocationPosition `bson:"position,omitempty" json:"position,omitempty" xml:"position,omitempty"`
	// The organization responsible for the provisioning and upkeep of the location.
	//
	// Comment: This can also be used as the part of the organization hierarchy where this location provides services. These services can be defined through the HealthcareService resource.
	ManagingOrganization *Reference `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty" xml:"managingOrganization,omitempty"`
	// Another Location of which this Location is physically a part of.
	PartOf *Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// Collection of characteristics (attributes).
	//
	// Comment: These could be such things as is wheelchair accessible.
	Characteristic []CodeableConcept `bson:"characteristic,omitempty" json:"characteristic,omitempty" xml:"characteristic,omitempty"`
	// What days/times during a week is this location usually open, and any exceptions where the location is not available.
	//
	// Comment: This type of information is commonly found published in directories and on websites informing customers when the facility is available.
	//
	// Specific services within the location may have their own hours which could be shorter (or longer) than the locations hours.
	//
	// Systems may choose to render availability differently than it is exchanged on the interface. For example, rather than "Mon, Tue, Wed, Thur, Fri from 9am-12am; Mon, Tue, Wed, Thur, Fri from 1pm-5pm" as would be implied by two availableTime repetitions, an application could render this information as "Mon-Fri 9-12am and 1-5pm".The availableStartTime is the opening time, and the availableEndTime is the closing time.
	HoursOfOperation []Availability `bson:"hoursOfOperation,omitempty" json:"hoursOfOperation,omitempty" xml:"hoursOfOperation,omitempty"`
	// Connection details of a virtual service (e.g.
	//
	// shared conference call facility with dedicated number/details).
	//
	// Comment: There are two types of virtual meetings that often exist:* a persistent, virtual meeting room that can only be used for a single purpose at a time, * and a dynamic virtual meeting room that is generated on demand for a specific purpose. Implementers may consider using Appointment.virtualService for virtual meeting rooms that are generated on-demand.
	VirtualService []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty" xml:"virtualService,omitempty"`
	// Technical endpoints providing access to services operated for the location.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
}

type LocationPosition struct {
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
	// Longitude.
	//
	// The value domain and the interpretation are the same as for the text of the longitude element in KML (see notes on Location main page).
	Longitude decimal.Decimal `bson:"longitude" json:"longitude" xml:"longitude"`
	// Latitude.
	//
	// The value domain and the interpretation are the same as for the text of the latitude element in KML (see notes on Location main page).
	Latitude decimal.Decimal `bson:"latitude" json:"latitude" xml:"latitude"`
	// Altitude.
	//
	// The value domain and the interpretation are the same as for the text of the altitude element in KML (see notes on Location main page).
	Altitude *decimal.Decimal `bson:"altitude,omitempty" json:"altitude,omitempty" xml:"altitude,omitempty"`
}

func (r Location) ResourceType() string {
	return "Location"
}

// MarshalJSON implements json.Marshaler
func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Location
		ResourceType string `json:"resourceType"`
	}{
		Location:     r,
		ResourceType: r.ResourceType(),
	})
}
