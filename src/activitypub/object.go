package activitypub

import (
	"encoding/json"
	"time"
)

type ObjectId string

const (
	ActivityBaseURI          URI                    = URI("https://www.w3.org/ns/activitystreams#")
	ObjectType               ActivityVocabularyType = "APObject"
	LinkType                 ActivityVocabularyType = "Link"
	ActivityType             ActivityVocabularyType = "Activity"
	IntransitiveActivityType ActivityVocabularyType = "IntransitiveActivity"
	ActorType                ActivityVocabularyType = "Actor"
	CollectionType           ActivityVocabularyType = "Collection"
	OrderedCollectionType    ActivityVocabularyType = "OrderedCollection"

	// APObject Types
	ArticleType      ActivityVocabularyType = "Article"
	AudioType        ActivityVocabularyType = "Audio"
	DocumentType     ActivityVocabularyType = "Document"
	EventType        ActivityVocabularyType = "Event"
	ImageType        ActivityVocabularyType = "Image"
	NoteType         ActivityVocabularyType = "Note"
	PageType         ActivityVocabularyType = "Page"
	PlaceType        ActivityVocabularyType = "Place"
	ProfileType      ActivityVocabularyType = "Profile"
	RelationshipType ActivityVocabularyType = "Relationship"
	TombstoneType    ActivityVocabularyType = "Tombstone"
	VideoType        ActivityVocabularyType = "Video"

	// Link Types
	MentionType ActivityVocabularyType = "Mention"
)

var validGenericObjectTypes = [...]ActivityVocabularyType{
	ActivityType,
	IntransitiveActivityType,
	ObjectType,
	ActorType,
	CollectionType,
	OrderedCollectionType,
}

var validGenericLinkTypes = [...]ActivityVocabularyType{
	LinkType,
}

var validGenericTypes = append(validGenericObjectTypes[:], validGenericLinkTypes[:]...)

var validObjectTypes = [...]ActivityVocabularyType{
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

type (
	ActivityVocabularyType string
	ActivityObject         interface{}
	ObjectOrLink           interface {
		IsLink() bool
		IsObject() bool
	}
	LinkOrUri            interface{}
	ImageOrLink          interface{}
	MimeType             string
	LangRef              string
	NaturalLanguageValue map[LangRef]string
)

func (o *APObject) IsLink() bool {
	return ValidLinkType(o.Type)
}

func (o *APObject) IsObject() bool {
	return ValidObjectType(o.Type)
}

func (n NaturalLanguageValue) MarshalJSON() ([]byte, error) {
	if len(n) == 1 {
		for _, v := range n {
			return json.Marshal(v)
		}
	}

	return json.Marshal(map[LangRef]string(n))
}

// Describes an object of any kind.
// The APObject type serves as the base type for most of the other kinds of objects defined in the Activity Vocabulary,
//  including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.
type APObject struct {
	// Provides the globally unique identifier for an APObject or Link.
	Id ObjectId `jsonld:"id,omitempty"`
	//  Identifies the APObject or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
	// Identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment ObjectOrLink `jsonld:"attachment,omitempty"`
	// Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo ObjectOrLink `jsonld:"attributedTo,omitempty"`
	// Identifies one or more entities that represent the total population of entities
	//  for which the object can considered to be relevant.
	Audience ObjectOrLink `jsonld:"audience,omitempty"`
	// The content or textual representation of the APObject encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValue `jsonld:"content,omitempty,collapsible"`
	// Identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	//  common originating context or purpose. An example could be all activities relating to a common project or event.
	//Context ObjectOrLink `jsonld:"_"`
	// The date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	//  the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Identifies the entity (e.g. an application) that generated the object.
	Generator ObjectOrLink `jsonld:"generator,omitempty"`
	// Indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	//  and should be suitable for presentation at a small size.
	Icon ImageOrLink `jsonld:"icon,omitempty"`
	// Indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image ImageOrLink `jsonld:"image,omitempty"`
	// Indicates one or more entities for which this object is considered a response.
	InReplyTo ObjectOrLink `jsonld:"inReplyTo,omitempty"`
	// Indicates one or more physical or logical locations associated with the object.
	Location ObjectOrLink `jsonld:"location,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview ObjectOrLink `jsonld:"preview,omitempty"`
	// The date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Identifies a Collection containing objects considered to be responses to this object.
	Replies Collection `jsonld:"replies,omitempty"`
	// The date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	//  the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// A natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValue `jsonld:"summary,omitempty,collapsible"`
	// One or more "tags" that have been associated with an objects. A tag can be any kind of APObject.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	//  while the latter implies associated by reference.
	Tag ObjectOrLink `jsonld:"tag,omitempty"`
	// The date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// Identifies one or more links to representations of the object
	Url LinkOrUri `jsonld:"url,omitempty"`
	// Identifies an entity considered to be part of the public primary audience of an APObject
	To ObjectOrLink `jsonld:"to,omitempty"`
	// Identifies an APObject that is part of the private primary audience of this APObject.
	Bto ObjectOrLink `jsonld:"bto,omitempty"`
	// Identifies an APObject that is part of the public secondary audience of this APObject.
	Cc ObjectOrLink `jsonld:"cc,omitempty"`
	// Identifies one or more Objects that are part of the private secondary audience of this APObject.
	Bcc ObjectOrLink `jsonld:"bcc,omitempty"`
	// When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	//  the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	//  section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
}

type ContentType string

type Source struct {
	Content   ContentType
	MediaType string
}

func ValidGenericType(_type ActivityVocabularyType) bool {
	for _, v := range validGenericObjectTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ValidObjectType(_type ActivityVocabularyType) bool {
	for _, v := range validObjectTypes {
		if v == _type {
			return true
		}
	}
	return ValidActivityType(_type) || ValidActorType(_type) || ValidCollectionType(_type) || ValidGenericType(_type)
}

func ObjectNew(id ObjectId, _type ActivityVocabularyType) *APObject {
	if !(ValidObjectType(_type)) {
		_type = ObjectType
	}
	o := APObject{Id: id, Type: _type}
	o.Name = make(NaturalLanguageValue)
	o.Content = make(NaturalLanguageValue)
	o.Summary = make(NaturalLanguageValue)
	return &o
}
