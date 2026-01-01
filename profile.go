package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/valyala/fastjson"
)

// Profile a Profile is a content object that describes another Object,
// typically used to describe CanReceiveActivities Type objects.
// The describes property is used to reference the object being described by the profile.
type Profile struct {
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
	// Describes On a Profile object, the describes property identifies the object described by the Profile.
	Describes Item `jsonld:"describes,omitempty"`
}

// IsLink returns false for Profile objects
func (p Profile) IsLink() bool {
	return false
}

// IsObject returns true for Profile objects
func (p Profile) IsObject() bool {
	return true
}

// IsCollection returns false for Profile objects
func (p Profile) IsCollection() bool {
	return false
}

// GetLink returns the IRI corresponding to the current Profile object
func (p Profile) GetLink() IRI {
	return IRI(p.ID)
}

// GetType returns the type of the current Profile
func (p Profile) GetType() ActivityVocabularyType {
	return p.Type
}

// GetID returns the ID corresponding to the current Profile
func (p Profile) GetID() ID {
	return p.ID
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (p *Profile) UnmarshalJSON(data []byte) error {
	par := fastjson.Parser{}
	val, err := par.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadProfile(val, p)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (p Profile) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false
	JSONWrite(&b, '{')

	OnObject(p, func(o *Object) error {
		return nil
	})

	if p.Describes != nil {
		notEmpty = JSONWriteItemProp(&b, "describes", p.Describes) || notEmpty
	}

	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (p *Profile) UnmarshalBinary(data []byte) error {
	return p.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (p Profile) MarshalBinary() ([]byte, error) {
	return p.GobEncode()
}

// GobEncode
func (p Profile) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapProfileProperties(mm, p)
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
func (p *Profile) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapProfileProperties(mm, p)
}

// Recipients performs recipient de-duplication on the Profile object's To, Bto, CC and BCC properties
func (p *Profile) Recipients() ItemCollection {
	aud := p.Audience
	return ItemCollectionDeduplication(&p.To, &p.CC, &p.Bto, &p.BCC, &aud)
}

// Clean removes Bto and BCC properties
func (p *Profile) Clean() {
	_ = OnObject(p, func(o *Object) error {
		o.Clean()
		return nil
	})
}

func (p Profile) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { }", p, p.Type)
	}
}

// ToProfile tries to convert the "it" Item to a Profile object
func ToProfile(it Item) (*Profile, error) {
	switch i := it.(type) {
	case *Profile:
		return i, nil
	case Profile:
		return &i, nil
	default:
		return reflectItemToType[Profile](it)
	}
}

type withProfileFn func(*Profile) error

// OnProfile calls function fn on it Item if it can be asserted to type *Profile
//
// This function should be called if trying to access the Profile specific properties
// like "describes".
// For the other properties OnObject should be used instead.
func OnProfile(it Item, fn func(*Profile) error) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return callOnItemCollection(it, OnProfile, fn)
	}
	ob, err := ToProfile(it)
	if err != nil {
		return err
	}
	return fn(ob)
}
