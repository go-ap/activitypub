package activitypub

import (
	"fmt"
	"github.com/buger/jsonparser"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

const (
	ObjectType                ActivityVocabularyType = "Object"
	LinkType                  ActivityVocabularyType = "Link"
	ActivityType              ActivityVocabularyType = "Activity"
	IntransitiveActivityType  ActivityVocabularyType = "IntransitiveActivity"
	ActorType                 ActivityVocabularyType = "Actor"
	CollectionType            ActivityVocabularyType = "Collection"
	OrderedCollectionType     ActivityVocabularyType = "OrderedCollection"
	CollectionPageType        ActivityVocabularyType = "CollectionPage"
	OrderedCollectionPageType ActivityVocabularyType = "OrderedCollectionPage"

	// ActivityPub Object Types
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

var GenericTypes = ActivityVocabularyTypes{
	ActivityType,
	IntransitiveActivityType,
	ObjectType,
	ActorType,
}

var ObjectTypes = ActivityVocabularyTypes{
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
	// ActivityObject is a subtype of Object that describes some form of action that may happen,
	// is currently happening, or has already happened
	ActivityObject interface {
		// GetID returns the dereferenceable ActivityStreams object id
		GetID() ID
		// GetType returns the ActivityStreams type
		GetType() ActivityVocabularyType
	}
	// LinkOrIRI is an interface that Object and Link structs implement, and at the same time
	// they are kept disjointed
	LinkOrIRI interface {
		// GetLink returns the object id in IRI type
		GetLink() IRI
	}
	// ObjectOrLink describes the requirements of an ActivityStreams object
	ObjectOrLink interface {
		ActivityObject
		LinkOrIRI
		// IsLink shows if current item represents a Link object or an IRI
		IsLink() bool
		// IsObject shows if current item represents an ActivityStreams object
		IsObject() bool
		// IsCollection shows if the current item represents an ItemCollection
		IsCollection() bool
	}
	// Mapper interface allows external objects to implement their own mechanism for loading information
	// from an ActivityStreams vocabulary object
	Mapper interface {
		// FromActivityStreams maps an ActivityStreams object to another struct representation
		FromActivityStreams(Item) error
	}

	// MimeType is the type for representing MIME types in certain ActivityStreams properties
	MimeType string
)

func (a ActivityVocabularyType) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return nil, nil
	}
	b := make([]byte, 0)
	writeStringValue(&b, string(a))
	return b, nil
}

// Object describes an ActivityPub object of any kind.
// It serves as the base type for most of the other kinds of objects defined in the Activity
// Vocabulary, including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.
type Object struct {
	// ID provides the globally unique identifier for anActivity Pub Object or Link.
	ID ID `jsonld:"id,omitempty"`
	// Type identifies the Activity Pub Object or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// Name a simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValues `jsonld:"name,omitempty,collapsible"`
	// Attachment identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment Item `jsonld:"attachment,omitempty"`
	// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo Item `jsonld:"attributedTo,omitempty"`
	// Audience identifies one or more entities that represent the total population of entities
	// for which the object can considered to be relevant.
	Audience ItemCollection `jsonld:"audience,omitempty"`
	// Content or textual representation of the Activity Pub Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValues `jsonld:"content,omitempty,collapsible"`
	// Context identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	// common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
	// MediaType when used on an Object, identifies the MIME media type of the value of the content property.
	// If not specified, the content property is assumed to contain text/html content.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// EndTime the date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	// the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Generator identifies the entity (e.g. an application) that generated the object.
	Generator Item `jsonld:"generator,omitempty"`
	// Icon indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	// and should be suitable for presentation at a small size.
	Icon Item `jsonld:"icon,omitempty"`
	// Image indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image Item `jsonld:"image,omitempty"`
	// InReplyTo indicates one or more entities for which this object is considered a response.
	InReplyTo Item `jsonld:"inReplyTo,omitempty"`
	// Location indicates one or more physical or logical locations associated with the object.
	Location Item `jsonld:"location,omitempty"`
	// Preview identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// Published the date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Replies identifies a Collection containing objects considered to be responses to this object.
	Replies Item `jsonld:"replies,omitempty"`
	// StartTime the date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	// the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// Summary a natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValues `jsonld:"summary,omitempty,collapsible"`
	// Tag one or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	// while the latter implies associated by reference.
	Tag ItemCollection `jsonld:"tag,omitempty"`
	// Updated the date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// URL identifies one or more links to representations of the object
	URL LinkOrIRI `jsonld:"url,omitempty"`
	// To identifies an entity considered to be part of the public primary audience of an Activity Pub Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Bto identifies anActivity Pub Object that is part of the private primary audience of this Activity Pub Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// CC identifies anActivity Pub Object that is part of the public secondary audience of this Activity Pub Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// BCC identifies one or more Objects that are part of the private secondary audience of this Activity Pub Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// Duration when the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	// the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	// section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// This is a list of all Like activities with this object as the object property, added as a side effect.
	// The likes collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Likes Item `jsonld:"likes,omitempty"`
	// This is a list of all Announce activities with this object as the object property, added as a side effect.
	// The shares collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Shares Item `jsonld:"shares,omitempty"`
	// Source property is intended to convey some sort of source from which the content markup was derived,
	// as a form of provenance, or to support future editing by clients.
	// In general, clients do the conversion from source to content, not the other way around.
	Source Source `jsonld:"source,omitempty"`
}

// ObjectNew initializes a new Object
func ObjectNew(typ ActivityVocabularyType) *Object {
	if !(ObjectTypes.Contains(typ)) {
		typ = ObjectType
	}
	o := Object{Type: typ}
	o.Name = NaturalLanguageValuesNew()
	o.Content = NaturalLanguageValuesNew()
	return &o
}

// GetID returns the ID corresponding to the current Object
func (o Object) GetID() ID {
	return o.ID
}

// GetLink returns the IRI corresponding to the current Object
func (o Object) GetLink() IRI {
	return IRI(o.ID)
}

// GetType returns the type of the current Object
func (o Object) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink validates if currentActivity Pub Object is a Link
func (o Object) IsLink() bool {
	return false
}

// IsObject validates if currentActivity Pub Object is an Object
func (o Object) IsObject() bool {
	return true
}

// IsCollection returns false for Object objects
func (o Object) IsCollection() bool {
	return false
}

// UnmarshalJSON
func (o *Object) UnmarshalJSON(data []byte) error {
	return loadObject(data, o)
}

// MarshalJSON
func (o Object) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	write(&b, '{')

	if writeObjectValue(&b, o) {
		write(&b, '}')
		return b, nil
	}
	return nil, nil
}

// Recipients performs recipient de-duplication on the Object's To, Bto, CC and BCC properties
func (o *Object) Recipients() ItemCollection {
	var aud ItemCollection
	return ItemCollectionDeduplication(&aud, &o.To, &o.Bto, &o.CC, &o.BCC, &o.Audience)
}

// Clean removes Bto and BCC properties
func (o *Object) Clean() {
	o.BCC = nil
	o.Bto = nil
}

type (
	// Article represents any kind of multi-paragraph written work.
	Article = Object
	// Audio represents an audio document of any kind.
	Audio = Document
	// Document represents a document of any kind.
	Document = Object
	// Event represents any kind of event.
	Event = Object
	// Image An image document of any kind
	Image = Document
	// Note represents a short written work typically less than a single paragraph in length.
	Note = Object
	// Page represents a Web Page.
	Page = Document
	// Video represents a video document of any kind
	Video = Document
)

// UnmarshalJSON
func (m *MimeType) UnmarshalJSON(data []byte) error {
	*m = MimeType(strings.Trim(string(data), "\""))
	return nil
}

// MarshalJSON
func (m MimeType) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return nil, nil
	}
	b := make([]byte, 0)
	writeStringValue(&b, string(m))
	return b, nil
}

// ToObject returns an Object pointer to the data in the current Item
// It relies on the fact that all the types in this package have a data layout compatible with Object.
func ToObject(it Item) (*Object, error) {
	switch i := it.(type) {
	case *Object:
		return i, nil
	case Object:
		return &i, nil
	case *Place:
		return (*Object)(unsafe.Pointer(i)), nil
	case Place:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Profile:
		return (*Object)(unsafe.Pointer(i)), nil
	case Profile:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Relationship:
		return (*Object)(unsafe.Pointer(i)), nil
	case Relationship:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Tombstone:
		return (*Object)(unsafe.Pointer(i)), nil
	case Tombstone:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Actor:
		return (*Object)(unsafe.Pointer(i)), nil
	case Actor:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Activity:
		return (*Object)(unsafe.Pointer(i)), nil
	case Activity:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *IntransitiveActivity:
		return (*Object)(unsafe.Pointer(i)), nil
	case IntransitiveActivity:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Question:
		return (*Object)(unsafe.Pointer(i)), nil
	case Question:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Collection:
		return (*Object)(unsafe.Pointer(i)), nil
	case Collection:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *CollectionPage:
		return (*Object)(unsafe.Pointer(i)), nil
	case CollectionPage:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *OrderedCollection:
		return (*Object)(unsafe.Pointer(i)), nil
	case OrderedCollection:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *OrderedCollectionPage:
		return (*Object)(unsafe.Pointer(i)), nil
	case OrderedCollectionPage:
		return (*Object)(unsafe.Pointer(&i)), nil
	default:
		// NOTE(marius): this is an ugly way of dealing with the interface conversion error: types from different scopes
		typ := reflect.TypeOf(new(Object))
		if reflect.TypeOf(it).ConvertibleTo(typ) {
			if i, ok := reflect.ValueOf(it).Convert(typ).Interface().(*Object); ok {
				return i, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to convert %q", it.GetType())
}

// FlattenObjectProperties flattens the Object's properties from Object types to IRI
func FlattenObjectProperties(o *Object) *Object {
	o.Replies = Flatten(o.Replies)
	o.AttributedTo = Flatten(o.AttributedTo)
	o.To = FlattenItemCollection(o.To)
	o.Bto = FlattenItemCollection(o.Bto)
	o.CC = FlattenItemCollection(o.CC)
	o.BCC = FlattenItemCollection(o.BCC)
	o.Audience = FlattenItemCollection(o.Audience)
	o.Tag = FlattenItemCollection(o.Tag)
	return o
}

// FlattenProperties flattens the Item's properties from Object types to IRI
func FlattenProperties(it Item) Item {
	if ActivityTypes.Contains(it.GetType()) {
		act, err := ToActivity(it)
		if err == nil {
			return FlattenActivityProperties(act)
		}
	}
	if ActorTypes.Contains(it.GetType()) || ObjectTypes.Contains(it.GetType()) {
		ob, err := ToObject(it)
		if err == nil {
			return FlattenObjectProperties(ob)
		}
	}
	return it
}

// Source is intended to convey some sort of source from which the content markup was derived,
// as a form of provenance, or to support future editing by clients.
type Source struct {
	// Content
	Content NaturalLanguageValues `jsonld:"content"`
	// MediaType
	MediaType MimeType `jsonld:"mediaType"`
}

// GetAPSource
func GetAPSource(data []byte) Source {
	s := Source{}

	if contBytes, _, _, err := jsonparser.Get(data, "source", "content"); err == nil {
		s.Content.UnmarshalJSON(contBytes)
	}
	if mimeBytes, _, _, err := jsonparser.Get(data, "source", "mediaType"); err == nil {
		s.MediaType.UnmarshalJSON(mimeBytes)
	}

	return s
}

// UnmarshalJSON
func (s *Source) UnmarshalJSON(data []byte) error {
	*s = GetAPSource(data)
	return nil
}

// MarshalJSON
func (s Source) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	empty := true
	write(&b, '{')
	if len(s.MediaType) > 0 {
		if v, err := s.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
			empty = !writeProp(&b, "mediaType", v)
		}
	}
	if len(s.Content) > 0 {
		empty = !writeNaturalLanguageProp(&b, "content", s.Content)
	}
	if !empty {
		write(&b, '}')
		return b, nil
	}
	return nil, nil
}

// Equals verifies if our receiver Object is equals with the "with" Object
func (o Object) Equals(with Item) bool {
	if with.IsCollection() {
		return false
	}
	if withID := with.GetID(); len(withID) > 0 && withID != o.ID {
		return false
	}
	if withType := with.GetType(); len(withType) > 0 && withType != o.Type {
		return false
	}
	if with.IsLink() && !with.GetLink().Equals(o.GetLink(), false) {
		return false
	}
	result := true
	OnObject(with, func(w *Object) error {
		if len(w.Name) > 0 {
			if !w.Name.Equals(o.Name) {
				result = false
				return nil
			}
		}
		if len(w.Summary) > 0 {
			if !w.Summary.Equals(o.Summary) {
				result = false
				return nil
			}
		}
		if len(w.Content) > 0 {
			if !w.Content.Equals(o.Content) {
				result = false
				return nil
			}
		}
		if w.Attachment != nil {
			if !ItemsEqual(o.Attachment, w.Attachment) {
				result = false
				return nil
			}
		}
		if w.AttributedTo != nil {
			if !ItemsEqual(o.AttributedTo, w.AttributedTo) {
				result = false
				return nil
			}
		}
		if w.Audience != nil {
			if !ItemsEqual(o.Audience, w.Audience) {
				result = false
				return nil
			}
		}
		if w.Context != nil {
			if !ItemsEqual(o.Context, w.Context) {
				result = false
				return nil
			}
		}
		if w.Generator != nil {
			if !ItemsEqual(o.Generator, w.Generator) {
				result = false
				return nil
			}
		}
		if w.Icon != nil {
			if !ItemsEqual(o.Icon, w.Icon) {
				result = false
				return nil
			}
		}
		if w.Image != nil {
			if !ItemsEqual(o.Image, w.Image) {
				result = false
				return nil
			}
		}
		if w.InReplyTo != nil {
			if !ItemsEqual(o.InReplyTo, w.InReplyTo) {
				result = false
				return nil
			}
		}
		if w.Location != nil {
			if !ItemsEqual(o.Location, w.Location) {
				result = false
				return nil
			}
		}
		if w.Preview != nil {
			if !ItemsEqual(o.Preview, w.Preview) {
				result = false
				return nil
			}
		}
		if w.Replies != nil {
			if !ItemsEqual(o.Replies, w.Replies) {
				result = false
				return nil
			}
		}
		if w.Tag != nil {
			if !ItemsEqual(o.Tag, w.Tag) {
				result = false
				return nil
			}
		}
		if w.URL != nil {
			if o.URL == nil {
				result = false
				return nil
			}
			if !w.URL.GetLink().Equals(o.URL.GetLink(), false) {
				result = false
				return nil
			}
		}
		if w.To != nil {
			if !ItemsEqual(o.To, w.To) {
				result = false
				return nil
			}
		}
		if w.Bto != nil {
			if !ItemsEqual(o.Bto, w.Bto) {
				result = false
				return nil
			}
		}
		if w.CC != nil {
			if !ItemsEqual(o.CC, w.CC) {
				result = false
				return nil
			}
		}
		if w.BCC != nil {
			if !ItemsEqual(o.BCC, w.BCC) {
				result = false
				return nil
			}
		}
		if !w.Published.IsZero() {
			if !w.Published.Equal(o.Published) {
				result = false
				return nil
			}
		}
		if !w.Updated.IsZero() {
			if !w.Updated.Equal(o.Updated) {
				result = false
				return nil
			}
		}
		if !w.StartTime.IsZero() {
			if !w.StartTime.Equal(o.StartTime) {
				result = false
				return nil
			}
		}
		if !w.EndTime.IsZero() {
			if !w.EndTime.Equal(o.EndTime) {
				result = false
				return nil
			}
		}
		if w.Duration != 0 {
			if w.Duration != o.Duration {
				result = false
				return nil
			}
		}
		if w.Likes != nil {
			if !ItemsEqual(o.Likes, w.Likes) {
				result = false
				return nil
			}
		}
		if w.Shares != nil {
			if !ItemsEqual(o.Shares, w.Shares) {
				result = false
				return nil
			}
		}
		return nil
	})
	return result
}
