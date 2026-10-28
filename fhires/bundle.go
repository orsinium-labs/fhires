package fhires

import (
	jsontext "encoding/json/jsontext"
	"encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// A container for a collection of resources.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
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
	// A persistent identifier for the bundle that won't change as a bundle is copied from server to server.
	//
	// Comment: Persistent identity generally only matters for batches of type Document, Message, and Collection. It would not normally be populated for search and history results and servers ignore Bundle.identifier when processing batches and transactions. For Documents  the .identifier SHALL be populated such that the .identifier is globally unique.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates the purpose of this bundle - how it is intended to be used.
	//
	// Comment: It's possible to use a bundle for other purposes (e.g. a document can be accepted as a transaction). This is primarily defined so that there can be specific rules for some of the bundle types.
	Type BundleType `bson:"type" json:"type" xml:"type"`
	// The date/time that the bundle was assembled - i.e.
	//
	// when the resources were placed in the bundle.
	//
	// Comment: For many bundles, the timestamp is equal to .meta.lastUpdated, because they are not stored (e.g. search results). When a bundle is placed in a persistent store, .meta.lastUpdated will be usually be changed by the server. When the bundle is a message, a middleware agent altering the message (even if not stored) SHOULD update .meta.lastUpdated. .timestamp is used to track the original time of the Bundle, and SHOULD be populated.
	//
	// Usage:
	//
	// * document : the date the document was created. Note: the composition may predate the document, or be associated with multiple documents. The date of the composition - the authoring time - may be earlier than the document assembly time
	//
	// * message : the date that the content of the message was assembled. This date is not changed by middleware engines unless they add additional data that changes the meaning of the time of the message
	//
	// * history : the date that the history was assembled. This time would be used as the _since time to ask for subsequent updates
	//
	// * searchset : the time that the search set was assembled. Note that different pages MAY have different timestamps but need not. Having different timestamps does not imply that subsequent pages will represent or include changes made since the initial query
	//
	// * transaction | transaction-response | batch | batch-response | collection : no particular assigned meaning
	//
	// The timestamp value should be greater than the lastUpdated and other timestamps in the resources in the bundle, and it should be equal or earlier than the .meta.lastUpdated on the Bundle itself.
	Timestamp *string `bson:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// If a set of search matches, this is the (potentially estimated) total number of entries of type 'match' across all pages in the search.
	//
	//  It does not include search.mode = 'include' or 'outcome' entries and it does not provide a count of the number of entries in the Bundle.
	//
	// Comment: Only used if the bundle is a search result set. The total does not include resources such as OperationOutcome and included resources, only the total number of matching resources.
	Total *int `bson:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// A series of links that provide context to this bundle.
	//
	// Comment: Both Bundle.link and Bundle.entry.link are defined to support providing additional context when Bundles are used (e.g. [HATEOAS](http://en.wikipedia.org/wiki/HATEOAS)).
	//
	// Bundle.entry.link corresponds to links found in the HTTP header if the resource in the entry was [read](http.html#read) directly.
	//
	// This specification defines some specific uses of Bundle.link for [searching](search.html#conformance) and [paging](http.html#paging), but no specific uses for Bundle.entry.link, and no defined function in a transaction - the meaning is implementation specific.  The behavior of navigation link types (next/prev/first/last) are well defined for searchset and history Bundles but are not currently defined for other types. Implementers who choose to use such link relationships for other bundle types will need to negotiate behavior with their interoperability partners.
	Link []BundleLink `bson:"link,omitempty" json:"link,omitempty" xml:"link,omitempty"`
	// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
	Entry []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty" xml:"entry,omitempty"`
	// Digital Signature - base64 encoded. XML-DSig or a JWS.
	//
	// Comment: The signature could be created by the "author" of the bundle or by the originating device.   Requirements around inclusion of a signature, verification of signatures and treatment of signed/non-signed bundles is implementation-environment specific.
	Signature *Signature `bson:"signature,omitempty" json:"signature,omitempty" xml:"signature,omitempty"`
	// Captures issues and warnings that relate to the construction of the Bundle and the content within it.
	//
	// Comment: Usage notes: These issues and warnings must apply to the Bundle as a whole, not to individual entries.Messages relating to the processing of individual entries (e.g. in a batch or transaction) SHALL be reported in the entry.response.outcome for that entry.  If there are errors that arise in the creation of the Bundle, then that should be handled by an OperationOutcome being returned instead of the Bundle.
	Issues jsontext.Value `bson:"issues,omitempty" json:"issues,omitempty" xml:"issues,omitempty"`
}

type BundleLink struct {
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
	// A name which details the functional use for this link - see [http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1](http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1).
	Relation LinkRelationTypes `bson:"relation" json:"relation" xml:"relation"`
	// The reference details for the link.
	URL string `bson:"url" json:"url" xml:"url"`
}

type BundleEntry struct {
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
	// A series of links that provide context to this entry.
	Link []BundleLink `bson:"link,omitempty" json:"link,omitempty"`
	// The Absolute URL for the resource.
	//
	// Except for transactions and batches, each entry in a Bundle must have a fullUrl. The fullUrl SHALL NOT disagree with the id in the resource - i.e. if the fullUrl is not a urn:uuid, the URL shall be version-independent URL consistent with the Resource.id. The fullUrl is a version independent reference to the resource. Even when not required, fullUrl MAY be set to a urn:uuid to allow referencing entries in a transaction. The fullUrl can be an arbitrary URI and is not limited to urn:uuid, urn:oid, http, and https. The fullUrl element SHALL have a value except when:
	//
	// * invoking a create
	//
	// * invoking or responding to an operation where the body is not a single identified resource
	//
	// * invoking or returning the results of a search or history operation.
	//
	// Comment: fullUrl might not be [unique in the context of a resource](bundle.html#bundle-unique). Note that since [FHIR resources do not need to be served through the FHIR API](references.html), the fullURL might be a URN or an absolute URL that does not end with the logical id of the resource (Resource.id). However, but if the fullUrl does look like a RESTful server URL (e.g. meets the [regex](references.html#regex), then the 'id' portion of the fullUrl SHALL end with the Resource.id.
	//
	// Note that the fullUrl is not the same as the canonical URL - it's an absolute url for an endpoint serving the resource (these will happen to have the same value on the canonical server for the resource with the canonical URL).
	FullUURL *string `bson:"fullUrl,omitempty" json:"fullUrl,omitempty" xml:"fullUrl,omitempty"`
	// The Resource for the entry.
	//
	// The purpose/meaning of the resource is determined by the Bundle.type. This is allowed to be a Parameters resource if and only if it is referenced by something else within the Bundle that provides context/meaning.
	Resource jsontext.Value `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// Information about the search process that lead to the creation of this entry.
	Search *BundleEntrySearch `bson:"search,omitempty" json:"search,omitempty" xml:"search,omitempty"`
	// Additional information about how this entry should be processed as part of a transaction or batch.
	//
	//  For history, it shows how the entry was processed to create the version contained in the entry.
	Request *BundleEntryRequest `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
	Response *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty" xml:"response,omitempty"`
}

type BundleEntrySearch struct {
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
	// Why this entry is in the result set - whether it's included as a match or because of an _include requirement, or to convey information or warning information about the search process.
	//
	// Comment: There is only one mode. In some corner cases, a resource may be included because it is both a match and an include. In these circumstances, 'match' takes precedence.
	Mode *SearchEntryMode `bson:"mode,omitempty" json:"mode,omitempty" xml:"mode,omitempty"`
	// When searching, the server's search ranking score for the entry.
	//
	// Comment: Servers are not required to return a ranking score. 1 is most relevant, and 0 is least relevant. Often, search results are sorted by score, but the client may specify a different sort order.
	//
	// See [Patient Match](patient-operation-match.html) for the EMPI search which relates to this element.
	Score *decimal.Decimal `bson:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

type BundleEntryRequest struct {
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
	// In a transaction or batch, this is the HTTP action to be executed for this entry.
	//
	// In a history bundle, this indicates the HTTP action that occurred.
	Method HTTPVerb `bson:"method" json:"method" xml:"method"`
	// The URL for this entry, relative to the root (the address to which the request is posted).
	//
	// Comment: E.g. for a Patient Create, the method would be "POST" and the URL would be "Patient". For a Patient Update, the method would be PUT and the URL would be "Patient/[id]".
	URL string `bson:"url" json:"url" xml:"url"`
	// If the ETag values match, return a 304 Not Modified status.
	//
	// See the API documentation for ["Conditional Read"](http.html#cread).
	IfNoneMatch *string `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty" xml:"ifNoneMatch,omitempty"`
	// Only perform the operation if the last updated date matches.
	//
	// See the API documentation for ["Conditional Read"](http.html#cread).
	IfModifiedSince *string `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty" xml:"ifModifiedSince,omitempty"`
	// Only perform the operation if the Etag value matches.
	//
	// For more information, see the API section ["Managing Resource Contention"](http.html#concurrency).
	IfMatch *string `bson:"ifMatch,omitempty" json:"ifMatch,omitempty" xml:"ifMatch,omitempty"`
	// Instruct the server not to perform the create if a specified resource already exists.
	//
	// For further information, see the API documentation for ["Conditional Create"](http.html#ccreate). This is just the query portion of the URL - what follows the "?" (not including the "?").
	IfNoneExist *string `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty" xml:"ifNoneExist,omitempty"`
}

type BundleEntryResponse struct {
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
	// The status code returned by processing this entry.
	//
	// The status SHALL start with a 3 digit HTTP code (e.g. 404) and may contain the standard HTTP description associated with the status code.
	Status string `bson:"status" json:"status" xml:"status"`
	// The location header created by processing this operation, populated if the operation returns a location.
	Location *string `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The Etag for the resource, if the operation for the entry produced a versioned resource (see [Resource Metadata and Versioning](http.html#versioning) and [Managing Resource Contention](http.html#concurrency)).
	//
	// Comment: Etags match the Resource.meta.versionId. The ETag has to match the version id in the header if a resource is included.
	Etag *string `bson:"etag,omitempty" json:"etag,omitempty" xml:"etag,omitempty"`
	// The date/time that the resource was modified on the server.
	//
	// Comment: This has to match the same time in the meta header (meta.lastUpdated) if a resource is included.
	LastModified *string `bson:"lastModified,omitempty" json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// An OperationOutcome containing hints and warnings produced as part of processing this entry in a batch or transaction.
	//
	// Comment: For a POST/PUT operation, this is the equivalent outcome that would be returned for prefer = operationoutcome - except that the resource is always returned whether or not the outcome is returned.
	//
	// This outcome is not used for error responses in batch/transaction, only for hints and warnings. In a batch operation, the error will be in Bundle.entry.response, and for transaction, there will be a single OperationOutcome instead of a bundle in the case of an error.
	Outcome jsontext.Value `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
}

func (r Bundle) ResourceType() string {
	return "Bundle"
}

// MarshalJSON implements json.Marshaler
func (r Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Bundle
		ResourceType string `json:"resourceType"`
	}{
		Bundle:       r,
		ResourceType: r.ResourceType(),
	})
}
