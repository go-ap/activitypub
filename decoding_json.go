package activitypub

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/valyala/fastjson"
)

var (
	apUnmarshalerType   = reflect.TypeOf(new(Item)).Elem()
	unmarshalerType     = reflect.TypeOf(new(json.Unmarshaler)).Elem()
	textUnmarshalerType = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()
)

// ItemTyperFunc will return an instance of a struct that implements activitystreams.Item
// The default for this package is GetItemByType but can be overwritten
var ItemTyperFunc TyperFn = GetItemByType

// TyperFn is the type of the function which returns an activitystreams.Item struct instance
// for a specific ActivityVocabularyType
type TyperFn func(ActivityVocabularyType) (Item, error)

func JSONGetID(data []byte) ID {
	i := fastjson.GetString(data, "id")
	return ID(i)
}

func JSONGetType(data []byte) ActivityVocabularyType {
	t := fastjson.GetBytes(data, "type")
	return ActivityVocabularyType(t)
}

func JSONGetMimeType(data []byte, prop string) MimeType {
	t := fastjson.GetString(data, prop)
	return MimeType(t)
}

func JSONGetInt(data []byte, prop string) int64 {
	if len(data) == 0 {
		return 0
	}
	val := fastjson.GetInt(data, prop)
	return int64(val)
}

func JSONGetFloat(data []byte, prop string) float64 {
	if len(data) == 0 {
		return 0
	}
	val := fastjson.GetFloat64(data, prop)
	return val
}

func JSONGetString(data []byte, prop string) string {
	val := fastjson.GetString(data, prop)
	return val
}

func JSONGetBytes(data []byte, prop string) []byte {
	val := fastjson.GetBytes(data, prop)
	return val
}

func JSONGetBoolean(data []byte, prop string) bool {
	val := fastjson.GetBool(data, prop)
	return val
}

func JSONGetNaturalLanguageField(data []byte, prop string) NaturalLanguageValues {
	n := NaturalLanguageValues{}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return nil
	}
	v := val.Get(prop)
	if v == nil {
		return nil
	}
	switch v.Type() {
	case fastjson.TypeObject:
		ob, _ := v.Object()
		ob.Visit(func(key []byte, v *fastjson.Value) {
			l := LangRefValue{}
			l.Ref = LangRef(key)
			if err := l.Value.UnmarshalJSON(v.GetStringBytes()); err == nil {
				if l.Ref != NilLangRef || len(l.Value) > 0 {
					n = append(n, l)
				}
			}
		})
	case fastjson.TypeString:
		l := LangRefValue{}
		if err := l.UnmarshalJSON(v.GetStringBytes()); err == nil {
			n = append(n, l)
		}
	}

	return n
}

func JSONGetTime(data []byte, prop string) time.Time {
	t := time.Time{}
	if str := fastjson.GetBytes(data, prop); len(str) > 0 {
		t.UnmarshalText(str)
		return t.UTC()
	}
	return t
}

func JSONGetDuration(data []byte, prop string) time.Duration {
	if str := fastjson.GetString(data, prop); len(str) > 0 {
		// TODO(marius): this needs to be replaced to be compatible with xsd:duration
		d, _ := time.ParseDuration(str)
		return d
	}
	return 0
}

func JSONGetPublicKey(data []byte, prop string) PublicKey {
	key := PublicKey{}
	key.UnmarshalJSON(JSONGetBytes(data, prop))
	return key
}

func JSONGetStreams(data []byte, prop string) []ItemCollection {
	// TODO(marius)
	return nil
}

func itemFn(data []byte) (Item, error) {
	if len(data) == 0 {
		return nil, nil
	}

	typ := JSONGetType(data)
	if typ == "" {
		// try to see if it's an IRI
		if i, ok := asIRI(data); ok {
			return i, nil
		}
	}
	i, err := ItemTyperFunc(typ)
	if err != nil || IsNil(i) {
		return nil, nil
	}

	p := reflect.PtrTo(reflect.TypeOf(i))
	if reflect.TypeOf(i).Implements(unmarshalerType) || p.Implements(unmarshalerType) {
		err = i.(json.Unmarshaler).UnmarshalJSON(data)
	}
	if reflect.TypeOf(i).Implements(textUnmarshalerType) || p.Implements(textUnmarshalerType) {
		err = i.(encoding.TextUnmarshaler).UnmarshalText(data)
	}
	if !NotEmpty(i) {
		return nil, nil
	}
	return i, err
}

func JSONUnmarshalToItem(data []byte) Item {
	if len(data) == 0 {
		return nil
	}
	if ItemTyperFunc == nil {
		return nil
	}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return nil
	}
	var i Item
	switch val.Type() {
	case fastjson.TypeArray:
		items := make(ItemCollection, 0)
		for _, v :=  range val.GetArray() {
			var it Item
			it, err = itemFn(v.GetStringBytes())
			if it != nil && err == nil {
				items.Append(it)
			}

		}
		if len(items) == 1 {
			i = items.First()
		}
		if len(items) > 1 {
			i = items
		}
	case fastjson.TypeObject:
		i, err = itemFn(data)
	case fastjson.TypeString:
		if iri, ok := asIRI(data); ok {
			// try to see if it's an IRI
			i = iri
		}
	}
	if err != nil {
		return nil
	}
	return i
}

func asIRI(val []byte) (IRI, bool) {
	if len(val) == 0 {
		return NilIRI, true
	}
	val = bytes.Trim(val, string('"'))
	u, err := url.ParseRequestURI(string(val))
	if err == nil && len(u.Scheme) > 0 && len(u.Host) > 0 {
		// try to see if it's an IRI
		return IRI(val), true
	}
	return EmptyIRI, false
}

func JSONGetItem(data []byte, prop string) Item {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return nil
	}
	val = val.Get(prop)
	if val == nil {
		return nil
	}
	switch val.Type() {
	case fastjson.TypeString:
		if i, ok := asIRI(val.GetStringBytes()); ok {
			// try to see if it's an IRI
			return i
		}
	case fastjson.TypeArray:
		return JSONGetItems(data, prop)
	case fastjson.TypeObject:
		return JSONUnmarshalToItem(val.GetStringBytes())
	case fastjson.TypeNumber:
		fallthrough
	case fastjson.TypeNull:
		fallthrough
	default:
		return nil
	}
	return nil
}

func JSONGetURIItem(data []byte, prop string) Item {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return nil
	}
	v := val.Get(prop)
	if v == nil {
		return nil
	}

	switch v.Type() {
	case fastjson.TypeObject:
		return JSONGetItem(data, prop)
	case fastjson.TypeArray:
		it := make(ItemCollection, 0)
		for _, ob := range v.GetArray() {
			value := ob.GetStringBytes()
			if i, ok := asIRI(value); ok {
				it.Append(i)
				continue
			}
			i, err := ItemTyperFunc(JSONGetType(value))
			if err != nil {
				continue
			}
			if err = i.(json.Unmarshaler).UnmarshalJSON(value); err != nil {
				continue
			}
			it.Append(i)

		}
		return it
	case fastjson.TypeString:
		return IRI(val.String())
	}

	return nil
}

func JSONGetItems(data []byte, prop string) ItemCollection {
	if len(data) == 0 {
		return nil
	}
	p := fastjson.Parser{}

	v, err := p.ParseBytes(data)
	if err != nil {
		return nil
	}

	v = v.Get(prop)
	if v == nil {
		return nil
	}

	it := make(ItemCollection, 0)
	switch v.Type() {
	case fastjson.TypeArray:
		val := v.GetArray(prop)
		if len(val) == 0 {
			return nil
		}
		for _, v := range val {
			if i, err := itemFn(v.GetStringBytes()); i != nil && err == nil {
				it.Append(i)
			}
		}
	case fastjson.TypeObject:
		if i := JSONGetItem(data, prop); i != nil {
			it.Append(i)
		}
	case fastjson.TypeString:
		if iri := v.GetStringBytes(); len(iri) > 0 {
			it.Append(IRI(iri))
		}
	}
	if len(it) == 0 {
		return nil
	}
	return it
}

func JSONGetLangRefField(data []byte, prop string) LangRef {
	val := fastjson.GetString(data, prop)
	return LangRef(val)
}

func JSONGetIRI(data []byte, prop string) IRI {
	val := fastjson.GetString(data, prop)
	return IRI(val)
}

// UnmarshalJSON tries to detect the type of the object in the json data and then outputs a matching
// ActivityStreams object, if possible
func UnmarshalJSON(data []byte) (Item, error) {
	if ItemTyperFunc == nil {
		ItemTyperFunc = GetItemByType
	}
	return JSONUnmarshalToItem(data), nil
}

func GetItemByType(typ ActivityVocabularyType) (Item, error) {
	switch typ {
	case ObjectType:
		return ObjectNew(typ), nil
	case LinkType:
		return &Link{Type: typ}, nil
	case ActivityType:
		return &Activity{Type: typ}, nil
	case IntransitiveActivityType:
		return &IntransitiveActivity{Type: typ}, nil
	case ActorType:
		return &Actor{Type: typ}, nil
	case CollectionType:
		return &Collection{Type: typ}, nil
	case OrderedCollectionType:
		return &OrderedCollection{Type: typ}, nil
	case CollectionPageType:
		return &CollectionPage{Type: typ}, nil
	case OrderedCollectionPageType:
		return &OrderedCollectionPage{Type: typ}, nil
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
		return &Place{Type: typ}, nil
	case ProfileType:
		return &Profile{Type: typ}, nil
	case RelationshipType:
		return &Relationship{Type: typ}, nil
	case TombstoneType:
		return &Tombstone{Type: typ}, nil
	case VideoType:
		return ObjectNew(typ), nil
	case MentionType:
		return &Mention{Type: typ}, nil
	case ApplicationType:
		return &Application{Type: typ}, nil
	case GroupType:
		return &Group{Type: typ}, nil
	case OrganizationType:
		return &Organization{Type: typ}, nil
	case PersonType:
		return &Person{Type: typ}, nil
	case ServiceType:
		return &Service{Type: typ}, nil
	case AcceptType:
		return &Accept{Type: typ}, nil
	case AddType:
		return &Add{Type: typ}, nil
	case AnnounceType:
		return &Announce{Type: typ}, nil
	case ArriveType:
		return &Arrive{Type: typ}, nil
	case BlockType:
		return &Block{Type: typ}, nil
	case CreateType:
		return &Create{Type: typ}, nil
	case DeleteType:
		return &Delete{Type: typ}, nil
	case DislikeType:
		return &Dislike{Type: typ}, nil
	case FlagType:
		return &Flag{Type: typ}, nil
	case FollowType:
		return &Follow{Type: typ}, nil
	case IgnoreType:
		return &Ignore{Type: typ}, nil
	case InviteType:
		return &Invite{Type: typ}, nil
	case JoinType:
		return &Join{Type: typ}, nil
	case LeaveType:
		return &Leave{Type: typ}, nil
	case LikeType:
		return &Like{Type: typ}, nil
	case ListenType:
		return &Listen{Type: typ}, nil
	case MoveType:
		return &Move{Type: typ}, nil
	case OfferType:
		return &Offer{Type: typ}, nil
	case QuestionType:
		return &Question{Type: typ}, nil
	case RejectType:
		return &Reject{Type: typ}, nil
	case ReadType:
		return &Read{Type: typ}, nil
	case RemoveType:
		return &Remove{Type: typ}, nil
	case TentativeRejectType:
		return &TentativeReject{Type: typ}, nil
	case TentativeAcceptType:
		return &TentativeAccept{Type: typ}, nil
	case TravelType:
		return &Travel{Type: typ}, nil
	case UndoType:
		return &Undo{Type: typ}, nil
	case UpdateType:
		return &Update{Type: typ}, nil
	case ViewType:
		return &View{Type: typ}, nil
	case "":
		// when no type is available use a plain Object
		return &Object{}, nil
	}
	return nil, fmt.Errorf("empty ActivityStreams type")
}

func JSONGetActorEndpoints(data []byte, prop string) *Endpoints {
	str := fastjson.GetBytes(data, prop)

	var e *Endpoints
	if len(str) > 0 {
		e = &Endpoints{}
		e.UnmarshalJSON(str)
	}

	return e
}

func loadObject(data []byte, o *Object) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = GetItemByType
	}
	o.ID = JSONGetID(data)
	o.Type = JSONGetType(data)
	o.Name = JSONGetNaturalLanguageField(data, "name")
	o.Content = JSONGetNaturalLanguageField(data, "content")
	o.Summary = JSONGetNaturalLanguageField(data, "summary")
	o.Context = JSONGetItem(data, "context")
	o.URL = JSONGetURIItem(data, "url")
	o.MediaType = JSONGetMimeType(data, "mediaType")
	o.Generator = JSONGetItem(data, "generator")
	o.AttributedTo = JSONGetItem(data, "attributedTo")
	o.Attachment = JSONGetItem(data, "attachment")
	o.Location = JSONGetItem(data, "location")
	o.Published = JSONGetTime(data, "published")
	o.StartTime = JSONGetTime(data, "startTime")
	o.EndTime = JSONGetTime(data, "endTime")
	o.Duration = JSONGetDuration(data, "duration")
	o.Icon = JSONGetItem(data, "icon")
	o.Preview = JSONGetItem(data, "preview")
	o.Image = JSONGetItem(data, "image")
	o.Updated = JSONGetTime(data, "updated")
	o.InReplyTo = JSONGetItem(data, "inReplyTo")
	o.To = JSONGetItems(data, "to")
	o.Audience = JSONGetItems(data, "audience")
	o.Bto = JSONGetItems(data, "bto")
	o.CC = JSONGetItems(data, "cc")
	o.BCC = JSONGetItems(data, "bcc")
	o.Replies = JSONGetItem(data, "replies")
	o.Tag = JSONGetItems(data, "tag")
	o.Likes = JSONGetItem(data, "likes")
	o.Shares = JSONGetItem(data, "shares")
	o.Source = GetAPSource(data)
	return nil
}

func loadIntransitiveActivity(data []byte, i *IntransitiveActivity) error {
	OnObject(i, func(o *Object) error {
		return loadObject(data, o)
	})
	i.Actor = JSONGetItem(data, "actor")
	i.Target = JSONGetItem(data, "target")
	i.Result = JSONGetItem(data, "result")
	i.Origin = JSONGetItem(data, "origin")
	i.Instrument = JSONGetItem(data, "instrument")
	return nil
}

func loadActivity(data []byte, a *Activity) error {
	OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		return loadIntransitiveActivity(data, i)
	})
	a.Object = JSONGetItem(data, "object")
	return nil
}

func loadQuestion(data []byte, q *Question) error {
	OnIntransitiveActivity(q, func(i *IntransitiveActivity) error {
		return loadIntransitiveActivity(data, i)
	})
	q.OneOf = JSONGetItem(data, "oneOf")
	q.AnyOf = JSONGetItem(data, "anyOf")
	q.Closed = JSONGetBoolean(data, "closed")
	return nil
}

func loadActor(data []byte, a *Actor) error {
	OnObject(a, func(o *Object) error {
		return loadObject(data, o)
	})
	a.PreferredUsername = JSONGetNaturalLanguageField(data, "preferredUsername")
	a.Followers = JSONGetItem(data, "followers")
	a.Following = JSONGetItem(data, "following")
	a.Inbox = JSONGetItem(data, "inbox")
	a.Outbox = JSONGetItem(data, "outbox")
	a.Liked = JSONGetItem(data, "liked")
	a.Endpoints = JSONGetActorEndpoints(data, "endpoints")
	a.Streams = JSONGetStreams(data, "streams")
	a.PublicKey = JSONGetPublicKey(data, "publicKey")
	return nil
}

func loadCollection(data []byte, c *Collection) error {
	OnObject(c, func(o *Object) error {
		return loadObject(data, o)
	})
	c.Current = JSONGetItem(data, "current")
	c.First = JSONGetItem(data, "first")
	c.Last = JSONGetItem(data, "last")
	c.TotalItems = uint(JSONGetInt(data, "totalItems"))
	c.Items = JSONGetItems(data, "items")
	return nil
}

func loadCollectionPage(data []byte, c *CollectionPage) error {
	OnCollection(c, func(c *Collection) error {
		return loadCollection(data, c)
	})
	c.Next = JSONGetItem(data, "next")
	c.Prev = JSONGetItem(data, "prev")
	c.PartOf = JSONGetItem(data, "partOf")
	return nil
}

func loadOrderedCollection(data []byte, c *OrderedCollection) error {
	OnObject(c, func(o *Object) error {
		return loadObject(data, o)
	})
	c.Current = JSONGetItem(data, "current")
	c.First = JSONGetItem(data, "first")
	c.Last = JSONGetItem(data, "last")
	c.TotalItems = uint(JSONGetInt(data, "totalItems"))
	c.OrderedItems = JSONGetItems(data, "orderedItems")
	return nil
}

func loadOrderedCollectionPage(data []byte, c *OrderedCollectionPage) error {
	OnOrderedCollection(c, func(c *OrderedCollection) error {
		return loadOrderedCollection(data, c)
	})
	c.Next = JSONGetItem(data, "next")
	c.Prev = JSONGetItem(data, "prev")
	c.PartOf = JSONGetItem(data, "partOf")
	c.StartIndex = uint(JSONGetInt(data, "startIndex"))
	return nil
}

func loadPlace(data []byte, p *Place) error {
	OnObject(p, func(o *Object) error {
		return loadObject(data, o)
	})
	p.Accuracy = JSONGetFloat(data, "accuracy")
	p.Altitude = JSONGetFloat(data, "altitude")
	p.Latitude = JSONGetFloat(data, "latitude")
	p.Longitude = JSONGetFloat(data, "longitude")
	p.Radius = JSONGetInt(data, "radius")
	p.Units = JSONGetString(data, "units")
	return nil
}

func loadProfile(data []byte, p *Profile) error {
	OnObject(p, func(o *Object) error {
		return loadObject(data, o)
	})
	p.Describes = JSONGetItem(data, "describes")
	return nil
}

func loadRelationship(data []byte, r *Relationship) error {
	OnObject(r, func(o *Object) error {
		return loadObject(data, o)
	})
	r.Subject = JSONGetItem(data, "subject")
	r.Object = JSONGetItem(data, "object")
	r.Relationship = JSONGetItem(data, "relationship")
	return nil
}

func loadTombstone(data []byte, t *Tombstone) error {
	OnObject(t, func(o *Object) error {
		return loadObject(data, o)
	})
	t.FormerType = ActivityVocabularyType(JSONGetString(data, "formerType"))
	t.Deleted = JSONGetTime(data, "deleted")
	return nil
}

func loadLink(data []byte, l *Link) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = GetItemByType
	}
	l.ID = JSONGetID(data)
	l.Type = JSONGetType(data)
	l.MediaType = JSONGetMimeType(data, "mediaType")
	l.Preview = JSONGetItem(data, "preview")
	h := JSONGetInt(data, "height")
	if h != 0 {
		l.Height = uint(h)
	}
	w := JSONGetInt(data, "width")
	if w != 0 {
		l.Width = uint(w)
	}
	l.Name = JSONGetNaturalLanguageField(data, "name")
	hrefLang := JSONGetLangRefField(data, "hrefLang")
	if len(hrefLang) > 0 {
		l.HrefLang = hrefLang
	}
	href := JSONGetURIItem(data, "href")
	if href != nil {
		ll := href.GetLink()
		if len(ll) > 0 {
			l.Href = ll
		}
	}
	rel := JSONGetURIItem(data, "rel")
	if rel != nil {
		rr := rel.GetLink()
		if len(rr) > 0 {
			l.Rel = rr
		}
	}
	return nil
}
