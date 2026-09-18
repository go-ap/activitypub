package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

type IntransitiveActivities interface {
	IntransitiveActivity | Question
}

// IntransitiveActivity Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions.
// The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
	// ID provides the globally unique identifier for the object.
	ID ID `jsonld:"id,omitempty"`
	// Type identifies the Object type. Multiple values may be specified.
	Type Typer `jsonld:"type,omitempty"`
	// Name represents a simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValues `jsonld:"name,omitempty,collapsible"`
	// Attachment identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment Item `jsonld:"attachment,omitempty"`
	// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo Item `jsonld:"attributedTo,omitempty"`
	// Audience identifies one or more entities that represent the total population of entities
	// for which the object can be considered to be relevant.
	Audience ItemCollection `jsonld:"audience,omitempty"`
	// Content represents a textual representation of the Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The MediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValues `jsonld:"content,omitempty,collapsible"`
	// Context identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	// common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
	// MediaType identifies the MIME media type of the value of the Content property.
	// If not specified, the content property is assumed to contain text/html content.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// EndTime represents the date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the EndTime property specifies the moment
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
	// StartTime represents the date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the StartTime property specifies
	// the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// Summary represents a natural language summarization of the object encoded as HTML.
	// Multiple language tagged summaries MAY be provided.
	Summary NaturalLanguageValues `jsonld:"summary,omitempty,collapsible"`
	// Tag identifies one or more "tags" that have been associated with an objects. A tag can be any kind of Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	// while the latter implies associated by reference.
	Tag Item `jsonld:"tag,omitempty"`
	// Updated represents the date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// URL identifies one or more links to representations of the object
	URL Item `jsonld:"url,omitempty"`
	// To identifies an entity considered to be part of the public primary audience of an Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Bto identifies an Object that is part of the private primary audience of this Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// CC identifies an Object that is part of the public secondary audience of this Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// BCC identifies one or more Objects that are part of the private secondary audience of this Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// Duration indicates the object's approximate duration when the Object describes a time-bound resource,
	// such as an Audio or Video, a meeting, etc,
	// The value must be expressed as an xsd:duration as defined by [xmlschema11-2],
	// section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// Likes identifies the collection containing a list of all Like activities with this object as the object property, added as a side effect.
	// The likes collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Likes Item `jsonld:"likes,omitempty"`
	// Shares identifies the collection containing a list of all Announce activities with this object as the object property, added as a side effect.
	// The shares collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Shares Item `jsonld:"shares,omitempty"`
	// Source property is intended to convey some sort of source from which the content markup was derived,
	// as a form of provenance, or to support future editing by clients.
	// In general, clients do the conversion from source to content, not the other way around.
	Source Source `jsonld:"source,omitempty"`
	// Actor describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor Item `jsonld:"actor,omitempty"`
	// Target describes the indirect object, or target, of the activity.
	// The precise meaning of the target is largely dependent on the type of action being described
	// but will often be the object of the English preposition "to".
	// For instance, in the activity "John added a movie to his wishlist",
	// the target of the activity is John's wishlist. An activity can have more than one target.
	Target Item `jsonld:"target,omitempty"`
	// Result describes the result of the Activity. For instance, if a particular action results in the creation
	// of a new resource, the result property can be used to describe that new resource.
	Result Item `jsonld:"result,omitempty"`
	// Origin describes an indirect object of the activity from which the activity is directed.
	// The precise meaning of the origin is the object of the English preposition "from".
	// For instance, in the activity "John moved an item to List B from List A", the origin of the activity is "List A".
	Origin Item `jsonld:"origin,omitempty"`
	// Instrument identifies one or more objects used (or to be used) in the completion of an Activity.
	Instrument Item `jsonld:"instrument,omitempty"`
}

type (
	// Arrive is an IntransitiveActivity that indicates that the actor has arrived at the location.
	// The origin can be used to identify the context from which the actor originated.
	// The target typically has no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-arrive
	Arrive = IntransitiveActivity

	// Travel indicates that the actor is traveling to target from origin.
	// Travel is an IntransitiveObject whose actor specifies the direct object.
	// If the target or origin are not specified, either can be determined by context.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-travel
	Travel = IntransitiveActivity
)

// Recipients performs recipient de-duplication on the IntransitiveActivity's To, Bto, CC and BCC properties
func (i *IntransitiveActivity) Recipients() ItemCollection {
	aud := i.Audience
	return ItemCollectionDeduplication(&i.To, &i.CC, &i.Bto, &i.BCC, &aud)
}

func cleanIntransitiveActivityProperties(aa *IntransitiveActivity) {
	_ = OnObject(aa, func(ob *Object) error {
		cleanObjectProperties(ob)
		return nil
	})
	aa.Actor = CleanRecipients(aa.Actor)
	aa.Target = CleanRecipients(aa.Target)
	aa.Instrument = CleanRecipients(aa.Instrument)
	aa.Origin = CleanRecipients(aa.Origin)
}

// Clean removes Bto and BCC properties
func (i *IntransitiveActivity) Clean() Item {
	aa := *i
	cleanIntransitiveActivityProperties(&aa)
	return &aa
}

// GetType returns the ActivityVocabulary type of the current Intransitive Activity
func (i IntransitiveActivity) GetType() Typer {
	return i.Type
}

// GetID returns the ID corresponding to the IntransitiveActivity object
func (i IntransitiveActivity) GetID() ID {
	return i.ID
}

// GetLink returns the IRI corresponding to the IntransitiveActivity object
func (i IntransitiveActivity) GetLink() IRI {
	return IRI(i.ID)
}

// Match returns whether the receiver matches the ActivityVocabularyType arguments.
func (i IntransitiveActivity) Match(tt ...ActivityVocabularyType) bool {
	return ActivityVocabularyTypes(tt).Match(i.Type)
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (i *IntransitiveActivity) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadIntransitiveActivity(val, i)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (i IntransitiveActivity) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	JSONWrite(&b, '{')

	if !JSONWriteIntransitiveActivityValue(&b, i) {
		return nil, nil
	}
	JSONWrite(&b, '}')
	return b.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (i *IntransitiveActivity) UnmarshalBinary(data []byte) error {
	return i.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (i IntransitiveActivity) MarshalBinary() ([]byte, error) {
	return i.GobEncode()
}

func (i IntransitiveActivity) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapIntransitiveActivityProperties(mm, &i)
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

func (i *IntransitiveActivity) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapIntransitiveActivityProperties(mm, i)
}

// ToIntransitiveActivity tries to convert it Item to an IntransitiveActivity object
func ToIntransitiveActivity(it LinkOrIRI) (*IntransitiveActivity, error) {
	switch i := it.(type) {
	case *IntransitiveActivity:
		return i, nil
	case IntransitiveActivity:
		return &i, nil
	case *Question:
		return (*IntransitiveActivity)(unsafe.Pointer(i)), nil
	case Question:
		return (*IntransitiveActivity)(unsafe.Pointer(&i)), nil
	case *Activity:
		return (*IntransitiveActivity)(unsafe.Pointer(i)), nil
	case Activity:
		return (*IntransitiveActivity)(unsafe.Pointer(&i)), nil
	default:
		return reflectItemToType[IntransitiveActivity](it)
	}
}

// Equals verifies if our receiver IntransitiveActivity is equals with the "with" Item
func (i IntransitiveActivity) Equals(with Item) bool {
	withActivity, err := ToIntransitiveActivity(with)
	if err != nil {
		return false
	}
	return i.equal(*withActivity)
}

// equal verifies if our receiver IntransitiveActivity is equals with the "with" IntransitiveActivity
func (i IntransitiveActivity) equal(with IntransitiveActivity) bool {
	result := true

	_ = OnObject(i, func(oa *Object) error {
		result = oa.Equals(with)
		return nil
	})
	if !result {
		return false
	}
	if !ItemsEqual(i.Actor, with.Actor) {
		return false
	}
	if !ItemsEqual(i.Target, with.Target) {
		return false
	}
	if !ItemsEqual(i.Result, with.Result) {
		return false
	}
	if !ItemsEqual(i.Origin, with.Origin) {
		return false
	}
	if !ItemsEqual(i.Instrument, with.Instrument) {
		return false
	}
	return true
}

func (i IntransitiveActivity) Format(s fmt.State, verb rune) {
	typ := i.Type
	switch verb {
	case 's':
		iri := i.ID
		if iri != "" {
			s.Write([]byte(iri))
		} else {
			_, _ = fmt.Fprintf(s, "%T[%v]", i, typ)
		}
	case 'v':
		n := 0
		if typ != nil {
			_, _ = fmt.Fprintf(s, "%T[%v] { ", i, typ)
			_ = fmtIntransitiveActivityProps(s, &n)(&i)
			_, _ = io.WriteString(s, " }")
		} else {
			_, _ = fmt.Fprintf(s, "%T { ", i)
			_ = fmtIntransitiveActivityProps(s, &n)(&i)
			_, _ = io.WriteString(s, " }")
		}
	}
}

func fmtIntransitiveActivityProps(w io.Writer, n *int) func(*IntransitiveActivity) error {
	return func(ia *IntransitiveActivity) error {
		_ = OnObject(ia, fmtObjectProps(w, n))
		comma := func() {
			if *n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}

		if !IsNil(ia.Actor) {
			comma()
			*n, _ = fmt.Fprintf(w, "actor: %s", ia.Actor)
		}
		if !IsNil(ia.Target) {
			comma()
			*n, _ = fmt.Fprintf(w, "target: %s", ia.Target)
		}
		if !IsNil(ia.Result) {
			comma()
			*n, _ = fmt.Fprintf(w, "result: %s", ia.Result)
		}
		if !IsNil(ia.Origin) {
			comma()
			*n, _ = fmt.Fprintf(w, "origin: %s", ia.Origin)
		}
		if !IsNil(ia.Instrument) {
			comma()
			*n, _ = fmt.Fprintf(w, "instrument: %s", ia.Instrument)
		}
		return nil
	}
}

// WithIntransitiveActivityFn represents a function type that can be used as a parameter for OnIntransitiveActivity helper function
type WithIntransitiveActivityFn func(*IntransitiveActivity) error

// OnIntransitiveActivity calls function fn on it Item if it can be asserted
// to type *IntransitiveActivity
//
// This function should be called if trying to access the IntransitiveActivity
// specific properties like "actor", for the other properties OnObject
// should be used instead.
func OnIntransitiveActivity(it LinkOrIRI, fn WithIntransitiveActivityFn) error {
	if IsNil(it) {
		return nil
	}
	actFn := func(it LinkOrIRI) error {
		act, err := ToIntransitiveActivity(it)
		if err != nil {
			return err
		}
		return fn(act)
	}

	if !IsItemCollection(it) {
		return actFn(it)
	}
	return OnItemCollection(it, func(col *ItemCollection) error {
		for _, ob := range *col {
			if err := actFn(ob); err != nil {
				return err
			}
		}
		return nil
	})
}

func notEmptyIntransitiveActivity(i *IntransitiveActivity) bool {
	notEmpty := i.Actor != nil ||
		i.Target != nil ||
		i.Result != nil ||
		i.Origin != nil ||
		i.Instrument != nil
	if notEmpty {
		return true
	}
	_ = OnObject(i, func(ob *Object) error {
		notEmpty = notEmptyObject(ob)
		return nil
	})
	return notEmpty
}
