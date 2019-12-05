package activitypub

import (
	"reflect"
	"testing"
)

type testPairs map[ActivityVocabularyType]reflect.Type

var objectPtrType = reflect.TypeOf(new(*Object)).Elem()
var tombstoneType = reflect.TypeOf(new(*Tombstone)).Elem()
var profileType = reflect.TypeOf(new(*Profile)).Elem()
var placeType = reflect.TypeOf(new(*Place)).Elem()
var relationshipType = reflect.TypeOf(new(*Relationship)).Elem()
var linkPtrType = reflect.TypeOf(new(*Link)).Elem()
var mentionPtrType = reflect.TypeOf(new(*Mention)).Elem()
var activityPtrType = reflect.TypeOf(new(*Activity)).Elem()
var intransitiveActivityPtrType = reflect.TypeOf(new(*IntransitiveActivity)).Elem()
var collectionPtrType = reflect.TypeOf(new(*Collection)).Elem()
var collectionPagePtrType = reflect.TypeOf(new(*CollectionPage)).Elem()
var orderedCollectionPtrType = reflect.TypeOf(new(*OrderedCollection)).Elem()
var orderedCollectionPagePtrType = reflect.TypeOf(new(*OrderedCollectionPage)).Elem()
var actorPtrType = reflect.TypeOf(new(*Actor)).Elem()
var applicationPtrType = reflect.TypeOf(new(*Application)).Elem()
var servicePtrType = reflect.TypeOf(new(*Service)).Elem()
var personPtrType = reflect.TypeOf(new(*Person)).Elem()
var groupPtrType = reflect.TypeOf(new(*Group)).Elem()
var organizationPtrType = reflect.TypeOf(new(*Organization)).Elem()
var acceptPtrType = reflect.TypeOf(new(*Accept)).Elem()
var addPtrType = reflect.TypeOf(new(*Add)).Elem()
var announcePtrType = reflect.TypeOf(new(*Announce)).Elem()
var arrivePtrType = reflect.TypeOf(new(*Arrive)).Elem()
var blockPtrType = reflect.TypeOf(new(*Block)).Elem()
var createPtrType = reflect.TypeOf(new(*Create)).Elem()
var deletePtrType = reflect.TypeOf(new(*Delete)).Elem()
var dislikePtrType = reflect.TypeOf(new(*Dislike)).Elem()
var flagPtrType = reflect.TypeOf(new(*Flag)).Elem()
var followPtrType = reflect.TypeOf(new(*Follow)).Elem()
var ignorePtrType = reflect.TypeOf(new(*Ignore)).Elem()
var invitePtrType = reflect.TypeOf(new(*Invite)).Elem()
var joinPtrType = reflect.TypeOf(new(*Join)).Elem()
var leavePtrType = reflect.TypeOf(new(*Leave)).Elem()
var likePtrType = reflect.TypeOf(new(*Like)).Elem()
var listenPtrType = reflect.TypeOf(new(*Listen)).Elem()
var movePtrType = reflect.TypeOf(new(*Move)).Elem()
var offerPtrType = reflect.TypeOf(new(*Offer)).Elem()
var questionPtrType = reflect.TypeOf(new(*Question)).Elem()
var rejectPtrType = reflect.TypeOf(new(*Reject)).Elem()
var readPtrType = reflect.TypeOf(new(*Read)).Elem()
var removePtrType = reflect.TypeOf(new(*Remove)).Elem()
var tentativeRejectPtrType = reflect.TypeOf(new(*TentativeReject)).Elem()
var tentativeAcceptPtrType = reflect.TypeOf(new(*TentativeAccept)).Elem()
var travelPtrType = reflect.TypeOf(new(*Travel)).Elem()
var undoPtrType = reflect.TypeOf(new(*Undo)).Elem()
var updatePtrType = reflect.TypeOf(new(*Update)).Elem()
var viewPtrType = reflect.TypeOf(new(*View)).Elem()

var tests = testPairs{
	ObjectType:                objectPtrType,
	ArticleType:               objectPtrType,
	AudioType:                 objectPtrType,
	DocumentType:              objectPtrType,
	ImageType:                 objectPtrType,
	NoteType:                  objectPtrType,
	PageType:                  objectPtrType,
	PlaceType:                 placeType,
	ProfileType:               profileType,
	RelationshipType:          relationshipType,
	TombstoneType:             tombstoneType,
	VideoType:                 objectPtrType,
	LinkType:                  linkPtrType,
	MentionType:               mentionPtrType,
	CollectionType:            collectionPtrType,
	CollectionPageType:        collectionPagePtrType,
	OrderedCollectionType:     orderedCollectionPtrType,
	OrderedCollectionPageType: orderedCollectionPagePtrType,
	ActorType:                 actorPtrType,
	ApplicationType:           applicationPtrType,
	ServiceType:               servicePtrType,
	PersonType:                personPtrType,
	GroupType:                 groupPtrType,
	OrganizationType:          organizationPtrType,
	ActivityType:              activityPtrType,
	IntransitiveActivityType:  intransitiveActivityPtrType,
	AcceptType:                acceptPtrType,
	AddType:                   addPtrType,
	AnnounceType:              announcePtrType,
	ArriveType:                arrivePtrType,
	BlockType:                 blockPtrType,
	CreateType:                createPtrType,
	DeleteType:                deletePtrType,
	DislikeType:               dislikePtrType,
	FlagType:                  flagPtrType,
	FollowType:                followPtrType,
	IgnoreType:                ignorePtrType,
	InviteType:                invitePtrType,
	JoinType:                  joinPtrType,
	LeaveType:                 leavePtrType,
	LikeType:                  likePtrType,
	ListenType:                listenPtrType,
	MoveType:                  movePtrType,
	OfferType:                 offerPtrType,
	QuestionType:              questionPtrType,
	RejectType:                rejectPtrType,
	ReadType:                  readPtrType,
	RemoveType:                removePtrType,
	TentativeRejectType:       tentativeRejectPtrType,
	TentativeAcceptType:       tentativeAcceptPtrType,
	TravelType:                travelPtrType,
	UndoType:                  undoPtrType,
	UpdateType:                updatePtrType,
	ViewType:                  viewPtrType,
}

func TestJSONGetItemByType(t *testing.T) {
	for typ, test := range tests {
		t.Run(string(typ), func(t *testing.T) {
			v, err := JSONGetItemByType(typ)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(v) != test {
				t.Errorf("Invalid type returned %T, expected %s", v, test.String())
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	dataEmpty := []byte("{}")
	i, err := UnmarshalJSON(dataEmpty)
	if err != nil {
		t.Errorf("invalid unmarshaling %s", err)
	}
	if i != nil {
		t.Errorf("invalid unmarshaling, expected nil")
	}
}

func TestJSONGetDuration(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetInt(t *testing.T) {

}

func TestJSONGetIRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetItem(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetItems(t *testing.T) {

}

func TestJSONGetLangRefField(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetMimeType(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetNaturalLanguageField(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetString(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetTime(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetURIItem(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONUnmarshalToItem(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetActorEndpoints(t *testing.T) {
	t.Skipf("TODO")
}
