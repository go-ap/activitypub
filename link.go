package activitypub

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/valyala/fastjson"
)

// LinkTypes represent the valid values for a Link object
var LinkTypes = ActivityVocabularyTypes{
	LinkType,
	MentionType,
}

// A Link is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988].
// Many of the properties defined by the Activity Vocabulary allow values that are either instances of APObject or Link.
// When a Link is used, it establishes a qualified relation connecting the subject
// (the containing object) to the resource identified by the href.
// Properties of the Link are properties of the reference as opposed to properties of the resource.
type Link struct {
	// Provides the globally unique identifier for an APObject or Link.
	ID ID `jsonld:"id,omitempty"`
	// Identifies the APObject or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValues `jsonld:"name,omitempty,collapsible"`
	// A link relation associated with a Link. The value must conform to both the [HTML5] and
	// [RFC5988](https://tools.ietf.org/html/rfc5988) "link relation" definitions.
	// In the [HTML5], any string not containing the "space" U+0020, "tab" (U+0009), "LF" (U+000A),
	// "FF" (U+000C), "CR" (U+000D) or "," (U+002C) characters can be used as a valid link relation.
	Rel IRI `jsonld:"rel,omitempty"`
	// When used on a Link, identifies the MIME media type of the referenced resource.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// On a Link, specifies a hint as to the rendering height in device-independent pixels of the linked resource.
	Height uint `jsonld:"height,omitempty"`
	// On a Link, specifies a hint as to the rendering width in device-independent pixels of the linked resource.
	Width uint `jsonld:"width,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// The target resource pointed to by a Link.
	Href IRI `jsonld:"href,omitempty"`
	// Hints as to the language used by the target resource.
	// Value must be a [BCP47](https://tools.ietf.org/html/bcp47) Language-Tag.
	HrefLang LangRef `jsonld:"hrefLang,omitempty"`
}

// Mention is a specialized Link that represents an @mention.
type Mention = Link

// LinkNew initializes a new Link
func LinkNew(id ID, typ ActivityVocabularyType) *Link {
	if !LinkTypes.Contains(typ) {
		typ = LinkType
	}
	return &Link{ID: id, Type: typ}
}

// MentionNew initializes a new Mention
func MentionNew(id ID) *Mention {
	return &Mention{ID: id, Type: MentionType}
}

// IsLink validates if current Link is a Link
func (l Link) IsLink() bool {
	return l.Type == LinkType || LinkTypes.Contains(l.Type)
}

// IsObject validates if current Link is an GetID
func (l Link) IsObject() bool {
	return l.Type == ObjectType || ObjectTypes.Contains(l.Type)
}

// IsCollection returns false for Link objects
func (l Link) IsCollection() bool {
	return false
}

// GetID returns the ID corresponding to the Link object
func (l Link) GetID() ID {
	return l.ID
}

// GetLink returns the IRI corresponding to the current Link
func (l Link) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the Type corresponding to the Mention object
func (l Link) GetType() ActivityVocabularyType {
	return l.Type
}

// MarshalJSON encodes the receiver object to a JSON document.
func (l Link) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	write(&b, '{')

	if writeLinkJSONValue(&b, l) {
		write(&b, '}')
		return b, nil
	}
	return nil, nil
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (l *Link) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return loadLink(val, l)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (l *Link) UnmarshalBinary(data []byte) error {
	return l.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (l Link) MarshalBinary() ([]byte, error) {
	return l.GobEncode()
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

func (l Link) GobEncode() ([]byte, error) {
	var (
		mm      = make(map[string][]byte)
		err     error
		hasData bool
	)
	if len(l.ID) > 0 {
		if mm["id"], err = l.ID.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(l.Type) > 0 {
		if mm["type"], err = l.Type.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(l.MediaType) > 0 {
		if mm["mediaType"], err = l.MediaType.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(l.Href) > 0 {
		if mm["href"], err = l.Href.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(l.HrefLang) > 0 {
		if mm["hrefLang"], err = l.HrefLang.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(l.Name) > 0 {
		if mm["name"], err = l.Name.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(l.Rel) > 0 {
		if mm["rel"], err = l.Rel.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if l.Width > 0 {
		if mm["width"], err = gobEncodeUint(l.Width); err != nil {
			return nil, err
		}
		hasData = true
	}
	if l.Height > 0 {
		if mm["height"], err = gobEncodeUint(l.Height); err != nil {
			return nil, err
		}
		hasData = true
	}
	if !hasData {
		return []byte{}, nil
	}
	bb := bytes.Buffer{}
	g := gob.NewEncoder(&bb)
	if err := g.Encode(mm); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func gobDecodeUint(i *uint, data []byte) error {
	g := gob.NewDecoder(bytes.NewReader(data))
	return g.Decode(i)
}

func (l *Link) GobDecode(data []byte) error {
	mm := make(map[string][]byte)
	g := gob.NewDecoder(bytes.NewReader(data))
	if err := g.Decode(&mm); err != nil {
		return err
	}
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
	return errors.New(fmt.Sprintf("GobDecode is not implemented for %T", *l))
}
