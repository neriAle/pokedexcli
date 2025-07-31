package repl

import(
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " Hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "SOME things NeVeR chanGE",
			expected: []string{"some", "things", "never", "change"},
		},
		{
			input: "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("size of slice is incorrect")
		} else {
			for i, _ := range actual {
				word := actual[i]
				exWord := c.expected[i]
				if word != exWord {
					t.Errorf("'%s' found, but expected '%s'", word, exWord)
					break
				}
 			}
		}
	}
}