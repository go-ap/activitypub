package activitypub

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"reflect"
	"time"
)

type typeId int32

type gobEncoder struct {
	sent map[reflect.Type]typeId // which types we've already sent
	w    *bytes.Buffer
	enc  *gob.Encoder
}

func (e *gobEncoder) encode(it Item) ([]byte, error) {
	err := e.enc.Encode(it)
	if err != nil {
		return nil, err
	}
	return e.w.Bytes(), nil
}

//// GobEncode
//func GobEncode(it Item) ([]byte, error) {
//	w := bytes.NewBuffer(make([]byte, 0))
//	enc := gobEncoder{
//		w:   w,
//		enc: gob.NewEncoder(w),
//	}
//	return enc.encode(it)
//}

func (e *gobEncoder) writeS(s string) error {
	return e.enc.Encode(s)
}

func (e *gobEncoder) writeIRIProp(i IRI) error {
	return e.enc.Encode(i.String())
}

func (e *gobEncoder) writeGobProp(p string, b []byte) bool {
	c, _ := e.w.Write([]byte(p))
	d, _ := e.w.Write(b)
	return c+d > 0
}
func (e *gobEncoder) writeItemGobProp(p string, it Item) bool {
	return true
}
func (e *gobEncoder) writeNaturalLanguageGobProp(p string, v NaturalLanguageValues) bool {
	return true
}
func (e *gobEncoder) writeIRIGobProp(p string, i LinkOrIRI) bool {
	return true
}
func (e *gobEncoder) writeTimeGobProp(p string, t time.Time) bool {
	return true
}
func (e *gobEncoder) writeDurationGobProp(p string, d time.Duration) bool {
	return true
}

func writeObjectGobValue(buf io.Writer, o *Object) (int, error) {
	return 0, errors.New(fmt.Sprintf("writeObjectGobValue is not implemented for %T", *o))
}

/*
func (e *gobEncoder) writeObjectGobValue(o Object) (bool, error) {
	notEmpty := true
	if v, err := o.ID.GobEncode(); err == nil && len(v) > 0 {
		notEmpty = e.writeGobProp("id", v) || notEmpty
	}
	if v, err := o.Type.GobEncode(); err == nil && len(v) > 0 {
		notEmpty = e.writeGobProp("type", v) || notEmpty
	}
	if v, err := o.MediaType.GobEncode(); err == nil && len(v) > 0 {
		notEmpty = e.writeGobProp("mediaType", v) || notEmpty
	}
	if len(o.Name) > 0 {
		notEmpty = e.writeNaturalLanguageGobProp("name", o.Name) || notEmpty
	}
	if len(o.Summary) > 0 {
		notEmpty = e.writeNaturalLanguageGobProp("summary", o.Summary) || notEmpty
	}
	if len(o.Content) > 0 {
		notEmpty = e.writeNaturalLanguageGobProp("content", o.Content) || notEmpty
	}
	if o.Attachment != nil {
		notEmpty = e.writeItemGobProp("attachment", o.Attachment) || notEmpty
	}
	if o.AttributedTo != nil {
		notEmpty = e.writeItemGobProp("attributedTo", o.AttributedTo) || notEmpty
	}
	if o.Audience != nil {
		notEmpty = e.writeItemGobProp("audience", o.Audience) || notEmpty
	}
	if o.Context != nil {
		notEmpty = e.writeItemGobProp("context", o.Context) || notEmpty
	}
	if o.Generator != nil {
		notEmpty = e.writeItemGobProp("generator", o.Generator) || notEmpty
	}
	if o.Icon != nil {
		notEmpty = e.writeItemGobProp("icon", o.Icon) || notEmpty
	}
	if o.Image != nil {
		notEmpty = e.writeItemGobProp("image", o.Image) || notEmpty
	}
	if o.InReplyTo != nil {
		notEmpty = e.writeItemGobProp("inReplyTo", o.InReplyTo) || notEmpty
	}
	if o.Location != nil {
		notEmpty = e.writeItemGobProp("location", o.Location) || notEmpty
	}
	if o.Preview != nil {
		notEmpty = e.writeItemGobProp("preview", o.Preview) || notEmpty
	}
	if o.Replies != nil {
		notEmpty = e.writeItemGobProp("replies", o.Replies) || notEmpty
	}
	if o.Tag != nil {
		notEmpty = e.writeItemGobProp("tag", o.Tag) || notEmpty
	}
	if o.URL != nil {
		notEmpty = e.writeIRIGobProp("url", o.URL) || notEmpty
	}
	if o.To != nil {
		notEmpty = e.writeItemGobProp("to", o.To) || notEmpty
	}
	if o.Bto != nil {
		notEmpty = e.writeItemGobProp("bto", o.Bto) || notEmpty
	}
	if o.CC != nil {
		notEmpty = e.writeItemGobProp("cc", o.CC) || notEmpty
	}
	if o.BCC != nil {
		notEmpty = e.writeItemGobProp("bcc", o.BCC) || notEmpty
	}
	if !o.Published.IsZero() {
		notEmpty = e.writeTimeGobProp("published", o.Published) || notEmpty
	}
	if !o.Updated.IsZero() {
		notEmpty = e.writeTimeGobProp("updated", o.Updated) || notEmpty
	}
	if !o.StartTime.IsZero() {
		notEmpty = e.writeTimeGobProp("startTime", o.StartTime) || notEmpty
	}
	if !o.EndTime.IsZero() {
		notEmpty = e.writeTimeGobProp("endTime", o.EndTime) || notEmpty
	}
	if o.Duration != 0 {
		// TODO(marius): maybe don't use 0 as a nil value for Object types
		//  which can have a valid duration of 0 - (Video, Audio, etc)
		notEmpty = e.writeDurationGobProp("duration", o.Duration) || notEmpty
	}
	if o.Likes != nil {
		notEmpty = e.writeItemGobProp("likes", o.Likes) || notEmpty
	}
	if o.Shares != nil {
		notEmpty = e.writeItemGobProp("shares", o.Shares) || notEmpty
	}
	if v, err := o.Source.GobEncode(); err == nil && len(v) > 0 {
		notEmpty = e.writeGobProp("source", v) || notEmpty
	}
	return notEmpty, nil
}
*/
