package ignite

import (
	"bytes"
	"errors"
	"testing"
)

func TestReadFile(t *testing.T) {
	type IgniteFileInput struct {
		path string
		cb func(string) ([]byte, error)
	}
	
	type ReadFileExp struct {
		file []byte
		err error
	}
	
	tests := map[string]struct {
		input IgniteFileInput
		exp ReadFileExp
	} {
		"simple": {
			input: IgniteFileInput{
				path: "asdf",
				cb: func(string) ([]byte, error) {return []byte("a"), nil},
			},
			exp: ReadFileExp{
				[]byte("a"),
				nil,
			},
		},
		"error": {
			input: IgniteFileInput{
				path: "asdf",
				cb: func(string) ([]byte, error) {return nil, errors.New("fail")},
			},
			exp: ReadFileExp{
				nil,
				errors.New("fail"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actFile, actErr := readFile(test.input.path, test.input.cb)

			if !bytes.Equal(actFile, test.exp.file) {
				t.Errorf("Returned file doesn't match expected. Expected %v, got %v", test.exp.file, actFile)
			}

			if test.exp.err != nil {
				if !(actErr.Error() == test.exp.err.Error()) {
					t.Errorf("Returned error doesn't match expected. Expected %v, got %v", test.exp.err, actErr)
				}
			}
		})
	}
}
