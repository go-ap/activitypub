package activitypub

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

type canErrorFunc func(format string, args ...interface{})

// See reflect.DeepEqual
func assertDeepEquals(t canErrorFunc, x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		t("%T != %T", x, y)
		return false
	}
	return deepValueEqual(t, v1, v2, make(map[visit]bool), 0)
}

// See reflect.deepValueEqual
func deepValueEqual(t canErrorFunc, v1, v2 reflect.Value, visited map[visit]bool, depth int) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		t("types differ %s != %s", v1.Type().Name(), v2.Type().Name())
		return false
	}

	hard := func(v1, v2 reflect.Value) bool {
		switch v1.Kind() {
		case reflect.Ptr:
			return false
		case reflect.Map, reflect.Slice, reflect.Interface:
			// Nil pointers cannot be cyclic. Avoid putting them in the visited map.
			return !v1.IsNil() && !v2.IsNil()
		}
		return false
	}

	if hard(v1, v2) {
		var addr1, addr2 unsafe.Pointer
		if v1.CanAddr() {
			addr1 = unsafe.Pointer(v1.UnsafeAddr())
		} else {
			addr1 = unsafe.Pointer(v1.Pointer())
		}
		if v2.CanAddr() {
			addr2 = unsafe.Pointer(v2.UnsafeAddr())
		} else {
			addr2 = unsafe.Pointer(v2.Pointer())
		}
		if uintptr(addr1) > uintptr(addr2) {
			// Canonicalize order to reduce number of entries in visited.
			// Assumes non-moving garbage collector.
			addr1, addr2 = addr2, addr1
		}
		// Short circuit if references are already seen.
		typ := v1.Type()
		v := visit{addr1, addr2, typ}
		if visited[v] {
			return true
		}

		// Remember for later.
		visited[v] = true
	}

	switch v1.Kind() {
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(t, v1.Index(i), v2.Index(i), visited, depth+1) {
				t("Arrays not equal at index %d %s %s", i, v1.Index(i), v2.Index(i))
				return false
			}
		}
		return true
	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			t("One of the slices is not nil %s[%d] vs %s[%d]", v1.Type().Name(), v1.Len(), v2.Type().Name(), v2.Len())
			return false
		}
		if v1.Len() != v2.Len() {
			t("Slices lengths are different %s[%d] vs %s[%d]", v1.Type().Name(), v1.Len(), v2.Type().Name(), v2.Len())
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(t, v1.Index(i), v2.Index(i), visited, depth+1) {
				t("Slices elements at pos %d are not equal %#v vs %#v", i, v1.Index(i), v2.Index(i))
				return false
			}
		}
		return true
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			if v1.IsNil() == v2.IsNil() {
				return true
			}
			var isNil1, isNil2 string
			if v1.IsNil() {
				isNil1 = "is"
			} else {
				isNil1 = "is not"
			}
			if v2.IsNil() {
				isNil2 = "is"
			} else {
				isNil2 = "is not"
			}
			t("Interface '%s' %s nil and '%s' %s nil", v1.Type().Name(), isNil1, v2.Type().Name(), isNil2)
			return false
		}
		return deepValueEqual(t, v1.Elem(), v2.Elem(), visited, depth+1)
	case reflect.Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return deepValueEqual(t, v1.Elem(), v2.Elem(), visited, depth+1)
	case reflect.Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			var (
				f1 = v1.Field(i); f2 = v2.Field(i)
				n1 = v1.Type().Field(i).Name; n2 = v2.Type().Field(i).Name
				t1 = f1.Type().Name(); t2 = f2.Type().Name()
			)
			if !deepValueEqual(t, v1.Field(i), v2.Field(i), visited, depth+1) {
				t("Struct fields at pos %d %s[%s] and %s[%s] are not deeply equal", i, n1, t1, n2, t2)
				if f1.CanInterface() && f2.CanInterface() {
					t("  Values: %#v - %#v", v1.Field(i).Interface(), v2.Field(i).Interface())
				}
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() != v2.IsNil() {
			t("Maps are not nil", v1.Type().Name(), v2.Type().Name())
			return false
		}
		if v1.Len() != v2.Len() {
			t("Maps don't have the same length %d vs %d", v1.Len(), v2.Len())
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(t, v1.MapIndex(k), v2.MapIndex(k), visited, depth+1) {
				t("Maps values at index %s are not equal", k.String())
				return false
			}
		}
		return true
	case reflect.Func:
		if v1.IsNil() && v2.IsNil() {
			return true
		}
		// Can't do better than this:
		return false
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	}
	return false
}

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
			v, err := GetItemByType(typ)
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
	tests := []struct {
		name string
		data []byte
		want Item
		err     error
	}{
		{
			name: "empty",
			data: []byte{'{', '}'},
			want: nil,
			err:  nil,
		},
		{
			name: "IRI",
			data: []byte(`"http://example.com"`),
			want: IRI("http://example.com"),
			err:  nil,
		},
		{
			name: "IRIs",
			data: []byte(fmt.Sprintf("[%q, %q]", "http://example.com", "http://example.net")),
			want:    ItemCollection{
				IRI("http://example.com"),
				IRI("http://example.net"),
			},
			err:     nil,
		},
		{
			name: "object",
			data: []byte(`{"type":"Note"}`),
			want: &Object{Type: NoteType},
			err:  nil,
		},
		{
			name: "activity",
			data: []byte(`{"type":"Like"}`),
			want: &Activity{Type: LikeType},
			err:  nil,
		},
		{
			name: "collection-2-items",
			data: []byte(`{ "@context": "https://www.w3.org/ns/activitystreams", "id": "https://federated.git/inbox", "type": "OrderedCollection", "updated": "2021-08-08T16:09:05Z", "first": "https://federated.git/inbox?maxItems=100", "totalItems": 2, "orderedItems": [ { "id": "https://federated.git/activities/07440c39-64b2-4492-89cf-f5c2872cf4ff", "type": "Create", "attributedTo": "https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91", "to": [ "https://www.w3.org/ns/activitystreams#Public" ], "cc": [ "https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91/followers" ], "published": "2021-08-08T16:09:05Z", "actor": "https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91", "object": "https://federated.git/objects/3eb69f77-3b08-4bf1-8760-c7333e2900c4" }, { "id": "https://federated.git/activities/ab9a5511-cdb5-4585-8a48-775d1bf20121", "type": "Like", "attributedTo": "https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91", "to": [ "https://www.w3.org/ns/activitystreams#Public", "https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91" ], "published": "2021-08-08T16:09:05Z", "actor": "https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91", "object": "https://federated.git/objects/3eb69f77-3b08-4bf1-8760-c7333e2900c4" }]}`),
			want: &OrderedCollection{
				ID: "https://federated.git/inbox",
				Type: OrderedCollectionType,
				Updated: time.Date(2021, 8, 8, 16, 9, 5, 0, time.UTC),
				First: IRI("https://federated.git/inbox?maxItems=100"),
				OrderedItems: ItemCollection{
					&Activity{
						ID:           "https://federated.git/activities/07440c39-64b2-4492-89cf-f5c2872cf4ff",
						Type:         CreateType,
						AttributedTo: IRI("https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91"),
						To:           ItemCollection{PublicNS},
						CC:           ItemCollection{IRI("https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91/followers")},
						Published:    time.Date(2021, 8, 8, 16, 9, 5, 0, time.UTC),
						Actor:        IRI("https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91"),
						Object:       IRI("https://federated.git/objects/3eb69f77-3b08-4bf1-8760-c7333e2900c4"),
					},
					&Activity{
						ID:           "https://federated.git/activities/ab9a5511-cdb5-4585-8a48-775d1bf20121",
						Type:         LikeType,
						AttributedTo: IRI("https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91"),
						To:           ItemCollection{PublicNS, IRI("https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91")},
						Published:    time.Date(2021, 8, 8, 16, 9, 5, 0, time.UTC),
						Actor:        IRI("https://federated.git/actors/b1757243-080a-49dc-b832-42905d554b91"),
						Object:       IRI("https://federated.git/objects/3eb69f77-3b08-4bf1-8760-c7333e2900c4"),
					},
				},
				TotalItems: 2,
			},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalJSON(tt.data)
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) {
				if !errors.Is(err, tt.err) {
					t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.err)
				}
				return
			}
			if !assertDeepEquals(t.Errorf, got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %#v, want %#v", got, tt.want)
			}
		})
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

func TestJSONGetBoolean(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetBytes(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetFloat(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetPublicKey(t *testing.T) {
	t.Skipf("TODO")
}

func TestJSONGetStreams(t *testing.T) {
	t.Skipf("TODO")
}
