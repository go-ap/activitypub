package activitypub

import (
	"bytes"
	"fmt"
	"time"
)

func writeProp(b *bytes.Buffer, name string, val []byte) (notEmpty bool) {
	if len(val) == 0 {
		return false
	}
	writePropName(b, name)
	return writeValue(b, val)
}

func writePropName(b *bytes.Buffer, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	b.Write([]byte{'"'})
	b.WriteString(s)
	b.Write([]byte{'"', ':'})
	return true
}

func writeValue(b *bytes.Buffer, s []byte) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	b.Write(s)
	return true
}

func writeNaturalLanguageProp(b *bytes.Buffer, n string, nl NaturalLanguageValues) (notEmpty bool) {
	l := nl.Count()
	if l > 1 {
		n += "Map"
	}
	if v, err := nl.MarshalJSON(); err == nil && len(v) > 0 {
		return writeProp(b, n, v)
	}
	return false
}

func writeBoolProp(b *bytes.Buffer, n string, t bool) (notEmpty bool) {
	return writeProp(b, n, []byte(fmt.Sprintf("%t", t)))
}
func writeIntProp(b *bytes.Buffer, n string, d int) (notEmpty bool) {
	return writeProp(b, n, []byte(fmt.Sprintf("%d", d)))
}
func writeTimeProp(b *bytes.Buffer, n string, t time.Time) (notEmpty bool) {
	if v, err := t.MarshalJSON(); err == nil {
		return writeProp(b, n, v)
	}
	return false
}

func writeDurationProp(b *bytes.Buffer, n string, d time.Duration) (notEmpty bool) {
	if v, err := marshalXSD(d); err == nil {
		return writeProp(b, n, v)
	}
	return false
}

func writeIRIProp(b *bytes.Buffer, n string, i LinkOrIRI) (notEmpty bool) {
	url := i.GetLink()
	if len(url) == 0 {
		return false
	}
	writePropName(b, n)
	b.Write([]byte{'"'})
	b.Write([]byte(url))
	b.Write([]byte{'"'})
	return true
}

func writeItemProp(b *bytes.Buffer, n string, i Item) (notEmpty bool) {
	if i == nil {
		return notEmpty
	}
	if i.IsObject() {
		OnObject(i, func(o *Object) error {
			v, err := o.MarshalJSON()
			if err != nil {
				return nil
			}
			notEmpty = writeProp(b, n, v)
			return nil
		})
	} else if i.IsCollection() {
		OnCollection(i, func(c CollectionInterface) error {
			notEmpty = writeItemCollectionProp(b, n, c.Collection()) || notEmpty
			return nil
		})
	}
	return notEmpty
}

func writeString(b *bytes.Buffer, s string) (notEmpty bool) {
	if len(s) == 0 {
		return false
	}
	b.Write([]byte{'"'})
	b.WriteString(s)
	b.Write([]byte{'"'})
	return true
}

func writeItemCollection(b *bytes.Buffer, col ItemCollection) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	writeComma := func() { b.WriteString(",") }
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeComma()
		}
	}
	b.Write([]byte{'['})
	for _, i := range col {
		if i.IsObject() {
			OnObject(i, func(o *Object) error {
				v, err := o.MarshalJSON()
				if err != nil {
					return nil
				}
				writeCommaIfNotEmpty(notEmpty)
				notEmpty = writeValue(b, v) || notEmpty
				return nil
			})
		} else if i.IsLink() {
			writeCommaIfNotEmpty(notEmpty)
			notEmpty = writeValue(b, []byte(i.GetLink())) || notEmpty
		}
	}
	b.Write([]byte{']'})
	return notEmpty
}
func writeItemCollectionProp(b *bytes.Buffer, n string, col ItemCollection) (notEmpty bool) {
	if len(col) == 0 {
		return notEmpty
	}
	writePropName(b, n)
	return writeItemCollection(b, col)
}

func writeObject(b *bytes.Buffer, o Object) (notEmpty bool) {
	writeComma := func() { b.WriteString(",") }
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeComma()
		}
	}
	if v, err := o.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = writeProp(b, "id", v) || notEmpty
	}
	if v, err := o.Type.MarshalJSON(); err == nil && len(v) > 0 {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeProp(b, "type", v) || notEmpty
	}
	if v, err := o.MediaType.MarshalJSON(); err == nil && len(v) > 0 {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeProp(b, "mediaType", v) || notEmpty
	}
	if len(o.Name) > 0 {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeNaturalLanguageProp(b, "name", o.Name) || notEmpty
	}
	if len(o.Summary) > 0 {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeNaturalLanguageProp(b, "summary", o.Summary) || notEmpty
	}
	if len(o.Content) > 0 {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeNaturalLanguageProp(b, "content", o.Content) || notEmpty
	}
	if o.Attachment != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "attachment", o.Attachment) || notEmpty
	}
	if o.AttributedTo != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "attributedTo", o.AttributedTo) || notEmpty
	}
	if o.Audience != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "audience", o.Audience) || notEmpty
	}
	if o.Context != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "context", o.Context) || notEmpty
	}
	if o.Generator != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "generator", o.Generator) || notEmpty
	}
	if o.Icon != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "icon", o.Icon) || notEmpty
	}
	if o.Image != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "image", o.Image) || notEmpty
	}
	if o.InReplyTo != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "inReplyTo", o.InReplyTo) || notEmpty
	}
	if o.Location != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "location", o.Location) || notEmpty
	}
	if o.Preview != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "preview", o.Preview) || notEmpty
	}
	if o.Replies != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "replies", o.Replies) || notEmpty
	}
	if o.Tag != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "tag", o.Tag) || notEmpty
	}
	if o.URL != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeIRIProp(b, "url", o.URL) || notEmpty
	}
	if o.To != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "to", o.To) || notEmpty
	}
	if o.Bto != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "bto", o.Bto) || notEmpty
	}
	if o.CC != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "cc", o.CC) || notEmpty
	}
	if o.BCC != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "bcc", o.BCC) || notEmpty
	}
	if !o.Published.IsZero() {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeTimeProp(b, "published", o.Published) || notEmpty
	}
	if !o.Updated.IsZero() {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeTimeProp(b, "updated", o.Updated) || notEmpty
	}
	if !o.StartTime.IsZero() {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeTimeProp(b, "startTime", o.StartTime) || notEmpty
	}
	if !o.EndTime.IsZero() {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeTimeProp(b, "endTime", o.EndTime) || notEmpty
	}
	if o.Duration != 0 {
		// TODO(marius): maybe don't use 0 as a nil value for Object types
		//  which can have a valid duration of 0 - (Video, Audio, etc)
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeDurationProp(b, "duration", o.Duration) || notEmpty
	}
	if o.Likes != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "likes", o.Likes) || notEmpty
	}
	if o.Shares != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "shares", o.Shares) || notEmpty
	}
	if v, err := o.Source.MarshalJSON(); err == nil && len(v) > 0 {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeProp(b, "source", v) || notEmpty
	}
	return notEmpty
}

func writeActivity(b *bytes.Buffer, a Activity) (notEmpty bool) {
	writeComma := func() { b.WriteString(",") }
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeComma()
		}
	}

	OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = writeIntransitiveActivity(b, *i) || notEmpty
		return nil
	})
	if a.Object != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "object", a.Object) || notEmpty
	}
	return notEmpty
}

func writeIntransitiveActivity(b *bytes.Buffer, i IntransitiveActivity) (notEmpty bool) {
	writeComma := func() { b.WriteString(",") }
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeComma()
		}
	}
	OnObject(i, func(o *Object) error {
		if o == nil {
			return nil
		}
		notEmpty = writeObject(b, *o) || notEmpty
		return nil
	})
	if i.Actor != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "actor", i.Actor) || notEmpty
	}
	if i.Target != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "target", i.Target) || notEmpty
	}
	if i.Result != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "result", i.Result) || notEmpty
	}
	if i.Origin != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "origin", i.Origin) || notEmpty
	}
	if i.Instrument != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "instrument", i.Instrument) || notEmpty
	}
	return notEmpty
}

func writeQuestion(b *bytes.Buffer, q Question) (notEmpty bool) {
	writeComma := func() { b.WriteString(",") }
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeComma()
		}
	}

	OnIntransitiveActivity(q, func(i *IntransitiveActivity) error {
		if i == nil {
			return nil
		}
		notEmpty = writeIntransitiveActivity(b, *i) || notEmpty
		return nil
	})
	if q.OneOf != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "oneOf", q.OneOf) || notEmpty
	} else if q.AnyOf != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(b, "oneOf", q.OneOf) || notEmpty
	}
	writeCommaIfNotEmpty(notEmpty)
	notEmpty = writeBoolProp(b, "closed", q.Closed) || notEmpty
	return notEmpty
}
