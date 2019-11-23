package activitystreams

import (
	"fmt"
	"testing"
)

func TestActivityNew(t *testing.T) {
	var testValue = ObjectID("test")
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

func TestIntransitiveActivityNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType ActivityVocabularyType = "Arrive"

	a := IntransitiveActivityNew(testValue, testType)

	if a.ID != testValue {
		t.Errorf("IntransitiveActivity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != testType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", a.Type, testType)
	}

	g := IntransitiveActivityNew(testValue, "")

	if g.ID != testValue {
		t.Errorf("IntransitiveActivity Id '%v' different than expected '%v'", g.ID, testValue)
	}
	if g.Type != IntransitiveActivityType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", g.Type, IntransitiveActivityType)
	}
}

func TestAcceptNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := AcceptNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != AcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AcceptType)
	}
}

func TestAddNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := AddNew(testValue, nil, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != AddType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AddType)
	}
}

func TestAnnounceNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := AnnounceNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != AnnounceType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AnnounceType)
	}
}

func TestArriveNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := ArriveNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ArriveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ArriveType)
	}
}

func TestBlockNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := BlockNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != BlockType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, BlockType)
	}
}

func TestCreateNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := CreateNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, CreateType)
	}
}

func TestDeleteNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := DeleteNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != DeleteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, DeleteType)
	}
}

func TestDislikeNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := DislikeNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != DislikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, DislikeType)
	}
}

func TestFlagNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := FlagNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != FlagType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, FlagType)
	}
}

func TestFollowNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := FollowNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != FollowType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, FollowType)
	}
}

func TestIgnoreNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := IgnoreNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != IgnoreType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, IgnoreType)
	}
}

func TestInviteNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := InviteNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != InviteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, InviteType)
	}
}

func TestJoinNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := JoinNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != JoinType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, JoinType)
	}
}

func TestLeaveNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := LeaveNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != LeaveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, LeaveType)
	}
}

func TestLikeNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := LikeNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, LikeType)
	}
}

func TestListenNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := ListenNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ListenType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ListenType)
	}
}

func TestMoveNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := MoveNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != MoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, MoveType)
	}
}

func TestOfferNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := OfferNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != OfferType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, OfferType)
	}
}

func TestQuestionNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := QuestionNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != QuestionType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, QuestionType)
	}
}

func TestRejectNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := RejectNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != RejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, RejectType)
	}
}

func TestReadNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := ReadNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ReadType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ReadType)
	}
}

func TestRemoveNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := RemoveNew(testValue, nil, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != RemoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, RemoveType)
	}
}

func TestTentativeRejectNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := TentativeRejectNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != TentativeRejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TentativeRejectType)
	}
}

func TestTentativeAcceptNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := TentativeAcceptNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != TentativeAcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TentativeAcceptType)
	}
}

func TestTravelNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := TravelNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != TravelType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TravelType)
	}
}

func TestUndoNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := UndoNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != UndoType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, UndoType)
	}
}

func TestUpdateNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := UpdateNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != UpdateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, UpdateType)
	}
}

func TestViewNew(t *testing.T) {
	var testValue = ObjectID("test")

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
		t.Errorf("%T.To should have exactly 3(four) elements, not %d", a, len(a.To))
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
	var err error
	recIds := make([]ObjectID, 0)
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

func TestIntransitiveActivityRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	a := IntransitiveActivityNew("test", "t")

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
	var err error
	recIds := make([]ObjectID, 0)
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
	recIds := make([]ObjectID, 0)
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
	recIds := make([]ObjectID, 0)
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
	recIds := make([]ObjectID, 0)
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
	recIds := make([]ObjectID, 0)
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

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}
func TestActivity_GetIDGetType(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if *a.GetID() != "test" || a.GetType() != ActivityType {
		t.Errorf("%T should not return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
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
func TestIntransitiveActivity_GetLink(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if *i.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", i, i, i)
	}
}
func TestIntransitiveActivity_GetObject(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if *i.GetID() != "test" || i.GetType() != QuestionType {
		t.Errorf("%T should not return an empty %T object. Received %#v", i, i, i)
	}
}
func TestIntransitiveActivity_IsLink(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if i.IsLink() {
		t.Errorf("%T should not respond true to IsLink", i)
	}
}
func TestIntransitiveActivity_IsObject(t *testing.T) {
	i := IntransitiveActivityNew("test", ActivityType)

	if !i.IsObject() {
		t.Errorf("%T should respond true to IsObject", i)
	}
}

func checkDedup(list ItemCollection, recIds *[]ObjectID) error {
	for _, rec := range list {
		for _, id := range *recIds {
			if *rec.GetID() == id {
				return fmt.Errorf("%T[%s] already stored in recipients list, Deduplication faild", rec, id)
			}
		}
		*recIds = append(*recIds, *rec.GetID())
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
	recIds := make([]ObjectID, 0)
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
func TestIntransitiveActivity_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	c := IntransitiveActivityNew("act", IntransitiveActivityType)
	c.To.Append(to)
	c.CC.Append(cc)
	c.BCC.Append(cc)

	c.Recipients()

	var err error
	recIds := make([]ObjectID, 0)
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
	recIds := make([]ObjectID, 0)
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

func TestRead_GetID(t *testing.T) {
	a := ReadNew("test", Person{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestAccept_GetID(t *testing.T) {
	a := AcceptNew("test", Person{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestAdd_GetID(t *testing.T) {
	a := AddNew("test", Person{}, Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), a)
	}
}

func TestAnnounce_GetID(t *testing.T) {
	a := AnnounceNew("test", Person{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestArrive_GetID(t *testing.T) {
	a := ArriveNew("test")

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestBlock_GetID(t *testing.T) {
	a := BlockNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestCreate_GetID(t *testing.T) {
	a := CreateNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestDelete_GetID(t *testing.T) {
	a := DeleteNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestDislike_GetID(t *testing.T) {
	a := DislikeNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestFlag_GetID(t *testing.T) {
	a := FlagNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestFollow_GetID(t *testing.T) {
	a := FollowNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestIgnore_GetID(t *testing.T) {
	a := IgnoreNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestInvite_GetID(t *testing.T) {
	a := InviteNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestJoin_GetID(t *testing.T) {
	a := JoinNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestLeave_GetID(t *testing.T) {
	a := LeaveNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestLike_GetID(t *testing.T) {
	a := LikeNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestListen_GetID(t *testing.T) {
	a := ListenNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestMove_GetID(t *testing.T) {
	a := MoveNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestOffer_GetID(t *testing.T) {
	a := OfferNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestQuestion_GetID(t *testing.T) {
	a := QuestionNew("test")

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestReject_GetID(t *testing.T) {
	a := RejectNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestRemove_GetID(t *testing.T) {
	a := RemoveNew("test", Object{}, Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestTravel_GetID(t *testing.T) {
	a := TravelNew("test")

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestUpdate_GetID(t *testing.T) {
	a := UpdateNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestView_GetID(t *testing.T) {
	a := ViewNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestIntransitiveActivity_GetID(t *testing.T) {
	a := IntransitiveActivityNew("test", IntransitiveActivityType)

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestTentativeAccept_GetID(t *testing.T) {
	a := TentativeAcceptNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestTentativeReject_GetID(t *testing.T) {
	a := TentativeRejectNew("test", Object{})

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestAccept_IsObject(t *testing.T) {
	a := AcceptNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestAdd_IsObject(t *testing.T) {
	a := AddNew("test", Object{}, Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestAnnounce_IsObject(t *testing.T) {
	a := AnnounceNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestArrive_IsObject(t *testing.T) {
	a := ArriveNew("test")

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestBlock_IsObject(t *testing.T) {
	a := BlockNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestCreate_IsObject(t *testing.T) {
	a := CreateNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestDelete_IsObject(t *testing.T) {
	a := DeleteNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestDislike_IsObject(t *testing.T) {
	a := DislikeNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestFlag_IsObject(t *testing.T) {
	a := FlagNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestFollow_IsObject(t *testing.T) {
	a := FollowNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestIgnore_IsObject(t *testing.T) {
	a := IgnoreNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestInvite_IsObject(t *testing.T) {
	a := InviteNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestJoin_IsObject(t *testing.T) {
	a := JoinNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestLeave_IsObject(t *testing.T) {
	a := LeaveNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestLike_IsObject(t *testing.T) {
	a := LikeNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestListen_IsObject(t *testing.T) {
	a := ListenNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestMove_IsObject(t *testing.T) {
	a := MoveNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestOffer_IsObject(t *testing.T) {
	a := OfferNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestQuestion_IsObject(t *testing.T) {
	a := QuestionNew("test")

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestRead_IsObject(t *testing.T) {
	a := ReadNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestReject_IsObject(t *testing.T) {
	a := RejectNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestRemove_IsObject(t *testing.T) {
	a := RemoveNew("test", Object{}, Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}
func TestTravel_IsObject(t *testing.T) {
	a := TravelNew("test")

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestUpdate_IsObject(t *testing.T) {
	a := UpdateNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestView_IsObject(t *testing.T) {
	a := ViewNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestTentativeAccept_IsObject(t *testing.T) {
	a := TentativeAcceptNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestTentativeReject_IsObject(t *testing.T) {
	a := TentativeRejectNew("test", Object{})

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestAccept_IsLink(t *testing.T) {
	a := AcceptNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestAdd_IsLink(t *testing.T) {
	a := AddNew("test", Object{}, Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestAnnounce_IsLink(t *testing.T) {
	a := AnnounceNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestArrive_IsLink(t *testing.T) {
	a := ArriveNew("test")

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestBlock_IsLink(t *testing.T) {
	a := BlockNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestCreate_IsLink(t *testing.T) {
	a := CreateNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestDelete_IsLink(t *testing.T) {
	a := DeleteNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestDislike_IsLink(t *testing.T) {
	a := DislikeNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestFlag_IsLink(t *testing.T) {
	a := FlagNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestFollow_IsLink(t *testing.T) {
	a := FollowNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestIgnore_IsLink(t *testing.T) {
	a := IgnoreNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestInvite_IsLink(t *testing.T) {
	a := InviteNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestJoin_IsLink(t *testing.T) {
	a := JoinNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestLeave_IsLink(t *testing.T) {
	a := LeaveNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestLike_IsLink(t *testing.T) {
	a := LikeNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestListen_IsLink(t *testing.T) {
	a := ListenNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestMove_IsLink(t *testing.T) {
	a := MoveNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestOffer_IsLink(t *testing.T) {
	a := OfferNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestQuestion_IsLink(t *testing.T) {
	a := QuestionNew("test")

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestRead_IsLink(t *testing.T) {
	a := ReadNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestReject_IsLink(t *testing.T) {
	a := RejectNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestRemove_IsLink(t *testing.T) {
	a := RemoveNew("test", Object{}, Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}
func TestTravel_IsLink(t *testing.T) {
	a := TravelNew("test")

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestUpdate_IsLink(t *testing.T) {
	a := UpdateNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestView_IsLink(t *testing.T) {
	a := ViewNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestTentativeAccept_IsLink(t *testing.T) {
	a := TentativeAcceptNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestTentativeReject_IsLink(t *testing.T) {
	a := TentativeRejectNew("test", Object{})

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestAccept_GetLink(t *testing.T) {
	a := AcceptNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestAdd_GetLink(t *testing.T) {
	a := AddNew("test", Object{}, Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestAnnounce_GetLink(t *testing.T) {
	a := AnnounceNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestArrive_GetLink(t *testing.T) {
	a := ArriveNew("test")

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestBlock_GetLink(t *testing.T) {
	a := BlockNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestCreate_GetLink(t *testing.T) {
	a := CreateNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestDelete_GetLink(t *testing.T) {
	a := DeleteNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestDislike_GetLink(t *testing.T) {
	a := DislikeNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestFlag_GetLink(t *testing.T) {
	a := FlagNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestFollow_GetLink(t *testing.T) {
	a := FollowNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestIgnore_GetLink(t *testing.T) {
	a := IgnoreNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestInvite_GetLink(t *testing.T) {
	a := InviteNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestJoin_GetLink(t *testing.T) {
	a := JoinNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestLeave_GetLink(t *testing.T) {
	a := LeaveNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestLike_GetLink(t *testing.T) {
	a := LikeNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestListen_GetLink(t *testing.T) {
	a := ListenNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestMove_GetLink(t *testing.T) {
	a := MoveNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestOffer_GetLink(t *testing.T) {
	a := OfferNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestQuestion_GetLink(t *testing.T) {
	a := QuestionNew("test")

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestRead_GetLink(t *testing.T) {
	a := ReadNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestReject_GetLink(t *testing.T) {
	a := RejectNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestRemove_GetLink(t *testing.T) {
	a := RemoveNew("test", Object{}, Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}
func TestTravel_GetLink(t *testing.T) {
	a := TravelNew("test")

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestUpdate_GetLink(t *testing.T) {
	a := UpdateNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\"for %T, received %q", a, a.GetLink())
	}
}

func TestView_GetLink(t *testing.T) {
	a := ViewNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestTentativeAccept_GetLink(t *testing.T) {
	a := TentativeAcceptNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestTentativeReject_GetLink(t *testing.T) {
	a := TentativeRejectNew("test", Object{})

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestAccept_GetType(t *testing.T) {
	a := AcceptNew("test", Object{})

	if a.GetType() != AcceptType {
		t.Errorf("GetType should return %q for %T, received %q", AcceptType, a, a.GetType())
	}
}

func TestAdd_GetType(t *testing.T) {
	a := AddNew("test", Object{}, Object{})

	if a.GetType() != AddType {
		t.Errorf("GetType should return %q for %T, received %q", AddType, a, a.GetType())
	}
}

func TestAnnounce_GetType(t *testing.T) {
	a := AnnounceNew("test", Object{})

	if a.GetType() != AnnounceType {
		t.Errorf("GetType should return %q for %T, received %q", AnnounceType, a, a.GetType())
	}
}

func TestArrive_GetType(t *testing.T) {
	a := ArriveNew("test")

	if a.GetType() != ArriveType {
		t.Errorf("GetType should return %q for %T, received %q", ArriveType, a, a.GetType())
	}
}

func TestBlock_GetType(t *testing.T) {
	a := BlockNew("test", Object{})

	if a.GetType() != BlockType {
		t.Errorf("GetType should return %q for %T, received %q", BlockType, a, a.GetType())
	}
}
func TestCreate_GetType(t *testing.T) {
	a := CreateNew("test", Object{})

	if a.GetType() != CreateType {
		t.Errorf("GetType should return %q for %T, received %q", CreateType, a, a.GetType())
	}
}

func TestDelete_GetType(t *testing.T) {
	a := DeleteNew("test", Object{})

	if a.GetType() != DeleteType {
		t.Errorf("GetType should return %q for %T, received %q", DeleteType, a, a.GetType())
	}
}
func TestDislike_GetType(t *testing.T) {
	a := DislikeNew("test", Object{})

	if a.GetType() != DislikeType {
		t.Errorf("GetType should return %q for %T, received %q", DislikeType, a, a.GetType())
	}
}

func TestFlag_GetType(t *testing.T) {
	a := FlagNew("test", Object{})

	if a.GetType() != FlagType {
		t.Errorf("GetType should return %q for %T, received %q", FlagType, a, a.GetType())
	}
}
func TestFollow_GetType(t *testing.T) {
	a := FollowNew("test", Object{})

	if a.GetType() != FollowType {
		t.Errorf("GetType should return %q for %T, received %q", FollowType, a, a.GetType())
	}
}
func TestIgnore_GetType(t *testing.T) {
	a := IgnoreNew("test", Object{})

	if a.GetType() != IgnoreType {
		t.Errorf("GetType should return %q for %T, received %q", IgnoreType, a, a.GetType())
	}
}
func TestInvite_GetType(t *testing.T) {
	a := InviteNew("test", Object{})

	if a.GetType() != InviteType {
		t.Errorf("GetType should return %q for %T, received %q", InviteType, a, a.GetType())
	}
}
func TestJoin_GetType(t *testing.T) {
	a := JoinNew("test", Object{})

	if a.GetType() != JoinType {
		t.Errorf("GetType should return %q for %T, received %q", JoinType, a, a.GetType())
	}
}
func TestLeave_GetType(t *testing.T) {
	a := LeaveNew("test", Object{})

	if a.GetType() != LeaveType {
		t.Errorf("GetType should return %q for %T, received %q", LeaveType, a, a.GetType())
	}
}
func TestLike_GetType(t *testing.T) {
	a := LikeNew("test", Object{})

	if a.GetType() != LikeType {
		t.Errorf("GetType should return %q for %T, received %q", LikeType, a, a.GetType())
	}
}
func TestListen_GetType(t *testing.T) {
	a := ListenNew("test", Object{})

	if a.GetType() != ListenType {
		t.Errorf("GetType should return %q for %T, received %q", ListenType, a, a.GetType())
	}
}
func TestMove_GetType(t *testing.T) {
	a := MoveNew("test", Object{})

	if a.GetType() != MoveType {
		t.Errorf("GetType should return %q for %T, received %q", MoveType, a, a.GetType())
	}
}
func TestOffer_GetType(t *testing.T) {
	a := OfferNew("test", Object{})

	if a.GetType() != OfferType {
		t.Errorf("GetType should return %q for %T, received %q", OfferType, a, a.GetType())
	}
}
func TestQuestion_GetType(t *testing.T) {
	a := QuestionNew("test")

	if a.GetType() != QuestionType {
		t.Errorf("GetType should return %q for %T, received %q", QuestionType, a, a.GetType())
	}
}
func TestRead_GetType(t *testing.T) {
	a := ReadNew("test", Object{})

	if a.GetType() != ReadType {
		t.Errorf("GetType should return %q for %T, received %q", ReadType, a, a.GetType())
	}
}
func TestReject_GetType(t *testing.T) {
	a := RejectNew("test", Object{})

	if a.GetType() != RejectType {
		t.Errorf("GetType should return %q for %T, received %q", RejectType, a, a.GetType())
	}
}
func TestRemove_GetType(t *testing.T) {
	a := RemoveNew("test", Object{}, Object{})

	if a.GetType() != RemoveType {
		t.Errorf("GetType should return %q for %T, received %q", RemoveType, a, a.GetType())
	}
}
func TestTravel_GetType(t *testing.T) {
	a := TravelNew("test")

	if a.GetType() != TravelType {
		t.Errorf("GetType should return %q for %T, received %q", TravelType, a, a.GetType())
	}
}

func TestUpdate_GetType(t *testing.T) {
	a := UpdateNew("test", Object{})

	if a.GetType() != UpdateType {
		t.Errorf("GetType should return %q for %T, received %q", UpdateType, a, a.GetType())
	}
}

func TestView_GetType(t *testing.T) {
	a := ViewNew("test", Object{})

	if a.GetType() != ViewType {
		t.Errorf("GetType should return %q for %T, received %q", ViewType, a, a.GetType())
	}
}

func TestTentativeAccept_GetType(t *testing.T) {
	a := TentativeAcceptNew("test", Object{})

	if a.GetType() != TentativeAcceptType {
		t.Errorf("GetType should return %q for %T, received %q", TentativeAcceptType, a, a.GetType())
	}
}

func TestTentativeReject_GetType(t *testing.T) {
	a := TentativeRejectNew("test", Object{})

	if a.GetType() != TentativeRejectType {
		t.Errorf("GetType should return %q for %T, received %q", TentativeRejectType, a, a.GetType())
	}
}

func TestActivity_GetLink(t *testing.T) {
	a := ActivityNew("test", ActivityType, Object{})
	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetType())
	}
}

func TestActivity_GetType(t *testing.T) {
	{
		a := ActivityNew("test", ActivityType, Object{})
		if a.GetType() != ActivityType {
			t.Errorf("GetType should return %q for %T, received %q", ActivityType, a, a.GetType())
		}
	}
	{
		a := ActivityNew("test", AcceptType, Object{})
		if a.GetType() != AcceptType {
			t.Errorf("GetType should return %q for %T, received %q", AcceptType, a, a.GetType())
		}
	}
	{
		a := ActivityNew("test", BlockType, Object{})
		if a.GetType() != BlockType {
			t.Errorf("GetType should return %q for %T, received %q", BlockType, a, a.GetType())
		}
	}
}

func TestIntransitiveActivity_GetType(t *testing.T) {
	{
		a := IntransitiveActivityNew("test", IntransitiveActivityType)
		if a.GetType() != IntransitiveActivityType {
			t.Errorf("GetType should return %q for %T, received %q", IntransitiveActivityType, a, a.GetType())
		}
	}
	{
		a := IntransitiveActivityNew("test", ArriveType)
		if a.GetType() != ArriveType {
			t.Errorf("GetType should return %q for %T, received %q", ArriveType, a, a.GetType())
		}
	}
	{
		a := IntransitiveActivityNew("test", QuestionType)
		if a.GetType() != QuestionType {
			t.Errorf("GetType should return %q for %T, received %q", QuestionType, a, a.GetType())
		}
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
	act := ActivityNew(ObjectID("test"), CreateType, nil)
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

func TestToIntransitiveActivity(t *testing.T) {
	var it Item
	act := IntransitiveActivityNew("test", TravelType)
	it = act

	a, err := ToIntransitiveActivity(it)
	if err != nil {
		t.Error(err)
	}
	if a != act {
		t.Errorf("Invalid activity returned by ToActivity #%v", a)
	}

	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToIntransitiveActivity(it)
	if err == nil {
		t.Errorf("Error returned when calling ToActivity with object should not be nil")
	}
	if o != nil {
		t.Errorf("Invalid return by ToActivity #%v, should have been nil", o)
	}
}

func TestToQuestion(t *testing.T) {
	var it Item
	act := QuestionNew("test")
	it = act

	a, err := ToQuestion(it)
	if err != nil {
		t.Error(err)
	}
	if a != act {
		t.Errorf("Invalid activity returned by ToActivity #%v", a)
	}

	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToQuestion(it)
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

func TestFlattenIntransitiveActivityProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidEventRSVPActivityType(t *testing.T) {
	t.Skipf("TODO")
}
func TestValidGroupManagementActivityType(t *testing.T) {
	t.Skipf("TODO")
}
