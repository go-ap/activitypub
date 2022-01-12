package activitypub

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"time"
)

func unmapActorProperties(mm map[string][]byte, a *Actor) error {
	return errors.New("TODO unmapActorProperties")
}
func unmapIntransitiveActivityProperties(mm map[string][]byte, act *IntransitiveActivity) error {
	return errors.New("TODO unmapIntransitiveActivityProperties")
}
func unmapActivityProperties(mm map[string][]byte, act *Activity) error {
	return errors.New("TODO unmapActivityProperties")
}
func unmapLinkProperties(mm map[string][]byte, l *Link) error {
	return errors.New("TODO unmapLinkProperties")
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
	iri := IRI("")
	if err := tryDecodeIRI(&iri, data); err == nil {
		return iri, err
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return nil, err
	}
	typ := ObjectType
	sTyp, isObject := mm["type"]
	if isObject {
		typ = ActivityVocabularyType(sTyp)
	} else {
		_, isObject = mm["id"]
	}
	if isObject {
		it, err := GetItemByType(typ)
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
				return fmt.Errorf("TODO: Implement decode of %T", c)
			})
		case OrderedCollectionType:
			err = OnOrderedCollection(it, func(c *OrderedCollection) error {
				return fmt.Errorf("TODO: Implement decode of %T", c)
			})
		case CollectionPageType:
			err = OnCollectionPage(it, func(p *CollectionPage) error {
				return fmt.Errorf("TODO: Implement decode of %T", p)
			})
		case OrderedCollectionPageType:
			err = OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
				return fmt.Errorf("TODO: Implement decode of %T", p)
			})
		case PlaceType:
			err = OnPlace(it, func(p *Place) error {
				return fmt.Errorf("TODO: Implement decode of %T", p)
			})
		case ProfileType:
			err = OnProfile(it, func(p *Profile) error {
				return fmt.Errorf("TODO: Implement decode of %T", p)
			})
		case RelationshipType:
			err = OnRelationship(it, func(r *Relationship) error {
				return fmt.Errorf("TODO: Implement decode of %T", r)
			})
		case TombstoneType:
			err = OnTombstone(it, func(t *Tombstone) error {
				return fmt.Errorf("TODO: Implement decode of %T", t)
			})
		case QuestionType:
			err = OnQuestion(it, func(q *Question) error {
				return fmt.Errorf("TODO: Implement decode of %T", q)
			})
		}
		return it, err
	}

	return nil, errors.New("unable to decode to any known ActivityPub types")
}

func gobDecodeObjectAsMap(data []byte) (map[string][]byte, error) {
	mm := make(map[string][]byte)
	g := gob.NewDecoder(bytes.NewReader(data))
	if err := g.Decode(&mm); err != nil {
		return nil, err
	}
	return mm, nil
}
