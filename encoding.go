package activitypub

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~mariusor/go-xsd-duration"
	"github.com/go-ap/jsonld"
	"time"
)

func writeComma(b *[]byte) {
	if len(*b) > 1 && (*b)[len(*b)-1] != ',' {
		*b = append(*b, ',')
	}
}
func writeProp(b *[]byte, name string, val []byte) (notEmpty bool) {
	if len(val) == 0 {
		return false
	}
	writeComma(b)
	success := writePropName(b, name) && writeValue(b, val)
	if !success {
		*b = (*b)[:len(*b)-1]
	}
	return success
}

func write(b *[]byte, c ...byte) {
	*b = append(*b, c...)
}

func writeS(b *[]byte, s string) {
	*b = append(*b, s...)
}

func writePropName(b *[]byte, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	write(b, '"')
	writeS(b, s)
	write(b, '"', ':')
	return true
}

func writeValue(b *[]byte, s []byte) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	write(b, s...)
	return true
}

func writeNaturalLanguageProp(b *[]byte, n string, nl NaturalLanguageValues) (notEmpty bool) {
	l := nl.Count()
	if l > 1 {
		n += "Map"
	}
	if v, err := nl.MarshalJSON(); err == nil && len(v) > 0 {
		return writeProp(b, n, v)
	}
	return false
}
func writeStringProp(b *[]byte, n string, s string) (notEmpty bool) {
	return writeProp(b, n, []byte(fmt.Sprintf(`"%s"`, s)))
}
func writeBoolProp(b *[]byte, n string, t bool) (notEmpty bool) {
	return writeProp(b, n, []byte(fmt.Sprintf(`"%t"`, t)))
}
func writeIntProp(b *[]byte, n string, d int64) (notEmpty bool) {
	return writeProp(b, n, []byte(fmt.Sprintf("%d", d)))
}
func writeFloatProp(b *[]byte, n string, f float64) (notEmpty bool) {
	return writeProp(b, n, []byte(fmt.Sprintf("%f", f)))
}
func writeTimeProp(b *[]byte, n string, t time.Time) (notEmpty bool) {
	if v, err := t.UTC().MarshalJSON(); err == nil {
		return writeProp(b, n, v)
	}
	return false
}

func writeDurationProp(b *[]byte, n string, d time.Duration) (notEmpty bool) {
	if v, err := xsd.Marshal(d); err == nil {
		return writeProp(b, n, v)
	}
	return false
}

func writeIRIProp(b *[]byte, n string, i LinkOrIRI) (notEmpty bool) {
	url := i.GetLink().String()
	if len(url) == 0 {
		return false
	}
	writeStringProp(b, n, url)
	return true
}

func writeItemProp(b *[]byte, n string, i Item) (notEmpty bool) {
	if i == nil {
		return notEmpty
	}
	if im, ok := i.(json.Marshaler); ok {
		v, err := im.MarshalJSON()
		if err != nil {
			return false
		}
		return writeProp(b, n, v)
	}
	return notEmpty
}

func writeStringValue(b *[]byte, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	write(b, '"')
	writeS(b, s)
	write(b, '"')
	return true
}

func writeItemCollectionValue(b *[]byte, col ItemCollection) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			write(b, ',')
		}
	}
	write(b, '[')
	for i, it := range col {
		if im, ok := it.(json.Marshaler); ok {
			v, err := im.MarshalJSON()
			if err != nil {
				return false
			}
			writeCommaIfNotEmpty(i > 0)
			write(b, v...)
		}
	}
	write(b, ']')
	return true
}
func writeItemCollectionProp(b *[]byte, n string, col ItemCollection) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	writeComma(b)
	success := writePropName(b, n) && writeItemCollectionValue(b, col)
	if !success {
		*b = (*b)[:len(*b)-1]
	}
	return success
}

func writeObjectValue(b *[]byte, o Object) (notEmpty bool) {
	if v, err := o.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "id", v) || notEmpty
	}
	if v, err := o.Type.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "type", v) || notEmpty
	}
	if v, err := o.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "mediaType", v) || notEmpty
	}
	if len(o.Name) > 0 {
		notEmpty = writeNaturalLanguageProp(b, "name", o.Name) || notEmpty
	}
	if len(o.Summary) > 0 {
		notEmpty = writeNaturalLanguageProp(b, "summary", o.Summary) || notEmpty
	}
	if len(o.Content) > 0 {
		notEmpty = writeNaturalLanguageProp(b, "content", o.Content) || notEmpty
	}
	if o.Attachment != nil {
		notEmpty = writeItemProp(b, "attachment", o.Attachment) || notEmpty
	}
	if o.AttributedTo != nil {
		notEmpty = writeItemProp(b, "attributedTo", o.AttributedTo) || notEmpty
	}
	if o.Audience != nil {
		notEmpty = writeItemProp(b, "audience", o.Audience) || notEmpty
	}
	if o.Context != nil {
		notEmpty = writeItemProp(b, "context", o.Context) || notEmpty
	}
	if o.Generator != nil {
		notEmpty = writeItemProp(b, "generator", o.Generator) || notEmpty
	}
	if o.Icon != nil {
		notEmpty = writeItemProp(b, "icon", o.Icon) || notEmpty
	}
	if o.Image != nil {
		notEmpty = writeItemProp(b, "image", o.Image) || notEmpty
	}
	if o.InReplyTo != nil {
		notEmpty = writeItemProp(b, "inReplyTo", o.InReplyTo) || notEmpty
	}
	if o.Location != nil {
		notEmpty = writeItemProp(b, "location", o.Location) || notEmpty
	}
	if o.Preview != nil {
		notEmpty = writeItemProp(b, "preview", o.Preview) || notEmpty
	}
	if o.Replies != nil {
		notEmpty = writeItemProp(b, "replies", o.Replies) || notEmpty
	}
	if o.Tag != nil {
		notEmpty = writeItemProp(b, "tag", o.Tag) || notEmpty
	}
	if o.URL != nil {
		notEmpty = writeIRIProp(b, "url", o.URL) || notEmpty
	}
	if o.To != nil {
		notEmpty = writeItemProp(b, "to", o.To) || notEmpty
	}
	if o.Bto != nil {
		notEmpty = writeItemProp(b, "bto", o.Bto) || notEmpty
	}
	if o.CC != nil {
		notEmpty = writeItemProp(b, "cc", o.CC) || notEmpty
	}
	if o.BCC != nil {
		notEmpty = writeItemProp(b, "bcc", o.BCC) || notEmpty
	}
	if !o.Published.IsZero() {
		notEmpty = writeTimeProp(b, "published", o.Published) || notEmpty
	}
	if !o.Updated.IsZero() {
		notEmpty = writeTimeProp(b, "updated", o.Updated) || notEmpty
	}
	if !o.StartTime.IsZero() {
		notEmpty = writeTimeProp(b, "startTime", o.StartTime) || notEmpty
	}
	if !o.EndTime.IsZero() {
		notEmpty = writeTimeProp(b, "endTime", o.EndTime) || notEmpty
	}
	if o.Duration != 0 {
		// TODO(marius): maybe don't use 0 as a nil value for Object types
		//  which can have a valid duration of 0 - (Video, Audio, etc)
		notEmpty = writeDurationProp(b, "duration", o.Duration) || notEmpty
	}
	if o.Likes != nil {
		notEmpty = writeItemProp(b, "likes", o.Likes) || notEmpty
	}
	if o.Shares != nil {
		notEmpty = writeItemProp(b, "shares", o.Shares) || notEmpty
	}
	if v, err := o.Source.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "source", v) || notEmpty
	}
	return notEmpty
}

func writeActivityValue(b *[]byte, a Activity) (notEmpty bool) {
	OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = writeIntransitiveActivityValue(b, *i) || notEmpty
		return nil
	})
	if a.Object != nil {
		notEmpty = writeItemProp(b, "object", a.Object) || notEmpty
	}
	return notEmpty
}

func writeIntransitiveActivityValue(b *[]byte, i IntransitiveActivity) (notEmpty bool) {
	OnObject(i, func(o *Object) error {
		if o == nil {
			return nil
		}
		notEmpty = writeObjectValue(b, *o) || notEmpty
		return nil
	})
	if i.Actor != nil {
		notEmpty = writeItemProp(b, "actor", i.Actor) || notEmpty
	}
	if i.Target != nil {
		notEmpty = writeItemProp(b, "target", i.Target) || notEmpty
	}
	if i.Result != nil {
		notEmpty = writeItemProp(b, "result", i.Result) || notEmpty
	}
	if i.Origin != nil {
		notEmpty = writeItemProp(b, "origin", i.Origin) || notEmpty
	}
	if i.Instrument != nil {
		notEmpty = writeItemProp(b, "instrument", i.Instrument) || notEmpty
	}
	return notEmpty
}

func writeQuestionValue(b *[]byte, q Question) (notEmpty bool) {
	OnIntransitiveActivity(q, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = writeIntransitiveActivityValue(b, *i) || notEmpty
		return nil
	})
	if q.OneOf != nil {
		notEmpty = writeItemProp(b, "oneOf", q.OneOf) || notEmpty
	} else if q.AnyOf != nil {
		notEmpty = writeItemProp(b, "anyOf", q.OneOf) || notEmpty
	}
	notEmpty = writeBoolProp(b, "closed", q.Closed) || notEmpty
	return notEmpty
}

func writeLinkValue(b *[]byte, l Link) (notEmpty bool) {
	if v, err := l.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "id", v) || notEmpty
	}
	if v, err := l.Type.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "type", v) || notEmpty
	}
	if v, err := l.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "mediaType", v) || notEmpty
	}
	if len(l.Name) > 0 {
		notEmpty = writeNaturalLanguageProp(b, "name", l.Name) || notEmpty
	}
	if v, err := l.Rel.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "rel", v) || notEmpty
	}
	if l.Height > 0 {
		notEmpty = writeIntProp(b, "height", int64(l.Height))
	}
	if l.Width > 0 {
		notEmpty = writeIntProp(b, "width", int64(l.Width))
	}
	if l.Preview != nil {
		notEmpty = writeItemProp(b, "rel", l.Preview) || notEmpty
	}
	if v, err := l.Href.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "href", v) || notEmpty
	}
	if len(l.HrefLang) > 0 {
		notEmpty = writeStringProp(b, "hrefLang", string(l.HrefLang)) || notEmpty
	}
	return notEmpty
}

// MarshalJSON wraps the jsonld.Marshal function
func MarshalJSON(it Item) ([]byte, error) {
	return jsonld.Marshal(it)
}
