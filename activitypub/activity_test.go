package activitypub

import "testing"

func TestActivityNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType string = "Accept"

	a := ActivityNew(testValue, testType)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != testType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, testType)
	}

	g := ActivityNew(testValue, "")

	if g.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", g.Id, testValue)
	}
	if g.Type != ActivityType {
		t.Errorf("Activity Type '%v' different than expected '%v'", g.Type, ActivityType)
	}
}

func TestIntransitiveActivityNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType string = "Accept"

	a := IntransitiveActivityNew(testValue, testType)

	if a.Id != testValue {
		t.Errorf("IntransitiveActivity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != testType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", a.Type, testType)
	}

	g := IntransitiveActivityNew(testValue, "")

	if g.Id != testValue {
		t.Errorf("IntransitiveActivity Id '%v' different than expected '%v'", g.Id, testValue)
	}
	if g.Type != IntransitiveActivityType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", g.Type, IntransitiveActivityType)
	}
}

func TestValidActivityType(t *testing.T) {
	var invalidType string = "RandomType"

	if ValidActivityType(ActivityType) {
		t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
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
	var testValue = ObjectId("test")

	a := AcceptNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != AcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AcceptType)
	}
}

func TestAddNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := AddNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != AddType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AddType)
	}
}

func TestAnnounceNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := AnnounceNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != AnnounceType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, AnnounceType)
	}
}

func TestArriveNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := ArriveNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != ArriveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ArriveType)
	}
}

func TestBlockNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := BlockNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != BlockType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, BlockType)
	}
}

func TestCreateNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := CreateNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, CreateType)
	}
}

func TestDeleteNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := DeleteNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != DeleteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, DeleteType)
	}
}

func TestDislikeNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := DislikeNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != DislikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, DislikeType)
	}
}

func TestFlagNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := FlagNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != FlagType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, FlagType)
	}
}

func TestFollowNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := FollowNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != FollowType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, FollowType)
	}
}

func TestIgnoreNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := IgnoreNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != IgnoreType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, IgnoreType)
	}
}

func TestInviteNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := InviteNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != InviteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, InviteType)
	}
}

func TestJoinNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := JoinNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != JoinType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, JoinType)
	}
}

func TestLeaveNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := LeaveNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != LeaveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, LeaveType)
	}
}

func TestLikeNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := LikeNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, LikeType)
	}
}

func TestListenNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := ListenNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != ListenType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ListenType)
	}
}

func TestMoveNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := MoveNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != MoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, MoveType)
	}
}

func TestOfferNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := OfferNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != OfferType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, OfferType)
	}
}

func TestQuestionNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := QuestionNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != QuestionType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, QuestionType)
	}
}

func TestRejectNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := RejectNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != RejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, RejectType)
	}
}

func TestReadNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := ReadNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != ReadType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ReadType)
	}
}

func TestRemoveNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := RemoveNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != RemoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, RemoveType)
	}
}

func TestTentativeRejectNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := TentativeRejectNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != TentativeRejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TentativeRejectType)
	}
}

func TestTentativeAcceptNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := TentativeAcceptNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != TentativeAcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TentativeAcceptType)
	}
}

func TestTravelNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := TravelNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != TravelType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TravelType)
	}
}

func TestUndoNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := UndoNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != UndoType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, UndoType)
	}
}

func TestUpdateNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := UpdateNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != UpdateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, UpdateType)
	}
}

func TestViewNew(t *testing.T) {
	var testValue = ObjectId("test")

	a := ViewNew(testValue)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != ViewType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ViewType)
	}
}
