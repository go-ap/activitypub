package activitypub

import (
	"errors"
	"time"
	"unsafe"
)

// OrderedCollection is a subtype of Collection in which members of the logical
// collection are assumed to always be strictly ordered.
type OrderedCollection struct {
	// ID provides the globally unique identifier for anActivity Pub Object or Link.
	ID ObjectID `jsonld:"id,omitempty"`
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
	OrderedItems ItemCollection `jsonld:"orderedItems,omitempty"`
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
func (o OrderedCollection) GetID() ObjectID {
	return o.ID
}

// GetLink returns the IRI corresponding to the OrderedCollection object
func (o OrderedCollection) GetLink() IRI {
	return IRI(o.ID)
}

// IsObject returns true for am OrderedCollection object
func (o OrderedCollection) IsObject() bool {
	return true
}

// Collection returns the underlying Collection type
func (o *OrderedCollection) Collection() ItemCollection {
	return o.OrderedItems
}

// IsCollection returns true for OrderedCollection objects
func (o OrderedCollection) IsCollection() bool {
	return true
}

// Contains verifies if OrderedCollection array contains the received one
func (o OrderedCollection) Contains(r IRI) bool {
	if len(o.OrderedItems) == 0 {
		return false
	}
	for _, iri := range o.OrderedItems {
		if r.Equals(iri.GetLink(), false) {
			return true
		}
	}
	return false
}

// Count returns the maximum between the length of Items in collection and its TotalItems property
func (o *OrderedCollection) Count() uint {
	if o.TotalItems > 0 {
		return o.TotalItems
	}
	return uint(len(o.OrderedItems))
}

// Append adds an element to an OrderedCollection
func (o *OrderedCollection) Append(ob Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	return nil
}

// UnmarshalJSON
func (o *OrderedCollection) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	o.ID = JSONGetObjectID(data)
	o.Type = JSONGetType(data)
	o.Name = JSONGetNaturalLanguageField(data, "name")
	o.Content = JSONGetNaturalLanguageField(data, "content")
	o.Summary = JSONGetNaturalLanguageField(data, "summary")
	o.Context = JSONGetItem(data, "context")
	o.URL = JSONGetURIItem(data, "url")
	o.MediaType = MimeType(JSONGetString(data, "mediaType"))
	o.Generator = JSONGetItem(data, "generator")
	o.AttributedTo = JSONGetItem(data, "attributedTo")
	o.Attachment = JSONGetItem(data, "attachment")
	o.Location = JSONGetItem(data, "location")
	o.Published = JSONGetTime(data, "published")
	o.StartTime = JSONGetTime(data, "startTime")
	o.EndTime = JSONGetTime(data, "endTime")
	o.Duration = JSONGetDuration(data, "duration")
	o.Icon = JSONGetItem(data, "icon")
	o.Preview = JSONGetItem(data, "preview")
	o.Image = JSONGetItem(data, "image")
	o.Updated = JSONGetTime(data, "updated")
	inReplyTo := JSONGetItems(data, "inReplyTo")
	if len(inReplyTo) > 0 {
		o.InReplyTo = inReplyTo
	}
	to := JSONGetItems(data, "to")
	if len(to) > 0 {
		o.To = to
	}
	audience := JSONGetItems(data, "audience")
	if len(audience) > 0 {
		o.Audience = audience
	}
	bto := JSONGetItems(data, "bto")
	if len(bto) > 0 {
		o.Bto = bto
	}
	cc := JSONGetItems(data, "cc")
	if len(cc) > 0 {
		o.CC = cc
	}
	bcc := JSONGetItems(data, "bcc")
	if len(bcc) > 0 {
		o.BCC = bcc
	}
	replies := JSONGetItem(data, "replies")
	if replies != nil {
		o.Replies = replies
	}
	tag := JSONGetItems(data, "tag")
	if len(tag) > 0 {
		o.Tag = tag
	}

	o.TotalItems = uint(JSONGetInt(data, "totalItems"))
	o.OrderedItems = JSONGetItems(data, "orderedItems")

	o.Current = JSONGetItem(data, "current")
	o.First = JSONGetItem(data, "first")
	o.Last = JSONGetItem(data, "last")

	return nil
}

// OrderedCollectionPageNew initializes a new OrderedCollectionPage
func OrderedCollectionPageNew(parent CollectionInterface) *OrderedCollectionPage {
	p := OrderedCollectionPage{
		PartOf: parent.GetLink(),
	}
	if pc, ok := parent.(*OrderedCollection); ok {
		copyOrderedCollectionToPage(pc, &p)
	}
	p.Type = OrderedCollectionPageType
	return &p
}

// ToOrderedCollection
func ToOrderedCollection(it Item) (*OrderedCollection, error) {
	switch i := it.(type) {
	case *OrderedCollection:
		return i, nil
	case OrderedCollection:
		return &i, nil
	case *OrderedCollectionPage:
		return (*OrderedCollection)(unsafe.Pointer(i)), nil
	case OrderedCollectionPage:
		return (*OrderedCollection)(unsafe.Pointer(&i)), nil
	}
	return nil, errors.New("unable to convert to ordered collection")
}

func copyOrderedCollectionToPage(c *OrderedCollection, p *OrderedCollectionPage) error {
	p.ID = c.ID
	return nil
}
