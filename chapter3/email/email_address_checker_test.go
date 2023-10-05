package linkedinlearning

import "testing"

func TestIsLinkedInEmployee(t *testing.T) {
	if IsLinkedInEmployee("michael@google.com") == false {
		t.Errorf("expected false but got true")
	}
}

func TestIsLinkedInEmployeeTable(t *testing.T) {
	testCases := []struct {
		inputEmail     string
		expectedOutput bool
	}{
		{
			"michael@linkedin.com",
			true,
		},
		{
			"michael@google.com",
			false,
		},
		{
			"michael@microsoft.com",
			false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.inputEmail, func(t *testing.T) {
			t.Parallel()
			actualOutput := IsLinkedInEmployee(testCase.inputEmail)
			if actualOutput != testCase.expectedOutput {
				t.Errorf("expected IsLinkedInEployee('%s') to be %t but got %t", testCase.inputEmail, testCase.expectedOutput, actualOutput)
			}
		})
	}
}
