package activitypub

import (
	"reflect"
	"testing"
)

func TestObjectNew(t *testing.T) {
	var testValue = ID("test")
	var testType = ArticleType

	o := ObjectNew(testType)
	o.ID = testValue

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ObjectNew("")
	n.ID = testValue
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if n.Type != ObjectType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ObjectType)
	}

}

func TestActivityVocabularyTypes_Contains(t *testing.T) {
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ObjectTypes {
			if ActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be invalid", inValidType)
			}
		}
		if ActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ActivityTypes {
			if !ActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if IntransitiveActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ActivityTypes {
			if IntransitiveActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be invalid", inValidType)
			}
		}
		if IntransitiveActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range IntransitiveActivityTypes {
			if !IntransitiveActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range CollectionManagementActivityTypes {
			if !CollectionManagementActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if CollectionManagementActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ContentManagementActivityTypes {
			if CollectionManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ReactionsActivityTypes {
			if CollectionManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}

	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ContentManagementActivityTypes {
			if !ContentManagementActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if ContentManagementActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range CollectionManagementActivityTypes {
			if ContentManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ReactionsActivityTypes {
			if ContentManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ReactionsActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ReactionsActivityTypes {
			if !ReactionsActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if ReactionsActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range CollectionManagementActivityTypes {
			if ReactionsActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ContentManagementActivityTypes {
			if ReactionsActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}
	{
		for _, validType := range CollectionTypes {
			if !CollectionTypes.Contains(validType) {
				t.Errorf("Generic Type '%#v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActorTypes.Contains(invalidType) {
			t.Errorf("APObject Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ActorTypes {
			if !ActorTypes.Contains(validType) {
				t.Errorf("APObject Type '%v' should be valid", validType)
			}
		}
	}
	{
		for _, validType := range GenericObjectTypes {
			if !GenericObjectTypes.Contains(validType) {
				t.Errorf("Generic Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ObjectTypes.Contains(invalidType) {
			t.Errorf("APObject Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ObjectTypes {
			if !ObjectTypes.Contains(validType) {
				t.Errorf("APObject Type '%v' should be valid", validType)
			}
		}
	}
}

func TestObject_IsLink(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if o.IsLink() {
		t.Errorf("%#v should not be a valid link", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if m.IsLink() {
		t.Errorf("%#v should not be a valid link", m.Type)
	}
}

func TestObject_IsObject(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if !o.IsObject() {
		t.Errorf("%#v should be a valid object", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestObjectsArr_Append(t *testing.T) {
	d := make(ItemCollection, 0)

	val := Object{ID: ID("grrr")}

	d.Append(val)

	if len(d) != 1 {
		t.Errorf("Objects array should have exactly an element")
	}
	if !reflect.DeepEqual(d[0], val) {
		t.Errorf("First item in object array does not match %q", val.ID)
	}
}

func TestRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	first := make(ItemCollection, 0)
	if len(first) != 0 {
		t.Errorf("Objects array should have exactly an element")
	}

	first.Append(bob)
	first.Append(alice)
	first.Append(foo)
	first.Append(bar)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	first.Append(bar)
	first.Append(alice)
	first.Append(foo)
	first.Append(bob)
	if len(first) != 8 {
		t.Errorf("Objects array should have exactly 8(eight) elements, not %d", len(first))
	}

	ItemCollectionDeduplication(&first)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	second := make(ItemCollection, 0)
	second.Append(bar)
	second.Append(foo)

	ItemCollectionDeduplication(&first, &second)
	if len(first) != 4 {
		t.Errorf("First Objects array should have exactly 8(eight) elements, not %d", len(first))
	}
	if len(second) != 0 {
		t.Errorf("Second Objects array should have exactly 0(zero) elements, not %d", len(second))
	}

	_, err := ItemCollectionDeduplication(&first, &second, nil)
	if err != nil {
		t.Errorf("Deduplication with empty array failed")
	}
}

func validateEmptyObject(o Object, t *testing.T) {
	if o.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", o, o.ID)
	}
	if o.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", o, o.Type)
	}
	if o.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", o, o.AttributedTo)
	}
	if len(o.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", o, o.Name)
	}
	if len(o.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", o, o.Summary)
	}
	if len(o.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", o, o.Content)
	}
	if o.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", o, o.URL)
	}
	if o.Icon != nil {
		t.Errorf("Unmarshaled object %T should have empty Icon, received %v", o, o.Icon)
	}
	if o.Image != nil {
		t.Errorf("Unmarshaled object %T should have empty Image, received %v", o, o.Image)
	}
	if !o.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", o, o.Published)
	}
	if !o.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty StartTime, received %q", o, o.StartTime)
	}
	if !o.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Updated, received %q", o, o.Updated)
	}
	if !o.EndTime.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty EndTime, received %q", o, o.EndTime)
	}
	if o.Duration != 0 {
		t.Errorf("Unmarshaled object %T should have empty Duration, received %q", o, o.Duration)
	}
	if len(o.To) > 0 {
		t.Errorf("Unmarshaled object %T should have empty To, received %q", o, o.To)
	}
	if len(o.Bto) > 0 {
		t.Errorf("Unmarshaled object %T should have empty Bto, received %q", o, o.Bto)
	}
	if len(o.CC) > 0 {
		t.Errorf("Unmarshaled object %T should have empty CC, received %q", o, o.CC)
	}
	if len(o.BCC) > 0 {
		t.Errorf("Unmarshaled object %T should have empty BCC, received %q", o, o.BCC)
	}
	validateEmptySource(o.Source, t)
}

func validateEmptySource(s Source, t *testing.T) {
	if s.MediaType != "" {
		t.Errorf("Unmarshalled object %T should have empty Source.MediaType, received %q", s, s.MediaType)
	}
	if s.Content != nil {
		t.Errorf("Unmarshalled object %T should have empty Source.Content, received %q", s, s.Content)
	}
}

func TestObject_UnmarshalJSON(t *testing.T) {
	o := Object{}

	dataEmpty := []byte("{}")
	o.UnmarshalJSON(dataEmpty)
	validateEmptyObject(o, t)
}

func TestMimeType_UnmarshalJSON(t *testing.T) {
	m := MimeType("")
	dataEmpty := []byte("")

	m.UnmarshalJSON(dataEmpty)
	if m != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", m, m)
	}
}

func TestLangRefValue_String(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRefValue_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRefValue_UnmarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRef_UnmarshalText(t *testing.T) {
	l := LangRef("")
	dataEmpty := []byte("")

	l.UnmarshalText(dataEmpty)
	if l != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", l, l)
	}
}

func TestObject_GetID(t *testing.T) {
	a := Object{}
	testVal := "crash$"
	a.ID = ID(testVal)
	if string(a.GetID()) != testVal {
		t.Errorf("%T should return %q, Received %q", a.GetID, testVal, a.GetID())
	}
}

func TestObject_GetLink(t *testing.T) {
	a := Object{}
	testVal := "crash$"
	a.ID = ID(testVal)
	if string(a.GetLink()) != testVal {
		t.Errorf("%T should return %q, Received %q", a.GetLink, testVal, a.GetLink())
	}
}

func TestObject_GetType(t *testing.T) {
	a := Object{}
	a.Type = ActorType
	if a.GetType() != ActorType {
		t.Errorf("%T should return %q, Received %q", a.GetType, ActorType, a.GetType())
	}
}

func TestToObject(t *testing.T) {
	var it Item
	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToObject(it)
	if err != nil {
		t.Error(err)
	}
	if o != ob {
		t.Errorf("Invalid activity returned by ToObject #%v", ob)
	}

	act := ActivityNew("test", CreateType, nil)
	it = act

	a, err := ToObject(it)
	if err != nil {
		t.Errorf("Error returned when calling ToObject with activity should be nil, received %s", err)
	}
	if a == nil {
		t.Errorf("Invalid return by ToObject #%v, should have not been nil", a)
	}
}

func TestFlattenObjectProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestToTombstone(t *testing.T) {
	t.Skipf("TODO")
}

func TestToRelationship(t *testing.T) {
	t.Skipf("TODO")
}

func TestObject_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollectionDeduplication(t *testing.T) {
	t.Skipf("TODO")
}

func TestSource_UnmarshalJSON(t *testing.T) {
	s := Source{}

	dataEmpty := []byte("{}")
	s.UnmarshalJSON(dataEmpty)
	validateEmptySource(s, t)
}

func TestGetAPSource(t *testing.T) {
	data := []byte(`{"source": {"content": "test", "mediaType": "text/plain" }}`)

	a := GetAPSource(data)

	if a.Content.First().String() != "test" {
		t.Errorf("Content didn't match test value. Received %q, expecting %q", a.Content, "test")
	}
	if a.MediaType != "text/plain" {
		t.Errorf("Content didn't match test value. Received %q, expecting %q", a.MediaType, "text/plain")
	}
}
