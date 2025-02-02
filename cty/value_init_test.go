package cty

import (
	"testing"
)

func TestSetVal(t *testing.T) {
	plain := SetVal([]Value{True})
	marked := SetVal([]Value{True}).Mark(1)
	deepMarked := SetVal([]Value{True.Mark(2), True.Mark(3)})

	if plain.RawEquals(marked) {
		t.Errorf("plain should be unequal to marked\nplain:  %#v\nmarked: %#v", plain, marked)
	}
	if marked.RawEquals(deepMarked) {
		t.Errorf("marked should be unequal to deepMarked\nmarked:      %#v\ndeepmarked: %#v", marked, deepMarked)
	}
	if got, want := marked.Marks(), NewValueMarks(1); !got.Equal(want) {
		t.Errorf("wrong marks for marked\ngot:  %#v\nwant: %#v", got, want)
	}
	if got, want := deepMarked.Marks(), NewValueMarks(2, 3); !got.Equal(want) {
		// Both 2 and 3 marks are preserved even though both of them are
		// marking the same value True, and thus the resulting set contains
		// only one element.
		t.Errorf("wrong marks for deepMarked\ngot:  %#v\nwant: %#v", got, want)
	}

	if got, want := deepMarked.unmarkForce(), SetVal([]Value{True}); !got.RawEquals(want) {
		t.Errorf("wrong unmarked value for deepMarked\ngot:  %#v\nwant: %#v", got, want)
	}
}
