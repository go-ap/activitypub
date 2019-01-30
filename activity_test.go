package activitypub

import (
	as "github.com/go-ap/activitystreams"
	"testing"
)

func TestActivityNew(t *testing.T) {
	var testValue = as.ObjectID("test")
	var testType as.ActivityVocabularyType = "Accept"

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
	if g.Type != as.ActivityType {
		t.Errorf("Activity Type '%v' different than expected '%v'", g.Type, as.ActivityType)
	}
}

func TestIntransitiveActivityNew(t *testing.T) {
	var testValue = as.ObjectID("test")
	var testType as.ActivityVocabularyType = "Arrive"

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
	if g.Type != as.IntransitiveActivityType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", g.Type, as.IntransitiveActivityType)
	}
}

func TestAcceptNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := AcceptNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.AcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.AcceptType)
	}
}

func TestAddNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := AddNew(testValue, nil, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.AddType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.AddType)
	}
}

func TestAnnounceNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := AnnounceNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.AnnounceType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.AnnounceType)
	}
}

func TestArriveNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := ArriveNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.ArriveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.ArriveType)
	}
}

func TestBlockNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := BlockNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.BlockType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.BlockType)
	}
}

func TestCreateNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := CreateNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.CreateType)
	}
}

func TestDeleteNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := DeleteNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.DeleteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.DeleteType)
	}
}

func TestDislikeNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := DislikeNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.DislikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.DislikeType)
	}
}

func TestFlagNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := FlagNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.FlagType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.FlagType)
	}
}

func TestFollowNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := FollowNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.FollowType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.FollowType)
	}
}

func TestIgnoreNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := IgnoreNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.IgnoreType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.IgnoreType)
	}
}

func TestInviteNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := InviteNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.InviteType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.InviteType)
	}
}

func TestJoinNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := JoinNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.JoinType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.JoinType)
	}
}

func TestLeaveNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := LeaveNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.LeaveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.LeaveType)
	}
}

func TestLikeNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := LikeNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.LikeType)
	}
}

func TestListenNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := ListenNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.ListenType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.ListenType)
	}
}

func TestMoveNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := MoveNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.MoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.MoveType)
	}
}

func TestOfferNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := OfferNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.OfferType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.OfferType)
	}
}

func TestQuestionNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := QuestionNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.QuestionType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.QuestionType)
	}
}

func TestRejectNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := RejectNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.RejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.RejectType)
	}
}

func TestReadNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := ReadNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.ReadType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.ReadType)
	}
}

func TestRemoveNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := RemoveNew(testValue, nil, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.RemoveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.RemoveType)
	}
}

func TestTentativeRejectNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := TentativeRejectNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.TentativeRejectType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.TentativeRejectType)
	}
}

func TestTentativeAcceptNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := TentativeAcceptNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.TentativeAcceptType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.TentativeAcceptType)
	}
}

func TestTravelNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := TravelNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.TravelType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.TravelType)
	}
}

func TestUndoNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := UndoNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.UndoType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.UndoType)
	}
}

func TestUpdateNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := UpdateNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.UpdateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.UpdateType)
	}
}

func TestViewNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	a := ViewNew(testValue, nil)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != as.ViewType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, as.ViewType)
	}
}

func TestActivity_UnmarshalJSON(t *testing.T) {
	a := Activity{}

	dataEmpty := []byte("{}")
	a.UnmarshalJSON(dataEmpty)
	if a.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", a, a.ID)
	}
	if a.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", a, a.Type)
	}
	if a.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", a, a.AttributedTo)
	}
	if len(a.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", a, a.Name)
	}
	if len(a.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", a, a.Summary)
	}
	if len(a.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", a, a.Content)
	}
	if a.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", a, a.URL)
	}
	if !a.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", a, a.Published)
	}
	if !a.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", a, a.StartTime)
	}
	if !a.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", a, a.Updated)
	}
}

func TestCreate_UnmarshalJSON(t *testing.T) {
	c := Create{}

	dataEmpty := []byte("{}")
	c.UnmarshalJSON(dataEmpty)
	if c.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", c, c.ID)
	}
	if c.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", c, c.Type)
	}
	if c.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", c, c.AttributedTo)
	}
	if len(c.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", c, c.Name)
	}
	if len(c.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", c, c.Summary)
	}
	if len(c.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", c, c.Content)
	}
	if c.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", c, c.URL)
	}
	if !c.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", c, c.Published)
	}
	if !c.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", c, c.StartTime)
	}
	if !c.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", c, c.Updated)
	}
}

func TestDislike_UnmarshalJSON(t *testing.T) {
	d := Dislike{}

	dataEmpty := []byte("{}")
	d.UnmarshalJSON(dataEmpty)
	if d.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", d, d.ID)
	}
	if d.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", d, d.Type)
	}
	if d.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", d, d.AttributedTo)
	}
	if len(d.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", d, d.Name)
	}
	if len(d.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", d, d.Summary)
	}
	if len(d.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", d, d.Content)
	}
	if d.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", d, d.URL)
	}
	if !d.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", d, d.Published)
	}
	if !d.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", d, d.StartTime)
	}
	if !d.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", d, d.Updated)
	}
}

func TestLike_UnmarshalJSON(t *testing.T) {
	l := Like{}

	dataEmpty := []byte("{}")
	l.UnmarshalJSON(dataEmpty)
	if l.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", l, l.ID)
	}
	if l.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", l, l.Type)
	}
	if l.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", l, l.AttributedTo)
	}
	if len(l.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", l, l.Name)
	}
	if len(l.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", l, l.Summary)
	}
	if len(l.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", l, l.Content)
	}
	if l.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", l, l.URL)
	}
	if !l.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", l, l.Published)
	}
	if !l.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", l, l.StartTime)
	}
	if !l.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", l, l.Updated)
	}
}

func TestUpdate_UnmarshalJSON(t *testing.T) {
	u := Update{}

	dataEmpty := []byte("{}")
	u.UnmarshalJSON(dataEmpty)
	if u.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", u, u.ID)
	}
	if u.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", u, u.Type)
	}
	if u.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", u, u.AttributedTo)
	}
	if len(u.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", u, u.Name)
	}
	if len(u.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", u, u.Summary)
	}
	if len(u.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", u, u.Content)
	}
	if u.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", u, u.URL)
	}
	if !u.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", u, u.Published)
	}
	if !u.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", u, u.StartTime)
	}
	if !u.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", u, u.Updated)
	}
}
