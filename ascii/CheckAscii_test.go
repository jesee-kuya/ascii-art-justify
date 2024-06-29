package ascii

import "testing"

func TestCheckAscii(t *testing.T) {
	testCases := []struct {
		testCase string
		data     []string
		expected bool
	}{
		{
			testCase: "Empty input",
			data:     []string{},
			expected: true,
		},
		{
			testCase: "All ASCII characters",
			data:     []string{"Hello", "World", "!@#$%^&*()"},
			expected: true,
		},
		{
			testCase: "Contains non-ASCII character",
			data:     []string{"Hello", "Wörld"},
			expected: false,
		},
		{
			testCase: "Multiple non-ASCII characters",
			data:     []string{"こんにちは", "世界"},
			expected: false,
		},
		{
			testCase: "Boundary ASCII characters",
			data:     []string{"Hello", string([]byte{126}), string([]byte{32})},
			expected: true,
		},
		{
			testCase: "Non-printable ASCII characters",
			data:     []string{"Hello", string([]byte{31})},
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.testCase, func(t *testing.T) {
			result := CheckAscii(test.data)
			if result != test.expected {
				t.Errorf("Test case '%s' failed: expected %v, got %v", test.testCase, test.expected, result)
			}
		})
	}
}
