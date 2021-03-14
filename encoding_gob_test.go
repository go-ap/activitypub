package activitypub

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"
)

func TestMarshalGob(t *testing.T) {
	tests := []struct {
		name    string
		it      Item
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty object",
			it:      &Object{
				ID: "test",
			},
			want:    []byte{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got, err := MarshalGob(tt.it)
			buf := bytes.NewBuffer(make([]byte,0))
			err := gob.NewEncoder(buf).Encode(tt.it)
			got := buf.Bytes()
			
			it := new(Object)
			gob.NewDecoder(bytes.NewReader(got)).Decode(it)
			if it.ID == tt.it.GetID() {
				t.Logf("Yay!")
			}
			
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalGob() error = %s, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalGob() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
