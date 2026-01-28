package activitypub

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/valyala/fastjson"
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
				f1 = v1.Field(i)
				f2 = v2.Field(i)
				n1 = v1.Type().Field(i).Name
				n2 = v2.Type().Field(i).Name
				t1 = f1.Type().Name()
				t2 = f2.Type().Name()
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

var (
	objectPtrType                = reflect.TypeOf(new(*Object)).Elem()
	tombstoneType                = reflect.TypeOf(new(*Tombstone)).Elem()
	profileType                  = reflect.TypeOf(new(*Profile)).Elem()
	placeType                    = reflect.TypeOf(new(*Place)).Elem()
	relationshipType             = reflect.TypeOf(new(*Relationship)).Elem()
	linkPtrType                  = reflect.TypeOf(new(*Link)).Elem()
	mentionPtrType               = reflect.TypeOf(new(*Mention)).Elem()
	activityPtrType              = reflect.TypeOf(new(*Activity)).Elem()
	intransitiveActivityPtrType  = reflect.TypeOf(new(*IntransitiveActivity)).Elem()
	collectionPtrType            = reflect.TypeOf(new(*Collection)).Elem()
	collectionPagePtrType        = reflect.TypeOf(new(*CollectionPage)).Elem()
	orderedCollectionPtrType     = reflect.TypeOf(new(*OrderedCollection)).Elem()
	orderedCollectionPagePtrType = reflect.TypeOf(new(*OrderedCollectionPage)).Elem()
	actorPtrType                 = reflect.TypeOf(new(*Actor)).Elem()
	applicationPtrType           = reflect.TypeOf(new(*Application)).Elem()
	servicePtrType               = reflect.TypeOf(new(*Service)).Elem()
	personPtrType                = reflect.TypeOf(new(*Person)).Elem()
	groupPtrType                 = reflect.TypeOf(new(*Group)).Elem()
	organizationPtrType          = reflect.TypeOf(new(*Organization)).Elem()
	acceptPtrType                = reflect.TypeOf(new(*Accept)).Elem()
	addPtrType                   = reflect.TypeOf(new(*Add)).Elem()
	announcePtrType              = reflect.TypeOf(new(*Announce)).Elem()
	arrivePtrType                = reflect.TypeOf(new(*Arrive)).Elem()
	blockPtrType                 = reflect.TypeOf(new(*Block)).Elem()
	createPtrType                = reflect.TypeOf(new(*Create)).Elem()
	deletePtrType                = reflect.TypeOf(new(*Delete)).Elem()
	dislikePtrType               = reflect.TypeOf(new(*Dislike)).Elem()
	flagPtrType                  = reflect.TypeOf(new(*Flag)).Elem()
	followPtrType                = reflect.TypeOf(new(*Follow)).Elem()
	ignorePtrType                = reflect.TypeOf(new(*Ignore)).Elem()
	invitePtrType                = reflect.TypeOf(new(*Invite)).Elem()
	joinPtrType                  = reflect.TypeOf(new(*Join)).Elem()
	leavePtrType                 = reflect.TypeOf(new(*Leave)).Elem()
	likePtrType                  = reflect.TypeOf(new(*Like)).Elem()
	listenPtrType                = reflect.TypeOf(new(*Listen)).Elem()
	movePtrType                  = reflect.TypeOf(new(*Move)).Elem()
	offerPtrType                 = reflect.TypeOf(new(*Offer)).Elem()
	questionPtrType              = reflect.TypeOf(new(*Question)).Elem()
	rejectPtrType                = reflect.TypeOf(new(*Reject)).Elem()
	readPtrType                  = reflect.TypeOf(new(*Read)).Elem()
	removePtrType                = reflect.TypeOf(new(*Remove)).Elem()
	tentativeRejectPtrType       = reflect.TypeOf(new(*TentativeReject)).Elem()
	tentativeAcceptPtrType       = reflect.TypeOf(new(*TentativeAccept)).Elem()
	travelPtrType                = reflect.TypeOf(new(*Travel)).Elem()
	undoPtrType                  = reflect.TypeOf(new(*Undo)).Elem()
	updatePtrType                = reflect.TypeOf(new(*Update)).Elem()
	viewPtrType                  = reflect.TypeOf(new(*View)).Elem()
)

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

const imageWithContent = `{
  "id" : "https://example.com/icon",
  "to" : [ "https://www.w3.org/ns/activitystreams#Public" ],
  "type" : "Image",
  "mediaType" : "image/jpeg",
  "content" : "data:image/jpeg;base64,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s0AwgAAgAAgAAQkIcME4UI8E0888888o4wMMI8QgcAAAAAIAAAAsA8kgIw0go0AI0IAw8Y0YAA4wAQwoAAAAUIggQwMEkAMk8QIggYAUk8QgoEogUkE0IckAAAAYwgkwoQ0EYkYIQsQkAk4EsoIUkg4AYYQ8YgAAAA84kE0Ew8EQMEk8IEocwsgQEIYgAUsAYE8IIoEMQAQco4A8I0Iw8IwEMg0g8oEo4koIUsIo84YskIAYEUYI48EM8E4AMIoIIY0ccscQAkcMw8EcQow8EYs0MYoMAg8gsQ4U80oo48884oQU80owoQ08EssIoIMMgc8Awgsc88McM4Uwgw8kIwgs4Ago8gcg8gsUA8UUMUwgYQ8MEU4o4UMcgososocok8UwA0kwQ0YoIA4MgIwIMkUUkEAAIQ408IIwgwwY8E48gII4Y4AwkUgEwAgYcAckU8oAAYMAoI8Y0cA8kk4c4IccAckAQM08Y0cU80UMkcMoAY4ckQcQA40AQwUwQIs4UgQwAkYoEkoUQsgk0Mc8scQIwIw8wAgIMQQsUs4g8skAMUoYwcgIEg4wcw888s488QAcw8EAE4AMEIQA8osgsE4wAoAIwUcYIEw88UwMcwUM0800w4AoE4EcYsgssUE4EAoQwwg4EAg8oUgcMooowYcAAIAEUY80MsYgo8skMwsU40oE00UAkgog0sEkkcIAgQgEg808sMUEsMIU08s8AA0YYkkssQoAUYwscoAAAgUs8k8IkYUYowk8woQIYsQE08AMMQUAo80UkAAAAwgAgUgcg4s0gko0Iwockks08YIk04Eg0k0UAAAQAAAAgQg80MYMcYsMUQkUEwcwQY44UUQA0s4AAAAAAAAAAAAAAk0AIsgs8UkYog8ckEQQ8c84MoAAUAAAAAAAAAAAAgQUcUEsgAcwYYoos8YQEwsUAEsAAAAAAAAAAAAAAYEogwEUsEwIUggkYgQ4EEcsoIIUIoIEMMAAEMAAA0AcIQgAkEUQockYYocQsQsEU4MUkwMI884kUgA8MscUkkocsMcAMgwAsoMEM4MwowYUoQUwIAMIAA8AUE0I840ccggA0ok8Y4oYk0cYUAQ8gA0s4AgMA//EABQRAQAAAAAAAAAAAAAAAAAAAKD/2gAIAQIRAT8QPj//xAAUEQEAAAAAAAAAAAAAAAAAAACg/9oACAEBEQE/ED4//8QAKhABAQEAAgICAgIDAAEFAQAAAQARITEQQVFhIHGBkTBAoVCxweHw8dH/2gAIAQAAAT8QUNiaWDaqkW+sEtMlOiAYS54ARWzg6bOWNgFqgDsn5p+gdsvO/q4ABJFQsSH+bQkiudLmpMs8C3yTHycwZDj/AMJje3gec3ydSd5ZPcB2x3wWx4TjwKlcuQJPu55kfrWRaq4tMDtjp2wtXFPEDJP7ENyoeYcbTd9Rpulst38U2HkcQZDn/wAK9PI48ZR3uQzN/iy7DHj4zpYEIfm2DebcJbG1E4FsxVAyYapsagwC+okVzJeFJlruzrwjgxcr1MlbOufmnlepeHb/AMKHXjp40d9tsnP+IMCVwXmZ2x4h3VhqJHEBdAQAgM3OWQpz8TwXklHigZm91snuj3CyuvxJmRby6lm18BzG+hbol/UJq8meMNxdyw6XL1C+JhwMBzdvwG9MNO0/AT0bfRfc/JDAYThQvBLRn4AvV9Uid/jr48HVjY+e0InyicqJRweYjiSNgTRwoLlAtgwBbo4xuy/VwmI47CYoXI2ie0MYCfUJgWhnxJT1e8dx6nDHEeaZIB6+I8RxN/a4XsxYGwxTv6uOuZ0Q4hiPOTFN42Othb9IYjyIxUCR9cSsfhnEWUcXQkDASJc+fKQDdixXn6l4nNunga5YOjM97YLnYS6fHUnhgOplg4+Cun92jyP1DKAtBNmBCQbcNi6zTkbfEz9ypYMuOod4RMcFj5/stDP7JwC8/MBRI+YvowXxJNBMNC2i2DHEeHxLL6vm0uPXzcJG3aIByWsSHNmHiy6EDTu7vAaxodbESBCfQ2soktfnTQbO8nab3C35XbznMEJxcTeZyDxb74OyAYJDcDiIibCyeOqMKC3WC2QOrrTu3YbmDOFzjLGmPJ7r/dh6ufuya8yo+oKfV1izuzA8a3WmQdZl1XL93O+1yJ6kAL+5vH/SJpf3ZKc5JP8A5W8FhKbt37hG2ZaW7AuTdLgClg73zFzAsRbm5LMLm2Q6dTgJh8+PYtV8sth8FLe13y/g60lTNbVe7X8epZHtnuP4K6ewGSXVOufDwbpwghDq5SE4d78gYDc8SU5Za+CYI/Fgi4kxiJW7WcSGcZrp3PVxX7nwG7GgjHBOLYtvbOwXG32LOBe/ABrHxNzONhe4DjbN9zh1drLJw2ftBswFzIwXLPm2xZ+5jl/2etYTtu/wJI5cPWfuWVTLsYZarS1+Zhnl0qkFysJjJC4Rt6CEKEA0hUDqKoELAj5nfU3xaigCPMSRmEYuSSMfM8c8wqYaLm2//paI1ZyeFjC6T17K7f8Agd8rJQc3R4XI58O5E87hekt4siQgY3PiBjE4YFwExT1E5nfxaxZ/EBzf9T3V/EZ5fYRQDXjQwiA4bc1k35isjCMvKA+UnpSXt/U+pCOmB6tDcfCe3/CzZM/2xy0h9y8S58dL+dhK4CxOju1ZLO4Vm6zD2yRRHpzbIduCHRtgOy/Gb4CwP/QjeATXoyYf00WvLdhLW9MxwrAHzZ2+lzjUCdObesTaSArEXhiSUmfn08Lr/um+pX8XJqdn/gKGgP6nWDf1I+HFmhwMZhEjvVtjpIXAhWBuDYcij5L2hItBCDzP3LVLt0SHTklYcSKKI3UFjLm0IDuB+Jb1IHiRkLJQjZ59LZbH4mE/MK18fhr89cz/AFe0mng5nhN8ClOtsV+0NADJzBrbgguwcIWeJpdziZW+7UD4hA656hFz/d6r/d1k4kc+CxtlptxlYyTisgmv/RB6RjgjPmcbLUuixlTlIrmSEQzLsZHzGzONm7Th/wDAB38U+JrQ+1kQy5SwnD4QIum2Dz1bluCNv9xgm2GJwZ/MNi/7vZL2iyDXmftpasI8wX0X6X6x9Yz4Ma1nKL0h+JUA/wAWm4M85Nk1Guq6mZabd54f+A1tYmli6G95PZE747+5h4FshoNnn1D8ATn3jY+0uJcLabVuycLk8m83Ew4iAB4Rv08S/r4MfAYUqXwJ+LvBa3Cb5EVI0iT+y5Z/3R347fjttqcyDcZcwW12diyOWeI7D4kLF5nu1qW60ZiXa2QhQnW5tFkBnm16cHzA+l/mV+SPvvr/AOzhS5x4/i1J+KWTHxQgfU+4Wy5J5/8AZM3izhj/AODLg+JGo4sbWrnh2WCc38CnxbSNG5VdoLSKPWO+DZwrmDkobbs5w3xsRg683tzoPjz9D+4R6R8Kdh/jzO0knhkkn4AaWrxJuEr4CZvF9Eg/+BWTDlkwv6jAT/U9w/1Z6Dm4B8XHA7hekwTGXY6bhTYNI2SM5l13mxMDPI58j+PGrA/uR7bvYow4E9+eH7kEBD4Z4fByuZHWJLiA6iwS8gsxHn6l6H/LdVZMg/h0ChNxu0/HQzf+V0JLrz3Jgkck8mOfxc4Nt9r+ruSeH8H+n9yfSy9JExs+g8DOdNciNTRuyM2xfMEUVsz8kgmGEHAuI8I9jrrzrcm2s+38A+W6+piW5paeFllxckspDm7rBXODXptmo6kjouLzcsR9+cWAC4mfBCBDEeM75MfB2zH/AAgAEzH35yPtKALAePXxOww2GKeRoINRxYPPU5xxbH4D3n+4fvYdKB4pnNYc08Rea/u3urcOYlEInltuVgnDYmaSyFwW4IAMODymTnO7Rj6yHq0X585TC0fBRhhMpbkQk3N6lg+A5kpD0wTkBE8GaLuF/k+SNfMZAQjiAwo2ep8Gx+YxwcW6TH0LMh78i5+br7qd6f7nnW3f567D2H1dSjiXJR8rp1cd3yPNyic4nUndtFmiGssgW+W27xs88WAv46f/ACjqC+p/X4IV45bDplg9SPCzPhtNRcbIbgJa2F6qy/aR7ZVefIo6QeCLs1Ldr5XYjuSzWv7k3UydvOS21JaWv+7tlZd8vBZOOTs7nHW02rdYm+Gk+Am7Ld4eV2gxOqTnZ4ECvc3NgE9EeoRwPj/IAxmWPgz1M/iDy4GfUNXxYeAu3+jv5tLXz5M4HMZ3ANfGXuDlqzCChw2bBf0tsES+oGEnhMicbfrKPUN/zGgfiEz7mfB8FYDcTc1uZDwaQw3b/VDb9oOYfhhP3EcDO9R4pzgnNg/KSA15tIPmepaD5lLIPuyNMgj1H+h/mTCPvwZSylKZcWO3dzaLsso8NhEzy/4j5eE/wj1Jkdw4ksSY+KZc22Oin9Ebub9XbgfqWdnNqfuMH1BzMbmbnu2W4D4h5Dwf9jx9T+4R6RuwQv8A7C/+gkMHnz2jn4bhDt8wxZSzJDiWDYq23x85Zsp9/wCI5YSbJlnjHPPa7u3g9R3dSbIaQut2g3ARxBfuywX8T4T/AFE2Ig2nhoKYepGc5MtRmdHM0cH1CxnwsMgkp7T+/FxP3mnhYop7TP3svfHxB9louvnGv38GYHzZgyywtPklbDMXb/EIM8JsmXaXr8evAukfLJILpPd0gXQuVgrj3PMH3lAdRF1AAC7AHAgHC6sJC9Rhi/a0uvgafgUc33z97k7t5h4eew8EmB5udR3LFhOmLqxe4mI2O3G2b/i9fPa355uG7Tu/WTmx8GDCHMbPqfhJF08NxiusKHncwM9ts5y4rsXqfQ4GJ5k8ImC9K1bj4hMtbhGw3CU2LcbbcI00wMkEQdyKxndLC8ItQ0yXdYHtWzu0PHS4G5mGP+J6LXe+cuSPVhPPrxTxYQR4kg8AeM2P3AEbMMdzb0FzNxaixoLJyeFuXhmg3Am4LgLJxGbeeBjrGCceE7ICR30WE4+JssTwE2WV0ubKaZaep/iZuIlvg6NzMc/yAzh6gt0hbdocuH14yPFN8PUHMEm2HmZQ2773c7HMHCRyfMltTlzBDB9+bIIHZRr3LjkuJ5Y9Ri4bRi7WU8W8ycmn12BcEpZYZzauguRCIm4pSybcKiSAbD1GxoS0ujDu9/8AAdkeLLPGWXayTPA+AghrDiSSSDmxsh9bmc8SYS3GQ4IN7EchZXDDgATwlMIN3mzLPxO0bEkk1yvVET2fg2Rf3cba77peNwr1KBJr1F6EaCzCVI6WuB/mzX/ebbtydx+Gc1nBhDDxPMmfgPqHgax4842TB4yfGFkGxjxZZ4Z78h7YXaxIp0L3PTYI+zEbHTsRkIk/cIZZtkB6Jz5jhcLYXuBBpHIBheCk1ZDGi+JEc5BinY0Ym/tkmry8c3CPC2OuzMDJWLBp1a+p/gl8hnDlhsNYevAbImJnnj+rBJbbJgtEyYoDkTMJkmQTGeCA3X9Wno3HuyDLtL0IbnV2BCxsbISIeJsXm5xyMBbq6RT/AKSvbZT22CZF2z7UhLjmcExl2xLMPLbfFgSwNYIwYOZ5baR14OdR5F3LVxZ/r/UPjEgA3JvvIZdYyRwXL0aaMzqwrt4PdnEwELLTDqFwMhseM0YYQbFXXU+AQx8HLCNunl0BsaZ5yCNC9R0pa3CyWyyzmZ3luXKInN2fQLhBSZTUck7YWl6PBCZZhB5she+YAWRGm25J6Pi0LVsYh9bY6uTq+iOcLlttMFKeT+AeJSlx4AVEgfTY+A6sPcvAYzlti3e6T3YzYmoH/wByV1b5eAyyWUDYlCIlblZBkpO2WN5mGa9s8cDAmtxmI3Nykkzxy48WY0kr9XSlistFlMt4yTFiwq62Dz0XC5RJ08Iewf4lzp/Up6T9X7r9kHk4Mwz4y+hOYM48HKLYkh2PiPFaxSRRG0S7WMq9x6sEp7mZMu673MmtRvmZ2ts+LMiGEEz227Ll3nscSnNYvfhaMPu7yHUztbNeW7eEgyQFsxz3ahFn34UZqE6xInwuLmzdTZ/EPG6LIXc+P8XEP8R8R4xDSTFlvgeAYoE6inhBrJpeLUy48et73vJrPdDqUvkuKyY8TxDvjW63aUk6WEw8PdPmOhd7pCz6hY2eFb3t2h39WAAI/HNnTJkpHemSPJHZU4bcOkOPlYZ79/4cHP7TjF8BREp7sJDgsF1iZl8NpRxHzxaOrHwfqG7d9vPpgp1CazxTxOebo8BngJNkg4z7vgaFxKSOIbGNuM7kOPCQ8Jskv1sR3kTtWdnudNvfHm9H/PHs+yNALACOlwYzp/X5bsx9s/I/uV7dvv8AHnAutyOMTmzuTF9wOdhQhlllLK6RYwcx7u96w22fEIXUlo4tjOMefAbJHjiBuMgkEOUeYCRlkW6XQuEdxs8Gff8A6vs/1fZ/q/8Ayr7P9T8v+r/8KS7R/F+lw9X6+Jovh6niQ+rfhfz+LXmeLKw92VzZWd4mxsV+EY+WZrszCWhLKXkMBu8t1gnrJt1hr4uGwNj45lGwvAJlnnGfAT0+PSSkE4sa93IlyQ0IdwRg5mJ/iQ7B/MC6Ff8A5FjXQ+pVauFizs/ce4S7RueCSjmycGW2gGow4ZhZx5JHc2BbT3Z/CYBD3Dp4uzgPcePNrvNss9YTwePJuC42Yug2xAlboz3Zx4DPCbJ4ermQYnNv4ebHDLkkeEzMhDnwYHrmH0P5hvqYvv8AAXh1vcfx4HP3Y8WvqR2QicAP7geOP6stk9zeThyVw5Hz4AYljfVM1GfqFlmqOIS4snu34kNMi5yhBKdyrLcZHebnXPweYnidbhsFOAJHpbpCZ4D8Msiiebc4/B8uQs6XFJ4F28JbZSTZBEeZH/8Apew2wYcFvIe2GLciCJKqyW9irZSdbp5bDs9cWvh1g7Gz4RGlOonEicz6QxaUdxCXNilxYbcz5J3fwbkuY75AHFMyTvgxsukln4gIDNzKvvwN5kxtsgDJ22wZaIQQ4g5gwhjHbXePlv3ufuV8F+raWywyGku0CszHCONs2Y9eX+5PQ+W20kfMicZU0ANvOLeRCOrBBmCI48QA7LYuVyPCuw8h0hrYMsryy6QlHqB4XFtXyG2IA8C+CdHwNEHO43EOY2kxzo8QJc4w+O8Kfm2WWPkNt26sy1scb17tlZa2rAgduxfF8nCaxPA+LxwgsvdySXJDSMCxPFiKFa1XIlELj4+/jqN6PGN5EC6iCVzQtp+A53/rQ6tCWeDhzbk21lmNziWl8GLPBlllhHwWxsrYRfctLLbJnvLIS6iP43Ow18fEjlIZxnlEVhP42GaEqRFGpYbRHyxWy42koc3Pnru0rnabhdZcbePOJghhbsuPA+zQPPbOXCkbjcy/ph26Q0XR4N04si5IfNwtt58BlkWRDI5+YXfUtgTSBlNhOeZNFs5LI7xlkowx8JSNjFm5xaOpOTOP68AXJAbECZaDMnZbRELTN4uEMWW4S5v6nQN5gtYVuWRDHxcCjAaouGeIL7jDuLO7U8z0tXJhLmzrZnhc8YEeyQ3jPE9pbKzhG5HDnxcpkyzhlmz8I3O2c+NuQxvmM6dloE8kukkNtwhRluxhw7ljfLf2oYvhW1ZrS/5rR9E/AtrC+LM27mymGxvRKzC4B1s7izXF+VgoujFrcERIazB8Qgsw4dXxMZiMHcvDYhhw25JYy3EqYs8gyd4S5iyM1WyqZJ3ZHLu+VJcC7EbuK1/4tqSL1Z7ywRY2kJUmeZZ4lQmH9NhEIuE8hjL4g3Mh4biR+0J1PctIsXSGAfBaBmtpCgwuXrxDX7fiFmXXgwHxRBqQtaWZNiYT4svdsObJ5hdpyBtxt52xdhcBbIOr5sLhbmTrAmNj9VvweJSBYLtP4DZ3RbvVqQyX9vgdE8gQvBbDDqyx83eTMFg/bPmGw1mRJliJlnMSNL0UAPMjz5mNTYdAZBOnMpgcLdmx+4Go3IyKu3EQnNBl3obOXGed1O7neLB8MWMNz1zEwtPwOHhqYoIwwxlIc9bG477raxJzO7XcM9bgM7ozhEdCsZQtphfEg9T4pMuH1b2n9znlhHqeIKMSCss0MnK83OkgdSjtHF2AuwlN8mtyl7u2O/DYkLvm2nDuPFmMaNvZllHAO2xFmy9uNnJ+HiR3n4SAbcB+dG8l3+yTJxH2sSMtvEQ8WqtTcYod2jGJ8lN08awp4/8Af+GqfGe3Kfia+LhV9sml7UKTO43MBObhbipISRXixeNuvlg3aQHIjpYurQsR8oA3bLN8Ccwp1C9yry92Dco2GbJUvhy8y9VzUQ2bS/s/NN75h+CbQk2HgY6t2SFoXOjzhe/hd8HUlg48bki9rmXL3PDuSMLT3cxsr22Vh0uQl5Q3ckIXPAs47D9qzws3IoYBPU7xhhzdn/O1eAfEOJ7ZNkyUX3JxtxrB/VoM7t4hLxMkYuqWg/JKHaF9T+/yNxOfm5EyeK7X2Rd45LjM6tzJ20tyyXwIYzB8x3OPBd8XZYbBvcSw7TDJsS+cjuXE+SOf4XLuHSw6QHC2CXA8QVwe4WRNbq0OZZTru4XT+7RmfXknMm9z/O5TEaw/dcS5G5c+/Cy8E062zBdNg/XFsvg68GHRxuq9/iJxIRT+rHfBTeG14YMiyFdzhbaF8LRXDPKJiwJTjzvc3Oy6zyHNEHHwFvg+LKPR4AwwuxPVtiRxWBtqBsMDbOyribMrxWUp6tZJDY2Ni2NnxSm/DERepwux5xM023pPFOG3VZj58Dy8Cmc8wn4fj8cp8l2DabxZgcWbdjO+LIebosg42K7Jy5c0uz3LhYPdhae/CNjVwU+XnwOPgHxr+DgIKh7iXjwTkv8Ahn4Mf7FxNzksX68PYfq03lwIwKP3Y+StfA4hrJng5MnXiOe3yF2MIJCEUbMYzw6i3iOJZTW3wGyTnmN+H2Qg0d/AX5C3hYyaHJ1JsacdmJ5etsmPaCgrCu93Mg+ZPzD8wHu09yl7nCp4Hju7zxZDfxLmLbxssdr2Pd1fqHeef6g8Ewbdm6k3FQQs+7OyywCYGEZ5QtXGwvNgsTYcOrm1hZFwiilwkzeZiluRh5A8FB07i2PJlwXw3lIdlkw8+JyPF2U4Ln3t2XslJgbtw8P2zsklRxZm5DwunEuUtXy8hK/K5ZMPMuIs+cs0TiAH8T72Bt+5wIF1bkPuy1lk4yAR5PAmWXFMW+vFLOY0J2MOURyyWYkvjSN4vWSAjxZHPNynJaRCD5JOjfL/AEtjNExdjxal8eltkagE1zYHdsyrufM25+5TKZH1cnUJnFxwy+8IOXnq157XTwNyukWjohhZtkbambJAeJxFsuIA7ZxnkPqHCMER4ssks2cc2+OBZsK7KYE4QozowPmcOeIwl2fgYgzhtl89IwoWa/uW2vVxI2xE252kMXknjyrhhy/4OTjfGsYBdbvooeoLWHLr41s7dRstzw8eAhtwvESCFnP3Hyn+V2lxGr46RObjliviYAsuOrhGaQE5MR31ZhGYpNl5jwNvWsgcifx4DdGcwPEhQTmNENT5u0or5nGPHh7Qg83Z/wAA4wwsxYT4xI9zO404bumn2XFuiDiSDx3hvgJswz355ra4L4DtbbxEQLCcNsrgwpeRyQSBjKcngXnMnQifJ8JvfPh7T7npNAoKR9Q1JMyS/EzZkMcjh4Gy4hux5f4RLfUIWtkqcZd7g5XthyXeLfNuE5CFfxiv87ljzDhBz4Bdth4bZaeCGse8m2qOEms/RKJBJB4mazI8EbBwGJ9+F+Hc8TifXlQNeCfi4m0WPV8XCiMdLe6o/jxYWcuD4vPmdHg6MeWe/wDAi0gDL20qvP4a+FLBMcMiA2/fjtDmPCSIWcyeBzdYxbQPMt5TMimkj48CCOEnRzPINWOfVhbJYAjYwY7LrwH4uy97pYDdlsMSKFoQ5/CjMJeIS9S4X6S7HhuRnt/M822M9/4MGDCW78HvLpGXN2iaQzydPB68b4B1CYdLc5nbCGBZrOIgIQyjHYsENpgaEbUMh7BEkZb47A+/w7J93db6R4hm38Gtj93a6Fg8Zo8HieN+k8vg8R7u789mzYf4RxhGQXV++A8rKQvaSfiaOcvSYD1usMZw7ZzJE1m3bGycJbiEH0ZstxsvqSvhDjKCwkMW5fEIDmT8yWY+HE/w/B4T7ldlzut7+C/6L08Tn4OtxT1h5hepQ7u7/W2B6ntLtliu0o9pDtMoQCI9+c5HBSza5I5Mhuha46PA4NnGWIueAXLeWWz0Ty8GuFiIWB98/hu8dF33bco8N3inu/3el6WHiLztDxDz4q6bcjPK/wC6iYxkmRFGy+dAc8ztdQOmzX1zeYNaRrnZG44aytRJe2FwLks7WAO233k83TaIf35UO0I8BFfD23Iw7Dhlm2ASSvx9SasNgzwUzDuZ/vrsSoOC1IkI93ys/wAwsH3Z3FvNLuwkOkywwnuYH9q4fEgBs6xF25Zw9BiOM4um/aLb11+LLrCZdW5PdoWQ26WXC3faaW0yuGeYDF2nOBkaxk7jaoWbR/kN0pE7M/2UdGdGorq+FIvBH0OQnpAepARFsWsju/MMDOpN/ErCsgRxDuy1zYZMoXP+BMbPMybPC9M8442ZYjGwQ+PSW+D4neDbJQxwsyOFsIWVEsNp/g7wmjFn6k6jnIlHp/2ty4E8AnJEc4NgXRY7ZjzBbztxDSJydWh2d67mc2EcnAkZB4ADptsfKzUuDLdvAbl2sVwtjTr5t2Q8aut2icHYFoWKixgTzO3Mu/gYZvNxk3Pi/Rlrvv8A23ljCOIMcXTXSxLYiOb3IB3DJdAwPtaFcyDGpkBADkiY+7lOYDy5+bRbRjmMFgJIzFhthtmzbVu7EnbfE/Akc6tY2wLErCKHmVPPmPuSJMIXGQx+klUGRF/2wEE8XAHUONiZFhyTD9z4jiAzLJ1OeJc0mQGuWIg4I4cWTIBky756+J4WyBkaWPbul1MqE8DFDUYIs21+OLRXGIsAs9uYpeRcMDN8KJaH7lc4S6SE6dxlzOLmgl2O5MBrvCRCPkF6NuvVx9v6u0Lr/SCOIEEn9oaImCstLge9Yi4cQ5ZDYU+aTYcAyGcW9t6vovqs1us92JbeGVrStIgcyfQ/VzTlCuEM8CGLgAuAD3H+tdLeSgMVDXZ+iQjjcyYWPEypK2yQ7gCSO5r6sfRPxHVpuCGIgwKVFWRrsjiZjAafVtn3IsLA5Nq9XFinznO5uYdUpwM5wN7VZ9v9Wns90j/AawOuzYurMlwsFKstknTZfgnkEc8x8K4S5gOTj6epXsWkjG7rZ8gS5L4Y1bgXFWlg8CEULgxA5hnRDJiME55sD9WJHzPU3O3mADY4AOLRbiAlTH/cUNPPjDsTQhPRzZDu74Bp82mHz4AB9Eyx4mbYrRuRPntnQw4vKyWLEEQwNTOIWjhtoXhBHCAGC9UmDhtsBavwYrt2jmUn7sQQyzNq/uAnLhMC0DZ8j3Y/jEV02OIyjeH1CgXF2z3Zsc5kjnhLPNvV2fzA8I0t3zbDY8FOuUn33OBWSWoyqMKBvcAdLrtAhgcMrkqc3mIgZhCjc4vuZc7Pd8hZA6sIkA37jBPiYb5vq7dJe6SLW24X5gDxZJEgTbAcdPMf3YeYkWJsBTjYyrZW2Uzbfxzg2Q/3dK9XW6zm2x/dkI4JsxZubYOnPh6mduC5zbBdvkeHozHVgcW7OsxBjCrHE7FAQnWLL9r4P9haFPU8Ebt9Ti3tjSkeOIACXS1I9XEfcYFygP4yAHiZN+ZUzLOsIGQwRXWKpZs4L4m2coyJE778fyyxnJEZnyy5A93VINs6JFeQamVHnm6RnECOuJ+V7/Ldpt8y/qjNuLsSMcR2EjLd9Shb4wI+5Lcg03JeTvrmD37aNtharlBxZZW3BH3a7YlcyLtjAjtkKD7tfQh8WClx8LOfUsIYahvJ5sAnOeOPNwmlp1CkjD9VoQgLcpO7LGNuyGE/QNn3HESwcEpNNnXRJcqc9xfPEvE92Bx6gQpmWzy4bIeZnE/K2DWZ+cbGZ9WoHLffy5TJjUzGQC8hFy2S8t61ZJpd+vYQG43CPcQZpAyB7iHRFXJJsTn5kPXiz8z6mXxHYwThLRuRs13z1EG+4WvCNFK/qyAE3EOp4A3bIbDNDlYANgEI2rZkhwfM3EUIk+7ag6TIGjcWu/i9aMLrSgC+2Ycc2TD4kmmSMIwzlcRC/YtoGSPiwiOPvxmM3mDherZPhIuGN8rD5Fhv1MDeJdV/IeDIvrB82f2ReCn3lJQur5cv2lXnZNgeRrRlB2k2DNgUIx0tkH7tCEyKldFrD0kubgFlZ103EjuLqkZ2t+Bu6/MbT0vXdsd4gWaEKLMpXUp7Q+nBtVkKrcxRBElgrMarPB+7Rg2nRtN1KgdyICQa3/ZYAy510CFrD+rSwYTCGrtm0+Jsu9SaH34L9uXA+JEYyIljYeFxAkg3gbI7tvn+f//Z"
}`

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want Item
		err  error
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
			want: ItemCollection{
				IRI("http://example.com"),
				IRI("http://example.net"),
			},
			err: nil,
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
				ID:      "https://federated.git/inbox",
				Type:    OrderedCollectionType,
				Updated: time.Date(2021, 8, 8, 16, 9, 5, 0, time.UTC),
				First:   IRI("https://federated.git/inbox?maxItems=100"),
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
			err: nil,
		},
		{
			name: "image with content",
			data: []byte(imageWithContent),
			want: &Object{
				ID:        "https://example.com/icon",
				To:        ItemCollection{IRI("https://www.w3.org/ns/activitystreams#Public")},
				Type:      ActivityVocabularyType("Image"),
				MediaType: "image/jpeg",
				Content:   NaturalLanguageValues{English: Content("data:image/jpeg;base64,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s0AwgAAgAAgAAQkIcME4UI8E0888888o4wMMI8QgcAAAAAIAAAAsA8kgIw0go0AI0IAw8Y0YAA4wAQwoAAAAUIggQwMEkAMk8QIggYAUk8QgoEogUkE0IckAAAAYwgkwoQ0EYkYIQsQkAk4EsoIUkg4AYYQ8YgAAAA84kE0Ew8EQMEk8IEocwsgQEIYgAUsAYE8IIoEMQAQco4A8I0Iw8IwEMg0g8oEo4koIUsIo84YskIAYEUYI48EM8E4AMIoIIY0ccscQAkcMw8EcQow8EYs0MYoMAg8gsQ4U80oo48884oQU80owoQ08EssIoIMMgc8Awgsc88McM4Uwgw8kIwgs4Ago8gcg8gsUA8UUMUwgYQ8MEU4o4UMcgososocok8UwA0kwQ0YoIA4MgIwIMkUUkEAAIQ408IIwgwwY8E48gII4Y4AwkUgEwAgYcAckU8oAAYMAoI8Y0cA8kk4c4IccAckAQM08Y0cU80UMkcMoAY4ckQcQA40AQwUwQIs4UgQwAkYoEkoUQsgk0Mc8scQIwIw8wAgIMQQsUs4g8skAMUoYwcgIEg4wcw888s488QAcw8EAE4AMEIQA8osgsE4wAoAIwUcYIEw88UwMcwUM0800w4AoE4EcYsgssUE4EAoQwwg4EAg8oUgcMooowYcAAIAEUY80MsYgo8skMwsU40oE00UAkgog0sEkkcIAgQgEg808sMUEsMIU08s8AA0YYkkssQoAUYwscoAAAgUs8k8IkYUYowk8woQIYsQE08AMMQUAo80UkAAAAwgAgUgcg4s0gko0Iwockks08YIk04Eg0k0UAAAQAAAAgQg80MYMcYsMUQkUEwcwQY44UUQA0s4AAAAAAAAAAAAAAk0AIsgs8UkYog8ckEQQ8c84MoAAUAAAAAAAAAAAAgQUcUEsgAcwYYoos8YQEwsUAEsAAAAAAAAAAAAAAYEogwEUsEwIUggkYgQ4EEcsoIIUIoIEMMAAEMAAA0AcIQgAkEUQockYYocQsQsEU4MUkwMI884kUgA8MscUkkocsMcAMgwAsoMEM4MwowYUoQUwIAMIAA8AUE0I840ccggA0ok8Y4oYk0cYUAQ8gA0s4AgMA//EABQRAQAAAAAAAAAAAAAAAAAAAKD/2gAIAQIRAT8QPj//xAAUEQEAAAAAAAAAAAAAAAAAAACg/9oACAEBEQE/ED4/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")},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalJSON(tt.data)
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) {
				if !cmp.Equal(err, tt.err, EquateWeakErrors) {
					t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.err)
				}
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %s", cmp.Diff(tt.want, got))
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

func TestJSONGetTypes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want TypeMatcher
		err  error
	}{
		{
			name: "empty",
			data: []byte{'{', '}'},
			want: nil,
			err:  nil,
		},
		{
			name: "single Activity type",
			data: []byte(`{"type":"Activity"}`),
			want: ActivityType,
			err:  nil,
		},
		{
			name: "multiple Activity type",
			data: []byte(`{"type":["Activity","Accept"]}`),
			want: ActivityVocabularyTypes{ActivityType, AcceptType},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := fastjson.Parser{}
			val, err := p.ParseBytes(tt.data)
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) {
				if !cmp.Equal(err, tt.err, EquateWeakErrors) {
					t.Errorf("JSONGetTypes() error = %v, wantErr %v", err, tt.err)
				}
				return
			}
			got := JSONGetTypes(val)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("JSONGetTypes() got = %s", cmp.Diff(got, tt.want))
			}
		})
	}
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
