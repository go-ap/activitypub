package activitypub

import "testing"

func TestID_UnmarshalJSON(t *testing.T) {
	o := ID("")
	dataEmpty := []byte("")

	o.UnmarshalJSON(dataEmpty)
	if o != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", o, o)
	}
}

func TestID_IsValid(t *testing.T) {
	t.Skip("TODO")
}

func TestID_MarshalJSON(t *testing.T) {
	t.Skip("TODO")
}
