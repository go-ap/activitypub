package activitypub

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

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
