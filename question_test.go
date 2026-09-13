package activitypub

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuestion_GetID(t *testing.T) {
	a := &Question{Type: QuestionType, ID: "test"}

	if a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), a.GetID())
	}
}

func TestQuestion_GetLink(t *testing.T) {
	a := &Question{Type: QuestionType, ID: "test"}

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestQuestion_GetType(t *testing.T) {
	a := &Question{Type: QuestionType, ID: "test"}

	if !QuestionType.Match(a.GetType()) {
		t.Errorf("GetType should return %q for %T, received %q", QuestionType, a, a.GetType())
	}
}

func TestToQuestion(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Question
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Question",
			it:   Question{ID: "test", Type: TravelType},
			want: &Question{ID: "test", Type: TravelType},
		},
		{
			name: "Valid *Question",
			it:   &Question{ID: "test", Type: ArriveType},
			want: &Question{ID: "test", Type: ArriveType},
		},
		{
			name: "Valid Question",
			it:   Question{ID: "test", Type: QuestionType},
			want: &Question{ID: "test", Type: QuestionType},
		},
		{
			name: "Valid *Question",
			it:   &Question{ID: "test", Type: QuestionType},
			want: &Question{ID: "test", Type: QuestionType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Question](IRI("")),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[Question](new(IntransitiveActivity)),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: UpdateType},
			wantErr: ErrorInvalidType[Question](new(Activity)),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Question](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Question](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Question](&Object{}),
		},
		{
			name:    "Actor",
			it:      &Actor{ID: "test", Type: PersonType},
			wantErr: ErrorInvalidType[Question](&Person{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToQuestion(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToQuestion() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
func TestToQuestion1(t *testing.T) {
	var it Item
	act := &Question{Type: QuestionType, ID: "test"}
	it = act

	a, err := ToQuestion(it)
	if err != nil {
		t.Error(err)
	}
	if a != act {
		t.Errorf("Invalid activity returned by ToActivity #%v", a)
	}

	ob := &Object{Type: ArticleType}
	it = ob

	o, err := ToQuestion(it)
	if err == nil {
		t.Errorf("Error returned when calling ToActivity with object should not be nil")
	}
	if o != nil {
		t.Errorf("Invalid return by ToActivity #%v, should have been nil", o)
	}
}

func TestQuestion_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestQuestion_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestQuestion_Clean(t *testing.T) {
	tests := []struct {
		name string
		i    Question
		want Item
	}{
		{
			name: "empty",
			i:    Question{},
			want: &Question{},
		},
		{
			name: "has Bto",
			i:    Question{Type: QuestionType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Question{Type: QuestionType},
		},
		{
			name: "has BCC",
			i:    Question{Type: QuestionType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Question{Type: QuestionType},
		},
		{
			name: "audience has BCC",
			i:    Question{Type: QuestionType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Question{Type: QuestionType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			i:    Question{Type: QuestionType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			i:    Question{Type: QuestionType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			i:    Question{Type: QuestionType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			i:    Question{Type: QuestionType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			i:    Question{Type: QuestionType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			i:    Question{Type: QuestionType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			i:    Question{Type: QuestionType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			i:    Question{Type: QuestionType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Question{Type: QuestionType, Tag: ItemCollection{&Object{}}},
		},
		{
			name: "actor has BCC",
			i:    Question{Type: QuestionType, Actor: &Person{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Actor: &Person{}},
		},
		{
			name: "origin has BCC",
			i:    Question{Type: QuestionType, Origin: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Origin: &Object{}},
		},
		{
			name: "target has BCC",
			i:    Question{Type: QuestionType, Target: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Target: &Object{}},
		},
		{
			name: "instrument has BCC",
			i:    Question{Type: QuestionType, Instrument: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Question{Type: QuestionType, Instrument: &Object{}},
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
