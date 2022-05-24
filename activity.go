package activitypub

import (
	"bytes"
	"encoding/gob"
	"errors"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

// Activity Types
const (
	AcceptType          ActivityVocabularyType = "Accept"
	AddType             ActivityVocabularyType = "Add"
	AnnounceType        ActivityVocabularyType = "Announce"
	ArriveType          ActivityVocabularyType = "Arrive"
	BlockType           ActivityVocabularyType = "Block"
	CreateType          ActivityVocabularyType = "Create"
	DeleteType          ActivityVocabularyType = "Delete"
	DislikeType         ActivityVocabularyType = "Dislike"
	FlagType            ActivityVocabularyType = "Flag"
	FollowType          ActivityVocabularyType = "Follow"
	IgnoreType          ActivityVocabularyType = "Ignore"
	InviteType          ActivityVocabularyType = "Invite"
	JoinType            ActivityVocabularyType = "Join"
	LeaveType           ActivityVocabularyType = "Leave"
	LikeType            ActivityVocabularyType = "Like"
	ListenType          ActivityVocabularyType = "Listen"
	MoveType            ActivityVocabularyType = "Move"
	OfferType           ActivityVocabularyType = "Offer"
	QuestionType        ActivityVocabularyType = "Question"
	RejectType          ActivityVocabularyType = "Reject"
	ReadType            ActivityVocabularyType = "Read"
	RemoveType          ActivityVocabularyType = "Remove"
	TentativeRejectType ActivityVocabularyType = "TentativeReject"
	TentativeAcceptType ActivityVocabularyType = "TentativeAccept"
	TravelType          ActivityVocabularyType = "Travel"
	UndoType            ActivityVocabularyType = "Undo"
	UpdateType          ActivityVocabularyType = "Update"
	ViewType            ActivityVocabularyType = "View"
)

func (a ActivityVocabularyTypes) Contains(typ ActivityVocabularyType) bool {
	for _, v := range a {
		if strings.ToLower(string(v)) == strings.ToLower(string(typ)) {
			return true
		}
	}
	return false
}

// ContentManagementActivityTypes use case primarily deals with activities that involve the creation, modification or deletion of content.
// This includes, for instance, activities such as "John created a new note", "Sally updated an article", and "Joe deleted the photo".
var ContentManagementActivityTypes = ActivityVocabularyTypes{
	CreateType,
	DeleteType,
	UpdateType,
}

// CollectionManagementActivityTypes use case primarily deals with activities involving the management of content within collections.
// Examples of collections include things like folders, albums, friend lists, etc.
// This includes, for instance, activities such as "Sally added a file to Folder A", "John moved the file from Folder A to Folder B", etc.
var CollectionManagementActivityTypes = ActivityVocabularyTypes{
	AddType,
	MoveType,
	RemoveType,
}

// ReactionsActivityTypes use case primarily deals with reactions to content.
// This can include activities such as liking or disliking content, ignoring updates, flagging content as being inappropriate, accepting or rejecting objects, etc.
var ReactionsActivityTypes = ActivityVocabularyTypes{
	AcceptType,
	BlockType,
	DislikeType,
	FlagType,
	IgnoreType,
	LikeType,
	RejectType,
	TentativeAcceptType,
	TentativeRejectType,
}

// EventRSVPActivityTypes use case primarily deals with invitations to events and RSVP type responses.
var EventRSVPActivityTypes = ActivityVocabularyTypes{
	AcceptType,
	IgnoreType,
	InviteType,
	RejectType,
	TentativeAcceptType,
	TentativeRejectType,
}

// GroupManagementActivityTypes use case primarily deals with management of groups.
// It can include, for instance, activities such as "John added Sally to Group A", "Sally joined Group A", "Joe left Group A", etc.
var GroupManagementActivityTypes = ActivityVocabularyTypes{
	AddType,
	JoinType,
	LeaveType,
	RemoveType,
}

// ContentExperienceActivityTypes use case primarily deals with describing activities involving listening to, reading, or viewing content.
// For instance, "Sally read the article", "Joe listened to the song".
var ContentExperienceActivityTypes = ActivityVocabularyTypes{
	ListenType,
	ReadType,
	ViewType,
}

// GeoSocialEventsActivityTypes use case primarily deals with activities involving geo-tagging type activities.
// For instance, it can include activities such as "Joe arrived at work", "Sally left work", and "John is travel from home to work".
var GeoSocialEventsActivityTypes = ActivityVocabularyTypes{
	ArriveType,
	LeaveType,
	TravelType,
}

// NotificationActivityTypes use case primarily deals with calling attention to particular objects or notifications.
var NotificationActivityTypes = ActivityVocabularyTypes{
	AnnounceType,
}

// QuestionActivityTypes use case primarily deals with representing inquiries of any type.
// See 5.4 Representing Questions for more information.
var QuestionActivityTypes = ActivityVocabularyTypes{
	QuestionType,
}

// RelationshipManagementActivityTypes use case primarily deals with representing activities involving the management of interpersonal and social relationships
// (e.g. friend requests, management of social network, etc). See 5.2 Representing Relationships Between Entities for more information.
var RelationshipManagementActivityTypes = ActivityVocabularyTypes{
	AcceptType,
	AddType,
	BlockType,
	CreateType,
	DeleteType,
	FollowType,
	IgnoreType,
	InviteType,
	RejectType,
}

// NegatingActivityTypes use case primarily deals with the ability to redact previously completed activities.
// See 5.5 Inverse Activities and "Undo" for more information.
var NegatingActivityTypes = ActivityVocabularyTypes{
	UndoType,
}

// OffersActivityTypes use case deals with activities involving offering one object to another.
// It can include, for instance, activities such as "Company A is offering a discount on purchase of Product Z to Sally", "Sally is offering to add a File to Folder A", etc.
var OffersActivityTypes = ActivityVocabularyTypes{
	OfferType,
}

var IntransitiveActivityTypes = ActivityVocabularyTypes{
	ArriveType,
	TravelType,
	QuestionType,
}

var ActivityTypes = ActivityVocabularyTypes{
	AcceptType,
	AddType,
	AnnounceType,
	BlockType,
	CreateType,
	DeleteType,
	DislikeType,
	FlagType,
	FollowType,
	IgnoreType,
	InviteType,
	JoinType,
	LeaveType,
	LikeType,
	ListenType,
	MoveType,
	OfferType,
	RejectType,
	ReadType,
	RemoveType,
	TentativeRejectType,
	TentativeAcceptType,
	UndoType,
	UpdateType,
	ViewType,
}

// HasRecipients is an interface implemented by activities to return their audience
// for further propagation
type HasRecipients interface {
	// Recipients is a method that should do a recipients de-duplication step and then return
	// the remaining recipients
	Recipients() ItemCollection
	Clean()
}

type Activities interface {
	Activity
}

// Activity is a subtype of Object that describes some form of action that may happen,
// is currently happening, or has already happened.
// The Activity type itself serves as an abstract base type for all types of activities.
// It is important to note that the Activity type itself does not carry any specific semantics
// about the kind of action being taken.
type Activity struct {
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
	// CanReceiveActivities describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor Item `jsonld:"actor,omitempty"`
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
	// Object When used within an Activity, describes the direct object of the activity.
	// For instance, in the activity "John added a movie to his wishlist",
	// the object of the activity is the movie added.
	// When used within a Relationship describes the entity to which the subject is related.
	Object Item `jsonld:"object,omitempty"`
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Activity) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink returns false for Activity objects
func (a Activity) IsLink() bool {
	return false
}

// GetID returns the ID corresponding to the Activity object
func (a Activity) GetID() ID {
	return a.ID
}

// GetLink returns the IRI corresponding to the Activity object
func (a Activity) GetLink() IRI {
	return IRI(a.ID)
}

// IsObject returns true for Activity objects
func (a Activity) IsObject() bool {
	return true
}

// IsCollection returns false for Activity objects
func (a Activity) IsCollection() bool {
	return false
}

func removeFromCollection(col ItemCollection, items ...Item) ItemCollection {
	result := make(ItemCollection, 0)
	if len(items) == 0 {
		return col
	}
	for _, ob := range col {
		found := false
		for _, it := range items {
			if ob.GetID().Equals(it.GetID(), false) {
				found = true
				break
			}
		}
		if !found {
			result = append(result, ob)
		}
	}
	return result
}

func removeFromAudience(a *Activity, items ...Item) error {
	if a.To != nil {
		a.To = removeFromCollection(a.To, items...)
	}
	if a.Bto != nil {
		a.Bto = removeFromCollection(a.Bto, items...)
	}
	if a.CC != nil {
		a.CC = removeFromCollection(a.CC, items...)
	}
	if a.BCC != nil {
		a.BCC = removeFromCollection(a.BCC, items...)
	}
	if a.Audience != nil {
		a.Audience = removeFromCollection(a.Audience, items...)
	}
	return nil
}

// Recipients performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (a *Activity) Recipients() ItemCollection {
	var alwaysRemove ItemCollection
	if a.GetType() == BlockType && a.Object != nil {
		alwaysRemove = append(alwaysRemove, a.Object)
	}
	if a.Actor != nil {
		alwaysRemove = append(alwaysRemove, a.Actor)
	}
	if len(alwaysRemove) > 0 {
		removeFromAudience(a, alwaysRemove...)
	}
	return ItemCollectionDeduplication(&a.To, &a.Bto, &a.CC, &a.BCC, &a.Audience)
}

// Clean removes Bto and BCC properties
func (a *Activity) Clean() {
	a.BCC = nil
	a.Bto = nil
	if a.Object != nil && a.Object.IsObject() {
		OnObject(a.Object, func(o *Object) error {
			o.Clean()
			return nil
		})
	}
	if a.Actor != nil && a.Actor.IsObject() {
		OnObject(a.Actor, func(o *Object) error {
			o.Clean()
			return nil
		})
	}
	if a.Target != nil && a.Target.IsObject() {
		OnObject(a.Target, func(o *Object) error {
			o.Clean()
			return nil
		})
	}
}

type (
	// Accept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate
	// the context into which the object has been accepted.
	Accept = Activity

	// Add indicates that the actor has added the object to the target. If the target property is not explicitly specified,
	// the target would need to be determined implicitly by context.
	// The origin can be used to identify the context from which the object originated.
	Add = Activity

	// Announce indicates that the actor is calling the target's attention the object.
	// The origin typically has no defined meaning.
	Announce = Activity

	// Block indicates that the actor is blocking the object. Blocking is a stronger form of Ignore.
	// The typical use is to support social systems that allow one user to block activities or content of other users.
	// The target and origin typically have no defined meaning.
	Block = Ignore

	// Create indicates that the actor has created the object.
	Create = Activity

	// Delete indicates that the actor has deleted the object.
	// If specified, the origin indicates the context from which the object was deleted.
	Delete = Activity

	// Dislike indicates that the actor dislikes the object.
	Dislike = Activity

	// Flag indicates that the actor is "flagging" the object.
	// Flagging is defined in the sense common to many social platforms as reporting content as being
	// inappropriate for any number of reasons.
	Flag = Activity

	// Follow indicates that the actor is "following" the object. Following is defined in the sense typically used within
	// Social systems in which the actor is interested in any activity performed by or on the object.
	// The target and origin typically have no defined meaning.
	Follow = Activity

	// Ignore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	Ignore = Activity

	// Invite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	Invite = Offer

	// Join indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	Join = Activity

	// Leave indicates that the actor has left the object. The target and origin typically have no meaning.
	Leave = Activity

	// Like indicates that the actor likes, recommends or endorses the object.
	// The target and origin typically have no defined meaning.
	Like = Activity

	// Listen inherits all properties from Activity.
	Listen = Activity

	// Move indicates that the actor has moved object from origin to target.
	// If the origin or target are not specified, either can be determined by context.
	Move = Activity

	// Offer indicates that the actor is offering the object.
	// If specified, the target indicates the entity to which the object is being offered.
	Offer = Activity

	// Reject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	Reject = Activity

	// Read indicates that the actor has read the object.
	Read = Activity

	// Remove indicates that the actor is removing the object. If specified,
	// the origin indicates the context from which the object is being removed.
	Remove = Activity

	// TentativeReject is a specialization of Reject in which the rejection is considered tentative.
	TentativeReject = Reject

	// TentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	TentativeAccept = Accept

	// Undo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing
	// some previously performed action (for instance, a person may have previously "liked" an article but,
	// for whatever reason, might choose to undo that like at some later point in time).
	// The target and origin typically have no defined meaning.
	Undo = Activity

	// Update indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism
	// for describing the actual set of modifications made to object.
	// The target and origin typically have no defined meaning.
	Update = Activity

	// View indicates that the actor has viewed the object.
	View = Activity
)

// AcceptNew initializes an Accept activity
func AcceptNew(id ID, ob Item) *Accept {
	a := ActivityNew(id, AcceptType, ob)
	o := Accept(*a)
	return &o
}

// AddNew initializes an Add activity
func AddNew(id ID, ob Item, trgt Item) *Add {
	a := ActivityNew(id, AddType, ob)
	o := Add(*a)
	o.Target = trgt
	return &o
}

// AnnounceNew initializes an Announce activity
func AnnounceNew(id ID, ob Item) *Announce {
	a := ActivityNew(id, AnnounceType, ob)
	o := Announce(*a)
	return &o
}

// BlockNew initializes a Block activity
func BlockNew(id ID, ob Item) *Block {
	a := ActivityNew(id, BlockType, ob)
	o := Block(*a)
	return &o
}

// CreateNew initializes a Create activity
func CreateNew(id ID, ob Item) *Create {
	a := ActivityNew(id, CreateType, ob)
	o := Create(*a)
	return &o
}

// DeleteNew initializes a Delete activity
func DeleteNew(id ID, ob Item) *Delete {
	a := ActivityNew(id, DeleteType, ob)
	o := Delete(*a)
	return &o
}

// DislikeNew initializes a Dislike activity
func DislikeNew(id ID, ob Item) *Dislike {
	a := ActivityNew(id, DislikeType, ob)
	o := Dislike(*a)
	return &o
}

// FlagNew initializes a Flag activity
func FlagNew(id ID, ob Item) *Flag {
	a := ActivityNew(id, FlagType, ob)
	o := Flag(*a)
	return &o
}

// FollowNew initializes a Follow activity
func FollowNew(id ID, ob Item) *Follow {
	a := ActivityNew(id, FollowType, ob)
	o := Follow(*a)
	return &o
}

// IgnoreNew initializes an Ignore activity
func IgnoreNew(id ID, ob Item) *Ignore {
	a := ActivityNew(id, IgnoreType, ob)
	o := Ignore(*a)
	return &o
}

// InviteNew initializes an Invite activity
func InviteNew(id ID, ob Item) *Invite {
	a := ActivityNew(id, InviteType, ob)
	o := Invite(*a)
	return &o
}

// JoinNew initializes a Join activity
func JoinNew(id ID, ob Item) *Join {
	a := ActivityNew(id, JoinType, ob)
	o := Join(*a)
	return &o
}

// LeaveNew initializes a Leave activity
func LeaveNew(id ID, ob Item) *Leave {
	a := ActivityNew(id, LeaveType, ob)
	o := Leave(*a)
	return &o
}

// LikeNew initializes a Like activity
func LikeNew(id ID, ob Item) *Like {
	a := ActivityNew(id, LikeType, ob)
	o := Like(*a)
	return &o
}

// ListenNew initializes a Listen activity
func ListenNew(id ID, ob Item) *Listen {
	a := ActivityNew(id, ListenType, ob)
	o := Listen(*a)
	return &o
}

// MoveNew initializes a Move activity
func MoveNew(id ID, ob Item) *Move {
	a := ActivityNew(id, MoveType, ob)
	o := Move(*a)
	return &o
}

// OfferNew initializes an Offer activity
func OfferNew(id ID, ob Item) *Offer {
	a := ActivityNew(id, OfferType, ob)
	o := Offer(*a)
	return &o
}

// RejectNew initializes a Reject activity
func RejectNew(id ID, ob Item) *Reject {
	a := ActivityNew(id, RejectType, ob)
	o := Reject(*a)
	return &o
}

// ReadNew initializes a Read activity
func ReadNew(id ID, ob Item) *Read {
	a := ActivityNew(id, ReadType, ob)
	o := Read(*a)
	return &o
}

// RemoveNew initializes a Remove activity
func RemoveNew(id ID, ob Item, trgt Item) *Remove {
	a := ActivityNew(id, RemoveType, ob)
	o := Remove(*a)
	o.Target = trgt
	return &o
}

// TentativeRejectNew initializes a TentativeReject activity
func TentativeRejectNew(id ID, ob Item) *TentativeReject {
	a := ActivityNew(id, TentativeRejectType, ob)
	o := TentativeReject(*a)
	return &o
}

// TentativeAcceptNew initializes a TentativeAccept activity
func TentativeAcceptNew(id ID, ob Item) *TentativeAccept {
	a := ActivityNew(id, TentativeAcceptType, ob)
	o := TentativeAccept(*a)
	return &o
}

// UndoNew initializes an Undo activity
func UndoNew(id ID, ob Item) *Undo {
	a := ActivityNew(id, UndoType, ob)
	o := Undo(*a)
	return &o
}

// UpdateNew initializes an Update activity
func UpdateNew(id ID, ob Item) *Update {
	a := ActivityNew(id, UpdateType, ob)
	u := Update(*a)
	return &u
}

// ViewNew initializes a View activity
func ViewNew(id ID, ob Item) *View {
	a := ActivityNew(id, ViewType, ob)
	o := View(*a)
	return &o
}

// ActivityNew initializes a basic activity
func ActivityNew(id ID, typ ActivityVocabularyType, ob Item) *Activity {
	if !ActivityTypes.Contains(typ) {
		typ = ActivityType
	}
	a := Activity{ID: id, Type: typ}
	a.Name = NaturalLanguageValuesNew()
	a.Content = NaturalLanguageValuesNew()

	a.Object = ob

	return &a
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (a *Activity) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return loadActivity(val, a)
}

// ToActivity
func ToActivity(it Item) (*Activity, error) {
	switch i := it.(type) {
	case *Activity:
		return i, nil
	case Activity:
		return &i, nil
	case *IntransitiveActivity:
		return (*Activity)(unsafe.Pointer(i)), nil
	case IntransitiveActivity:
		return (*Activity)(unsafe.Pointer(&i)), nil
	case *Question:
		return (*Activity)(unsafe.Pointer(i)), nil
	case Question:
		return (*Activity)(unsafe.Pointer(&i)), nil
	default:
		// NOTE(marius): this is an ugly way of dealing with the interface conversion error: types from different scopes
		typ := reflect.TypeOf(new(Activity))
		if reflect.TypeOf(it).ConvertibleTo(typ) {
			if i, ok := reflect.ValueOf(it).Convert(typ).Interface().(*Activity); ok {
				return i, nil
			}
		}
	}
	return nil, errors.New("unable to convert activity")
}

// MarshalJSON encodes the receiver object to a JSON document.
func (a Activity) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	write(&b, '{')

	if !writeActivityJSONValue(&b, a) {
		return nil, nil
	}
	write(&b, '}')
	return b, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (a *Activity) UnmarshalBinary(data []byte) error {
	return a.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (a Activity) MarshalBinary() ([]byte, error) {
	return a.GobEncode()
}

func mapIntransitiveActivityProperties(mm map[string][]byte, a *IntransitiveActivity) (hasData bool, err error) {
	err = OnObject(a, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if a.Actor != nil {
		if mm["actor"], err = gobEncodeItem(a.Actor); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Target != nil {
		if mm["target"], err = gobEncodeItem(a.Target); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Result != nil {
		if mm["result"], err = gobEncodeItem(a.Result); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Instrument != nil {
		if mm["instrument"], err = gobEncodeItem(a.Instrument); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return hasData, err
}

func mapActivityProperties(mm map[string][]byte, a *Activity) (hasData bool, err error) {
	err = OnIntransitiveActivity(a, func(a *IntransitiveActivity) error {
		hasData, err = mapIntransitiveActivityProperties(mm, a)
		return err
	})
	if a.Object != nil {
		if mm["object"], err = gobEncodeItem(a.Object); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return hasData, err
}

// GobEncode
func (a Activity) GobEncode() ([]byte, error) {
	var mm = make(map[string][]byte)
	hasData, err := mapActivityProperties(mm, &a)
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
func (a *Activity) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapActivityProperties(mm, a)
}

// Equals verifies if our receiver Object is equals with the "with" Object
func (a Activity) Equals(with Item) bool {
	result := true
	err := OnActivity(with, func(w *Activity) error {
		OnObject(a, func(oa *Object) error {
			result = oa.Equals(w)
			return nil
		})
		if w.Object != nil {
			if !ItemsEqual(a.Object, w.Object) {
				result = false
				return nil
			}
		}
		if w.Actor != nil {
			if !ItemsEqual(a.Actor, w.Actor) {
				result = false
				return nil
			}
		}
		if w.Target != nil {
			if !ItemsEqual(a.Target, w.Target) {
				result = false
				return nil
			}
		}
		if w.Result != nil {
			if !ItemsEqual(a.Result, w.Result) {
				result = false
				return nil
			}
		}
		if w.Origin != nil {
			if !ItemsEqual(a.Origin, w.Origin) {
				result = false
				return nil
			}
		}
		if w.Result != nil {
			if !ItemsEqual(a.Result, w.Result) {
				result = false
				return nil
			}
		}
		if w.Instrument != nil {
			if !ItemsEqual(a.Instrument, w.Instrument) {
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
