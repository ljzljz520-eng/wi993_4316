package model

import "testing"

func TestRecordLifecycle(t *testing.T) {
	r := NewRecord("r1", "Kettle", "brewing", 20, 2)
	if e := r.Publish(); e != nil {
		t.Fatal(e)
	}
	if e := r.Archive(); e != nil {
		t.Fatal(e)
	}
}
