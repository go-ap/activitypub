package activitypub

var validLinkTypes = [...]ActivityVocabularyType{
	MentionType,
}

// A GetLink is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988].
// Many of the properties defined by the Activity Vocabulary allow values that are either instances of APObject or GetLink.
// When a GetLink is used, it establishes a qualified relation connecting the subject
//  (the containing object) to the resource identified by the href.
// Properties of the GetLink are properties of the reference as opposed to properties of the resource.
type Link struct {
	// Provides the globally unique identifier for an APObject or GetLink.
	Id ObjectID `jsonld:"id,omitempty"`
	//  Identifies the APObject or GetLink type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
	// A link relation associated with a GetLink. The value must conform to both the [HTML5] and
	//  [RFC5988](https://tools.ietf.org/html/rfc5988) "link relation" definitions.
	// In the [HTML5], any string not containing the "space" U+0020, "tab" (U+0009), "LF" (U+000A),
	//  "FF" (U+000C), "CR" (U+000D) or "," (U+002C) characters can be used as a valid link relation.
	Rel *Link `jsonld:"rel"`
	// When used on a GetLink, identifies the MIME media type of the referenced resource.
	// When used on an APObject, identifies the MIME media type of the value of the content property.
	// If not specified, the content property is assumed to contain text/html content.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// On a GetLink, specifies a hint as to the rendering height in device-independent pixels of the linked resource.
	Height uint `jsonld:"height,omitempty"`
	// On a GetLink, specifies a hint as to the rendering width in device-independent pixels of the linked resource.
	Width uint `jsonld:"width,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview ObjectOrLink `jsonld:"preview,omitempty"`
	// The target resource pointed to by a GetLink.
	Href URI `jsonld:"href,omitempty"`
	// Hints as to the language used by the target resource.
	// Value must be a [BCP47](https://tools.ietf.org/html/bcp47) Language-Tag.
	HrefLang LangRef `jsonld:"hrefLang,omitempty"`
}

// Mention is a specialized GetLink that represents an @mention.
type Mention Link

// ValidLinkType validates a type against the valid link types
func ValidLinkType(_type ActivityVocabularyType) bool {
	for _, v := range validLinkTypes {
		if v == _type {
			return true
		}
	}
	return false
}

// LinkNew initializes a new GetLink
func LinkNew(id ObjectID, _type ActivityVocabularyType) *Link {
	if !ValidLinkType(_type) {
		_type = LinkType
	}
	return &Link{Id: id, Type: _type}
}

// MentionNew initializes a new Mention
func MentionNew(id ObjectID) *Mention {
	return &Mention{Id: id, Type: MentionType}
}

// IsLink validates if current GetLink is a GetLink
func (l Link) IsLink() bool {
	return l.Type == LinkType || ValidLinkType(l.Type)
}

// IsObject validates if current GetLink is an GetObject
func (l Link) IsObject() bool {
	return l.Type == ObjectType || ValidObjectType(l.Type)
}

// IsLink validates if current Mention is a GetLink
func (m Mention) IsLink() bool {
	return m.Type == MentionType || ValidLinkType(m.Type)
}

// IsObject validates if current Mention is an GetObject
func (m Mention) IsObject() bool {
	return m.Type == ObjectType || ValidObjectType(m.Type)
}

// GetObject returns the GetObject corresponding to the Mention object
func (m Mention) GetObject() Object {
	return Object{}
}

// GetLink returns the GetLink corresponding to the Mention object
func (m Mention) GetLink() Link {
	return Link(m)
}
