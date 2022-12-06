package activitypub

import "testing"

func TestFlatten(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemsEqual(t *testing.T) {
	type args struct {
		it   Item
		with Item
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil_items_equal",
			args: args{nil, nil},
			want: true,
		},
		{
			name: "nil_item_with_object",
			args: args{nil, &Object{}},
			want: false,
		},
		{
			name: "nil_item_with_object#1",
			args: args{&Object{}, nil},
			want: false,
		},
		{
			name: "empty_objects",
			args: args{&Object{}, &Object{}},
			want: true,
		},
		{
			name: "empty_objects_different_alias_type",
			args: args{&Activity{}, &Object{}},
			want: true,
		},
		{
			name: "empty_objects_different_alias_type#1",
			args: args{&Actor{}, &Object{}},
			want: true,
		},
		{
			name: "same_id_object",
			args: args{&Object{ID: "test"}, &Object{ID: "test"}},
			want: true,
		},
		{
			name: "same_id_object_different_alias",
			args: args{&Activity{ID: "test"}, &Object{ID: "test"}},
			want: true,
		},
		{
			name: "same_id_object_different_alias#1",
			args: args{&Activity{ID: "test"}, &Actor{ID: "test"}},
			want: true,
		},
		{
			name: "different_id_objects",
			args: args{&Object{ID: "test1"}, &Object{ID: "test"}},
			want: false,
		},
		{
			name: "different_id_types",
			args: args{&Object{ID: "test", Type: NoteType}, &Object{ID: "test", Type: ArticleType}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemsEqual(tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("ItemsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	type args struct {
		it Item
	}
	var (
		o      *Object
		col    *ItemCollection
		obNil  Item = o
		colNil Item = col
	)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "plain-nil",
			args: args{
				it: nil,
			},
			want: true,
		},
		{
			name: "interface-nil",
			args: args{
				it: Item(nil),
			},
			want: true,
		},
		{
			name: "object-nil",
			args: args{
				it: obNil,
			},
			want: true,
		},
		{
			name: "collection-nil",
			args: args{
				it: colNil,
			},
			want: true,
		},
		{
			name: "collection-not-nil",
			args: args{
				it: ItemCollection{},
			},
			want: false,
		},
		{
			name: "object-not-nil",
			args: args{
				it: &Object{},
			},
			want: false,
		},
		{
			name: "place-not-nil",
			args: args{
				it: &Place{},
			},
			want: false,
		},
		{
			name: "tombstone-not-nil",
			args: args{
				it: &Tombstone{},
			},
			want: false,
		},
		{
			name: "collection-not-nil",
			args: args{
				it: &Collection{},
			},
			want: false,
		},
		{
			name: "activity-not-nil",
			args: args{
				it: &Activity{},
			},
			want: false,
		},
		{
			name: "intransitive-activity-not-nil",
			args: args{
				it: &IntransitiveActivity{},
			},
			want: false,
		},
		{
			name: "actor-not-nil",
			args: args{
				it: &Actor{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.it); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemsEqual1(t *testing.T) {
	type args struct {
		it   Item
		with Item
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{},
			want: true,
		},
		{
			name: "equal empty items",
			args: args{
				it:   &Object{},
				with: &Actor{},
			},
			want: true,
		},
		{
			name: "equal same ID items",
			args: args{
				it:   &Object{ID: "example-1"},
				with: &Object{ID: "example-1"},
			},
			want: true,
		},
		{
			name: "different IDs",
			args: args{
				it:   &Object{ID: "example-1"},
				with: &Object{ID: "example-2"},
			},
			want: false,
		},
		{
			name: "different properties",
			args: args{
				it:   &Object{ID: "example-1"},
				with: &Object{Type: ArticleType},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemsEqual(tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("ItemsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsObject(t *testing.T) {
	type args struct {
		it Item
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{},
			want: false,
		},
		{
			name: "interface with nil value",
			args: args{Item(nil)},
			want: false,
		},
		{
			name: "empty object",
			args: args{Object{}},
			want: true,
		},
		{
			name: "pointer to empty object",
			args: args{&Object{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsObject(tt.args.it); got != tt.want {
				t.Errorf("IsObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
