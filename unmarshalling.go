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
	var ret Item
	var err error

	switch typ {
	case ObjectType:
		ret = ObjectNew(typ)
	case LinkType:
		ret = &Link{}
		o := ret.(*Link)
		o.Type = typ
	case ActivityType:
		ret = &Activity{}
		o := ret.(*Activity)
		o.Type = typ
	case IntransitiveActivityType:
		ret = &IntransitiveActivity{}
		o := ret.(*IntransitiveActivity)
		o.Type = typ
	case ActorType:
		ret = &Object{}
		o := ret.(*Object)
		o.Type = typ
	case CollectionType:
		ret = &Collection{}
		o := ret.(*Collection)
		o.Type = typ
	case OrderedCollectionType:
		ret = &OrderedCollection{}
		o := ret.(*OrderedCollection)
		o.Type = typ
	case CollectionPageType:
		ret = &CollectionPage{}
		o := ret.(*CollectionPage)
		o.Type = typ
	case OrderedCollectionPageType:
		ret = &OrderedCollectionPage{}
		o := ret.(*OrderedCollectionPage)
		o.Type = typ
	case ArticleType:
		ret = ObjectNew(typ)
	case AudioType:
		ret = ObjectNew(typ)
	case DocumentType:
		ret = ObjectNew(typ)
	case EventType:
		ret = ObjectNew(typ)
	case ImageType:
		ret = ObjectNew(typ)
	case NoteType:
		ret = ObjectNew(typ)
	case PageType:
		ret = ObjectNew(typ)
	case PlaceType:
		p := &Place{}
		p.Type = typ
		ret = p
	case ProfileType:
		p := &Profile{}
		p.Type = typ
		ret = p
	case RelationshipType:
		r := &Relationship{}
		r.Type = typ
		ret = r
	case TombstoneType:
		t := &Tombstone{}
		t.Type = typ
		ret = t
	case VideoType:
		ret = ObjectNew(typ)
	case MentionType:
		ret = &Mention{}
		o := ret.(*Mention)
		o.Type = typ
	case ApplicationType:
		ret = &Application{}
		o := ret.(*Application)
		o.Type = typ
	case GroupType:
		ret = &Group{}
		o := ret.(*Group)
		o.Type = typ
	case OrganizationType:
		ret = &Organization{}
		o := ret.(*Organization)
		o.Type = typ
	case PersonType:
		ret = &Person{}
		o := ret.(*Person)
		o.Type = typ
	case ServiceType:
		ret = &Service{}
		o := ret.(*Service)
		o.Type = typ
	case AcceptType:
		ret = &Accept{}
		o := ret.(*Accept)
		o.Type = typ
	case AddType:
		ret = &Add{}
		o := ret.(*Add)
		o.Type = typ
	case AnnounceType:
		ret = &Announce{}
		o := ret.(*Announce)
		o.Type = typ
	case ArriveType:
		ret = &Arrive{}
		o := ret.(*Arrive)
		o.Type = typ
	case BlockType:
		ret = &Block{}
		o := ret.(*Block)
		o.Type = typ
	case CreateType:
		ret = &Create{}
		o := ret.(*Create)
		o.Type = typ
	case DeleteType:
		ret = &Delete{}
		o := ret.(*Delete)
		o.Type = typ
	case DislikeType:
		ret = &Dislike{}
		o := ret.(*Dislike)
		o.Type = typ
	case FlagType:
		ret = &Flag{}
		o := ret.(*Flag)
		o.Type = typ
	case FollowType:
		ret = &Follow{}
		o := ret.(*Follow)
		o.Type = typ
	case IgnoreType:
		ret = &Ignore{}
		o := ret.(*Ignore)
		o.Type = typ
	case InviteType:
		ret = &Invite{}
		o := ret.(*Invite)
		o.Type = typ
	case JoinType:
		ret = &Join{}
		o := ret.(*Join)
		o.Type = typ
	case LeaveType:
		ret = &Leave{}
		o := ret.(*Leave)
		o.Type = typ
	case LikeType:
		ret = &Like{}
		o := ret.(*Like)
		o.Type = typ
	case ListenType:
		ret = &Listen{}
		o := ret.(*Listen)
		o.Type = typ
	case MoveType:
		ret = &Move{}
		o := ret.(*Move)
		o.Type = typ
	case OfferType:
		ret = &Offer{}
		o := ret.(*Offer)
		o.Type = typ
	case QuestionType:
		ret = &Question{}
		o := ret.(*Question)
		o.Type = typ
	case RejectType:
		ret = &Reject{}
		o := ret.(*Reject)
		o.Type = typ
	case ReadType:
		ret = &Read{}
		o := ret.(*Read)
		o.Type = typ
	case RemoveType:
		ret = &Remove{}
		o := ret.(*Remove)
		o.Type = typ
	case TentativeRejectType:
		ret = &TentativeReject{}
		o := ret.(*TentativeReject)
		o.Type = typ
	case TentativeAcceptType:
		ret = &TentativeAccept{}
		o := ret.(*TentativeAccept)
		o.Type = typ
	case TravelType:
		ret = &Travel{}
		o := ret.(*Travel)
		o.Type = typ
	case UndoType:
		ret = &Undo{}
		o := ret.(*Undo)
		o.Type = typ
	case UpdateType:
		ret = &Update{}
		o := ret.(*Update)
		o.Type = typ
	case ViewType:
		ret = &View{}
		o := ret.(*View)
		o.Type = typ
	case "":
		// when no type is available use a plain Object
		ret = &Object{}
	default:
		ret = nil
		err = fmt.Errorf("unrecognized ActivityStreams type %s", typ)
	}
	return ret, err
}
