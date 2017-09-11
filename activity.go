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

type Activity struct {
	BaseObject
	Actor  Actor
	Object BaseObject
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

func ActivityNew(id ObjectId, _type string) Activity {
	if !ValidActivityType(_type) {
		_type = ActivityType
	}
	o := BaseObject{Id: id, Type: _type}

	return Activity{BaseObject: o}
}

