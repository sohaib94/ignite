package repo

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Meta struct {
	Host string `yaml:"host"`
	Org string  `yaml:"org"`
	Conn string `yaml:"conn"`
	Name string `yaml:"name"`
}

type Directory struct {
	Files []string `yaml:"files"`
	Directories map[string]Directory `yaml:"directories"`
}

type Repo struct {
	Meta Meta `yaml:"meta"`
	Root Directory `yaml:"root"`
}

func (r Repo) Create(path string) error {

	if path == "" {
		if cwd, err := os.Getwd(); err != nil {
			log.Println(err)
		} else {
			path = cwd
		}
	}

	// Create the local repo
	if err := r.CreateLocalRepo(path); err != nil {
		log.Println("Failed to create local repo.")
		log.Println("Exiting")
		return err
	}
	
	// Create the remote repo

	return nil
}

func (r Repo) CreateLocalRepo(path string) error {

	log.Println(r)

	pathParts := strings.Split(path, "/")

	// Create the root dir
	repoName := cleanRepoName(pathParts[len(pathParts)-1])

	// Create dirs that need creating
	if err := createDirectory(filepath.Join(path), r.Root); err != nil {
		log.Printf("Failed to create repo %v", repoName)
	}
	
	return nil
}

func cleanRepoName(name string) string {

	r := regexp.MustCompile("[. ]")

	return strings.ToLower(r.ReplaceAllString(name, "_"))
}

func createDirectory(path string, dir Directory) error {
	
	// Create Directory
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Printf("Unable to create directory %v", path)
		log.Printf("Error: %v", err)
		return err
	}
	
	// TODO Figure out cloning
	// TODO Clone anything that needs cloning

	// Create files in directory
	for _, f := range dir.Files {
		if _, err := os.Create(filepath.Join(path, f)); err != nil {
			log.Printf("Unable to create file %v", f)
			log.Printf("Error: %v", err)
			log.Println("Continuing to next file")
		}
	}

	// Create any sub-directories
	for k, d := range dir.Directories {
		if err := createDirectory(filepath.Join(path, k), d); err != nil {
			log.Println("Continuing to next dir")
		}
	}

	return nil
}
