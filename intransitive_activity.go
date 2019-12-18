package activitypub

import (
	"errors"
	"time"
	"unsafe"
)

// IntransitiveActivity Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions.
// The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
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
	// CanReceiveActivities describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor CanReceiveActivities `jsonld:"actor,omitempty"`
	// Target describes the indirect object, or target, of the activity.
	// The precise meaning of the target is largely dependent on the type of action being described
	// but will often be the object of the English preposition "to".
	// For instance, in the activity "John added a movie to his wishlist",
	// the target of the activity is John's wishlist. An activity can have more than one target.
	Target Item `jsonld:"target,omitempty"`
	// Result describes the result of the activity. For instance, if a particular action results in the creation
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
	Arrive = IntransitiveActivity

	// Travel indicates that the actor is traveling to target from origin.
	// Travel is an IntransitiveObject whose actor specifies the direct object.
	// If the target or origin are not specified, either can be determined by context.
	Travel = IntransitiveActivity
)

// Recipients performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (i *IntransitiveActivity) Recipients() ItemCollection {
	actor := make(ItemCollection, 0)
	actor.Append(i.Actor)
	rec, _ := ItemCollectionDeduplication(&actor, &i.To, &i.Bto, &i.CC, &i.BCC, &i.Audience)
	return rec
}

// Clean removes Bto and BCC properties
func (i *IntransitiveActivity) Clean() {
	i.BCC = nil
	i.Bto = nil
}

// GetType returns the ActivityVocabulary type of the current Intransitive Activity
func (i IntransitiveActivity) GetType() ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for Activity objects
func (i IntransitiveActivity) IsLink() bool {
	return false
}

// GetID returns the ID corresponding to the IntransitiveActivity object
func (i IntransitiveActivity) GetID() ID {
	return i.ID
}

// GetLink returns the IRI corresponding to the IntransitiveActivity object
func (i IntransitiveActivity) GetLink() IRI {
	return IRI(i.ID)
}

// IsObject returns true for IntransitiveActivity objects
func (i IntransitiveActivity) IsObject() bool {
	return true
}

// IsCollection returns false for IntransitiveActivity objects
func (i IntransitiveActivity) IsCollection() bool {
	return false
}

// UnmarshalJSON
func (i *IntransitiveActivity) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	i.ID = JSONGetID(data)
	i.Type = JSONGetType(data)
	i.Name = JSONGetNaturalLanguageField(data, "name")
	i.Content = JSONGetNaturalLanguageField(data, "content")
	i.Summary = JSONGetNaturalLanguageField(data, "summary")
	i.Context = JSONGetItem(data, "context")
	i.URL = JSONGetURIItem(data, "url")
	i.MediaType = MimeType(JSONGetString(data, "mediaType"))
	i.Generator = JSONGetItem(data, "generator")
	i.AttributedTo = JSONGetItem(data, "attributedTo")
	i.Attachment = JSONGetItem(data, "attachment")
	i.Location = JSONGetItem(data, "location")
	i.Published = JSONGetTime(data, "published")
	i.StartTime = JSONGetTime(data, "startTime")
	i.EndTime = JSONGetTime(data, "endTime")
	i.Duration = JSONGetDuration(data, "duration")
	i.Icon = JSONGetItem(data, "icon")
	i.Preview = JSONGetItem(data, "preview")
	i.Image = JSONGetItem(data, "image")
	i.Updated = JSONGetTime(data, "updated")
	inReplyTo := JSONGetItems(data, "inReplyTo")
	if len(inReplyTo) > 0 {
		i.InReplyTo = inReplyTo
	}
	to := JSONGetItems(data, "to")
	if len(to) > 0 {
		i.To = to
	}
	audience := JSONGetItems(data, "audience")
	if len(audience) > 0 {
		i.Audience = audience
	}
	bto := JSONGetItems(data, "bto")
	if len(bto) > 0 {
		i.Bto = bto
	}
	cc := JSONGetItems(data, "cc")
	if len(cc) > 0 {
		i.CC = cc
	}
	bcc := JSONGetItems(data, "bcc")
	if len(bcc) > 0 {
		i.BCC = bcc
	}
	replies := JSONGetItem(data, "replies")
	if replies != nil {
		i.Replies = replies
	}
	tag := JSONGetItems(data, "tag")
	if len(tag) > 0 {
		i.Tag = tag
	}
	i.Actor = JSONGetItem(data, "actor")
	i.Target = JSONGetItem(data, "target")
	i.Result = JSONGetItem(data, "result")
	i.Origin = JSONGetItem(data, "origin")
	i.Instrument = JSONGetItem(data, "instrument")
	return nil
}

// IntransitiveActivityNew initializes a intransitive activity
func IntransitiveActivityNew(id ID, typ ActivityVocabularyType) *IntransitiveActivity {
	if !IntransitiveActivityTypes.Contains(typ) {
		typ = IntransitiveActivityType
	}
	i := IntransitiveActivity{ID: id, Type: typ}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	return &i
}

// ToIntransitiveActivity
func ToIntransitiveActivity(it Item) (*IntransitiveActivity, error) {
	switch i := it.(type) {
	case *Activity:
		return (*IntransitiveActivity)(unsafe.Pointer(i)), nil
	case Activity:
		return (*IntransitiveActivity)(unsafe.Pointer(&i)), nil
	case *Question:
		return (*IntransitiveActivity)(unsafe.Pointer(i)), nil
	case Question:
		return (*IntransitiveActivity)(unsafe.Pointer(&i)), nil
	case *IntransitiveActivity:
		return i, nil
	case IntransitiveActivity:
		return &i, nil
	}
	return nil, errors.New("unable to convert to intransitive activity")
}

// FlattenIntransitiveActivityProperties flattens the IntransitiveActivity's properties from Object type to IRI
func FlattenIntransitiveActivityProperties(act *IntransitiveActivity) *IntransitiveActivity {
	act.Actor = Flatten(act.Actor)
	act.Target = Flatten(act.Target)
	act.Result = Flatten(act.Result)
	act.Origin = Flatten(act.Origin)
	act.Result = Flatten(act.Result)
	act.Instrument = Flatten(act.Instrument)
	return act
}

// ArriveNew initializes an Arrive activity
func ArriveNew(id ID) *Arrive {
	a := IntransitiveActivityNew(id, ArriveType)
	o := Arrive(*a)
	return &o
}

// TravelNew initializes a Travel activity
func TravelNew(id ID) *Travel {
	a := IntransitiveActivityNew(id, TravelType)
	o := Travel(*a)
	return &o
}
