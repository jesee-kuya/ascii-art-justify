package ascii_test

import (
	"testing"

	"ascii/ascii"
)

func TestArrange(t *testing.T) {
	var testCases = []struct {
		testCase string
		data     []string
		expected string
	}{
		{
			testCase: "Empty Input",
			data:     []string{},
			expected: "",
		},
		{
			testCase: "Single Word",
			data:     []string{"hello"},
			expected: "hello",
		},
		{
			testCase: "Multiple Words",
			data:     []string{"hello", "world"},
			expected: "hello world",
		},
		{
			testCase: "Strings with Spaces",
			data:     []string{"hello ", " world"},
			expected: "hello   world",
		},
		{
			testCase: "Special Characters",
			data:     []string{"foo!", "@bar", "#baz"},
			expected: "foo! @bar #baz",
		},
		{
			testCase: "Mixed Case",
			data:     []string{"Go", "is", "Awesome"},
			expected: "Go is Awesome",
		},
		{
			testCase: "Leading and Trailing Spaces",
			data:     []string{" leading", "trailing "},
			expected: " leading trailing ",
		},
	}

	for _, test := range testCases {
		result := ascii.Arrange(test.data)
		if result != test.expected {
			t.Errorf("Test case '%s' failed: Arrange(%v) = %q; want %q", test.testCase, test.data, result, test.expected)
		}
	}
}
