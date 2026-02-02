package cleaner
import ("testing"; "fmt")

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:		"  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:		"Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:		"chariZarD      veNUSaur       rAICHu",
			expected:	[]string{"charizard", "venusaur", "raichu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of slice does not match expected slice. %d is not equal to %d (expected).", actualLen, expectedLen)
			t.FailNow()
		} else {
			fmt.Println("Length ok")
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match. %s does not match %s", word, expectedWord)
				t.FailNow()
			} else {
				fmt.Printf("Match ok. Case %d ", i)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
