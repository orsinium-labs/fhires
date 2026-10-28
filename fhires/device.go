package fhires

import v2 "encoding/json/v2"

// This resource describes the properties (regulated, has real time clock, etc.), adminstrative (manufacturer name, model number, serial number, firmware, etc.), and type (knee replacement, blood pressure cuff, MRI, etc.) of a physical unit (these values do not change much within a given module, for example the serail number, manufacturer name, and model number).
//
// An actual unit may consist of several modules in a distinct hierarchy and these are represented by multiple Device resources and bound through the 'parent' element.
//
// Purpose: Allows institutions to track their devices.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Device
type Device struct {
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
	// Unique instance identifiers assigned to a device by manufacturers other organizations or owners.
	//
	// Comment: Certain attributes, like serial number and UDI Carrier (the HRF or AIDC barcode string) are not device instance identifiers as they are not consistently able to uniquely identify the instance of a device, thus are not appropriate to be used to value Device.identifier. The barcode string from a barcode present on a device label or package may identify the instance, include names given to the device in local usage, or may identify the type of device. If the identifier identifies the type of device, Device.type element should be used. The identifier is typically valued if the serialNumber or lotNumber is not valued and represents a different type of identifier.  However, it is permissible to still include those identifiers in DeviceDefinition.identifier with the appropriate identifier.type.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The name used to display by default when the device is referenced.
	//
	// Based on intent of use by the resource creator, this may reflect one of the names in Device.name, or may be another simple name.
	DisplayName *string `bson:"displayName,omitempty" json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The reference to the definition for the device.
	Definition *CodeableReference `bson:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Unique device identifier (UDI) assigned to device label or package.
	//
	//  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
	//
	// Comment: UDI may identify an unique instance of a device, or it may only identify the type of the device.  See [UDI mappings](device-mappings.html#udi) for a complete mapping of UDI parts to Device.
	UdiCarrier []DeviceUdiCarrier `bson:"udiCarrier,omitempty" json:"udiCarrier,omitempty" xml:"udiCarrier,omitempty"`
	// The Device record status. This is not the status of the device like availability.
	Status *FHIRDeviceStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The availability of the device.
	AvailabilityStatus *CodeableConcept `bson:"availabilityStatus,omitempty" json:"availabilityStatus,omitempty" xml:"availabilityStatus,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled.
	//
	// Comment: Necessary to support mandatory requirements for traceability from donor/source to recipient and vice versa, while also satisfying donor anonymity requirements. The element is defined consistently across BiologicallyDerivedProduct, NutritionProduct, and Device.  The identifier references an event that links to a single biological entity such as a blood donor, or to multiple biological entities (e.g. when the product is an embryo or a pooled platelet product).  A single biologicalSourceEvent identifier may appear on multiple products of many types derived from a single donation event or source extraction.  As an example, a single donation event may provide 2 kidneys and a liver for organ transplantation, 2 corneas for eye surgery, heart valves and arterial tissue for cardiovascular surgery, multiple skin grafts, tendons, multiple shaped bone grafts and a large number of bone putty/paste products; and each of them may be assigned to the same biological source event identifier.
	BiologicalSourceEvent *Identifier `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty" xml:"biologicalSourceEvent,omitempty"`
	// A name of the manufacturer or entity legally responsible for the device.
	Manufacturer *string `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// The date and time when the device was manufactured.
	ManufactureDate *string `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty" xml:"manufactureDate,omitempty"`
	// The date and time beyond which this device is no longer valid or should not be used (if applicable).
	ExpirationDate *string `bson:"expirationDate,omitempty" json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// Lot number assigned by the manufacturer.
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty" xml:"lotNumber,omitempty"`
	// The serial number assigned by the organization when the device was manufactured.
	//
	// Comment: Alphanumeric Maximum 20. While a serial number is a type of identifier, in the context of devices for which a UDI is required, thus a serial number may be part of the production identifier, it is more informative to have the serial number as a dedicated attribute together with the other UDI production identifier.  When the device is not subject to including the UDI (e.g., its presence and exchange is not mandated by local regulatory requirements or specific use case at hand), thus the concept of production identifiers is not relevant as a set, it remains helpful for consistency to still use Device.serialNumber rather than using Device.identifier with an appropriate type.  Systems that do not realize an identifier is a serial number may place it in Device.identifier or if the identifier.system is known for the serial number, it may also be populated in the Device.identifier as a globally unique id_.
	SerialNumber *string `bson:"serialNumber,omitempty" json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// This represents the manufacturer's name of the device as provided by the device, from a UDI label, or by a person describing the Device.
	//
	//  This typically would be used when a person provides the name(s) or when the device represents one of the names available from DeviceDefinition.
	Name []DeviceName `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The manufacturer's model number for the device.
	ModelNumber *string `bson:"modelNumber,omitempty" json:"modelNumber,omitempty" xml:"modelNumber,omitempty"`
	// The part number or catalog number of the device.
	//
	// Comment: Alphanumeric Maximum 20.
	PartNumber *string `bson:"partNumber,omitempty" json:"partNumber,omitempty" xml:"partNumber,omitempty"`
	// Devices may be associated with one or more categories.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// The kind or type of device.
	//
	// A device instance may have more than one type - in which case those are the types that apply to the specific instance of the device.
	//
	// Comment: Multiple device types are use when a device is categorized in more than one context – for example, hybrid devices in which the device is both of type gateway and sensor.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The actual design of the device or software version running on the device.
	Version []DeviceVersion `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Identifies the standards, specifications, or formal guidances for the capabilities supported by the device.
	//
	// The device may be certified as conformant to these specifications e.g., communication, performance, process, measurement, or specialization standards.
	ConformsTo []DeviceConformsTo `bson:"conformsTo,omitempty" json:"conformsTo,omitempty" xml:"conformsTo,omitempty"`
	// Static or essentially fixed characteristics or features of the device (e.g., time or timing attributes, resolution, accuracy, intended use or instructions for use, and physical attributes) that are not otherwise captured in more specific attributes.
	//
	// Comment: Dynamic or current properties, such as settings, of an individual device are described using a Device instance-specific [DeviceMetric] and recorded using [Observation].  Static characteristics of a device could also be documented in an associated [DeviceDefinition] instance. The Device instance's properties, and their values, could be, but need not be, the same as those described in an associated DeviceDefinition.
	Property []DeviceProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// The designated condition for performing a task with the device.
	Mode *CodeableConcept `bson:"mode,omitempty" json:"mode,omitempty" xml:"mode,omitempty"`
	// The series of occurrences that repeats during the operation of the device.
	Cycle *Count `bson:"cycle,omitempty" json:"cycle,omitempty" xml:"cycle,omitempty"`
	// A measurement of time during the device's operation (e.g., days, hours, mins, etc.).
	Duration *Duration `bson:"duration,omitempty" json:"duration,omitempty" xml:"duration,omitempty"`
	// An organization that is responsible for the provision and ongoing maintenance of the device.
	Owner *Reference `bson:"owner,omitempty" json:"owner,omitempty" xml:"owner,omitempty"`
	// Contact details for an organization or a particular human that is responsible for the device.
	//
	// Comment: used for troubleshooting etc.
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// The place where the device can be found.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// A network address on which the device may be contacted directly.
	//
	// Comment: If the device is running a FHIR server, the network address should  be the Base URL from which a conformance statement may be retrieved.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// Technical endpoints providing access to services provided by the device defined at this resource.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The linked device acting as a communication controller, data collector, translator, or concentrator for the current device (e.g., mobile phone application that relays a blood pressure device's data).
	//
	// Comment: The Device.gateway may be used when the Device being referenced has a gateway and/or to capture one or more gateways associated with the device. If the Gateway is included in the Observation resource, the Observation.gatewayDevice should be used instead.
	Gateway []CodeableReference `bson:"gateway,omitempty" json:"gateway,omitempty" xml:"gateway,omitempty"`
	// Descriptive information, usage information or implantation information that is not captured in an existing element.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Provides additional safety characteristics about a medical device.
	//
	//  For example devices containing latex.
	Safety []CodeableConcept `bson:"safety,omitempty" json:"safety,omitempty" xml:"safety,omitempty"`
	// The higher level or encompassing device that this device is a logical part of.
	//
	// Comment: For example a vital signs monitor (parent) where three separate modules can be plugged into such as interchangeable blood pressure, oximeter, temperature modules.  These modules are represented as devices with the .parent valued to the vital signs monitor when plugged in.
	Parent *Reference `bson:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
}

type DeviceUdiCarrier struct {
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
	// The device identifier (DI) is a mandatory, fixed portion of a UDI that identifies the labeler and the specific version or model of a device.
	DeviceIdentifier string `bson:"deviceIdentifier" json:"deviceIdentifier" xml:"deviceIdentifier"`
	// Organization that is charged with issuing UDIs for devices.
	//
	// For example, the US FDA issuers include:
	//
	// 1) GS1: http://hl7.org/fhir/NamingSystem/gs1-di,
	//
	// 2) HIBCC: http://hl7.org/fhir/NamingSystem/hibcc-diI,
	//
	// 3) ICCBBA for blood containers: http://hl7.org/fhir/NamingSystem/iccbba-blood-di,
	//
	// 4) ICCBA for other devices: http://hl7.org/fhir/NamingSystem/iccbba-other-di # Informationsstelle für Arzneispezialitäten (IFA GmbH) (EU only): http://hl7.org/fhir/NamingSystem/ifa-gmbh-di.
	Issuer string `bson:"issuer" json:"issuer" xml:"issuer"`
	// The identity of the authoritative source for UDI generation within a jurisdiction.
	//
	// All UDIs are globally unique within a single namespace with the appropriate repository uri as the system. For example, UDIs of devices managed in the U.S. by the FDA, the value is http://hl7.org/fhir/NamingSystem/us-fda-udi or in the European Union by the European Commission http://hl7.org/fhir/NamingSystem/eu-ec-udi.
	Jurisdiction *string `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// The full UDI carrier of the Automatic Identification and Data Capture (AIDC) technology representation of the barcode string as printed on the packaging of the device - e.g., a barcode or RFID.
	//
	//   Because of limitations on character sets in XML and the need to round-trip JSON data through XML, AIDC Formats *SHALL* be base64 encoded.
	//
	// Comment: The AIDC form of UDIs should be scanned or otherwise used for the identification of the device whenever possible to minimize errors in records resulting from manual transcriptions. If separate barcodes for DI and PI are present, concatenate the string with DI first and in order of human readable expression on label.
	CarrierAIDC *string `bson:"carrierAIDC,omitempty" json:"carrierAIDC,omitempty" xml:"carrierAIDC,omitempty"`
	// The full UDI carrier as the human readable form (HRF) representation of the barcode string as printed on the packaging of the device.
	//
	// Comment: If separate barcodes for DI and PI are present, concatenate the string with DI first and in order of human readable expression on label.
	CarrierHRF *string `bson:"carrierHRF,omitempty" json:"carrierHRF,omitempty" xml:"carrierHRF,omitempty"`
	// A coded entry to indicate how the data was entered.
	EntryType *UDIEntryType `bson:"entryType,omitempty" json:"entryType,omitempty" xml:"entryType,omitempty"`
}

type DeviceName struct {
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
	// The actual name that identifies the device.
	Value string `bson:"value" json:"value" xml:"value"`
	// Indicates the kind of name. RegisteredName | UserFriendlyName | PatientReportedName.
	Type DeviceNameType `bson:"type" json:"type" xml:"type"`
	// Indicates the default or preferred name to be displayed.
	Display *bool `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
}

type DeviceVersion struct {
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
	// The type of the device version, e.g. manufacturer, approved, internal.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The hardware or software module of the device to which the version applies.
	//
	// Comment: Note that the module of the device would not need to be tracked as a separate device, e.g., using different UDI, thus would typically involve one or more software modules. For example, a device may involve two software modules each on a different version.
	Component *Identifier `bson:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
	// The date the version was installed on the device.
	InstallDate *string `bson:"installDate,omitempty" json:"installDate,omitempty" xml:"installDate,omitempty"`
	// The version text.
	Value string `bson:"value" json:"value" xml:"value"`
}

type DeviceConformsTo struct {
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
	// Describes the type of the standard, specification, or formal guidance.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Code that identifies the specific standard, specification, protocol, formal guidance, regulation, legislation, or certification scheme to which the device adheres.
	Specification CodeableConcept `bson:"specification" json:"specification" xml:"specification"`
	// Identifies the specific form or variant of the standard, specification, or formal guidance.
	//
	// This may be a 'version number', release, document edition, publication year, or other label.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

type DeviceProperty struct {
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
	// Code that specifies the property, such as resolution, color, size, being represented.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The value of the property specified by the associated property.type code.
	//
	// Comment: The text element in CodeableConcept.text is used for properties which would usually be coded (such as the size of an implant, or a security classification) but for which there is not currently an appropriate concept in the vocabulary (e.g. a custom implant size, or a security classification which depends on configuration). Otherwise the string choice type is used for descriptive properties, or instructions.
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueString          string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueAttachment      Attachment      `bson:"valueAttachment" json:"valueAttachment" xml:"valueAttachment"`
}

func (r Device) ResourceType() string {
	return "Device"
}

// MarshalJSON implements json.Marshaler
func (r Device) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Device
		ResourceType string `json:"resourceType"`
	}{
		Device:       r,
		ResourceType: r.ResourceType(),
	})
}
