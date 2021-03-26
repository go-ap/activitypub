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

func writeJSONProp(b *[]byte, name string, val []byte) (notEmpty bool) {
	if len(val) == 0 {
		return false
	}
	writeComma(b)
	success := writePropJSONName(b, name) && writeJSONValue(b, val)
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

func writePropJSONName(b *[]byte, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	write(b, '"')
	writeS(b, s)
	write(b, '"', ':')
	return true
}

func writeJSONValue(b *[]byte, s []byte) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	write(b, s...)
	return true
}

func writeNaturalLanguageJSONProp(b *[]byte, n string, nl NaturalLanguageValues) (notEmpty bool) {
	l := nl.Count()
	if l > 1 {
		n += "Map"
	}
	if v, err := nl.MarshalJSON(); err == nil && len(v) > 0 {
		return writeJSONProp(b, n, v)
	}
	return false
}

func writeStringJSONProp(b *[]byte, n string, s string) (notEmpty bool) {
	return writeJSONProp(b, n, []byte(fmt.Sprintf(`"%s"`, s)))
}

func writeBoolJSONProp(b *[]byte, n string, t bool) (notEmpty bool) {
	return writeJSONProp(b, n, []byte(fmt.Sprintf(`"%t"`, t)))
}

func writeIntJSONProp(b *[]byte, n string, d int64) (notEmpty bool) {
	return writeJSONProp(b, n, []byte(fmt.Sprintf("%d", d)))
}

func writeFloatJSONProp(b *[]byte, n string, f float64) (notEmpty bool) {
	return writeJSONProp(b, n, []byte(fmt.Sprintf("%f", f)))
}

func writeTimeJSONProp(b *[]byte, n string, t time.Time) (notEmpty bool) {
	var tb []byte
	write(&tb, '"')
	writeS(&tb, t.UTC().Format(time.RFC3339))
	write(&tb, '"')
	return writeJSONProp(b, n, tb)
}

func writeDurationJSONProp(b *[]byte, n string, d time.Duration) (notEmpty bool) {
	if v, err := xsd.Marshal(d); err == nil {
		return writeJSONProp(b, n, v)
	}
	return false
}

func writeIRIJSONProp(b *[]byte, n string, i LinkOrIRI) (notEmpty bool) {
	url := i.GetLink().String()
	if len(url) == 0 {
		return false
	}
	writeStringJSONProp(b, n, url)
	return true
}

func writeItemJSONProp(b *[]byte, n string, i Item) (notEmpty bool) {
	if i == nil {
		return notEmpty
	}
	if im, ok := i.(json.Marshaler); ok {
		v, err := im.MarshalJSON()
		if err != nil {
			return false
		}
		return writeJSONProp(b, n, v)
	}
	return notEmpty
}

func writeStringJSONValue(b *[]byte, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	write(b, '"')
	writeS(b, s)
	write(b, '"')
	return true
}

func writeItemCollectionJSONValue(b *[]byte, col ItemCollection) (notEmpty bool) {
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

func writeItemCollectionJSONProp(b *[]byte, n string, col ItemCollection) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	writeComma(b)
	success := writePropJSONName(b, n) && writeItemCollectionJSONValue(b, col)
	if !success {
		*b = (*b)[:len(*b)-1]
	}
	return success
}

func writeObjectJSONValue(b *[]byte, o Object) (notEmpty bool) {
	if v, err := o.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "id", v) || notEmpty
	}
	if v, err := o.Type.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "type", v) || notEmpty
	}
	if v, err := o.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "mediaType", v) || notEmpty
	}
	if len(o.Name) > 0 {
		notEmpty = writeNaturalLanguageJSONProp(b, "name", o.Name) || notEmpty
	}
	if len(o.Summary) > 0 {
		notEmpty = writeNaturalLanguageJSONProp(b, "summary", o.Summary) || notEmpty
	}
	if len(o.Content) > 0 {
		notEmpty = writeNaturalLanguageJSONProp(b, "content", o.Content) || notEmpty
	}
	if o.Attachment != nil {
		notEmpty = writeItemJSONProp(b, "attachment", o.Attachment) || notEmpty
	}
	if o.AttributedTo != nil {
		notEmpty = writeItemJSONProp(b, "attributedTo", o.AttributedTo) || notEmpty
	}
	if o.Audience != nil {
		notEmpty = writeItemJSONProp(b, "audience", o.Audience) || notEmpty
	}
	if o.Context != nil {
		notEmpty = writeItemJSONProp(b, "context", o.Context) || notEmpty
	}
	if o.Generator != nil {
		notEmpty = writeItemJSONProp(b, "generator", o.Generator) || notEmpty
	}
	if o.Icon != nil {
		notEmpty = writeItemJSONProp(b, "icon", o.Icon) || notEmpty
	}
	if o.Image != nil {
		notEmpty = writeItemJSONProp(b, "image", o.Image) || notEmpty
	}
	if o.InReplyTo != nil {
		notEmpty = writeItemJSONProp(b, "inReplyTo", o.InReplyTo) || notEmpty
	}
	if o.Location != nil {
		notEmpty = writeItemJSONProp(b, "location", o.Location) || notEmpty
	}
	if o.Preview != nil {
		notEmpty = writeItemJSONProp(b, "preview", o.Preview) || notEmpty
	}
	if o.Replies != nil {
		notEmpty = writeItemJSONProp(b, "replies", o.Replies) || notEmpty
	}
	if o.Tag != nil {
		notEmpty = writeItemJSONProp(b, "tag", o.Tag) || notEmpty
	}
	if o.URL != nil {
		notEmpty = writeIRIJSONProp(b, "url", o.URL) || notEmpty
	}
	if o.To != nil {
		notEmpty = writeItemJSONProp(b, "to", o.To) || notEmpty
	}
	if o.Bto != nil {
		notEmpty = writeItemJSONProp(b, "bto", o.Bto) || notEmpty
	}
	if o.CC != nil {
		notEmpty = writeItemJSONProp(b, "cc", o.CC) || notEmpty
	}
	if o.BCC != nil {
		notEmpty = writeItemJSONProp(b, "bcc", o.BCC) || notEmpty
	}
	if !o.Published.IsZero() {
		notEmpty = writeTimeJSONProp(b, "published", o.Published) || notEmpty
	}
	if !o.Updated.IsZero() {
		notEmpty = writeTimeJSONProp(b, "updated", o.Updated) || notEmpty
	}
	if !o.StartTime.IsZero() {
		notEmpty = writeTimeJSONProp(b, "startTime", o.StartTime) || notEmpty
	}
	if !o.EndTime.IsZero() {
		notEmpty = writeTimeJSONProp(b, "endTime", o.EndTime) || notEmpty
	}
	if o.Duration != 0 {
		// TODO(marius): maybe don't use 0 as a nil value for Object types
		//  which can have a valid duration of 0 - (Video, Audio, etc)
		notEmpty = writeDurationJSONProp(b, "duration", o.Duration) || notEmpty
	}
	if o.Likes != nil {
		notEmpty = writeItemJSONProp(b, "likes", o.Likes) || notEmpty
	}
	if o.Shares != nil {
		notEmpty = writeItemJSONProp(b, "shares", o.Shares) || notEmpty
	}
	if v, err := o.Source.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "source", v) || notEmpty
	}
	return notEmpty
}

func writeActivityJSONValue(b *[]byte, a Activity) (notEmpty bool) {
	OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = writeIntransitiveActivityJSONValue(b, *i) || notEmpty
		return nil
	})
	if a.Object != nil {
		notEmpty = writeItemJSONProp(b, "object", a.Object) || notEmpty
	}
	return notEmpty
}

func writeIntransitiveActivityJSONValue(b *[]byte, i IntransitiveActivity) (notEmpty bool) {
	OnObject(i, func(o *Object) error {
		if o == nil {
			return nil
		}
		notEmpty = writeObjectJSONValue(b, *o) || notEmpty
		return nil
	})
	if i.Actor != nil {
		notEmpty = writeItemJSONProp(b, "actor", i.Actor) || notEmpty
	}
	if i.Target != nil {
		notEmpty = writeItemJSONProp(b, "target", i.Target) || notEmpty
	}
	if i.Result != nil {
		notEmpty = writeItemJSONProp(b, "result", i.Result) || notEmpty
	}
	if i.Origin != nil {
		notEmpty = writeItemJSONProp(b, "origin", i.Origin) || notEmpty
	}
	if i.Instrument != nil {
		notEmpty = writeItemJSONProp(b, "instrument", i.Instrument) || notEmpty
	}
	return notEmpty
}

func writeQuestionJSONValue(b *[]byte, q Question) (notEmpty bool) {
	OnIntransitiveActivity(q, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = writeIntransitiveActivityJSONValue(b, *i) || notEmpty
		return nil
	})
	if q.OneOf != nil {
		notEmpty = writeItemJSONProp(b, "oneOf", q.OneOf) || notEmpty
	} else if q.AnyOf != nil {
		notEmpty = writeItemJSONProp(b, "anyOf", q.OneOf) || notEmpty
	}
	notEmpty = writeBoolJSONProp(b, "closed", q.Closed) || notEmpty
	return notEmpty
}

func writeLinkJSONValue(b *[]byte, l Link) (notEmpty bool) {
	if v, err := l.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "id", v) || notEmpty
	}
	if v, err := l.Type.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "type", v) || notEmpty
	}
	if v, err := l.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "mediaType", v) || notEmpty
	}
	if len(l.Name) > 0 {
		notEmpty = writeNaturalLanguageJSONProp(b, "name", l.Name) || notEmpty
	}
	if v, err := l.Rel.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "rel", v) || notEmpty
	}
	if l.Height > 0 {
		notEmpty = writeIntJSONProp(b, "height", int64(l.Height))
	}
	if l.Width > 0 {
		notEmpty = writeIntJSONProp(b, "width", int64(l.Width))
	}
	if l.Preview != nil {
		notEmpty = writeItemJSONProp(b, "rel", l.Preview) || notEmpty
	}
	if v, err := l.Href.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeJSONProp(b, "href", v) || notEmpty
	}
	if len(l.HrefLang) > 0 {
		notEmpty = writeStringJSONProp(b, "hrefLang", string(l.HrefLang)) || notEmpty
	}
	return notEmpty
}

// MarshalJSON wraps the jsonld.Marshal function
func MarshalJSON(it Item) ([]byte, error) {
	return jsonld.Marshal(it)
}
