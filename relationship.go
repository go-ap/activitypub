package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

// Relationship describes a relationship between two individuals.
// The subject and object properties are used to identify the connected individuals.
// See 5.2 Representing Relationships Between Entities for additional information.
//
//	5.2: The relationship property specifies the kind of relationship that exists between the two individuals identified
//	by the subject and object properties. Used together, these three properties form what is commonly known
//	as a "reified statement" where subject identifies the subject, relationship identifies the predicate,
//	and object identifies the object.
type Relationship struct {
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
	// Subject property identifies one of the connected individuals.
	// For instance, for a Relationship object describing "John is related to Sally", subject would refer to John.
	Subject Item `jsonld:"subject,omitempty"`
	// Object  property identifies one of the connected individuals.
	// For instance, for a Relationship object describing "John is related to Sally", object would refer to Sally.
	Object Item `jsonld:"object,omitempty"`
	// Relationship property identifies the kind of relationship that exists between subject and object.
	Relationship Item `jsonld:"relationship,omitempty"`
}

// IsLink returns false for Relationship objects
func (r Relationship) IsLink() bool {
	return false
}

// IsObject returns true for Relationship objects
func (r Relationship) IsObject() bool {
	return true
}

// IsCollection returns false for Relationship objects
func (r Relationship) IsCollection() bool {
	return false
}

// GetLink returns the IRI corresponding to the current Relationship object
func (r Relationship) GetLink() IRI {
	return IRI(r.ID)
}

// GetType returns the type of the current Relationship
func (r Relationship) GetType() ActivityVocabularyType {
	return r.Type
}

// GetID returns the ID corresponding to the current Relationship
func (r Relationship) GetID() ID {
	return r.ID
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (r *Relationship) UnmarshalJSON(data []byte) error {
	par := fastjson.Parser{}
	val, err := par.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadRelationship(val, r)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (r Relationship) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false
	JSONWrite(&b, '{')

	OnObject(r, func(o *Object) error {
		notEmpty = JSONWriteObjectValue(&b, *o)
		return nil
	})

	if r.Subject != nil {
		notEmpty = JSONWriteItemProp(&b, "subject", r.Subject) || notEmpty
	}
	if r.Object != nil {
		notEmpty = JSONWriteItemProp(&b, "object", r.Object) || notEmpty
	}
	if r.Relationship != nil {
		notEmpty = JSONWriteItemProp(&b, "relationship", r.Relationship) || notEmpty
	}

	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (r *Relationship) UnmarshalBinary(data []byte) error {
	return r.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (r Relationship) MarshalBinary() ([]byte, error) {
	return r.GobEncode()
}

// GobEncode
func (r Relationship) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapRelationshipProperties(mm, r)
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
func (r *Relationship) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapRelationshipProperties(mm, r)
}

// Recipients performs recipient de-duplication on the Relationship object's To, Bto, CC and BCC properties
func (r *Relationship) Recipients() ItemCollection {
	aud := r.Audience
	return ItemCollectionDeduplication(&r.To, &r.CC, &r.Bto, &r.BCC, &aud)
}

// Clean removes Bto and BCC properties
func (r *Relationship) Clean() {
	_ = OnObject(r, func(o *Object) error {
		o.Clean()
		return nil
	})
}

func (r Relationship) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { }", r, r.Type)
	}
}

// ToRelationship tries to convert the "it" Item to a Relationship object.
func ToRelationship(it Item) (*Relationship, error) {
	switch i := it.(type) {
	case *Relationship:
		return i, nil
	case Relationship:
		return &i, nil
	case *Object:
		return (*Relationship)(unsafe.Pointer(i)), nil
	case Object:
		return (*Relationship)(unsafe.Pointer(&i)), nil
	default:
		return reflectItemToType[Relationship](it)
	}
}

type withRelationshipFn func(*Relationship) error

// OnRelationship calls function fn on it Item if it can be asserted to type *Relationship
//
// This function should be called if trying to access the Relationship specific properties
// like "subject", "object", or "relationship".
// For the other properties OnObject should be used instead.
func OnRelationship(it Item, fn withRelationshipFn) error {
	if it == nil {
		return nil
	}
	ob, err := ToRelationship(it)
	if err != nil {
		return err
	}
	return fn(ob)
}
