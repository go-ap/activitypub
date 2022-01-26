package activitypub

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"
)

func GobDecode(data []byte) (Item, error) {
	return gobDecodeItem(data)
}

func gobDecodeUint(i *uint, data []byte) error {
	g := gob.NewDecoder(bytes.NewReader(data))
	return g.Decode(i)
}

func gobDecodeFloat64(f *float64, data []byte) error {
	g := gob.NewDecoder(bytes.NewReader(data))
	return g.Decode(f)
}

func gobDecodeInt64(i *int64, data []byte) error {
	g := gob.NewDecoder(bytes.NewReader(data))
	return g.Decode(i)
}

func gobDecodeBool(b *bool, data []byte) error {
	g := gob.NewDecoder(bytes.NewReader(data))
	return g.Decode(b)
}

func unmapActorProperties(mm map[string][]byte, a *Actor) error {
	err := OnObject(a, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["inbox"]; ok {
		if a.Inbox, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["outbox"]; ok {
		if a.Outbox, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["following"]; ok {
		if a.Following, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["followers"]; ok {
		if a.Followers, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["liked"]; ok {
		if a.Liked, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["preferredUsername"]; ok {
		if a.PreferredUsername, err = gobDecodeNaturalLanguageValues(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["endpoints"]; ok {
		if err = a.Endpoints.GobDecode(raw); err != nil {
			return err
		}
	}
	//if raw, ok := mm["streams"]; ok {
	//	if err = a.Streams.GobDecode(raw); err != nil {
	//		return err
	//	}
	//}
	if raw, ok := mm["publicKey"]; ok {
		if err = a.PublicKey.GobDecode(raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapIntransitiveActivityProperties(mm map[string][]byte, act *IntransitiveActivity) error {
	err := OnObject(act, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["actor"]; ok {
		if act.Actor, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["target"]; ok {
		if act.Target, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["result"]; ok {
		if act.Result, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["origin"]; ok {
		if act.Origin, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["instrument"]; ok {
		if act.Instrument, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapActivityProperties(mm map[string][]byte, act *Activity) error {
	err := OnIntransitiveActivity(act, func(act *IntransitiveActivity) error {
		return unmapIntransitiveActivityProperties(mm, act)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["object"]; ok {
		if act.Object, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapLinkProperties(mm map[string][]byte, l *Link) error {
	if raw, ok := mm["id"]; ok {
		if err := l.ID.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["type"]; ok {
		if err := l.Type.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["mediaType"]; ok {
		if err := l.MediaType.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["href"]; ok {
		if err := l.Href.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["hrefLang"]; ok {
		if err := l.HrefLang.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["name"]; ok {
		if err := l.Name.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["rel"]; ok {
		if err := l.Rel.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["width"]; ok {
		if err := gobDecodeUint(&l.Width, raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["height"]; ok {
		if err := gobDecodeUint(&l.Height, raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapObjectProperties(mm map[string][]byte, o *Object) error {
	var err error
	if raw, ok := mm["id"]; ok {
		if err = o.ID.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["type"]; ok {
		if err = o.Type.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["name"]; ok {
		if err = o.Name.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["attachment"]; ok {
		if o.Attachment, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["attributedTo"]; ok {
		if o.AttributedTo, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["audience"]; ok {
		if o.Audience, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["content"]; ok {
		if o.Content, err = gobDecodeNaturalLanguageValues(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["context"]; ok {
		if o.Context, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["mediaType"]; ok {
		if err = o.MediaType.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["endTime"]; ok {
		if err = o.EndTime.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["generator"]; ok {
		if o.Generator, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["icon"]; ok {
		if o.Icon, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["image"]; ok {
		if o.Image, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["inReplyTo"]; ok {
		if o.InReplyTo, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["location"]; ok {
		if o.Location, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["preview"]; ok {
		if o.Preview, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["published"]; ok {
		if err = o.Published.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["replies"]; ok {
		if o.Replies, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["startTime"]; ok {
		if err = o.StartTime.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["summary"]; ok {
		if o.Summary, err = gobDecodeNaturalLanguageValues(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["tag"]; ok {
		if o.Tag, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["updated"]; ok {
		if err = o.Updated.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["url"]; ok {
		if o.URL, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["to"]; ok {
		if o.To, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["bto"]; ok {
		if o.Bto, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["cc"]; ok {
		if o.CC, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["bcc"]; ok {
		if o.BCC, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["duration"]; ok {
		if o.Duration, err = gobDecodeDuration(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["likes"]; ok {
		if o.Likes, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["shares"]; ok {
		if o.Shares, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["source"]; ok {
		if err := o.Source.GobDecode(raw); err != nil {
			return err
		}
	}
	return nil
}

func tryDecodeObject(ob *Object, data []byte) error {
	if err := ob.GobDecode(data); err != nil {
		return err
	}
	return nil
}

func tryDecodeItems(items *ItemCollection, data []byte) error {
	tt := make([][]byte, 0)
	g := gob.NewDecoder(bytes.NewReader(data))
	if err := g.Decode(&tt); err != nil {
		return err
	}
	for _, it := range tt {
		ob, err := gobDecodeItem(it)
		if err != nil {
			return err
		}
		*items = append(*items, ob)
	}
	return nil
}

func tryDecodeIRIs(iris *IRIs, data []byte) error {
	return iris.GobDecode(data)
}

func tryDecodeIRI(iri *IRI, data []byte) error {
	return iri.GobDecode(data)
}

func gobDecodeDuration(data []byte) (time.Duration, error) {
	var d time.Duration
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&d)
	return d, err
}

func gobDecodeNaturalLanguageValues(data []byte) (NaturalLanguageValues, error) {
	n := make(NaturalLanguageValues, 0)
	err := n.GobDecode(data)
	return n, err
}

func gobDecodeItems(data []byte) (ItemCollection, error) {
	items := make(ItemCollection, 0)
	if err := tryDecodeItems(&items, data); err != nil {
		return nil, err
	}
	return items, nil
}

func gobDecodeItem(data []byte) (Item, error) {
	items := make(ItemCollection, 0)
	if err := tryDecodeItems(&items, data); err == nil {
		return items, nil
	}
	iris := make(IRIs, 0)
	if err := tryDecodeIRIs(&iris, data); err == nil {
		it := make(ItemCollection, len(iris))
		for i, iri := range iris {
			it[i] = iri
		}
		return it, nil
	}
	isObject := false
	typ := ObjectType
	mm, err := gobDecodeObjectAsMap(data)
	if err == nil {
		var sTyp []byte
		sTyp, isObject = mm["type"]
		if isObject {
			typ = ActivityVocabularyType(sTyp)
		} else {
			_, isObject = mm["id"]
		}
	}
	if isObject {
		it, err := ItemTyperFunc(typ)
		if err != nil {
			return nil, err
		}
		switch it.GetType() {
		case IRIType:
		case "", ObjectType, ArticleType, AudioType, DocumentType, EventType, ImageType, NoteType, PageType, VideoType:
			err = OnObject(it, func(ob *Object) error {
				return unmapObjectProperties(mm, ob)
			})
		case LinkType, MentionType:
			err = OnLink(it, func(l *Link) error {
				return unmapLinkProperties(mm, l)
			})
		case ActivityType, AcceptType, AddType, AnnounceType, BlockType, CreateType, DeleteType, DislikeType,
			FlagType, FollowType, IgnoreType, InviteType, JoinType, LeaveType, LikeType, ListenType, MoveType, OfferType,
			RejectType, ReadType, RemoveType, TentativeRejectType, TentativeAcceptType, UndoType, UpdateType, ViewType:
			err = OnActivity(it, func(act *Activity) error {
				return unmapActivityProperties(mm, act)
			})
		case IntransitiveActivityType, ArriveType, TravelType:
			err = OnIntransitiveActivity(it, func(act *IntransitiveActivity) error {
				return unmapIntransitiveActivityProperties(mm, act)
			})
		case ActorType, ApplicationType, GroupType, OrganizationType, PersonType, ServiceType:
			err = OnActor(it, func(a *Actor) error {
				return unmapActorProperties(mm, a)
			})
		case CollectionType:
			err = OnCollection(it, func(c *Collection) error {
				return unmapCollectionProperties(mm, c)
			})
		case OrderedCollectionType:
			err = OnOrderedCollection(it, func(c *OrderedCollection) error {
				return unmapOrderedCollectionProperties(mm, c)
			})
		case CollectionPageType:
			err = OnCollectionPage(it, func(p *CollectionPage) error {
				return unmapCollectionPageProperties(mm, p)
			})
		case OrderedCollectionPageType:
			err = OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
				return unmapOrderedCollectionPageProperties(mm, p)
			})
		case PlaceType:
			err = OnPlace(it, func(p *Place) error {
				return unmapPlaceProperties(mm, p)
			})
		case ProfileType:
			err = OnProfile(it, func(p *Profile) error {
				return unmapProfileProperties(mm, p)
			})
		case RelationshipType:
			err = OnRelationship(it, func(r *Relationship) error {
				return unmapRelationshipProperties(mm, r)
			})
		case TombstoneType:
			err = OnTombstone(it, func(t *Tombstone) error {
				return unmapTombstoneProperties(mm, t)
			})
		case QuestionType:
			err = OnQuestion(it, func(q *Question) error {
				return unmapQuestionProperties(mm, q)
			})
		}
		return it, err
	}
	iri := IRI("")
	if err := tryDecodeIRI(&iri, data); err == nil {
		return iri, err
	}

	return nil, errors.New("unable to gob decode to any known ActivityPub types")
}

func gobDecodeObjectAsMap(data []byte) (map[string][]byte, error) {
	mm := make(map[string][]byte)
	g := gob.NewDecoder(bytes.NewReader(data))
	if err := g.Decode(&mm); err != nil {
		return nil, err
	}
	return mm, nil
}

func unmapIncompleteCollectionProperties(mm map[string][]byte, c *Collection) error {
	err := OnObject(c, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["current"]; ok {
		if c.Current, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["first"]; ok {
		if c.First, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["last"]; ok {
		if c.Last, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["totalItems"]; ok {
		if err = gobDecodeUint(&c.TotalItems, raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapCollectionProperties(mm map[string][]byte, c *Collection) error {
	err := unmapIncompleteCollectionProperties(mm, c)
	if err != nil {
		return err
	}
	if raw, ok := mm["items"]; ok {
		if c.Items, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	return err
}

func unmapCollectionPageProperties(mm map[string][]byte, c *CollectionPage) error {
	err := OnCollection(c, func(c *Collection) error {
		return unmapCollectionProperties(mm, c)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["partOf"]; ok {
		if c.PartOf, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["next"]; ok {
		if c.Next, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["prev"]; ok {
		if c.Prev, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	return err
}

func unmapOrderedCollectionProperties(mm map[string][]byte, o *OrderedCollection) error {
	err := OnCollection(o, func(c *Collection) error {
		return unmapIncompleteCollectionProperties(mm, c)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["orderedItems"]; ok {
		if o.OrderedItems, err = gobDecodeItems(raw); err != nil {
			return err
		}
	}
	return err
}

func unmapOrderedCollectionPageProperties(mm map[string][]byte, c *OrderedCollectionPage) error {
	err := OnOrderedCollection(c, func(c *OrderedCollection) error {
		return unmapOrderedCollectionProperties(mm, c)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["partOf"]; ok {
		if c.PartOf, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["next"]; ok {
		if c.Next, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["prev"]; ok {
		if c.Prev, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	return err
}

func unmapPlaceProperties(mm map[string][]byte, p *Place) error {
	err := OnObject(p, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["accuracy"]; ok {
		if err = gobDecodeFloat64(&p.Accuracy, raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["altitude"]; ok {
		if err = gobDecodeFloat64(&p.Altitude, raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["latitude"]; ok {
		if err = gobDecodeFloat64(&p.Latitude, raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["radius"]; ok {
		if err = gobDecodeInt64(&p.Radius, raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["units"]; ok {
		p.Units = string(raw)
	}
	return nil
}

func unmapProfileProperties(mm map[string][]byte, p *Profile) error {
	err := OnObject(p, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["Describes"]; ok {
		if p.Describes, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapRelationshipProperties(mm map[string][]byte, r *Relationship) error {
	err := OnObject(r, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["subject"]; ok {
		if r.Subject, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["object"]; ok {
		if r.Object, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["relationship"]; ok {
		if r.Relationship, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapTombstoneProperties(mm map[string][]byte, t *Tombstone) error {
	err := OnObject(t, func(ob *Object) error {
		return unmapObjectProperties(mm, ob)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["formerType"]; ok {
		if err = t.FormerType.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["deleted"]; ok {
		if err = t.Deleted.GobDecode(raw); err != nil {
			return err
		}
	}
	return nil
}

func unmapQuestionProperties(mm map[string][]byte, q *Question) error {
	err := OnIntransitiveActivity(q, func(act *IntransitiveActivity) error {
		return unmapIntransitiveActivityProperties(mm, act)
	})
	if err != nil {
		return err
	}
	if raw, ok := mm["oneOf"]; ok {
		if q.OneOf, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["anyOf"]; ok {
		if q.AnyOf, err = gobDecodeItem(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["closed"]; ok {
		if err = gobDecodeBool(&q.Closed, raw); err != nil {
			return err
		}
	}
	return nil
}
