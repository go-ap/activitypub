package activitypub

type ObjectId string

const (
	ObjectType   string = "Object"
	LinkType     string = "Link"
	ActivityType string = "Activity"
	ActorType    string = "Actor"
	CollectionType string = "Collection"
	OrderedCollectionType string = "OrderedCollection"

	// Object Types
	ArticleType string = "Article"
	AudioType string = "Audio"
	DocumentType string = "Document"
	EventType string = "Event"
	ImageType string = "Image"
	NoteType string = "Note"
	PageType string = "Page"
	PlaceType string = "Place"
	ProfileType string = "Profile"
	RelationshipType string = "Relationship"
	TombstoneType string = "Tombstone"
	VideoType string = "Video"

	// Link Types
	MentionType string = "Mention"
)

var validObjectTypes = [...]string{
	ArticleType,
	AudioType,
	DocumentType,
	EventType,
	ImageType,
	NoteType,
	PageType,
	PlaceType,
	ProfileType,
	RelationshipType,
	TombstoneType,
	VideoType,
}

var validLinkTypes = [...]string{
	MentionType,
}

type NaturalLanguageValue map[string]string

type BaseObject struct {
	Id   ObjectId
	Type string
	Name NaturalLanguageValue

	Href      string
	HrefLang  string
	MediaType string
}

type ContentType string

type Source struct {
	Content   ContentType
	MediaType string
}

func ValidObjectType(_type string) bool {
	for _, v := range validObjectTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ValidLinkType(_type string) bool {
	for _, v := range validLinkTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ObjectNew(id ObjectId, _type string) BaseObject {
	if !ValidObjectType(_type) {
		_type = ObjectType
	}
	return BaseObject{Id: id, Type: _type}
}

func LinkNew(id ObjectId, _type string) BaseObject {
	if !ValidLinkType(_type) {
		_type = LinkType
	}
	return BaseObject{Id: id, Type:_type}
}
