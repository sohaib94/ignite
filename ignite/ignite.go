package ignite

import (
	"log"

	"github.com/sohaib94/ignite/repo"
	"gopkg.in/yaml.v3"
)

type Ignite struct {
	Repo repo.Repo `yaml:"repo"`
	OutputParentPath string
}

type read func(string) ([]byte, error)

type IgniteFile struct {
	Path    string
	Reader  read
	content []byte
}

// TODO handle converting yaml file
func readFile(path string, readcb read) (a []byte, err error) {

	file, err := readcb(path)
	if err != nil {
		log.Printf("Provided ignite file %s doesn't exist. Error: %v \n", path, err)
		return nil, err
	}

	return file, nil
}

func (i Ignite) Handle(f *IgniteFile) error {

	var err error
	f.content, err = readFile(f.Path, f.Reader)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f.content, &i)
	if err  != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	i.Repo.Create(i.OutputParentPath)

	return nil
}
