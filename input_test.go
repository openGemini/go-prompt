package prompt

import (
	"bytes"
	"testing"
)

func TestRemoveAllControlSequences(t *testing.T) {
	f := NewFilter()
	scenarioTable := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "case1",
			input:    []byte{0x1b, 0x5b, 0x41},
			expected: []byte{},
		},
		{
			name:     "case2",
			input:    []byte{0x5b, 0x1b, 0x5b, 0x41, 0x5b},
			expected: []byte{0x5b, 0x5b},
		},
		{
			name:     "case3",
			input:    []byte{0x1b, 0x5b, 0x41, 0x1b, 0x5b, 0x41},
			expected: []byte{},
		},
		{
			name:     "case4",
			input:    []byte{0x1b, 0x5b, 0x41, 0x5b, 0x1b, 0x5b, 0x41},
			expected: []byte{0x5b},
		},
		{
			name:     "case5",
			input:    []byte{},
			expected: []byte{},
		},
		{
			name:     "case6",
			input:    []byte{'a'},
			expected: []byte{'a'},
		},
		{
			name:     "case7",
			input:    []byte{0x1b, 'a'},
			expected: []byte{'a'},
		},
	}
	for _, s := range scenarioTable {
		ret := RemoveAllControlSequences(s.input, f)
		if !bytes.Equal(ret, s.expected) {
			t.Errorf("%s: Should be %#v, but got %#v", s.name, s.expected, ret)
		}
	}
}

func TestPosixParserGetKey(t *testing.T) {
	scenarioTable := []struct {
		name     string
		input    []byte
		expected Key
	}{
		{
			name:     "escape",
			input:    []byte{0x1b},
			expected: Escape,
		},
		{
			name:     "undefined",
			input:    []byte{'a'},
			expected: NotDefined,
		},
	}

	for _, s := range scenarioTable {
		t.Run(s.name, func(t *testing.T) {
			key := GetKey(s.input)
			if key != s.expected {
				t.Errorf("Should be %s, but got %s", key, s.expected)
			}
		})
	}
}
