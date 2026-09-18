package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"time"

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

// ContentManagementActivityTypes use case primarily deals with activities that involve the creation, modification or
// deletion of content.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-crud
//
// This includes, for instance, activities such as "John created a new note", "Sally updated an article", and
// "Joe deleted the photo".
var ContentManagementActivityTypes = ActivityVocabularyTypes{
	CreateType,
	DeleteType,
	UpdateType,
}

// CollectionManagementActivityTypes use case primarily deals with activities involving the management of content within
// collections.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-collection
//
// Examples of collections include things like folders, albums, friend lists, etc.
// This includes, for instance, activities such as "Sally added a file to Folder A", "John moved the file from Folder A
// to Folder B", etc.
var CollectionManagementActivityTypes = ActivityVocabularyTypes{
	AddType,
	MoveType,
	RemoveType,
}

// ReactionsActivityTypes use case primarily deals with reactions to content.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-reactions
//
// This can include activities such as liking or disliking content, ignoring updates, flagging content as being
// inappropriate, accepting or rejecting objects, etc.
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
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-rsvp
var EventRSVPActivityTypes = ActivityVocabularyTypes{
	AcceptType,
	IgnoreType,
	InviteType,
	RejectType,
	TentativeAcceptType,
	TentativeRejectType,
}

// GroupManagementActivityTypes use case primarily deals with management of groups.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-group
//
// It can include, for instance, activities such as "John added Sally to Group A", "Sally joined Group A",
// "Joe left Group A", etc.
var GroupManagementActivityTypes = ActivityVocabularyTypes{
	AddType,
	JoinType,
	LeaveType,
	RemoveType,
}

// ContentExperienceActivityTypes use case primarily deals with describing activities involving listening to, reading,
// or viewing content.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-experience
//
// For instance, "Sally read the article", "Joe listened to the song".
var ContentExperienceActivityTypes = ActivityVocabularyTypes{
	ListenType,
	ReadType,
	ViewType,
}

// GeoSocialEventsActivityTypes use case primarily deals with activities involving geo-tagging type activities.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-geo
//
// For instance, it can include activities such as "Joe arrived at work", "Sally left work", and
// "John is travel from home to work".
var GeoSocialEventsActivityTypes = ActivityVocabularyTypes{
	ArriveType,
	LeaveType,
	TravelType,
}

// NotificationActivityTypes use case primarily deals with calling attention to particular objects or notifications.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-notification
var NotificationActivityTypes = ActivityVocabularyTypes{
	AnnounceType,
}

// QuestionActivityTypes use case primarily deals with representing inquiries of any type.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-questions
//
// See 5.4 Representing Questions for more information.
// https://www.w3.org/TR/activitystreams-vocabulary/#questions
var QuestionActivityTypes = ActivityVocabularyTypes{
	QuestionType,
}

// RelationshipManagementActivityTypes use case primarily deals with representing activities involving the management of
// interpersonal and social relationships
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-relationships
//
// (e.g. friend requests, management of social network, etc). See 5.2 Representing Relationships Between Entities
// for more information.
// https://www.w3.org/TR/activitystreams-vocabulary/#connections
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
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-undo
//
// See 5.5 Inverse Activities and "Undo" for more information.
// https://www.w3.org/TR/activitystreams-vocabulary/#inverse
var NegatingActivityTypes = ActivityVocabularyTypes{
	UndoType,
}

// OffersActivityTypes use case deals with activities involving offering one object to another.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#motivations-offers
//
// It can include, for instance, activities such as "Company A is offering a discount on purchase of Product Z to Sally",
// "Sally is offering to add a File to Folder A", etc.
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

// HasRecipients is an interface implemented by objects to return their audience
// for further propagation.
//
// Please take care to the fact that the de-duplication functionality requires a pointer receiver
// therefore a valid Item interface that wraps around an Object struct, can not be type asserted
// to HasRecipients.
type HasRecipients interface {
	// Recipients is a method that should do a recipients de-duplication step and then return
	// the remaining recipients.
	Recipients() ItemCollection
}

type Activities interface {
	Activity
}

// Activity is a subtype of Object that describes some form of action that may happen,
// is currently happening, or has already happened.
// The Activity type itself serves as an abstract base type for all types of activities.
// It is important to note that the Activity type itself does not carry any specific semantics
// about the kind of action being taken.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-activity
//
// Activity objects are specializations of the base Object type that provide information about
// actions that have either already occurred, are in the process of occurring, or may occur in the future.
//
// In addition to common properties supported by all Object instances, Activity objects support the following
// additional properties defined by the Vocabulary: actor | object | target | origin | result | instrument
//
// The type property is used to identify the type of action the Activity Statement represents.
//
// https://www.w3.org/TR/activitystreams-core/#activities
type Activity struct {
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
	// Likes identifies the collection containing a list of all Like activities with this object as the object property,
	// added as a side effect.
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
	// Object describes the direct object of the activity.
	// For instance, in the activity "John added a movie to his wishlist", the object of the activity is the movie added.
	Object Item `jsonld:"object,omitempty"`
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Activity) GetType() Typer {
	return a.Type
}

// GetID returns the ID corresponding to the Activity object
func (a Activity) GetID() ID {
	return a.ID
}

// GetLink returns the IRI corresponding to the Activity object
func (a Activity) GetLink() IRI {
	return IRI(a.ID)
}

// Match returns whether the receiver matches the ActivityVocabularyType arguments.
func (a Activity) Match(tt ...ActivityVocabularyType) bool {
	return ActivityVocabularyTypes(tt).Match(a.Type)
}

func removeFromCollection(col ItemCollection, items ...Item) ItemCollection {
	result := make(ItemCollection, 0)
	if len(items) == 0 {
		return col
	}
	for _, ob := range col {
		if !ItemCollection(items).Contains(ob) {
			result = append(result, ob)
		}
	}
	return result
}

func removeFromAllRecipients(a *Activity, items ...Item) error {
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
	toRemove := make(ItemCollection, 0)
	if BlockType.Match(a.GetType()) && !IsNil(a.Object) {
		_ = OnItem(a.Object, func(object Item) error {
			_ = toRemove.Append(object)
			return nil
		})
	}
	if len(toRemove) > 0 {
		_ = removeFromAllRecipients(a, toRemove...)
	}
	aud := a.Audience
	return ItemCollectionDeduplication(&a.To, &a.CC, &a.Bto, &a.BCC, &aud)
}

type Cleanable interface {
	// Clean is a method that removes BCC/Bto recipients in preparation for public consumption of
	// the Object.
	Clean() Item
}

// CleanRecipients checks if the "it" Item has recipients and cleans them if it does
func CleanRecipients(it Item) Item {
	if s, ok := it.(Cleanable); ok {
		it = s.Clean()
	}
	return it
}

func cleanActivityProperties(aa *Activity) {
	_ = OnIntransitiveActivity(aa, func(i *IntransitiveActivity) error {
		cleanIntransitiveActivityProperties(i)
		return nil
	})
	aa.Object = CleanRecipients(aa.Object)
}

// Clean removes Bto and BCC properties
func (a *Activity) Clean() Item {
	aa := *a
	cleanActivityProperties(&aa)
	return aa
}

type (
	// Accept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate
	// the context into which the object has been accepted.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-accept
	Accept = Activity

	// Add indicates that the actor has added the object to the target. If the target property is not explicitly specified,
	// the target would need to be determined implicitly by context.
	// The origin can be used to identify the context from which the object originated.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-add
	Add = Activity

	// Announce indicates that the actor is calling the target's attention the object.
	// The origin typically has no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-announce
	Announce = Activity

	// Block indicates that the actor is blocking the object. Blocking is a stronger form of Ignore.
	// The typical use is to support social systems that allow one user to block activities or content of other users.
	// The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-block
	Block = Ignore

	// Create indicates that the actor has created the object.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-create
	Create = Activity

	// Delete indicates that the actor has deleted the object.
	// If specified, the origin indicates the context from which the object was deleted.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-delete
	Delete = Activity

	// Dislike indicates that the actor dislikes the object.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-dislike
	Dislike = Activity

	// Flag indicates that the actor is "flagging" the object.
	// Flagging is defined in the sense common to many social platforms as reporting content as being
	// inappropriate for any number of reasons.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-flag
	Flag = Activity

	// Follow indicates that the actor is "following" the object. Following is defined in the sense typically used within
	// Social systems in which the actor is interested in any activity performed by or on the object.
	// The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-follow
	Follow = Activity

	// Ignore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-ignore
	Ignore = Activity

	// Invite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-invite
	Invite = Offer

	// Join indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-join
	Join = Activity

	// Leave indicates that the actor has left the object. The target and origin typically have no meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-leave
	Leave = Activity

	// Like indicates that the actor likes, recommends or endorses the object.
	// The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-like
	Like = Activity

	// Listen indicates that the actor has listened to the object.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-listen
	Listen = Activity

	// Move indicates that the actor has moved object from origin to target.
	// If the origin or target are not specified, either can be determined by context.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-move
	Move = Activity

	// Offer indicates that the actor is offering the object.
	// If specified, the target indicates the entity to which the object is being offered.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-offer
	Offer = Activity

	// Reject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-reject
	Reject = Activity

	// Read indicates that the actor has read the object.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-read
	Read = Activity

	// Remove indicates that the actor is removing the object. If specified,
	// the origin indicates the context from which the object is being removed.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-remove
	Remove = Activity

	// TentativeReject is a specialization of Reject in which the rejection is considered tentative.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-tentativereject
	TentativeReject = Reject

	// TentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-tentativeaccept
	TentativeAccept = Accept

	// Undo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing
	// some previously performed action (for instance, a person may have previously "liked" an article but,
	// for whatever reason, might choose to undo that like at some later point in time).
	// The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-undo
	Undo = Activity

	// Update indicates that the actor has updated the object.
	// Please note, however, that this vocabulary does not define a mechanism
	// for describing the actual set of modifications made to object.
	// The target and origin typically have no defined meaning.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-update
	Update = Activity

	// View indicates that the actor has viewed the object.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-view
	View = Activity
)

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (a *Activity) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadActivity(val, a)
}

func fmtActivityProps(w io.Writer, n *int) func(*Activity) error {
	return func(a *Activity) error {
		_ = OnIntransitiveActivity(a, fmtIntransitiveActivityProps(w, n))

		comma := func() {
			if *n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}

		if !IsNil(a.Object) {
			comma()
			*n, _ = fmt.Fprintf(w, "object: %s", a.Object)
		}
		return nil
	}
}

func (a Activity) Format(s fmt.State, verb rune) {
	typ := a.Type
	switch verb {
	case 's':
		iri := a.ID
		if iri != "" {
			s.Write([]byte(iri))
		} else {
			_, _ = fmt.Fprintf(s, "%T[%v]", a, typ)
		}
	case 'v':
		n := 0
		if typ != nil {
			_, _ = fmt.Fprintf(s, "%T[%v] { ", a, typ)
			_ = fmtActivityProps(s, &n)(&a)
			_, _ = io.WriteString(s, " }")
		} else {
			_, _ = fmt.Fprintf(s, "%T { ", a)
			_ = fmtActivityProps(s, &n)(&a)
			_, _ = io.WriteString(s, " }")
		}
	}
}

// ToActivity
func ToActivity(it LinkOrIRI) (*Activity, error) {
	switch i := it.(type) {
	case *Activity:
		return i, nil
	case Activity:
		return &i, nil
	default:
		return reflectItemToType[Activity](it)
	}
}

// MarshalJSON encodes the receiver object to a JSON document.
func (a Activity) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	JSONWrite(&b, '{')

	if !JSONWriteActivityValue(&b, a) {
		return nil, nil
	}
	JSONWrite(&b, '}')
	return b.Bytes(), nil
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
	mm := make(map[string][]byte)
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

// Equals verifies if our receiver Activity is equals with the "with" Item
func (a Activity) Equals(with Item) bool {
	withActivity, err := ToActivity(with)
	if err != nil {
		return false
	}
	return a.equal(*withActivity)
}

// equal verifies if our receiver Activity is equals with the "with" Activity
func (a *Activity) equal(with Activity) bool {
	result := true
	_ = OnIntransitiveActivity(a, func(oi *IntransitiveActivity) error {
		result = oi.Equals(with)
		return nil
	})
	if !result {
		return false
	}
	if !ItemsEqual(a.Object, with.Object) {
		return false
	}
	return true
}

// WithActivityFn represents a function type that can be used as a parameter for OnActivity helper function
type WithActivityFn func(*Activity) error

// OnActivity calls function fn on it Item if it can be asserted to type *Activity
//
// This function should be called if trying to access the Activity specific properties
// like "object", for the other properties OnObject, or OnIntransitiveActivity
// should be used instead.
func OnActivity(it LinkOrIRI, fn WithActivityFn) error {
	if IsNil(it) {
		return nil
	}

	actFn := func(it LinkOrIRI) error {
		act, err := ToActivity(it)
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

func notEmptyActivity(a *Activity) bool {
	var notEmpty bool
	_ = OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		notEmpty = notEmptyIntransitiveActivity(i)
		return nil
	})
	return notEmpty || a.Object != nil
}
