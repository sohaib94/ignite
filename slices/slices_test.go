package slices

import "testing"

func TestUnorderedEqualsStrings(t *testing.T) {

	tests := map[string]struct {
		input1 []string
		input2 []string
		exp bool
	}{
		"unequal strings length": {
			input1: []string{"a", "b"},
			input2: []string{"a"},
			exp: false,
		},
		"unequal strings": {
			input1: []string{"a", "b"},
			input2: []string{"a", "c"},
			exp: false,
		},
		"equal strings same order": {
			input1: []string{"a", "b"},
			input2: []string{"a", "b"},
			exp: true,
		},
		"equal strings diff order": {
			input1: []string{"a", "b"},
			input2: []string{"b", "a"},
			exp: true,
		},
		"empty slices equal": {
			input1: []string{},
			input2: []string{},
			exp: true,
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			r := UnorderedEquals(test.input1, test.input2)

			if r != test.exp {
				t.Errorf("Failed test %v. \nExpected: %v \nGot: %v", name, test.exp, r)
			}
		})
	}
}

func TestUnorderedEqualsInt(t *testing.T) {
	tests := map[string]struct {
		input1 []int
		input2 []int
		exp bool
	}{
		"unequal ints length": {
			input1: []int{12, 13},
			input2: []int{13},
			exp: false,
		},
		"unequal ints": {
			input1: []int{12, 13},
			input2: []int{13, 14},
			exp: false,
		},
		"equal ints same order": {
			input1: []int{12, 13},
			input2: []int{12, 13},
			exp: true,
		},
		"equal ints diff order": {
			input1: []int{12, 13},
			input2: []int{13, 12},
			exp: true,
		},
		"equal ints diff types": {
			input1: []int{12, 13},
			input2: []int{13, 12},
			exp: true,
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			r := UnorderedEquals(test.input1, test.input2)

			if r != test.exp {
				t.Errorf("Failed test %v. \nExpected: %v \nGot: %v", name, test.exp, r)
			}
		})
	}
}
