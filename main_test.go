package main

import "testing"

func TestMaxWidth(t *testing.T) {
	// Level 2: nodes at positions 0 and 3 -> width 4 (two gaps inside).
	r1 := tn(1, tn(3, tn(5, nil, nil), tn(3, nil, nil)), tn(2, nil, tn(9, nil, nil)))
	if got := MaxWidth(r1); got != 4 {
		t.Errorf("MaxWidth = %d, want 4", got)
	}

	// Only left children: width stays 1.
	r2 := tn(1, tn(2, tn(3, nil, nil), nil), nil)
	if got := MaxWidth(r2); got != 1 {
		t.Errorf("MaxWidth = %d, want 1", got)
	}

	// Gap spans the root's two subtrees.
	r3 := tn(1, tn(3, tn(5, nil, nil), nil), tn(2, nil, tn(9, nil, nil)))
	if got := MaxWidth(r3); got != 4 {
		t.Errorf("MaxWidth = %d, want 4", got)
	}

	if got := MaxWidth(nil); got != 0 {
		t.Errorf("MaxWidth(nil) = %d, want 0", got)
	}
	if got := MaxWidth(tn(7, nil, nil)); got != 1 {
		t.Errorf("single node = %d, want 1", got)
	}
}
