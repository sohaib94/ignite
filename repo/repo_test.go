package repo

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sl "github.com/sohaib94/ignite/slices"
)

func TestCreate(t *testing.T) {

}

func TestCreateLocal(t *testing.T) {

	type ExpOutput struct {
		createdDirs  []string
		createdFiles []string
		err          error
	}

	tests := map[string]struct {
		input Repo
		exp   ExpOutput
	}{
		"single file": {
			input: Repo{
				Meta: Meta{
					Host: "github.com",
					Org:  "test",
					Conn: "ssh",
				},
				Root: Directory{
					Files: []string{"a"},
				},
			},
			exp: ExpOutput{
				createdDirs:  []string{},
				createdFiles: []string{"a"},
				err:          nil,
			},
		},
		"single subdirectory": {
			input: Repo{
				Meta: Meta{
					Host: "github.com",
					Org:  "test",
					Conn: "ssh",
				},
				Root: Directory{
					Directories: map[string]Directory{
						"a": {},
					},
				},
			},
			exp: ExpOutput{
				createdDirs:  []string{"a"},
				createdFiles: []string{},
				err:          nil,
			},
		},
		"nested": {
			input: Repo{
				Meta: Meta{
					Host: "github.com",
					Org:  "test",
					Conn: "ssh",
				},
				Root: Directory{
					Directories: map[string]Directory{
						"a": {
							Directories: map[string]Directory{
								"b": {
									Files: []string{"e"},
								},
							},
							Files: []string{"d"},
						},
					},
					Files: []string{"c"},
				},
			},
			exp: ExpOutput{
				createdDirs:  []string{"a", "a/b"},
				createdFiles: []string{"c", "a/d", "a/b/e"},
				err:          nil,
			},
		},
		"no cloning": {
			input: Repo{
				Meta: Meta{
					Host: "github.com",
					Org:  "test",
					Conn: "ssh",
				},
				Root: Directory{
					Files: []string{"a", "b"},
					Directories: map[string]Directory{
						"test": {
							Files: []string{"c", "d"},
						},
					},
				},
			},
			exp: ExpOutput{
				createdDirs:  []string{"test"},
				createdFiles: []string{"a", "b", "test/c", "test/d"},
				err:          nil,
			},
		},
	}

	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalf("Unable to create temp dir for testing")
	}
	defer os.RemoveAll(tempDir)

	for name, test := range tests {

		testDir, err := os.MkdirTemp(tempDir, strings.ReplaceAll(name, " ", "_"))
		defer os.RemoveAll(testDir)
		if err != nil {
			log.Fatalf("Unable to create temp dir for testing %v", name)
		}

		t.Run(name, func(t *testing.T) {
			err := test.input.CreateLocalRepo(testDir)
			var createdDirs []string
			var createdFiles []string

			walkErr := filepath.WalkDir(testDir, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}

				cleanPath := path[len(testDir):]

				if d.IsDir() {
					if path != testDir {
						createdDirs = append(createdDirs, cleanPath[1:])
					}
				} else {
					createdFiles = append(createdFiles, cleanPath[1:])
				}

				return nil
			})

			if walkErr != nil {
				log.Fatalf("Failed to walk temp dir in test %v", name)
			}

			if test.exp.err != nil {
				if !(err.Error() == test.exp.err.Error()) {
					t.Errorf("Returned error doesn't match expected. \nExpected %v \nGot %v", test.exp.err, err)
				}
			}

			if !sl.UnorderedEquals(createdDirs, test.exp.createdDirs) || !sl.UnorderedEquals(createdFiles, test.exp.createdFiles) {
				t.Errorf("Created dirs don't match expected dirs. \nExpected %v \nGot %v", test.exp.createdDirs, createdDirs)
			}
			if !sl.UnorderedEquals(createdFiles, test.exp.createdFiles) {
				t.Errorf("Created files don't match expected files. \nExpected %v \nGot %v", test.exp.createdFiles, createdFiles)
			}
		})
	}
}

func TestCreateRepo(t *testing.T) {

}

func TestCleanRepoName(t *testing.T) {
	tests := map[string]struct {
		input string
		exp   string
	}{
		"only spaces": {
			input: "hello world",
			exp:   "hello_world",
		},
		"upper case": {
			input: "Hello World",
			exp:   "hello_world",
		},
		"full stops": {
			input: "hello.world",
			exp:   "hello_world",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := cleanRepoName(test.input)

			if out != test.exp {
				log.Fatalf("Cleaned name incorrect. \nExpected %v \nGot %v", test.exp, out)
			}
		})
	}
}
