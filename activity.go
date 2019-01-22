package activitystreams

import (
	"time"
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

var validContentManagementActivityTypes = [...]ActivityVocabularyType{
	CreateType,
	DeleteType,
	UpdateType,
}

var validCollectionManagementActivityTypes = [...]ActivityVocabularyType{
	AddType,
	MoveType,
	RemoveType,
}

var validReactionsActivityTypes = [...]ActivityVocabularyType{
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

var validEventRSVPActivityTypes = [...]ActivityVocabularyType{
	AcceptType,
	IgnoreType,
	InviteType,
	RejectType,
	TentativeAcceptType,
	TentativeRejectType,
}

var validGroupManagementActivityTypes = [...]ActivityVocabularyType{
	AddType,
	JoinType,
	LeaveType,
	RemoveType,
}

var validContentExperienceActivityTypes = [...]ActivityVocabularyType{
	ArriveType,
	LeaveType,
	TravelType,
}

var validGeoSocialEventsActivityTypes = [...]ActivityVocabularyType{
	ArriveType,
	LeaveType,
	TravelType,
}

var validNotificationActivityTypes = [...]ActivityVocabularyType{
	AnnounceType,
}

var validQuestionActivityTypes = [...]ActivityVocabularyType{
	QuestionType,
}

var validRelationshipManagementActivityTypes = [...]ActivityVocabularyType{
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

var validNegatingActivityTypes = [...]ActivityVocabularyType{
	UndoType,
}

var validOffersActivityTypes = [...]ActivityVocabularyType{
	OfferType,
}
var validIntransitiveActivityTypes = [...]ActivityVocabularyType{
	ArriveType,
	TravelType,
	QuestionType,
}
var validActivityTypes = [...]ActivityVocabularyType{
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
	// Actor Types
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
	Accept Activity

	// Add indicates that the actor has added the object to the target. If the target property is not explicitly specified,
	// the target would need to be determined implicitly by context.
	// The origin can be used to identify the context from which the object originated.
	Add Activity

	// Announce indicates that the actor is calling the target's attention the object.
	// The origin typically has no defined meaning.
	Announce Activity

	// Arrive is an IntransitiveActivity that indicates that the actor has arrived at the location.
	// The origin can be used to identify the context from which the actor originated.
	// The target typically has no defined meaning.
	Arrive IntransitiveActivity

	// Block indicates that the actor is blocking the object. Blocking is a stronger form of Ignore.
	// The typical use is to support social systems that allow one user to block activities or content of other users.
	// The target and origin typically have no defined meaning.
	Block Ignore

	// Create indicates that the actor has created the object.
	Create Activity

	// Delete indicates that the actor has deleted the object.
	// If specified, the origin indicates the context from which the object was deleted.
	Delete Activity

	// Dislike indicates that the actor dislikes the object.
	Dislike Activity

	// Flag indicates that the actor is "flagging" the object.
	// Flagging is defined in the sense common to many social platforms as reporting content as being
	// inappropriate for any number of reasons.
	Flag Activity

	// Follow indicates that the actor is "following" the object. Following is defined in the sense typically used within
	// Social systems in which the actor is interested in any activity performed by or on the object.
	// The target and origin typically have no defined meaning.
	Follow Activity

	// Ignore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	Ignore Activity

	// Invite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	Invite Offer

	// Join indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	Join Activity

	// Leave indicates that the actor has left the object. The target and origin typically have no meaning.
	Leave Activity

	// Like indicates that the actor likes, recommends or endorses the object.
	// The target and origin typically have no defined meaning.
	Like Activity

	// Listen inherits all properties from Activity.
	Listen Activity

	// Move indicates that the actor has moved object from origin to target.
	// If the origin or target are not specified, either can be determined by context.
	Move Activity

	// Offer indicates that the actor is offering the object.
	// If specified, the target indicates the entity to which the object is being offered.
	Offer Activity

	// Reject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	Reject Activity

	// Read indicates that the actor has read the object.
	Read Activity

	// Remove indicates that the actor is removing the object. If specified,
	// the origin indicates the context from which the object is being removed.
	Remove Activity

	// TentativeReject is a specialization of Reject in which the rejection is considered tentative.
	TentativeReject Reject

	// TentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	TentativeAccept Accept

	// Travel indicates that the actor is traveling to target from origin.
	// Travel is an IntransitiveObject whose actor specifies the direct object.
	// If the target or origin are not specified, either can be determined by context.
	Travel IntransitiveActivity

	// Undo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing
	// some previously performed action (for instance, a person may have previously "liked" an article but,
	// for whatever reason, might choose to undo that like at some later point in time).
	// The target and origin typically have no defined meaning.
	Undo Activity

	// Update indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism
	// for describing the actual set of modifications made to object.
	// The target and origin typically have no defined meaning.
	Update Activity

	// View indicates that the actor has viewed the object.
	View Activity
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
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
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
	Content NaturalLanguageValue `jsonld:"content,omitempty,collapsible"`
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
	Summary NaturalLanguageValue `jsonld:"summary,omitempty,collapsible"`
	// Tag One or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	// while the latter implies associated by reference.
	Tag Item `jsonld:"tag,omitempty"`
	// Updated the date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// URL identifies one or more links to representations of the object
	URL LinkOrURI `jsonld:"url,omitempty"`
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
	q.Name = NaturalLanguageValueNew()
	q.Content = NaturalLanguageValueNew()
	return &q
}

// ValidContentManagementType is a validation function for content management Activity objects
func ValidContentManagementType(typ ActivityVocabularyType) bool {
	for _, v := range validContentManagementActivityTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ValidCollectionManagementType is a validation function for collection management Activity objects
func ValidCollectionManagementType(typ ActivityVocabularyType) bool {
	for _, v := range validCollectionManagementActivityTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ValidReactionsType is a validation function for reactions Activity objects
func ValidReactionsType(typ ActivityVocabularyType) bool {
	for _, v := range validReactionsActivityTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ValidActivityType is a validation function for Activity objects
func ValidActivityType(typ ActivityVocabularyType) bool {
	for _, v := range validActivityTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ValidIntransitiveActivityType is a validation function for IntransitiveActivity objects
func ValidIntransitiveActivityType(typ ActivityVocabularyType) bool {
	for _, v := range validIntransitiveActivityTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ActivityNew initializes a basic activity
func ActivityNew(id ObjectID, typ ActivityVocabularyType, ob Item) *Activity {
	if !ValidActivityType(typ) {
		typ = ActivityType
	}
	a := Activity{Parent: Parent{ID: id, Type: typ}}
	a.Name = NaturalLanguageValueNew()
	a.Content = NaturalLanguageValueNew()

	a.Object = ob

	return &a
}

// IntransitiveActivityNew initializes a intransitive activity
func IntransitiveActivityNew(id ObjectID, typ ActivityVocabularyType) *IntransitiveActivity {
	if !ValidIntransitiveActivityType(typ) {
		typ = IntransitiveActivityType
	}
	i := IntransitiveActivity{Parent: Parent{ID: id, Type: typ}}
	i.Name = NaturalLanguageValueNew()
	i.Content = NaturalLanguageValueNew()

	return &i
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (a *Activity) RecipientsDeduplication() {
	var actor ItemCollection
	actor.Append(a.Actor)
	recipientsDeduplication(&actor, &a.To, &a.Bto, &a.CC, &a.BCC)
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (i *IntransitiveActivity) RecipientsDeduplication() {
	var actor ItemCollection
	actor.Append(i.Actor)
	recipientsDeduplication(&actor, &i.To, &i.Bto, &i.CC, &i.BCC)
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (b *Block) RecipientsDeduplication() {
	var dedupObjects ItemCollection
	dedupObjects.Append(b.Actor)
	dedupObjects.Append(b.Object)
	recipientsDeduplication(&dedupObjects, &b.To, &b.Bto, &b.CC, &b.BCC)
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (c *Create) RecipientsDeduplication() {
	var dedupObjects ItemCollection
	dedupObjects.Append(c.Actor)
	dedupObjects.Append(c.Object)
	recipientsDeduplication(&dedupObjects, &c.To, &c.Bto, &c.CC, &c.BCC)
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (l *Like) RecipientsDeduplication() {
	var dedupObjects ItemCollection
	dedupObjects.Append(l.Actor)
	dedupObjects.Append(l.Object)
	recipientsDeduplication(&dedupObjects, &l.To, &l.Bto, &l.CC, &l.BCC)
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (d *Dislike) RecipientsDeduplication() {
	var dedupObjects ItemCollection
	dedupObjects.Append(d.Actor)
	dedupObjects.Append(d.Object)
	recipientsDeduplication(&dedupObjects, &d.To, &d.Bto, &d.CC, &d.BCC)
}

// RecipientsDeduplication performs recipient de-duplication on the Activity's To, Bto, CC and BCC properties
func (u *Update) RecipientsDeduplication() {
	var dedupObjects ItemCollection
	dedupObjects.Append(u.Actor)
	dedupObjects.Append(u.Object)
	recipientsDeduplication(&dedupObjects, &u.To, &u.Bto, &u.CC, &u.BCC)
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

// IsObject returns true for Activity objects
func (i IntransitiveActivity) IsObject() bool {
	return true
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

// GetID returns the ObjectID corresponding to the Like object
func (l Like) GetID() *ObjectID {
	return Activity(l).GetID()
}

// GetLink returns the IRI corresponding to the Like object
func (l Like) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (l Like) GetType() ActivityVocabularyType {
	return l.Type
}

// IsObject returns true for Like objects
func (l Like) IsObject() bool {
	return true
}

// IsLink returns false for Like objects
func (l Like) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Dislike object
func (d Dislike) GetID() *ObjectID {
	return Activity(d).GetID()
}

// GetLink returns the IRI corresponding to the Dislike object
func (d Dislike) GetLink() IRI {
	return IRI(d.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (d Dislike) GetType() ActivityVocabularyType {
	return d.Type
}

// IsObject returns true for Dislike objects
func (d Dislike) IsObject() bool {
	return true
}

// IsLink returns false for Dislike objects
func (d Dislike) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Accept object
func (a Accept) GetID() *ObjectID {
	return Activity(a).GetID()
}

// GetLink returns the IRI corresponding to the Accept object
func (a Accept) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Accept) GetType() ActivityVocabularyType {
	return a.Type
}

// IsObject returns true for Accept objects
func (a Accept) IsObject() bool {
	return true
}

// IsLink returns false for Accept objects
func (a Accept) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Add object
func (a Add) GetID() *ObjectID {
	return Activity(a).GetID()
}

// GetLink returns the IRI corresponding to the Add object
func (a Add) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Add) GetType() ActivityVocabularyType {
	return a.Type
}

// IsObject returns true for Add objects
func (a Add) IsObject() bool {
	return true
}

// IsLink returns false for Add objects
func (a Add) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Announce object
func (a Announce) GetID() *ObjectID {
	return Activity(a).GetID()
}

// GetLink returns the IRI corresponding to the Announce object
func (a Announce) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Announce) GetType() ActivityVocabularyType {
	return a.Type
}

// IsObject returns true for Announce objects
func (a Announce) IsObject() bool {
	return true
}

// IsLink returns false for Announce objects
func (a Announce) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Arrive object
func (a Arrive) GetID() *ObjectID {
	return IntransitiveActivity(a).GetID()
}

// GetLink returns the IRI corresponding to the Arrive object
func (a Arrive) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (a Arrive) GetType() ActivityVocabularyType {
	return a.Type
}

// IsObject returns true for Arrive objects
func (a Arrive) IsObject() bool {
	return true
}

// IsLink returns false for Arrive objects
func (a Arrive) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Block object
func (b Block) GetID() *ObjectID {
	return Activity(b).GetID()
}

// GetLink returns the IRI corresponding to the Block object
func (b Block) GetLink() IRI {
	return IRI(b.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (b Block) GetType() ActivityVocabularyType {
	return b.Type
}

// IsObject returns true for Block objects
func (b Block) IsObject() bool {
	return true
}

// IsLink returns false for Block objects
func (b Block) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Create object
func (c Create) GetID() *ObjectID {
	return Activity(c).GetID()
}

// GetLink returns the IRI corresponding to the Create object
func (c Create) GetLink() IRI {
	return IRI(c.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (c Create) GetType() ActivityVocabularyType {
	return c.Type
}

// IsObject returns true for Create objects
func (c Create) IsObject() bool {
	return true
}

// IsLink returns false for Create objects
func (c Create) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Delete object
func (d Delete) GetID() *ObjectID {
	return Activity(d).GetID()
}

// GetLink returns the IRI corresponding to the Delete object
func (d Delete) GetLink() IRI {
	return IRI(d.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (d Delete) GetType() ActivityVocabularyType {
	return d.Type
}

// IsObject returns true for Delete objects
func (d Delete) IsObject() bool {
	return true
}

// IsLink returns false for Delete objects
func (d Delete) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Flag object
func (f Flag) GetID() *ObjectID {
	return Activity(f).GetID()
}

// GetLink returns the IRI corresponding to the Flag object
func (f Flag) GetLink() IRI {
	return IRI(f.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (f Flag) GetType() ActivityVocabularyType {
	return f.Type
}

// IsObject returns true for Flag objects
func (f Flag) IsObject() bool {
	return true
}

// IsLink returns false for Flag objects
func (f Flag) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Follow object
func (f Follow) GetID() *ObjectID {
	return Activity(f).GetID()
}

// GetLink returns the IRI corresponding to the Follow object
func (f Follow) GetLink() IRI {
	return IRI(f.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (f Follow) GetType() ActivityVocabularyType {
	return f.Type
}

// IsObject returns true for Follow objects
func (f Follow) IsObject() bool {
	return true
}

// IsLink returns false for Follow objects
func (f Follow) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Ignore object
func (i Ignore) GetID() *ObjectID {
	return Activity(i).GetID()
}

// GetLink returns the IRI corresponding to the Ignore object
func (i Ignore) GetLink() IRI {
	return IRI(i.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (i Ignore) GetType() ActivityVocabularyType {
	return i.Type
}

// IsObject returns true for Ignore objects
func (i Ignore) IsObject() bool {
	return true
}

// IsLink returns false for Ignore objects
func (i Ignore) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Invite object
func (i Invite) GetID() *ObjectID {
	return Activity(i).GetID()
}

// GetLink returns the IRI corresponding to the Invite object
func (i Invite) GetLink() IRI {
	return IRI(i.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (i Invite) GetType() ActivityVocabularyType {
	return i.Type
}

// IsObject returns true for Invite objects
func (i Invite) IsObject() bool {
	return true
}

// IsLink returns false for Invite objects
func (i Invite) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Join object
func (j Join) GetID() *ObjectID {
	return Activity(j).GetID()
}

// GetLink returns the IRI corresponding to the Join object
func (j Join) GetLink() IRI {
	return IRI(j.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (j Join) GetType() ActivityVocabularyType {
	return j.Type
}

// IsObject returns true for Join objects
func (j Join) IsObject() bool {
	return true
}

// IsLink returns false for Join objects
func (j Join) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Leave object
func (l Leave) GetID() *ObjectID {
	return Activity(l).GetID()
}

// GetLink returns the IRI corresponding to the Leave object
func (l Leave) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (l Leave) GetType() ActivityVocabularyType {
	return l.Type
}

// IsObject returns true for Leave objects
func (l Leave) IsObject() bool {
	return true
}

// IsLink returns false for Leave objects
func (l Leave) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Listen object
func (l Listen) GetID() *ObjectID {
	return Activity(l).GetID()
}

// GetLink returns the IRI corresponding to the Listen object
func (l Listen) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (l Listen) GetType() ActivityVocabularyType {
	return l.Type
}

// IsObject returns true for Listen objects
func (l Listen) IsObject() bool {
	return true
}

// IsLink returns false for Listen objects
func (l Listen) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Move object
func (m Move) GetID() *ObjectID {
	return Activity(m).GetID()
}

// GetLink returns the IRI corresponding to the Move object
func (m Move) GetLink() IRI {
	return IRI(m.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (m Move) GetType() ActivityVocabularyType {
	return m.Type
}

// IsObject returns true for Move objects
func (m Move) IsObject() bool {
	return true
}

// IsLink returns false for Move objects
func (m Move) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Offer object
func (o Offer) GetID() *ObjectID {
	return Activity(o).GetID()
}

// GetLink returns the IRI corresponding to the Offer object
func (o Offer) GetLink() IRI {
	return IRI(o.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (o Offer) GetType() ActivityVocabularyType {
	return o.Type
}

// IsObject returns true for Offer objects
func (o Offer) IsObject() bool {
	return true
}

// IsLink returns false for Offer objects
func (o Offer) IsLink() bool {
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

// GetID returns the ObjectID corresponding to the Reject object
func (r Reject) GetID() *ObjectID {
	return Activity(r).GetID()
}

// GetLink returns the IRI corresponding to the Reject object
func (r Reject) GetLink() IRI {
	return IRI(r.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (r Reject) GetType() ActivityVocabularyType {
	return r.Type
}

// IsObject returns true for Reject objects
func (r Reject) IsObject() bool {
	return true
}

// IsLink returns false for Reject objects
func (r Reject) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Remove object
func (r Remove) GetID() *ObjectID {
	return Activity(r).GetID()
}

// GetLink returns the IRI corresponding to the Remove object
func (r Remove) GetLink() IRI {
	return IRI(r.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (r Remove) GetType() ActivityVocabularyType {
	return r.Type
}

// IsObject returns true for Remove objects
func (r Remove) IsObject() bool {
	return true
}

// IsLink returns false for Remove objects
func (r Remove) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Read object
func (r Read) GetID() *ObjectID {
	return Activity(r).GetID()
}

// GetLink returns the IRI corresponding to the Read object
func (r Read) GetLink() IRI {
	return IRI(r.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (r Read) GetType() ActivityVocabularyType {
	return r.Type
}

// IsObject returns true for Read objects
func (r Read) IsObject() bool {
	return true
}

// IsLink returns false for Read objects
func (r Read) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the TentativeAccept object
func (t TentativeAccept) GetID() *ObjectID {
	return Activity(t).GetID()
}

// GetLink returns the IRI corresponding to the TentativeAccept object
func (t TentativeAccept) GetLink() IRI {
	return IRI(t.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (t TentativeAccept) GetType() ActivityVocabularyType {
	return t.Type
}

// IsObject returns true for TentativeAccept objects
func (t TentativeAccept) IsObject() bool {
	return true
}

// IsLink returns false for TentativeAccept objects
func (t TentativeAccept) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the TentativeReject object
func (t TentativeReject) GetID() *ObjectID {
	return Activity(t).GetID()
}

// GetLink returns the IRI corresponding to the TentativeReject object
func (t TentativeReject) GetLink() IRI {
	return IRI(t.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (t TentativeReject) GetType() ActivityVocabularyType {
	return t.Type
}

// IsObject returns true for TentativeReject objects
func (t TentativeReject) IsObject() bool {
	return true
}

// IsLink returns false for TentativeReject objects
func (t TentativeReject) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Travel object
func (t Travel) GetID() *ObjectID {
	return IntransitiveActivity(t).GetID()
}

// GetLink returns the IRI corresponding to the Travel object
func (t Travel) GetLink() IRI {
	return IRI(t.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (t Travel) GetType() ActivityVocabularyType {
	return t.Type
}

// IsObject returns true for Travel objects
func (t Travel) IsObject() bool {
	return true
}

// IsLink returns false for Travel objects
func (t Travel) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Undo object
func (u Undo) GetID() *ObjectID {
	return Activity(u).GetID()
}

// GetLink returns the IRI corresponding to the Unto object
func (u Undo) GetLink() IRI {
	return IRI(u.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (u Undo) GetType() ActivityVocabularyType {
	return u.Type
}

// IsObject returns true for Undo objects
func (u Undo) IsObject() bool {
	return true
}

// IsLink returns false for Undo objects
func (u Undo) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Update object
func (u Update) GetID() *ObjectID {
	return Activity(u).GetID()
}

// GetLink returns the IRI corresponding to the Update object
func (u Update) GetLink() IRI {
	return IRI(u.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (u Update) GetType() ActivityVocabularyType {
	return u.Type
}

// IsObject returns true for Update objects
func (u Update) IsObject() bool {
	return true
}

// IsLink returns false for Update objects
func (u Update) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the View object
func (v View) GetID() *ObjectID {
	return Activity(v).GetID()
}

// GetLink returns the IRI corresponding to the View object
func (v View) GetLink() IRI {
	return IRI(v.ID)
}

// GetType returns the ActivityVocabulary type of the current Activity
func (v View) GetType() ActivityVocabularyType {
	return v.Type
}

// IsObject returns true for View objects
func (v View) IsObject() bool {
	return true
}

// IsLink returns false for View objects
func (v View) IsLink() bool {
	return false
}

// UnmarshalJSON
func (a *Activity) UnmarshalJSON(data []byte) error {
	a.Parent.UnmarshalJSON(data)
	a.Actor = getAPItem(data, "actor")
	a.Object = getAPItem(data, "object")
	return nil
}

// UnmarshalJSON
func (l *Like) UnmarshalJSON(data []byte) error {
	a := Activity(*l)
	err := a.UnmarshalJSON(data)

	*l = Like(a)

	return err
}

// UnmarshalJSON
func (d *Dislike) UnmarshalJSON(data []byte) error {
	a := Activity(*d)
	err := a.UnmarshalJSON(data)

	*d = Dislike(a)

	return err
}

// UnmarshalJSON
func (u *Update) UnmarshalJSON(data []byte) error {
	a := Activity(*u)
	err := a.UnmarshalJSON(data)

	*u = Update(a)

	return err
}

// UnmarshalJSON
func (c *Create) UnmarshalJSON(data []byte) error {
	a := Activity(*c)
	err := a.UnmarshalJSON(data)

	*c = Create(a)

	return err
}
