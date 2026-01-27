package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

// OrderedCollection is a subtype of Collection in which members of the logical
// collection are assumed to always be strictly ordered.
type OrderedCollection struct {
	// ID provides the globally unique identifier for an Activity Pub Object or Link.
	ID ID `jsonld:"id,omitempty"`
	// Type identifies the Activity Pub Object or Link type. Multiple values may be specified.
	Type ActivityVocabularyTypes `jsonld:"type,omitempty"`
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
	URL Item `jsonld:"url,omitempty"`
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
	// In a paged Collection, indicates the furthest preceding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	OrderedItems ItemCollection `jsonld:"orderedItems,omitempty"`
}

type (
	// InboxStream contains all activities received by the actor.
	// The server SHOULD filter content according to the requester's permission.
	// In general, the owner of an inbox is likely to be able to access all of their inbox contents.
	// Depending on access control, some other content may be public, whereas other content may
	// require authentication for non-owner users, if they can access the inbox at all.
	InboxStream = OrderedCollection

	// LikedCollection is a list of every object from all of the actor's Like activities,
	// added as a side effect. The liked collection MUST be either an OrderedCollection or
	// a Collection and MAY be filtered on privileges of an authenticated user or as
	// appropriate when no authentication is given.
	LikedCollection = OrderedCollection

	// LikesCollection is a list of all Like activities with this object as the object property,
	// added as a side effect. The likes collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when
	// no authentication is given.
	LikesCollection = OrderedCollection

	// OutboxStream contains activities the user has published,
	// subject to the ability of the requestor to retrieve the activity (that is,
	// the contents of the outbox are filtered by the permissions of the person reading it).
	OutboxStream = OrderedCollection

	// SharesCollection is a list of all Announce activities with this object as the object property,
	// added as a side effect. The shares collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication
	// is given.
	SharesCollection = OrderedCollection
)

// GetType returns the OrderedCollection's type
func (o OrderedCollection) GetType() ActivityVocabularyType {
	return o.Type.GetType()
}

// GetTypes returns the OrderedCollection's types
func (o OrderedCollection) GetTypes() ActivityVocabularyTypes {
	return o.Type
}

// IsLink returns false for an OrderedCollection object
func (o OrderedCollection) IsLink() bool {
	return false
}

// GetID returns the ID corresponding to the OrderedCollection
func (o OrderedCollection) GetID() ID {
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
func (o OrderedCollection) Collection() ItemCollection {
	return o.OrderedItems
}

// IsCollection returns true for OrderedCollection objects.
func (o OrderedCollection) IsCollection() bool {
	return true
}

// Contains verifies if OrderedCollection array contains the received item r.
func (o OrderedCollection) Contains(r Item) bool {
	if len(o.OrderedItems) == 0 {
		return false
	}
	for _, it := range o.OrderedItems {
		if ItemsEqual(it, r) {
			return true
		}
	}
	return false
}

// Count returns the maximum between the length of Items in collection and its TotalItems property
func (o *OrderedCollection) Count() uint {
	if o == nil {
		return 0
	}
	return uint(len(o.OrderedItems))
}

// Append adds an element to an the receiver collection object.
func (o *OrderedCollection) Append(it ...Item) error {
	for _, ob := range it {
		if !o.OrderedItems.Contains(ob) {
			o.OrderedItems = append(o.OrderedItems, ob)
			o.TotalItems += 1
		}
	}
	return nil
}

// Remove removes items from an OrderedCollection
func (o *OrderedCollection) Remove(it ...Item) {
	for _, ob := range it {
		if o.OrderedItems.Contains(ob) {
			o.OrderedItems.Remove(ob)
			o.TotalItems -= 1
		}
	}
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (o *OrderedCollection) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadOrderedCollection(val, o)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (o OrderedCollection) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false
	JSONWrite(&b, '{')

	OnObject(o, func(o *Object) error {
		notEmpty = JSONWriteObjectValue(&b, *o)
		return nil
	})
	if o.Current != nil {
		notEmpty = JSONWriteItemProp(&b, "current", o.Current) || notEmpty
	}
	if o.First != nil {
		notEmpty = JSONWriteItemProp(&b, "first", o.First) || notEmpty
	}
	if o.Last != nil {
		notEmpty = JSONWriteItemProp(&b, "last", o.Last) || notEmpty
	}
	notEmpty = JSONWriteIntProp(&b, "totalItems", int64(o.TotalItems)) || notEmpty
	if o.OrderedItems != nil {
		notEmpty = JSONWriteItemCollectionProp(&b, "orderedItems", o.OrderedItems, false) || notEmpty
	}
	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (o *OrderedCollection) UnmarshalBinary(data []byte) error {
	return o.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (o OrderedCollection) MarshalBinary() ([]byte, error) {
	return o.GobEncode()
}

// GobEncode
func (o OrderedCollection) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapOrderedCollectionProperties(mm, o)
	if err != nil {
		return nil, err
	}
	if !hasData {
		return []byte{}, nil
	}
	bb := bytes.Buffer{}
	g := gob.NewEncoder(&bb)
	if err := g.Encode(mm); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

// GobDecode
func (o *OrderedCollection) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapOrderedCollectionProperties(mm, o)
}

// OrderedCollectionPageNew initializes a new OrderedCollectionPage
func OrderedCollectionPageNew(parent CollectionInterface) *OrderedCollectionPage {
	p := OrderedCollectionPage{
		PartOf: parent.GetLink(),
	}
	if pc, ok := parent.(*OrderedCollection); ok {
		copyOrderedCollectionToPage(pc, &p)
	}
	p.Type = OrderedCollectionPageType.ToTypes()
	return &p
}

// ToOrderedCollection
func ToOrderedCollection(it LinkOrIRI) (*OrderedCollection, error) {
	switch i := it.(type) {
	case *OrderedCollection:
		return i, nil
	case OrderedCollection:
		return &i, nil
	case *OrderedCollectionPage:
		return (*OrderedCollection)(unsafe.Pointer(i)), nil
	case OrderedCollectionPage:
		return (*OrderedCollection)(unsafe.Pointer(&i)), nil
	// NOTE(marius): let's try again to convert Collection -> OrderedCollection, as they have the same
	// shape in memory.
	case *Collection:
		return (*OrderedCollection)(unsafe.Pointer(i)), nil
	case Collection:
		return (*OrderedCollection)(unsafe.Pointer(&i)), nil
	case *CollectionPage:
		return (*OrderedCollection)(unsafe.Pointer(i)), nil
	case CollectionPage:
		return (*OrderedCollection)(unsafe.Pointer(&i)), nil
	default:
		return reflectItemToType[OrderedCollection](it)
	}
}

func copyOrderedCollectionToPage(c *OrderedCollection, p *OrderedCollectionPage) error {
	p.Type = OrderedCollectionPageType.ToTypes()
	p.Name = c.Name
	p.Content = c.Content
	p.Summary = c.Summary
	p.Context = c.Context
	p.URL = c.URL
	p.MediaType = c.MediaType
	p.Generator = c.Generator
	p.AttributedTo = c.AttributedTo
	p.Attachment = c.Attachment
	p.Location = c.Location
	p.Published = c.Published
	p.StartTime = c.StartTime
	p.EndTime = c.EndTime
	p.Duration = c.Duration
	p.Icon = c.Icon
	p.Preview = c.Preview
	p.Image = c.Image
	p.Updated = c.Updated
	p.InReplyTo = c.InReplyTo
	p.To = c.To
	p.Audience = c.Audience
	p.Bto = c.Bto
	p.CC = c.CC
	p.BCC = c.BCC
	p.Replies = c.Replies
	p.Tag = c.Tag
	p.TotalItems = c.TotalItems
	p.OrderedItems = c.OrderedItems
	p.Current = c.Current
	p.First = c.First
	p.PartOf = c.GetLink()
	return nil
}

// ItemsMatch
func (o OrderedCollection) ItemsMatch(col ...Item) bool {
	for _, it := range col {
		if match := o.OrderedItems.Contains(it); !match {
			return false
		}
	}
	return true
}

// Equals verifies if our receiver OrderedCollection is equals with the "with" Item
func (o *OrderedCollection) Equals(with Item) bool {
	if IsNil(with) {
		return o == nil
	}
	if !with.IsCollection() {
		return false
	}
	withCollection, err := ToOrderedCollection(with)
	if err != nil {
		return false
	}
	return o.equal(*withCollection)
}

// equal verifies if our receiver OrderedCollection is equals with the "with" OrderedCollection
func (o OrderedCollection) equal(with OrderedCollection) bool {
	result := true
	_ = OnObject(with, func(wo *Object) error {
		if !wo.Equals(o) {
			result = false
			return nil
		}
		return nil
	})
	if with.TotalItems > 0 {
		if with.TotalItems != o.TotalItems {
			result = false
		}
	}
	if with.Current != nil {
		if !ItemsEqual(o.Current, with.Current) {
			result = false
		}
	}
	if with.First != nil {
		if !ItemsEqual(o.First, with.First) {
			result = false
		}
	}
	if with.Last != nil {
		if !ItemsEqual(o.Last, with.Last) {
			result = false
		}
	}
	if with.OrderedItems != nil {
		if !ItemsEqual(o.OrderedItems, with.OrderedItems) {
			result = false
		}
	}
	return result
}

func (o OrderedCollection) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { totalItems: %d }", o, o.GetType(), o.TotalItems)
	}
}
func (o *OrderedCollection) Recipients() ItemCollection {
	aud := o.Audience
	return ItemCollectionDeduplication(&o.To, &o.CC, &o.Bto, &o.BCC, &aud)
}

func (o *OrderedCollection) Clean() {
	_ = OnObject(o, func(o *Object) error {
		o.Clean()
		return nil
	})
}

// OnOrderedCollection calls function fn on it Item if it can be asserted
// to type *OrderedCollection
//
// This function should be called if trying to access the Collection specific
// properties like "totalItems", "orderedItems", etc. For the other properties
// OnObject should be used instead.
func OnOrderedCollection(it Item, fn WithOrderedCollectionFn) error {
	if it == nil {
		return nil
	}
	col, err := ToOrderedCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}
