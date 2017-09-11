package activitypub

type ObjectId string

const (
	ObjectType   string = "Object"
	LinkType     string = "Link"
	ActivityType string = "Activity"
	ActorType    string = "Actor"
	
	// Object Types
	ArticleType string = "Article"
	AudioType string = "Audio"
	DocumentType string = "Document"
	EventType string = "Event"
	ImageType string = "Image"
	NoteType string = "Note"
	PageType string = "Page"
	PlaceType string = "Place"
	ProfileType string = "Profile"
	RelationshipType string = "Relationship"
	TombstoneType string = "Tombstone"
	VideoType string = "Video"

	// Link Types
	MentionType string = "Mention"

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

	// Actor Types
	ApplicationType  string = "Application"
	GroupType        string = "Group"
	OrganizationType string = "Organization"
	PersonType       string = "Person"
	ServiceType      string = "Service"
)

var validObjectTypes = [...]string{
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
var validLinkTypes = [...]string{
	MentionType,
}
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

type NaturalLanguageValue map[string]string

type BaseObject struct {
	Id   ObjectId
	Type string
	Name NaturalLanguageValue

	Href      string
	HrefLang  string
	MediaType string
}

type ContentType string
type Source struct {
	Content   ContentType
	MediaType string
}

type ActivityObject struct {
	BaseObject
	Actor  Actor
	Object BaseObject
	Source Source
}

func ValidObjectType(_type string) bool {
	for _, v := range validObjectTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ValidLinkType(_type string) bool {
	for _, v := range validLinkTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ValidActivityType(_type string) bool {
	for _, v := range validActivityTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ObjectNew(id ObjectId, _type string) BaseObject {
	if !ValidObjectType(_type) {
		_type = ObjectType
	}
	return BaseObject{Id: id, Type: _type}
}

func LinkNew(id ObjectId, _type string) BaseObject {
	if !ValidLinkType(_type) {
		_type = LinkType
	}
	return BaseObject{Id: id, Type:_type}
}

func ActivityNew(id ObjectId, _type string) ActivityObject {
	if !ValidActivityType(_type) {
		_type = ActivityType
	}
	o := BaseObject{Id: id, Type: _type}

	return ActivityObject{BaseObject: o}
}
