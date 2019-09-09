package activitystreams

import (
	"encoding"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/buger/jsonparser"
)

var (
	apUnmarshalerType   = reflect.TypeOf(new(Item)).Elem()
	unmarshalerType     = reflect.TypeOf(new(json.Unmarshaler)).Elem()
	textUnmarshalerType = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()
)

// ItemTyperFunc will return an instance of a struct that implements activitystreams.Item
// The default for this package is JSONGetItemByType but can be overwritten
var ItemTyperFunc TyperFunction

// TyperFunction is the type of the function which returns an activitystreams.Item struct instance
// for a specific ActivityVocabularyType
type TyperFunction func(ActivityVocabularyType) (Item, error)

func JSONGetObjectID(data []byte) ObjectID {
	i, err := jsonparser.GetString(data, "id")
	if err != nil {
		return ObjectID("")
	}
	return ObjectID(i)
}

func JSONGetType(data []byte) ActivityVocabularyType {
	t, err := jsonparser.GetString(data, "type")
	typ := ActivityVocabularyType(t)
	if err != nil {
		return ActivityVocabularyType("")
	}
	return typ
}

func JSONGetMimeType(data []byte) MimeType {
	t, err := jsonparser.GetString(data, "mediaType")
	if err != nil {
		return MimeType("")
	}
	return MimeType(t)
}

func JSONGetInt(data []byte, prop string) int64 {
	val, err := jsonparser.GetInt(data, prop)
	if err != nil {
	}
	return val
}

func JSONGetString(data []byte, prop string) string {
	val, err := jsonparser.GetString(data, prop)
	if err != nil {
	}
	return val
}

func JSONGetNaturalLanguageField(data []byte, prop string) NaturalLanguageValues {
	n := NaturalLanguageValues{}
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			if dataType == jsonparser.String {
				n.Append(LangRef(key), string(value))
			}
			return err
		})
	case jsonparser.String:
		n.Append(NilLangRef, string(val))
	}

	return n
}

func JSONGetTime(data []byte, prop string) time.Time {
	t := time.Time{}
	str, _ := jsonparser.GetUnsafeString(data, prop)
	t.UnmarshalText([]byte(str))
	return t
}

func JSONGetDuration(data []byte, prop string) time.Duration {
	str, _ := jsonparser.GetUnsafeString(data, prop)
	d, _ := time.ParseDuration(str)
	return d
}

func JSONUnmarshalToItem(data []byte) Item {
	if _, err := url.ParseRequestURI(string(data)); err == nil {
		// try to see if it's an IRI
		return IRI(data)
	}

	i, err := ItemTyperFunc(JSONGetType(data))
	if err != nil {
		return nil
	}
	p := reflect.PtrTo(reflect.TypeOf(i))
	if reflect.TypeOf(i).Implements(unmarshalerType) || p.Implements(unmarshalerType) {
		err = i.(json.Unmarshaler).UnmarshalJSON(data)
	}
	if reflect.TypeOf(i).Implements(textUnmarshalerType) || p.Implements(textUnmarshalerType) {
		err = i.(encoding.TextUnmarshaler).UnmarshalText(data)
	}
	if err != nil {
		return nil
	}
	return i
}

func JSONGetItem(data []byte, prop string) Item {
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}
	switch typ {
	case jsonparser.String:
		if _, err = url.ParseRequestURI(string(val)); err == nil {
			// try to see if it's an IRI
			return IRI(val)
		}
	case jsonparser.Object:
		return JSONUnmarshalToItem(val)
	case jsonparser.Number:
		fallthrough
	case jsonparser.Array:
		fallthrough
	case jsonparser.Boolean:
		fallthrough
	case jsonparser.Null:
		fallthrough
	case jsonparser.Unknown:
		fallthrough
	default:
		return nil
	}
	return nil
}

func JSONGetItems(data []byte, prop string) ItemCollection {
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}

	var it ItemCollection
	switch typ {
	case jsonparser.Array:
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			i := JSONUnmarshalToItem(value)
			if i != nil {
				it.Append(i)
			}
		}, prop)
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			i := JSONUnmarshalToItem(value)
			if i != nil {
				it.Append(i)
			}
			return err
		}, prop)
	case jsonparser.String:
		s, _ := jsonparser.GetString(val)
		it.Append(IRI(s))
	}
	return it
}

func JSONGetURIItem(data []byte, prop string) Item {
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}

	switch typ {
	case jsonparser.Object:
		return JSONGetItem(data, prop)
	case jsonparser.Array:
		var it ItemCollection
		jsonparser.ArrayEach(val, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if _, err := url.Parse(string(value)); err == nil {
				it.Append(IRI(value))
				return
			}
			i, err := ItemTyperFunc(JSONGetType(value))
			if err != nil {
				return
			}
			err = i.(json.Unmarshaler).UnmarshalJSON(value)
			if err != nil {
				return
			}
			it.Append(i)
		})

		return it
	case jsonparser.String:
		return IRI(val)
	}

	return nil
}

func JSONGetLangRefField(data []byte, prop string) LangRef {
	val, err := jsonparser.GetString(data, prop)
	if err != nil {
		return LangRef("")
	}
	return LangRef(val)
}

func JSONGetIRI(data []byte, prop string) IRI {
	val, err := jsonparser.GetString(data, prop)
	if err != nil {
		return IRI("")
	}
	return IRI(val)
}

// UnmarshalJSON tries to detect the type of the object in the json data and then outputs a matching
// ActivityStreams object, if possible
func UnmarshalJSON(data []byte) (Item, error) {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	return JSONUnmarshalToItem(data), nil
}

func JSONGetItemByType(typ ActivityVocabularyType) (Item, error) {
	switch typ {
	case ObjectType:
		return ObjectNew(typ), nil
	case LinkType:
		return &Link{Type: typ}, nil
	case ActivityType:
		return &Activity{Parent: Parent{Type: typ}}, nil
	case IntransitiveActivityType:
		return &IntransitiveActivity{Parent: Parent{Type: typ}}, nil
	case ActorType:
		return ObjectNew(typ), nil
	case CollectionType:
		return &Collection{Parent: Parent{Type: typ}}, nil
	case OrderedCollectionType:
		return &OrderedCollection{Parent: Parent{Type: typ}}, nil
	case CollectionPageType:
		return &CollectionPage{ParentCollection: ParentCollection{Parent: Parent{Type: typ}}}, nil
	case OrderedCollectionPageType:
		return &OrderedCollectionPage{OrderedCollection: OrderedCollection{Parent: Parent{Type: typ}}}, nil
	case ArticleType:
		return ObjectNew(typ), nil
	case AudioType:
		return ObjectNew(typ), nil
	case DocumentType:
		return ObjectNew(typ), nil
	case EventType:
		return ObjectNew(typ), nil
	case ImageType:
		return ObjectNew(typ), nil
	case NoteType:
		return ObjectNew(typ), nil
	case PageType:
		return ObjectNew(typ), nil
	case PlaceType:
		return &Place{Parent: Parent{Type: typ}}, nil
	case ProfileType:
		return &Profile{Parent: Parent{Type: typ}}, nil
	case RelationshipType:
		return &Relationship{Parent: Parent{Type: typ}}, nil
	case TombstoneType:
		return &Tombstone{Parent: Parent{Type: typ}}, nil
	case VideoType:
		return ObjectNew(typ), nil
	case MentionType:
		return &Mention{Type:typ}, nil
	case ApplicationType:
		return ObjectNew(typ), nil
	case GroupType:
		return ObjectNew(typ), nil
	case OrganizationType:
		return ObjectNew(typ), nil
	case PersonType:
		return ObjectNew(typ), nil
	case ServiceType:
		return ObjectNew(typ), nil
	case AcceptType:
		return &Accept{Parent:Parent{Type: typ}}, nil
	case AddType:
		return &Add{Parent:Parent{Type: typ}}, nil
	case AnnounceType:
		return &Announce{Parent:Parent{Type: typ}}, nil
	case ArriveType:
		return &Arrive{Parent:Parent{Type: typ}}, nil
	case BlockType:
		return &Block{Parent:Parent{Type: typ}}, nil
	case CreateType:
		return &Create{Parent:Parent{Type: typ}}, nil
	case DeleteType:
		return &Delete{Parent:Parent{Type: typ}}, nil
	case DislikeType:
		return &Dislike{Parent:Parent{Type: typ}}, nil
	case FlagType:
		return &Flag{Parent:Parent{Type: typ}}, nil
	case FollowType:
		return &Follow{Parent:Parent{Type: typ}}, nil
	case IgnoreType:
		return &Ignore{Parent:Parent{Type: typ}}, nil
	case InviteType:
		return &Invite{Parent:Parent{Type: typ}}, nil
	case JoinType:
		return &Join{Parent:Parent{Type: typ}}, nil
	case LeaveType:
		return &Leave{Parent:Parent{Type: typ}}, nil
	case LikeType:
		return &Like{Parent:Parent{Type: typ}}, nil
	case ListenType:
		return &Listen{Parent:Parent{Type: typ}}, nil
	case MoveType:
		return &Move{Parent:Parent{Type: typ}}, nil
	case OfferType:
		return &Offer{Parent:Parent{Type: typ}}, nil
	case QuestionType:
		return &Question{Type: typ}, nil
	case RejectType:
		return &Reject{Parent:Parent{Type: typ}}, nil
	case ReadType:
		return &Read{Parent:Parent{Type: typ}}, nil
	case RemoveType:
		return &Remove{Parent:Parent{Type: typ}}, nil
	case TentativeRejectType:
		return &TentativeReject{Parent:Parent{Type: typ}}, nil
	case TentativeAcceptType:
		return &TentativeAccept{Parent:Parent{Type: typ}}, nil
	case TravelType:
		return &Travel{Parent:Parent{Type: typ}}, nil
	case UndoType:
		return &Undo{Parent:Parent{Type: typ}}, nil
	case UpdateType:
		return &Update{Parent:Parent{Type: typ}}, nil
	case ViewType:
		return &View{Parent:Parent{Type: typ}}, nil
	case "":
		// when no type is available use a plain Object
		return &Object{}, nil
	}
	return nil, fmt.Errorf("unrecognized ActivityStreams type %s", typ)
}
