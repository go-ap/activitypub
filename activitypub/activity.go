package activitypub

const (
	// Activity Types
	AcceptType          string = "Accept"
	AddType             string = "Add"
	AnnounceType        string = "Announce"
	ArriveType          string = "Arrive"
	BlockType           string = "Block"
	CreateType          string = "Create"
	DeleteType          string = "Delete"
	DislikeType         string = "Dislike"
	FlagType            string = "Flag"
	FollowType          string = "Follow"
	IgnoreType          string = "Ignore"
	InviteType          string = "Invite"
	JoinType            string = "Join"
	LeaveType           string = "Leave"
	LikeType            string = "Like"
	ListenType          string = "Listen"
	MoveType            string = "Move"
	OfferType           string = "Offer"
	QuestionType        string = "Question"
	RejectType          string = "Reject"
	ReadType            string = "Read"
	RemoveType          string = "Remove"
	TentativeRejectType string = "TentativeReject"
	TentativeAcceptType string = "TentativeAccept"
	TravelType          string = "Travel"
	UndoType            string = "Undo"
	UpdateType          string = "Update"
	ViewType            string = "View"
)

var validActivityTypes = [...]string{
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

// An Activity is a subtype of Object that describes some form of action that may happen,
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
	Object 			ObjectOrLink 			`jsonld:"object,omitempty"`
}

// Indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate
//  the context into which the object has been accepted.
type Accept Activity

// Indicates that the actor has added the object to the target. If the target property is not explicitly specified,
//  the target would need to be determined implicitly by context.
// The origin can be used to identify the context from which the object originated.
type Add Activity

// Indicates that the actor is calling the target's attention the object.
// The origin typically has no defined meaning.
type Announce Activity

// An IntransitiveActivity that indicates that the actor has arrived at the location.
// The origin can be used to identify the context from which the actor originated.
// The target typically has no defined meaning.
type Arrive IntransitiveActivity

// Indicates that the actor is blocking the object. Blocking is a stronger form of Ignore.
// The typical use is to support social systems that allow one user to block activities or content of other users.
// The target and origin typically have no defined meaning.
type Block Ignore

// Indicates that the actor has created the object.
type Create Activity

// Indicates that the actor has deleted the object.
// If specified, the origin indicates the context from which the object was deleted.
type Delete Activity

// Indicates that the actor dislikes the object.
type Dislike Activity

// Indicates that the actor is "flagging" the object.
// Flagging is defined in the sense common to many social platforms as reporting content as being
//  inappropriate for any number of reasons.
type Flag Activity

// Indicates that the actor is "following" the object. Following is defined in the sense typically used within
//  Social systems in which the actor is interested in any activity performed by or on the object.
// The target and origin typically have no defined meaning.
type Follow Activity

// Indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
type Ignore Activity

// A specialization of Offer in which the actor is extending an invitation for the object to the target.
type Invite Offer

// Indicates that the actor has joined the object. The target and origin typically have no defined meaning.
type Join Activity

// Indicates that the actor has left the object. The target and origin typically have no meaning.
type Leave Activity

// Indicates that the actor likes, recommends or endorses the object.
// The target and origin typically have no defined meaning.
type Like Activity

// Inherits all properties from Activity.
type Listen Activity

// Indicates that the actor has moved object from origin to target.
// If the origin or target are not specified, either can be determined by context.
type Move Activity

// Indicates that the actor is offering the object.
// If specified, the target indicates the entity to which the object is being offered.
type Offer Activity

// Indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
type Reject Activity

// Indicates that the actor has read the object.
type Read Activity

// Indicates that the actor is removing the object. If specified,
//  the origin indicates the context from which the object is being removed.
type Remove Activity

// A specialization of Reject in which the rejection is considered tentative.
type TentativeReject Reject

// A specialization of Accept indicating that the acceptance is tentative.
// A specialization of Accept indicating that the acceptance is tentative.
type TentativeAccept Accept

// Indicates that the actor is traveling to target from origin.
// Travel is an IntransitiveObject whose actor specifies the direct object.
// If the target or origin are not specified, either can be determined by context.
type Travel IntransitiveActivity

// Indicates that the actor is undoing the object. In most cases, the object will be an Activity describing
//  some previously performed action (for instance, a person may have previously "liked" an article but,
//  for whatever reason, might choose to undo that like at some later point in time).
// The target and origin typically have no defined meaning.
type Undo Activity

// Indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism
//  for describing the actual set of modifications made to object.
// The target and origin typically have no defined meaning.
type Update Activity

// Indicates that the actor has viewed the object.
type View Activity

// Represents a question being asked. Question objects are an extension of IntransitiveActivity.
// That is, the Question object is an Activity, but the direct object is the question
//  itself and therefore it would not contain an object property.
// Either of the anyOf and oneOf properties may be used to express possible answers,
//  but a Question object must not have both properties.
type Question struct {
	*IntransitiveActivity
	// Identifies an exclusive option for a Question. Use of oneOf implies that the Question
	//  can have only a single answer. To indicate that a Question can have multiple answers, use anyOf.
	OneOf	ObjectOrLink	`jsonld:"oneOf,omitempty"`
	// Identifies an inclusive option for a Question. Use of anyOf implies that the Question can have multiple answers.
	// To indicate that a Question can have only one answer, use oneOf.
	AnyOf	ObjectOrLink	`jsonld:"anyOf,omitempty"`
	// Indicates that a question has been closed, and answers are no longer accepted.
	Closed	bool			`jsonld:"closed,omitempty"`
}

// Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions.
// The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
	*BaseObject
	// Describes one or more entities that either performed or are expected to perform the activity.
	// Any single activity can have multiple actors. The actor may be specified using an indirect Link.
	Actor  			Actor 					`jsonld:"actor,omitempty"`
	// Describes the indirect object, or target, of the activity.
	// The precise meaning of the target is largely dependent on the type of action being described
	//  but will often be the object of the English preposition "to".
	// For instance, in the activity "John added a movie to his wishlist",
	//  the target of the activity is John's wishlist. An activity can have more than one target.
	Target 			ObjectOrLink 			`jsonld:"actor,omitempty"`
	// Describes the result of the activity. For instance, if a particular action results in the creation
	//  of a new resource, the result property can be used to describe that new resource.
	Result 			ObjectOrLink 			`jsonld:"actor,omitempty"`
	// Describes an indirect object of the activity from which the activity is directed.
	// The precise meaning of the origin is the object of the English preposition "from".
	// For instance, in the activity "John moved an item to List B from List A", the origin of the activity is "List A".
	Origin 			ObjectOrLink			`jsonld:"origin,omitempty"`
	// Identifies one or more objects used (or to be used) in the completion of an Activity.
	Instrument 		ObjectOrLink			`jsonld:"instrument,omitempty"`
	Source Source
}

func AcceptNew(id ObjectId) *Accept {
	a := ActivityNew(id, AcceptType)
	o := Accept(*a)
	return &o
}
func AddNew(id ObjectId) *Add {
	a := ActivityNew(id, AddType)
	o := Add(*a)
	return &o
}
func AnnounceNew(id ObjectId) *Announce {
	a := ActivityNew(id, AnnounceType)
	o := Announce(*a)
	return &o
}
func ArriveNew(id ObjectId) *Arrive {
	a := IntransitiveActivityNew(id, ArriveType)
	o := Arrive(*a)
	return &o
}
func BlockNew(id ObjectId) *Block {
	a := ActivityNew(id, BlockType)
	o := Block(*a)
	return &o
}
func CreateNew(id ObjectId) *Create {
	a := ActivityNew(id, CreateType)
	o := Create(*a)
	return &o
}
func DeleteNew(id ObjectId) *Delete {
	a := ActivityNew(id, DeleteType)
	o := Delete(*a)
	return &o
}
func DislikeNew(id ObjectId) *Dislike {
	a := ActivityNew(id, DislikeType)
	o := Dislike(*a)
	return &o
}
func FlagNew(id ObjectId) *Flag {
	a := ActivityNew(id, FlagType)
	o := Flag(*a)
	return &o
}
func FollowNew(id ObjectId) *Follow {
	a := ActivityNew(id, FollowType)
	o := Follow(*a)
	return &o
}
func IgnoreNew(id ObjectId) *Ignore {
	a := ActivityNew(id, IgnoreType)
	o := Ignore(*a)
	return &o
}
func InviteNew(id ObjectId) *Invite {
	a := ActivityNew(id, InviteType)
	o := Invite(*a)
	return &o
}
func JoinNew(id ObjectId) *Join {
	a := ActivityNew(id, JoinType)
	o := Join(*a)
	return &o
}
func LeaveNew(id ObjectId) *Leave {
	a := ActivityNew(id, LeaveType)
	o := Leave(*a)
	return &o
}
func LikeNew(id ObjectId) *Like {
	a := ActivityNew(id, LikeType)
	o := Like(*a)
	return &o
}
func ListenNew(id ObjectId) *Listen {
	a := ActivityNew(id, ListenType)
	o := Listen(*a)
	return &o
}
func MoveNew(id ObjectId) *Move {
	a := ActivityNew(id, MoveType)
	o := Move(*a)
	return &o
}
func OfferNew(id ObjectId) *Offer {
	a := ActivityNew(id, OfferType)
	o := Offer(*a)
	return &o
}
func RejectNew(id ObjectId) *Reject {
	a := ActivityNew(id, RejectType)
	o := Reject(*a)
	return &o
}
func ReadNew(id ObjectId) *Read {
	a := ActivityNew(id, ReadType)
	o := Read(*a)
	return &o
}
func RemoveNew(id ObjectId) *Remove {
	a := ActivityNew(id, RemoveType)
	o := Remove(*a)
	return &o
}
func TentativeRejectNew(id ObjectId) *TentativeReject {
	a := ActivityNew(id, TentativeRejectType)
	o := TentativeReject(*a)
	return &o
}
func TentativeAcceptNew(id ObjectId) *TentativeAccept {
	a := ActivityNew(id, TentativeAcceptType)
	o := TentativeAccept(*a)
	return &o
}
func TravelNew(id ObjectId) *Travel {
	a := IntransitiveActivityNew(id, TravelType)
	o := Travel(*a)
	return &o
}
func UndoNew(id ObjectId) *Undo {
	a := ActivityNew(id, UndoType)
	o := Undo(*a)
	return &o
}
func UpdateNew(id ObjectId) *Accept {
	a := ActivityNew(id, UpdateType)
	o := Accept(*a)
	return &o
}
func ViewNew(id ObjectId) *View {
	a := ActivityNew(id, ViewType)
	o := View(*a)
	return &o
}
func QuestionNew(id ObjectId) *Question {
	a := IntransitiveActivityNew(id, QuestionType)
	o := Question{IntransitiveActivity: a}
	return &o
}

func ValidActivityType(_type string) bool {
	for _, v := range validActivityTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ActivityNew(id ObjectId, _type string) *Activity {
	if !ValidActivityType(_type) {
		_type = ActivityType
	}
	o := BaseObject{Id: id, Type: _type}
	a := IntransitiveActivity{BaseObject: &o}

	return &Activity{IntransitiveActivity: &a}
}

func IntransitiveActivityNew(id ObjectId, _type string) *IntransitiveActivity {
	if !ValidActivityType(_type) {
		_type = IntransitiveActivityType
	}
	o := BaseObject{Id: id, Type: _type}

	return &IntransitiveActivity{BaseObject: &o}
}

