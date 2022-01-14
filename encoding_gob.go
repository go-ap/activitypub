package activitypub

import (
	"bytes"
	"encoding/gob"
)

func GobEncode(it Item) ([]byte, error) {
	return gobEncodeItem(it)
}

// TODO(marius): when migrating to go1.18, use a numeric constraint for this
func gobEncodeInt64(i int64) ([]byte, error) {
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gg.Encode(i); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// TODO(marius): when migrating to go1.18, use a numeric constraint for this
func gobEncodeUint(i uint) ([]byte, error) {
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gg.Encode(i); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func gobEncodeFloat64(f float64) ([]byte, error) {
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gg.Encode(f); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func gobEncodeBool(t bool) ([]byte, error) {
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gg.Encode(t); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func gobEncodeBytes(s []byte) ([]byte, error) {
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	if err := gg.Encode(s); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func gobEncodeStringLikeType(g *gob.Encoder, s []byte) error {
	if err := g.Encode(s); err != nil {
		return err
	}
	return nil
}

func gobEncodeItems(col ItemCollection) ([]byte, error) {
	b := bytes.Buffer{}
	tt := make([][]byte, 0)
	for _, it := range col.Collection() {
		single, err := gobEncodeItem(it)
		if err != nil {
			return nil, err
		}
		tt = append(tt, single)
	}
	err := gob.NewEncoder(&b).Encode(tt)
	return b.Bytes(), err
}

func gobEncodeItem(it Item) ([]byte, error) {
	if IsIRI(it) {
		if i, ok := it.(IRI); ok {
			return []byte(i), nil
		}
		return []byte{}, nil
	}
	b := bytes.Buffer{}
	var err error
	if IsItemCollection(it) {
		err = OnItemCollection(it, func(col *ItemCollection) error {
			bytes, err := gobEncodeItems(*col)
			b.Write(bytes)
			return err
		})
	}
	if IsObject(it) {
		switch it.GetType() {
		case IRIType:
			var bytes []byte
			bytes, err = it.(IRI).GobEncode()
			b.Write(bytes)
		case "", ObjectType, ArticleType, AudioType, DocumentType, EventType, ImageType, NoteType, PageType, VideoType:
			err = OnObject(it, func(ob *Object) error {
				bytes, err := ob.GobEncode()
				b.Write(bytes)
				return err
			})
		case LinkType, MentionType:
			err = OnLink(it, func(l *Link) error {
				bytes, err := l.GobEncode()
				b.Write(bytes)
				return err
			})
		case ActivityType, AcceptType, AddType, AnnounceType, BlockType, CreateType, DeleteType, DislikeType,
			FlagType, FollowType, IgnoreType, InviteType, JoinType, LeaveType, LikeType, ListenType, MoveType, OfferType,
			RejectType, ReadType, RemoveType, TentativeRejectType, TentativeAcceptType, UndoType, UpdateType, ViewType:
			err = OnActivity(it, func(act *Activity) error {
				bytes, err := act.GobEncode()
				b.Write(bytes)
				return err
			})
		case IntransitiveActivityType, ArriveType, TravelType:
			err = OnIntransitiveActivity(it, func(act *IntransitiveActivity) error {
				bytes, err := act.GobEncode()
				b.Write(bytes)
				return err
			})
		case ActorType, ApplicationType, GroupType, OrganizationType, PersonType, ServiceType:
			err = OnActor(it, func(a *Actor) error {
				bytes, err := a.GobEncode()
				b.Write(bytes)
				return err
			})
		case CollectionType:
			err = OnCollection(it, func(c *Collection) error {
				bytes, err := c.GobEncode()
				b.Write(bytes)
				return err
			})
		case OrderedCollectionType:
			err = OnOrderedCollection(it, func(c *OrderedCollection) error {
				bytes, err := c.GobEncode()
				b.Write(bytes)
				return err
			})
		case CollectionPageType:
			err = OnCollectionPage(it, func(p *CollectionPage) error {
				bytes, err := p.GobEncode()
				b.Write(bytes)
				return err
			})
		case OrderedCollectionPageType:
			err = OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
				bytes, err := p.GobEncode()
				b.Write(bytes)
				return err
			})
		case PlaceType:
			err = OnPlace(it, func(p *Place) error {
				bytes, err := p.GobEncode()
				b.Write(bytes)
				return err
			})
		case ProfileType:
			err = OnProfile(it, func(p *Profile) error {
				bytes, err := p.GobEncode()
				b.Write(bytes)
				return err
			})
		case RelationshipType:
			err = OnRelationship(it, func(r *Relationship) error {
				bytes, err := r.GobEncode()
				b.Write(bytes)
				return err
			})
		case TombstoneType:
			err = OnTombstone(it, func(t *Tombstone) error {
				bytes, err := t.GobEncode()
				b.Write(bytes)
				return err
			})
		case QuestionType:
			err = OnQuestion(it, func(q *Question) error {
				bytes, err := q.GobEncode()
				b.Write(bytes)
				return err
			})
		}
	}
	return b.Bytes(), err
}

func mapObjectProperties(mm map[string][]byte, o *Object) (hasData bool, err error) {
	if len(o.ID) > 0 {
		if mm["id"], err = o.ID.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Type) > 0 {
		if mm["type"], err = o.Type.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.MediaType) > 0 {
		if mm["mediaType"], err = o.MediaType.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Name) > 0 {
		if mm["name"], err = o.Name.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Attachment != nil {
		if mm["attachment"], err = gobEncodeItem(o.Attachment); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.AttributedTo != nil {
		if mm["attributedTo"], err = gobEncodeItem(o.AttributedTo); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Audience != nil {
		if mm["audience"], err = gobEncodeItem(o.Audience); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Content != nil {
		if mm["content"], err = o.Content.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Context != nil {
		if mm["context"], err = gobEncodeItem(o.Context); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.MediaType) > 0 {
		if mm["mediaType"], err = o.MediaType.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.EndTime.IsZero() {
		if mm["endTime"], err = o.EndTime.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Generator != nil {
		if mm["generator"], err = gobEncodeItem(o.Generator); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Icon != nil {
		if mm["icon"], err = gobEncodeItem(o.Icon); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Image != nil {
		if mm["image"], err = gobEncodeItem(o.Image); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.InReplyTo != nil {
		if mm["inReplyTo"], err = gobEncodeItem(o.InReplyTo); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Location != nil {
		if mm["location"], err = gobEncodeItem(o.Location); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Preview != nil {
		if mm["preview"], err = gobEncodeItem(o.Preview); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.Published.IsZero() {
		if mm["published"], err = o.Published.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Replies != nil {
		if mm["replies"], err = gobEncodeItem(o.Replies); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.StartTime.IsZero() {
		if mm["startTime"], err = o.StartTime.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Summary) > 0 {
		if mm["summary"], err = o.Summary.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Tag != nil {
		if mm["tag"], err = gobEncodeItem(o.Tag); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.Updated.IsZero() {
		if mm["updated"], err = o.Updated.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Tag != nil {
		if mm["tag"], err = gobEncodeItem(o.Tag); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !o.Updated.IsZero() {
		if mm["updated"], err = o.Updated.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.URL != nil {
		if mm["url"], err = gobEncodeItem(o.URL); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.To != nil {
		if mm["to"], err = gobEncodeItem(o.To); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Bto != nil {
		if mm["bto"], err = gobEncodeItem(o.Bto); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.CC != nil {
		if mm["cc"], err = gobEncodeItem(o.CC); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.BCC != nil {
		if mm["bcc"], err = gobEncodeItem(o.BCC); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Duration > 0 {
		if mm["duration"], err = gobEncodeInt64(int64(o.Duration)); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Likes != nil {
		if mm["likes"], err = gobEncodeItem(o.Likes); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Shares != nil {
		if mm["shares"], err = gobEncodeItem(o.Shares); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if o.Shares != nil {
		if mm["shares"], err = gobEncodeItem(o.Shares); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(o.Source.MediaType)+len(o.Source.Content) > 0 {
		if mm["source"], err = o.Source.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}

	return hasData, nil
}

func mapActorProperties(mm map[string][]byte, a *Actor) (hasData bool, err error) {
	err = OnObject(a, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if a.Inbox != nil {
		if mm["inbox"], err = gobEncodeItem(a.Inbox); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Inbox != nil {
		if mm["inbox"], err = gobEncodeItem(a.Inbox); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Outbox != nil {
		if mm["outbox"], err = gobEncodeItem(a.Outbox); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Following != nil {
		if mm["following"], err = gobEncodeItem(a.Following); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Followers != nil {
		if mm["followers"], err = gobEncodeItem(a.Followers); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Liked != nil {
		if mm["liked"], err = gobEncodeItem(a.Liked); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(a.PreferredUsername) > 0 {
		if mm["preferredUsername"], err = a.PreferredUsername.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if a.Endpoints != nil {
		if mm["endpoints"], err = a.Endpoints.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if len(a.Streams) > 0 {
		//if mm["streams"], err = gobDecodeItem(a.Streams); err != nil {
		//	return hasData, err
		//}
		//hasData = true
	}
	if len(a.PublicKey.PublicKeyPem)+len(a.PublicKey.ID) > 0 {
		if mm["publicKey"], err = a.PublicKey.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return hasData, err
}

func mapIncompleteCollectionProperties(mm map[string][]byte, c Collection) (hasData bool, err error) {
	err = OnObject(c, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if c.Current != nil {
		if mm["current"], err = gobEncodeItem(c.Current); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.First != nil {
		if mm["first"], err = gobEncodeItem(c.First); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.Last != nil {
		if mm["last"], err = gobEncodeItem(c.Last); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.TotalItems > 0 {
		hasData = true
	}
	if mm["totalItems"], err = gobEncodeUint(c.TotalItems); err != nil {
		return hasData, err
	}
	return
}

func mapCollectionProperties(mm map[string][]byte, c Collection) (hasData bool, err error) {
	hasData, err = mapIncompleteCollectionProperties(mm, c)
	if err != nil {
		return hasData, err
	}
	if c.Items != nil {
		if mm["items"], err = gobEncodeItems(c.Items); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return
}

func mapOrderedCollectionProperties(mm map[string][]byte, c OrderedCollection) (hasData bool, err error) {
	err = OnCollection(c, func(c *Collection) error {
		hasData, err = mapIncompleteCollectionProperties(mm, *c)
		return err
	})
	if c.OrderedItems != nil {
		if mm["orderedItems"], err = gobEncodeItems(c.OrderedItems); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return
}

func mapCollectionPageProperties(mm map[string][]byte, c CollectionPage) (hasData bool, err error) {
	err = OnCollection(c, func(c *Collection) error {
		hasData, err = mapCollectionProperties(mm, *c)
		return err
	})
	if c.PartOf != nil {
		if mm["partOf"], err = gobEncodeItem(c.PartOf); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.Next != nil {
		if mm["next"], err = gobEncodeItem(c.Next); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.Prev != nil {
		if mm["prev"], err = gobEncodeItem(c.Prev); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return
}

func mapOrderedCollectionPageProperties(mm map[string][]byte, c OrderedCollectionPage) (hasData bool, err error) {
	err = OnOrderedCollection(c, func(c *OrderedCollection) error {
		hasData, err = mapOrderedCollectionProperties(mm, *c)
		return err
	})
	if c.PartOf != nil {
		if mm["partOf"], err = gobEncodeItem(c.PartOf); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.Next != nil {
		if mm["next"], err = gobEncodeItem(c.Next); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if c.Prev != nil {
		if mm["prev"], err = gobEncodeItem(c.Prev); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return
}

func mapLinkProperties(mm map[string][]byte, l Link) (hasData bool, err error) {
	if len(l.ID) > 0 {
		if mm["id"], err = l.ID.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if len(l.Type) > 0 {
		if mm["type"], err = l.Type.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if len(l.MediaType) > 0 {
		if mm["mediaType"], err = l.MediaType.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if len(l.Href) > 0 {
		if mm["href"], err = l.Href.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if len(l.HrefLang) > 0 {
		if mm["hrefLang"], err = l.HrefLang.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if len(l.Name) > 0 {
		if mm["name"], err = l.Name.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if len(l.Rel) > 0 {
		if mm["rel"], err = l.Rel.GobEncode(); err != nil {
			return
		}
		hasData = true
	}
	if l.Width > 0 {
		if mm["width"], err = gobEncodeUint(l.Width); err != nil {
			return
		}
		hasData = true
	}
	if l.Height > 0 {
		if mm["height"], err = gobEncodeUint(l.Height); err != nil {
			return
		}
		hasData = true
	}
	return
}

func mapPlaceProperties(mm map[string][]byte, p Place) (hasData bool, err error) {
	err = OnObject(p, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if p.Accuracy > 0 {
		if mm["accuracy"], err = gobEncodeFloat64(p.Accuracy); err != nil {
			return
		}
		hasData = true
	}
	if p.Altitude > 0 {
		if mm["altitude"], err = gobEncodeFloat64(p.Altitude); err != nil {
			return
		}
		hasData = true
	}
	if p.Latitude > 0 {
		if mm["latitude"], err = gobEncodeFloat64(p.Latitude); err != nil {
			return
		}
		hasData = true
	}
	if p.Longitude > 0 {
		if mm["longitude"], err = gobEncodeFloat64(p.Longitude); err != nil {
			return
		}
		hasData = true
	}
	if p.Radius > 0 {
		if mm["radius"], err = gobEncodeInt64(p.Radius); err != nil {
			return
		}
		hasData = true
	}
	if len(p.Units) > 0 {
		if mm["units"], err = gobEncodeBytes([]byte(p.Units)); err != nil {
			return
		}
		hasData = true
	}
	return
}

func mapProfileProperties(mm map[string][]byte, p Profile) (hasData bool, err error) {
	err = OnObject(p, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if p.Describes != nil {
		if mm["describes"], err = gobEncodeItem(p.Describes); err != nil {
			return
		}
		hasData = true
	}
	return
}

func mapRelationshipProperties(mm map[string][]byte, r Relationship) (hasData bool, err error) {
	err = OnObject(r, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if r.Subject != nil {
		if mm["subject"], err = gobEncodeItem(r.Subject); err != nil {
			return
		}
		hasData = true
	}
	if r.Object != nil {
		if mm["object"], err = gobEncodeItem(r.Object); err != nil {
			return
		}
		hasData = true
	}
	if r.Relationship != nil {
		if mm["relationship"], err = gobEncodeItem(r.Relationship); err != nil {
			return
		}
		hasData = true
	}
	return
}

func mapTombstoneProperties(mm map[string][]byte, t Tombstone) (hasData bool, err error) {
	err = OnObject(t, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if len(t.FormerType) > 0 {
		if mm["formerType"], err = t.FormerType.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	if !t.Deleted.IsZero() {
		if mm["deleted"], err = t.Deleted.GobEncode(); err != nil {
			return hasData, err
		}
		hasData = true
	}
	return
}

func mapQuestionProperties(mm map[string][]byte, q Question) (hasData bool, err error) {
	err = OnObject(q, func(o *Object) error {
		hasData, err = mapObjectProperties(mm, o)
		return err
	})
	if q.OneOf != nil {
		if mm["oneOf"], err = gobEncodeItem(q.OneOf); err != nil {
			return
		}
		hasData = true
	}
	if q.AnyOf != nil {
		if mm["anyOf"], err = gobEncodeItem(q.AnyOf); err != nil {
			return
		}
		hasData = true
	}
	if q.Closed {
		hasData = true
	}
	if hasData {
		if mm["closed"], err = gobEncodeBool(q.Closed); err != nil {
			return
		}
	}
	return
}
