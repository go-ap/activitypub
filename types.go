package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"slices"

	"github.com/go-ap/errors"
	"github.com/valyala/fastjson"
)

// ActivityVocabularyType is the data type for an Activity type object
type ActivityVocabularyType string

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

func (t ActivityVocabularyTypes) MatchOther(typ TypeMatcher) bool {
	if typ == nil {
		return len(t) == 0 || len(t) == 1 && t[0] == NilType
	}
	return typ.Matches(t...)
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
	e1 := EmptyTypes(t...)
	e2 := EmptyTypes(tt...)
	if e1 || e2 {
		return e1 == e2
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
	if types := JSONGetTypes(val); types != nil {
		if typ, ok := types.(ActivityVocabularyType); ok {
			*t = ActivityVocabularyTypes{typ}
		}
		if typ, ok := types.(ActivityVocabularyTypes); ok {
			*t = typ
		}
	}
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

func (a ActivityVocabularyTypes) Contains(typ ActivityVocabularyType) bool {
	return slices.Contains(a, typ)
}

func OnType(t TypeMatcher, fn func(typ ...ActivityVocabularyType) error) error {
	if fn == nil {
		return errors.Newf("nil TypeFn")
	}
	var onRun ActivityVocabularyTypes
	if tt, ok := t.(ActivityVocabularyType); ok {
		onRun = ActivityVocabularyTypes{tt}
	}
	if tt, ok := t.(ActivityVocabularyTypes); ok {
		onRun = tt
	}
	return fn(onRun...)
}

func HasTypes(it ActivityObject) bool {
	if it == nil {
		return false
	}
	result := false
	_ = OnType(it.GetType(), func(typ ...ActivityVocabularyType) error {
		result = !EmptyTypes(typ...)
		return nil
	})
	return result
}

func TypesMatch(m1, m2 TypeMatcher) bool {
	matcherIsNilOrZero := func(m TypeMatcher) bool {
		if typ, ok := m.(ActivityVocabularyType); ok {
			return EmptyTypes(typ)
		}
		if types, ok := m.(ActivityVocabularyTypes); ok {
			return EmptyTypes(types...)
		}
		return m == nil
	}
	if m1 == nil || m2 == nil {
		return matcherIsNilOrZero(m1) && matcherIsNilOrZero(m2)
	}

	result := false
	_ = OnType(m1, func(typ ...ActivityVocabularyType) error {
		result = m2.Matches(typ...)
		return nil
	})
	return result
}

func (a ActivityVocabularyType) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return nil, nil
	}
	b := make([]byte, 0)
	JSONWriteStringValue(&b, string(a))
	return b, nil
}

// GobEncode
func (a ActivityVocabularyType) GobEncode() ([]byte, error) {
	return []byte(a), nil
}

// GobDecode
func (a *ActivityVocabularyType) GobDecode(data []byte) error {
	*a = ActivityVocabularyType(data)
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (a *ActivityVocabularyType) UnmarshalBinary(data []byte) error {
	return a.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (a ActivityVocabularyType) MarshalBinary() ([]byte, error) {
	return a.GobEncode()
}

// Matches returns whether the receiver matches the ActivityVocabularyType arguments.
func (a ActivityVocabularyType) Matches(tt ...ActivityVocabularyType) (match bool) {
	if a == NilType && EmptyTypes(tt...) {
		return true
	}
	for _, search := range tt {
		if match = a == search; match {
			break
		}
	}
	return match
}

func (a ActivityVocabularyType) MatchOther(m TypeMatcher) bool {
	if m == nil {
		return a == NilType
	}
	return m.Matches(a)
}
