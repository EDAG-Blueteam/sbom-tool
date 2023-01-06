package structs

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var METADATA_FILES = map[string]string{
	"package.json":     "npm",    // node.js
	"requirements.txt": "pypi",   // python
	"pom.xml":          "maven",  // java(mvn)
	"conan.conf":       "conan",  // c++/conan
	"Cargo.toml":       "cargo",  // rust/cargo
	"java.config":      "maven",  // maven
	"maven.config":     "maven",  // duh
	"build.gradle":     "gradle", // java (gradle)
}

type Filesystem struct {
	Root string `json:"root"`
}

func NewFilesystem(root string) Filesystem {

	var filesystem Filesystem

	filesystem.SetRoot(root)

	return filesystem

}

func (filesystem *Filesystem) SetRoot(root string) {

	stat, err := os.Stat(root)

	if err == nil && stat.IsDir() {
		filesystem.Root = root
	}

}

func (filesystem *Filesystem) Scan() []ResultInfo {

	var result []ResultInfo

	err := filepath.WalkDir(filesystem.Root, func(path string, info fs.DirEntry, err error) error {

		// skip unused path
		match, _ := regexp.MatchString("\\.\\w+", info.Name())
		if info.IsDir() &&
			(info.Name() == "SBOMWorkingDir" || match) {
			return filepath.SkipDir
		}
		log.Printf("Visited: %s\n", path)

		if info.IsDir() == false {

			var name = info.Name()

			for key, val := range METADATA_FILES {
				if key == name {
					result = append(result, ResultInfo{
						Path: path,
						Type: val,
						Uuid: CreateProjectUuid(path),
					})
					break
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return result

}

func (filesystem *Filesystem) Read(file string) []byte {

	var result []byte

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	stat, err1 := os.Stat(filesystem.Root + file)

	if err1 == nil && stat.IsDir() == false {

		buffer, err2 := os.ReadFile(filesystem.Root + file)

		if err2 == nil {
			result = buffer
		}

	}

	return result

}
