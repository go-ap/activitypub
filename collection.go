package activitypub

import (
	"errors"
	"reflect"
	"time"
	"unsafe"
)

const CollectionOfItems ActivityVocabularyType = "ItemCollection"

var CollectionTypes = ActivityVocabularyTypes{
	CollectionOfItems,
	CollectionType,
	OrderedCollectionType,
	CollectionPageType,
	OrderedCollectionPageType,
}

type CollectionInterface interface {
	ObjectOrLink
	Collection() ItemCollection
	Append(ob Item) error
	Count() uint
	Contains(Item) bool
}

// Collection is a subtype of Activity Pub Object that represents ordered or unordered sets of Activity Pub Object or Link instances.
type Collection struct {
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
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current ObjectOrLink `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items ItemCollection `jsonld:"items,omitempty"`
}

type (
	// FollowersCollection is a collection of followers
	FollowersCollection = Followers

	// Followers is a Collection type
	Followers = Collection

	// FollowingCollection is a list of everybody that the actor has followed, added as a side effect.
	// The following collection MUST be either an OrderedCollection or a Collection and MAY
	// be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	FollowingCollection = Following

	// Following is a type alias for a simple Collection
	Following = Collection
)

// CollectionNew initializes a new Collection
func CollectionNew(id ID) *Collection {
	c := Collection{ID: id, Type: CollectionType}
	c.Name = NaturalLanguageValuesNew()
	c.Content = NaturalLanguageValuesNew()
	c.Summary = NaturalLanguageValuesNew()
	return &c
}

// OrderedCollectionNew initializes a new OrderedCollection
func OrderedCollectionNew(id ID) *OrderedCollection {
	o := OrderedCollection{ID: id, Type: OrderedCollectionType}
	o.Name = NaturalLanguageValuesNew()
	o.Content = NaturalLanguageValuesNew()

	return &o
}

// GetID returns the ID corresponding to the Collection object
func (c Collection) GetID() ID {
	return c.ID
}

// GetType returns the Collection's type
func (c Collection) GetType() ActivityVocabularyType {
	return c.Type
}

// IsLink returns false for a Collection object
func (c Collection) IsLink() bool {
	return false
}

// IsObject returns true for a Collection object
func (c Collection) IsObject() bool {
	return true
}

// IsCollection returns true for Collection objects
func (c Collection) IsCollection() bool {
	return true
}

// GetLink returns the IRI corresponding to the Collection object
func (c Collection) GetLink() IRI {
	return IRI(c.ID)
}

// Collection returns the Collection's items
func (c Collection) Collection() ItemCollection {
	return c.Items
}

// Append adds an element to a Collection
func (c *Collection) Append(ob Item) error {
	c.Items = append(c.Items, ob)
	return nil
}

// Count returns the maximum between the length of Items in collection and its TotalItems property
func (c *Collection) Count() uint {
	if c.TotalItems > 0 {
		return c.TotalItems
	}
	return uint(len(c.Items))
}

// Contains verifies if Collection array contains the received one
func (c Collection) Contains(r Item) bool {
	if len(c.Items) == 0 {
		return false
	}
	for _, it := range c.Items {
		if ItemsEqual(it, r) {
			return true
		}
	}
	return false
}

// UnmarshalJSON
func (c *Collection) UnmarshalJSON(data []byte) error {
	return loadCollection(data, c)
}

// MarshalJSON
func (c Collection) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false
	write(&b, '{')

	OnObject(c, func(o *Object) error {
		notEmpty = writeObjectJSONValue(&b, *o)
		return nil
	})
	if c.Current != nil {
		notEmpty = writeItemJSONProp(&b, "current", c.Current) || notEmpty
	}
	if c.First != nil {
		notEmpty = writeItemJSONProp(&b, "first", c.First) || notEmpty
	}
	if c.Last != nil {
		notEmpty = writeItemJSONProp(&b, "last", c.Last) || notEmpty
	}
	notEmpty = writeIntJSONProp(&b, "totalItems", int64(c.TotalItems)) || notEmpty
	if c.Items != nil {
		notEmpty = writeItemCollectionJSONProp(&b, "items", c.Items) || notEmpty
	}
	if notEmpty {
		write(&b, '}')
		return b, nil
	}
	return nil, nil
}

// ToCollection
func ToCollection(it Item) (*Collection, error) {
	switch i := it.(type) {
	case *Collection:
		return i, nil
	case Collection:
		return &i, nil
	case *CollectionPage:
		return (*Collection)(unsafe.Pointer(i)), nil
	case CollectionPage:
		return (*Collection)(unsafe.Pointer(&i)), nil
	default:
		// NOTE(marius): this is an ugly way of dealing with the interface conversion error: types from different scopes
		typ := reflect.TypeOf(new(Collection))
		val := reflect.ValueOf(it)
		if val.IsValid() && typ.Elem().Name() == val.Type().Elem().Name() {
			conv := val.Convert(typ)
			if i, ok := conv.Interface().(*Collection); ok {
				return i, nil
			}
		}
	}
	return nil, errors.New("unable to convert to collection")
}

// FollowingNew initializes a new Following
func FollowingNew() *Following {
	id := ID("following")

	i := Following{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()
	i.Summary = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// ItemsMatch
func (c Collection) ItemsMatch(col ...Item) bool {
	for _, it := range col {
		if match := c.Items.Contains(it); !match {
			return false
		}
	}
	return true
}

// Equals
func (c Collection) Equals(with Item) bool {
	if IsNil(with) {
		return false
	}
	if !with.IsCollection() {
		return false
	}
	result := true
	OnCollection(with, func(w *Collection) error {
		OnObject(w, func(wo *Object) error {
			if !wo.Equals(c) {
				result = false
				return nil
			}
			return nil
		})
		if w.TotalItems > 0 {
			if w.TotalItems != c.TotalItems {
				result = false
				return nil
			}
		}
		if w.Current != nil {
			if !ItemsEqual(c.Current, w.Current) {
				result = false
				return nil
			}
		}
		if w.First != nil {
			if !ItemsEqual(c.First, w.First) {
				result = false
				return nil
			}
		}
		if w.Last != nil {
			if !ItemsEqual(c.Last, w.Last) {
				result = false
				return nil
			}
		}
		if w.Items != nil {
			if !c.Items.Equals(w.Items) {
				result = false
				return nil
			}
		}
		return nil
	})
	return result
}
