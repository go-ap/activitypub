package activitypub

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	fnIA = func(_ *IntransitiveActivity) error { return nil }

	notIntransitiveActivity   Item = new(Activity)
	maybeIntransitiveActivity Item = new(IntransitiveActivity)
)

func Benchmark_ToIntransitiveActivityHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToIntransitiveActivity(maybeIntransitiveActivity)
	}
}

func Benchmark_To_T_IntransitiveActivityHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[IntransitiveActivity](maybeIntransitiveActivity)
	}
}

func Benchmark_ToIntransitiveActivityNotHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToIntransitiveActivity(notIntransitiveActivity)
	}
}

func Benchmark_To_T_IntransitiveActivityNotHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[IntransitiveActivity](notIntransitiveActivity)
	}
}

func Benchmark_OnIntransitiveActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnIntransitiveActivity(maybeIntransitiveActivity, fnIA)
	}
}

func Benchmark_On_T_IntransitiveActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[IntransitiveActivity](maybeIntransitiveActivity, fnIA)
	}
}

func TestIntransitiveActivityRecipients(t *testing.T) {
	bob := &Person{Type: PersonType, ID: "bob"}
	alice := &Person{Type: PersonType, ID: "alice"}
	foo := &Organization{Type: OrganizationType, ID: "foo"}
	bar := &Group{Type: GroupType, ID: "bar"}

	a := &IntransitiveActivity{ID: "test", Type: ActivityVocabularyType("t")}

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
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(eight) elements, not %d", a, len(a.To))
	}

	a.Recipients()
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", a, len(a.To))
	}

	b := &Activity{Type: ActivityVocabularyType("t"), ID: "test"}

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

func TestIntransitiveActivity_GetLink(t *testing.T) {
	i := &IntransitiveActivity{Type: QuestionType, ID: "test"}

	if i.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", i, i, i)
	}
}

func TestIntransitiveActivity_GetObject(t *testing.T) {
	i := &IntransitiveActivity{Type: QuestionType, ID: "test"}

	if i.GetID() != "test" || !i.Match(QuestionType) {
		t.Errorf("%T should not return an empty %T object. Received %#v", i, i, i)
	}
}

func TestIntransitiveActivity_Recipients(t *testing.T) {
	to := &Person{Type: PersonType, ID: "bob"}
	o := Object{Type: ArticleType}
	cc := &Person{Type: PersonType, ID: "alice"}

	o.ID = "something"

	c := &IntransitiveActivity{Type: IntransitiveActivityType, ID: "act"}
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

func TestIntransitiveActivity_GetID(t *testing.T) {
	a := &IntransitiveActivity{Type: IntransitiveActivityType, ID: "test"}

	if a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), a.GetID())
	}
}

func TestIntransitiveActivity_GetType(t *testing.T) {
	{
		a := &IntransitiveActivity{Type: IntransitiveActivityType, ID: "test"}
		if !a.Match(IntransitiveActivityType) {
			t.Errorf("GetType should return %q for %T, received %q", IntransitiveActivityType, a, a.GetType())
		}
	}
	{
		a := &IntransitiveActivity{Type: ArriveType, ID: "test"}
		if !a.Match(ArriveType) {
			t.Errorf("GetType should return %q for %T, received %q", ArriveType, a, a.GetType())
		}
	}
	{
		a := &IntransitiveActivity{Type: QuestionType, ID: "test"}
		if !a.Match(QuestionType) {
			t.Errorf("GetType should return %q for %T, received %q", QuestionType, a, a.GetType())
		}
	}
}

func TestToIntransitiveActivity(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *IntransitiveActivity
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid IntransitiveActivity",
			it:   IntransitiveActivity{ID: "test", Type: TravelType},
			want: &IntransitiveActivity{ID: "test", Type: TravelType},
		},
		{
			name: "Valid *IntransitiveActivity",
			it:   &IntransitiveActivity{ID: "test", Type: ArriveType},
			want: &IntransitiveActivity{ID: "test", Type: ArriveType},
		},
		{
			name: "Valid Question",
			it:   Question{ID: "test", Type: QuestionType},
			want: &IntransitiveActivity{ID: "test", Type: QuestionType},
		},
		{
			name: "Valid *Question",
			it:   &Question{ID: "test", Type: QuestionType},
			want: &IntransitiveActivity{ID: "test", Type: QuestionType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[IntransitiveActivity](IRI("")),
		},
		{
			name: "Activity",
			it:   &Activity{ID: "test", Type: UpdateType},
			want: &IntransitiveActivity{ID: "test", Type: UpdateType},
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[IntransitiveActivity](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[IntransitiveActivity](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[IntransitiveActivity](&Object{}),
		},
		{
			name:    "Actor",
			it:      &Actor{ID: "test", Type: PersonType},
			wantErr: ErrorInvalidType[IntransitiveActivity](&Person{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntransitiveActivity(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToIntransitiveActivity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToIntransitiveActivity() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestIntransitiveActivity_Clean(t *testing.T) {
	tests := []struct {
		name string
		i    IntransitiveActivity
		want Item
	}{
		{
			name: "empty",
			i:    IntransitiveActivity{},
			want: &IntransitiveActivity{},
		},
		{
			name: "has Bto",
			i:    IntransitiveActivity{Type: ArriveType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &IntransitiveActivity{Type: ArriveType},
		},
		{
			name: "has BCC",
			i:    IntransitiveActivity{Type: ArriveType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &IntransitiveActivity{Type: ArriveType},
		},
		{
			name: "audience has BCC",
			i:    IntransitiveActivity{Type: ArriveType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &IntransitiveActivity{Type: ArriveType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			i:    IntransitiveActivity{Type: ArriveType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: ArriveType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			i:    IntransitiveActivity{Type: TravelType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: TravelType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			i:    IntransitiveActivity{Type: TravelType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: TravelType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			i:    IntransitiveActivity{Type: TravelType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: TravelType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			i:    IntransitiveActivity{Type: ArriveType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: ArriveType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			i:    IntransitiveActivity{Type: ArriveType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: ArriveType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			i:    IntransitiveActivity{Type: ArriveType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: ArriveType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			i:    IntransitiveActivity{Type: QuestionType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &IntransitiveActivity{Type: QuestionType, Tag: ItemCollection{&Object{}}},
		},
		{
			name: "actor has BCC",
			i:    IntransitiveActivity{Type: QuestionType, Actor: &Person{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: QuestionType, Actor: &Person{}},
		},
		{
			name: "origin has BCC",
			i:    IntransitiveActivity{Type: QuestionType, Origin: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: QuestionType, Origin: &Object{}},
		},
		{
			name: "target has BCC",
			i:    IntransitiveActivity{Type: QuestionType, Target: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: QuestionType, Target: &Object{}},
		},
		{
			name: "instrument has BCC",
			i:    IntransitiveActivity{Type: QuestionType, Instrument: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &IntransitiveActivity{Type: QuestionType, Instrument: &Object{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Clean(); !cmp.Equal(got, tt.want, EquateItems) {
				t.Errorf("Clean() = %s", cmp.Diff(tt.want, got, EquateItems))
			}
		})
	}
}

func TestIntransitiveActivity_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestIntransitiveActivity_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestIntransitiveActivity_Equals(t *testing.T) {
	type fields struct {
		ID           ID
		Type         Typer
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
		URL          Item
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
	}
	tests := []struct {
		name   string
		fields fields
		arg    Item
		want   bool
	}{
		{
			name:   "equal-empty-intransitive-activity",
			fields: fields{},
			arg:    IntransitiveActivity{},
			want:   true,
		},
		{
			name:   "equal-intransitive-activity-just-id",
			fields: fields{ID: "test"},
			arg:    IntransitiveActivity{ID: "test"},
			want:   true,
		},
		{
			name:   "equal-intransitive-activity-id",
			fields: fields{ID: "test", URL: IRI("example.com")},
			arg:    IntransitiveActivity{ID: "test"},
			want:   false,
		},
		{
			name:   "equal-false-with-id-and-url",
			fields: fields{ID: "test"},
			arg:    IntransitiveActivity{ID: "test", URL: IRI("example.com")},
			want:   false,
		},
		{
			name:   "not a valid intransitive-activity",
			fields: fields{ID: "http://example.com"},
			arg:    Link{ID: "http://example.com"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := IntransitiveActivity{
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
			if got := a.Equals(tt.arg); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIntransitiveActivity_initialization() {
	// intransitiveActivity1 is a struct literal which can be operated on directly.
	// For example, we can set the Actor:
	intransitiveActivity1 := IntransitiveActivity{ID: "http://example.com/1"}
	intransitiveActivity1.Actor = IRI("http://example.com/~jdoe")
	fmt.Printf("IntransitiveActivity1: %v\n", intransitiveActivity1)

	// intransitiveActivity2 is wrapped in an Item interface.
	// It is probably the most common way of interacting with objects in the library,
	// because usually they get unmarshaled from an HTTP request, or another representation.
	var intransitiveActivity2 Item = &IntransitiveActivity{Type: IntransitiveActivityType}

	// That means we can't set any properties directly, so
	// if you uncomment the next line you will get a compiler error.
	// intransitiveActivity2.Actor = IRI("http://example.com/~jdoe")

	_ = OnIntransitiveActivity(intransitiveActivity2, func(intransitiveActivity *IntransitiveActivity) error {
		// In order to operate on it, we must wrap it in a call to OnIntransitiveActivity
		intransitiveActivity.Actor = IRI("http://example.com/~jdoe")
		return nil
	})
	fmt.Printf("IntransitiveActivity2: %v\n", intransitiveActivity2)

	// Output:
	// IntransitiveActivity1: activitypub.IntransitiveActivity { id: http://example.com/1, actor: http://example.com/~jdoe }
	// IntransitiveActivity2: activitypub.IntransitiveActivity[IntransitiveActivity] { actor: http://example.com/~jdoe }
}

func ExampleToIntransitiveActivity() {
	// IntransitiveActivities form a special case of activities that do not contain an Object
	// property that they operate on.

	// We can see here an initialization for the Arrive struct which is an alias for
	// the IntransitiveActivity type.
	// Using it makes no difference on a semantic level, but it can provide more context about the
	// intention to other developers.
	var intransitiveActivity1 Item = &Arrive{Type: UpdateType}

	// There are two other intransitive activities in the Activity Vocabulary, the Travel type
	// which is also an alias, and the Question which is a disjoint type as it contains additional
	// properties.
	//
	// We mentioned before that the Type needs to be set manually, so developers also need to take
	// care about their correctness. An invalid type can be set onto an activity and there's no
	// mechanism for validating that - at least at the moment.
	// As an example, the above initialization contains uses the semantically invalid Update type.

	// As we've seen previously, we can't operate directly on intransitiveActivity1 as it's an Item
	// instance, so uncommenting the following line will trigger a compiler error:
	//activity1.Actor = IRI("http://example.com/~jdoe")

	ia1, _ := ToIntransitiveActivity(intransitiveActivity1)
	ia1.Type = ArriveType
	ia1.Target = IRI("http://example.com/ys")
	fmt.Printf("IntransitiveActivity1: %v\n", intransitiveActivity1)
	fmt.Printf("                     : %v\n\n", ia1)

	// Here is how a Question value can interact with the conversion to IntransitiveActivity.
	var question Item = &Question{ID: "http://example.com/huh", AnyOf: ItemCollection{}}
	q, _ := ToIntransitiveActivity(question)
	// We can set intransitive activity properties on the new value,
	// and they will be reflected in the original object.
	q.Actor = IRI("http://example.com/~jdoe")
	fmt.Printf("Question: %v\n", question)
	// But also it results in loss of data, as the anyOf properties
	// are no longer accessible in the converted value.
	fmt.Printf("        : %v\n", q)

	// Normally this is not a problem, because, as we mentioned in the ExampleToActor,
	// assigning back to the interface value should be avoided.
	// Losing the data of the specific type being the major reason as to why not.
	//
	// If we uncomment the following line, the test fails as
	// we lose the question specific data permanently:
	//question = q
	fmt.Printf("Question: %v\n\n", question)

	// Another consideration is that the hierarchy of types that can be converted is unidirectional.
	// As an example, an Object type can't be converted to an IntransitiveActivity type,
	// and trying to do so results in an error.
	var notWhatWeWant Item = &Object{Type: TravelType}
	na, err := ToIntransitiveActivity(notWhatWeWant)
	fmt.Printf("NotWhatWeWant: %v\n", notWhatWeWant)
	fmt.Printf("             : %v\n", na)
	fmt.Printf("Error        : %v\n\n", err)

	// Output:
	// IntransitiveActivity1: activitypub.IntransitiveActivity[Arrive] { target: http://example.com/ys }
	//                      : activitypub.IntransitiveActivity[Arrive] { target: http://example.com/ys }
	//
	// Question: activitypub.Question { id: http://example.com/huh, actor: http://example.com/~jdoe, anyOf: [] }
	//         : activitypub.IntransitiveActivity { id: http://example.com/huh, actor: http://example.com/~jdoe }
	// Question: activitypub.Question { id: http://example.com/huh, actor: http://example.com/~jdoe, anyOf: [] }
	//
	// NotWhatWeWant: activitypub.Object[Travel] {  }
	//              : <nil>
	// Error        : unable to convert *activitypub.Object to *activitypub.IntransitiveActivity
}

func ExampleToIntransitiveActivity_returning() {
	// Similarly to the ExampleToObject_returning(), returning from
	// functions that deal with IntransitiveActivities will result
	// in lost information when the pointer wraps an Activity type.

	mangleItems := func(it Item) Item {
		ob, _ := ToIntransitiveActivity(it)
		ob.Actor = IRI("http://example.com/~jdoe")

		// This is a bug, the Object property is lost for Activity objects.
		return ob
	}

	var originalA Item = &Activity{Type: LikeType, Object: IRI("http://example.com/bob")}
	mangledA := mangleItems(originalA)
	fmt.Printf("OriginalA: %v\n", originalA)
	fmt.Printf(" MangledA: %v\n", mangledA)

	var originalQ Item = &Question{Type: QuestionType, AnyOf: IRIs{"http://example.com/1", "http://example.com/2"}}
	mangledQ := mangleItems(originalQ)
	fmt.Printf("OriginalQ: %v\n", originalQ)
	fmt.Printf(" MangledQ: %v\n", mangledQ)

	// Output:
	// OriginalA: activitypub.Activity[Like] { actor: http://example.com/~jdoe, object: http://example.com/bob }
	//  MangledA: activitypub.IntransitiveActivity[Like] { actor: http://example.com/~jdoe }
	// OriginalQ: activitypub.Question[Question] { actor: http://example.com/~jdoe, anyOf: [http://example.com/1 http://example.com/2] }
	//  MangledQ: activitypub.IntransitiveActivity[Question] { actor: http://example.com/~jdoe }
}

func ExampleOnIntransitiveActivity() {
	// In the ExampleToIntransitiveActivity() function, we saw how we can convert data types
	// to IntransitiveActivity pointer values and be allowed to use their properties in that way.
	//
	// Here we can see how this mechanism can be used to build specific logic when dealing
	// with opaque Item interface values.

	// As we've seen, you can not access this IntransitiveActivity's properties
	// because it's wrapped in the Item interface.
	var intransitiveActivity1 Item = &IntransitiveActivity{Type: ArriveType, Target: IRI("http://example.com/ys")}
	// Uncommenting this line will trigger a compiler error.
	//intransitiveActivity1.Actor = IRI("http://example.com/~jdoe")

	_ = OnIntransitiveActivity(intransitiveActivity1, func(ia *IntransitiveActivity) error {
		// Instead we can wrap it in an OnIntransitiveActivity() call in which
		// we can modify it, and the changes will be visible outside its scope.
		ia.Actor = IRI("http://example.com/~jdoe")

		// Similarly, as we've seen in the ExampleToIntransitiveActivity(), we can also modify
		// the properties in common with the Object type, without needing
		// a call to OnObject().
		ia.Summary = DefaultNaturalLanguage("I made it!")
		return nil
	})
	fmt.Printf("IntransitiveActivity1: %v\n", intransitiveActivity1)

	// We can also modify the Question type, which is disjoint to IntransitiveActivity,
	// but only the properties they have in common.
	var question Item = &Question{Type: QuestionType, AnyOf: IRIs{"http://example.com/yay", "http://example.com/nay"}}
	_ = OnIntransitiveActivity(question, func(q *IntransitiveActivity) error {
		q.Actor = IRI("http://example.com/~jdoe")
		// We can't access Question specific properties, so
		// uncommenting this will trigger a compiler error.
		//q.AnyOf = nil
		return nil
	})
	fmt.Printf("Question: %v\n", question)

	// Output:
	// IntransitiveActivity1: activitypub.IntransitiveActivity[Arrive] { summary: I made it!, actor: http://example.com/~jdoe, target: http://example.com/ys }
	// Question: activitypub.Question[Question] { actor: http://example.com/~jdoe, anyOf: [http://example.com/yay http://example.com/nay] }
}
