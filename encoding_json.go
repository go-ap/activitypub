package activitypub

import (
	"encoding/json"
	"fmt"
	"time"

	"git.sr.ht/~mariusor/go-xsd-duration"
	"github.com/go-ap/jsonld"
)

func JSONWriteComma(b *[]byte) {
	if len(*b) > 1 && (*b)[len(*b)-1] != ',' {
		*b = append(*b, ',')
	}
}

func JSONWriteProp(b *[]byte, name string, val []byte) (notEmpty bool) {
	if len(val) == 0 {
		return false
	}
	JSONWriteComma(b)
	success := JSONWritePropName(b, name) && JSONWriteValue(b, val)
	if !success {
		*b = (*b)[:len(*b)-1]
	}
	return success
}

func JSONWrite(b *[]byte, c ...byte) {
	*b = append(*b, c...)
}

func JSONWriteS(b *[]byte, s string) {
	*b = append(*b, s...)
}

func JSONWritePropName(b *[]byte, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	JSONWrite(b, '"')
	JSONWriteS(b, s)
	JSONWrite(b, '"', ':')
	return true
}

func JSONWriteValue(b *[]byte, s []byte) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	JSONWrite(b, s...)
	return true
}

func JSONWriteNaturalLanguageProp(b *[]byte, n string, nl NaturalLanguageValues) (notEmpty bool) {
	l := nl.Count()
	if l > 1 {
		n += "Map"
	}
	if v, err := nl.MarshalJSON(); err == nil && len(v) > 0 {
		return JSONWriteProp(b, n, v)
	}
	return false
}

func JSONWriteStringProp(b *[]byte, n string, s string) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf(`"%s"`, s)))
}

func JSONWriteBoolProp(b *[]byte, n string, t bool) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf(`"%t"`, t)))
}

func JSONWriteIntProp(b *[]byte, n string, d int64) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf("%d", d)))
}

func JSONWriteFloatProp(b *[]byte, n string, f float64) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf("%f", f)))
}

func JSONWriteTimeProp(b *[]byte, n string, t time.Time) (notEmpty bool) {
	var tb []byte
	JSONWrite(&tb, '"')
	JSONWriteS(&tb, t.UTC().Format(time.RFC3339))
	JSONWrite(&tb, '"')
	return JSONWriteProp(b, n, tb)
}

func JSONWriteDurationProp(b *[]byte, n string, d time.Duration) (notEmpty bool) {
	var tb []byte
	if v, err := xsd.Marshal(d); err == nil {
		JSONWrite(&tb, '"')
		JSONWrite(&tb, v...)
		JSONWrite(&tb, '"')
	}
	return JSONWriteProp(b, n, tb)
}

func JSONWriteIRIProp(b *[]byte, n string, i LinkOrIRI) (notEmpty bool) {
	url := i.GetLink().String()
	if len(url) == 0 {
		return false
	}
	JSONWriteStringProp(b, n, url)
	return true
}

func JSONWriteItemProp(b *[]byte, n string, i Item) (notEmpty bool) {
	if i == nil {
		return notEmpty
	}
	if im, ok := i.(json.Marshaler); ok {
		v, err := im.MarshalJSON()
		if err != nil {
			return false
		}
		return JSONWriteProp(b, n, v)
	}
	return notEmpty
}

func byteInsertAt(raw []byte, b byte, p int) []byte {
	return append(raw[:p], append([]byte{b}, raw[p:]...)...)
}

func escapeQuote(s string) string {
	raw := []byte(s)
	end := len(s)
	for i := 0; i < end; i++ {
		c := raw[i]
		if c == '"' && (i > 0 && s[i-1] != '\\') {
			raw = byteInsertAt(raw, '\\', i)
			i++
			end++
		}
	}
	return string(raw)
}

func JSONWriteStringValue(b *[]byte, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	JSONWrite(b, '"')
	JSONWriteS(b, escapeQuote(s))
	JSONWrite(b, '"')
	return true
}

func JSONWriteItemCollectionValue(b *[]byte, col ItemCollection, compact bool) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	if len(col) == 1 && compact {
		it := col[0]
		im, ok := it.(json.Marshaler)
		if !ok {
			return false
		}
		v, err := im.MarshalJSON()
		if err != nil {
			return false
		}
		if len(v) == 0 {
			return false
		}
		JSONWrite(b, v...)
		return true
	}
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			JSONWrite(b, ',')
		}
	}
	JSONWrite(b, '[')
	skipComma := true
	for _, it := range col {
		im, ok := it.(json.Marshaler)
		if !ok {
			continue
		}
		v, err := im.MarshalJSON()
		if err != nil {
			return false
		}
		if len(v) == 0 {
			continue
		}
		writeCommaIfNotEmpty(!skipComma)
		JSONWrite(b, v...)
		skipComma = false
	}
	JSONWrite(b, ']')
	return true
}

func JSONWriteItemCollectionProp(b *[]byte, n string, col ItemCollection, compact bool) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	JSONWriteComma(b)
	success := JSONWritePropName(b, n) && JSONWriteItemCollectionValue(b, col, compact)
	if !success {
		*b = (*b)[:len(*b)-1]
	}
	return success
}

func JSONWriteObjectValue(b *[]byte, o Object) (notEmpty bool) {
	if v, err := o.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "id", v)
	}
	if v, err := o.Type.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "type", v) || notEmpty
	}
	if v, err := o.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "mediaType", v) || notEmpty
	}
	if len(o.Name) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "name", o.Name) || notEmpty
	}
	if len(o.Summary) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "summary", o.Summary) || notEmpty
	}
	if len(o.Content) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "content", o.Content) || notEmpty
	}
	if o.Attachment != nil {
		notEmpty = JSONWriteItemProp(b, "attachment", o.Attachment) || notEmpty
	}
	if o.AttributedTo != nil {
		notEmpty = JSONWriteItemProp(b, "attributedTo", o.AttributedTo) || notEmpty
	}
	if o.Audience != nil {
		notEmpty = JSONWriteItemProp(b, "audience", o.Audience) || notEmpty
	}
	if o.Context != nil {
		notEmpty = JSONWriteItemProp(b, "context", o.Context) || notEmpty
	}
	if o.Generator != nil {
		notEmpty = JSONWriteItemProp(b, "generator", o.Generator) || notEmpty
	}
	if o.Icon != nil {
		notEmpty = JSONWriteItemProp(b, "icon", o.Icon) || notEmpty
	}
	if o.Image != nil {
		notEmpty = JSONWriteItemProp(b, "image", o.Image) || notEmpty
	}
	if o.InReplyTo != nil {
		notEmpty = JSONWriteItemProp(b, "inReplyTo", o.InReplyTo) || notEmpty
	}
	if o.Location != nil {
		notEmpty = JSONWriteItemProp(b, "location", o.Location) || notEmpty
	}
	if o.Preview != nil {
		notEmpty = JSONWriteItemProp(b, "preview", o.Preview) || notEmpty
	}
	if o.Replies != nil {
		notEmpty = JSONWriteItemProp(b, "replies", o.Replies) || notEmpty
	}
	if o.Tag != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "tag", o.Tag, false) || notEmpty
	}
	if o.URL != nil {
		notEmpty = JSONWriteItemProp(b, "url", o.URL) || notEmpty
	}
	if o.To != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "to", o.To, false) || notEmpty
	}
	if o.Bto != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "bto", o.Bto, false) || notEmpty
	}
	if o.CC != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "cc", o.CC, false) || notEmpty
	}
	if o.BCC != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "bcc", o.BCC, false) || notEmpty
	}
	if !o.Published.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "published", o.Published) || notEmpty
	}
	if !o.Updated.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "updated", o.Updated) || notEmpty
	}
	if !o.StartTime.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "startTime", o.StartTime) || notEmpty
	}
	if !o.EndTime.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "endTime", o.EndTime) || notEmpty
	}
	if o.Duration != 0 {
		// TODO(marius): maybe don't use 0 as a nil value for Object types
		//  which can have a valid duration of 0 - (Video, Audio, etc)
		notEmpty = JSONWriteDurationProp(b, "duration", o.Duration) || notEmpty
	}
	if o.Likes != nil {
		notEmpty = JSONWriteItemProp(b, "likes", o.Likes) || notEmpty
	}
	if o.Shares != nil {
		notEmpty = JSONWriteItemProp(b, "shares", o.Shares) || notEmpty
	}
	if v, err := o.Source.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "source", v) || notEmpty
	}
	return notEmpty
}

func JSONWriteActivityValue(b *[]byte, a Activity) (notEmpty bool) {
	_ = OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = JSONWriteIntransitiveActivityValue(b, *i) || notEmpty
		return nil
	})
	if a.Object != nil {
		notEmpty = JSONWriteItemProp(b, "object", a.Object) || notEmpty
	}
	return notEmpty
}

func JSONWriteIntransitiveActivityValue(b *[]byte, i IntransitiveActivity) (notEmpty bool) {
	_ = OnObject(i, func(o *Object) error {
		if o == nil {
			return nil
		}
		notEmpty = JSONWriteObjectValue(b, *o) || notEmpty
		return nil
	})
	if i.Actor != nil {
		notEmpty = JSONWriteItemProp(b, "actor", i.Actor) || notEmpty
	}
	if i.Target != nil {
		notEmpty = JSONWriteItemProp(b, "target", i.Target) || notEmpty
	}
	if i.Result != nil {
		notEmpty = JSONWriteItemProp(b, "result", i.Result) || notEmpty
	}
	if i.Origin != nil {
		notEmpty = JSONWriteItemProp(b, "origin", i.Origin) || notEmpty
	}
	if i.Instrument != nil {
		notEmpty = JSONWriteItemProp(b, "instrument", i.Instrument) || notEmpty
	}
	return notEmpty
}

func JSONWriteQuestionValue(b *[]byte, q Question) (notEmpty bool) {
	_ = OnIntransitiveActivity(q, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = JSONWriteIntransitiveActivityValue(b, *i) || notEmpty
		return nil
	})
	if q.OneOf != nil {
		notEmpty = JSONWriteItemProp(b, "oneOf", q.OneOf) || notEmpty
	}
	if q.AnyOf != nil {
		notEmpty = JSONWriteItemProp(b, "anyOf", q.AnyOf) || notEmpty
	}
	notEmpty = JSONWriteBoolProp(b, "closed", q.Closed) || notEmpty
	return notEmpty
}

func JSONWriteLinkValue(b *[]byte, l Link) (notEmpty bool) {
	if v, err := l.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "id", v)
	}
	if v, err := l.Type.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "type", v) || notEmpty
	}
	if v, err := l.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "mediaType", v) || notEmpty
	}
	if len(l.Name) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "name", l.Name) || notEmpty
	}
	if v, err := l.Rel.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "rel", v) || notEmpty
	}
	if l.Height > 0 {
		notEmpty = JSONWriteIntProp(b, "height", int64(l.Height))
	}
	if l.Width > 0 {
		notEmpty = JSONWriteIntProp(b, "width", int64(l.Width))
	}
	if l.Preview != nil {
		notEmpty = JSONWriteItemProp(b, "rel", l.Preview) || notEmpty
	}
	if v, err := l.Href.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "href", v) || notEmpty
	}
	if l.HrefLang.Valid() {
		notEmpty = JSONWriteStringProp(b, "hrefLang", l.HrefLang.String()) || notEmpty
	}
	return notEmpty
}

// MarshalJSON represents just a wrapper for the jsonld.Marshal function
func MarshalJSON(it LinkOrIRI) ([]byte, error) {
	return jsonld.Marshal(it)
}
