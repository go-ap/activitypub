package activitystreams

import (
	"errors"
	"time"
	"unsafe"
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

type ActivityVocabularyTypes []ActivityVocabularyType

func (a ActivityVocabularyTypes) Contains(typ ActivityVocabularyType) bool {
	for _, v := range a {
		if v == typ {
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

// Activity is a subtype of Object that describes some form of action that may happen,
// is currently happening, or has already happened.
// The Activity type itself serves as an abstract base type for all types of activities.
// It is important to note that the Activity type itself does not carry any specific semantics
// about the kind of action being taken.
type Activity struct {
	Parent
	// Actor describes one or more entities that either performed or are expected to perform the activity.
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

// IntransitiveActivity Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions.
// The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
	Parent
	// Actor describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor Actor `jsonld:"actor,omitempty"`
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

	// Arrive is an IntransitiveActivity that indicates that the actor has arrived at the location.
	// The origin can be used to identify the context from which the actor originated.
	// The target typically has no defined meaning.
	Arrive = IntransitiveActivity

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

	// Travel indicates that the actor is traveling to target from origin.
	// Travel is an IntransitiveObject whose actor specifies the direct object.
	// If the target or origin are not specified, either can be determined by context.
	Travel = IntransitiveActivity

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

// Question represents a question being asked. Question objects are an extension of IntransitiveActivity.
// That is, the Question object is an Activity, but the direct object is the question
// itself and therefore it would not contain an object property.
// Either of the anyOf and oneOf properties may be used to express possible answers,
// but a Question object must not have both properties.
type Question struct {
	// ID providesthe globally unique identifier for an Activity Pub Object or Link.
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
	Audience Item `jsonld:"audience,omitempty"`
	// Content the content or textual representation of the Activity Pub Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValues `jsonld:"content,omitempty,collapsible"`
	// Context identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	// common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
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
	// Preview identifies an entity that providesa preview of this object.
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
	// Tag One or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	// while the latter implies associated by reference.
	Tag Item `jsonld:"tag,omitempty"`
	// Updated the date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// URL identifies one or more links to representations of the object
	URL LinkOrIRI `jsonld:"url,omitempty"`
	// To identifies an entity considered to be part of the public primary audience of an Activity Pub Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Bto identifies an Activity Pub Object that is part of the private primary audience of this Activity Pub Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// CC identifies an Activity Pub Object that is part of the public secondary audience of this Activity Pub Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// BCC identifies one or more Objects that are part of the private secondary audience of this Activity Pub Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// Duration When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	// the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	// section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// Actor describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor Actor `jsonld:"actor,omitempty"`
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
	// OneOf identifies an exclusive option for a Question. Use of oneOf implies that the Question
	// can have only a single answer. To indicate that a Question can have multiple answers, use anyOf.
	OneOf Item `jsonld:"oneOf,omitempty"`
	// AnyOf identifies an inclusive option for a Question. Use of anyOf implies that the Question can have multiple answers.
	// To indicate that a Question can have only one answer, use oneOf.
	AnyOf Item `jsonld:"anyOf,omitempty"`
	// Closed indicates that a question has been closed, and answers are no longer accepted.
	Closed bool `jsonld:"closed,omitempty"`
}

// AcceptNew initializes an Accept activity
func AcceptNew(id ObjectID, ob Item) *Accept {
	a := ActivityNew(id, AcceptType, ob)
	o := Accept(*a)
	return &o
}

// AddNew initializes an Add activity
func AddNew(id ObjectID, ob Item, trgt Item) *Add {
	a := ActivityNew(id, AddType, ob)
	o := Add(*a)
	o.Target = trgt
	return &o
}

// AnnounceNew initializes an Announce activity
func AnnounceNew(id ObjectID, ob Item) *Announce {
	a := ActivityNew(id, AnnounceType, ob)
	o := Announce(*a)
	return &o
}

// ArriveNew initializes an Arrive activity
func ArriveNew(id ObjectID) *Arrive {
	a := IntransitiveActivityNew(id, ArriveType)
	o := Arrive(*a)
	return &o
}

// BlockNew initializes a Block activity
func BlockNew(id ObjectID, ob Item) *Block {
	a := ActivityNew(id, BlockType, ob)
	o := Block(*a)
	return &o
}

// CreateNew initializes a Create activity
func CreateNew(id ObjectID, ob Item) *Create {
	a := ActivityNew(id, CreateType, ob)
	o := Create(*a)
	return &o
}

// DeleteNew initializes a Delete activity
func DeleteNew(id ObjectID, ob Item) *Delete {
	a := ActivityNew(id, DeleteType, ob)
	o := Delete(*a)
	return &o
}

// DislikeNew initializes a Dislike activity
func DislikeNew(id ObjectID, ob Item) *Dislike {
	a := ActivityNew(id, DislikeType, ob)
	o := Dislike(*a)
	return &o
}

// FlagNew initializes a Flag activity
func FlagNew(id ObjectID, ob Item) *Flag {
	a := ActivityNew(id, FlagType, ob)
	o := Flag(*a)
	return &o
}

// FollowNew initializes a Follow activity
func FollowNew(id ObjectID, ob Item) *Follow {
	a := ActivityNew(id, FollowType, ob)
	o := Follow(*a)
	return &o
}

// IgnoreNew initializes an Ignore activity
func IgnoreNew(id ObjectID, ob Item) *Ignore {
	a := ActivityNew(id, IgnoreType, ob)
	o := Ignore(*a)
	return &o
}

// InviteNew initializes an Invite activity
func InviteNew(id ObjectID, ob Item) *Invite {
	a := ActivityNew(id, InviteType, ob)
	o := Invite(*a)
	return &o
}

// JoinNew initializes a Join activity
func JoinNew(id ObjectID, ob Item) *Join {
	a := ActivityNew(id, JoinType, ob)
	o := Join(*a)
	return &o
}

// LeaveNew initializes a Leave activity
func LeaveNew(id ObjectID, ob Item) *Leave {
	a := ActivityNew(id, LeaveType, ob)
	o := Leave(*a)
	return &o
}

// LikeNew initializes a Like activity
func LikeNew(id ObjectID, ob Item) *Like {
	a := ActivityNew(id, LikeType, ob)
	o := Like(*a)
	return &o
}

// ListenNew initializes a Listen activity
func ListenNew(id ObjectID, ob Item) *Listen {
	a := ActivityNew(id, ListenType, ob)
	o := Listen(*a)
	return &o
}

// MoveNew initializes a Move activity
func MoveNew(id ObjectID, ob Item) *Move {
	a := ActivityNew(id, MoveType, ob)
	o := Move(*a)
	return &o
}

// OfferNew initializes an Offer activity
func OfferNew(id ObjectID, ob Item) *Offer {
	a := ActivityNew(id, OfferType, ob)
	o := Offer(*a)
	return &o
}

// RejectNew initializes a Reject activity
func RejectNew(id ObjectID, ob Item) *Reject {
	a := ActivityNew(id, RejectType, ob)
	o := Reject(*a)
	return &o
}

// ReadNew initializes a Read activity
func ReadNew(id ObjectID, ob Item) *Read {
	a := ActivityNew(id, ReadType, ob)
	o := Read(*a)
	return &o
}

// RemoveNew initializes a Remove activity
func RemoveNew(id ObjectID, ob Item, trgt Item) *Remove {
	a := ActivityNew(id, RemoveType, ob)
	o := Remove(*a)
	o.Target = trgt
	return &o
}

// TentativeRejectNew initializes a TentativeReject activity
func TentativeRejectNew(id ObjectID, ob Item) *TentativeReject {
	a := ActivityNew(id, TentativeRejectType, ob)
	o := TentativeReject(*a)
	return &o
}

// TentativeAcceptNew initializes a TentativeAccept activity
func TentativeAcceptNew(id ObjectID, ob Item) *TentativeAccept {
	a := ActivityNew(id, TentativeAcceptType, ob)
	o := TentativeAccept(*a)
	return &o
}

// TravelNew initializes a Travel activity
func TravelNew(id ObjectID) *Travel {
	a := IntransitiveActivityNew(id, TravelType)
	o := Travel(*a)
	return &o
}

// UndoNew initializes an Undo activity
func UndoNew(id ObjectID, ob Item) *Undo {
	a := ActivityNew(id, UndoType, ob)
	o := Undo(*a)
	return &o
}

// UpdateNew initializes an Update activity
func UpdateNew(id ObjectID, ob Item) *Update {
	a := ActivityNew(id, UpdateType, ob)
	u := Update(*a)
	return &u
}

// ViewNew initializes a View activity
func ViewNew(id ObjectID, ob Item) *View {
	a := ActivityNew(id, ViewType, ob)
	o := View(*a)
	return &o
}

// QuestionNew initializes a Question activity
func QuestionNew(id ObjectID) *Question {
	q := Question{ID: id, Type: QuestionType}
	q.Name = NaturalLanguageValuesNew()
	q.Content = NaturalLanguageValuesNew()
	return &q
}

// ActivityNew initializes a basic activity
func ActivityNew(id ObjectID, typ ActivityVocabularyType, ob Item) *Activity {
	if !ActivityTypes.Contains(typ) {
		typ = ActivityType
	}
	a := Activity{Parent: Parent{ID: id, Type: typ}}
	a.Name = NaturalLanguageValuesNew()
	a.Content = NaturalLanguageValuesNew()

	a.Object = ob

	return &a
}

// IntransitiveActivityNew initializes a intransitive activity
func IntransitiveActivityNew(id ObjectID, typ ActivityVocabularyType) *IntransitiveActivity {
	if !IntransitiveActivityTypes.Contains(typ) {
		typ = IntransitiveActivityType
	}
	i := IntransitiveActivity{Parent: Parent{ID: id, Type: typ}}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	return &i
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Activity) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink returns false for Activity objects
func (a Activity) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Activity object
func (a Activity) GetID() *ObjectID {
	return &a.ID
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

// Recipients performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (a *Activity) Recipients() ItemCollection {
	actor := make(ItemCollection, 0)
	actor.Append(a.Actor)
	rec, _ := ItemCollectionDeduplication(&actor, &a.To, &a.Bto, &a.CC, &a.BCC, &a.Audience)
	return rec
}

// Clean removes Bto and BCC properties
func (a *Activity) Clean() {
	a.BCC = nil
	a.Bto = nil
}

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

// GetID returns the ObjectID corresponding to the IntransitiveActivity object
func (i IntransitiveActivity) GetID() *ObjectID {
	return &i.ID
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

// GetID returns the ObjectID corresponding to the Question object
func (q Question) GetID() *ObjectID {
	return &q.ID
}

// GetLink returns the IRI corresponding to the Question object
func (q Question) GetLink() IRI {
	return IRI(q.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (q Question) GetType() ActivityVocabularyType {
	return q.Type
}

// IsObject returns true for Question objects
func (q Question) IsObject() bool {
	return true
}

// IsLink returns false for Question objects
func (q Question) IsLink() bool {
	return false
}

// IsCollection returns false for Question objects
func (q Question) IsCollection() bool {
	return false
}

// UnmarshalJSON
func (a *Activity) UnmarshalJSON(data []byte) error {
	a.Parent.UnmarshalJSON(data)
	a.Actor = JSONGetItem(data, "actor")
	a.Object = JSONGetItem(data, "object")
	a.Target = JSONGetItem(data, "target")
	a.Instrument = JSONGetItem(data, "instrument")
	a.Origin = JSONGetItem(data, "origin")
	a.Result = JSONGetItem(data, "result")
	return nil
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
	}
	return nil, errors.New("unable to convert activity")
}

// ToQuestion
func ToQuestion(it Item) (*Question, error) {
	switch i := it.(type) {
	case *Question:
		return i, nil
	case Question:
		return &i, nil
	}
	return nil, errors.New("unable to convert to question activity")
}

// ToIntransitiveActivity
func ToIntransitiveActivity(it Item) (*IntransitiveActivity, error) {
	switch i := it.(type) {
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

// FlattenActivityProperties flattens the Activity's properties from Object type to IRI
func FlattenActivityProperties(act *Activity) *Activity {
	act.Object = Flatten(act.Object)
	act.Actor = Flatten(act.Actor)
	act.Target = Flatten(act.Target)
	act.Result = Flatten(act.Result)
	act.Origin = Flatten(act.Origin)
	act.Result = Flatten(act.Result)
	act.Instrument = Flatten(act.Instrument)
	return act
}
