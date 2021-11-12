package activitypub

// ActivityVocabularyTypes is a type alias for a slice of ActivityVocabularyType elements
type ActivityVocabularyTypes []ActivityVocabularyType

// Types contains all valid types in the ActivityPub vocabulary
var Types = ActivityVocabularyTypes{
	LinkType,
	MentionType,

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

	QuestionType,

	CollectionType,
	OrderedCollectionType,
	CollectionPageType,
	OrderedCollectionPageType,

	ApplicationType,
	GroupType,
	OrganizationType,
	PersonType,
	ServiceType,

	AcceptType,
	AddType,
	AnnounceType,
	BlockType,
	CreateType,
	DeleteType,
	DislikeType,
	FlagType,
	FollowType,
	IgnoreType,
	InviteType,
	JoinType,
	LeaveType,
	LikeType,
	ListenType,
	MoveType,
	OfferType,
	RejectType,
	ReadType,
	RemoveType,
	TentativeRejectType,
	TentativeAcceptType,
	UndoType,
	UpdateType,
	ViewType,

	ArriveType,
	TravelType,
	QuestionType,
}
