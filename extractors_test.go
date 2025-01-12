package activitypub

import "testing"

func TestContentOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContentOf(tt.arg); got != tt.want {
				t.Errorf("ContentOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameOf(tt.arg); got != tt.want {
				t.Errorf("NameOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPreferredNameOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreferredNameOf(tt.arg); got != tt.want {
				t.Errorf("PreferredNameOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummaryOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SummaryOf(tt.arg); got != tt.want {
				t.Errorf("SummaryOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
