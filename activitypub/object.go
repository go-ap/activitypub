package activitypub

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/buger/jsonparser"
)

// ObjectID designates an unique global identifier.
// All Objects in [ActivityStreams] should have unique global identifiers.
// ActivityPub extends this requirement; all objects distributed by the ActivityPub protocol MUST
// have unique global identifiers, unless they are intentionally transient
// (short lived activities that are not intended to be able to be looked up,
// such as some kinds of chat messages or game notifications).
// These identifiers must fall into one of the following groups:
//
// 1. Publicly dereferencable URIs, such as HTTPS URIs, with their authority belonging
// to that of their originating server. (Publicly facing content SHOULD use HTTPS URIs).
// 2. An ID explicitly specified as the JSON null object, which implies an anonymous object
// (a part of its parent context)
type ObjectID string

const (
	// ActivityBaseURI the basic URI for the activity streams namespaces
	ActivityBaseURI                                 = URI("https://www.w3.org/ns/activitystreams#")
	ObjectType               ActivityVocabularyType = "Object"
	LinkType                 ActivityVocabularyType = "Link"
	ActivityType             ActivityVocabularyType = "Activity"
	IntransitiveActivityType ActivityVocabularyType = "IntransitiveActivity"
	ActorType                ActivityVocabularyType = "Actor"
	CollectionType           ActivityVocabularyType = "Collection"
	OrderedCollectionType    ActivityVocabularyType = "OrderedCollection"

	// Activity Pub GetID Types
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

	// MentionType is a link type for @mentions
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
	// ActivityVocabularyType is the data type for an Activity type object
	ActivityVocabularyType string
	// ActivityObject is a subtype of GetID that describes some form of action that may happen,
	//  is currently happening, or has already happened
	ActivityObject interface{}
	// ObjectOrLink describes an object of any kind.
	ObjectOrLink interface {
		GetID() ObjectID
		GetType() ActivityVocabularyType
	}
	// ObjectsArr is a named type for matching an ObjectOrLink slice type to Collection interface
	ObjectsArr []ObjectOrLink
	// LinkOrURI is an interface that GetID and GetLink structs implement, and at the same time
	// they are kept disjointed
	LinkOrURI interface{}
	// ImageOrLink is an interface that Image and GetLink structs implement
	ImageOrLink interface{}
	// MimeType is the type for MIME types
	MimeType string
	// LangRef is the type for a language reference, should be ISO 639-1 language specifier.
	LangRef string
	// NaturalLanguageValue is a mapping for multiple language values
	NaturalLanguageValue map[LangRef]string
)

// IsLink validates if current Activity Pub GetID is a GetLink
func (o Object) IsLink() bool {
	return false
}

// IsObject validates if current Activity Pub GetID is an GetID
func (o Object) IsObject() bool {
	return true
}

// MarshalJSON serializes the NaturalLanguageValue into JSON
func (n NaturalLanguageValue) MarshalJSON() ([]byte, error) {
	if len(n) == 1 {
		for _, v := range n {
			return json.Marshal(v)
		}
	}

	return json.Marshal(map[LangRef]string(n))
}

// Append is syntactic sugar for resizing the NaturalLanguageValue map
//  and appending an element
func (n *NaturalLanguageValue) Append(lang LangRef, value string) error {
	var t NaturalLanguageValue
	if len(*n) == 0 {
		t = make(NaturalLanguageValue, 1)
	} else {
		t = *n
	}
	t[lang] = value
	*n = t

	return nil
}

// UnmarshalJSON tries to load the NaturalLanguage array from the incoming json value
func (l *LangRef) UnmarshalJSON(data []byte) error {
	*l = LangRef(data[1 : len(data)-1])
	return nil
}

// UnmarshalJSON tries to load the NaturalLanguage array from the incoming json value
func (n *NaturalLanguageValue) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		// a quoted string - loading it to c.URL
		if data[len(data)-1] != '"' {
			return fmt.Errorf("invalid string value when unmarshalling %T value", n)
		}
		n.Append(LangRef("-"), string(data[1:len(data)-1]))
	}
	if data[0] == '{' {
		// an object - trying to load it to the struct
		if data[len(data)-1] != '}' {
			return fmt.Errorf("invalid object value when unmarshalling %T value", n)
		}
	}
	if data[0] == '[' {
		// an object - trying to load it to the struct
		if data[len(data)-1] != ']' {
			return fmt.Errorf("invalid object value when unmarshalling %T value", n)
		}
	}
	return nil
}

// Describes an object of any kind.
// The Activity Pub GetID type serves as the base type for most of the other kinds of objects defined in the Activity Vocabulary,
//  including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.
type Object struct {
	// Provides the globally unique identifier for an Activity Pub GetID or GetLink.
	ID ObjectID `jsonld:"id,omitempty"`
	//  Identifies the Activity Pub GetID or GetLink type. Multiple values may be specified.
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
	// The content or textual representation of the Activity Pub GetID encoded as a JSON string.
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
	Replies ObjectOrLink `jsonld:"replies,omitempty"`
	// The date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	//  the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// A natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValue `jsonld:"summary,omitempty,collapsible"`
	// One or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub GetID.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	//  while the latter implies associated by reference.
	Tag ObjectOrLink `jsonld:"tag,omitempty"`
	// The date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// Identifies one or more links to representations of the object
	URL LinkOrURI `jsonld:"url,omitempty"`
	// Identifies an entity considered to be part of the public primary audience of an Activity Pub GetID
	To ObjectsArr `jsonld:"to,omitempty"`
	// Identifies an Activity Pub GetID that is part of the private primary audience of this Activity Pub GetID.
	Bto ObjectsArr `jsonld:"bto,omitempty"`
	// Identifies an Activity Pub GetID that is part of the public secondary audience of this Activity Pub GetID.
	CC ObjectsArr `jsonld:"cc,omitempty"`
	// Identifies one or more Objects that are part of the private secondary audience of this Activity Pub GetID.
	BCC ObjectsArr `jsonld:"bcc,omitempty"`
	// When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	//  the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	//  section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
}

// ContentType represents the content type for a Source object
type ContentType string

// Source is intended to convey some sort of source from which the content markup was derived,
// as a form of provenance, or to support future editing by clients.
type Source struct {
	Content   ContentType
	MediaType string
}

// ValidGenericType validates the type against the valid generic object types
func ValidGenericType(typ ActivityVocabularyType) bool {
	for _, v := range validGenericObjectTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ValidObjectType validates the type against the valid object types
func ValidObjectType(typ ActivityVocabularyType) bool {
	for _, v := range validObjectTypes {
		if v == typ {
			return true
		}
	}
	return ValidActivityType(typ) || ValidActorType(typ) || ValidCollectionType(typ) || ValidGenericType(typ)
}

// ObjectNew initializes a new GetID
func ObjectNew(id ObjectID, typ ActivityVocabularyType) *Object {
	if !(ValidObjectType(typ)) {
		typ = ObjectType
	}
	o := Object{ID: id, Type: typ}
	o.Name = make(NaturalLanguageValue)
	o.Content = make(NaturalLanguageValue)
	o.Summary = make(NaturalLanguageValue)
	return &o
}

// GetID returns the GetID corresponding to the GetID object
func (o Object) GetID() ObjectID {
	return o.ID
}

// GetLink returns the GetLink corresponding to the GetID object
func (o Object) GetType() ActivityVocabularyType {
	return o.Type
}

// Append facilitates adding elements to ObjectOrLink arrays
// and ensures ObjectsArr implements the Collection interface
func (c *ObjectsArr) Append(o ObjectOrLink) error {
	oldLen := len(*c)
	d := make(ObjectsArr, oldLen+1)
	for k, it := range *c {
		d[k] = it
	}
	d[oldLen] = o
	*c = d
	return nil
}

// recipientsDeduplication normalizes the received arguments lists
func recipientsDeduplication(recArgs ...*ObjectsArr) error {
	recIds := make([]ObjectID, 0)

	for _, recList := range recArgs {
		if recList == nil {
			continue
		}

		toRemove := make([]int, 0)
		for i, rec := range *recList {
			save := true
			for _, id := range recIds {
				if rec.GetID() == id {
					// mark the element for removal
					toRemove = append(toRemove, i)
					save = false
				}
			}
			if save {
				recIds = append(recIds, rec.GetID())
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(toRemove)))
		for _, idx := range toRemove {
			*recList = append((*recList)[:idx], (*recList)[idx+1:]...)
		}
	}
	return nil
}

func (i *ObjectID) MarshalJSON(data []byte) error {
	*i = ObjectID(data)
	return nil
}

func (c *ContentType) MarshalJSON(data []byte) error {
	*c = ContentType(data)
	return nil
}

func (o *Object) MarshalJSON(data []byte) error {
	val, _, _, _ := jsonparser.Get(data, "ID")
	o.ID.MarshalJSON(val)
	return nil
}
