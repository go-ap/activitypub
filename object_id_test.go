package activitypub

import (
	"bytes"
	"testing"
)

func TestID_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want ID
	}{
		{
			name: "nil",
			data: []byte(nil),
			want: "",
		},
		{
			name: "empty",
			data: []byte(""),
			want: "",
		},
		{
			name: "something",
			data: []byte("something"),
			want: "something",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ID("")
			got.UnmarshalJSON(tt.data)
			if got != tt.want {
				t.Errorf("UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		i       ID
		want    []byte
		wantErr error
	}{
		{
			name: "nil",
			i:    "",
			want: []byte(nil),
		},
		{
			name: "empty",
			i:    "",
			want: []byte(""),
		},
		{
			name: "something",
			i:    "something",
			want: []byte(`"something"`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.MarshalJSON()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("MarshalJSON() returned no error but expected %v", tt.wantErr)
				}
				if tt.wantErr.Error() != err.Error() {
					t.Errorf("MarshalJSON() returned error %v but expected %v", err, tt.wantErr)
				}
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
	t.Skip("TODO")
}

func TestID_IsValid(t *testing.T) {
	tests := []struct {
		name string
		i    ID
		want bool
	}{
		{
			name: "empty",
			i:    "",
			want: false,
		},
		{
			name: "something",
			i:    "something",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
