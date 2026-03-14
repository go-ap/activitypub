package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"slices"
	"strings"

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

// EmptyTypes returns whether the collection of [ActivityVocabularyType]s is empty
func EmptyTypes(tt ...ActivityVocabularyType) bool {
	if len(tt) == 0 {
		return true
	}
	return len(flattenTypes(tt...)) == 0
}

func (at ActivityVocabularyTypes) AsTypes() ActivityVocabularyTypes {
	return at
}

// Match returns whether the receiver matches the ActivityVocabularyType arguments.
func (at ActivityVocabularyTypes) Match(other Typer) bool {
	return AnyTypes(at...).Match(types(other)...)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (at ActivityVocabularyTypes) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	if !JSONWriteActivityVocabularyTypes(&b, at) {
		return nil, fmt.Errorf("error JSON encoding ActivityVocabularyTypes")
	}
	return b.Bytes(), nil
}

// UnmarshalJSON decodes the receiver type from the JSON document.
func (at *ActivityVocabularyTypes) UnmarshalJSON(b []byte) error {
	if at == nil {
		return fmt.Errorf("nil ActivityVocabularyTypes receiver")
	}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(b)
	if err != nil {
		return err
	}
	if types := JSONGetTypes(val); types != nil {
		if typ, ok := types.(ActivityVocabularyType); ok {
			*at = ActivityVocabularyTypes{typ}
		}
		if typ, ok := types.(ActivityVocabularyTypes); ok {
			*at = typ
		}
	}
	return nil
}

// GobEncode
func (at ActivityVocabularyTypes) GobEncode() ([]byte, error) {
	switch len(at) {
	case 0:
		return nil, nil
	case 1:
		return at[0].GobEncode()
	default:
		tt := make([][]byte, len(at))
		for i, ty := range at {
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
func (at *ActivityVocabularyTypes) GobDecode(data []byte) error {
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
		if at == nil {
			at = &types
		} else {
			*at = types
		}
	} else {
		t := NilType
		if err := t.GobDecode(data); err != nil {
			return err
		}
		*at = ActivityVocabularyTypes{t}
	}

	return nil
}

func HasTypes(it ActivityObject) bool {
	if it == nil || it.GetType() == nil {
		return false
	}
	return !EmptyTypes(it.GetType().AsTypes()...)
}

func TypesEqual(m1, m2 Typer) bool {
	matcherIsNilOrZero := func(m Typer) bool {
		if m == nil {
			return true
		}
		all := m.AsTypes()
		if len(all) == 0 {
			return true
		}
		for _, t := range all {
			if t != NilType {
				return false
			}
		}
		return true
	}
	if m1 == nil || m2 == nil {
		return matcherIsNilOrZero(m1) && matcherIsNilOrZero(m2)
	}

	t1 := m1.AsTypes()
	t2 := m2.AsTypes()
	return AnyTypes(t1...).Match(t2...)
}

func (a ActivityVocabularyType) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return nil, nil
	}
	b := bytes.Buffer{}
	JSONWriteStringValue(&b, string(a))
	return b.Bytes(), nil
}

func (t ActivityVocabularyType) String() string {
	return string(t)
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

func (a ActivityVocabularyType) AsTypes() ActivityVocabularyTypes {
	if a == NilType {
		return ActivityVocabularyTypes{}
	}
	return ActivityVocabularyTypes{a}
}

func (at ActivityVocabularyTypes) String() string {
	if len(at) == 0 {
		return ""
	}
	s := make([]string, 0, len(at))
	for _, tt := range at {
		s = append(s, string(tt))
	}
	return strings.Join(s, ", ")
}

type TypesMatcher interface {
	Match(...ActivityVocabularyType) bool
}

type MatcherFn func(...ActivityVocabularyType) bool

func (mfn MatcherFn) Match(tt ...ActivityVocabularyType) bool {
	return mfn(tt...)
}

func AllTypes(toMatch ...ActivityVocabularyType) MatcherFn {
	return func(toCheck ...ActivityVocabularyType) bool {
		if len(toMatch) == 0 {
			if len(toCheck) == 0 {
				return true
			}
		}
		for _, search := range toMatch {
			if !slices.Contains(toCheck, search) {
				return false
			}
		}
		return true
	}
}

func flattenTypes(typ ...ActivityVocabularyType) ActivityVocabularyTypes {
	result := make(ActivityVocabularyTypes, 0, len(typ))
	for _, t := range typ {
		if t == NilType {
			continue
		}
		result = append(result, t)
	}
	return result
}

func AnyTypes(toMatch ...ActivityVocabularyType) MatcherFn {
	return func(toCheck ...ActivityVocabularyType) bool {
		if len(toMatch) == 0 || len(toCheck) == 0 {
			// NOTE(marius): if one of the type slices is empty expresses that the code wants
			// to match for an empty Type property.
			return len(flattenTypes(toCheck...)) == 0 && len(flattenTypes(toMatch...)) == 0
		}
		for _, search := range toMatch {
			if slices.Contains(toCheck, search) {
				return true
			}
		}
		return false
	}
}

// Match returns whether the receiver matches the ActivityVocabularyType arguments.
func (a ActivityVocabularyType) Match(other Typer) (match bool) {
	if other == nil {
		return a == NilType
	}
	tt := types(other)
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
