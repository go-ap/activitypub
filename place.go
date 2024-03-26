package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/valyala/fastjson"
)

// Place represents a logical or physical location. See 5.3 Representing Places for additional information.
type Place struct {
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
	// Accuracy indicates the accuracy of position coordinates on a Place objects.
	// Expressed in properties of percentage. e.g. "94.0" means "94.0% accurate".
	Accuracy float64 `jsonld:"accuracy,omitempty"`
	// Altitude indicates the altitude of a place. The measurement units is indicated using the units property.
	// If units is not specified, the default is assumed to be "m" indicating meters.
	Altitude float64 `jsonld:"altitude,omitempty"`
	// Latitude the latitude of a place
	Latitude float64 `jsonld:"latitude,omitempty"`
	// Longitude the longitude of a place
	Longitude float64 `jsonld:"longitude,omitempty"`
	// Radius the radius from the given latitude and longitude for a Place.
	// The units is expressed by the units property. If units is not specified,
	// the default is assumed to be "m" indicating "meters".
	Radius int64 `jsonld:"radius,omitempty"`
	// Specifies the measurement units for the radius and altitude properties on a Place object.
	// If not specified, the default is assumed to be "m" for "meters".
	// Values "cm" | " feet" | " inches" | " km" | " m" | " miles" | xsd:anyURI
	Units string `jsonld:"units,omitempty"`
}

// IsLink returns false for Place objects
func (p Place) IsLink() bool {
	return false
}

// IsObject returns true for Place objects
func (p Place) IsObject() bool {
	return true
}

// IsCollection returns false for Place objects
func (p Place) IsCollection() bool {
	return false
}

// GetLink returns the IRI corresponding to the current Place object
func (p Place) GetLink() IRI {
	return IRI(p.ID)
}

// GetType returns the type of the current Place
func (p Place) GetType() ActivityVocabularyType {
	return p.Type
}

// GetID returns the ID corresponding to the current Place
func (p Place) GetID() ID {
	return p.ID
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (p *Place) UnmarshalJSON(data []byte) error {
	par := fastjson.Parser{}
	val, err := par.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadPlace(val, p)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (p Place) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false
	JSONWrite(&b, '{')

	OnObject(p, func(o *Object) error {
		notEmpty = JSONWriteObjectValue(&b, *o)
		return nil
	})
	if p.Accuracy > 0 {
		notEmpty = JSONWriteFloatProp(&b, "accuracy", p.Accuracy) || notEmpty
	}
	if p.Altitude > 0 {
		notEmpty = JSONWriteFloatProp(&b, "altitude", p.Altitude) || notEmpty
	}
	if p.Latitude > 0 {
		notEmpty = JSONWriteFloatProp(&b, "latitude", p.Latitude) || notEmpty
	}
	if p.Longitude > 0 {
		notEmpty = JSONWriteFloatProp(&b, "longitude", p.Longitude) || notEmpty
	}
	if p.Radius > 0 {
		notEmpty = JSONWriteIntProp(&b, "radius", p.Radius) || notEmpty
	}
	if len(p.Units) > 0 {
		notEmpty = JSONWriteStringProp(&b, "radius", p.Units) || notEmpty
	}
	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (p *Place) UnmarshalBinary(data []byte) error {
	return p.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (p Place) MarshalBinary() ([]byte, error) {
	return p.GobEncode()
}

// GobEncode
func (p Place) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapPlaceProperties(mm, p)
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
func (p *Place) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapPlaceProperties(mm, p)
}

// Recipients performs recipient de-duplication on the Place object's To, Bto, CC and BCC properties
func (p *Place) Recipients() ItemCollection {
	aud := p.Audience
	return ItemCollectionDeduplication(&p.To, &p.CC, &p.Bto, &p.BCC, &aud)
}

// Clean removes Bto and BCC properties
func (p *Place) Clean() {
	_ = OnObject(p, func(o *Object) error {
		o.Clean()
		return nil
	})
}

func (p Place) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { }", p, p.Type)
	}
}

// ToPlace
func ToPlace(it Item) (*Place, error) {
	switch i := it.(type) {
	case *Place:
		return i, nil
	case Place:
		return &i, nil
	default:
		return reflectItemToType[Place](it)
	}
}

type withPlaceFn func(*Place) error

// OnPlace calls function fn on it Item if it can be asserted to type *Place
//
// This function should be called if trying to access the Place specific properties
// like "accuracy", "altitude", "latitude", "longitude", "radius", or "units".
// For the other properties OnObject should be used instead.
func OnPlace(it Item, fn withPlaceFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if IsLink(it) {
					continue
				}
				if err := OnPlace(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	ob, err := ToPlace(it)
	if err != nil {
		return err
	}
	return fn(ob)
}
