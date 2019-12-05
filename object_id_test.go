package activitypub

import "testing"

func TestObjectID_UnmarshalJSON(t *testing.T) {
	o := ID("")
	dataEmpty := []byte("")

	o.UnmarshalJSON(dataEmpty)
	if o != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", o, o)
	}
}
