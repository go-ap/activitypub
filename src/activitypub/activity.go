package activitypub

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

var validActivityTypes = [...]ActivityVocabularyType{
	AcceptType,
	AddType,
	AnnounceType,
	ArriveType,
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
	QuestionType,
	RejectType,
	ReadType,
	RemoveType,
	TentativeRejectType,
	TentativeAcceptType,
	TravelType,
	UndoType,
	UpdateType,
	ViewType,
	// Actor Types
}

// IntransitiveActivity Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions.
// The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
	apObject
	// Describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor Actor `jsonld:"actor,omitempty"`
	// Describes the indirect object, or target, of the activity.
	// The precise meaning of the target is largely dependent on the type of action being described
	//  but will often be the object of the English preposition "to".
	// For instance, in the activity "John added a movie to his wishlist",
	//  the target of the activity is John's wishlist. An activity can have more than one target.
	Target ObjectOrLink `jsonld:"target,omitempty"`
	// Describes the result of the activity. For instance, if a particular action results in the creation
	//  of a new resource, the result property can be used to describe that new resource.
	Result ObjectOrLink `jsonld:"result,omitempty"`
	// Describes an indirect object of the activity from which the activity is directed.
	// The precise meaning of the origin is the object of the English preposition "from".
	// For instance, in the activity "John moved an item to List B from List A", the origin of the activity is "List A".
	Origin ObjectOrLink `jsonld:"origin,omitempty"`
	// Identifies one or more objects used (or to be used) in the completion of an Activity.
	Instrument ObjectOrLink `jsonld:"instrument,omitempty"`
	Source     Source       `jsonld:"source,omitempty"`
}

// Activity is a subtype of apObject that describes some form of action that may happen,
//  is currently happening, or has already happened.
// The Activity type itself serves as an abstract base type for all types of activities.
// It is important to note that the Activity type itself does not carry any specific semantics
//  about the kind of action being taken.
type Activity struct {
	*IntransitiveActivity
	// When used within an Activity, describes the direct object of the activity.
	// For instance, in the activity "John added a movie to his wishlist",
	//  the object of the activity is the movie added.
	// When used within a Relationship describes the entity to which the subject is related.
	Object ObjectOrLink `jsonld:"object,omitempty"`
}

type (
	// Accept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate
	//  the context into which the object has been accepted.
	Accept Activity

	// Add indicates that the actor has added the object to the target. If the target property is not explicitly specified,
	//  the target would need to be determined implicitly by context.
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
	//  inappropriate for any number of reasons.
	Flag Activity

	// Follow indicates that the actor is "following" the object. Following is defined in the sense typically used within
	//  Social systems in which the actor is interested in any activity performed by or on the object.
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
	//  the origin indicates the context from which the object is being removed.
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
	//  some previously performed action (for instance, a person may have previously "liked" an article but,
	//  for whatever reason, might choose to undo that like at some later point in time).
	// The target and origin typically have no defined meaning.
	Undo Activity

	// Update indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism
	//  for describing the actual set of modifications made to object.
	// The target and origin typically have no defined meaning.
	Update Activity

	// View indicates that the actor has viewed the object.
	View Activity
)

// Question represents a question being asked. Question objects are an extension of IntransitiveActivity.
// That is, the Question object is an Activity, but the direct object is the question
//  itself and therefore it would not contain an object property.
// Either of the anyOf and oneOf properties may be used to express possible answers,
//  but a Question object must not have both properties.
type Question struct {
	*IntransitiveActivity
	// Identifies an exclusive option for a Question. Use of oneOf implies that the Question
	//  can have only a single answer. To indicate that a Question can have multiple answers, use anyOf.
	OneOf ObjectOrLink `jsonld:"oneOf,omitempty"`
	// Identifies an inclusive option for a Question. Use of anyOf implies that the Question can have multiple answers.
	// To indicate that a Question can have only one answer, use oneOf.
	AnyOf ObjectOrLink `jsonld:"anyOf,omitempty"`
	// Indicates that a question has been closed, and answers are no longer accepted.
	Closed bool `jsonld:"closed,omitempty"`
}

// AcceptNew initializes an Accept activity
func AcceptNew(id ObjectID, ob ObjectOrLink) *Accept {
	a := ActivityNew(id, AcceptType, ob)
	o := Accept(*a)
	return &o
}

// AddNew initializes an Add activity
func AddNew(id ObjectID, ob ObjectOrLink, trgt ObjectOrLink) *Add {
	a := ActivityNew(id, AddType, ob)
	o := Add(*a)
	o.Target = trgt
	return &o
}

// AnnounceNew initializes an Announce activity
func AnnounceNew(id ObjectID, ob ObjectOrLink) *Announce {
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
func BlockNew(id ObjectID, ob ObjectOrLink) *Block {
	a := ActivityNew(id, BlockType, ob)
	o := Block(*a)
	return &o
}

// CreateNew initializes a Create activity
func CreateNew(id ObjectID, ob ObjectOrLink) *Create {
	a := ActivityNew(id, CreateType, ob)
	o := Create(*a)
	return &o
}

// DeleteNew initializes a Delete activity
func DeleteNew(id ObjectID, ob ObjectOrLink) *Delete {
	a := ActivityNew(id, DeleteType, ob)
	o := Delete(*a)
	return &o
}

// DislikeNew initializes a Dislike activity
func DislikeNew(id ObjectID, ob ObjectOrLink) *Dislike {
	a := ActivityNew(id, DislikeType, ob)
	o := Dislike(*a)
	return &o
}

// FlagNew initializes a Flag activity
func FlagNew(id ObjectID, ob ObjectOrLink) *Flag {
	a := ActivityNew(id, FlagType, ob)
	o := Flag(*a)
	return &o
}

// FollowNew initializes a Follow activity
func FollowNew(id ObjectID, ob ObjectOrLink) *Follow {
	a := ActivityNew(id, FollowType, ob)
	o := Follow(*a)
	return &o
}

// IgnoreNew initializes an Ignore activity
func IgnoreNew(id ObjectID, ob ObjectOrLink) *Ignore {
	a := ActivityNew(id, IgnoreType, ob)
	o := Ignore(*a)
	return &o
}

// InviteNew initializes an Invite activity
func InviteNew(id ObjectID, ob ObjectOrLink) *Invite {
	a := ActivityNew(id, InviteType, ob)
	o := Invite(*a)
	return &o
}

// JoinNew initializes a Join activity
func JoinNew(id ObjectID, ob ObjectOrLink) *Join {
	a := ActivityNew(id, JoinType, ob)
	o := Join(*a)
	return &o
}

// LeaveNew initializes a Leave activity
func LeaveNew(id ObjectID, ob ObjectOrLink) *Leave {
	a := ActivityNew(id, LeaveType, ob)
	o := Leave(*a)
	return &o
}

// LikeNew initializes a Like activity
func LikeNew(id ObjectID, ob ObjectOrLink) *Like {
	a := ActivityNew(id, LikeType, ob)
	o := Like(*a)
	return &o
}

// ListenNew initializes a Listen activity
func ListenNew(id ObjectID, ob ObjectOrLink) *Listen {
	a := ActivityNew(id, ListenType, ob)
	o := Listen(*a)
	return &o
}

// MoveNew initializes a Move activity
func MoveNew(id ObjectID, ob ObjectOrLink) *Move {
	a := ActivityNew(id, MoveType, ob)
	o := Move(*a)
	return &o
}

// OfferNew initializes an Offer activity
func OfferNew(id ObjectID, ob ObjectOrLink) *Offer {
	a := ActivityNew(id, OfferType, ob)
	o := Offer(*a)
	return &o
}

// RejectNew initializes a Reject activity
func RejectNew(id ObjectID, ob ObjectOrLink) *Reject {
	a := ActivityNew(id, RejectType, ob)
	o := Reject(*a)
	return &o
}

// ReadNew initializes a Read activity
func ReadNew(id ObjectID, ob ObjectOrLink) *Read {
	a := ActivityNew(id, ReadType, ob)
	o := Read(*a)
	return &o
}

// RemoveNew initializes a Remove activity
func RemoveNew(id ObjectID, ob ObjectOrLink, trgt ObjectOrLink) *Remove {
	a := ActivityNew(id, RemoveType, ob)
	o := Remove(*a)
	o.Target = trgt
	return &o
}

// TentativeRejectNew initializes a TentativeReject activity
func TentativeRejectNew(id ObjectID, ob ObjectOrLink) *TentativeReject {
	a := ActivityNew(id, TentativeRejectType, ob)
	o := TentativeReject(*a)
	return &o
}

// TentativeAcceptNew initializes a TentativeAccept activity
func TentativeAcceptNew(id ObjectID, ob ObjectOrLink) *TentativeAccept {
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
func UndoNew(id ObjectID, ob ObjectOrLink) *Undo {
	a := ActivityNew(id, UndoType, ob)
	o := Undo(*a)
	return &o
}

// UpdateNew initializes an Update activity
func UpdateNew(id ObjectID, ob ObjectOrLink) *Accept {
	a := ActivityNew(id, UpdateType, ob)
	o := Accept(*a)
	return &o
}

// ViewNew initializes a View activity
func ViewNew(id ObjectID, ob ObjectOrLink) *View {
	a := ActivityNew(id, ViewType, ob)
	o := View(*a)
	return &o
}

// QuestionNew initializes a Question activity
func QuestionNew(id ObjectID) *Question {
	a := IntransitiveActivityNew(id, QuestionType)
	o := Question{IntransitiveActivity: a}
	return &o
}

// ValidActivityType is a validation function for Activity objects
func ValidActivityType(_type ActivityVocabularyType) bool {
	for _, v := range validActivityTypes {
		if v == _type {
			return true
		}
	}
	return false
}

// ActivityNew initializes a basic activity
func ActivityNew(id ObjectID, _type ActivityVocabularyType, ob ObjectOrLink) *Activity {
	if !ValidActivityType(_type) {
		_type = ActivityType
	}
	o := ObjectNew(id, _type)
	i := IntransitiveActivity{apObject: o}

	a := Activity{IntransitiveActivity: &i}
	a.Object = ob

	return &a
}

// IntransitiveActivityNew initializes a intransitive activity
func IntransitiveActivityNew(id ObjectID, _type ActivityVocabularyType) *IntransitiveActivity {
	if !ValidActivityType(_type) {
		_type = IntransitiveActivityType
	}
	o := ObjectNew(id, _type)

	return &IntransitiveActivity{apObject: o}
}
