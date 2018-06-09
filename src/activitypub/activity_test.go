package activitypub

import "testing"

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
	var testType ActivityVocabularyType = "Accept"

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

func TestValidActivityType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidActivityType(ActivityType) {
		t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
	}
	for _, inValidType := range validObjectTypes {
		if ValidActivityType(inValidType) {
			t.Errorf("APObject Type '%v' should be invalid", inValidType)
		}
	}
	if ValidActivityType(invalidType) {
		t.Errorf("Activity Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validActivityTypes {
		if !ValidActivityType(validType) {
			t.Errorf("Activity Type '%v' should be valid", validType)
		}
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

func TestActivityRecipientsDeduplication(t *testing.T) {
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

	a.RecipientsDeduplication()
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

	b.RecipientsDeduplication()
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

func TestBlockRecipientsDeduplication(t *testing.T) {
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

	a.RecipientsDeduplication()
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

	b.RecipientsDeduplication()
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
}

func TestIntransitiveActivityRecipientsDeduplication(t *testing.T) {
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

	a.RecipientsDeduplication()
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

	b.RecipientsDeduplication()
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
func TestActivity_GetLink(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if a.GetLink().ID != "" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetLink(), a.GetLink())
	}
}
func TestActivity_GetObject(t *testing.T) {
	a := ActivityNew("test", ActivityType, Person{})

	if a.GetObject().ID != "test" || a.GetObject().Type != ActivityType {
		t.Errorf("%T should not return an empty %T object. Received %#v", a, a.GetObject(), a.GetObject())
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

	if i.GetLink().ID != "" {
		t.Errorf("%T should return an empty %T object. Received %#v", i, i.GetLink(), i.GetLink())
	}
}
func TestIntransitiveActivity_GetObject(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if i.GetObject().ID != "test" || i.GetObject().Type != QuestionType {
		t.Errorf("%T should not return an empty %T object. Received %#v", i, i.GetObject(), i.GetObject())
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
func TestActivity_RecipientsDeduplication(t *testing.T) {
	t.Skip("See TestDeduplication")
}
func TestIntransitiveActivity_RecipientsDeduplication(t *testing.T) {
	t.Skip("See TestDeduplication")
}
func TestBlock_RecipientsDeduplication(t *testing.T) {
	t.Skip("See TestDeduplication")
}
