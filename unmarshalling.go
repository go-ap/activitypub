package activitystreams

import (
	"encoding"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

var (
	apUnmarshalerType   = reflect.TypeOf(new(Item)).Elem()
	unmarshalerType     = reflect.TypeOf(new(json.Unmarshaler)).Elem()
	textUnmarshalerType = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()
)

type mockObj map[string]json.RawMessage

func getType(j json.RawMessage) ActivityVocabularyType {
	mock := make(mockObj, 0)
	json.Unmarshal([]byte(j), &mock)

	for key, val := range mock {
		if strings.ToLower(key) == "type" {
			return ActivityVocabularyType(strings.Trim(string(val), "\""))
		}
	}
	return ""
}

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

func JSONGetNaturalLanguageField(data []byte, prop string) NaturalLanguageValue {
	n := NaturalLanguageValue{}
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

	i, err := getAPObjectByType(JSONGetType(data))
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
			i, err := getAPObjectByType(JSONGetType(value))
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

// UnmarshalJSON tries to detect the type of the object in the json data and then outputs a matching
// ActivityStreams object, if possible
func UnmarshalJSON(data []byte) (Item, error) {
	return JSONUnmarshalToItem(data), nil
}

/*
func unmarshal(data []byte, a interface{}) (interface{}, error) {
	ta := make(mockObj, 0)
	err := jsonld.Unmarshal(data, &ta)
	if err != nil {
		return nil, err
	}

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		cField := typ.Field(i)
		cValue := val.Field(i)
		cTag := cField.Tag
		tag, _ := jsonld.LoadTag(cTag)

		var vv reflect.Value
		for key, j := range ta {
			if j == nil {
				continue
			}
			if key == tag.Name {
				if cField.Type.Implements(textUnmarshalerType) {
					m, _ := cValue.Interface().(encoding.TextUnmarshaler)
					m.UnmarshalText(j)
					vv = reflect.ValueOf(m)
				}
				if cField.Type.Implements(unmarshalerType) {
					m, _ := cValue.Interface().(json.Unmarshaler)
					m.UnmarshalJSON(j)
					vv = reflect.ValueOf(m)
				}
				if cField.Type.Implements(apUnmarshalerType) {
					o := getAPObjectByType(getType(j))
					if o != nil {
						jsonld.Unmarshal([]byte(j), o)
						vv = reflect.ValueOf(o)
					}
				}
			}
			if vv.CanAddr() {
				cValue.Set(vv)
				fmt.Printf("\n\nReflected %q %q => %#v\n\n%#v\n", cField.Name, cField.Type, vv, tag.Name)
			}
		}
	}
	return a, nil
}
*/

func getAPObjectByType(typ ActivityVocabularyType) (Item, error) {
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
		o := Object{}
		o.Type = typ
	case ImageType:
		ret = ObjectNew(typ)
		o := ret.(*Object)
		o.Type = typ
	case NoteType:
		ret = ObjectNew(typ)
	case PageType:
		ret = ObjectNew(typ)
	case PlaceType:
		ret = ObjectNew(typ)
	case ProfileType:
		ret = ObjectNew(typ)
	case RelationshipType:
		ret = ObjectNew(typ)
	case TombstoneType:
		ret = ObjectNew(typ)
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
