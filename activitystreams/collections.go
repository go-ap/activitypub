package activitystreams

import (
	"time"
)

var validCollectionTypes = [...]ActivityVocabularyType{CollectionType, OrderedCollectionType}

// Page represents a Web Page.
type Page Item

type CollectionInterface interface {
	ObjectOrLink
	Collection() CollectionInterface
	Append(ob Item) error
}

// Collection is a subtype of Activity Pub Object that represents ordered or unordered sets of Activity Pub Object or Link instances.
type Collection struct {
	// Provides the globally unique identifier for an Activity Pub Activity Pub Object or Link.
	ID ObjectID `jsonld:"id,omitempty"`
	//  Identifies the Activity Pub Activity Pub Object or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
	// Identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment Item `jsonld:"attachment,omitempty"`
	// Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo Item `jsonld:"attributedTo,omitempty"`
	// Identifies one or more entities that represent the total population of entities
	//  for which the object can considered to be relevant.
	Audience Item `jsonld:"audience,omitempty"`
	// The content or textual representation of the Activity Pub Activity Pub Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValue `jsonld:"content,omitempty,collapsible"`
	// Identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	//  common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
	// The date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	//  the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Identifies the entity (e.g. an application) that generated the object.
	Generator Item `jsonld:"generator,omitempty"`
	// Indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	//  and should be suitable for presentation at a small size.
	Icon ImageOrLink `jsonld:"icon,omitempty"`
	// Indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image ImageOrLink `jsonld:"image,omitempty"`
	// Indicates one or more entities for which this object is considered a response.
	InReplyTo Item `jsonld:"inReplyTo,omitempty"`
	// Indicates one or more physical or logical locations associated with the object.
	Location Item `jsonld:"location,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// The date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Identifies a Collection containing objects considered to be responses to this object.
	Replies Item `jsonld:"replies,omitempty"`
	// The date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	//  the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// A natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValue `jsonld:"summary,omitempty,collapsible"`
	// One or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	//  while the latter implies associated by reference.
	Tag Item `jsonld:"tag,omitempty"`
	// The date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// Identifies one or more links to representations of the object
	URL LinkOrURI `jsonld:"url,omitempty"`
	// Identifies an entity considered to be part of the public primary audience of an Activity Pub Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Identifies an Activity Pub Activity Pub Object that is part of the private primary audience of this Activity Pub Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// Identifies an Activity Pub Activity Pub Object that is part of the public secondary audience of this Activity Pub Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// Identifies one or more Objects that are part of the private secondary audience of this Activity Pub Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	//  the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	//  section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current ObjectOrLink `jsonld:"current,omitempty"`
	// 	In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items ItemCollection `jsonld:"items,omitempty"`
}

// OrderedCollection is a subtype of Collection in which members of the logical
// collection are assumed to always be strictly ordered.
type OrderedCollection struct {
	// Provides the globally unique identifier for an Activity Pub Activity Pub Object or Link.
	ID ObjectID `jsonld:"id,omitempty"`
	//  Identifies the Activity Pub Activity Pub Object or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
	// Identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment Item `jsonld:"attachment,omitempty"`
	// Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo Item `jsonld:"attributedTo,omitempty"`
	// Identifies one or more entities that represent the total population of entities
	//  for which the object can considered to be relevant.
	Audience Item `jsonld:"audience,omitempty"`
	// The content or textual representation of the Activity Pub Activity Pub Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValue `jsonld:"content,omitempty,collapsible"`
	// Identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	//  common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
	// The date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	//  the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Identifies the entity (e.g. an application) that generated the object.
	Generator Item `jsonld:"generator,omitempty"`
	// Indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	//  and should be suitable for presentation at a small size.
	Icon ImageOrLink `jsonld:"icon,omitempty"`
	// Indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image ImageOrLink `jsonld:"image,omitempty"`
	// Indicates one or more entities for which this object is considered a response.
	InReplyTo Item `jsonld:"inReplyTo,omitempty"`
	// Indicates one or more physical or logical locations associated with the object.
	Location Item `jsonld:"location,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// The date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Identifies a Collection containing objects considered to be responses to this object.
	Replies Item `jsonld:"replies,omitempty"`
	// The date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	//  the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// A natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValue `jsonld:"summary,omitempty,collapsible"`
	// One or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	//  while the latter implies associated by reference.
	Tag Item `jsonld:"tag,omitempty"`
	// The date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// Identifies one or more links to representations of the object
	URL LinkOrURI `jsonld:"url,omitempty"`
	// Identifies an entity considered to be part of the public primary audience of an Activity Pub Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Identifies an Activity Pub Activity Pub Object that is part of the private primary audience of this Activity Pub Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// Identifies an Activity Pub Activity Pub Object that is part of the public secondary audience of this Activity Pub Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// Identifies one or more Objects that are part of the private secondary audience of this Activity Pub Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	//  the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	//  section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current ObjectOrLink `jsonld:"current,omitempty"`
	// 	In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	OrderedItems ItemCollection `jsonld:"orderedItems,omitempty"`
}

// CollectionPage is a Collection that contains a large number of items and when it becomes impractical
// for an implementation to serialize every item contained by a Collection using the items (or orderedItems)
// property alone. In such cases, the items within a Collection can be divided into distinct subsets or "pages".
type CollectionPage struct {
	PartOf *Collection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First Page `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last Page `jsonld:"last,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next Page `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev Page `jsonld:"prev,omitempty"`
}

// OrderedCollectionPage type extends from both CollectionPage and OrderedCollection.
// In addition to the properties inherited from each of those, the OrderedCollectionPage
// may contain an additional startIndex property whose value indicates the relative index position
// of the first item contained by the page within the OrderedCollection to which the page belongs.
type OrderedCollectionPage struct {
	PartOf *OrderedCollection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First Page `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last Page `jsonld:"last,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next Page `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev Page `jsonld:"prev,omitempty"`
	// A non-negative integer value identifying the relative position within the logical view of a strictly ordered collection.
	StartIndex uint `jsonld:"startIndex,omitempty"`
}

// ValidCollectionType validates against the valid collection types
func ValidCollectionType(typ ActivityVocabularyType) bool {
	for _, v := range validCollectionTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// CollectionNew initializes a new Collection
func CollectionNew(id ObjectID) *Collection {
	c := Collection{ID: id, Type: CollectionType}
	c.Name = NaturalLanguageValueNew()
	c.Content = NaturalLanguageValueNew()
	c.Summary = NaturalLanguageValueNew()
	return &c
}

// CollectionNew initializes a new Collection
func OrderedCollectionNew(id ObjectID) *OrderedCollection {
	o := OrderedCollection{ID: id, Type: OrderedCollectionType}
	o.Name = NaturalLanguageValueNew()
	o.Content = NaturalLanguageValueNew()

	return &o
}

// CollectionNew initializes a new Collection
func CollectionPageNew(parent *Collection) *CollectionPage {
	return &CollectionPage{PartOf: parent}
}

// CollectionNew initializes a new Collection
func OrderedCollectionPageNew(parent *OrderedCollection) *OrderedCollectionPage {
	return &OrderedCollectionPage{PartOf: parent}
}

// Append adds an element to an OrderedCollection
func (o *OrderedCollection) Append(ob Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// Append adds an element to an Collection
func (c *Collection) Append(ob Item) error {
	c.Items = append(c.Items, ob)
	c.TotalItems++
	return nil
}

// GetType returns the Collection's type
func (c Collection) GetType() ActivityVocabularyType {
	return c.Type
}

// IsLink returns false for a Collection object
func (c Collection) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Collection object
func (c *Collection) GetID() *ObjectID {
	return &c.ID
}

// IsObject returns true for a Collection object
func (c Collection) IsObject() bool {
	return true
}

// GetType returns the OrderedCollection's type
func (o OrderedCollection) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an OrderedCollection object
func (o OrderedCollection) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the OrderedCollection
func (o OrderedCollection) GetID() *ObjectID {
	return &o.ID
}

// IsObject returns true for am OrderedCollection object
func (o OrderedCollection) IsObject() bool {
	return true
}

// UnmarshalJSON
func (o *OrderedCollection) UnmarshalJSON(data []byte) error {
	o.ID = getAPObjectID(data)
	o.Type = getAPType(data)
	o.Name = getAPNaturalLanguageField(data, "name")
	o.Content = getAPNaturalLanguageField(data, "content")
	u := getURIField(data, "url")
	if len(u) > 0 {
		o.URL = u
	}
	o.TotalItems = uint(getAPInt(data, "totalItems"))
	it := getAPItems(data, "orderedItems")
	if it != nil {
		o.OrderedItems = it
	}
	o.Published = getAPTime(data, "published")
	o.StartTime = getAPTime(data, "startTime")
	o.Updated = getAPTime(data, "updated")

	return nil
}

// UnmarshalJSON
func (c *Collection) UnmarshalJSON(data []byte) error {
	c.ID = getAPObjectID(data)
	c.Type = getAPType(data)
	c.Name = getAPNaturalLanguageField(data, "name")
	c.Content = getAPNaturalLanguageField(data, "content")
	u := getURIField(data, "url")
	if len(u) > 0 {
		c.URL = u
	}
	c.TotalItems = uint(getAPInt(data, "totalItems"))
	it := getAPItems(data, "items")
	if it != nil {
		c.Items = it
	}
	c.Published = getAPTime(data, "published")
	c.StartTime = getAPTime(data, "startTime")
	c.Updated = getAPTime(data, "updated")

	return nil
}

/*
func (c *Collection) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (o *OrderedCollection) MarshalJSON() ([]byte, error) {
	return nil, nil
}
*/

// Collection returns the underlying Collection type
func (c *Collection) Collection() CollectionInterface {
	return c
}

// Collection returns the underlying Collection type
func (o *OrderedCollection) Collection() CollectionInterface {
	return o
}
