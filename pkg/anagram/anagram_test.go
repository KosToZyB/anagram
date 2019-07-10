package anagram

import (
	"reflect"
	"testing"
)

func TestAnagram_GetAnagram(t *testing.T) {
	type fields struct {
		data []string
	}
	type args struct {
		word string
	}

	type test struct {
		name   string
		fields fields
		args   args
		want   []string
	}

	array1 := []string{"foobar", "aabb", "baba", "boofar", "test"}
	test1 := test{
		name:   "test1",
		fields: fields{
			data: array1,
		},
		args:   args{
			word: "foobar",
		},
		want: []string{"boofar", "foobar"},
	}

	test2 := test{
		name:   "test2",
		fields: fields{
			data: array1,
		},
		args:   args{
			word: "raboof",
		},
		want: []string{"boofar", "foobar"},
	}

	test3 := test{
		name:   "test3",
		fields: fields{
			data: array1,
		},
		args:   args{
			word: "abba",
		},
		want: []string{"aabb", "baba"},
	}

	test4 := test{
		name:   "test4",
		fields: fields{
			data: array1,
		},
		args:   args{
			word: "test",
		},
		want: []string{"test"},
	}

	test5 := test{
		name:   "test5",
		fields: fields{
			data: array1,
		},
		args:   args{
			word: "qwerty",
		},
		want: []string{},
	}

	tests := []test{test1, test2, test3, test4, test5}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAnagram()
			a.Load(tt.fields.data)

			if got := a.GetAnagram(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Anagram.GetAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
