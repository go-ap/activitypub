package activitypub

/*
func TestMarshalGob(t *testing.T) {
	tests := []struct {
		name    string
		it      Item
		want    []byte
		wantErr error
	}{
		{
			name: "empty object",
			it: &Object{
				ID: "test",
			},
			want:    []byte{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(make([]byte, 0))
			err := gob.NewEncoder(buf).Encode(tt.it)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("MarshalGob() error = %s, wantErr %v", err, tt.wantErr)
				return
			}

			it := new(Object)
			got := buf.Bytes()
			if err := gob.NewDecoder(bytes.NewReader(got)).Decode(it); err != nil {
				t.Errorf("Gob Decoding failed for previously generated output %v", err)
			}
			if tt.wantErr == nil {
				if !assertDeepEquals(t.Errorf, it, tt.it) {
					t.Errorf("Gob Decoded value is different got = %#v, want %#v", it, tt.it)
				}
			}
		})
	}
}
*/
