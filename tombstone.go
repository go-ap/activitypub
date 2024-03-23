package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

// Tombstone a Tombstone represents a content object that has been deleted.
// It can be used in Collections to signify that there used to be an object at this position,
// but it has been deleted.
type Tombstone struct {
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
	// FormerType On a Tombstone object, the formerType property identifies the type of the object that was deleted.
	FormerType ActivityVocabularyType `jsonld:"formerType,omitempty"`
	// Deleted On a Tombstone object, the deleted property is a timestamp for when the object was deleted.
	Deleted time.Time `jsonld:"deleted,omitempty"`
}

// IsLink returns false for Tombstone objects
func (t Tombstone) IsLink() bool {
	return false
}

// IsObject returns true for Tombstone objects
func (t Tombstone) IsObject() bool {
	return true
}

// IsCollection returns false for Tombstone objects
func (t Tombstone) IsCollection() bool {
	return false
}

// GetLink returns the IRI corresponding to the current Tombstone object
func (t Tombstone) GetLink() IRI {
	return IRI(t.ID)
}

// GetType returns the type of the current Tombstone
func (t Tombstone) GetType() ActivityVocabularyType {
	return t.Type
}

// GetID returns the ID corresponding to the current Tombstone
func (t Tombstone) GetID() ID {
	return t.ID
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (t *Tombstone) UnmarshalJSON(data []byte) error {
	par := fastjson.Parser{}
	val, err := par.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadTombstone(val, t)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (t Tombstone) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false
	JSONWrite(&b, '{')

	OnObject(t, func(o *Object) error {
		notEmpty = JSONWriteObjectValue(&b, *o)
		return nil
	})
	if len(t.FormerType) > 0 {
		if v, err := t.FormerType.MarshalJSON(); err == nil && len(v) > 0 {
			notEmpty = JSONWriteProp(&b, "formerType", v) || notEmpty
		}
	}
	if !t.Deleted.IsZero() {
		notEmpty = JSONWriteTimeProp(&b, "deleted", t.Deleted) || notEmpty
	}
	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *Tombstone) UnmarshalBinary(data []byte) error {
	return t.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t Tombstone) MarshalBinary() ([]byte, error) {
	return t.GobEncode()
}

// GobEncode
func (t Tombstone) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapTombstoneProperties(mm, t)
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
func (t *Tombstone) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapTombstoneProperties(mm, t)
}

// Recipients performs recipient de-duplication on the Tombstone object's To, Bto, CC and BCC properties
func (t *Tombstone) Recipients() ItemCollection {
	aud := t.Audience
	return ItemCollectionDeduplication(&t.To, &t.CC, &t.Bto, &t.BCC, &aud)
}

// Clean removes Bto and BCC properties
func (t *Tombstone) Clean() {
	_ = OnObject(t, func(o *Object) error {
		o.Clean()
		return nil
	})
}

func (t Tombstone) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { formerType: %q }", t, t.Type, t.FormerType)
	}
}

// ToTombstone
func ToTombstone(it Item) (*Tombstone, error) {
	switch i := it.(type) {
	case *Tombstone:
		return i, nil
	case Tombstone:
		return &i, nil
	case *Object:
		return (*Tombstone)(unsafe.Pointer(i)), nil
	case Object:
		return (*Tombstone)(unsafe.Pointer(&i)), nil
	default:
		// NOTE(marius): this is an ugly way of dealing with the interface conversion error: types from different scopes
		typ := reflect.TypeOf(new(Tombstone))
		if reflect.TypeOf(it).ConvertibleTo(typ) {
			if i, ok := reflect.ValueOf(it).Convert(typ).Interface().(*Tombstone); ok {
				return i, nil
			}
		}
	}
	return nil, ErrorInvalidType[Tombstone](it)
}

type withTombstoneFn func(*Tombstone) error

// OnTombstone calls function fn on it Item if it can be asserted to type *Tombstone
//
// This function should be called if trying to access the Tombstone specific properties
// like "formerType" or "deleted".
// For the other properties OnObject should be used instead.
func OnTombstone(it Item, fn withTombstoneFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if err := OnTombstone(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	ob, err := ToTombstone(it)
	if err != nil {
		return err
	}
	return fn(ob)
}
