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
