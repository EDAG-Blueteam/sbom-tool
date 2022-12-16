package structs

import (
	"log"
	"os"
	"path/filepath"
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

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		// Exclude unused path
		// TODO exclude all file that start with "."
		if info.IsDir() &&
			(info.Name() == ".git" || info.Name() == "SBOMWorkingDir" || info.Name() == ".idea") {
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
