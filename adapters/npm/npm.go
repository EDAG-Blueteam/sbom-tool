package npm

import (
	"fmt"
	"path/filepath"
	"sbom-tool/structs"
)

type NPM struct {
}

func (npm *NPM) Generate(file string) []byte {

	var shell = structs.NewShell(filepath.Dir(file))

	_, err := shell.Execute("npm", []string{"install", "--save-dev", "@cyclonedx/cyclonedx-npm"})

	fmt.Println(err)

	return nil
}

func IsPackage(file string) bool {

	var result bool = false

	return result

}

func (npm *NPM) BuildToolsExist() bool {
	//TODO implement me
	return true
}
