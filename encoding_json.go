package activitypub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"git.sr.ht/~mariusor/go-xsd-duration"
	"github.com/go-ap/jsonld"
)

func JSONWriteComma(b *bytes.Buffer) {
	b.WriteRune(',')
}

func JSONWriteProp(b *bytes.Buffer, name string, val []byte, needsComma bool) (notEmpty bool) {
	if len(val) == 0 {
		return false
	}
	if needsComma {
		JSONWriteComma(b)
	}
	success := JSONWritePropName(b, name) && JSONWriteValue(b, val)
	//if !success {
	//	_ = b.UnreadByte()
	//}
	return success
}

func JSONWrite(b *bytes.Buffer, c ...byte) {
	b.Write(c)
}

func JSONWriteS(b *bytes.Buffer, s string) {
	b.WriteString(s)
}

func JSONWritePropName(b *bytes.Buffer, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	JSONWrite(b, '"')
	JSONWriteS(b, s)
	JSONWrite(b, '"', ':')
	return true
}

func JSONWriteValue(b *bytes.Buffer, s []byte) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	JSONWrite(b, s...)
	return true
}

func JSONWriteNaturalLanguageProp(b *bytes.Buffer, n string, nl NaturalLanguageValues, needsComma bool) (notEmpty bool) {
	l := nl.Count()
	if l > 1 {
		n += "Map"
	}
	if v, err := nl.MarshalJSON(); err == nil && len(v) > 0 {
		return JSONWriteProp(b, n, v, needsComma)
	}
	return false
}

func JSONWriteStringProp(b *bytes.Buffer, n string, s string, needsComma bool) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf(`"%s"`, s)), needsComma)
}

func JSONWriteBoolProp(b *bytes.Buffer, n string, t bool, needsComma bool) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf(`"%t"`, t)), needsComma)
}

func JSONWriteIntProp(b *bytes.Buffer, n string, d int64, needsComma bool) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf("%d", d)), needsComma)
}

func JSONWriteFloatProp(b *bytes.Buffer, n string, f float64, needsComma bool) (notEmpty bool) {
	return JSONWriteProp(b, n, []byte(fmt.Sprintf("%f", f)), needsComma)
}

func JSONWriteTimeProp(b *bytes.Buffer, n string, t time.Time, needsComma bool) (notEmpty bool) {
	tb := bytes.Buffer{}
	JSONWrite(&tb, '"')
	JSONWriteS(&tb, t.UTC().Format(time.RFC3339))
	JSONWrite(&tb, '"')
	return JSONWriteProp(b, n, tb.Bytes(), needsComma)
}

func JSONWriteDurationProp(b *bytes.Buffer, n string, d time.Duration, needsComma bool) (notEmpty bool) {
	tb := bytes.Buffer{}
	if v, err := xsd.Marshal(d); err == nil {
		JSONWrite(&tb, '"')
		JSONWrite(&tb, v...)
		JSONWrite(&tb, '"')
	}
	return JSONWriteProp(b, n, tb.Bytes(), needsComma)
}

func JSONWriteIRIProp(b *bytes.Buffer, n string, i LinkOrIRI, needsComma bool) (notEmpty bool) {
	url := i.GetLink().String()
	if len(url) == 0 {
		return false
	}
	JSONWriteStringProp(b, n, url, needsComma)
	return true
}

func JSONWriteItemProp(b *bytes.Buffer, n string, i Item, needsComma bool) (notEmpty bool) {
	if i == nil {
		return notEmpty
	}
	if im, ok := i.(json.Marshaler); ok {
		v, err := im.MarshalJSON()
		if err != nil {
			return false
		}
		return JSONWriteProp(b, n, v, needsComma)
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

func JSONWriteStringValue(b *bytes.Buffer, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	JSONWrite(b, '"')
	JSONWriteS(b, escapeQuote(s))
	JSONWrite(b, '"')
	return true
}

func JSONWriteItemCollectionValue(b *bytes.Buffer, col ItemCollection, compact bool) (notEmpty bool) {
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

func JSONWriteItemCollectionProp(b *bytes.Buffer, n string, col ItemCollection, compact, needsComma bool) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	if needsComma {
		JSONWriteComma(b)
	}
	success := JSONWritePropName(b, n) && JSONWriteItemCollectionValue(b, col, compact)
	if !success {
		_ = b.UnreadByte()
	}
	return success
}

func JSONWriteObjectValue(b *bytes.Buffer, o Object) (notEmpty bool) {
	if v, err := o.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "id", v, false)
	}
	if HasTypes(o) {
		notEmpty = JSONWriteTypes(b, "type", o.Type, notEmpty) || notEmpty
	}
	if v, err := o.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "mediaType", v, notEmpty) || notEmpty
	}
	if len(o.Name) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "name", o.Name, notEmpty) || notEmpty
	}
	if len(o.Summary) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "summary", o.Summary, notEmpty) || notEmpty
	}
	if len(o.Content) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "content", o.Content, notEmpty) || notEmpty
	}
	if o.Attachment != nil {
		notEmpty = JSONWriteItemProp(b, "attachment", o.Attachment, notEmpty) || notEmpty
	}
	if o.AttributedTo != nil {
		notEmpty = JSONWriteItemProp(b, "attributedTo", o.AttributedTo, notEmpty) || notEmpty
	}
	if o.Audience != nil {
		notEmpty = JSONWriteItemProp(b, "audience", o.Audience, notEmpty) || notEmpty
	}
	if o.Context != nil {
		notEmpty = JSONWriteItemProp(b, "context", o.Context, notEmpty) || notEmpty
	}
	if o.Generator != nil {
		notEmpty = JSONWriteItemProp(b, "generator", o.Generator, notEmpty) || notEmpty
	}
	if o.Icon != nil {
		notEmpty = JSONWriteItemProp(b, "icon", o.Icon, notEmpty) || notEmpty
	}
	if o.Image != nil {
		notEmpty = JSONWriteItemProp(b, "image", o.Image, notEmpty) || notEmpty
	}
	if o.InReplyTo != nil {
		notEmpty = JSONWriteItemProp(b, "inReplyTo", o.InReplyTo, notEmpty) || notEmpty
	}
	if o.Location != nil {
		notEmpty = JSONWriteItemProp(b, "location", o.Location, notEmpty) || notEmpty
	}
	if o.Preview != nil {
		notEmpty = JSONWriteItemProp(b, "preview", o.Preview, notEmpty) || notEmpty
	}
	if o.Replies != nil {
		notEmpty = JSONWriteItemProp(b, "replies", o.Replies, notEmpty) || notEmpty
	}
	if o.Tag != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "tag", o.Tag, false, notEmpty) || notEmpty
	}
	if o.URL != nil {
		notEmpty = JSONWriteItemProp(b, "url", o.URL, notEmpty) || notEmpty
	}
	if o.To != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "to", o.To, false, notEmpty) || notEmpty
	}
	if o.Bto != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "bto", o.Bto, false, notEmpty) || notEmpty
	}
	if o.CC != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "cc", o.CC, false, notEmpty) || notEmpty
	}
	if o.BCC != nil {
		notEmpty = JSONWriteItemCollectionProp(b, "bcc", o.BCC, false, notEmpty) || notEmpty
	}
	if !o.Published.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "published", o.Published, notEmpty) || notEmpty
	}
	if !o.Updated.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "updated", o.Updated, notEmpty) || notEmpty
	}
	if !o.StartTime.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "startTime", o.StartTime, notEmpty) || notEmpty
	}
	if !o.EndTime.IsZero() {
		notEmpty = JSONWriteTimeProp(b, "endTime", o.EndTime, notEmpty) || notEmpty
	}
	if o.Duration != 0 {
		// TODO(marius): maybe don't use 0 as a nil value for Object types
		//  which can have a valid duration of 0 - (Video, Audio, etc)
		notEmpty = JSONWriteDurationProp(b, "duration", o.Duration, notEmpty) || notEmpty
	}
	if o.Likes != nil {
		notEmpty = JSONWriteItemProp(b, "likes", o.Likes, notEmpty) || notEmpty
	}
	if o.Shares != nil {
		notEmpty = JSONWriteItemProp(b, "shares", o.Shares, notEmpty) || notEmpty
	}
	if v, err := o.Source.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "source", v, notEmpty) || notEmpty
	}
	return notEmpty
}

func JSONWriteActivityValue(b *bytes.Buffer, a Activity) (notEmpty bool) {
	_ = OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = JSONWriteIntransitiveActivityValue(b, *i) || notEmpty
		return nil
	})
	if a.Object != nil {
		notEmpty = JSONWriteItemProp(b, "object", a.Object, notEmpty) || notEmpty
	}
	return notEmpty
}

func JSONWriteIntransitiveActivityValue(b *bytes.Buffer, i IntransitiveActivity) (notEmpty bool) {
	_ = OnObject(i, func(o *Object) error {
		if o == nil {
			return nil
		}
		notEmpty = JSONWriteObjectValue(b, *o) || notEmpty
		return nil
	})
	if i.Actor != nil {
		notEmpty = JSONWriteItemProp(b, "actor", i.Actor, notEmpty) || notEmpty
	}
	if i.Target != nil {
		notEmpty = JSONWriteItemProp(b, "target", i.Target, notEmpty) || notEmpty
	}
	if i.Result != nil {
		notEmpty = JSONWriteItemProp(b, "result", i.Result, notEmpty) || notEmpty
	}
	if i.Origin != nil {
		notEmpty = JSONWriteItemProp(b, "origin", i.Origin, notEmpty) || notEmpty
	}
	if i.Instrument != nil {
		notEmpty = JSONWriteItemProp(b, "instrument", i.Instrument, notEmpty) || notEmpty
	}
	return notEmpty
}

func JSONWriteQuestionValue(b *bytes.Buffer, q Question) (notEmpty bool) {
	_ = OnIntransitiveActivity(q, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = JSONWriteIntransitiveActivityValue(b, *i) || notEmpty
		return nil
	})
	if q.OneOf != nil {
		notEmpty = JSONWriteItemProp(b, "oneOf", q.OneOf, notEmpty) || notEmpty
	}
	if q.AnyOf != nil {
		notEmpty = JSONWriteItemProp(b, "anyOf", q.AnyOf, notEmpty) || notEmpty
	}
	notEmpty = JSONWriteBoolProp(b, "closed", q.Closed, notEmpty) || notEmpty
	return notEmpty
}

func JSONWriteTypes(b *bytes.Buffer, n string, ty Typer, needsComma bool) (notEmpty bool) {
	typ := ty.AsTypes()
	if v, err := typ.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, n, v, needsComma)
	}
	return true
}

func JSONWriteLinkValue(b *bytes.Buffer, l Link) (notEmpty bool) {
	if v, err := l.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "id", v, false)
	}
	if HasTypes(l) {
		notEmpty = JSONWriteTypes(b, "type", l.Type, notEmpty) || notEmpty
	}
	if v, err := l.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "mediaType", v, notEmpty) || notEmpty
	}
	if len(l.Name) > 0 {
		notEmpty = JSONWriteNaturalLanguageProp(b, "name", l.Name, notEmpty) || notEmpty
	}
	if v, err := l.Rel.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "rel", v, notEmpty) || notEmpty
	}
	if l.Height > 0 {
		notEmpty = JSONWriteIntProp(b, "height", int64(l.Height), notEmpty) || notEmpty
	}
	if l.Width > 0 {
		notEmpty = JSONWriteIntProp(b, "width", int64(l.Width), notEmpty) || notEmpty
	}
	if l.Preview != nil {
		notEmpty = JSONWriteItemProp(b, "rel", l.Preview, notEmpty) || notEmpty
	}
	if v, err := l.Href.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(b, "href", v, notEmpty) || notEmpty
	}
	if l.HrefLang.Valid() {
		notEmpty = JSONWriteStringProp(b, "hrefLang", l.HrefLang.String(), notEmpty) || notEmpty
	}
	return notEmpty
}

func JSONWriteActivityVocabularyTypes(b *bytes.Buffer, t ActivityVocabularyTypes) (notEmpty bool) {
	if b == nil {
		return notEmpty
	}
	tLen := len(t)
	switch tLen {
	case 0:
		return notEmpty
	case 1:
		return JSONWriteStringValue(b, string(t[0]))
	default:
		JSONWrite(b, '[')
		for i, ty := range t {
			if !JSONWriteStringValue(b, string(ty)) {
				return false
			}
			if i < tLen-1 {
				JSONWriteComma(b)
			}
			notEmpty = true
		}
		JSONWrite(b, ']')
		return notEmpty
	}
}

// MarshalJSON represents just a wrapper for the jsonld.Marshal function
func MarshalJSON(it LinkOrIRI) ([]byte, error) {
	return jsonld.Marshal(it)
}
