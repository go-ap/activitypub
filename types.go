package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"slices"

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

func (t ActivityVocabularyTypes) GetType() ActivityVocabularyType {
	switch len(t) {
	case 0:
		return NilType
	default:
		return t[0]
	}
}

func (t ActivityVocabularyTypes) GetTypes() ActivityVocabularyTypes {
	return t
}

// EmptyTypes returns whether the collection of [ActivityVocabularyType]s is empty
func EmptyTypes(tt ...ActivityVocabularyType) bool {
	empty := len(tt) == 0
	allEmpty := !empty
	for _, ty := range tt {
		allEmpty = allEmpty && ty == NilType
	}
	return empty || allEmpty
}

// Matches returns whether the receiver matches the ActivityVocabularyType arguments.
func (t ActivityVocabularyTypes) Matches(tt ...ActivityVocabularyType) bool {
	if EmptyTypes(t...) && EmptyTypes(tt...) {
		return true
	}
	match := true
	for _, search := range tt {
		match = match && slices.Contains(t, search)
	}
	return match
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

// GobEncode
func (a ActivityVocabularyTypes) GobEncode() ([]byte, error) {
	switch len(a) {
	case 0:
		return nil, nil
	case 1:
		return a[0].GobEncode()
	default:
		tt := make([][]byte, len(a))
		for i, ty := range a {
			b, err := ty.GobEncode()
			if err != nil {
				return nil, err
			}
			tt[i] = b
		}
		b := bytes.Buffer{}
		err := gob.NewEncoder(&b).Encode(tt)
		return b.Bytes(), err
	}
}

// GobDecode
func (a *ActivityVocabularyTypes) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		tt := [][]byte{}
		g := gob.NewDecoder(bytes.NewReader(data))
		if err := g.Decode(&tt); err != nil {
			return err
		}
		types := make(ActivityVocabularyTypes, len(tt))
		for i, it := range tt {
			if err := types[i].GobDecode(it); err != nil {
				return err
			}
		}
		if a == nil {
			a = &types
		} else {
			*a = types
		}
	} else {
		at := NilType
		if err := at.GobDecode(data); err != nil {
			return err
		}
		if a == nil {
			a = &ActivityVocabularyTypes{at}
		} else {
			*a = ActivityVocabularyTypes{at}
		}
	}

	return nil
}
