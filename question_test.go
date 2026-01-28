package activitypub

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuestionNew(t *testing.T) {
	testValue := ID("test")

	a := QuestionNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if !a.Matches(QuestionType) {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.GetType(), QuestionType)
	}
}

func TestQuestion_GetID(t *testing.T) {
	a := QuestionNew("test")

	if a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), a.GetID())
	}
}

func TestQuestion_IsObject(t *testing.T) {
	a := QuestionNew("test")

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestQuestion_IsLink(t *testing.T) {
	a := QuestionNew("test")

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestQuestion_GetLink(t *testing.T) {
	a := QuestionNew("test")

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestQuestion_GetType(t *testing.T) {
	a := QuestionNew("test")

	if !a.GetType().Matches(QuestionType) {
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

func TestQuestion_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestQuestion_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}
