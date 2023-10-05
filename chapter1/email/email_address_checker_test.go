package linkedinlearning

import "testing"

func TestIsLinkedInEmployee(t *testing.T) {
	if IsLinkedInEmployee("michael@google.com") == false {
		t.Errorf("expected false but got true")
	}
}
