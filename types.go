package activitypub

import (
	"fmt"

	"github.com/valyala/fastjson"
)

// ActivityVocabularyTypes is a type alias for a slice of ActivityVocabularyType elements
type ActivityVocabularyTypes []ActivityVocabularyType

// Types contains all valid types in the ActivityPub vocabulary
var Types = ActivityVocabularyTypes{
	LinkType,
	MentionType,

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

	QuestionType,

	CollectionType,
	OrderedCollectionType,
	CollectionPageType,
	OrderedCollectionPageType,

	ApplicationType,
	GroupType,
	OrganizationType,
	PersonType,
	ServiceType,

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

	ArriveType,
	TravelType,
	QuestionType,
}

// MarshalJSON encodes the receiver object to a JSON document.
func (t ActivityVocabularyTypes) MarshalJSON() ([]byte, error) {
	b := []byte{}
	if !JSONWriteActivityVocabularyTypes(&b, t) {
		return nil, fmt.Errorf("error JSON encoding ActivityVocabularyTypes")
	}
	return b, nil
}

// UnmarshalJSON decodes the receiver type from the JSON document.
func (t *ActivityVocabularyTypes) UnmarshalJSON(b []byte) error {
	if t == nil {
		return fmt.Errorf("nil ActivityVocabularyTypes receiver")
	}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(b)
	if err != nil {
		return err
	}
	*t = JSONGetTypes(val)
	return nil
}
