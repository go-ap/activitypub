package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

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

func gobEncodeItem(it Item) ([]byte, error) {
	b := bytes.Buffer{}
	var err error
	if IsIRI(it) {
		g := gob.NewEncoder(&b)
		err = gobEncodeStringLikeType(g, []byte(it.GetLink()))
	}
	if IsItemCollection(it) {
		g := gob.NewEncoder(&b)
		tt := make([][]byte, 0)
		err = OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range col.Collection() {
				single, _ := gobEncodeItem(it)
				tt = append(tt, single)
			}
			return nil
		})
		err = g.Encode(tt)
	}
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
			return fmt.Errorf("TODO: Implement encode of %T", c)
		})
	case OrderedCollectionType:
		err = OnOrderedCollection(it, func(c *OrderedCollection) error {
			return fmt.Errorf("TODO: Implement encode of %T", c)
		})
	case CollectionPageType:
		err = OnCollectionPage(it, func(p *CollectionPage) error {
			return fmt.Errorf("TODO: Implement encode of %T", p)
		})
	case OrderedCollectionPageType:
		err = OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
			return fmt.Errorf("TODO: Implement encode of %T", p)
		})
	case PlaceType:
		err = OnPlace(it, func(p *Place) error {
			return fmt.Errorf("TODO: Implement encode of %T", p)
		})
	case ProfileType:
		err = OnProfile(it, func(p *Profile) error {
			return fmt.Errorf("TODO: Implement encode of %T", p)
		})
	case RelationshipType:
		err = OnRelationship(it, func(r *Relationship) error {
			return fmt.Errorf("TODO: Implement encode of %T", r)
		})
	case TombstoneType:
		err = OnTombstone(it, func(t *Tombstone) error {
			return fmt.Errorf("TODO: Implement encode of %T", t)
		})
	case QuestionType:
		err = OnQuestion(it, func(q *Question) error {
			return fmt.Errorf("TODO: Implement encode of %T", q)
		})
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
