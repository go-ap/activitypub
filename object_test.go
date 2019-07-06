package activitypub

import (
	"testing"
)

func validateEmptyObject(o Object, t *testing.T) {
	if o.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", o, o.ID)
	}
	if o.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", o, o.Type)
	}
	if o.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", o, o.AttributedTo)
	}
	if len(o.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", o, o.Name)
	}
	if len(o.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", o, o.Summary)
	}
	if len(o.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", o, o.Content)
	}
	if o.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", o, o.URL)
	}
	if !o.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", o, o.Published)
	}
	if !o.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", o, o.StartTime)
	}
	if !o.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", o, o.Updated)
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
