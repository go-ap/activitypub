package activitystreams

import (
	"bytes"
	"fmt"
	"github.com/buger/jsonparser"
	json "github.com/go-ap/jsonld"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// ObjectID designates an unique global identifier.
// All Objects in [ActivityStreams] should have unique global identifiers.
// ActivityPub extends this requirement; all objects distributed by the ActivityPub protocol MUST
// have unique global identifiers, unless they are intentionally transient
// (short lived activities that are not intended to be able to be looked up,
// such as some kinds of chat messages or game notifications).
// These identifiers must fall into one of the following groups:
//
// 1. Publicly dereferenceable URIs, such as HTTPS URIs, with their authority belonging
// to that of their originating server. (Publicly facing content SHOULD use HTTPS URIs).
// 2. An ID explicitly specified as the JSON null object, which implies an anonymous object
// (a part of its parent context)
type ObjectID IRI

const (
	// ActivityBaseURI the basic URI for the activity streams namespaces
	ActivityBaseURI                                  = IRI("https://www.w3.org/ns/activitystreams")
	ObjectType                ActivityVocabularyType = "Object"
	LinkType                  ActivityVocabularyType = "Link"
	ActivityType              ActivityVocabularyType = "Activity"
	IntransitiveActivityType  ActivityVocabularyType = "IntransitiveActivity"
	ActorType                 ActivityVocabularyType = "Actor"
	CollectionType            ActivityVocabularyType = "Collection"
	OrderedCollectionType     ActivityVocabularyType = "OrderedCollection"
	CollectionPageType        ActivityVocabularyType = "CollectionPage"
	OrderedCollectionPageType ActivityVocabularyType = "OrderedCollectionPage"

	// Activity Pub Object Types
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

const (
	NilLangRef LangRef = "-"
)

var GenericObjectTypes = ActivityVocabularyTypes{
	ActivityType,
	IntransitiveActivityType,
	ObjectType,
	ActorType,
	CollectionType,
	OrderedCollectionType,
}

var GenericLinkTypes = ActivityVocabularyTypes{
	LinkType,
}

var GenericTypes = append(GenericObjectTypes[:], GenericLinkTypes[:]...)

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
		GetID() *ObjectID
		// GetType returns the ActivityStreams type
		GetType() ActivityVocabularyType
	}
	// LinkOrIRI is an interface that Object and Link structs implement, and at the same time
	// they are kept disjointed
	LinkOrIRI interface {
		// GetLink returns the object id in IRI type
		GetLink() IRI
	}
	// Item describes the requirements of an ActivityStreams object
	ObjectOrLink interface {
		ActivityObject
		LinkOrIRI
		// IsLink shows if current item represents a Link object or an IRI
		IsLink() bool
		// IsObject shows if current item represents an ActivityStreams object
		IsObject() bool
		// IsCollection shows if the current item represents an ItemCollection
		IsCollection()bool
	}
	// Mapper interface allows external objects to implement their own mechanism for loading information
	// from an ActivityStreams vocabulary object
	Mapper interface {
		// FromActivityStreams maps an ActivityStreams object to another struct representation
		FromActivityStreams(Item) error
	}

	// MimeType is the type for representing MIME types in certain ActivityStreams properties
	MimeType string
	// LangRef is the type for a language reference code, should be an ISO639-1 language specifier.
	LangRef string

	// LangRefValue is a type for storing per language values
	LangRefValue struct {
		Ref   LangRef
		Value string
	}
	// NaturalLanguageValues is a mapping for multiple language values
	NaturalLanguageValues []LangRefValue
)

func NaturalLanguageValuesNew() NaturalLanguageValues {
	return make(NaturalLanguageValues, 0)
}

func (n NaturalLanguageValues) String() string {
	cnt := len(n)
	if cnt == 1 {
		return n[0].String()
	}
	s := strings.Builder{}
	s.Write([]byte{'['})
	for k, v := range n {
		s.WriteString(v.String())
		if k != cnt-1 {
			s.Write([]byte{','})
		}
	}
	s.Write([]byte{']'})
	return s.String()
}

func (n NaturalLanguageValues) Get(ref LangRef) string {
	for _, val := range n {
		if val.Ref == ref {
			return val.Value
		}
	}
	return ""
}

// Set sets a language, value pair in a NaturalLanguageValues array
func (n *NaturalLanguageValues) Set(ref LangRef, v string) error {
	found := false
	for k, vv := range *n {
		if vv.Ref == ref {
			(*n)[k] = LangRefValue{ref, v}
			found = true
		}
	}
	if !found {
		n.Append(ref, v)
	}
	return nil
}

// MarshalJSON serializes the NaturalLanguageValues into JSON
func (n NaturalLanguageValues) MarshalJSON() ([]byte, error) {
	if len(n) == 0 {
		return json.Marshal(nil)
	}
	if len(n) == 1 {
		v := n[0]
		if v.Ref == NilLangRef {
			return v.MarshalJSON()
		}
	}
	mm := make(map[LangRef]string)
	for _, val := range n {
		mm[val.Ref] = val.Value
	}

	return json.Marshal(mm)
}

// First returns the first element in the array
func (n NaturalLanguageValues) First() LangRefValue {
	for _, v := range n {
		return v
	}
	return LangRefValue{}
}

// MarshalText serializes the NaturalLanguageValues into Text
func (n NaturalLanguageValues) MarshalText() ([]byte, error) {
	for _, v := range n {
		return []byte(fmt.Sprintf("%q", v)), nil
	}
	return nil, nil
}

// Append is syntactic sugar for resizing the NaturalLanguageValues map
// and appending an element
func (n *NaturalLanguageValues) Append(lang LangRef, value string) error {
	var t NaturalLanguageValues
	if len(*n) == 0 {
		t = make(NaturalLanguageValues, 0)
	} else {
		t = *n
	}
	t = append(*n, LangRefValue{lang, value})
	*n = t

	return nil
}

// Count returns the length of Items in the item collection
func (n *NaturalLanguageValues) Count() uint {
	return uint(len(*n))
}

// String adds support for Stringer interface. It returns the Value[LangRef] text or just Value if LangRef is NIL
func (l LangRefValue) String() string {
	if l.Ref == NilLangRef {
		return l.Value
	}
	return fmt.Sprintf("%s[%s]", l.Value, l.Ref)
}

// UnmarshalJSON implements the JsonEncoder interface
func (l *LangRefValue) UnmarshalJSON(data []byte) error {
	_, typ, _, err := jsonparser.Get(data)
	if err != nil {
		l.Ref = NilLangRef
		l.Value = string(unescape(data))
		return nil
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			l.Ref = LangRef(key)
			l.Value = string(unescape(value))
			return err
		})
	case jsonparser.String:
		l.Ref = NilLangRef
		l.Value = string(unescape(data))
	}

	return nil
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRefValue) UnmarshalText(data []byte) error {
	return nil
}

// MarshalJSON serializes the LangRefValue into JSON
func (l LangRefValue) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	if l.Ref != NilLangRef {
		if l.Value == "" {
			return nil, nil
		}
		buf.WriteByte('"')
		buf.WriteString(string(l.Ref))
		buf.Write([]byte{'"', ':'})
	}
	l.Value = string(unescape([]byte(l.Value)))
	buf.WriteString(strconv.Quote(l.Value))
	return buf.Bytes(), nil
}

// MarshalText serializes the LangRefValue into JSON
func (l LangRefValue) MarshalText() ([]byte, error) {
	if l.Ref != NilLangRef && l.Value == "" {
		return nil, nil
	}
	buf := bytes.Buffer{}
	buf.WriteString(l.Value)
	if l.Ref != NilLangRef {
		buf.WriteByte('[')
		buf.WriteString(string(l.Ref))
		buf.WriteByte(']')
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON implements the JsonEncoder interface
func (l *LangRef) UnmarshalJSON(data []byte) error {
	return l.UnmarshalText(data)
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRef) UnmarshalText(data []byte) error {
	*l = LangRef("")
	if len(data) == 0 {
		return nil
	}
	if len(data) > 2 {
		if data[0] == '"' && data[len(data)-1] == '"' {
			*l = LangRef(data[1 : len(data)-1])
		}
	} else {
		*l = LangRef(data)
	}
	return nil
}

func unescape(b []byte) []byte {
	// FIMXE(marius): I feel like I'm missing something really obvious about encoding/decoding from Json regarding
	//    escape characters, and that this function is just a hack. Be better future Marius, find the real problem!
	b = bytes.ReplaceAll(b, []byte{'\\', 'a'}, []byte{'\a'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'f'}, []byte{'\f'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'n'}, []byte{'\n'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'r'}, []byte{'\r'})
	b = bytes.ReplaceAll(b, []byte{'\\', 't'}, []byte{'\t'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'v'}, []byte{'\v'})
	b = bytes.ReplaceAll(b, []byte{'\\', '"'}, []byte{'"'})
	b = bytes.ReplaceAll(b, []byte{'\\', '\\'}, []byte{'\\'}) // this should cover the case of \\u -> \u
	return b
}

// UnmarshalJSON tries to load the NaturalLanguage array from the incoming json value
func (n *NaturalLanguageValues) UnmarshalJSON(data []byte) error {
	val, typ, _, err := jsonparser.Get(data)
	if err != nil {
		// try our luck if data contains an unquoted string
		n.Append(NilLangRef, string(unescape(data)))
		return nil
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, val []byte, dataType jsonparser.ValueType, offset int) error {
			n.Append(LangRef(key), string(unescape(val)))
			return err
		})
	case jsonparser.String:
		n.Append(NilLangRef, string(unescape(val)))
	case jsonparser.Array:
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			l := LangRefValue{}
			l.UnmarshalJSON(value)
			n.Append(l.Ref, l.Value)
		})
	}

	return nil
}

// UnmarshalText tries to load the NaturalLanguage array from the incoming Text value
func (n *NaturalLanguageValues) UnmarshalText(data []byte) error {
	if data[0] == '"' {
		// a quoted string - loading it to c.URL
		if data[len(data)-1] != '"' {
			return fmt.Errorf("invalid string value when unmarshaling %T value", n)
		}
		n.Append(LangRef(NilLangRef), string(data[1:len(data)-1]))
	}
	return nil
}
// Describes an object of any kind.
// The Activity Pub Object type serves as the base type for most of the other kinds of objects defined in the Activity
// Vocabulary, including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.
type Object struct {
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

// GetID returns the ObjectID corresponding to the current Object
func (o Object) GetID() *ObjectID {
	return &o.ID
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

// UnmarshalJSON
func (o *Object) UnmarshalJSON(data []byte) error {
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
	return nil
}

// Recipients performs recipient de-duplication on the Object's To, Bto, CC and BCC properties
func (o *Object) Recipients() ItemCollection {
	var aud ItemCollection
	rec, _ := ItemCollectionDeduplication(&aud, &o.To, &o.Bto, &o.CC, &o.BCC, &o.Audience)
	return rec
}

// Clean removes Bto and BCC properties
func (o *Object) Clean(){
	o.BCC = nil
	o.Bto = nil
}

type (
	object = Object
	Parent = Object
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

// ItemCollectionDeduplication normalizes the received arguments lists into a single unified one
func ItemCollectionDeduplication(recCols ...*ItemCollection) (ItemCollection, error) {
	rec := make(ItemCollection, 0)

	for _, recCol := range recCols {
		if recCol == nil {
			continue
		}

		toRemove := make([]int, 0)
		for i, cur := range *recCol {
			save := true
			if cur == nil {
				continue
			}
			var testIt Item
			if cur.IsObject() {
				testIt = cur
			} else if cur.IsLink() {
				testIt = cur.GetLink()
			} else {
				continue
			}
			for _, it := range rec {
				if testIt == it {
					// mark the element for removal
					toRemove = append(toRemove, i)
					save = false
				}
			}
			if save {
				rec = append(rec, testIt)
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(toRemove)))
		for _, idx := range toRemove {
			*recCol = append((*recCol)[:idx], (*recCol)[idx+1:]...)
		}
	}
	return rec, nil
}

// UnmarshalJSON
func (i *ObjectID) UnmarshalJSON(data []byte) error {
	*i = ObjectID(strings.Trim(string(data), "\""))
	return nil
}

// UnmarshalJSON
func (c *MimeType) UnmarshalJSON(data []byte) error {
	*c = MimeType(strings.Trim(string(data), "\""))
	return nil
}

// ToObject
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
	case *Activity:
		return &i.Parent, nil
	case Activity:
		return &i.Parent, nil
	case *IntransitiveActivity:
		return (*Object)(unsafe.Pointer(&i.Parent)), nil
	case IntransitiveActivity:
		return (*Object)(unsafe.Pointer(&i.Parent)), nil
	case *Question:
		return (*Object)(unsafe.Pointer(&i)), nil
	case Question:
		return (*Object)(unsafe.Pointer(&i)), nil
	}
	return nil, fmt.Errorf("unable to convert %q", it.GetType())
}

// FlattenObjectProperties flattens the Object's properties from Object types to IRI
func FlattenObjectProperties(o *Object) *Object {
	if o.Replies != nil && o.Replies.IsObject() {
		if len(o.Replies.GetLink()) > 0 {
			o.Replies = o.Replies.GetLink()
		} else {
			o.Replies = IRI(fmt.Sprintf("%s/replies", o.GetLink()))
		}
	}
	o.AttributedTo = Flatten(o.AttributedTo)
	o.To = FlattenItemCollection(o.To)
	o.Bto = FlattenItemCollection(o.Bto)
	o.CC = FlattenItemCollection(o.CC)
	o.BCC = FlattenItemCollection(o.BCC)
	o.Audience = FlattenItemCollection(o.Audience)

	o.Tag = FlattenItemCollection(o.Tag)

	return o
}

// FlattenProperties flattens the Item's properties from Object types to IRI
func FlattenProperties(it Item) Item {
	if ActivityTypes.Contains(it.GetType()) {
		act, err := ToActivity(it)
		if err == nil {
			return FlattenActivityProperties(act)
		}
	}
	if ActorTypes.Contains(it.GetType()) || ObjectTypes.Contains(it.GetType()) {
		ob, err := ToObject(it)
		if err == nil {
			return FlattenObjectProperties(ob)
		}
	}
	return it
}
