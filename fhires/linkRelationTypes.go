package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Link Relation Types defined at https://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/iana-link-relations
type LinkRelationTypes int

const (
	LinkRelationTypesAbout LinkRelationTypes = iota
	// Asserts that the link target provides an access control description for the link context.
	LinkRelationTypesAcl
	// Refers to a substitute for this context
	LinkRelationTypesAlternate
	// Used to reference alternative content that uses the AMP profile of the HTML format.
	LinkRelationTypesAmphtml
	// Refers to an appendix.
	LinkRelationTypesAppendix
	// Refers to an icon for the context. Synonym for icon.
	LinkRelationTypesAppleTouchIcon
	// Refers to a launch screen for the context.
	LinkRelationTypesAppleTouchStartupImage
	/*
	   Refers to a collection of records, documents, or other
	         materials of historical interest.
	*/
	LinkRelationTypesArchives
	// Refers to the context's author.
	LinkRelationTypesAuthor
	/*
	   Identifies the entity that blocks access to a resource
	         following receipt of a legal demand.
	*/
	LinkRelationTypesBlockedBy
	// Gives a permanent link to use for bookmarking purposes.
	LinkRelationTypesBookmark
	// Designates the preferred version of a resource (the IRI and its contents).
	LinkRelationTypesCanonical
	// Refers to a chapter in a collection of resources.
	LinkRelationTypesChapter
	// Indicates that the link target is preferred over the link context for the purpose of permanent citation.
	LinkRelationTypesCiteAs
	// The target IRI points to a resource which represents the collection resource for the context IRI.
	LinkRelationTypesCollection
	// Refers to a table of contents.
	LinkRelationTypesContents
	/*
	   The document linked to was later converted to the
	         document that contains this link relation.  For example, an RFC can
	         have a link to the Internet-Draft that became the RFC; in that case,
	         the link relation would be "convertedFrom".
	*/
	LinkRelationTypesConvertedFrom
	/*
	   Refers to a copyright statement that applies to the
	       link's context.
	*/
	LinkRelationTypesCopyright
	// The target IRI points to a resource where a submission form can be obtained.
	LinkRelationTypesCreateForm
	/*
	   Refers to a resource containing the most recent
	         item(s) in a collection of resources.
	*/
	LinkRelationTypesCurrent
	/*
	   Refers to a resource providing information about the
	         link's context.
	*/
	LinkRelationTypesDescribedby
	/*
	   The relationship A 'describes' B asserts that
	         resource A provides a description of resource B. There are no
	         constraints on the format or representation of either A or B,
	         neither are there any further constraints on either resource.
	*/
	LinkRelationTypesDescribes
	/*
	   Refers to a list of patent disclosures made with respect to
	         material for which 'disclosure' relation is specified.
	*/
	LinkRelationTypesDisclosure
	/*
	   Used to indicate an origin that will be used to fetch required
	         resources for the link context, and that the user agent ought to resolve
	         as early as possible.
	*/
	LinkRelationTypesDnsPrefetch
	/*
	   Refers to a resource whose available representations
	         are byte-for-byte identical with the corresponding representations of
	         the context IRI.
	*/
	LinkRelationTypesDuplicate
	/*
	   Refers to a resource that can be used to edit the
	         link's context.
	*/
	LinkRelationTypesEdit
	/*
	   The target IRI points to a resource where a submission form for
	         editing associated resource can be obtained.
	*/
	LinkRelationTypesEditForm
	/*
	   Refers to a resource that can be used to edit media
	         associated with the link's context.
	*/
	LinkRelationTypesEditMedia
	/*
	   Identifies a related resource that is potentially
	         large and might require special handling.
	*/
	LinkRelationTypesEnclosure
	// Refers to a resource that is not part of the same site as the current context.
	LinkRelationTypesExternal
	/*
	   An IRI that refers to the furthest preceding resource
	       in a series of resources.
	*/
	LinkRelationTypesFirst
	// Refers to a glossary of terms.
	LinkRelationTypesGlossary
	// Refers to context-sensitive help.
	LinkRelationTypesHelp
	/*
	   Refers to a resource hosted by the server indicated by
	         the link context.
	*/
	LinkRelationTypesHosts
	/*
	   Refers to a hub that enables registration for
	       notification of updates to the context.
	*/
	LinkRelationTypesHub
	// Refers to an icon representing the link's context.
	LinkRelationTypesIcon
	// Refers to an index.
	LinkRelationTypesIndex
	// refers to a resource associated with a time interval that ends before the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalAfter
	// refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalBefore
	// refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource
	LinkRelationTypesIntervalContains
	// refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource, or ends before the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalDisjoint
	// refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalDuring
	// refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalEquals
	// refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalFinishedBy
	// refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalFinishes
	// refers to a resource associated with a time interval that begins before or is coincident with the beginning of the time interval associated with the context resource, and ends after or is coincident with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalIn
	// refers to a resource associated with a time interval whose beginning coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalMeets
	// refers to a resource associated with a time interval whose end coincides with the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalMetBy
	// refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalOverlappedBy
	// refers to a resource associated with a time interval that begins before the end of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalOverlaps
	// refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource
	LinkRelationTypesIntervalStartedBy
	// refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalStarts
	// The target IRI points to a resource that is a member of the collection represented by the context IRI.
	LinkRelationTypesItem
	/*
	   An IRI that refers to the furthest following resource
	         in a series of resources.
	*/
	LinkRelationTypesLast
	/*
	   Points to a resource containing the latest (e.g.,
	         current) version of the context.
	*/
	LinkRelationTypesLatestVersion
	// Refers to a license associated with this context.
	LinkRelationTypesLicense
	/*
	   The link target of a link with the "linkset" relation
	         type provides a set of links, including links in which the link
	         context of the link participates.

	*/
	LinkRelationTypesLinkset
	/*
	   Refers to further information about the link's context,
	         expressed as a LRDD ("Link-based Resource Descriptor Document")
	         resource.  See  for information about
	         processing this relation type in host-meta documents. When used
	         elsewhere, it refers to additional links and other metadata.
	         Multiple instances indicate additional LRDD resources. LRDD
	         resources MUST have an "application/xrd+xml" representation, and
	         MAY have others.
	*/
	LinkRelationTypesLrdd
	// Links to a manifest file for the context.
	LinkRelationTypesManifest
	// Refers to a mask that can be applied to the icon for the context.
	LinkRelationTypesMaskIcon
	// Refers to a feed of personalised media recommendations relevant to the link context.
	LinkRelationTypesMediaFeed
	// The Target IRI points to a Memento, a fixed resource that will not change state anymore.
	LinkRelationTypesMemento
	// Links to the context's Micropub endpoint.
	LinkRelationTypesMicropub
	// Refers to a module that the user agent is to preemptively fetch and store for use in the current context.
	LinkRelationTypesModulepreload
	/*
	   Refers to a resource that can be used to monitor changes in an HTTP resource.

	*/
	LinkRelationTypesMonitor
	/*
	   Refers to a resource that can be used to monitor changes in a specified group of HTTP resources.

	*/
	LinkRelationTypesMonitorGroup
	/*
	   Indicates that the link's context is a part of a series, and
	         that the next in the series is the link target.

	*/
	LinkRelationTypesNext
	// Refers to the immediately following archive resource.
	LinkRelationTypesNextArchive
	// Indicates that the context’s original author or publisher does not endorse the link target.
	LinkRelationTypesNofollow
	// Indicates that any newly created top-level browsing context which results from following the link will not be an auxiliary browsing context.
	LinkRelationTypesNoopener
	// Indicates that no referrer information is to be leaked when following the link.
	LinkRelationTypesNoreferrer
	// Indicates that any newly created top-level browsing context which results from following the link will be an auxiliary browsing context.
	LinkRelationTypesOpener
	// Refers to an OpenID Authentication server on which the context relies for an assertion that the end user controls an Identifier.
	LinkRelationTypesOpenid2_Local_id
	// Refers to a resource which accepts OpenID Authentication protocol messages for the context.
	LinkRelationTypesOpenid2_Provider
	// The Target IRI points to an Original Resource.
	LinkRelationTypesOriginal
	// Refers to a P3P privacy policy for the context.
	LinkRelationTypesP3Pv1
	// Indicates a resource where payment is accepted.
	LinkRelationTypesPayment
	// Gives the address of the pingback resource for the link context.
	LinkRelationTypesPingback
	/*
	   Used to indicate an origin that will be used to fetch required
	         resources for the link context. Initiating an early connection, which
	         includes the DNS lookup, TCP handshake, and optional TLS negotiation,
	         allows the user agent to mask the high latency costs of establishing a
	         connection.
	*/
	LinkRelationTypesPreconnect
	/*
	   Points to a resource containing the predecessor
	         version in the version history.

	*/
	LinkRelationTypesPredecessorVersion
	/*
	   The prefetch link relation type is used to identify a resource
	         that might be required by the next navigation from the link context, and
	         that the user agent ought to fetch, such that the user agent can deliver a
	         faster response once the resource is requested in the future.
	*/
	LinkRelationTypesPrefetch
	/*
	   Refers to a resource that should be loaded early in the
	         processing of the link's context, without blocking rendering.
	*/
	LinkRelationTypesPreload
	/*
	   Used to identify a resource that might be required by the next
	         navigation from the link context, and that the user agent ought to fetch
	         and execute, such that the user agent can deliver a faster response once
	         the resource is requested in the future.
	*/
	LinkRelationTypesPrerender
	/*
	   Indicates that the link's context is a part of a series, and
	         that the previous in the series is the link target.

	*/
	LinkRelationTypesPrev
	// Refers to a resource that provides a preview of the link's context.
	LinkRelationTypesPreview
	/*
	   Refers to the previous resource in an ordered series
	         of resources.  Synonym for "prev".
	*/
	LinkRelationTypesPrevious
	// Refers to the immediately preceding archive resource.
	LinkRelationTypesPrevArchive
	// Refers to a privacy policy associated with the link's context.
	LinkRelationTypesPrivacyPolicy
	/*
	   Identifying that a resource representation conforms
	   to a certain profile, without affecting the non-profile semantics
	   of the resource representation.
	*/
	LinkRelationTypesProfile
	/*
	   Links to a publication manifest. A manifest represents
	         structured information about a publication, such as informative metadata,
	         a list of resources, and a default reading order.
	*/
	LinkRelationTypesPublication
	// Identifies a related resource.
	LinkRelationTypesRelated
	/*
	   Identifies the root of RESTCONF API as configured on this HTTP server.
	         The "restconf" relation defines the root of the API defined in RFC8040.
	         Subsequent revisions of RESTCONF will use alternate relation values to support
	         protocol versioning.
	*/
	LinkRelationTypesRestconf
	/*
	   Identifies a resource that is a reply to the context
	         of the link.

	*/
	LinkRelationTypesReplies
	/*
	   The resource identified by the link target provides an input value to an
	       instance of a rule, where the resource which represents the rule instance is
	       identified by the link context.

	*/
	LinkRelationTypesRuleinput
	/*
	   Refers to a resource that can be used to search through
	         the link's context and related resources.
	*/
	LinkRelationTypesSearch
	// Refers to a section in a collection of resources.
	LinkRelationTypesSection
	/*
	   Conveys an identifier for the link's context.

	*/
	LinkRelationTypesSelf
	/*
	   Indicates a URI that can be used to retrieve a
	         service document.
	*/
	LinkRelationTypesService
	/*
	   Identifies service description for the context that
	         is primarily intended for consumption by machines.
	*/
	LinkRelationTypesServiceDesc
	/*
	   Identifies service documentation for the context that
	         is primarily intended for human consumption.
	*/
	LinkRelationTypesServiceDoc
	/*
	   Identifies general metadata for the context that is
	         primarily intended for consumption by machines.
	*/
	LinkRelationTypesServiceMeta
	/*
	   Refers to a resource that is within a context that is
	   		sponsored (such as advertising or another compensation agreement).
	*/
	LinkRelationTypesSponsored
	/*
	   Refers to the first resource in a collection of
	         resources.
	*/
	LinkRelationTypesStart
	/*
	   Identifies a resource that represents the context's
	         status.
	*/
	LinkRelationTypesStatus
	// Refers to a stylesheet.
	LinkRelationTypesStylesheet
	/*
	   Refers to a resource serving as a subsection in a
	         collection of resources.
	*/
	LinkRelationTypesSubsection
	/*
	   Points to a resource containing the successor version
	         in the version history.

	*/
	LinkRelationTypesSuccessorVersion
	/*
	   Identifies a resource that provides information about
	         the context's retirement policy.

	*/
	LinkRelationTypesSunset
	/*
	   Gives a tag (identified by the given address) that applies to
	         the current document.

	*/
	LinkRelationTypesTag
	// Refers to the terms of service associated with the link's context.
	LinkRelationTypesTermsOfService
	// The Target IRI points to a TimeGate for an Original Resource.
	LinkRelationTypesTimegate
	// The Target IRI points to a TimeMap for an Original Resource.
	LinkRelationTypesTimemap
	// Refers to a resource identifying the abstract semantic type of which the link's context is considered to be an instance.
	LinkRelationTypesType
	/*
	   Refers to a resource that is within a context that is User Generated Content.

	*/
	LinkRelationTypesUgc
	/*
	   Refers to a parent document in a hierarchy of
	         documents.

	*/
	LinkRelationTypesUp
	/*
	   Points to a resource containing the version history
	         for the context.

	*/
	LinkRelationTypesVersionHistory
	/*
	   Identifies a resource that is the source of the
	         information in the link's context.

	*/
	LinkRelationTypesVia
	/*
	   Identifies a target URI that supports the Webmention protocol.
	       This allows clients that mention a resource in some form of publishing process
	       to contact that endpoint and inform it that this resource has been mentioned.
	*/
	LinkRelationTypesWebmention
	// Points to a working copy for this resource.
	LinkRelationTypesWorkingCopy
	/*
	   Points to the versioned resource from which this
	         working copy was obtained.

	*/
	LinkRelationTypesWorkingCopyOf
)

func (code LinkRelationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *LinkRelationTypes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "about":
		*code = LinkRelationTypesAbout
	case "acl":
		*code = LinkRelationTypesAcl
	case "alternate":
		*code = LinkRelationTypesAlternate
	case "amphtml":
		*code = LinkRelationTypesAmphtml
	case "appendix":
		*code = LinkRelationTypesAppendix
	case "apple-touch-icon":
		*code = LinkRelationTypesAppleTouchIcon
	case "apple-touch-startup-image":
		*code = LinkRelationTypesAppleTouchStartupImage
	case "archives":
		*code = LinkRelationTypesArchives
	case "author":
		*code = LinkRelationTypesAuthor
	case "blocked-by":
		*code = LinkRelationTypesBlockedBy
	case "bookmark":
		*code = LinkRelationTypesBookmark
	case "canonical":
		*code = LinkRelationTypesCanonical
	case "chapter":
		*code = LinkRelationTypesChapter
	case "cite-as":
		*code = LinkRelationTypesCiteAs
	case "collection":
		*code = LinkRelationTypesCollection
	case "contents":
		*code = LinkRelationTypesContents
	case "convertedFrom":
		*code = LinkRelationTypesConvertedFrom
	case "copyright":
		*code = LinkRelationTypesCopyright
	case "create-form":
		*code = LinkRelationTypesCreateForm
	case "current":
		*code = LinkRelationTypesCurrent
	case "describedby":
		*code = LinkRelationTypesDescribedby
	case "describes":
		*code = LinkRelationTypesDescribes
	case "disclosure":
		*code = LinkRelationTypesDisclosure
	case "dns-prefetch":
		*code = LinkRelationTypesDnsPrefetch
	case "duplicate":
		*code = LinkRelationTypesDuplicate
	case "edit":
		*code = LinkRelationTypesEdit
	case "edit-form":
		*code = LinkRelationTypesEditForm
	case "edit-media":
		*code = LinkRelationTypesEditMedia
	case "enclosure":
		*code = LinkRelationTypesEnclosure
	case "external":
		*code = LinkRelationTypesExternal
	case "first":
		*code = LinkRelationTypesFirst
	case "glossary":
		*code = LinkRelationTypesGlossary
	case "help":
		*code = LinkRelationTypesHelp
	case "hosts":
		*code = LinkRelationTypesHosts
	case "hub":
		*code = LinkRelationTypesHub
	case "icon":
		*code = LinkRelationTypesIcon
	case "index":
		*code = LinkRelationTypesIndex
	case "intervalAfter":
		*code = LinkRelationTypesIntervalAfter
	case "intervalBefore":
		*code = LinkRelationTypesIntervalBefore
	case "intervalContains":
		*code = LinkRelationTypesIntervalContains
	case "intervalDisjoint":
		*code = LinkRelationTypesIntervalDisjoint
	case "intervalDuring":
		*code = LinkRelationTypesIntervalDuring
	case "intervalEquals":
		*code = LinkRelationTypesIntervalEquals
	case "intervalFinishedBy":
		*code = LinkRelationTypesIntervalFinishedBy
	case "intervalFinishes":
		*code = LinkRelationTypesIntervalFinishes
	case "intervalIn":
		*code = LinkRelationTypesIntervalIn
	case "intervalMeets":
		*code = LinkRelationTypesIntervalMeets
	case "intervalMetBy":
		*code = LinkRelationTypesIntervalMetBy
	case "intervalOverlappedBy":
		*code = LinkRelationTypesIntervalOverlappedBy
	case "intervalOverlaps":
		*code = LinkRelationTypesIntervalOverlaps
	case "intervalStartedBy":
		*code = LinkRelationTypesIntervalStartedBy
	case "intervalStarts":
		*code = LinkRelationTypesIntervalStarts
	case "item":
		*code = LinkRelationTypesItem
	case "last":
		*code = LinkRelationTypesLast
	case "latest-version":
		*code = LinkRelationTypesLatestVersion
	case "license":
		*code = LinkRelationTypesLicense
	case "linkset":
		*code = LinkRelationTypesLinkset
	case "lrdd":
		*code = LinkRelationTypesLrdd
	case "manifest":
		*code = LinkRelationTypesManifest
	case "mask-icon":
		*code = LinkRelationTypesMaskIcon
	case "media-feed":
		*code = LinkRelationTypesMediaFeed
	case "memento":
		*code = LinkRelationTypesMemento
	case "micropub":
		*code = LinkRelationTypesMicropub
	case "modulepreload":
		*code = LinkRelationTypesModulepreload
	case "monitor":
		*code = LinkRelationTypesMonitor
	case "monitor-group":
		*code = LinkRelationTypesMonitorGroup
	case "next":
		*code = LinkRelationTypesNext
	case "next-archive":
		*code = LinkRelationTypesNextArchive
	case "nofollow":
		*code = LinkRelationTypesNofollow
	case "noopener":
		*code = LinkRelationTypesNoopener
	case "noreferrer":
		*code = LinkRelationTypesNoreferrer
	case "opener":
		*code = LinkRelationTypesOpener
	case "openid2.local_id":
		*code = LinkRelationTypesOpenid2_Local_id
	case "openid2.provider":
		*code = LinkRelationTypesOpenid2_Provider
	case "original":
		*code = LinkRelationTypesOriginal
	case "P3Pv1":
		*code = LinkRelationTypesP3Pv1
	case "payment":
		*code = LinkRelationTypesPayment
	case "pingback":
		*code = LinkRelationTypesPingback
	case "preconnect":
		*code = LinkRelationTypesPreconnect
	case "predecessor-version":
		*code = LinkRelationTypesPredecessorVersion
	case "prefetch":
		*code = LinkRelationTypesPrefetch
	case "preload":
		*code = LinkRelationTypesPreload
	case "prerender":
		*code = LinkRelationTypesPrerender
	case "prev":
		*code = LinkRelationTypesPrev
	case "preview":
		*code = LinkRelationTypesPreview
	case "previous":
		*code = LinkRelationTypesPrevious
	case "prev-archive":
		*code = LinkRelationTypesPrevArchive
	case "privacy-policy":
		*code = LinkRelationTypesPrivacyPolicy
	case "profile":
		*code = LinkRelationTypesProfile
	case "publication":
		*code = LinkRelationTypesPublication
	case "related":
		*code = LinkRelationTypesRelated
	case "restconf":
		*code = LinkRelationTypesRestconf
	case "replies":
		*code = LinkRelationTypesReplies
	case "ruleinput":
		*code = LinkRelationTypesRuleinput
	case "search":
		*code = LinkRelationTypesSearch
	case "section":
		*code = LinkRelationTypesSection
	case "self":
		*code = LinkRelationTypesSelf
	case "service":
		*code = LinkRelationTypesService
	case "service-desc":
		*code = LinkRelationTypesServiceDesc
	case "service-doc":
		*code = LinkRelationTypesServiceDoc
	case "service-meta":
		*code = LinkRelationTypesServiceMeta
	case "sponsored":
		*code = LinkRelationTypesSponsored
	case "start":
		*code = LinkRelationTypesStart
	case "status":
		*code = LinkRelationTypesStatus
	case "stylesheet":
		*code = LinkRelationTypesStylesheet
	case "subsection":
		*code = LinkRelationTypesSubsection
	case "successor-version":
		*code = LinkRelationTypesSuccessorVersion
	case "sunset":
		*code = LinkRelationTypesSunset
	case "tag":
		*code = LinkRelationTypesTag
	case "terms-of-service":
		*code = LinkRelationTypesTermsOfService
	case "timegate":
		*code = LinkRelationTypesTimegate
	case "timemap":
		*code = LinkRelationTypesTimemap
	case "type":
		*code = LinkRelationTypesType
	case "ugc":
		*code = LinkRelationTypesUgc
	case "up":
		*code = LinkRelationTypesUp
	case "version-history":
		*code = LinkRelationTypesVersionHistory
	case "via":
		*code = LinkRelationTypesVia
	case "webmention":
		*code = LinkRelationTypesWebmention
	case "working-copy":
		*code = LinkRelationTypesWorkingCopy
	case "working-copy-of":
		*code = LinkRelationTypesWorkingCopyOf
	default:
		return fmt.Errorf("unknown LinkRelationTypes code `%s`", s)
	}
	return nil
}

func (code LinkRelationTypes) String() string {
	return code.Code()
}

func (code LinkRelationTypes) Code() string {
	switch code {
	case LinkRelationTypesAbout:
		return "about"
	case LinkRelationTypesAcl:
		return "acl"
	case LinkRelationTypesAlternate:
		return "alternate"
	case LinkRelationTypesAmphtml:
		return "amphtml"
	case LinkRelationTypesAppendix:
		return "appendix"
	case LinkRelationTypesAppleTouchIcon:
		return "apple-touch-icon"
	case LinkRelationTypesAppleTouchStartupImage:
		return "apple-touch-startup-image"
	case LinkRelationTypesArchives:
		return "archives"
	case LinkRelationTypesAuthor:
		return "author"
	case LinkRelationTypesBlockedBy:
		return "blocked-by"
	case LinkRelationTypesBookmark:
		return "bookmark"
	case LinkRelationTypesCanonical:
		return "canonical"
	case LinkRelationTypesChapter:
		return "chapter"
	case LinkRelationTypesCiteAs:
		return "cite-as"
	case LinkRelationTypesCollection:
		return "collection"
	case LinkRelationTypesContents:
		return "contents"
	case LinkRelationTypesConvertedFrom:
		return "convertedFrom"
	case LinkRelationTypesCopyright:
		return "copyright"
	case LinkRelationTypesCreateForm:
		return "create-form"
	case LinkRelationTypesCurrent:
		return "current"
	case LinkRelationTypesDescribedby:
		return "describedby"
	case LinkRelationTypesDescribes:
		return "describes"
	case LinkRelationTypesDisclosure:
		return "disclosure"
	case LinkRelationTypesDnsPrefetch:
		return "dns-prefetch"
	case LinkRelationTypesDuplicate:
		return "duplicate"
	case LinkRelationTypesEdit:
		return "edit"
	case LinkRelationTypesEditForm:
		return "edit-form"
	case LinkRelationTypesEditMedia:
		return "edit-media"
	case LinkRelationTypesEnclosure:
		return "enclosure"
	case LinkRelationTypesExternal:
		return "external"
	case LinkRelationTypesFirst:
		return "first"
	case LinkRelationTypesGlossary:
		return "glossary"
	case LinkRelationTypesHelp:
		return "help"
	case LinkRelationTypesHosts:
		return "hosts"
	case LinkRelationTypesHub:
		return "hub"
	case LinkRelationTypesIcon:
		return "icon"
	case LinkRelationTypesIndex:
		return "index"
	case LinkRelationTypesIntervalAfter:
		return "intervalAfter"
	case LinkRelationTypesIntervalBefore:
		return "intervalBefore"
	case LinkRelationTypesIntervalContains:
		return "intervalContains"
	case LinkRelationTypesIntervalDisjoint:
		return "intervalDisjoint"
	case LinkRelationTypesIntervalDuring:
		return "intervalDuring"
	case LinkRelationTypesIntervalEquals:
		return "intervalEquals"
	case LinkRelationTypesIntervalFinishedBy:
		return "intervalFinishedBy"
	case LinkRelationTypesIntervalFinishes:
		return "intervalFinishes"
	case LinkRelationTypesIntervalIn:
		return "intervalIn"
	case LinkRelationTypesIntervalMeets:
		return "intervalMeets"
	case LinkRelationTypesIntervalMetBy:
		return "intervalMetBy"
	case LinkRelationTypesIntervalOverlappedBy:
		return "intervalOverlappedBy"
	case LinkRelationTypesIntervalOverlaps:
		return "intervalOverlaps"
	case LinkRelationTypesIntervalStartedBy:
		return "intervalStartedBy"
	case LinkRelationTypesIntervalStarts:
		return "intervalStarts"
	case LinkRelationTypesItem:
		return "item"
	case LinkRelationTypesLast:
		return "last"
	case LinkRelationTypesLatestVersion:
		return "latest-version"
	case LinkRelationTypesLicense:
		return "license"
	case LinkRelationTypesLinkset:
		return "linkset"
	case LinkRelationTypesLrdd:
		return "lrdd"
	case LinkRelationTypesManifest:
		return "manifest"
	case LinkRelationTypesMaskIcon:
		return "mask-icon"
	case LinkRelationTypesMediaFeed:
		return "media-feed"
	case LinkRelationTypesMemento:
		return "memento"
	case LinkRelationTypesMicropub:
		return "micropub"
	case LinkRelationTypesModulepreload:
		return "modulepreload"
	case LinkRelationTypesMonitor:
		return "monitor"
	case LinkRelationTypesMonitorGroup:
		return "monitor-group"
	case LinkRelationTypesNext:
		return "next"
	case LinkRelationTypesNextArchive:
		return "next-archive"
	case LinkRelationTypesNofollow:
		return "nofollow"
	case LinkRelationTypesNoopener:
		return "noopener"
	case LinkRelationTypesNoreferrer:
		return "noreferrer"
	case LinkRelationTypesOpener:
		return "opener"
	case LinkRelationTypesOpenid2_Local_id:
		return "openid2.local_id"
	case LinkRelationTypesOpenid2_Provider:
		return "openid2.provider"
	case LinkRelationTypesOriginal:
		return "original"
	case LinkRelationTypesP3Pv1:
		return "P3Pv1"
	case LinkRelationTypesPayment:
		return "payment"
	case LinkRelationTypesPingback:
		return "pingback"
	case LinkRelationTypesPreconnect:
		return "preconnect"
	case LinkRelationTypesPredecessorVersion:
		return "predecessor-version"
	case LinkRelationTypesPrefetch:
		return "prefetch"
	case LinkRelationTypesPreload:
		return "preload"
	case LinkRelationTypesPrerender:
		return "prerender"
	case LinkRelationTypesPrev:
		return "prev"
	case LinkRelationTypesPreview:
		return "preview"
	case LinkRelationTypesPrevious:
		return "previous"
	case LinkRelationTypesPrevArchive:
		return "prev-archive"
	case LinkRelationTypesPrivacyPolicy:
		return "privacy-policy"
	case LinkRelationTypesProfile:
		return "profile"
	case LinkRelationTypesPublication:
		return "publication"
	case LinkRelationTypesRelated:
		return "related"
	case LinkRelationTypesRestconf:
		return "restconf"
	case LinkRelationTypesReplies:
		return "replies"
	case LinkRelationTypesRuleinput:
		return "ruleinput"
	case LinkRelationTypesSearch:
		return "search"
	case LinkRelationTypesSection:
		return "section"
	case LinkRelationTypesSelf:
		return "self"
	case LinkRelationTypesService:
		return "service"
	case LinkRelationTypesServiceDesc:
		return "service-desc"
	case LinkRelationTypesServiceDoc:
		return "service-doc"
	case LinkRelationTypesServiceMeta:
		return "service-meta"
	case LinkRelationTypesSponsored:
		return "sponsored"
	case LinkRelationTypesStart:
		return "start"
	case LinkRelationTypesStatus:
		return "status"
	case LinkRelationTypesStylesheet:
		return "stylesheet"
	case LinkRelationTypesSubsection:
		return "subsection"
	case LinkRelationTypesSuccessorVersion:
		return "successor-version"
	case LinkRelationTypesSunset:
		return "sunset"
	case LinkRelationTypesTag:
		return "tag"
	case LinkRelationTypesTermsOfService:
		return "terms-of-service"
	case LinkRelationTypesTimegate:
		return "timegate"
	case LinkRelationTypesTimemap:
		return "timemap"
	case LinkRelationTypesType:
		return "type"
	case LinkRelationTypesUgc:
		return "ugc"
	case LinkRelationTypesUp:
		return "up"
	case LinkRelationTypesVersionHistory:
		return "version-history"
	case LinkRelationTypesVia:
		return "via"
	case LinkRelationTypesWebmention:
		return "webmention"
	case LinkRelationTypesWorkingCopy:
		return "working-copy"
	case LinkRelationTypesWorkingCopyOf:
		return "working-copy-of"
	}
	return "<unknown>"
}

func (code LinkRelationTypes) Display() string {
	switch code {
	case LinkRelationTypesAbout:
		return "Refers to a resource that is the subject of the link's context."
	case LinkRelationTypesAcl:
		return "Asserts that the link target provides an access control description for the link context."
	case LinkRelationTypesAlternate:
		return "Refers to a substitute for this context"
	case LinkRelationTypesAmphtml:
		return "Used to reference alternative content that uses the AMP profile of the HTML format."
	case LinkRelationTypesAppendix:
		return "Refers to an appendix."
	case LinkRelationTypesAppleTouchIcon:
		return "Refers to an icon for the context. Synonym for icon."
	case LinkRelationTypesAppleTouchStartupImage:
		return "Refers to a launch screen for the context."
	case LinkRelationTypesArchives:
		return "Refers to a collection of records, documents, or other\n      materials of historical interest."
	case LinkRelationTypesAuthor:
		return "Refers to the context's author."
	case LinkRelationTypesBlockedBy:
		return "Identifies the entity that blocks access to a resource\n      following receipt of a legal demand."
	case LinkRelationTypesBookmark:
		return "Gives a permanent link to use for bookmarking purposes."
	case LinkRelationTypesCanonical:
		return "Designates the preferred version of a resource (the IRI and its contents)."
	case LinkRelationTypesChapter:
		return "Refers to a chapter in a collection of resources."
	case LinkRelationTypesCiteAs:
		return "Indicates that the link target is preferred over the link context for the purpose of permanent citation."
	case LinkRelationTypesCollection:
		return "The target IRI points to a resource which represents the collection resource for the context IRI."
	case LinkRelationTypesContents:
		return "Refers to a table of contents."
	case LinkRelationTypesConvertedFrom:
		return "The document linked to was later converted to the\n      document that contains this link relation.  For example, an RFC can\n      have a link to the Internet-Draft that became the RFC; in that case,\n      the link relation would be \"convertedFrom\"."
	case LinkRelationTypesCopyright:
		return "Refers to a copyright statement that applies to the\n    link's context."
	case LinkRelationTypesCreateForm:
		return "The target IRI points to a resource where a submission form can be obtained."
	case LinkRelationTypesCurrent:
		return "Refers to a resource containing the most recent\n      item(s) in a collection of resources."
	case LinkRelationTypesDescribedby:
		return "Refers to a resource providing information about the\n      link's context."
	case LinkRelationTypesDescribes:
		return "The relationship A 'describes' B asserts that\n      resource A provides a description of resource B. There are no\n      constraints on the format or representation of either A or B,\n      neither are there any further constraints on either resource."
	case LinkRelationTypesDisclosure:
		return "Refers to a list of patent disclosures made with respect to \n      material for which 'disclosure' relation is specified."
	case LinkRelationTypesDnsPrefetch:
		return "Used to indicate an origin that will be used to fetch required \n      resources for the link context, and that the user agent ought to resolve \n      as early as possible."
	case LinkRelationTypesDuplicate:
		return "Refers to a resource whose available representations\n      are byte-for-byte identical with the corresponding representations of\n      the context IRI."
	case LinkRelationTypesEdit:
		return "Refers to a resource that can be used to edit the\n      link's context."
	case LinkRelationTypesEditForm:
		return "The target IRI points to a resource where a submission form for\n      editing associated resource can be obtained."
	case LinkRelationTypesEditMedia:
		return "Refers to a resource that can be used to edit media\n      associated with the link's context."
	case LinkRelationTypesEnclosure:
		return "Identifies a related resource that is potentially\n      large and might require special handling."
	case LinkRelationTypesExternal:
		return "Refers to a resource that is not part of the same site as the current context."
	case LinkRelationTypesFirst:
		return "An IRI that refers to the furthest preceding resource\n    in a series of resources."
	case LinkRelationTypesGlossary:
		return "Refers to a glossary of terms."
	case LinkRelationTypesHelp:
		return "Refers to context-sensitive help."
	case LinkRelationTypesHosts:
		return "Refers to a resource hosted by the server indicated by\n      the link context."
	case LinkRelationTypesHub:
		return "Refers to a hub that enables registration for\n    notification of updates to the context."
	case LinkRelationTypesIcon:
		return "Refers to an icon representing the link's context."
	case LinkRelationTypesIndex:
		return "Refers to an index."
	case LinkRelationTypesIntervalAfter:
		return "refers to a resource associated with a time interval that ends before the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalBefore:
		return "refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalContains:
		return "refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalDisjoint:
		return "refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource, or ends before the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalDuring:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalEquals:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalFinishedBy:
		return "refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalFinishes:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalIn:
		return "refers to a resource associated with a time interval that begins before or is coincident with the beginning of the time interval associated with the context resource, and ends after or is coincident with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalMeets:
		return "refers to a resource associated with a time interval whose beginning coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalMetBy:
		return "refers to a resource associated with a time interval whose end coincides with the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalOverlappedBy:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalOverlaps:
		return "refers to a resource associated with a time interval that begins before the end of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalStartedBy:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalStarts:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case LinkRelationTypesItem:
		return "The target IRI points to a resource that is a member of the collection represented by the context IRI."
	case LinkRelationTypesLast:
		return "An IRI that refers to the furthest following resource\n      in a series of resources."
	case LinkRelationTypesLatestVersion:
		return "Points to a resource containing the latest (e.g.,\n      current) version of the context."
	case LinkRelationTypesLicense:
		return "Refers to a license associated with this context."
	case LinkRelationTypesLinkset:
		return "The link target of a link with the \"linkset\" relation\n      type provides a set of links, including links in which the link\n      context of the link participates.\n    "
	case LinkRelationTypesLrdd:
		return "Refers to further information about the link's context,\n      expressed as a LRDD (\"Link-based Resource Descriptor Document\")\n      resource.  See  for information about\n      processing this relation type in host-meta documents. When used\n      elsewhere, it refers to additional links and other metadata.\n      Multiple instances indicate additional LRDD resources. LRDD\n      resources MUST have an \"application/xrd+xml\" representation, and\n      MAY have others."
	case LinkRelationTypesManifest:
		return "Links to a manifest file for the context."
	case LinkRelationTypesMaskIcon:
		return "Refers to a mask that can be applied to the icon for the context."
	case LinkRelationTypesMediaFeed:
		return "Refers to a feed of personalised media recommendations relevant to the link context."
	case LinkRelationTypesMemento:
		return "The Target IRI points to a Memento, a fixed resource that will not change state anymore."
	case LinkRelationTypesMicropub:
		return "Links to the context's Micropub endpoint."
	case LinkRelationTypesModulepreload:
		return "Refers to a module that the user agent is to preemptively fetch and store for use in the current context."
	case LinkRelationTypesMonitor:
		return "Refers to a resource that can be used to monitor changes in an HTTP resource.\n    "
	case LinkRelationTypesMonitorGroup:
		return "Refers to a resource that can be used to monitor changes in a specified group of HTTP resources.\n    "
	case LinkRelationTypesNext:
		return "Indicates that the link's context is a part of a series, and\n      that the next in the series is the link target.\n    "
	case LinkRelationTypesNextArchive:
		return "Refers to the immediately following archive resource."
	case LinkRelationTypesNofollow:
		return "Indicates that the context’s original author or publisher does not endorse the link target."
	case LinkRelationTypesNoopener:
		return "Indicates that any newly created top-level browsing context which results from following the link will not be an auxiliary browsing context."
	case LinkRelationTypesNoreferrer:
		return "Indicates that no referrer information is to be leaked when following the link."
	case LinkRelationTypesOpener:
		return "Indicates that any newly created top-level browsing context which results from following the link will be an auxiliary browsing context."
	case LinkRelationTypesOpenid2_Local_id:
		return "Refers to an OpenID Authentication server on which the context relies for an assertion that the end user controls an Identifier."
	case LinkRelationTypesOpenid2_Provider:
		return "Refers to a resource which accepts OpenID Authentication protocol messages for the context."
	case LinkRelationTypesOriginal:
		return "The Target IRI points to an Original Resource."
	case LinkRelationTypesP3Pv1:
		return "Refers to a P3P privacy policy for the context."
	case LinkRelationTypesPayment:
		return "Indicates a resource where payment is accepted."
	case LinkRelationTypesPingback:
		return "Gives the address of the pingback resource for the link context."
	case LinkRelationTypesPreconnect:
		return "Used to indicate an origin that will be used to fetch required \n      resources for the link context. Initiating an early connection, which \n      includes the DNS lookup, TCP handshake, and optional TLS negotiation, \n      allows the user agent to mask the high latency costs of establishing a \n      connection."
	case LinkRelationTypesPredecessorVersion:
		return "Points to a resource containing the predecessor\n      version in the version history.\n    "
	case LinkRelationTypesPrefetch:
		return "The prefetch link relation type is used to identify a resource \n      that might be required by the next navigation from the link context, and \n      that the user agent ought to fetch, such that the user agent can deliver a \n      faster response once the resource is requested in the future."
	case LinkRelationTypesPreload:
		return "Refers to a resource that should be loaded early in the \n      processing of the link's context, without blocking rendering."
	case LinkRelationTypesPrerender:
		return "Used to identify a resource that might be required by the next \n      navigation from the link context, and that the user agent ought to fetch \n      and execute, such that the user agent can deliver a faster response once \n      the resource is requested in the future."
	case LinkRelationTypesPrev:
		return "Indicates that the link's context is a part of a series, and\n      that the previous in the series is the link target.\n    "
	case LinkRelationTypesPreview:
		return "Refers to a resource that provides a preview of the link's context."
	case LinkRelationTypesPrevious:
		return "Refers to the previous resource in an ordered series\n      of resources.  Synonym for \"prev\"."
	case LinkRelationTypesPrevArchive:
		return "Refers to the immediately preceding archive resource."
	case LinkRelationTypesPrivacyPolicy:
		return "Refers to a privacy policy associated with the link's context."
	case LinkRelationTypesProfile:
		return "Identifying that a resource representation conforms\nto a certain profile, without affecting the non-profile semantics\nof the resource representation."
	case LinkRelationTypesPublication:
		return "Links to a publication manifest. A manifest represents \n      structured information about a publication, such as informative metadata, \n      a list of resources, and a default reading order."
	case LinkRelationTypesRelated:
		return "Identifies a related resource."
	case LinkRelationTypesRestconf:
		return "Identifies the root of RESTCONF API as configured on this HTTP server.\n      The \"restconf\" relation defines the root of the API defined in RFC8040.\n      Subsequent revisions of RESTCONF will use alternate relation values to support \n      protocol versioning."
	case LinkRelationTypesReplies:
		return "Identifies a resource that is a reply to the context\n      of the link.\n    "
	case LinkRelationTypesRuleinput:
		return "The resource identified by the link target provides an input value to an \n    instance of a rule, where the resource which represents the rule instance is \n    identified by the link context.\n    "
	case LinkRelationTypesSearch:
		return "Refers to a resource that can be used to search through\n      the link's context and related resources."
	case LinkRelationTypesSection:
		return "Refers to a section in a collection of resources."
	case LinkRelationTypesSelf:
		return "Conveys an identifier for the link's context.\n    "
	case LinkRelationTypesService:
		return "Indicates a URI that can be used to retrieve a\n      service document."
	case LinkRelationTypesServiceDesc:
		return "Identifies service description for the context that\n      is primarily intended for consumption by machines."
	case LinkRelationTypesServiceDoc:
		return "Identifies service documentation for the context that\n      is primarily intended for human consumption."
	case LinkRelationTypesServiceMeta:
		return "Identifies general metadata for the context that is\n      primarily intended for consumption by machines."
	case LinkRelationTypesSponsored:
		return "Refers to a resource that is within a context that is \n\t\tsponsored (such as advertising or another compensation agreement)."
	case LinkRelationTypesStart:
		return "Refers to the first resource in a collection of\n      resources."
	case LinkRelationTypesStatus:
		return "Identifies a resource that represents the context's\n      status."
	case LinkRelationTypesStylesheet:
		return "Refers to a stylesheet."
	case LinkRelationTypesSubsection:
		return "Refers to a resource serving as a subsection in a\n      collection of resources."
	case LinkRelationTypesSuccessorVersion:
		return "Points to a resource containing the successor version\n      in the version history.\n    "
	case LinkRelationTypesSunset:
		return "Identifies a resource that provides information about\n      the context's retirement policy.\n    "
	case LinkRelationTypesTag:
		return "Gives a tag (identified by the given address) that applies to\n      the current document.\n    "
	case LinkRelationTypesTermsOfService:
		return "Refers to the terms of service associated with the link's context."
	case LinkRelationTypesTimegate:
		return "The Target IRI points to a TimeGate for an Original Resource."
	case LinkRelationTypesTimemap:
		return "The Target IRI points to a TimeMap for an Original Resource."
	case LinkRelationTypesType:
		return "Refers to a resource identifying the abstract semantic type of which the link's context is considered to be an instance."
	case LinkRelationTypesUgc:
		return "Refers to a resource that is within a context that is User Generated Content.\n    "
	case LinkRelationTypesUp:
		return "Refers to a parent document in a hierarchy of\n      documents.\n    "
	case LinkRelationTypesVersionHistory:
		return "Points to a resource containing the version history\n      for the context.\n    "
	case LinkRelationTypesVia:
		return "Identifies a resource that is the source of the\n      information in the link's context.\n    "
	case LinkRelationTypesWebmention:
		return "Identifies a target URI that supports the Webmention protocol.\n    This allows clients that mention a resource in some form of publishing process\n    to contact that endpoint and inform it that this resource has been mentioned."
	case LinkRelationTypesWorkingCopy:
		return "Points to a working copy for this resource."
	case LinkRelationTypesWorkingCopyOf:
		return "Points to the versioned resource from which this\n      working copy was obtained.\n    "
	}
	return "<unknown>"
}

func (code LinkRelationTypes) Definition() string {
	switch code {
	case LinkRelationTypesAbout:
		return "Refers to a resource that is the subject of the link's context."
	case LinkRelationTypesAcl:
		return "Asserts that the link target provides an access control description for the link context."
	case LinkRelationTypesAlternate:
		return "Refers to a substitute for this context"
	case LinkRelationTypesAmphtml:
		return "Used to reference alternative content that uses the AMP profile of the HTML format."
	case LinkRelationTypesAppendix:
		return "Refers to an appendix."
	case LinkRelationTypesAppleTouchIcon:
		return "Refers to an icon for the context. Synonym for icon."
	case LinkRelationTypesAppleTouchStartupImage:
		return "Refers to a launch screen for the context."
	case LinkRelationTypesArchives:
		return "Refers to a collection of records, documents, or other\n      materials of historical interest."
	case LinkRelationTypesAuthor:
		return "Refers to the context's author."
	case LinkRelationTypesBlockedBy:
		return "Identifies the entity that blocks access to a resource\n      following receipt of a legal demand."
	case LinkRelationTypesBookmark:
		return "Gives a permanent link to use for bookmarking purposes."
	case LinkRelationTypesCanonical:
		return "Designates the preferred version of a resource (the IRI and its contents)."
	case LinkRelationTypesChapter:
		return "Refers to a chapter in a collection of resources."
	case LinkRelationTypesCiteAs:
		return "Indicates that the link target is preferred over the link context for the purpose of permanent citation."
	case LinkRelationTypesCollection:
		return "The target IRI points to a resource which represents the collection resource for the context IRI."
	case LinkRelationTypesContents:
		return "Refers to a table of contents."
	case LinkRelationTypesConvertedFrom:
		return "The document linked to was later converted to the\n      document that contains this link relation.  For example, an RFC can\n      have a link to the Internet-Draft that became the RFC; in that case,\n      the link relation would be \"convertedFrom\"."
	case LinkRelationTypesCopyright:
		return "Refers to a copyright statement that applies to the\n    link's context."
	case LinkRelationTypesCreateForm:
		return "The target IRI points to a resource where a submission form can be obtained."
	case LinkRelationTypesCurrent:
		return "Refers to a resource containing the most recent\n      item(s) in a collection of resources."
	case LinkRelationTypesDescribedby:
		return "Refers to a resource providing information about the\n      link's context."
	case LinkRelationTypesDescribes:
		return "The relationship A 'describes' B asserts that\n      resource A provides a description of resource B. There are no\n      constraints on the format or representation of either A or B,\n      neither are there any further constraints on either resource."
	case LinkRelationTypesDisclosure:
		return "Refers to a list of patent disclosures made with respect to \n      material for which 'disclosure' relation is specified."
	case LinkRelationTypesDnsPrefetch:
		return "Used to indicate an origin that will be used to fetch required \n      resources for the link context, and that the user agent ought to resolve \n      as early as possible."
	case LinkRelationTypesDuplicate:
		return "Refers to a resource whose available representations\n      are byte-for-byte identical with the corresponding representations of\n      the context IRI."
	case LinkRelationTypesEdit:
		return "Refers to a resource that can be used to edit the\n      link's context."
	case LinkRelationTypesEditForm:
		return "The target IRI points to a resource where a submission form for\n      editing associated resource can be obtained."
	case LinkRelationTypesEditMedia:
		return "Refers to a resource that can be used to edit media\n      associated with the link's context."
	case LinkRelationTypesEnclosure:
		return "Identifies a related resource that is potentially\n      large and might require special handling."
	case LinkRelationTypesExternal:
		return "Refers to a resource that is not part of the same site as the current context."
	case LinkRelationTypesFirst:
		return "An IRI that refers to the furthest preceding resource\n    in a series of resources."
	case LinkRelationTypesGlossary:
		return "Refers to a glossary of terms."
	case LinkRelationTypesHelp:
		return "Refers to context-sensitive help."
	case LinkRelationTypesHosts:
		return "Refers to a resource hosted by the server indicated by\n      the link context."
	case LinkRelationTypesHub:
		return "Refers to a hub that enables registration for\n    notification of updates to the context."
	case LinkRelationTypesIcon:
		return "Refers to an icon representing the link's context."
	case LinkRelationTypesIndex:
		return "Refers to an index."
	case LinkRelationTypesIntervalAfter:
		return "refers to a resource associated with a time interval that ends before the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalBefore:
		return "refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalContains:
		return "refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalDisjoint:
		return "refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource, or ends before the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalDuring:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalEquals:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalFinishedBy:
		return "refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalFinishes:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalIn:
		return "refers to a resource associated with a time interval that begins before or is coincident with the beginning of the time interval associated with the context resource, and ends after or is coincident with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalMeets:
		return "refers to a resource associated with a time interval whose beginning coincides with the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalMetBy:
		return "refers to a resource associated with a time interval whose end coincides with the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalOverlappedBy:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the beginning of the time interval associated with the context resource"
	case LinkRelationTypesIntervalOverlaps:
		return "refers to a resource associated with a time interval that begins before the end of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalStartedBy:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource"
	case LinkRelationTypesIntervalStarts:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case LinkRelationTypesItem:
		return "The target IRI points to a resource that is a member of the collection represented by the context IRI."
	case LinkRelationTypesLast:
		return "An IRI that refers to the furthest following resource\n      in a series of resources."
	case LinkRelationTypesLatestVersion:
		return "Points to a resource containing the latest (e.g.,\n      current) version of the context."
	case LinkRelationTypesLicense:
		return "Refers to a license associated with this context."
	case LinkRelationTypesLinkset:
		return "The link target of a link with the \"linkset\" relation\n      type provides a set of links, including links in which the link\n      context of the link participates.\n    "
	case LinkRelationTypesLrdd:
		return "Refers to further information about the link's context,\n      expressed as a LRDD (\"Link-based Resource Descriptor Document\")\n      resource.  See  for information about\n      processing this relation type in host-meta documents. When used\n      elsewhere, it refers to additional links and other metadata.\n      Multiple instances indicate additional LRDD resources. LRDD\n      resources MUST have an \"application/xrd+xml\" representation, and\n      MAY have others."
	case LinkRelationTypesManifest:
		return "Links to a manifest file for the context."
	case LinkRelationTypesMaskIcon:
		return "Refers to a mask that can be applied to the icon for the context."
	case LinkRelationTypesMediaFeed:
		return "Refers to a feed of personalised media recommendations relevant to the link context."
	case LinkRelationTypesMemento:
		return "The Target IRI points to a Memento, a fixed resource that will not change state anymore."
	case LinkRelationTypesMicropub:
		return "Links to the context's Micropub endpoint."
	case LinkRelationTypesModulepreload:
		return "Refers to a module that the user agent is to preemptively fetch and store for use in the current context."
	case LinkRelationTypesMonitor:
		return "Refers to a resource that can be used to monitor changes in an HTTP resource.\n    "
	case LinkRelationTypesMonitorGroup:
		return "Refers to a resource that can be used to monitor changes in a specified group of HTTP resources.\n    "
	case LinkRelationTypesNext:
		return "Indicates that the link's context is a part of a series, and\n      that the next in the series is the link target.\n    "
	case LinkRelationTypesNextArchive:
		return "Refers to the immediately following archive resource."
	case LinkRelationTypesNofollow:
		return "Indicates that the context’s original author or publisher does not endorse the link target."
	case LinkRelationTypesNoopener:
		return "Indicates that any newly created top-level browsing context which results from following the link will not be an auxiliary browsing context."
	case LinkRelationTypesNoreferrer:
		return "Indicates that no referrer information is to be leaked when following the link."
	case LinkRelationTypesOpener:
		return "Indicates that any newly created top-level browsing context which results from following the link will be an auxiliary browsing context."
	case LinkRelationTypesOpenid2_Local_id:
		return "Refers to an OpenID Authentication server on which the context relies for an assertion that the end user controls an Identifier."
	case LinkRelationTypesOpenid2_Provider:
		return "Refers to a resource which accepts OpenID Authentication protocol messages for the context."
	case LinkRelationTypesOriginal:
		return "The Target IRI points to an Original Resource."
	case LinkRelationTypesP3Pv1:
		return "Refers to a P3P privacy policy for the context."
	case LinkRelationTypesPayment:
		return "Indicates a resource where payment is accepted."
	case LinkRelationTypesPingback:
		return "Gives the address of the pingback resource for the link context."
	case LinkRelationTypesPreconnect:
		return "Used to indicate an origin that will be used to fetch required \n      resources for the link context. Initiating an early connection, which \n      includes the DNS lookup, TCP handshake, and optional TLS negotiation, \n      allows the user agent to mask the high latency costs of establishing a \n      connection."
	case LinkRelationTypesPredecessorVersion:
		return "Points to a resource containing the predecessor\n      version in the version history.\n    "
	case LinkRelationTypesPrefetch:
		return "The prefetch link relation type is used to identify a resource \n      that might be required by the next navigation from the link context, and \n      that the user agent ought to fetch, such that the user agent can deliver a \n      faster response once the resource is requested in the future."
	case LinkRelationTypesPreload:
		return "Refers to a resource that should be loaded early in the \n      processing of the link's context, without blocking rendering."
	case LinkRelationTypesPrerender:
		return "Used to identify a resource that might be required by the next \n      navigation from the link context, and that the user agent ought to fetch \n      and execute, such that the user agent can deliver a faster response once \n      the resource is requested in the future."
	case LinkRelationTypesPrev:
		return "Indicates that the link's context is a part of a series, and\n      that the previous in the series is the link target.\n    "
	case LinkRelationTypesPreview:
		return "Refers to a resource that provides a preview of the link's context."
	case LinkRelationTypesPrevious:
		return "Refers to the previous resource in an ordered series\n      of resources.  Synonym for \"prev\"."
	case LinkRelationTypesPrevArchive:
		return "Refers to the immediately preceding archive resource."
	case LinkRelationTypesPrivacyPolicy:
		return "Refers to a privacy policy associated with the link's context."
	case LinkRelationTypesProfile:
		return "Identifying that a resource representation conforms\nto a certain profile, without affecting the non-profile semantics\nof the resource representation."
	case LinkRelationTypesPublication:
		return "Links to a publication manifest. A manifest represents \n      structured information about a publication, such as informative metadata, \n      a list of resources, and a default reading order."
	case LinkRelationTypesRelated:
		return "Identifies a related resource."
	case LinkRelationTypesRestconf:
		return "Identifies the root of RESTCONF API as configured on this HTTP server.\n      The \"restconf\" relation defines the root of the API defined in RFC8040.\n      Subsequent revisions of RESTCONF will use alternate relation values to support \n      protocol versioning."
	case LinkRelationTypesReplies:
		return "Identifies a resource that is a reply to the context\n      of the link.\n    "
	case LinkRelationTypesRuleinput:
		return "The resource identified by the link target provides an input value to an \n    instance of a rule, where the resource which represents the rule instance is \n    identified by the link context.\n    "
	case LinkRelationTypesSearch:
		return "Refers to a resource that can be used to search through\n      the link's context and related resources."
	case LinkRelationTypesSection:
		return "Refers to a section in a collection of resources."
	case LinkRelationTypesSelf:
		return "Conveys an identifier for the link's context.\n    "
	case LinkRelationTypesService:
		return "Indicates a URI that can be used to retrieve a\n      service document."
	case LinkRelationTypesServiceDesc:
		return "Identifies service description for the context that\n      is primarily intended for consumption by machines."
	case LinkRelationTypesServiceDoc:
		return "Identifies service documentation for the context that\n      is primarily intended for human consumption."
	case LinkRelationTypesServiceMeta:
		return "Identifies general metadata for the context that is\n      primarily intended for consumption by machines."
	case LinkRelationTypesSponsored:
		return "Refers to a resource that is within a context that is \n\t\tsponsored (such as advertising or another compensation agreement)."
	case LinkRelationTypesStart:
		return "Refers to the first resource in a collection of\n      resources."
	case LinkRelationTypesStatus:
		return "Identifies a resource that represents the context's\n      status."
	case LinkRelationTypesStylesheet:
		return "Refers to a stylesheet."
	case LinkRelationTypesSubsection:
		return "Refers to a resource serving as a subsection in a\n      collection of resources."
	case LinkRelationTypesSuccessorVersion:
		return "Points to a resource containing the successor version\n      in the version history.\n    "
	case LinkRelationTypesSunset:
		return "Identifies a resource that provides information about\n      the context's retirement policy.\n    "
	case LinkRelationTypesTag:
		return "Gives a tag (identified by the given address) that applies to\n      the current document.\n    "
	case LinkRelationTypesTermsOfService:
		return "Refers to the terms of service associated with the link's context."
	case LinkRelationTypesTimegate:
		return "The Target IRI points to a TimeGate for an Original Resource."
	case LinkRelationTypesTimemap:
		return "The Target IRI points to a TimeMap for an Original Resource."
	case LinkRelationTypesType:
		return "Refers to a resource identifying the abstract semantic type of which the link's context is considered to be an instance."
	case LinkRelationTypesUgc:
		return "Refers to a resource that is within a context that is User Generated Content.\n    "
	case LinkRelationTypesUp:
		return "Refers to a parent document in a hierarchy of\n      documents.\n    "
	case LinkRelationTypesVersionHistory:
		return "Points to a resource containing the version history\n      for the context.\n    "
	case LinkRelationTypesVia:
		return "Identifies a resource that is the source of the\n      information in the link's context.\n    "
	case LinkRelationTypesWebmention:
		return "Identifies a target URI that supports the Webmention protocol.\n    This allows clients that mention a resource in some form of publishing process\n    to contact that endpoint and inform it that this resource has been mentioned."
	case LinkRelationTypesWorkingCopy:
		return "Points to a working copy for this resource."
	case LinkRelationTypesWorkingCopyOf:
		return "Points to the versioned resource from which this\n      working copy was obtained.\n    "
	}
	return "<unknown>"
}
