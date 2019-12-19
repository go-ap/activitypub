package activitypub

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestActivityNew(t *testing.T) {
	var testValue = ID("test")
	var testType ActivityVocabularyType = "Accept"

	a := ActivityNew(testValue, testType, nil)
	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != testType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, testType)
	}

	g := ActivityNew(testValue, "", nil)

	if g.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", g.ID, testValue)
	}
	if g.Type != ActivityType {
		t.Errorf("Activity Type '%v' different than expected '%v'", g.Type, ActivityType)
	}
}

func TestAcceptNew(t *testing.T) {
	var testValue = ID("test")

	a := AcceptNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != AcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AcceptType)
	}
}

func TestAddNew(t *testing.T) {
	var testValue = ID("test")

	a := AddNew(testValue, nil, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != AddType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AddType)
	}
}

func TestAnnounceNew(t *testing.T) {
	var testValue = ID("test")

	a := AnnounceNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != AnnounceType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AnnounceType)
	}
}

func TestBlockNew(t *testing.T) {
	var testValue = ID("test")

	a := BlockNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != BlockType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, BlockType)
	}
}

func TestCreateNew(t *testing.T) {
	var testValue = ID("test")

	a := CreateNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, CreateType)
	}
}

func TestDeleteNew(t *testing.T) {
	var testValue = ID("test")

	a := DeleteNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != DeleteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, DeleteType)
	}
}

func TestDislikeNew(t *testing.T) {
	var testValue = ID("test")

	a := DislikeNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != DislikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, DislikeType)
	}
}

func TestFlagNew(t *testing.T) {
	var testValue = ID("test")

	a := FlagNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != FlagType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, FlagType)
	}
}

func TestFollowNew(t *testing.T) {
	var testValue = ID("test")

	a := FollowNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != FollowType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, FollowType)
	}
}

func TestIgnoreNew(t *testing.T) {
	var testValue = ID("test")

	a := IgnoreNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != IgnoreType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, IgnoreType)
	}
}

func TestInviteNew(t *testing.T) {
	var testValue = ID("test")

	a := InviteNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != InviteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, InviteType)
	}
}

func TestJoinNew(t *testing.T) {
	var testValue = ID("test")

	a := JoinNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != JoinType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, JoinType)
	}
}

func TestLeaveNew(t *testing.T) {
	var testValue = ID("test")

	a := LeaveNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != LeaveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, LeaveType)
	}
}

func TestLikeNew(t *testing.T) {
	var testValue = ID("test")

	a := LikeNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, LikeType)
	}
}

func TestListenNew(t *testing.T) {
	var testValue = ID("test")

	a := ListenNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ListenType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ListenType)
	}
}

func TestMoveNew(t *testing.T) {
	var testValue = ID("test")

	a := MoveNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != MoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, MoveType)
	}
}

func TestOfferNew(t *testing.T) {
	var testValue = ID("test")

	a := OfferNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != OfferType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, OfferType)
	}
}

func TestRejectNew(t *testing.T) {
	var testValue = ID("test")

	a := RejectNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != RejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, RejectType)
	}
}

func TestReadNew(t *testing.T) {
	var testValue = ID("test")

	a := ReadNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ReadType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ReadType)
	}
}

func TestRemoveNew(t *testing.T) {
	var testValue = ID("test")

	a := RemoveNew(testValue, nil, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != RemoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, RemoveType)
	}
}

func TestTentativeRejectNew(t *testing.T) {
	var testValue = ID("test")

	a := TentativeRejectNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != TentativeRejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TentativeRejectType)
	}
}

func TestTentativeAcceptNew(t *testing.T) {
	var testValue = ID("test")

	a := TentativeAcceptNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != TentativeAcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TentativeAcceptType)
	}
}

func TestUndoNew(t *testing.T) {
	var testValue = ID("test")

	a := UndoNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != UndoType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, UndoType)
	}
}

func TestUpdateNew(t *testing.T) {
	var testValue = ID("test")

	a := UpdateNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != UpdateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, UpdateType)
	}
}

func TestViewNew(t *testing.T) {
	var testValue = ID("test")

	a := ViewNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ViewType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ViewType)
	}
}

func TestActivityRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	a := ActivityNew("t", "test", nil)

	a.To.Append(bob)
	a.To.Append(alice)
	a.To.Append(foo)
	a.To.Append(bar)
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", a, len(a.To))
	}

	a.To.Append(bar)
	a.To.Append(alice)
	a.To.Append(foo)
	a.To.Append(bob)
	if len(a.To) != 8 {
		t.Errorf("%T.To should have exactly 8(eight) elements, not %d", a, len(a.To))
	}

	a.Recipients()
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", a, len(a.To))
	}

	b := ActivityNew("t", "test", nil)

	b.To.Append(bar)
	b.To.Append(alice)
	b.To.Append(foo)
	b.To.Append(bob)
	b.Bto.Append(bar)
	b.Bto.Append(alice)
	b.Bto.Append(foo)
	b.Bto.Append(bob)
	b.CC.Append(bar)
	b.CC.Append(alice)
	b.CC.Append(foo)
	b.CC.Append(bob)
	b.BCC.Append(bar)
	b.BCC.Append(alice)
	b.BCC.Append(foo)
	b.BCC.Append(bob)

	b.Recipients()
	if len(b.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", b, len(b.To))
	}
	if len(b.Bto) != 0 {
		t.Errorf("%T.Bto should have exactly 0(zero) elements, not %d", b, len(b.Bto))
	}
	if len(b.CC) != 0 {
		t.Errorf("%T.CC should have exactly 0(zero) elements, not %d", b, len(b.CC))
	}
	if len(b.BCC) != 0 {
		t.Errorf("%T.BCC should have exactly 0(zero) elements, not %d", b, len(b.BCC))
	}
}

func TestBlockRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	a := BlockNew("bbb", bob)

	a.To.Append(bob)
	a.To.Append(alice)
	a.To.Append(foo)
	a.To.Append(bar)
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", a, len(a.To))
	}

	a.To.Append(bar)
	a.To.Append(alice)
	a.To.Append(foo)
	a.To.Append(bob)
	if len(a.To) != 8 {
		t.Errorf("%T.To should have exactly 8(eight) elements, not %d", a, len(a.To))
	}

	a.Recipients()
	if len(a.To) != 3 {
		t.Errorf("%T.To should have exactly 3(three) elements, not %d", a, len(a.To))
	}

	b := BlockNew("t", bob)

	b.To.Append(bar)
	b.To.Append(alice)
	b.To.Append(foo)
	b.To.Append(bob)
	b.Bto.Append(bar)
	b.Bto.Append(alice)
	b.Bto.Append(foo)
	b.Bto.Append(bob)
	b.CC.Append(bar)
	b.CC.Append(alice)
	b.CC.Append(foo)
	b.CC.Append(bob)
	b.BCC.Append(bar)
	b.BCC.Append(alice)
	b.BCC.Append(foo)
	b.BCC.Append(bob)

	b.Recipients()
	if len(b.To) != 3 {
		t.Errorf("%T.To should have exactly 3(three) elements, not %d", b, len(b.To))
	}
	if len(b.Bto) != 0 {
		t.Errorf("%T.Bto should have exactly 0(zero) elements, not %d", b, len(b.Bto))
	}
	if len(b.CC) != 0 {
		t.Errorf("%T.CC should have exactly 0(zero) elements, not %d", b, len(b.CC))
	}
	if len(b.BCC) != 0 {
		t.Errorf("%T.BCC should have exactly 0(zero) elements, not %d", b, len(b.BCC))
	}
	var err error
	recIds := make([]ID, 0)
	err = checkDedup(b.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestCreate_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	c := CreateNew("act", o)
	c.To.Append(to)
	c.CC.Append(cc)
	c.BCC.Append(cc)

	c.Recipients()

	var err error
	recIds := make([]ID, 0)
	err = checkDedup(c.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestDislike_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	d := DislikeNew("act", o)
	d.To.Append(to)
	d.CC.Append(cc)
	d.BCC.Append(cc)

	d.Recipients()

	var err error
	recIds := make([]ID, 0)
	err = checkDedup(d.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(d.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(d.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(d.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestLike_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	l := LikeNew("act", o)
	l.To.Append(to)
	l.CC.Append(cc)
	l.BCC.Append(cc)

	l.Recipients()

	var err error
	recIds := make([]ID, 0)
	err = checkDedup(l.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(l.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(l.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(l.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	u := UpdateNew("act", o)
	u.To.Append(to)
	u.CC.Append(cc)
	u.BCC.Append(cc)

	u.Recipients()

	var err error
	recIds := make([]ID, 0)
	err = checkDedup(u.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(u.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(u.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(u.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestActivity_GetID(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), a.GetID())
	}
}
func TestActivity_GetIDGetType(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if a.GetID() != "test" || a.GetType() != ActivityType {
		t.Errorf("%T should not return an empty %T object. Received %#v", a, a.GetID(), a.GetID())
	}
}
func TestActivity_IsLink(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if a.IsLink() {
		t.Errorf("%T should not respond true to IsLink", a)
	}
}
func TestActivity_IsObject(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func checkDedup(list ItemCollection, recIds *[]ID) error {
	for _, rec := range list {
		for _, id := range *recIds {
			if rec.GetID() == id {
				return fmt.Errorf("%T[%s] already stored in recipients list, Deduplication faild", rec, id)
			}
		}
		*recIds = append(*recIds, rec.GetID())
	}
	return nil
}

func TestActivity_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	c := ActivityNew("act", ActivityType, o)
	c.To.Append(to)
	c.CC.Append(cc)
	c.BCC.Append(cc)

	c.Recipients()

	var err error
	recIds := make([]ID, 0)
	err = checkDedup(c.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestBlock_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	b := BlockNew("act", o)
	b.To.Append(to)
	b.CC.Append(cc)
	b.BCC.Append(cc)

	b.Recipients()

	var err error
	recIds := make([]ID, 0)
	err = checkDedup(b.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestActivity_UnmarshalJSON(t *testing.T) {
	a := Activity{}

	dataEmpty := []byte("{}")
	a.UnmarshalJSON(dataEmpty)
	if a.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", a, a.ID)
	}
	if a.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", a, a.Type)
	}
	if a.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", a, a.AttributedTo)
	}
	if len(a.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", a, a.Name)
	}
	if len(a.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", a, a.Summary)
	}
	if len(a.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", a, a.Content)
	}
	if a.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", a, a.URL)
	}
	if !a.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", a, a.Published)
	}
	if !a.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty StartTime, received %q", a, a.StartTime)
	}
	if !a.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty Updated, received %q", a, a.Updated)
	}
}

func TestCreate_UnmarshalJSON(t *testing.T) {
	c := Create{}

	dataEmpty := []byte("{}")
	c.UnmarshalJSON(dataEmpty)
	if c.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", c, c.ID)
	}
	if c.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", c, c.Type)
	}
	if c.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", c, c.AttributedTo)
	}
	if len(c.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", c, c.Name)
	}
	if len(c.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", c, c.Summary)
	}
	if len(c.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", c, c.Content)
	}
	if c.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", c, c.URL)
	}
	if !c.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", c, c.Published)
	}
	if !c.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty StartTime, received %q", c, c.StartTime)
	}
	if !c.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty Updated, received %q", c, c.Updated)
	}
}

func TestDislike_UnmarshalJSON(t *testing.T) {
	d := Dislike{}

	dataEmpty := []byte("{}")
	d.UnmarshalJSON(dataEmpty)
	if d.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", d, d.ID)
	}
	if d.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", d, d.Type)
	}
	if d.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", d, d.AttributedTo)
	}
	if len(d.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", d, d.Name)
	}
	if len(d.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", d, d.Summary)
	}
	if len(d.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", d, d.Content)
	}
	if d.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", d, d.URL)
	}
	if !d.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", d, d.Published)
	}
	if !d.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty StartTime, received %q", d, d.StartTime)
	}
	if !d.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty Updated, received %q", d, d.Updated)
	}
}

func TestLike_UnmarshalJSON(t *testing.T) {
	l := Like{}

	dataEmpty := []byte("{}")
	l.UnmarshalJSON(dataEmpty)
	if l.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", l, l.ID)
	}
	if l.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", l, l.Type)
	}
	if l.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", l, l.AttributedTo)
	}
	if len(l.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", l, l.Name)
	}
	if len(l.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", l, l.Summary)
	}
	if len(l.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", l, l.Content)
	}
	if l.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", l, l.URL)
	}
	if !l.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", l, l.Published)
	}
	if !l.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty StartTime, received %q", l, l.StartTime)
	}
	if !l.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty Updated, received %q", l, l.Updated)
	}
}

func TestUpdate_UnmarshalJSON(t *testing.T) {
	u := Update{}

	dataEmpty := []byte("{}")
	u.UnmarshalJSON(dataEmpty)
	if u.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", u, u.ID)
	}
	if u.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", u, u.Type)
	}
	if u.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", u, u.AttributedTo)
	}
	if len(u.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", u, u.Name)
	}
	if len(u.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", u, u.Summary)
	}
	if len(u.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", u, u.Content)
	}
	if u.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", u, u.URL)
	}
	if !u.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", u, u.Published)
	}
	if !u.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty StartTime, received %q", u, u.StartTime)
	}
	if !u.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty Updated, received %q", u, u.Updated)
	}
}

func TestToActivity(t *testing.T) {
	var it Item
	act := ActivityNew(ID("test"), CreateType, nil)
	it = act

	a, err := ToActivity(it)
	if err != nil {
		t.Error(err)
	}
	if a != act {
		t.Errorf("Invalid activity returned by ToActivity #%v", a)
	}

	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToActivity(it)
	if err == nil {
		t.Errorf("Error returned when calling ToActivity with object should not be nil")
	}
	if o != nil {
		t.Errorf("Invalid return by ToActivity #%v, should have been nil", o)
	}
}

func TestFlattenActivityProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidEventRSVPActivityType(t *testing.T) {
	t.Skipf("TODO")
}
func TestValidGroupManagementActivityType(t *testing.T) {
	t.Skipf("TODO")
}

func TestActivity_Clean(t *testing.T) {
	t.Skipf("TODO")
}

func TestActivity_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestActivity_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestActivity_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestActivity_MarshalJSON(t *testing.T) {
	type fields struct {
		ID           ID
		Type         ActivityVocabularyType
		Name         NaturalLanguageValues
		Attachment   Item
		AttributedTo Item
		Audience     ItemCollection
		Content      NaturalLanguageValues
		Context      Item
		MediaType    MimeType
		EndTime      time.Time
		Generator    Item
		Icon         Item
		Image        Item
		InReplyTo    Item
		Location     Item
		Preview      Item
		Published    time.Time
		Replies      Item
		StartTime    time.Time
		Summary      NaturalLanguageValues
		Tag          ItemCollection
		Updated      time.Time
		URL          LinkOrIRI
		To           ItemCollection
		Bto          ItemCollection
		CC           ItemCollection
		BCC          ItemCollection
		Duration     time.Duration
		Likes        Item
		Shares       Item
		Source       Source
		Actor        Item
		Target       Item
		Result       Item
		Origin       Item
		Instrument   Item
		Object       Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Empty",
			fields:  fields{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "JustID",
			fields: fields{
				ID: ID("example.com"),
			},
			want:    []byte(`{"id":"example.com"}`),
			wantErr: false,
		},
		{
			name: "JustType",
			fields: fields{
				Type: ActivityVocabularyType("myType"),
			},
			want:    []byte(`{"type":"myType"}`),
			wantErr: false,
		},
		{
			name: "JustOneName",
			fields: fields{
				Name: NaturalLanguageValues{
					{Ref: NilLangRef, Value: "ana"},
				},
			},
			want:    []byte(`{"name":"ana"}`),
			wantErr: false,
		},
		{
			name: "MoreNames",
			fields: fields{
				Name: NaturalLanguageValues{
					{Ref: "en", Value: "anna"},
					{Ref: "fr", Value: "anne"},
				},
			},
			want:    []byte(`{"nameMap":{"en":"anna","fr":"anne"}}`),
			wantErr: false,
		},
		{
			name: "JustOneSummary",
			fields: fields{
				Summary: NaturalLanguageValues{
					{Ref: NilLangRef, Value: "test summary"},
				},
			},
			want:    []byte(`{"summary":"test summary"}`),
			wantErr: false,
		},
		{
			name: "MoreSummaryEntries",
			fields: fields{
				Summary: NaturalLanguageValues{
					{Ref: "en", Value: "test summary"},
					{Ref: "fr", Value: "teste summary"},
				},
			},
			want:    []byte(`{"summaryMap":{"en":"test summary","fr":"teste summary"}}`),
			wantErr: false,
		},
		{
			name: "JustOneContent",
			fields: fields{
				Content: NaturalLanguageValues{
					{Ref: NilLangRef, Value: "test content"},
				},
			},
			want:    []byte(`{"content":"test content"}`),
			wantErr: false,
		},
		{
			name: "MoreContentEntries",
			fields: fields{
				Content: NaturalLanguageValues{
					{Ref: "en", Value: "test content"},
					{Ref: "fr", Value: "teste content"},
				},
			},
			want:    []byte(`{"contentMap":{"en":"test content","fr":"teste content"}}`),
			wantErr: false,
		},
		{
			name: "MediaType",
			fields: fields{
				MediaType: MimeType("text/stupid"),
			},
			want:    []byte(`{"mediaType":"text/stupid"}`),
			wantErr: false,
		},
		{
			name: "Attachment",
			fields: fields{
				Attachment: &Object{
					ID:   "some example",
					Type: VideoType,
				},
			},
			want:    []byte(`{"attachment":{"id":"some example","type":"Video"}}`),
			wantErr: false,
		},
		{
			name: "AttributedTo",
			fields: fields{
				AttributedTo: &Actor{
					ID:   "http://example.com/ana",
					Type: PersonType,
				},
			},
			want:    []byte(`{"attributedTo":{"id":"http://example.com/ana","type":"Person"}}`),
			wantErr: false,
		},
		{
			name: "AttributedToDouble",
			fields: fields{
				AttributedTo: ItemCollection{
					&Actor{
						ID:   "http://example.com/ana",
						Type: PersonType,
					},
					&Actor{
						ID:   "http://example.com/GGG",
						Type: GroupType,
					},
				},
			},
			want:    []byte(`{"attributedTo":[{"id":"http://example.com/ana","type":"Person"},{"id":"http://example.com/GGG","type":"Group"}]}`),
			wantErr: false,
		},
		{
			name: "Source",
			fields: fields{
				Source: Source{
					MediaType: MimeType("text/plain"),
					Content:   NaturalLanguageValues{},
				},
			},
			want:    []byte(`{"source":{"mediaType":"text/plain"}}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Activity{
				ID:           tt.fields.ID,
				Type:         tt.fields.Type,
				Name:         tt.fields.Name,
				Attachment:   tt.fields.Attachment,
				AttributedTo: tt.fields.AttributedTo,
				Audience:     tt.fields.Audience,
				Content:      tt.fields.Content,
				Context:      tt.fields.Context,
				MediaType:    tt.fields.MediaType,
				EndTime:      tt.fields.EndTime,
				Generator:    tt.fields.Generator,
				Icon:         tt.fields.Icon,
				Image:        tt.fields.Image,
				InReplyTo:    tt.fields.InReplyTo,
				Location:     tt.fields.Location,
				Preview:      tt.fields.Preview,
				Published:    tt.fields.Published,
				Replies:      tt.fields.Replies,
				StartTime:    tt.fields.StartTime,
				Summary:      tt.fields.Summary,
				Tag:          tt.fields.Tag,
				Updated:      tt.fields.Updated,
				URL:          tt.fields.URL,
				To:           tt.fields.To,
				Bto:          tt.fields.Bto,
				CC:           tt.fields.CC,
				BCC:          tt.fields.BCC,
				Duration:     tt.fields.Duration,
				Likes:        tt.fields.Likes,
				Shares:       tt.fields.Shares,
				Source:       tt.fields.Source,
				Actor:        tt.fields.Actor,
				Target:       tt.fields.Target,
				Result:       tt.fields.Result,
				Origin:       tt.fields.Origin,
				Instrument:   tt.fields.Instrument,
				Object:       tt.fields.Object,
			}
			got, err := a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestIntransitiveActivity_MarshalJSON(t *testing.T) {
	type fields struct {
		ID           ID
		Type         ActivityVocabularyType
		Name         NaturalLanguageValues
		Attachment   Item
		AttributedTo Item
		Audience     ItemCollection
		Content      NaturalLanguageValues
		Context      Item
		MediaType    MimeType
		EndTime      time.Time
		Generator    Item
		Icon         Item
		Image        Item
		InReplyTo    Item
		Location     Item
		Preview      Item
		Published    time.Time
		Replies      Item
		StartTime    time.Time
		Summary      NaturalLanguageValues
		Tag          ItemCollection
		Updated      time.Time
		URL          LinkOrIRI
		To           ItemCollection
		Bto          ItemCollection
		CC           ItemCollection
		BCC          ItemCollection
		Duration     time.Duration
		Likes        Item
		Shares       Item
		Source       Source
		Actor        CanReceiveActivities
		Target       Item
		Result       Item
		Origin       Item
		Instrument   Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Empty",
			fields:  fields{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "JustID",
			fields: fields{
				ID: ID("example.com"),
			},
			want:    []byte(`{"id":"example.com"}`),
			wantErr: false,
		},
		{
			name: "JustType",
			fields: fields{
				Type: ActivityVocabularyType("myType"),
			},
			want:    []byte(`{"type":"myType"}`),
			wantErr: false,
		},
		{
			name: "JustOneName",
			fields: fields{
				Name: NaturalLanguageValues{
					{Ref: NilLangRef, Value: "ana"},
				},
			},
			want:    []byte(`{"name":"ana"}`),
			wantErr: false,
		},
		{
			name: "MoreNames",
			fields: fields{
				Name: NaturalLanguageValues{
					{Ref: "en", Value: "anna"},
					{Ref: "fr", Value: "anne"},
				},
			},
			want:    []byte(`{"nameMap":{"en":"anna","fr":"anne"}}`),
			wantErr: false,
		},
		{
			name: "JustOneSummary",
			fields: fields{
				Summary: NaturalLanguageValues{
					{Ref: NilLangRef, Value: "test summary"},
				},
			},
			want:    []byte(`{"summary":"test summary"}`),
			wantErr: false,
		},
		{
			name: "MoreSummaryEntries",
			fields: fields{
				Summary: NaturalLanguageValues{
					{Ref: "en", Value: "test summary"},
					{Ref: "fr", Value: "teste summary"},
				},
			},
			want:    []byte(`{"summaryMap":{"en":"test summary","fr":"teste summary"}}`),
			wantErr: false,
		},
		{
			name: "JustOneContent",
			fields: fields{
				Content: NaturalLanguageValues{
					{Ref: NilLangRef, Value: "test content"},
				},
			},
			want:    []byte(`{"content":"test content"}`),
			wantErr: false,
		},
		{
			name: "MoreContentEntries",
			fields: fields{
				Content: NaturalLanguageValues{
					{Ref: "en", Value: "test content"},
					{Ref: "fr", Value: "teste content"},
				},
			},
			want:    []byte(`{"contentMap":{"en":"test content","fr":"teste content"}}`),
			wantErr: false,
		},
		{
			name: "MediaType",
			fields: fields{
				MediaType: MimeType("text/stupid"),
			},
			want:    []byte(`{"mediaType":"text/stupid"}`),
			wantErr: false,
		},
		{
			name: "Attachment",
			fields: fields{
				Attachment: &Object{
					ID:   "some example",
					Type: VideoType,
				},
			},
			want:    []byte(`{"attachment":{"id":"some example","type":"Video"}}`),
			wantErr: false,
		},
		{
			name: "AttributedTo",
			fields: fields{
				AttributedTo: &Actor{
					ID:   "http://example.com/ana",
					Type: PersonType,
				},
			},
			want:    []byte(`{"attributedTo":{"id":"http://example.com/ana","type":"Person"}}`),
			wantErr: false,
		},
		{
			name: "AttributedToDouble",
			fields: fields{
				AttributedTo: ItemCollection{
					&Actor{
						ID:   "http://example.com/ana",
						Type: PersonType,
					},
					&Actor{
						ID:   "http://example.com/GGG",
						Type: GroupType,
					},
				},
			},
			want:    []byte(`{"attributedTo":[{"id":"http://example.com/ana","type":"Person"},{"id":"http://example.com/GGG","type":"Group"}]}`),
			wantErr: false,
		},
		{
			name: "Source",
			fields: fields{
				Source: Source{
					MediaType: MimeType("text/plain"),
					Content:   NaturalLanguageValues{},
				},
			},
			want:    []byte(`{"source":{"mediaType":"text/plain"}}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := IntransitiveActivity{
				ID:           tt.fields.ID,
				Type:         tt.fields.Type,
				Name:         tt.fields.Name,
				Attachment:   tt.fields.Attachment,
				AttributedTo: tt.fields.AttributedTo,
				Audience:     tt.fields.Audience,
				Content:      tt.fields.Content,
				Context:      tt.fields.Context,
				MediaType:    tt.fields.MediaType,
				EndTime:      tt.fields.EndTime,
				Generator:    tt.fields.Generator,
				Icon:         tt.fields.Icon,
				Image:        tt.fields.Image,
				InReplyTo:    tt.fields.InReplyTo,
				Location:     tt.fields.Location,
				Preview:      tt.fields.Preview,
				Published:    tt.fields.Published,
				Replies:      tt.fields.Replies,
				StartTime:    tt.fields.StartTime,
				Summary:      tt.fields.Summary,
				Tag:          tt.fields.Tag,
				Updated:      tt.fields.Updated,
				URL:          tt.fields.URL,
				To:           tt.fields.To,
				Bto:          tt.fields.Bto,
				CC:           tt.fields.CC,
				BCC:          tt.fields.BCC,
				Duration:     tt.fields.Duration,
				Likes:        tt.fields.Likes,
				Shares:       tt.fields.Shares,
				Source:       tt.fields.Source,
				Actor:        tt.fields.Actor,
				Target:       tt.fields.Target,
				Result:       tt.fields.Result,
				Origin:       tt.fields.Origin,
				Instrument:   tt.fields.Instrument,
			}
			got, err := i.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}
