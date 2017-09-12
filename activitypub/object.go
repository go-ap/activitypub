package activitypub

import (
	"time"
)

type ObjectId string

const (
	ActivityBaseURI          URI    = URI("https://www.w3.org/ns/activitystreams#")
	ObjectType               string = "Object"
	LinkType                 string = "Link"
	ActivityType             string = "Activity"
	IntransitiveActivityType string = "IntransitiveActivity"
	ActorType                string = "Actor"
	CollectionType           string = "Collection"
	OrderedCollectionType    string = "OrderedCollection"

	// Object Types
	ArticleType string = "Article"
	AudioType string = "Audio"
	DocumentType string = "Document"
	EventType string = "Event"
	ImageType string = "Image"
	NoteType string = "Note"
	PageType string = "Page"
	PlaceType string = "Place"
	ProfileType string = "Profile"
	RelationshipType string = "Relationship"
	TombstoneType string = "Tombstone"
	VideoType string = "Video"

	// Link Types
	MentionType string = "Mention"
)

var validObjectTypes = [...]string{
	ArticleType,
	AudioType,
	DocumentType,
	EventType,
	ImageType,
	NoteType,
	PageType,
	PlaceType,
	ProfileType,
	RelationshipType,
	TombstoneType,
	VideoType,
}

var validLinkTypes = [...]string{
	MentionType,
}

type ObjectOrLink interface{}
type LinkOrUri interface{}
type ImageOrLink interface{}

type MimeType string
type LangRef string
type NaturalLanguageValue map[LangRef]string

// Describes an object of any kind.
// The Object type serves as the base type for most of the other kinds of objects defined in the Activity Vocabulary,
//  including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.
type BaseObject struct {
	// Provides the globally unique identifier for an Object or Link.
	Id ObjectId `jsonld:"id"`
	//  Identifies the Object or Link type. Multiple values may be specified.
	Type string `jsonld:"type"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name"`
}

type Object struct {
	*BaseObject
	// Identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment		ObjectOrLink			`jsonld:"attachment"`
	// Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo	ObjectOrLink			`jsonld:"attributedTo"`
	// Identifies one or more entities that represent the total population of entities
	//  for which the object can considered to be relevant.
	Audience 		ObjectOrLink			`jsonld:"audience"`
	// The content or textual representation of the Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content 		NaturalLanguageValue	`jsonld:"content"`
	// Identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	//  common originating context or purpose. An example could be all activities relating to a common project or event.
	Context 		ObjectOrLink			`jsonld:"context"`
	// The date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	//  the activity concluded or is expected to conclude.
	EndTime 		time.Time				`jsonld:"endTime"`
	// Identifies the entity (e.g. an application) that generated the object.
	Generator 		ObjectOrLink			`jsonld:"generator"`
	// Indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	//  and should be suitable for presentation at a small size.
	Icon 			ImageOrLink				`jsonld:"icon"`
	// Indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image 			ImageOrLink				`jsonld:"image"`
	// Indicates one or more entities for which this object is considered a response.
	InReplyTo 		ObjectOrLink			`jsonld:"inReplyTo"`
	// Indicates one or more physical or logical locations associated with the object.
	Location 		ObjectOrLink			`jsonld:"location"`
	// Identifies an entity that provides a preview of this object.
	Preview 		ObjectOrLink			`jsonld:"preview"`
	// The date and time at which the object was published
	Published 		time.Time				`jsonld:"published"`
	// Identifies a Collection containing objects considered to be responses to this object.
	Replies 		Collection				`jsonld:"replies"`
	// The date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	//  the moment the activity began or is scheduled to begin.
	StartTime 		time.Time				`jsonld:"startTime"`
	// A natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary 		NaturalLanguageValue	`jsonld:"summary"`
	// One or more "tags" that have been associated with an objects. A tag can be any kind of Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	//  while the latter implies associated by reference.
	Tag 			ObjectOrLink			`jsonld:"tag"`
	// The date and time at which the object was updated
	Updated 		time.Time				`jsonld:"updated"`
	// Identifies one or more links to representations of the object
	Url 			LinkOrUri				`jsonld:"url"`
	// Identifies an entity considered to be part of the public primary audience of an Object
	To 				ObjectOrLink			`jsonld:"to"`
	// Identifies an Object that is part of the private primary audience of this Object.
	Bto 			ObjectOrLink			`jsonld:"bto"`
	// Identifies an Object that is part of the public secondary audience of this Object.
	Cc 				ObjectOrLink			`jsonld:"cc"`
	// Identifies one or more Objects that are part of the private secondary audience of this Object.
	Bcc 			ObjectOrLink			`jsonld:"bcc"`
	// When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	//  the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	//  section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration 		time.Duration			`jsonld:"duration"`
}

// A Link is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988].
// Many of the properties defined by the Activity Vocabulary allow values that are either instances of Object or Link.
// When a Link is used, it establishes a qualified relation connecting the subject
//  (the containing object) to the resource identified by the href.
// Properties of the Link are properties of the reference as opposed to properties of the resource.
type Link struct {
	*BaseObject
	// A link relation associated with a Link. The value must conform to both the [HTML5] and
	//  [RFC5988](https://tools.ietf.org/html/rfc5988) "link relation" definitions.
	// In the [HTML5], any string not containing the "space" U+0020, "tab" (U+0009), "LF" (U+000A),
	//  "FF" (U+000C), "CR" (U+000D) or "," (U+002C) characters can be used as a valid link relation.
	Rel 			*Link					`jsonld:"rel"`
	// When used on a Link, identifies the MIME media type of the referenced resource.
	// When used on an Object, identifies the MIME media type of the value of the content property.
	// If not specified, the content property is assumed to contain text/html content.
	MediaType		MimeType				`jsonld:"mediaType"`
	// On a Link, specifies a hint as to the rendering height in device-independent pixels of the linked resource.
	Height			uint					`jsonld:"height"`
	// On a Link, specifies a hint as to the rendering width in device-independent pixels of the linked resource.
	Width			uint					`jsonld:"width"`
	// Identifies an entity that provides a preview of this object.
	Preview			ObjectOrLink			`jsonld:"preview"`
	// The target resource pointed to by a Link.
	Href URI `jsonld:"href"`
	// Hints as to the language used by the target resource.
	// Value must be a [BCP47](https://tools.ietf.org/html/bcp47) Language-Tag.
	HrefLang   		LangRef                	`jsonld:"hrefLang"`
}

type ContentType string

type Source struct {
	Content   ContentType
	MediaType string
}

func ValidObjectType(_type string) bool {
	for _, v := range validObjectTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ValidLinkType(_type string) bool {
	for _, v := range validLinkTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ObjectNew(id ObjectId, _type string) *Object {
	if !ValidObjectType(_type) {
		_type = ObjectType
	}
	p := BaseObject{Id: id, Type:_type}
	return &Object{BaseObject: &p}
}

func LinkNew(id ObjectId, _type string) *Link {
	if !ValidLinkType(_type) {
		_type = LinkType
	}
	p := BaseObject{Id: id, Type:_type}
	return &Link{BaseObject: &p}
}

