package build

import (
	"time"

	vocab "github.com/go-ap/activitypub"
)

type InitFn = any

func Lang[T ~string](l vocab.LangRef) func(T) vocab.LangRefValue {
	return func(t T) vocab.LangRefValue {
		return vocab.LangRefValue{Ref: l, Value: []byte(t)}
	}
}

func LangMap(v ...vocab.LangRefValue) vocab.NaturalLanguageValues {
	res := make(vocab.NaturalLanguageValues)
	for _, vv := range v {
		res[vv.Ref] = vv.Value
	}
	return res
}

var (
	EN = Lang[string](vocab.English)
	FR = Lang[string](vocab.French)
)

func Object(initFn ...InitFn) *vocab.Object {
	ob := new(vocab.Object)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		}
	}
	return ob
}

func types[T ~string | vocab.ActivityVocabularyTypes](cc ...T) vocab.Typer {
	result := make(vocab.ActivityVocabularyTypes, 0, len(cc))
	for _, c := range cc {
		switch v := any(c).(type) {
		case string:
			if v != "" {
				result = append(result, vocab.ActivityVocabularyType(v))
			}
		case []byte:
			result = append(result, vocab.ActivityVocabularyType(v))
		case vocab.ActivityVocabularyType:
			result = append(result, v)
		case vocab.ActivityVocabularyTypes:
			result = append(result, v...)
		}
	}
	if len(result) == 1 {
		return result[0]
	}
	return result
}

func WithFormerType(t ...vocab.ActivityVocabularyType) vocab.WithTombstoneFn {
	typ := types(t...)
	return func(ob *vocab.Tombstone) error {
		ob.FormerType = typ
		return nil
	}
}

func Tombstone(initFn ...InitFn) *vocab.Tombstone {
	initFn = append([]InitFn{WithType(vocab.TombstoneType)}, initFn...)
	ob := vocab.Tombstone{}
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(&ob, fn)
		case vocab.WithTombstoneFn:
			_ = vocab.OnTombstone(&ob, fn)
		}
	}
	return &ob
}

func WithAccuracy(acc float64) vocab.WithPlaceFn {
	return func(pl *vocab.Place) error {
		pl.Accuracy = acc
		return nil
	}
}

func WithAltitude(alt float64) vocab.WithPlaceFn {
	return func(pl *vocab.Place) error {
		pl.Altitude = alt
		return nil
	}
}

func WithLatitude(lat float64) vocab.WithPlaceFn {
	return func(pl *vocab.Place) error {
		pl.Latitude = lat
		return nil
	}
}

func WithLongitude(lng float64) vocab.WithPlaceFn {
	return func(pl *vocab.Place) error {
		pl.Longitude = lng
		return nil
	}
}

func WithRadius(rad int64) vocab.WithPlaceFn {
	return func(pl *vocab.Place) error {
		pl.Radius = rad
		return nil
	}
}

func WithUnits(u string) vocab.WithPlaceFn {
	return func(pl *vocab.Place) error {
		pl.Units = u
		return nil
	}
}

func Place(initFn ...InitFn) *vocab.Place {
	initFn = append([]InitFn{WithType(vocab.PlaceType)}, initFn...)

	ob := new(vocab.Place)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithPlaceFn:
			_ = vocab.OnPlace(ob, fn)
		}
	}
	return ob
}

func WithDescribes(p ...vocab.Item) vocab.WithProfileFn {
	return func(pr *vocab.Profile) error {
		pr.Describes = vocab.ItemCollection(p).Normalize()
		return nil
	}
}

func Profile(initFn ...InitFn) *vocab.Profile {
	initFn = append([]InitFn{WithType(vocab.ProfileType)}, initFn...)

	ob := new(vocab.Profile)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithProfileFn:
			_ = vocab.OnProfile(ob, fn)
		}
	}
	return ob
}

func WithRelationship(rr ...vocab.Item) vocab.WithRelationshipFn {
	return func(rel *vocab.Relationship) error {
		rel.Relationship = vocab.ItemCollection(rr).Normalize()
		return nil
	}
}
func Relationship(initFn ...InitFn) *vocab.Relationship {
	initFn = append([]InitFn{WithType(vocab.RelationshipType)}, initFn...)

	ob := new(vocab.Relationship)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithRelationshipFn:
			_ = vocab.OnRelationship(ob, fn)
		}
	}
	return ob
}

func Video(initFn ...InitFn) *vocab.Video {
	initFn = append([]InitFn{WithType(vocab.VideoType)}, initFn...)
	return Object(initFn...)
}

func Audio(initFn ...InitFn) *vocab.Audio {
	initFn = append([]InitFn{WithType(vocab.AudioType)}, initFn...)
	return Object(initFn...)
}

func Image(initFn ...InitFn) *vocab.Image {
	initFn = append([]InitFn{WithType(vocab.ImageType)}, initFn...)
	return Object(initFn...)
}

func Document(initFn ...InitFn) *vocab.Document {
	initFn = append([]InitFn{WithType(vocab.DocumentType)}, initFn...)
	return Object(initFn...)
}

func Note(initFn ...InitFn) *vocab.Note {
	initFn = append([]InitFn{WithType(vocab.NoteType)}, initFn...)
	return Object(initFn...)
}

func Article(initFn ...InitFn) *vocab.Article {
	initFn = append([]InitFn{WithType(vocab.ArticleType)}, initFn...)
	return Object(initFn...)
}

func Event(initFn ...InitFn) *vocab.Event {
	initFn = append([]InitFn{WithType(vocab.EventType)}, initFn...)
	return Object(initFn...)
}

func Page(initFn ...InitFn) *vocab.Page {
	initFn = append([]InitFn{WithType(vocab.PageType)}, initFn...)
	return Object(initFn...)
}

func WithID(i vocab.IRI) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		ob.ID = i
		return nil
	}
}

func WithAttributedTo(i ...vocab.Item) vocab.WithObjectFn {
	ato := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.AttributedTo = ato.Normalize()
		return nil
	}
}

func WithContext(i ...vocab.Item) vocab.WithObjectFn {
	c := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.Context = c.Normalize()
		return nil
	}
}

func WithAudience(i ...vocab.Item) vocab.WithObjectFn {
	aud := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.Audience = aud
		return nil
	}
}

func WithGenerator(i ...vocab.Item) vocab.WithObjectFn {
	g := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.Generator = g.Normalize()
		return nil
	}
}

func WithURL(i ...vocab.Item) vocab.WithObjectFn {
	u := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.URL = u.Normalize()
		return nil
	}
}

func WithIcon(i ...vocab.Item) vocab.WithObjectFn {
	ato := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.Icon = ato.Normalize()
		return nil
	}
}

func WithImage(i ...vocab.Item) vocab.WithObjectFn {
	ato := vocab.ItemCollection(i)
	return func(ob *vocab.Object) error {
		ob.Image = ato.Normalize()
		return nil
	}
}

func WithAttachment(a ...vocab.Item) vocab.WithObjectFn {
	at := vocab.ItemCollection(a)
	return func(ob *vocab.Object) error {
		ob.Attachment = at.Normalize()
		return nil
	}
}

func WithTag(t ...vocab.Item) vocab.WithObjectFn {
	tags := vocab.ItemCollection(t)
	return func(ob *vocab.Object) error {
		ob.Tag = tags.Normalize()
		return nil
	}
}

func WithTo(i ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		return ob.To.Append(i...)
	}
}

func WithBto(i ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		return ob.Bto.Append(i...)
	}
}

func WithCC(i ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		return ob.CC.Append(i...)
	}
}

func WithBCC(i ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		return ob.BCC.Append(i...)
	}
}

func argNLV[T ~string | []byte | vocab.LangRefValue | vocab.NaturalLanguageValues](c T) vocab.NaturalLanguageValues {
	var result vocab.NaturalLanguageValues
	switch v := any(c).(type) {
	case string:
		if len(v) > 0 {
			result = vocab.DefaultLangValue(v)
		}
	case []byte:
		if len(v) > 0 {
			result = vocab.DefaultLangValue(string(v))
		}
	case vocab.LangRefValue:
		result[v.Ref] = v.Value
	case vocab.NaturalLanguageValues:
		result = v
	}
	return result
}

func WithName[T ~string | []byte | vocab.LangRefValue | vocab.NaturalLanguageValues](c T) vocab.WithObjectFn {
	v := argNLV(c)
	return func(ob *vocab.Object) error {
		ob.Name = v
		return nil
	}
}

func WithSummary[T ~string | []byte | vocab.LangRefValue | vocab.NaturalLanguageValues](c T) vocab.WithObjectFn {
	v := argNLV(c)
	return func(ob *vocab.Object) error {
		ob.Summary = v
		return nil
	}
}

func WithContent[T ~string | []byte | vocab.LangRefValue | vocab.NaturalLanguageValues](c T) vocab.WithObjectFn {
	v := argNLV(c)
	return func(ob *vocab.Object) error {
		ob.Content = v
		return nil
	}
}

func WithSource(s vocab.Source) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		ob.Source = s
		return nil
	}
}

func WithSourceContent[T ~string | []byte | vocab.LangRefValue | vocab.NaturalLanguageValues](c T) vocab.WithObjectFn {
	v := argNLV(c)
	return func(ob *vocab.Object) error {
		ob.Source.Content = v
		return nil
	}
}

func WithSourceMediaType(m string) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		ob.Source.MediaType = vocab.MimeType(m)
		return nil
	}
}

func WithMediaType[T ~string](m T) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		ob.MediaType = vocab.MimeType(m)
		return nil
	}
}

func argTime[T string | time.Time](s T) time.Time {
	switch ss := any(s).(type) {
	case string:
		p, _ := time.Parse(time.RFC3339Nano, ss)
		return p
	case time.Time:
		return ss
	}
	return time.Time{}
}

func WithPublished[T string | time.Time](s T) vocab.WithObjectFn {
	p := argTime(s)
	return func(ob *vocab.Object) error {
		ob.Published = p
		return nil
	}
}

func WithUpdated[T string | time.Time](s T) vocab.WithObjectFn {
	u := argTime(s)
	return func(ob *vocab.Object) error {
		ob.Updated = u
		return nil
	}
}

func WithReplies(r ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		if r == nil && !vocab.EmptyID.Equal(ob.ID) {
			r = vocab.ItemCollection{vocab.Replies.IRI(ob)}
		}
		ob.Replies = vocab.ItemCollection(r).Normalize()
		return nil
	}
}

func WithLikes(l ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		if l == nil && !vocab.EmptyID.Equal(ob.ID) {
			l = vocab.ItemCollection{vocab.Likes.IRI(ob)}
		}
		ob.Likes = vocab.ItemCollection(l).Normalize()
		return nil
	}
}

func WithShares(s ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		if s == nil && !vocab.EmptyID.Equal(ob.ID) {
			s = vocab.ItemCollection{vocab.Shares.IRI(ob)}
		}
		ob.Shares = vocab.ItemCollection(s).Normalize()
		return nil
	}
}

func WithType(t ...vocab.ActivityVocabularyType) vocab.WithObjectFn {
	typ := types(t...)
	return func(ob *vocab.Object) error {
		ob.Type = typ
		return nil
	}
}

func WithPreview(c ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		ob.Preview = vocab.ItemCollection(c).Normalize()
		return nil
	}
}

func WithInReplyTo(c ...vocab.Item) vocab.WithObjectFn {
	return func(ob *vocab.Object) error {
		ob.InReplyTo = vocab.ItemCollection(c).Normalize()
		return nil
	}
}
