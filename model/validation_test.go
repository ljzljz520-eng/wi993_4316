package model

import "testing"

func TestValidationRules(t *testing.T) {
	if ValidateID("x!") == nil {
		t.Fatal("expected invalid")
	}
	if PriceBand(80) != "standard" {
		t.Fatal("band")
	}
	if !CanTransition("draft", "published") {
		t.Fatal("transition")
	}
}
