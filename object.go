package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

const (
	IRIType                   ActivityVocabularyType = "IRI"
	ObjectType                ActivityVocabularyType = "Object"
	LinkType                  ActivityVocabularyType = "Link"
	ActivityType              ActivityVocabularyType = "Activity"
	IntransitiveActivityType  ActivityVocabularyType = "IntransitiveActivity"
	ActorType                 ActivityVocabularyType = "Actor"
	CollectionType            ActivityVocabularyType = "Collection"
	OrderedCollectionType     ActivityVocabularyType = "OrderedCollection"
	CollectionPageType        ActivityVocabularyType = "CollectionPage"
	OrderedCollectionPageType ActivityVocabularyType = "OrderedCollectionPage"

	// ActivityPub Object Types
	ArticleType      ActivityVocabularyType = "Article"
	AudioType        ActivityVocabularyType = "Audio"
	DocumentType     ActivityVocabularyType = "Document"
	EventType        ActivityVocabularyType = "Event"
	ImageType        ActivityVocabularyType = "Image"
	NoteType         ActivityVocabularyType = "Note"
	PageType         ActivityVocabularyType = "Page"
	PlaceType        ActivityVocabularyType = "Place"
	ProfileType      ActivityVocabularyType = "Profile"
	RelationshipType ActivityVocabularyType = "Relationship"
	TombstoneType    ActivityVocabularyType = "Tombstone"
	VideoType        ActivityVocabularyType = "Video"

	// MentionType is a link type for @mentions
	MentionType ActivityVocabularyType = "Mention"
)

var GenericTypes = ActivityVocabularyTypes{
	ActivityType,
	IntransitiveActivityType,
	ObjectType,
	ActorType,
}

var ObjectTypes = ActivityVocabularyTypes{
	ArticleType,
	AudioType,
	DocumentType,
	EventType,
	ImageType,
	NoteType,
	PageType,
	PlaceType,
	ProfileType,
	RelationshipType,
	TombstoneType,
	VideoType,
}

type (
	// ActivityVocabularyType is the data type for an Activity type object
	ActivityVocabularyType string
	// ActivityObject is a subtype of Object that describes some form of action that may happen,
	// is currently happening, or has already happened
	ActivityObject interface {
		// GetID returns the dereferenceable ActivityStreams object id
		GetID() ID
		// GetType returns the ActivityStreams type
		GetType() ActivityVocabularyType
	}
	// LinkOrIRI is an interface that Object and Link structs implement, and at the same time
	// they are kept disjointed
	LinkOrIRI interface {
		// GetLink returns the object id in IRI type
		GetLink() IRI
	}
	// ObjectOrLink describes the requirements of an ActivityStreams object
	ObjectOrLink interface {
		ActivityObject
		LinkOrIRI
		// IsLink shows if current item represents a Link object or an IRI
		IsLink() bool
		// IsObject shows if current item represents an ActivityStreams object
		IsObject() bool
		// IsCollection shows if the current item represents an ItemCollection
		IsCollection() bool
	}
	// Mapper interface allows external objects to implement their own mechanism for loading information
	// from an ActivityStreams vocabulary object
	Mapper interface {
		// FromActivityStreams maps an ActivityStreams object to another struct representation
		FromActivityStreams(Item) error
	}

	// MimeType is the type for representing MIME types in certain ActivityStreams properties
	MimeType string
)

func (a ActivityVocabularyType) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return nil, nil
	}
	b := make([]byte, 0)
	writeStringJSONValue(&b, string(a))
	return b, nil
}

// GobEncode
func (a ActivityVocabularyType) GobEncode() ([]byte, error) {
	if len(a) == 0 {
		return []byte{}, nil
	}
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gobEncodeStringLikeType(gg, []byte(a)); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// GobDecode
func (a *ActivityVocabularyType) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	var bb []byte
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&bb); err != nil {
		return err
	}
	*a = ActivityVocabularyType(bb)
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (a *ActivityVocabularyType) UnmarshalBinary(data []byte) error {
	return a.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (a ActivityVocabularyType) MarshalBinary() ([]byte, error) {
	return a.GobEncode()
}

// Object describes an ActivityPub object of any kind.
// It serves as the base type for most of the other kinds of objects defined in the Activity
// Vocabulary, including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.
type Object struct {
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
}

// ObjectNew initializes a new Object
func ObjectNew(typ ActivityVocabularyType) *Object {
	if !(ObjectTypes.Contains(typ)) {
		typ = ObjectType
	}
	o := Object{Type: typ}
	o.Name = NaturalLanguageValuesNew()
	o.Content = NaturalLanguageValuesNew()
	return &o
}

// GetID returns the ID corresponding to the current Object
func (o Object) GetID() ID {
	return o.ID
}

// GetLink returns the IRI corresponding to the current Object
func (o Object) GetLink() IRI {
	return IRI(o.ID)
}

// GetType returns the type of the current Object
func (o Object) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink validates if currentActivity Pub Object is a Link
func (o Object) IsLink() bool {
	return false
}

// IsObject validates if currentActivity Pub Object is an Object
func (o Object) IsObject() bool {
	return true
}

// IsCollection returns false for Object objects
func (o Object) IsCollection() bool {
	return false
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (o *Object) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return loadObject(val, o)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (o Object) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	write(&b, '{')

	if writeObjectJSONValue(&b, o) {
		write(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (o *Object) UnmarshalBinary(data []byte) error {
	return o.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (o Object) MarshalBinary() ([]byte, error) {
	return o.GobEncode()
}

func gobDecodeItem(it Item, data []byte) error {
	return nil
}

func gobEncodeItemCollection(g *gob.Encoder, col ItemCollection) error {
	return g.Encode(col)
}

func gobEncodeItem(it Item) ([]byte, error) {
	b := bytes.Buffer{}
	var err error
	if IsIRI(it) {
		g := gob.NewEncoder(&b)
		err = gobEncodeStringLikeType(g, []byte(it.GetLink()))
	}
	if IsItemCollection(it) {
		g := gob.NewEncoder(&b)
		err = OnItemCollection(it, func(col *ItemCollection) error {
			return gobEncodeItemCollection(g, *col)
		})
	}
	switch it.GetType() {
	case "", ObjectType, ArticleType, AudioType, DocumentType, EventType, ImageType, NoteType, PageType, VideoType:
		err = OnObject(it, func(ob *Object) error {
			bytes, err := ob.GobEncode()
			b.Write(bytes)
			return err
		})
	case LinkType, MentionType:
		err = OnLink(it, func(l *Link) error {
			bytes, err := l.GobEncode()
			b.Write(bytes)
			return err
		})
	case ActivityType, AcceptType, AddType, AnnounceType, BlockType, CreateType, DeleteType, DislikeType,
		FlagType, FollowType, IgnoreType, InviteType, JoinType, LeaveType, LikeType, ListenType, MoveType, OfferType,
		RejectType, ReadType, RemoveType, TentativeRejectType, TentativeAcceptType, UndoType, UpdateType, ViewType:
		err = OnActivity(it, func(act *Activity) error {
			bytes, err := act.GobEncode()
			b.Write(bytes)
			return err
		})
	case IntransitiveActivityType, ArriveType, TravelType:
		err = OnIntransitiveActivity(it, func(act *IntransitiveActivity) error {
			bytes, err := act.GobEncode()
			b.Write(bytes)
			return err
		})
	case ActorType, ApplicationType, GroupType, OrganizationType, PersonType, ServiceType:
		err = OnActor(it, func(a *Actor) error {
			bytes, err := a.GobEncode()
			b.Write(bytes)
			return err
		})
	case CollectionType:
		err = OnCollection(it, func(c *Collection) error {
			return nil
		})
	case OrderedCollectionType:
		err = OnOrderedCollection(it, func(c *OrderedCollection) error {
			return nil
		})
	case CollectionPageType:
		err = OnCollectionPage(it, func(p *CollectionPage) error {
			return nil
		})
	case OrderedCollectionPageType:
		err = OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
			return nil
		})
	case PlaceType:
		err = OnPlace(it, func(p *Place) error {
			return nil
		})
	case ProfileType:
		err = OnProfile(it, func(p *Profile) error {
			return nil
		})
	case RelationshipType:
		err = OnRelationship(it, func(r *Relationship) error {
			return nil
		})
	case TombstoneType:
		err = OnTombstone(it, func(t *Tombstone) error {
			return nil
		})
	case QuestionType:
		err = OnQuestion(it, func(q *Question) error {
			return nil
		})
	}
	return b.Bytes(), err
}

func mapObjectProperties(mm map[string][]byte, o *Object) (hasData bool, err error) {
	if len(o.ID) > 0 {
		if mm["id"], err = o.ID.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Type) > 0 {
		if mm["type"], err = o.Type.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.MediaType) > 0 {
		if mm["mediaType"], err = o.MediaType.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Name) > 0 {
		if mm["name"], err = o.Name.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Attachment != nil {
		if mm["attachment"], err = gobEncodeItem(o.Attachment); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.AttributedTo != nil {
		if mm["attributedTo"], err = gobEncodeItem(o.AttributedTo); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Audience != nil {
		if mm["audience"], err = gobEncodeItem(o.Audience); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Content != nil {
		if mm["content"], err = o.Content.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Context != nil {
		if mm["context"], err = gobEncodeItem(o.Context); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.MediaType) > 0 {
		if mm["mediaType"], err = o.MediaType.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.EndTime.IsZero() {
		if mm["endTime"], err = o.EndTime.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Generator != nil {
		if mm["generator"], err = gobEncodeItem(o.Generator); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Icon != nil {
		if mm["icon"], err = gobEncodeItem(o.Icon); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Image != nil {
		if mm["image"], err = gobEncodeItem(o.Image); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.InReplyTo != nil {
		if mm["inReplyTo"], err = gobEncodeItem(o.InReplyTo); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Location != nil {
		if mm["location"], err = gobEncodeItem(o.Location); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Preview != nil {
		if mm["preview"], err = gobEncodeItem(o.Preview); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.Published.IsZero() {
		if mm["published"], err = o.Published.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Replies != nil {
		if mm["replies"], err = gobEncodeItem(o.Replies); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.StartTime.IsZero() {
		if mm["startTime"], err = o.StartTime.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Summary) > 0 {
		if mm["summary"], err = o.Summary.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Tag != nil {
		if mm["tag"], err = gobEncodeItem(o.Tag); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.Updated.IsZero() {
		if mm["updated"], err = o.Updated.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Tag != nil {
		if mm["tag"], err = gobEncodeItem(o.Tag); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.Updated.IsZero() {
		if mm["updated"], err = o.Updated.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.URL != nil {
		if mm["url"], err = gobEncodeItem(o.URL); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.To != nil {
		if mm["to"], err = gobEncodeItem(o.To); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Bto != nil {
		if mm["bto"], err = gobEncodeItem(o.Bto); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.CC != nil {
		if mm["cc"], err = gobEncodeItem(o.CC); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.BCC != nil {
		if mm["bcc"], err = gobEncodeItem(o.BCC); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Duration > 0 {
		if mm["duration"], err = gobEncodeInt64(int64(o.Duration)); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Likes != nil {
		if mm["likes"], err = gobEncodeItem(o.Likes); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Shares != nil {
		if mm["shares"], err = gobEncodeItem(o.Shares); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Shares != nil {
		if mm["shares"], err = gobEncodeItem(o.Shares); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Source.MediaType)+len(o.Source.Content) > 0 {
		if mm["source"], err = o.Source.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}

	return hasData, nil
}

// GobEncode
func (o Object) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapObjectProperties(mm, &o)
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

func unmapObjectProperties(mm map[string][]byte, o *Object) error {
	if raw, ok := mm["id"]; ok {
		if err := o.ID.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["type"]; ok {
		if err := o.Type.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["mediaType"]; ok {
		if err := o.MediaType.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["name"]; ok {
		if err := o.Name.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["attachment"]; ok {
		if err := gobDecodeItem(o.Attachment, raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["source"]; ok {
		if err := o.Source.GobDecode(raw); err != nil {
			return err
		}
	}
	return nil
}

// GobDecode
func (o *Object) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm := make(map[string][]byte)
	g := gob.NewDecoder(bytes.NewReader(data))
	if err := g.Decode(&mm); err != nil {
		return err
	}
	return unmapObjectProperties(mm, o)
}

// Recipients performs recipient de-duplication on the Object's To, Bto, CC and BCC properties
func (o *Object) Recipients() ItemCollection {
	var aud ItemCollection
	return ItemCollectionDeduplication(&aud, &o.To, &o.Bto, &o.CC, &o.BCC, &o.Audience)
}

// Clean removes Bto and BCC properties
func (o *Object) Clean() {
	o.BCC = nil
	o.Bto = nil
}

type (
	// Article represents any kind of multi-paragraph written work.
	Article = Object
	// Audio represents an audio document of any kind.
	Audio = Document
	// Document represents a document of any kind.
	Document = Object
	// Event represents any kind of event.
	Event = Object
	// Image An image document of any kind
	Image = Document
	// Note represents a short written work typically less than a single paragraph in length.
	Note = Object
	// Page represents a Web Page.
	Page = Document
	// Video represents a video document of any kind
	Video = Document
)

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (m *MimeType) UnmarshalJSON(data []byte) error {
	*m = MimeType(strings.Trim(string(data), "\""))
	return nil
}

// MarshalJSON encodes the receiver object to a JSON document.
func (m MimeType) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return nil, nil
	}
	b := make([]byte, 0)
	writeStringJSONValue(&b, string(m))
	return b, nil
}

// GobEncode
func (m MimeType) GobEncode() ([]byte, error) {
	if len(m) == 0 {
		return []byte{}, nil
	}
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gobEncodeStringLikeType(gg, []byte(m)); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// GobDecode
func (m *MimeType) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	var bb []byte
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&bb); err != nil {
		return err
	}
	*m = MimeType(bb)
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (m *MimeType) UnmarshalBinary(data []byte) error {
	return m.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (m MimeType) MarshalBinary() ([]byte, error) {
	return m.GobEncode()
}

// ToLink returns a Link pointer to the data in the current Item
func ToLink(it Item) (*Link, error) {
	switch i := it.(type) {
	case *Link:
		return i, nil
	case Link:
		return &i, nil
	}
	return nil, fmt.Errorf("unable to convert %q", it.GetType())
}

// ToObject returns an Object pointer to the data in the current Item
// It relies on the fact that all the types in this package have a data layout compatible with Object.
func ToObject(it Item) (*Object, error) {
	switch i := it.(type) {
	case *Object:
		return i, nil
	case Object:
		return &i, nil
	case *Place:
		return (*Object)(unsafe.Pointer(i)), nil
	case Place:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Profile:
		return (*Object)(unsafe.Pointer(i)), nil
	case Profile:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Relationship:
		return (*Object)(unsafe.Pointer(i)), nil
	case Relationship:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Tombstone:
		return (*Object)(unsafe.Pointer(i)), nil
	case Tombstone:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Actor:
		return (*Object)(unsafe.Pointer(i)), nil
	case Actor:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Activity:
		return (*Object)(unsafe.Pointer(i)), nil
	case Activity:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *IntransitiveActivity:
		return (*Object)(unsafe.Pointer(i)), nil
	case IntransitiveActivity:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Question:
		return (*Object)(unsafe.Pointer(i)), nil
	case Question:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *Collection:
		return (*Object)(unsafe.Pointer(i)), nil
	case Collection:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *CollectionPage:
		return (*Object)(unsafe.Pointer(i)), nil
	case CollectionPage:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *OrderedCollection:
		return (*Object)(unsafe.Pointer(i)), nil
	case OrderedCollection:
		return (*Object)(unsafe.Pointer(&i)), nil
	case *OrderedCollectionPage:
		return (*Object)(unsafe.Pointer(i)), nil
	case OrderedCollectionPage:
		return (*Object)(unsafe.Pointer(&i)), nil
	default:
		// NOTE(marius): this is an ugly way of dealing with the interface conversion error: types from different scopes
		typ := reflect.TypeOf(new(Object))
		if reflect.TypeOf(it).ConvertibleTo(typ) {
			if i, ok := reflect.ValueOf(it).Convert(typ).Interface().(*Object); ok {
				return i, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to convert %q", it.GetType())
}

// Source is intended to convey some sort of source from which the content markup was derived,
// as a form of provenance, or to support future editing by clients.
type Source struct {
	// Content
	Content NaturalLanguageValues `jsonld:"content"`
	// MediaType
	MediaType MimeType `jsonld:"mediaType"`
}

// GetAPSource
func GetAPSource(val *fastjson.Value) Source {
	s := Source{}
	if val == nil {
		return s
	}

	if contBytes := val.Get("source", "content").GetStringBytes(); len(contBytes) > 0 {
		s.Content.UnmarshalJSON(contBytes)
	}
	if mimeBytes := val.Get("source", "mediaType").GetStringBytes(); len(mimeBytes) > 0 {
		s.MediaType.UnmarshalJSON(mimeBytes)
	}

	return s
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (s *Source) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	*s = GetAPSource(val)
	return nil
}

// MarshalJSON encodes the receiver object to a JSON document.
func (s Source) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	empty := true
	write(&b, '{')
	if len(s.MediaType) > 0 {
		if v, err := s.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
			empty = !writeJSONProp(&b, "mediaType", v)
		}
	}
	if len(s.Content) > 0 {
		empty = !writeNaturalLanguageJSONProp(&b, "content", s.Content)
	}
	if !empty {
		write(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (s *Source) UnmarshalBinary(data []byte) error {
	return s.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (s Source) MarshalBinary() ([]byte, error) {
	return s.GobEncode()
}

// GobDecode
func (s *Source) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm := make(map[string][]byte)
	g := gob.NewDecoder(bytes.NewReader(data))
	if err := g.Decode(&mm); err != nil {
		return err
	}
	if raw, ok := mm["mediaType"]; ok {
		if err := s.MediaType.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["content"]; ok {
		if err := s.Content.GobDecode(raw); err != nil {
			return err
		}
	}
	return nil
}

// GobEncode
func (s Source) GobEncode() ([]byte, error) {
	var (
		mm      = make(map[string][]byte)
		err     error
		hasData bool
	)
	if len(s.MediaType) > 0 {
		if mm["mediaType"], err = s.MediaType.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(s.Content) > 0 {
		if mm["content"], err = s.Content.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
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

// Equals verifies if our receiver Object is equals with the "with" Object
func (o Object) Equals(with Item) bool {
	if with.IsCollection() {
		return false
	}
	if withID := with.GetID(); len(withID) > 0 && withID != o.ID {
		return false
	}
	if withType := with.GetType(); len(withType) > 0 && withType != o.Type {
		return false
	}
	if with.IsLink() && !with.GetLink().Equals(o.GetLink(), false) {
		return false
	}
	result := true
	OnObject(with, func(w *Object) error {
		if len(w.Name) > 0 {
			if !w.Name.Equals(o.Name) {
				result = false
				return nil
			}
		}
		if len(w.Summary) > 0 {
			if !w.Summary.Equals(o.Summary) {
				result = false
				return nil
			}
		}
		if len(w.Content) > 0 {
			if !w.Content.Equals(o.Content) {
				result = false
				return nil
			}
		}
		if w.Attachment != nil {
			if !ItemsEqual(o.Attachment, w.Attachment) {
				result = false
				return nil
			}
		}
		if w.AttributedTo != nil {
			if !ItemsEqual(o.AttributedTo, w.AttributedTo) {
				result = false
				return nil
			}
		}
		if w.Audience != nil {
			if !ItemsEqual(o.Audience, w.Audience) {
				result = false
				return nil
			}
		}
		if w.Context != nil {
			if !ItemsEqual(o.Context, w.Context) {
				result = false
				return nil
			}
		}
		if w.Generator != nil {
			if !ItemsEqual(o.Generator, w.Generator) {
				result = false
				return nil
			}
		}
		if w.Icon != nil {
			if !ItemsEqual(o.Icon, w.Icon) {
				result = false
				return nil
			}
		}
		if w.Image != nil {
			if !ItemsEqual(o.Image, w.Image) {
				result = false
				return nil
			}
		}
		if w.InReplyTo != nil {
			if !ItemsEqual(o.InReplyTo, w.InReplyTo) {
				result = false
				return nil
			}
		}
		if w.Location != nil {
			if !ItemsEqual(o.Location, w.Location) {
				result = false
				return nil
			}
		}
		if w.Preview != nil {
			if !ItemsEqual(o.Preview, w.Preview) {
				result = false
				return nil
			}
		}
		if w.Replies != nil {
			if !ItemsEqual(o.Replies, w.Replies) {
				result = false
				return nil
			}
		}
		if w.Tag != nil {
			if !ItemsEqual(o.Tag, w.Tag) {
				result = false
				return nil
			}
		}
		if w.URL != nil {
			if o.URL == nil {
				result = false
				return nil
			}
			if !w.URL.GetLink().Equals(o.URL.GetLink(), false) {
				result = false
				return nil
			}
		}
		if w.To != nil {
			if !ItemsEqual(o.To, w.To) {
				result = false
				return nil
			}
		}
		if w.Bto != nil {
			if !ItemsEqual(o.Bto, w.Bto) {
				result = false
				return nil
			}
		}
		if w.CC != nil {
			if !ItemsEqual(o.CC, w.CC) {
				result = false
				return nil
			}
		}
		if w.BCC != nil {
			if !ItemsEqual(o.BCC, w.BCC) {
				result = false
				return nil
			}
		}
		if !w.Published.IsZero() {
			if !w.Published.Equal(o.Published) {
				result = false
				return nil
			}
		}
		if !w.Updated.IsZero() {
			if !w.Updated.Equal(o.Updated) {
				result = false
				return nil
			}
		}
		if !w.StartTime.IsZero() {
			if !w.StartTime.Equal(o.StartTime) {
				result = false
				return nil
			}
		}
		if !w.EndTime.IsZero() {
			if !w.EndTime.Equal(o.EndTime) {
				result = false
				return nil
			}
		}
		if w.Duration != 0 {
			if w.Duration != o.Duration {
				result = false
				return nil
			}
		}
		if w.Likes != nil {
			if !ItemsEqual(o.Likes, w.Likes) {
				result = false
				return nil
			}
		}
		if w.Shares != nil {
			if !ItemsEqual(o.Shares, w.Shares) {
				result = false
				return nil
			}
		}
		return nil
	})
	return result
}
