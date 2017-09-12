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
	// @see activitypub.Object
	Object 			ObjectOrLink 			`jsonld:"object,omitempty"`
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

