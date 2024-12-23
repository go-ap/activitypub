package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
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
	JSONWriteStringValue(&b, string(a))
	return b, nil
}

// GobEncode
func (a ActivityVocabularyType) GobEncode() ([]byte, error) {
	return []byte(a), nil
}

// GobDecode
func (a *ActivityVocabularyType) GobDecode(data []byte) error {
	*a = ActivityVocabularyType(data)
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

type Objects interface {
	Object | Tombstone | Place | Profile | Relationship |
		Actors |
		Activities |
		IntransitiveActivities |
		Collections |
		IRI
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
	return JSONLoadObject(val, o)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (o Object) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	JSONWrite(&b, '{')

	if JSONWriteObjectValue(&b, o) {
		JSONWrite(&b, '}')
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

// GobDecode
func (o *Object) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapObjectProperties(mm, o)
}

func fmtObjectProps(w io.Writer) func(*Object) error {
	return func(o *Object) error {
		if len(o.ID) > 0 {
			if n, _ := fmt.Fprintf(w, "ID:%s", o.ID); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if len(o.Name) > 0 {
			if n, _ := fmt.Fprintf(w, "%s: [%s]", "name", o.Name); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if len(o.Summary) > 0 {
			if n, _ := fmt.Fprintf(w, "%s: [%s]", "summary", o.Summary); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if len(o.Content) > 0 {
			if n, _ := fmt.Fprintf(w, "%s: [%s]", "content", o.Content); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Attachment) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "attachment", o.Attachment); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.AttributedTo) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "attributedTo", o.AttributedTo); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Audience) && o.Audience.Count() > 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "audience", o.Audience); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Context) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "context", o.Context); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Generator) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "generator", o.Generator); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Icon) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "icon", o.Icon); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Image) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "image", o.Image); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.InReplyTo) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "inReplyTo", o.InReplyTo); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Location) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "location", o.Location); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Preview) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "preview", o.Preview); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Replies) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "replies", o.Replies); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Tag) && o.Tag.Count() > 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "tag", o.Tag); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.URL) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "url", o.URL); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.To) && o.To.Count() > 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "to", o.To); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Bto) && o.Bto.Count() > 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "bto", o.Bto); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.CC) && o.CC.Count() > 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "cc", o.CC); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.BCC) && o.BCC.Count() > 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "bcc", o.BCC); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !o.Published.IsZero() {
			if n, _ := fmt.Fprintf(w, "%s: %s", "published", o.Published); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !o.Updated.IsZero() {
			if n, _ := fmt.Fprintf(w, "%s: %s", "updated", o.Updated); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !o.StartTime.IsZero() {
			if n, _ := fmt.Fprintf(w, "%s: %s", "startTime", o.StartTime); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !o.EndTime.IsZero() {
			if n, _ := fmt.Fprintf(w, "%s: %s", "endTime", o.EndTime); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if o.Duration != 0 {
			if n, _ := fmt.Fprintf(w, "%s: %s", "duration", o.Duration); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Likes) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "likes", o.Likes); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		if !IsNil(o.Shares) {
			if n, _ := fmt.Fprintf(w, "%s: %s", "shares", o.Shares); n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}
		return nil
	}
}

func (o Object) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		if o.Type != "" && o.ID != "" {
			_, _ = fmt.Fprintf(s, "%T[%s]( %s )", o, o.Type, o.ID)
		} else if o.ID != "" {
			_, _ = fmt.Fprintf(s, "%T( %s )", o, o.ID)
		} else {
			_, _ = fmt.Fprintf(s, "%T[%p]", o, &o)
		}
	case 'v':
		if o.Type != "" && o.ID != "" {
			_, _ = fmt.Fprintf(s, "%T[%s] {", o, o.Type)
			_ = fmtObjectProps(s)(&o)
			_, _ = io.WriteString(s, " }")
		} else if o.ID != "" {
			_, _ = fmt.Fprintf(s, "%T { ", o)
			_ = fmtObjectProps(s)(&o)
			_, _ = io.WriteString(s, " }")
		}
	}
}

// Recipients performs recipient de-duplication on the Object's To, Bto, CC and BCC properties
func (o *Object) Recipients() ItemCollection {
	aud := o.Audience
	return ItemCollectionDeduplication(&o.To, &o.CC, &o.Bto, &o.BCC, &aud)
}

// Clean removes Bto and BCC properties
func (o *Object) Clean() {
	o.BCC = o.BCC[:0]
	o.Bto = o.Bto[:0]
	CleanRecipients(o.Audience)
	CleanRecipients(o.Attachment)
	CleanRecipients(o.Icon)
	CleanRecipients(o.Image)
	CleanRecipients(o.Context)
	CleanRecipients(o.Generator)
	CleanRecipients(o.AttributedTo)
	CleanRecipients(o.Preview)
	CleanRecipients(o.Tag)
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
	JSONWriteStringValue(&b, string(m))
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
func ToLink(it LinkOrIRI) (*Link, error) {
	switch i := it.(type) {
	case *Link:
		return i, nil
	case Link:
		return &i, nil
	}
	return nil, fmt.Errorf("unable to convert %T to %T", it, new(Link))
}

// ToObject returns an Object pointer to the data in the current Item
// It relies on the fact that all the types in this package have a data layout compatible with Object.
func ToObject(it LinkOrIRI) (*Object, error) {
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
		return reflectItemToType[Object](it)
	}
}

func reflectItemToType[T Objects | Links](it LinkOrIRI) (*T, error) {
	if IsNil(it) {
		return nil, nil
	}
	tTyp := reflect.TypeFor[*T]()
	if !reflect.TypeOf(it).ConvertibleTo(tTyp) {
		return nil, ErrorInvalidType[T](it)
	}

	iVal := reflect.ValueOf(it)
	if !iVal.IsValid() {
		return nil, ErrorInvalidType[T](it)
	}
	if i, ok := iVal.Convert(tTyp).Interface().(*T); ok {
		return i, nil
	}
	return nil, ErrorInvalidType[T](it)
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
	JSONWrite(&b, '{')
	if len(s.MediaType) > 0 {
		if v, err := s.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
			empty = !JSONWriteProp(&b, "mediaType", v)
		}
	}
	if len(s.Content) > 0 {
		empty = !JSONWriteNaturalLanguageProp(&b, "content", s.Content)
	}
	if !empty {
		JSONWrite(&b, '}')
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
	if IsItemCollection(with) {
		return false
	}
	if withID := with.GetID(); !o.ID.Equals(withID, true) {
		return false
	}
	if withType := with.GetType(); !strings.EqualFold(string(o.Type), string(withType)) {
		return false
	}
	if with.IsLink() && !with.GetLink().Equals(o.GetLink(), false) {
		return false
	}
	result := true
	err := OnObject(with, func(w *Object) error {
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
	if err != nil {
		result = false
	}
	return result
}
