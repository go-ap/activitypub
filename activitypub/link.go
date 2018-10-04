package activitypub

var validLinkTypes = [...]ActivityVocabularyType{
	MentionType,
}

// A Link is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988].
// Many of the properties defined by the Activity Vocabulary allow values that are either instances of APObject or Link.
// When a Link is used, it establishes a qualified relation connecting the subject
//  (the containing object) to the resource identified by the href.
// Properties of the Link are properties of the reference as opposed to properties of the resource.
type Link struct {
	// Provides the globally unique identifier for an APObject or Link.
	ID ObjectID `jsonld:"id,omitempty"`
	//  Identifies the APObject or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
	// A link relation associated with a Link. The value must conform to both the [HTML5] and
	//  [RFC5988](https://tools.ietf.org/html/rfc5988) "link relation" definitions.
	// In the [HTML5], any string not containing the "space" U+0020, "tab" (U+0009), "LF" (U+000A),
	//  "FF" (U+000C), "CR" (U+000D) or "," (U+002C) characters can be used as a valid link relation.
	Rel *Link `jsonld:"rel,omitempty"`
	// When used on a Link, identifies the MIME media type of the referenced resource.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// On a Link, specifies a hint as to the rendering height in device-independent pixels of the linked resource.
	Height uint `jsonld:"height,omitempty"`
	// On a Link, specifies a hint as to the rendering width in device-independent pixels of the linked resource.
	Width uint `jsonld:"width,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// The target resource pointed to by a Link.
	Href URI `jsonld:"href,omitempty"`
	// Hints as to the language used by the target resource.
	// Value must be a [BCP47](https://tools.ietf.org/html/bcp47) Language-Tag.
	HrefLang LangRef `jsonld:"hrefLang,omitempty"`
}

// Mention is a specialized Link that represents an @mention.
type Mention Link

// ValidLinkType validates a type against the valid link types
func ValidLinkType(typ ActivityVocabularyType) bool {
	for _, v := range validLinkTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// LinkNew initializes a new Link
func LinkNew(id ObjectID, typ ActivityVocabularyType) *Link {
	if !ValidLinkType(typ) {
		typ = LinkType
	}
	return &Link{ID: id, Type: typ}
}

// MentionNew initializes a new Mention
func MentionNew(id ObjectID) *Mention {
	return &Mention{ID: id, Type: MentionType}
}

// IsLink validates if current Link is a Link
func (l Link) IsLink() bool {
	return l.Type == LinkType || ValidLinkType(l.Type)
}

// IsObject validates if current Link is an GetID
func (l Link) IsObject() bool {
	return l.Type == ObjectType || ValidObjectType(l.Type)
}

// GetID returns the ObjectID corresponding to the Mention object
func (l Link) GetID() *ObjectID {
	return &l.ID
}

// GetType returns the Type corresponding to the Mention object
func (l Link) GetType() ActivityVocabularyType {
	return l.Type
}

// IsLink validates if current Mention is a Link
func (m Mention) IsLink() bool {
	return m.Type == MentionType || ValidLinkType(m.Type)
}

// IsObject validates if current Mention is an GetID
func (m Mention) IsObject() bool {
	return m.Type == ObjectType || ValidObjectType(m.Type)
}

// GetID returns the ObjectID corresponding to the Mention object
func (m Mention) GetID() *ObjectID {
	return Link(m).GetID()
}

// GetType returns the Type corresponding to the Mention object
func (m Mention) GetType() ActivityVocabularyType {
	return m.Type
}

// UnmarshalJSON
func (l *Link) UnmarshalJSON(data []byte) error {
	l.ID = getAPObjectID(data)
	l.Type = getAPType(data)
	l.MediaType = getAPMimeType(data)
	l.Name = getAPNaturalLanguageField(data, "name")
	l.HrefLang = getAPLangRefField(data, "hrefLang")
	u := getURIField(data, "href")
	if len(u) > 0 {
		l.Href = u
	}

	//fmt.Printf("%s\n %#v", data, l)

	return nil
}
